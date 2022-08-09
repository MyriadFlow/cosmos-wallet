// Package cosmos_transfer provides function to transfer tokens
package cosmos_transfer

import (
	"context"
	"fmt"
	"strconv"

	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/env"
	apiBaseTendermint "github.com/cosmos/cosmos-sdk/api/cosmos/base/tendermint/v1beta1"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"google.golang.org/grpc"
)

// TransacParams defines transfer request containing from address, to address, private key
// denom and amount
type TransacParams struct {
	FromAddr sdk.AccAddress
	ToAddr   sdk.AccAddress
	PrivKey  *secp256k1.PrivKey
	Denom    string
	Amount   int64
}

// Transfer signs transfer request and returns tx hash or error if any
func Transfer(p *TransacParams) (string, error) {
	// Connect to gRPC server
	grpcServerUrl := env.MustGetEnv("NODE_GRPC_URL")
	grpcConn, err := grpc.Dial(
		grpcServerUrl,
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		err = fmt.Errorf("failed to dial grpc url %s: %w", grpcServerUrl, err)
		return "", err
	}
	defer grpcConn.Close()

	// Create new send transaction message
	trasactionMsg := banktypes.NewMsgSend(p.FromAddr, p.ToAddr, sdk.NewCoins(sdk.NewInt64Coin(p.Denom, p.Amount)))
	encCfg := simapp.MakeTestEncodingConfig()

	// Create transaction builder and set transaction message
	txBuilder := encCfg.TxConfig.NewTxBuilder()
	err = txBuilder.SetMsgs(trasactionMsg)
	if err != nil {
		err = fmt.Errorf("failed to set trasaction msg: %w", err)
		return "", err
	}

	// Get gas limit from environment
	gasLimit, err := strconv.ParseUint(env.MustGetEnv("GAS_LIMIT"), 10, 64)
	if err != nil {
		err = fmt.Errorf("failed to parse uint from env string for gas limit: %w", err)
		return "", err
	}
	txBuilder.SetGasLimit(gasLimit)

	// Create base tendermint clinet to query latest block
	baseTendermintClient := apiBaseTendermint.NewServiceClient(grpcConn)
	getLatestBlockRes, err := baseTendermintClient.GetLatestBlock(context.Background(), &apiBaseTendermint.GetLatestBlockRequest{})
	if err != nil {
		err = fmt.Errorf("failed to get latest block: %w", err)
		return "", err
	}

	// Set timeout height to latest+100
	timeOutHeight := getLatestBlockRes.Block.Header.Height + 100
	txBuilder.SetTimeoutHeight(uint64(timeOutHeight))

	txClient := tx.NewServiceClient(grpcConn)

	txBytes, err := processSignature(txBuilder, p, encCfg, grpcConn)
	if err != nil {
		err = fmt.Errorf("failed to process signature and get transaction bytes: %w", err)
		return "", err
	}

	simulateRes, err := txClient.Simulate(context.Background(), &tx.SimulateRequest{
		TxBytes: txBytes,
	})
	if err != nil {
		err = fmt.Errorf("failed to get gas fee using simulate transaction: %w", err)
		return "", err
	}
	txBuilder.SetFeeAmount(sdk.NewCoins(sdk.NewInt64Coin(env.MustGetEnv("SMALLEST_DENOM"), int64(simulateRes.GasInfo.GasUsed))))

	txBytes, err = processSignature(txBuilder, p, encCfg, grpcConn)
	if err != nil {
		err = fmt.Errorf("failed to process signature and get transaction bytes: %w", err)
		return "", err
	}
	// Broadcast the tx via gRPC. We use txClient for the Protobuf Tx
	// service.
	// We then call the BroadcastTx method on this client.
	grpcRes, err := txClient.BroadcastTx(
		context.Background(),
		&tx.BroadcastTxRequest{
			Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes, // Proto-binary of the signed transaction, see previous step.
		},
	)
	if err != nil {
		err = fmt.Errorf("failed to broadcast tx: %w", err)
		return "", err
	}
	if grpcRes.TxResponse.Code != 0 {
		err = fmt.Errorf("transaction failed with status code %d: %s", grpcRes.TxResponse.Code, grpcRes.TxResponse.RawLog)
		return "", err
	}
	return grpcRes.TxResponse.TxHash, nil
}

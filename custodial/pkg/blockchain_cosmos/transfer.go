package blockchain_cosmos

import (
	"context"
	"fmt"
	"strconv"

	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/env"
	apiAuth "github.com/cosmos/cosmos-sdk/api/cosmos/auth/v1beta1"
	apiBaseTendermint "github.com/cosmos/cosmos-sdk/api/cosmos/base/tendermint/v1beta1"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"google.golang.org/grpc"
)

type TransferParams struct {
	FromAddr sdk.AccAddress
	ToAddr   sdk.AccAddress
	PrivKey  *secp256k1.PrivKey
	Denom    string
	Amount   int64
}

type txHash string

// Signs transfer request and returns tx hash or error if any
func Transfer(p *TransferParams) (string, error) {
	grpcServerUrl := env.MustGetEnv("NODE_GRPC_URL")
	grpcConn, err := grpc.Dial(
		grpcServerUrl,       // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		err = fmt.Errorf("failed to dial grpc url %s: %w", grpcServerUrl, err)
		return "", err
	}
	defer grpcConn.Close()
	queryClient := apiAuth.NewQueryClient(grpcConn)
	accountQueryRes, err := queryClient.Account(context.Background(), &apiAuth.QueryAccountRequest{
		Address: p.FromAddr.String(),
	})
	if err != nil {
		err = fmt.Errorf("failed to create auth query client: %w", err)
		return "", err
	}
	var baseAccount apiAuth.BaseAccount
	err = accountQueryRes.GetAccount().UnmarshalTo(&baseAccount)
	if err != nil {
		err = fmt.Errorf("failed to get account details: %w", err)
		return "", err
	}
	trasactionMsg := banktypes.NewMsgSend(p.FromAddr, p.ToAddr, sdk.NewCoins(sdk.NewInt64Coin(p.Denom, p.Amount)))
	encCfg := simapp.MakeTestEncodingConfig()
	txBuilder := encCfg.TxConfig.NewTxBuilder()
	err = txBuilder.SetMsgs(trasactionMsg)
	if err != nil {
		err = fmt.Errorf("failed to set trasaction msg: %w", err)
		return "", err
	}

	gasLimit, err := strconv.ParseUint(env.MustGetEnv("GAS_LIMIT"), 10, 64)
	if err != nil {
		err = fmt.Errorf("failed to parse uint from env string for gas limit: %w", err)
		return "", err
	}
	txBuilder.SetGasLimit(gasLimit)
	txBuilder.SetFeeAmount(sdk.NewCoins(sdk.NewInt64Coin(env.MustGetEnv("SMALLEST_DENOM"), 1)))

	baseTendermintClient := apiBaseTendermint.NewServiceClient(grpcConn)
	getLatestBlockRes, err := baseTendermintClient.GetLatestBlock(context.Background(), &apiBaseTendermint.GetLatestBlockRequest{})
	if err != nil {
		err = fmt.Errorf("failed to get latest block: %w", err)
		return "", err
	}

	timeOutHeight := getLatestBlockRes.Block.Header.Height + 100
	txBuilder.SetTimeoutHeight(uint64(timeOutHeight))
	sigV2 := signing.SignatureV2{
		PubKey: p.PrivKey.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: baseAccount.Sequence,
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		err = fmt.Errorf("failed to set initial signatures: %w", err)
		return "", err
	}

	signerData := xauthsigning.SignerData{
		ChainID:       env.MustGetEnv("CHAIN_ID"),
		AccountNumber: baseAccount.AccountNumber,
	}
	sigV2, err = clienttx.SignWithPrivKey(
		encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
		txBuilder, p.PrivKey, encCfg.TxConfig, baseAccount.Sequence)
	if err != nil {
		err = fmt.Errorf("failed to sign transaction: %w", err)
		return "", err
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		err = fmt.Errorf("failed to set final signatures: %w", err)
		return "", err
	}

	// Encode tx.
	txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		err = fmt.Errorf("failed to get transaction bytes: %w", err)
		return "", err
	}

	// Broadcast the tx via gRPC. We create a new client for the Protobuf Tx
	// service.
	txClient := tx.NewServiceClient(grpcConn)
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
		err = fmt.Errorf("transaction failed: %s", grpcRes.TxResponse.RawLog)
		return "", err
	}
	return grpcRes.TxResponse.TxHash, nil
}

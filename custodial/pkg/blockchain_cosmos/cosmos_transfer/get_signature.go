package cosmos_transfer

import (
	"fmt"

	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/env"
	"github.com/cosmos/cosmos-sdk/client"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"google.golang.org/grpc"
)

// processSignature processes the signature in transaction builder by signing it in two rounds
func processSignature(txBuilder client.TxBuilder, p *TransacParams, encCfg params.EncodingConfig, grpcConn *grpc.ClientConn) (txBytes []byte, e error) {

	// Get BaseAccount for account number and sequence
	baseAccount, err := getAccountDetails(p.FromAddr, grpcConn)
	if err != nil {
		err = fmt.Errorf("failed to get base account details: %w", err)
		return nil, err
	}
	// First round: we gather all the signer info. We use the "set empty
	// signature" hack to do that.
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
		return nil, err
	}

	// Second round: all signer infos are set, so signer can sign.
	signerData := xauthsigning.SignerData{
		ChainID:       env.MustGetEnv("CHAIN_ID"),
		AccountNumber: baseAccount.AccountNumber,
	}
	sigV2, err = clienttx.SignWithPrivKey(
		encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
		txBuilder, p.PrivKey, encCfg.TxConfig, baseAccount.Sequence)
	if err != nil {
		err = fmt.Errorf("failed to sign transaction: %w", err)
		return nil, err
	}
	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		err = fmt.Errorf("failed to sign transaction: %w", err)
		return nil, err
	}

	// Encode tx and return
	return encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
}

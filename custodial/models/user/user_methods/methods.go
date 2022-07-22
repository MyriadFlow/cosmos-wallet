package usermethods

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/MyriadFlow/cosmos-wallet/custodial/models/user"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/blockchain_cosmos"
	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/google/uuid"
)

// Creates and stores mnemonic and returns public Key and user Id
func Create() (*cryptotypes.PubKey, string, error) {
	mnemonic, err := blockchain_cosmos.GenerateMnemonic()
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate mnemonic: %w", err)
	}
	uid := uuid.NewString()

	err = user.Add(uid, *mnemonic)
	if err != nil {
		return nil, "", fmt.Errorf("failed to add user into database: %w", err)
	}

	privKey, err := blockchain_cosmos.GetWallet(*mnemonic)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get privateKey for mnemonic: %w", err)
	}
	pubKey := privKey.PubKey()
	return &pubKey, uid, nil
}

func Transfer(uid string, from string, to string, amount int64) error {
	user, err := user.Get(uid)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	fromAddr, err := sdk.AccAddressFromBech32(from)
	if err != nil {
		return fmt.Errorf("failed to get from address: %w", err)
	}

	toAddr, err := sdk.AccAddressFromBech32(to)
	if err != nil {
		return fmt.Errorf("failed to get to address: %w", err)
	}
	privKey, err := blockchain_cosmos.GetWallet(user.Mnemonic)
	if err != nil {
		return fmt.Errorf("failed to get private key: %w", err)
	}

	trasactionMsg := banktypes.NewMsgSend(fromAddr, toAddr, sdk.NewCoins(sdk.NewInt64Coin("uatom", amount)))

	encCfg := simapp.MakeTestEncodingConfig()
	txBuilder := encCfg.TxConfig.NewTxBuilder()
	err = txBuilder.SetMsgs(trasactionMsg)
	if err != nil {
		return err
	}

	//TODO: check gas limit
	txBuilder.SetGasLimit(210000)
	txBuilder.SetFeeAmount(sdk.NewCoins(sdk.NewInt64Coin("uatom", 1)))
	txBuilder.SetTimeoutHeight(13269949)
	sigV2 := signing.SignatureV2{
		PubKey: privKey.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: 1,
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return err
	}

	//TODO get account no from network
	signerData := xauthsigning.SignerData{
		ChainID:       "theta-testnet-001",
		AccountNumber: 697738,
	}
	//TODO get account sequence from network
	sigV2, err = tx.SignWithPrivKey(
		encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
		txBuilder, privKey, encCfg.TxConfig, 0)
	if err != nil {
		return err
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return err
	}

	// Encode tx.
	txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return err
	}
	jsonBody := fmt.Sprintf(`{
	"jsonrpc": "2.0",
	"id": "1",
	"method": "check_tx",
	"params": {
	  "tx": "%s"
	}
  }`, base64.StdEncoding.EncodeToString(txBytes))

	resp, err := http.Post("https://rpc.sentry-01.theta-testnet.polypore.xyz", "application/json", strings.NewReader(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to send transaction: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("status code not 200")
	}
	//TODO handle error _ > err
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	logo.Info(string(bodyBytes))
	// txJSON := string(txJSONBytes)
	// logo.Info(base64.StdEncoding.EncodeToString(txBytes))
	//TODO handle json rpc error, for e.g. with status codes or missing hash

	return errors.New("not implemented")
}

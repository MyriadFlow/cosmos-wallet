![myriad-flow-x-assetmantle_auto_x2](https://user-images.githubusercontent.com/51229945/181774432-369d1722-18dc-46dd-a59b-89f76e3c1fe2.jpg)
# Cosmos Network

Wallet modules for Cosmos Network - AssetMantle
Different modules to help you integrate many features including
off-chain verification, transfer of tokens, custodial wallet.

[![Go Report Card](https://goreportcard.com/badge/github.com/MyriadFlow/cosmos-wallet)](https://goreportcard.com/report/github.com/MyriadFlow/cosmos-wallet)

## Running up

- Install dependencies - `go get ./...`
- Set up `.env` file for the modules you want to start
- Run `go run main.go`

# Design decisions

- `o` prefix is used where there might be conflict, for e.g. package names like errorso, logo

# Modules

## Custodial

This module consist of methods which are used to manage custodial wallet with postgres data base
It contains methods such as `GenerateMnemonic()`, `GetPrivKey(mnemonic string)` and `Transfer(p *TransferParams)`

## Sign Auth

This module consist of methods which are used to manage the off-chain verification of user using signatures mostly signed with `ADR036` for example `keplr.signArbitrary`

This modules contains integration with postgres so verfication can be completed in two steps which are

1. Request for challenge

- The user passes his wallet address and gets flow ID and EULA to sign

2. Sign the challenge

- The user signs the challenge by using `signArbitrary` such that message will be EULA+Flow ID
- The signature obtained is then send along with the public key

After this steps if the signature is valid then user receives the PASETO token.

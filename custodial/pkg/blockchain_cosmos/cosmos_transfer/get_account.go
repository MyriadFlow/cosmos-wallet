package cosmos_transfer

import (
	"context"
	"fmt"

	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/errorso"
	apiAuth "github.com/cosmos/cosmos-sdk/api/cosmos/auth/v1beta1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
)

// getAccountDetails returns BaseAccount for given wallet address
func getAccountDetails(walletAddress sdk.Address, grpcConn *grpc.ClientConn) (*apiAuth.BaseAccount, error) {
	// Initiat new query client to query account details
	queryClient := apiAuth.NewQueryClient(grpcConn)
	accountQueryRes, err := queryClient.Account(context.Background(), &apiAuth.QueryAccountRequest{
		Address: walletAddress.String(),
	})
	if err != nil {
		if grpcStatus.Code(err) == codes.NotFound {
			return nil, errorso.AccountNotFound
		}
		err = fmt.Errorf("failed to create auth query client: %w", err)
		return nil, err
	}

	var baseAccount apiAuth.BaseAccount
	// Unmarshal proto bytes into BaseAccount
	err = accountQueryRes.GetAccount().UnmarshalTo(&baseAccount)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal proto bytes into BaseAccount: %w", err)
		return nil, err
	}

	return &baseAccount, nil
}

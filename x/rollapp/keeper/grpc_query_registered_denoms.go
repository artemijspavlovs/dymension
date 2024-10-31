package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dymensionxyz/dymension/v3/x/rollapp/types"
)

func (k Keeper) RegisteredDenoms(c context.Context, req *types.QueryRegisteredDenomsRequest) (*types.QueryRegisteredDenomsResponse, error) {
	if req == nil || req.RollappId == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	denoms, err := k.GetAllRegisteredDenoms(sdk.UnwrapSDKContext(c), req.RollappId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryRegisteredDenomsResponse{
		Denoms: denoms,
	}, nil
}
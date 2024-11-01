package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/dymensionxyz/dymension/v3/x/sequencer/types"
	"github.com/dymensionxyz/gerr-cosmos/gerrc"
)

func (k Keeper) SequencersByRollapp(c context.Context, req *types.QueryGetSequencersByRollappRequest) (*types.QueryGetSequencersByRollappResponse, error) {
	if req == nil {
		return nil, gerrc.ErrInvalidArgument
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryGetSequencersByRollappResponse{
		Sequencers: k.RollappSequencers(ctx, req.RollappId),
	}, nil
}

func (k Keeper) SequencersByRollappByStatus(c context.Context, req *types.QueryGetSequencersByRollappByStatusRequest) (*types.QueryGetSequencersByRollappByStatusResponse, error) {
	if req == nil {
		return nil, gerrc.ErrInvalidArgument
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryGetSequencersByRollappByStatusResponse{
		Sequencers: k.RollappSequencersByStatus(ctx, req.RollappId, req.Status),
	}, nil
}

func (k Keeper) GetProposerByRollapp(c context.Context, req *types.QueryGetProposerByRollappRequest) (*types.QueryGetProposerByRollappResponse, error) {
	if req == nil {
		return nil, gerrc.ErrInvalidArgument
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryGetProposerByRollappResponse{
		ProposerAddr: k.GetProposer(ctx, req.RollappId).Address,
	}, nil
}

func (k Keeper) GetNextProposerByRollapp(c context.Context, req *types.QueryGetNextProposerByRollappRequest) (*types.QueryGetNextProposerByRollappResponse, error) {
	if req == nil {
		return nil, gerrc.ErrInvalidArgument
	}
	ctx := sdk.UnwrapSDKContext(c)

	successor := k.GetSuccessor(ctx, req.RollappId)
	inProgress := k.awaitingLastProposerBlock(ctx, req.RollappId)

	return &types.QueryGetNextProposerByRollappResponse{
		NextProposerAddr:   successor.Address,
		RotationInProgress: inProgress,
	}, nil
}

func (k Keeper) Proposers(c context.Context, req *types.QueryProposersRequest) (*types.QueryProposersResponse, error) {
	if req == nil {
		return nil, gerrc.ErrInvalidArgument
	}
	ctx := sdk.UnwrapSDKContext(c)

	var proposers []types.Sequencer

	store := ctx.KVStore(k.storeKey)
	sequencerStore := prefix.NewStore(store, types.ProposerByRollappKey(""))

	pageRes, err := query.Paginate(sequencerStore, req.Pagination, func(key []byte, value []byte) error {
		proposer, err := k.RealSequencer(ctx, string(value))
		if err != nil {
			return err
		}
		proposers = append(proposers, proposer)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &types.QueryProposersResponse{Proposers: proposers, Pagination: pageRes}, nil
}

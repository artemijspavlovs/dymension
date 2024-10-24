package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/dymension/v3/x/sequencer/types"
)

func (k Keeper) SentinelSequencer(ctx sdk.Context, rollapp string) types.Sequencer {
	s := k.NewSequencer(ctx, rollapp)
	s.Status = types.Bonded
	s.Address = types.SentinelSeqAddr
	return *s
}

func (k Keeper) NewSequencer(ctx sdk.Context, rollapp string) *types.Sequencer {
	return &types.Sequencer{
		RollappId: rollapp,
		Tokens:    sdk.NewCoins(sdk.NewCoin(k.bondDenom(ctx), sdk.NewInt(0))),
	}
}

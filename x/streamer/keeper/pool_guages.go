package keeper

import (
	"time"

	"github.com/dymensionxyz/dymension/x/streamer/types"
	gammtypes "github.com/osmosis-labs/osmosis/v15/x/gamm/types"
	lockuptypes "github.com/osmosis-labs/osmosis/v15/x/lockup/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	defaultLockDuration = 60 * time.Second
)

func (k Keeper) CreatePoolGauge(ctx sdk.Context, poolId uint64) error {
	_, err := k.ik.CreateGauge(
		ctx,
		true,
		k.ak.GetModuleAddress(types.ModuleName),
		sdk.Coins{},
		lockuptypes.QueryCondition{
			LockQueryType: lockuptypes.ByDuration,
			Denom:         gammtypes.GetPoolShareDenom(poolId),
			Duration:      defaultLockDuration,
			Timestamp:     time.Time{},
		},
		ctx.BlockTime(),
		1,
	)
	if err != nil {
		return err
	}

	return nil
}

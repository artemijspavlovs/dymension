package types

import (
	time "time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewPlan(id uint64, rollappId string, allocation sdk.Coin, curve BondingCurve, start time.Time, end time.Time) Plan {
	plan := Plan{
		Id:               id,
		RollappId:        rollappId,
		ModuleAccAddress: "",
		TotalAllocation:  allocation,
		BondingCurve:     curve,
		StartTime:        start,
		EndTime:          end,
		SoldAmt:          math.ZeroInt(),
		ClaimedAmt:       math.ZeroInt(),
	}
	return plan
}

func (p Plan) IsSettled() bool {
	return p.SettledDenom != ""
}
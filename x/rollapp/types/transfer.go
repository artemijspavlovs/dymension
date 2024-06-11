package types

import (
	"fmt"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
)

type TransferData struct {
	transfertypes.FungibleTokenPacketData
	// RollappID will be the empty string if the packet does not pertain to a registered rollapp
	RollappID string
}

// IsRollapp returns whether the transfer came from a rollapp or was sent to a rollapp
func (d TransferData) IsRollapp() bool {
	return d.RollappID != ""
}

// MustAmountInt returns the int amount. Should call validateBasic first!
func (d TransferData) MustAmountInt() math.Int {
	x, ok := sdk.NewIntFromString(d.Amount)
	if !ok {
		panic(fmt.Sprintf("parse transfer amount to Int: %s", d.Amount))
	}
	return x
}
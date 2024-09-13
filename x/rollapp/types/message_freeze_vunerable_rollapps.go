package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dymensionxyz/gerr-cosmos/gerrc"
)

var _ sdk.Msg = new(MsgFreezeVulnerableRollapps)

func (m MsgFreezeVulnerableRollapps) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf(
			"authority '%s' must be a valid bech32 address: %s",
			m.Authority, err.Error(),
		)
	}

	if len(m.DrsVersions) == 0 {
		return gerrc.ErrInvalidArgument.Wrap("at least one DRS version is required")
	}

	return nil
}

func (m MsgFreezeVulnerableRollapps) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{signer}
}

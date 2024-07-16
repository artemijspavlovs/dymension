package types

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dymensionxyz/dymension/v3/testutil/sample"
)

func NewRollapp(
	creator,
	rollappId,
	initSequencerAddress,
	bech32Prefix,
	genesisChecksum,
	alias string,
	metadata *RollappMetadata,
	transfersEnabled bool,
) Rollapp {
	return Rollapp{
		RollappId:               rollappId,
		Creator:                 creator,
		InitialSequencerAddress: initSequencerAddress,
		GenesisChecksum:         genesisChecksum,
		Bech32Prefix:            bech32Prefix,
		Alias:                   alias,
		Metadata:                metadata,
		GenesisState: RollappGenesisState{
			TransfersEnabled: transfersEnabled,
		},
	}
}

const (
	maxAliasLength       = 64
	maxDescriptionLength = 256
	maxDataURILength     = 25 * 1024 // 25KB
	dataURIPattern       = `^data:(?P<mimeType>[\w/]+);base64,(?P<data>[A-Za-z0-9+/=]+)$`
)

func (r Rollapp) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(r.Creator)
	if err != nil {
		return errorsmod.Wrap(ErrInvalidCreatorAddress, err.Error())
	}

	// validate rollappId
	_, err = NewChainID(r.RollappId)
	if err != nil {
		return err
	}

	_, err = sdk.AccAddressFromBech32(r.InitialSequencerAddress)
	if err != nil {
		return errorsmod.Wrap(ErrInvalidInitialSequencerAddress, err.Error())
	}

	if err = validateBech32Prefix(r.Bech32Prefix); err != nil {
		return errorsmod.Wrap(ErrInvalidBech32Prefix, err.Error())
	}

	if r.GenesisChecksum == "" {
		return errorsmod.Wrap(ErrEmptyGenesisChecksum, "GenesisChecksum")
	}

	if l := len(r.Alias); l == 0 || l > maxAliasLength {
		return ErrInvalidAlias
	}

	if err = validateMetadata(r.Metadata); err != nil {
		return errorsmod.Wrap(ErrInvalidMetadata, err.Error())
	}

	return nil
}

func validateBech32Prefix(prefix string) error {
	bechAddr, err := sdk.Bech32ifyAddressBytes(prefix, sample.Acc())
	if err != nil {
		return err
	}

	bAddr, err := sdk.GetFromBech32(bechAddr, prefix)
	if err != nil {
		return err
	}

	if err = sdk.VerifyAddressFormat(bAddr); err != nil {
		return err
	}
	return nil
}

func validateMetadata(metadata *RollappMetadata) error {
	if metadata == nil {
		return nil
	}

	if _, err := url.Parse(metadata.Website); err != nil {
		return errorsmod.Wrap(ErrInvalidWebsiteURL, err.Error())
	}

	if len(metadata.Description) > maxDescriptionLength {
		return ErrInvalidDescription
	}

	if err := validateBaseURI(metadata.LogoDataUri); err != nil {
		return errorsmod.Wrap(ErrInvalidLogoURI, err.Error())
	}

	if err := validateBaseURI(metadata.TokenLogoUri); err != nil {
		return errorsmod.Wrap(ErrInvalidTokenLogoURI, err.Error())
	}

	return nil
}

func validateBaseURI(dataURI string) error {
	if dataURI == "" {
		return nil
	}

	if len(dataURI) > maxDataURILength {
		return fmt.Errorf("data URI exceeds maximum length")
	}

	matched, _ := regexp.MatchString(dataURIPattern, dataURI)
	if !matched {
		return fmt.Errorf("invalid data URI format")
	}

	commaIndex := strings.Index(dataURI, ",")
	if commaIndex == -1 {
		return fmt.Errorf("no comma found in data URI")
	}
	base64Data := dataURI[commaIndex+1:]

	_, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return fmt.Errorf("invalid base64 data: %w", err)
	}

	return nil
}

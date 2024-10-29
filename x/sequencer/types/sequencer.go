package types

import (
	"time"

	errorsmod "cosmossdk.io/errors"
	tmprotocrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	cometbfttypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/gerr-cosmos/gerrc"
	"github.com/dymensionxyz/sdk-utils/utils/uevent"
)

const (
	SentinelSeqAddr = "sentinel"
)

func NewTestSequencer(
	pk cryptotypes.PubKey,
) Sequencer {
	pkAny, err := codectypes.NewAnyWithValue(pk)
	if err != nil {
		panic(err)
	}
	return Sequencer{
		Address:      sdk.AccAddress(pk.Address()).String(),
		DymintPubKey: pkAny,
	}

}

// ValidateBasic performs basic validation of the sequencer object
func (seq Sequencer) ValidateBasic() error {
	if seq.Tokens.Len() != 1 {
		return gerrc.ErrInvalidArgument.Wrap("expect one coin")
	}
	return nil
}

func (seq *Sequencer) SetOptedIn(ctx sdk.Context, x bool) error {
	if err := uevent.EmitTypedEvent(ctx, &EventOptInStatusChange{
		seq.RollappId,
		seq.Address,
		seq.OptedIn,
		x,
	}); err != nil {
		return err
	}
	seq.OptedIn = x
	return nil
}

func (seq Sequencer) Sentinel() bool {
	return seq.Address == SentinelSeqAddr
}

func (seq Sequencer) Bonded() bool {
	return seq.Status == Bonded
}

func (seq Sequencer) IsPotentialProposer() bool {
	return seq.Bonded() && seq.OptedIn
}

func (seq Sequencer) TokensCoin() sdk.Coin {
	return seq.Tokens[0]
}

func (seq Sequencer) SetTokensCoin(c sdk.Coin) {
	seq.Tokens[0] = c
}

func (seq Sequencer) AccAddr() sdk.AccAddress {
	return sdk.MustAccAddressFromBech32(seq.Address)
}

func (seq Sequencer) NoticeInProgress(now time.Time) bool {
	return seq.NoticeStarted() && !seq.NoticeElapsed(now)
}

func (seq Sequencer) NoticeElapsed(now time.Time) bool {
	return seq.NoticeStarted() && !now.Before(seq.NoticePeriodTime)
}

func (seq Sequencer) NoticeStarted() bool {
	return seq.NoticePeriodTime != time.Time{}
}

// MustValsetHash : intended for tests
func (seq Sequencer) MustValsetHash() []byte {
	x, err := seq.ValsetHash()
	if err != nil {
		panic(err)
	}
	return x
}

func (seq Sequencer) ValsetHash() ([]byte, error) {
	pubKey, err := seq.PubKey()
	if err != nil {
		return nil, errorsmod.Wrap(err, "pub key")
	}

	// convert the pubkey to tmPubKey
	tmPubKey, err := cryptocodec.ToTmPubKeyInterface(pubKey)
	if err != nil {
		return nil, errorsmod.Wrap(err, "tm pub key")
	}

	val := cometbfttypes.NewValidator(tmPubKey, 1)

	valset, err := cometbfttypes.ValidatorSetFromExistingValidators([]*cometbfttypes.Validator{
		val,
	})
	if err != nil {
		return nil, errorsmod.Wrap(err, "validator set")
	}
	return valset.Hash(), nil
}

// GetDymintPubKeyHash returns the hash of the sequencer
// as expected to be written on the rollapp ibc client headers
func (seq Sequencer) GetDymintPubKeyHash() ([]byte, error) {
	pubKey, err := seq.PubKey()
	if err != nil {
		return nil, err
	}

	// convert the pubkey to tmPubKey
	tmPubKey, err := cryptocodec.ToTmPubKeyInterface(pubKey)
	if err != nil {
		return nil, err
	}
	// Create a new tmValidator with fixed voting power of 1
	// TODO: Make sure the voting power is a param coming from hub and
	// not being set independently in dymint and hub
	tmValidator := cometbfttypes.NewValidator(tmPubKey, 1)
	tmValidatorSet := cometbfttypes.NewValidatorSet([]*cometbfttypes.Validator{tmValidator})
	return tmValidatorSet.Hash(), nil
}

// CometPubKey returns the bytes of the sequencer's dymint pubkey
func (seq Sequencer) CometPubKey() (tmprotocrypto.PublicKey, error) {
	pubKey, err := seq.PubKey()
	if err != nil {
		return tmprotocrypto.PublicKey{}, err
	}

	// convert the pubkey to tmPubKey
	tmPubKey, err := cryptocodec.ToTmProtoPublicKey(pubKey)
	return tmPubKey, err
}

func (seq Sequencer) PubKey() (cryptotypes.PubKey, error) {
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(interfaceRegistry)
	protoCodec := codec.NewProtoCodec(interfaceRegistry)

	var pubKey cryptotypes.PubKey
	err := protoCodec.UnpackAny(seq.DymintPubKey, &pubKey)
	return pubKey, err
}

// MustPubKey is intended for tests
func (seq Sequencer) MustPubKey() cryptotypes.PubKey {
	ret, err := seq.PubKey()
	if err != nil {
		panic(err)
	}
	return ret
}

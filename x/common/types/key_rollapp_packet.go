package types

import (
	"encoding/binary"
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

const (
	// KeySeparator defines the separator for keys
	KeySeparator = "/"
)

var (
	// PendingRollappPacketKeyPrefix is the prefix for pending rollapp packets
	PendingRollappPacketKeyPrefix = []byte{0x00, 0x01}
	// FinalizedRollappPacketKeyPrefix is the prefix for finalized rollapp packets
	FinalizedRollappPacketKeyPrefix = []byte{0x00, 0x02}
	// RevertedRollappPacketKeyPrefix is the prefix for reverted rollapp packets
	RevertedRollappPacketKeyPrefix = []byte{0x00, 0x03}
)

// RollappPacketKey constructs a key for a specific RollappPacket
func RollappPacketKey(rollappPacket *RollappPacket) []byte {
	rollappIdBytes := []byte(rollappPacket.RollappId)
	// Get the relevant key prefix based on the packet status
	statusBytes := mustGetStatusBytes(rollappPacket.Status)
	// Build the key bytes repr. Convert each uint64 to big endian bytes to ensure lexicographic ordering.
	keySeparatorBytes := []byte(KeySeparator)
	proofHeightBytes := sdk.Uint64ToBigEndian(rollappPacket.ProofHeight)
	// Build the packetUID from the destination channel and sequence number.
	packetSequenceBytes := sdk.Uint64ToBigEndian(rollappPacket.Packet.Sequence)
	packetDestinationChannelBytes := []byte(rollappPacket.Packet.DestinationChannel)
	packetUIDBytes := append(packetDestinationChannelBytes, packetSequenceBytes...)

	// Concatenate the byte slices directly.
	// rollappID/status/proofHeight/packetUID
	result := append(rollappIdBytes, keySeparatorBytes...)
	result = append(result, statusBytes...)
	result = append(result, keySeparatorBytes...)
	result = append(result, proofHeightBytes...)
	result = append(result, keySeparatorBytes...)
	result = append(result, packetUIDBytes...)

	return result
}

func RollappIDAndStatusPacketPrefix(rollappID string, status Status) []byte {
	rollappIdBytes := []byte(rollappID)
	statusPrefix := mustGetStatusBytes(status)
	keySeparatorBytes := []byte(KeySeparator)
	result := append(rollappIdBytes, keySeparatorBytes...)
	result = append(result, statusPrefix...)
	return result
}

// mustGetStatusBytes returns the byte representation of the status
func mustGetStatusBytes(status Status) []byte {
	switch status {
	case Status_PENDING:
		return PendingRollappPacketKeyPrefix
	case Status_FINALIZED:
		return FinalizedRollappPacketKeyPrefix
	case Status_REVERTED:
		return RevertedRollappPacketKeyPrefix
	default:
		panic(fmt.Sprintf("invalid packet status: %s", status))
	}
}

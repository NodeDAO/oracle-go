package chaintime

import (
	"time"

	"github.com/attestantio/go-eth2-client/spec/phase0"
)

// Service provides a number of functions for calculating chain-related times.
//
// nolint: interface bloat
type Service interface {
	// GenesisTime provides the time of the chain's genesis.
	GenesisTime() time.Time
	// SlotsPerEpoch provides the number of slots in the chain's epoch.
	SlotsPerEpoch() uint64
	// SlotDuration provides the duration of the chain's slot.
	SlotDuration() time.Duration

	// StartOfSlot provides the time at which a given slot starts.
	StartOfSlot(slot phase0.Slot) time.Time
	// StartOfEpoch provides the time at which a given epoch starts.
	StartOfEpoch(epoch phase0.Epoch) time.Time
	// CurrentSlot provides the current slot.
	CurrentSlot() phase0.Slot
	// CurrentEpoch provides the current epoch.
	CurrentEpoch() phase0.Epoch
	// CurrentSyncCommitteePeriod provides the current sync committee period.
	CurrentSyncCommitteePeriod() uint64
	// SlotToEpoch provides the epoch of the given slot.
	SlotToEpoch(slot phase0.Slot) phase0.Epoch
	// EpochToSlot provides the slot of the given epoch.
	EpochToSlot(epoch phase0.Epoch) phase0.Slot
	// SlotToSyncCommitteePeriod provides the sync committee period of the given slot.
	SlotToSyncCommitteePeriod(slot phase0.Slot) uint64
	// FirstSlotOfEpoch provides the first slot of the given epoch.
	FirstSlotOfEpoch(epoch phase0.Epoch) phase0.Slot
	// TimestampToSlot provides the slot of the given timestamp.
	TimestampToSlot(timestamp time.Time) phase0.Slot
	// TimestampToEpoch provides the epoch of the given timestamp.
	TimestampToEpoch(timestamp time.Time) phase0.Epoch
	// FirstEpochOfSyncPeriod provides the first epoch of the given sync period.
	FirstEpochOfSyncPeriod(period uint64) phase0.Epoch
	// AltairInitialEpoch provides the epoch at which the Altair hard fork takes place.
	AltairInitialEpoch() phase0.Epoch
	// AltairInitialSyncCommitteePeriod provides the sync committee period in which the Altair hard fork takes place.
	AltairInitialSyncCommitteePeriod() uint64
	// CapellaInitialEpoch provides the epoch at which the Capella hard fork takes place.
	CapellaInitialEpoch() phase0.Epoch
}

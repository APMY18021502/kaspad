package consensusstatemanager

import (
	"github.com/kaspanet/kaspad/domain/consensus/database"
	"github.com/kaspanet/kaspad/domain/consensus/model"
	"github.com/kaspanet/kaspad/domain/consensus/model/externalapi"
	"github.com/kaspanet/kaspad/domain/dagconfig"
)

// consensusStateManager manages the node's consensus state
type consensusStateManager struct {
	dagParams *dagconfig.Params

	databaseContext     *database.DomainDBContext
	consensusStateStore model.ConsensusStateStore
	multisetStore       model.MultisetStore
	utxoDiffStore       model.UTXODiffStore
	blockStore          model.BlockStore
	ghostdagManager     model.GHOSTDAGManager
	acceptanceManager   model.AcceptanceManager
}

// New instantiates a new ConsensusStateManager
func New(
	databaseContext *database.DomainDBContext,
	dagParams *dagconfig.Params,
	consensusStateStore model.ConsensusStateStore,
	multisetStore model.MultisetStore,
	utxoDiffStore model.UTXODiffStore,
	blockStore model.BlockStore,
	ghostdagManager model.GHOSTDAGManager,
	acceptanceManager model.AcceptanceManager) model.ConsensusStateManager {

	return &consensusStateManager{
		dagParams: dagParams,

		databaseContext:     databaseContext,
		consensusStateStore: consensusStateStore,
		multisetStore:       multisetStore,
		utxoDiffStore:       utxoDiffStore,
		blockStore:          blockStore,
		ghostdagManager:     ghostdagManager,
		acceptanceManager:   acceptanceManager,
	}
}

// UTXOByOutpoint returns a UTXOEntry matching the given outpoint
func (csm *consensusStateManager) UTXOByOutpoint(outpoint *externalapi.DomainOutpoint) (*externalapi.UTXOEntry, error) {
	return nil, nil
}

// CalculateConsensusStateChanges returns a set of changes that must occur in order
// to transition the current consensus state into the one including the given block
func (csm *consensusStateManager) CalculateConsensusStateChanges(block *externalapi.DomainBlock, isDisqualified bool) (
	stateChanges *model.ConsensusStateChanges, utxoDiffChanges *model.UTXODiffChanges,
	virtualGHOSTDAGData *model.BlockGHOSTDAGData, err error) {

	return nil, nil, nil, nil
}

// VirtualData returns the medianTime and blueScore of the current virtual block
func (csm *consensusStateManager) VirtualData() (medianTime int64, blueScore uint64, err error) {
	return 0, 0, nil
}

// RestoreUTXOSet calculates and returns the UTXOSet of the given blockHash
func (csm *consensusStateManager) RestorePastUTXOSet(blockHash *externalapi.DomainHash) (model.ReadOnlyUTXOSet, error) {
	return nil, nil
}

// RestoreDiffFromVirtual restores the diff between the given virtualDiffParentHash
// and the virtual
func (csm *consensusStateManager) RestoreDiffFromVirtual(utxoDiff *model.UTXODiff, virtualDiffParentHash *externalapi.DomainHash) (*model.UTXODiff, error) {
	panic("implement me")
}

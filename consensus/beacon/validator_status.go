// description:
// @author renshiwei
// Date: 2023/6/30

package beacon

import (
	"context"
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/NodeDAO/oracle-go/eth1"
	consensusApi "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/pkg/errors"
	"math/big"
	"strconv"
)

// ValidatorIsFullExited Whether the validator has completed the lifecycle of the exit
func ValidatorIsFullExited(validator *consensusApi.Validator) bool {
	return validator.Status == consensusApi.ValidatorStateWithdrawalDone
}

// ValidatorIsRequestExiting Whether the validator request the exit
func ValidatorIsRequestExiting(validator *consensusApi.Validator) bool {
	return validator.Status == consensusApi.ValidatorStateActiveExiting ||
		validator.Status == consensusApi.ValidatorStateActiveSlashed ||
		validator.Status == consensusApi.ValidatorStateExitedUnslashed ||
		validator.Status == consensusApi.ValidatorStateExitedSlashed ||
		validator.Status == consensusApi.ValidatorStateWithdrawalPossible ||
		validator.Status == consensusApi.ValidatorStateWithdrawalDone
}

// ValidatorSlashedAmount 32ETh - ValidatorWithdrawAbleBalance
func ValidatorSlashedAmount(ctx context.Context, validator *consensusApi.Validator) (*big.Int, error) {
	if !validator.Validator.Slashed {
		return nil, errors.New("validator not slashed.")
	}

	pubkey := validator.Validator.PublicKey.String()

	withdrawAbleSlot := consensus.ConsensusClient.ChainTimeService.EpochToSlot(validator.Validator.WithdrawableEpoch)
	pubkeys := make([]string, 1)
	pubkeys[0] = pubkey
	slashedValidator, err := consensus.ConsensusClient.CustomizeBeaconService.ValidatorsByPubKey(ctx, strconv.FormatInt(int64(withdrawAbleSlot), 10), pubkeys)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get validators info.")
	}

	slashAmount := new(big.Int).Sub(eth1.ETH32().BigInt(), eth1.GweiToWei(slashedValidator[pubkey].Balance))

	return slashAmount, nil
}

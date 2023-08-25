// desc:
// @author renshiwei
// Date: 2023/4/10 15:02

package withdraw

import (
	"context"
	"github.com/NodeDAO/oracle-go/common/errs"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/contracts"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/NodeDAO/oracle-go/eth1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
)

func (v *WithdrawHelper) calculationForOperator(ctx context.Context) error {

	// Whether an operator is required to distribute rewards
	if v.clVaultBalance.Cmp(v.clVaultMinSettleLimit) == -1 {
		logger.Debugf("[WithdrawOracle] Not distribute Operator rewards. clVaultBalance < clVaultMinSettleLimit."+
			" clVaultBalance:%s clVaultMinSettleLimit:%s",
			v.clVaultBalance.String(),
			v.clVaultMinSettleLimit.String())
	} else {
		v.isComputeOperatorReward = true
	}

	// Help computing k:operatorId v:EffectiveOperator
	effectiveOperators := make(map[int64]*EffectiveOperator)

	if err := v.calculationOperatorClCapital(ctx, effectiveOperators); err != nil {
		return errors.Wrap(err, "calculationOperatorClCapital err.")
	}

	if v.isComputeOperatorReward {
		if err := v.calculationOperatorWeight(ctx, effectiveOperators); err != nil {
			return errors.Wrap(err, "calculationOperatorWeight err.")
		}

		// calculationOperatorClReward and dealWithdrawInfo
		if err := v.calculationOperatorClReward(ctx, effectiveOperators); err != nil {
			return errors.Wrap(err, "calculationOperatorClReward err.")
		}
	}

	if err := v.calculationWithdrawInfos(ctx, effectiveOperators); err != nil {
		return errors.Wrap(err, "calculationWithdrawInfos err.")
	}

	return nil
}

// if ExitedAmountForClCapital >= 32 ether  ClCapital = 32 ether
// if ExitedAmountForClCapital < 32 ether  ClCapital = ExitedAmountForClCapital
func (v *WithdrawHelper) calculationOperatorClCapital(ctx context.Context, effectiveOperators map[int64]*EffectiveOperator) error {
	for _, exa := range v.requireReportValidator {
		if exa.IsNeedOracleReportExit {
			//  ClCapital Set only system-owned
			owner, err := contracts.VnftContract.Contract.OwnerOf(nil, exa.TokenId)
			if owner != common.HexToAddress(contracts.LiqContract.Address) {
				continue
			}
			if err != nil {
				return errors.Wrapf(err, "Failed to get VnftContract OwnerOf tokenId:%s", exa.TokenId.String())
			}

			_cap := eth1.ETH32().BigInt()
			// less 32 ether
			if exa.ExitedAmountForClCapital.Cmp(eth1.ETH32().BigInt()) == -1 {
				_cap = exa.ExitedAmountForClCapital
			}

			if effectiveOperators[exa.OperatorId.Int64()] == nil {
				effectiveOperators[exa.OperatorId.Int64()] = &EffectiveOperator{
					OperatorReward: withdrawOracle.WithdrawInfo{
						OperatorId: exa.OperatorId.Uint64(),
						ClReward:   big.NewInt(0),
						ClCapital:  _cap,
					},
				}
			} else {
				effectiveOperators[exa.OperatorId.Int64()].OperatorReward.ClCapital = new(big.Int).Add(effectiveOperators[exa.OperatorId.Int64()].OperatorReward.ClCapital, _cap)
			}

			v.totalOperatorClCapital = new(big.Int).Add(v.totalOperatorClCapital, _cap)
		}
	}

	v.clSettleAmount = new(big.Int).Add(v.clSettleAmount, v.totalOperatorClCapital)
	return nil
}

func (v *WithdrawHelper) calculationOperatorWeight(ctx context.Context, effectiveOperators map[int64]*EffectiveOperator) error {
	// Get all the operators
	operatorCount, err := v.getAllOperatorId(ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to get OperatorContract GetNodeOperatorsCount.")
	}

	// operator id start 1.
	for i := int64(1); i < operatorCount.Int64()+1; i++ {
		operatorId := big.NewInt(i)
		isTrusted, err := contracts.OperatorContract.Contract.IsTrustedOperator(nil, operatorId)
		if err != nil {
			return errors.Wrapf(err, "Failed to get OperatorContract IsTrustedOperator operatorId:%v.", i)
		}
		isQuit, err := contracts.OperatorContract.Contract.IsQuitOperator(nil, operatorId)
		if err != nil {
			return errors.Wrapf(err, "Failed to get OperatorContract IsQuitOperator operatorId:%v.", i)
		}

		// effective Operator init
		if isTrusted && !isQuit {
			totalNftCountOfOperator, err := contracts.VnftContract.Contract.GetActiveNftCountsOfOperator(nil, operatorId)
			if err != nil {
				return errors.Wrapf(err, "Failed to get VnftContract GetActiveNftCountsOfOperator operatorId:%v.", i)
			}
			nftUserCountOfOperator, err := contracts.VnftContract.Contract.GetUserActiveNftCountsOfOperator(nil, operatorId)
			if err != nil {
				return errors.Wrapf(err, "Failed to get VnftContract GetUserActiveNftCountsOfOperator operatorId:%v.", i)
			}
			nftCountOperatorOfStakingPool := new(big.Int).Sub(totalNftCountOfOperator, nftUserCountOfOperator)

			clCapital := big.NewInt(0)
			if effectiveOperators[i] != nil {
				clCapital = effectiveOperators[i].OperatorReward.ClCapital
			}

			effectiveOperators[i] = &EffectiveOperator{
				VnftCountOfStakingPool: nftCountOperatorOfStakingPool.Uint64(),
				OperatorReward: withdrawOracle.WithdrawInfo{
					OperatorId: uint64(i),
					ClReward:   big.NewInt(0),
					ClCapital:  clCapital,
				},
			}
		}

	}

	return nil
}

func (v *WithdrawHelper) calculationOperatorClReward(ctx context.Context, effectiveOperators map[int64]*EffectiveOperator) error {
	activeNftsOfStakingPool, err := contracts.VnftContract.Contract.ActiveNftsOfStakingPool(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get VnftContract ActiveNftsOfStakingPool.")
	}
	v.totalNftCountOfStakingPool = big.NewInt(int64(len(activeNftsOfStakingPool)))

	sumReward := new(big.Int).Sub(v.clVaultBalance, v.totalOperatorClCapital)
	if sumReward.Cmp(big.NewInt(0)) == -1 {
		return errs.NewSleepError("OperatorReward.ClReward < 0", RandomSleepTime())
	} else if sumReward.Cmp(eth1.ETH(20)) > 1 {
		return errs.NewSleepError("OperatorReward.ClReward >= 20 ETH", RandomSleepTime())
	}

	for _, op := range effectiveOperators {
		mul := new(big.Int).Mul(sumReward, big.NewInt(int64(op.VnftCountOfStakingPool)))
		op.OperatorReward.ClReward = new(big.Int).Div(mul, v.totalNftCountOfStakingPool)
	}

	v.clSettleAmount = new(big.Int).Add(v.clSettleAmount, sumReward)

	// deal accuracy
	sumSettle := big.NewInt(0)

	for _, op := range effectiveOperators {
		opSettle := new(big.Int).Add(op.OperatorReward.ClReward, op.OperatorReward.ClCapital)
		sumSettle = new(big.Int).Add(sumSettle, opSettle)
	}

	if v.clSettleAmount.Cmp(sumSettle) != 0 {
		accuracy := new(big.Int).Sub(v.clSettleAmount, sumSettle)
		effectiveOperators[1].OperatorReward.ClReward = new(big.Int).Add(effectiveOperators[1].OperatorReward.ClReward, accuracy)
	}

	return nil
}

func (v *WithdrawHelper) calculationWithdrawInfos(ctx context.Context, effectiveOperators map[int64]*EffectiveOperator) error {
	for _, op := range effectiveOperators {
		v.withdrawInfos = append(v.withdrawInfos, op.OperatorReward)
	}

	sumSettleCheck := big.NewInt(0)
	for _, info := range v.withdrawInfos {
		opSettle := new(big.Int).Add(info.ClReward, info.ClCapital)
		sumSettleCheck = new(big.Int).Add(sumSettleCheck, opSettle)
	}
	if v.clSettleAmount.Cmp(sumSettleCheck) != 0 {
		return errors.Errorf("clSettleAmount(%s) != sumSettleCheck(%s)", v.clSettleAmount, sumSettleCheck)
	}

	return nil
}

func (v *WithdrawHelper) getAllOperatorId(ctx context.Context) (*big.Int, error) {
	return contracts.OperatorContract.Contract.GetNodeOperatorsCount(nil)
}

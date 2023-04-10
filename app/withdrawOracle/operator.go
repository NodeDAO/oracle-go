// desc:
// @author renshiwei
// Date: 2023/4/10 15:02

package withdrawOracle

import (
	"context"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/contracts"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"math/big"
)

func (v *WithdrawHelper) calculationForOperator(ctx context.Context) error {
	// Whether an operator is required to distribute rewards
	if v.clVaultBalance.Cmp(v.clVaultMinSettleLimit) == -1 {
		logger.Infof("clVaultBalance less than clVaultMinSettleLimit, There is no need to distribute Operator rewards."+
			" clVaultBalance:%s clVaultMinSettleLimit:%s",
			v.clBalance.String(),
			v.clVaultMinSettleLimit.String())
		return nil
	}

	// Help computing
	effectiveOperators := make(map[int64]*EffectiveOperator)

	if err := v.calculationOperatorWeight(ctx, effectiveOperators); err != nil {
		return errors.Unwrap(err)
	}

	if err := v.calculationOperatorClCapital(ctx, effectiveOperators); err != nil {
		return errors.Unwrap(err)
	}

	// calculationOperatorClReward and dealWithdrawInfo
	if err := v.calculationOperatorClReward(ctx, effectiveOperators); err != nil {
		return errors.Unwrap(err)
	}

	return nil
}

func (v *WithdrawHelper) calculationOperatorWeight(ctx context.Context, effectiveOperators map[int64]*EffectiveOperator) error {
	// Get all the operators
	operatorCount, err := v.getAllOperatorId(ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to get OperatorContract GetNodeOperatorsCount.")
	}

	// operator id start 0.
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
			nftCount, err := contracts.VnftContract.Contract.GetNftCountsOfOperator(nil, operatorId)
			if err != nil {
				return errors.Wrapf(err, "Failed to get VnftContract GetNftCountsOfOperator operatorId:%v.", i)
			}

			effectiveOperators[i] = &EffectiveOperator{
				VnftCount: nftCount.Uint64(),
				OperatorReward: &withdrawOracle.WithdrawInfo{
					OperatorId: uint64(i),
				},
			}
		}

	}

	return nil
}

func (v *WithdrawHelper) calculationOperatorClCapital(ctx context.Context, effectiveOperators map[int64]*EffectiveOperator) error {
	for _, exa := range v.requireReportValidator {
		if exa.IsNeedOracleReportExit {
			effectiveOperators[exa.TokenId.Int64()].OperatorReward.ClCapital = new(big.Int).Add(effectiveOperators[exa.TokenId.Int64()].OperatorReward.ClCapital, exa.ExitedAmount)
			v.totalOperatorClCapital = new(big.Int).Add(v.totalOperatorClCapital, exa.ExitedAmount)
		}
	}

	return nil
}

func (v *WithdrawHelper) calculationOperatorClReward(ctx context.Context, effectiveOperators map[int64]*EffectiveOperator) error {
	sumReward := new(big.Int).Sub(v.clVaultBalance, v.totalOperatorClCapital)
	for _, op := range effectiveOperators {
		op.OperatorReward.ClReward = decimal.NewFromBigInt(sumReward, 0).Mul(decimal.NewFromInt(int64(op.VnftCount))).Div(decimal.NewFromBigInt(v.totalNftCount, 0)).BigInt()
		//dealWithdrawInfo
		v.withdrawInfos = append(v.withdrawInfos, op.OperatorReward)
	}

	return nil
}

func (v *WithdrawHelper) getAllOperatorId(ctx context.Context) (*big.Int, error) {
	return contracts.OperatorContract.Contract.GetNodeOperatorsCount(nil)
}

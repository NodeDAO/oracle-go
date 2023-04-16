// desc:
// @author renshiwei
// Date: 2023/4/10 15:02

package withdraw

import (
	"context"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/contracts"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/NodeDAO/oracle-go/eth1"
	"github.com/ethereum/go-ethereum/common"
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
	} else {
		v.isComputeOperatorReward = true
	}

	// Help computing k:operatorId v:EffectiveOperator
	effectiveOperators := make(map[int64]*EffectiveOperator)

	if err := v.calculationOperatorClCapital(ctx, effectiveOperators); err != nil {
		return errors.Wrap(err, "")
	}

	if v.isComputeOperatorReward {
		if err := v.calculationOperatorWeight(ctx, effectiveOperators); err != nil {
			return errors.Wrap(err, "")
		}

		// calculationOperatorClReward and dealWithdrawInfo
		if err := v.calculationOperatorClReward(ctx, effectiveOperators); err != nil {
			return errors.Wrap(err, "")
		}
	}

	if err := v.calculationWithdrawInfos(ctx, effectiveOperators); err != nil {
		return errors.Wrap(err, "")
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
			nftCount, err := contracts.VnftContract.Contract.GetActiveNftCountsOfOperator(nil, operatorId)
			if err != nil {
				return errors.Wrapf(err, "Failed to get VnftContract GetActiveNftCountsOfOperator operatorId:%v.", i)
			}

			effectiveOperators[i] = &EffectiveOperator{
				VnftCount: nftCount.Uint64(),
				OperatorReward: withdrawOracle.WithdrawInfo{
					OperatorId: uint64(i),
				},
			}
		}

	}

	return nil
}

// if ExitedAmount >= 32 ether  ClCapital = 32 ether
// if ExitedAmount < 32 ether  ClCapital = ExitedAmount
func (v *WithdrawHelper) calculationOperatorClCapital(ctx context.Context, effectiveOperators map[int64]*EffectiveOperator) error {
	for _, exa := range v.requireReportValidator {
		if exa.IsNeedOracleReportExit {
			//  ClCapital Set only system-owned
			owner, err := contracts.VnftContract.Contract.OwnerOf(nil, exa.TokenId)
			if owner != common.HexToAddress(contracts.LiqContract.Address) {
				break
			}
			if err != nil {
				return errors.Wrapf(err, "Failed to get VnftContract OwnerOf tokenId:%s", exa.TokenId.String())
			}

			_cap := eth1.ETH32().BigInt()
			// less 32 ether
			if exa.ExitedAmount.Cmp(eth1.ETH32().BigInt()) == -1 {
				_cap = exa.ExitedAmount
			}
			effectiveOperators[exa.OperatorId.Int64()].OperatorReward.ClCapital = new(big.Int).Add(effectiveOperators[exa.OperatorId.Int64()].OperatorReward.ClCapital, _cap)
			v.totalOperatorClCapital = new(big.Int).Add(v.totalOperatorClCapital, exa.ExitedAmount)
		}
	}

	v.clSettleAmount = new(big.Int).Add(v.clSettleAmount, v.totalOperatorClCapital)
	return nil
}

func (v *WithdrawHelper) calculationOperatorClReward(ctx context.Context, effectiveOperators map[int64]*EffectiveOperator) error {
	sumReward := new(big.Int).Sub(v.clVaultBalance, v.totalOperatorClCapital)
	for _, op := range effectiveOperators {
		op.OperatorReward.ClReward = decimal.NewFromBigInt(sumReward, 0).Mul(decimal.NewFromInt(int64(op.VnftCount))).Div(decimal.NewFromBigInt(v.totalNftCount, 0)).BigInt()
	}

	v.clSettleAmount = new(big.Int).Add(v.clSettleAmount, sumReward)

	return nil
}

func (v *WithdrawHelper) calculationWithdrawInfos(ctx context.Context, effectiveOperators map[int64]*EffectiveOperator) error {
	for _, op := range effectiveOperators {
		v.withdrawInfos = append(v.withdrawInfos, op.OperatorReward)
	}
	return nil
}

func (v *WithdrawHelper) getAllOperatorId(ctx context.Context) (*big.Int, error) {
	return contracts.OperatorContract.Contract.GetNodeOperatorsCount(nil)
}

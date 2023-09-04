// description:
// @author renshiwei
// Date: 2023/6/29

package largestake

import (
	"context"
	"github.com/NodeDAO/oracle-go/common/errs"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/NodeDAO/oracle-go/consensus/beacon"
	"github.com/NodeDAO/oracle-go/contracts"
	"github.com/NodeDAO/oracle-go/contracts/largeStakeOracle"
	"github.com/NodeDAO/oracle-go/eth1"
	consensusApi "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"math/big"
	"sort"
	"strings"
	"time"
)

func EncodeLargeStakeReportData(reportData *largeStakeOracle.LargeStakeOracleReportData) ([32]byte, error) {
	json, err := abi.JSON(strings.NewReader("[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"consensusVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"stakingId\",\"type\":\"uint128\"},{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"}],\"internalType\":\"struct CLStakingExitInfo[]\",\"name\":\"clStakingExitInfos\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"stakingId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"slashAmount\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"internalType\":\"struct CLStakingSlashInfo[]\",\"name\":\"clStakingSlashInfos\",\"type\":\"tuple[]\"}],\"internalType\":\"struct LargeStakeOracle.ReportData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"submitReportData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"))
	encodedData, err := json.Methods["submitReportData"].Inputs.Pack(reportData)

	// Code the data using the Pack method of the abi.Arguments structure
	//encodedData, err := reportDataArguments.Pack(reportData)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "encode report data err.")
	}

	// Calculate the keccak 256 hash value
	hash := crypto.Keccak256(encodedData)
	var res [32]byte
	copy(res[:], hash)

	return res, nil
}

func NewLargeStakeHelper(refSlot, consensusVersion, refBlockNumber *big.Int) *LargeStakeHelper {
	return &LargeStakeHelper{
		refSlot:          refSlot,
		consensusVersion: consensusVersion,
		refBlockNumber:   refBlockNumber,
	}
}

func (v *LargeStakeHelper) ProcessReportData(ctx context.Context) (*LargeStakeReportRes, error) {
	stakeValidatorMap, err := v.getLargeStakeValidator(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get large stake validator.")
	}

	clStakingExitInfos, clStakingSlashInfos, err := v.filterExitedSlashedValidator(ctx, stakeValidatorMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to filter exited slashed validator.")
	}

	if len(clStakingExitInfos) == 0 && len(clStakingSlashInfos) == 0 {
		logger.Debug("[LargeStakeOracle] no validator exited and slashed.")
		return &LargeStakeReportRes{
			IsNeedReport:   false,
			ReportData:     &largeStakeOracle.LargeStakeOracleReportData{},
			ReportDataHash: eth1.ZERO_HASH,
		}, nil
	}

	exitLimit, err := contracts.LargeStakeOracleContract.Contract.ExitLimit(nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get LargeStakeOracleContract ExitLimit.")
	}

	// If the exit validator is greater than the exit limit, only the exit limit is taken
	if len(clStakingExitInfos) > int(exitLimit.Int64()) {
		logger.Warn("[LargeStakeOracle] exited validator count is more than exit limit.")
		clStakingExitInfos = clStakingExitInfos[:exitLimit.Int64()]

		if len(clStakingSlashInfos) > int(exitLimit.Int64()) {
			sort.Slice(clStakingSlashInfos, func(i, j int) bool {
				return string(clStakingSlashInfos[i].Pubkey) < string(clStakingSlashInfos[j].Pubkey)
			})
			clStakingSlashInfos = clStakingSlashInfos[:exitLimit.Int64()]
		}
	}

	reportData := &largeStakeOracle.LargeStakeOracleReportData{
		ConsensusVersion:    v.consensusVersion,
		RefSlot:             v.refSlot,
		ClStakingExitInfos:  clStakingExitInfos,
		ClStakingSlashInfos: clStakingSlashInfos,
	}
	reportDataHash, err := EncodeLargeStakeReportData(reportData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode report data.")
	}

	res := &LargeStakeReportRes{
		IsNeedReport:   true,
		ReportData:     reportData,
		ReportDataHash: reportDataHash,
	}

	return res, nil
}

func (v *LargeStakeHelper) getLargeStakeValidator(ctx context.Context) (map[string]*LargeStakeValidator, error) {
	pubkeys := make([]string, 0)
	validatorExaMap := make(map[string]*LargeStakeValidator)

	// Get all the largeStake info
	stakingCounts, err := contracts.LargeStakingContract.Contract.TotalLargeStakingCounts(nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get LargeStakingContract TotalLargeStakingCounts.")
	}

	if stakingCounts.Cmp(big.NewInt(0)) == 0 {
		logger.Warn("[LargeStakeOracle] stakingCounts is zero.")
		return validatorExaMap, nil
	}

	for i := uint64(1); i <= stakingCounts.Uint64(); i++ {
		validatorBytes, err := contracts.LargeStakingContract.Contract.GetValidatorsOfStakingId(nil, big.NewInt(int64(i)))
		if err != nil {
			return nil, errors.Wrap(err, "failed to get LargeStakingContract GetValidatorsOfStakingId.")
		}
		for _, validatorByte := range validatorBytes {
			pubkey := hexutil.Encode(validatorByte)
			pubkeys = append(pubkeys, pubkey)

			validatorExa := &LargeStakeValidator{
				StakingId: big.NewInt(int64(i)),
			}
			validatorExaMap[pubkey] = validatorExa
		}
	}

	validators, err := consensus.ConsensusClient.CustomizeBeaconService.ValidatorsByPubKey(ctx, v.refSlot.String(), pubkeys)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get validators info. pubkeys:%+v", pubkeys)
	}
	for _, pubkey := range pubkeys {
		validator, ok := validators[pubkey]
		// Handle pubkeys that are not query by Beacon
		// set balance = 32 GWEI
		if !ok {
			validator = &consensusApi.Validator{
				Balance: 32e9,
			}
		}
		validatorExaMap[pubkey].Validator = validator
	}

	return validatorExaMap, nil
}

func (v *LargeStakeHelper) filterExitedSlashedValidator(
	ctx context.Context,
	validatorExaMap map[string]*LargeStakeValidator,
) ([]largeStakeOracle.CLStakingExitInfo, []largeStakeOracle.CLStakingSlashInfo, error) {
	clStakingExitInfos := make([]largeStakeOracle.CLStakingExitInfo, 0)
	clStakingSlashInfos := make([]largeStakeOracle.CLStakingSlashInfo, 0)

	exitStakingIdMap := make(map[uint64][][]byte)

	for pubkey, validatorExa := range validatorExaMap {
		pubkeyByte, err := hexutil.Decode(pubkey)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "failed to decode pubkey:%s", pubkey)
		}

		validatorInfo, err := contracts.LargeStakingContract.Contract.ValidatorInfo(nil, pubkeyByte)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "failed to get LargeStakingContract ValidatorInfo.pubkey:%s", pubkey)
		}

		if validatorExa.Validator.Status == consensusApi.ValidatorStateUnknown {
			if validatorInfo.RegisterBlock.Cmp(v.refBlockNumber) < 0 && validatorInfo.ExitBlock.Cmp(big.NewInt(0)) == 0 {
				if new(big.Int).Sub(v.refBlockNumber, validatorInfo.RegisterBlock).Cmp(big.NewInt(TWO_DAY_BLOCK_NUMBER)) > 0 {
					if _, ok := exitStakingIdMap[validatorExa.StakingId.Uint64()]; !ok {
						exitStakingIdMap[validatorExa.StakingId.Uint64()] = make([][]byte, 0)
					}
					exitStakingIdMap[validatorExa.StakingId.Uint64()] = append(exitStakingIdMap[validatorExa.StakingId.Uint64()], pubkeyByte)
				}
			}
			continue
		}

		if beacon.ValidatorIsFullExited(validatorExa.Validator) {
			if validatorInfo.ExitBlock.Cmp(big.NewInt(0)) == 0 {
				if _, ok := exitStakingIdMap[validatorExa.StakingId.Uint64()]; !ok {
					exitStakingIdMap[validatorExa.StakingId.Uint64()] = make([][]byte, 0)
				}
				exitStakingIdMap[validatorExa.StakingId.Uint64()] = append(exitStakingIdMap[validatorExa.StakingId.Uint64()], pubkeyByte)
			}
		}

		if validatorExa.Validator.Status != consensusApi.ValidatorStateUnknown && validatorExa.Validator.Validator.Slashed {
			if validatorInfo.SlashAmount.Cmp(big.NewInt(0)) == 0 {
				validatorSlashedAmount, err := consensus.ConsensusClient.CustomizeBeaconService.ValidatorSlashedAmount(ctx, validatorExa.Validator)
				if err != nil {
					return nil, nil, errors.Wrapf(err, "failed to get ValidatorSlashedAmount:%s", pubkey)
				}
				if validatorSlashedAmount.Cmp(eth1.ETH(2)) == 1 {
					logger.Warn("[LargeStakeOracle] slash amount > 2 ETH.",
						zap.String("pubkey", pubkey),
						zap.String("slashAmount", validatorSlashedAmount.String()),
					)
					return nil, nil, errs.NewSleepError("[LargeStakeOracle] slash amount > 2 ETH", 10*time.Minute)
				}

				clStakingSlashInfos = append(clStakingSlashInfos, largeStakeOracle.CLStakingSlashInfo{
					StakingId:   validatorExa.StakingId,
					SlashAmount: validatorSlashedAmount,
					Pubkey:      pubkeyByte,
				})
			}
		}
	}

	for stakingId, pubkeyBytes := range exitStakingIdMap {
		sort.Slice(pubkeyBytes, func(i, j int) bool {
			return string(pubkeyBytes[i]) < string(pubkeyBytes[j])
		})
		clStakingExitInfos = append(clStakingExitInfos, largeStakeOracle.CLStakingExitInfo{
			StakingId: big.NewInt(int64(stakingId)),
			Pubkeys:   pubkeyBytes,
		})
	}

	return clStakingExitInfos, clStakingSlashInfos, nil
}

package withdraw

import (
	"context"
	"fmt"
	"github.com/NodeDAO/oracle-go/app"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/config"
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/NodeDAO/oracle-go/contracts"
	"github.com/NodeDAO/oracle-go/eth1"
	consensusApi "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"math/big"
	"strconv"
	"testing"
)

func TestValidatorBalance(t *testing.T) {
	var err error
	ctx := context.Background()
	config.InitConfig("../../conf/config-mainnet.dev.yaml")
	logger.InitLog("debug", "console")
	app.InitServer(ctx)

	w := new(WithdrawHelper)
	w.clBalance = big.NewInt(0)
	w.validatorExaMap = make(map[string]*ValidatorExa)

	consensusContract, err := w.oracle.GetConsensusContract(ctx)
	require.NoError(t, err)

	currentFrame, err := consensusContract.GetCurrentFrame(nil)
	require.NoError(t, err)
	logger.Debug("current frame",
		zap.String("RefSlot", currentFrame.RefSlot.String()),
		zap.String("DeadlineSlot", currentFrame.ReportProcessingDeadlineSlot.String()),
	)

	headSlot, err := consensus.ConsensusClient.CustomizeBeaconService.HeadSlot(ctx)
	if headSlot.Cmp(currentFrame.RefSlot) == -1 || headSlot.Cmp(currentFrame.ReportProcessingDeadlineSlot) == 1 {
		panic(fmt.Sprintf("slot is not correct. headslot:%s", headSlot.String()))
	}

	w.refSlot = currentFrame.RefSlot

	err = w.obtainValidatorConsensusInfo(ctx)
	require.NoError(t, err)

	logger.Debug("-------start calculate validator balance-----", zap.Int("count", len(w.validatorExaMap)))
	for pubkey, exa := range w.validatorExaMap {
		if exa.VnftOwner == LiquidStaking {
			// sum cl balance
			w.clBalance = new(big.Int).Add(w.clBalance, big.NewInt(int64(exa.Validator.Balance)))
		}
		fmt.Printf("pubkey:%s balance:%v status:%s \n", pubkey, exa.Validator.Balance, exa.Validator.Status)
	}
	logger.Debug("-------end calculate validator balance-----")

	logger.Debugf("clBalance:%s", w.clBalance.String())
}

func TestNodeDAOValidatorForSlot(t *testing.T) {
	ctx := context.Background()
	config.InitConfig("../../conf/config-goerli.dev.yaml")
	logger.InitLog("debug", "console")
	app.InitServer(ctx)

	elBlockNumber := big.NewInt(9566221)

	v := &WithdrawHelper{
		refSlot:      big.NewInt(6355903),
		deadlineSlot: big.NewInt(6356223),
	}

	logger.Info("block info",
		zap.String("block number", elBlockNumber.String()),
		zap.String("refSlot", v.refSlot.String()),
	)

	v.validatorExaMap = make(map[string]*ValidatorExa)
	v.requireReportValidator = make(map[string]*ValidatorExa)

	opt := &bind.CallOpts{
		BlockNumber: elBlockNumber,
	}

	validatorBytesOfStakingPool, err := contracts.VnftContract.Contract.ActiveValidatorsOfStakingPool(opt)
	require.NoError(t, err)
	err = v.parseValidatorExaMap(ctx, validatorBytesOfStakingPool, LiquidStaking)
	require.NoError(t, err)

	validatorBytesOfUser, err := contracts.VnftContract.Contract.ActiveValidatorOfUser(opt)
	require.NoError(t, err)
	err = v.parseValidatorExaMap(ctx, validatorBytesOfUser, USER)
	require.NoError(t, err)

	exitTokenIds := make([]*big.Int, 0)
	validatorExaArr := make([]*ValidatorExa, 0)
	for pubkey, exa := range v.validatorExaMap {
		pubkeyBytes, err := hexutil.Decode(pubkey)
		require.NoError(t, err)

		// get tokenId
		tokenId, err := contracts.VnftContract.Contract.TokenOfValidator(opt, pubkeyBytes)
		require.NoError(t, err)
		exa.TokenId = tokenId

		operatorId, err := contracts.VnftContract.Contract.OperatorOf(opt, tokenId)
		require.NoError(t, err)
		exa.OperatorId = operatorId

		// IsExited
		if (exa.Validator.Status == consensusApi.ValidatorStateWithdrawalDone || exa.Validator.Status == consensusApi.ValidatorStateWithdrawalPossible) && exa.Validator.Balance == 0 {
			exa.IsExited = true
			// ExitedSlot
			exitedSlot := consensus.ConsensusClient.ChainTimeService.FirstSlotOfEpoch(exa.Validator.Validator.ExitEpoch)
			exa.ExitedSlot = big.NewInt(int64(exitedSlot))

			// ExitedBlockHeight
			executionBlock, err := consensus.ConsensusClient.CustomizeBeaconService.ExecutionBlock(ctx, strconv.Itoa(int(exitedSlot)))
			require.NoError(t, err)

			require.NotEqual(t, executionBlock.BlockNumber.Uint64(), 0)
			exa.ExitedBlockHeight = executionBlock.BlockNumber

			// IsNeedOracleReportExit
			exitTokenIds = append(exitTokenIds, tokenId)
		}

		validatorExaArr = append(validatorExaArr, exa)
	}

	// IsNeedOracleReportExit
	oracleReportExitNumbers, err := contracts.VnftContract.Contract.GetNftExitBlockNumbers(opt, exitTokenIds)
	require.NoError(t, err)
	for i, number := range oracleReportExitNumbers {
		if number.Cmp(decimal.Zero.BigInt()) == 0 {
			for _, exa := range validatorExaArr {
				pubkey := exa.Validator.Validator.PublicKey.String()
				if exitTokenIds[i] == exa.TokenId {
					exa.IsNeedOracleReportExit = true

					logger.Info("IsNeedOracleReportExit.",
						zap.String("pubkey", pubkey),
						zap.String("tokenId", exa.TokenId.String()),
					)
					break
				}
			}
		}
	}

	clVaultBalance, err := eth1.ElClient.BalanceAt(ctx, contracts.GetClVaultAddress(), elBlockNumber)
	require.NoError(t, err)
	logger.Info("clVaultBalance", zap.String("clVaultBalance", clVaultBalance.String()))
}

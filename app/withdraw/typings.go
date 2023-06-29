// desc:
// @author renshiwei
// Date: 2023/4/7 15:33

package withdraw

import (
	"github.com/NodeDAO/oracle-go/app/consensusModule"
	"github.com/NodeDAO/oracle-go/consensus/beacon"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/NodeDAO/oracle-go/utils/timetool"
	consensusApi "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"time"
)

type WithdrawHelper struct {
	// param
	refSlot          *big.Int
	deadlineSlot     *big.Int
	consensusVersion *big.Int

	// initial
	executionBlock           *beacon.ExecutionBlock
	delayedExitSlashStandard *big.Int
	clVaultMinSettleLimit    *big.Int
	exitRequestLimit         *big.Int
	keyTransactOpts          *bind.TransactOpts
	withdrawOracleModuleId   *big.Int

	// compute process
	validatorExaMap         map[string]*ValidatorExa
	requireReportValidator  map[string]*ValidatorExa
	totalOperatorClCapital  *big.Int
	totalNftCount           *big.Int
	isComputeOperatorReward bool

	// process res
	clBalance                  *big.Int
	clVaultBalance             *big.Int
	delayedExitTokenIds        []*big.Int
	largeExitDelayedRequestIds []*big.Int
	clSettleAmount             *big.Int
	withdrawInfos              []withdrawOracle.WithdrawInfo
	exitValidatorInfos         []withdrawOracle.ExitValidatorInfo

	// res
	reportData *withdrawOracle.WithdrawOracleReportData

	// report
	oracle              *Oracle
	hashConsensusHelper *consensusModule.HashConsensusHelper
}

type ValidatorExa struct {
	Validator *consensusApi.Validator

	IsExited          bool
	ExitedSlot        *big.Int
	ExitedBlockHeight *big.Int

	TokenId    *big.Int
	OperatorId *big.Int
	// Whether oracle needs to make a report
	IsNeedOracleReportExit bool
	// IsNeedOracleReportExit = true And then to calculate
	ExitedAmount *big.Int
	// Whether the tokenId is owned by the pledge pool
	IsOwnerLiqPool bool

	// 1.slashed 2.exited 3.Not OracleReportExit
	SlashAmount *big.Int

	IsDelayedExit bool
}

type EffectiveOperator struct {
	VnftCount      uint64
	OperatorReward withdrawOracle.WithdrawInfo
}

type Oracle struct {
}

// version
const (
	CONSENSUS_VERSION               = 2
	WITHRAW_ORACLE_CONTRACT_VERSION = 2

	SLOTS_PER_EPOCH   = 32
	SECONDS_PER_SLOT  = 12
	SECONDS_PER_EPOCH = SLOTS_PER_EPOCH * SECONDS_PER_SLOT
)

func DefaultRandomSleep() {
	minSleep := time.Second * SECONDS_PER_EPOCH
	maxSleep := time.Second * SECONDS_PER_EPOCH * 2
	timetool.SleepWithRandom(minSleep, maxSleep)
}

func RandomSleepTime() time.Duration {
	minSleep := time.Second * SECONDS_PER_EPOCH
	maxSleep := time.Second * SECONDS_PER_EPOCH * 2
	return timetool.RandomTime(minSleep, maxSleep)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package withdrawOracle

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ExitValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type ExitValidatorInfo struct {
	ExitTokenId     uint64
	ExitBlockNumber *big.Int
	SlashAmount     *big.Int
}

// WithdrawInfo is an auto generated low-level Go binding around an user-defined struct.
type WithdrawInfo struct {
	OperatorId uint64
	ClReward   *big.Int
	ClCapital  *big.Int
}

// WithdrawOracleProcessingState is an auto generated low-level Go binding around an user-defined struct.
type WithdrawOracleProcessingState struct {
	CurrentFrameRefSlot    *big.Int
	ProcessingDeadlineTime *big.Int
	DataHash               [32]byte
	DataSubmitted          bool
	ReportExitedCount      *big.Int
}

// WithdrawOracleReportData is an auto generated low-level Go binding around an user-defined struct.
type WithdrawOracleReportData struct {
	ConsensusVersion           *big.Int
	RefSlot                    *big.Int
	ClBalance                  *big.Int
	ClVaultBalance             *big.Int
	ClSettleAmount             *big.Int
	ReportExitedCount          *big.Int
	WithdrawInfos              []WithdrawInfo
	ExitValidatorInfos         []ExitValidatorInfo
	DelayedExitTokenIds        []*big.Int
	LargeExitDelayedRequestIds []*big.Int
}

// WithdrawOracleMetaData contains all meta data concerning the WithdrawOracle contract.
var WithdrawOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AddressCannotBeSame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ArgumentOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClVaultBalanceNotMinSettleLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClVaultMinSettleLimitNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DaoCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExitRequestLimitNotZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingRefSlot\",\"type\":\"uint256\"}],\"name\":\"InitialRefSlotCannotBeLessThanProcessingOne\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContractVersionIncrement\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestsData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestsDataLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotal\",\"type\":\"uint256\"}],\"name\":\"InvalidTotalBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ModuleIdIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ModuleIdNotEqual\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroContractVersionOnInit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyConsensusContractCanSubmitReport\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ProcessingDeadlineMissed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefSlotAlreadyProcessing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevRefSlot\",\"type\":\"uint256\"}],\"name\":\"RefSlotCannotDecrease\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingRefSlot\",\"type\":\"uint256\"}],\"name\":\"RefSlotMustBeGreaterThanProcessingOne\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedChainConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedVersion\",\"type\":\"uint256\"}],\"name\":\"UnexpectedConsensusVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"UnexpectedContractVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"consensusHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receivedHash\",\"type\":\"bytes32\"}],\"name\":\"UnexpectedDataHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"consensusRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataRefSlot\",\"type\":\"uint256\"}],\"name\":\"UnexpectedRefSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedRequestsDataLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"format\",\"type\":\"uint256\"}],\"name\":\"UnsupportedRequestsDataFormat\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ValidatorReportedExited\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionCannotBeSame\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevAddr\",\"type\":\"address\"}],\"name\":\"ConsensusHashContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prevVersion\",\"type\":\"uint256\"}],\"name\":\"ConsensusVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"ContractVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldDao\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"DaoAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldLiq\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLiq\",\"type\":\"address\"}],\"name\":\"LiquidStakingChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_addBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalBalance\",\"type\":\"uint256\"}],\"name\":\"PendingBalancesAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBalance\",\"type\":\"uint256\"}],\"name\":\"PendingBalancesReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ProcessingStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reportExitedCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clVaultBalance\",\"type\":\"uint256\"}],\"name\":\"ReportDataSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"}],\"name\":\"ReportSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clVaultMinSettleLimit\",\"type\":\"uint256\"}],\"name\":\"UpdateClVaultMinSettleLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitRequestLimit\",\"type\":\"uint256\"}],\"name\":\"UpdateExitRequestLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"old\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBalanceTolerate\",\"type\":\"uint256\"}],\"name\":\"UpdateTotalBalanceTolerate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVaultManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVaultManager\",\"type\":\"address\"}],\"name\":\"VaultManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitRequestLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reportExitedCount\",\"type\":\"uint256\"}],\"name\":\"WarnDataIncompleteProcessing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"}],\"name\":\"WarnProcessingMissed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GENESIS_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_PER_SLOT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pendingBalance\",\"type\":\"uint256\"}],\"name\":\"addPendingBalances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clVaultBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clVaultMinSettleLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exitRequestLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClVaultBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusReport\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"processingStarted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastClSettleAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastProcessingRefSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProcessingState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currentFrameRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"dataSubmitted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"reportExitedCount\",\"type\":\"uint256\"}],\"internalType\":\"structWithdrawOracle.ProcessingState\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"secondsPerSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"genesisTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consensusContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"consensusVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastProcessingRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_exitRequestLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_clVaultMinSettleLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_clBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pendingBalance\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensus\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lastProcessingRefSlot\",\"type\":\"uint256\"}],\"name\":\"initializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastClSettleAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRefSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidStakingContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_clVaultMinSettleLimit\",\"type\":\"uint256\"}],\"name\":\"setClVaultMinSettleLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setConsensusContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"setConsensusVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"setDaoAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_exitRequestLimit\",\"type\":\"uint256\"}],\"name\":\"setExitRequestLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidStakingContractAddress\",\"type\":\"address\"}],\"name\":\"setLiquidStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalBalanceTolerate\",\"type\":\"uint256\"}],\"name\":\"setTotalBalanceTolerate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultManagerContractAddress\",\"type\":\"address\"}],\"name\":\"setVaultManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reportHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_moduleId\",\"type\":\"uint256\"}],\"name\":\"submitConsensusReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"consensusVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"clBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"clVaultBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"clSettleAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reportExitedCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"clReward\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"clCapital\",\"type\":\"uint96\"}],\"internalType\":\"structWithdrawInfo[]\",\"name\":\"withdrawInfos\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"exitTokenId\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"exitBlockNumber\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"slashAmount\",\"type\":\"uint96\"}],\"internalType\":\"structExitValidatorInfo[]\",\"name\":\"exitValidatorInfos\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"delayedExitTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"largeExitDelayedRequestIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structWithdrawOracle.ReportData\",\"name\":\"data\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_contractVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_moduleId\",\"type\":\"uint256\"}],\"name\":\"submitReportData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBalanceTolerate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateContractVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WithdrawOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use WithdrawOracleMetaData.ABI instead.
var WithdrawOracleABI = WithdrawOracleMetaData.ABI

// WithdrawOracle is an auto generated Go binding around an Ethereum contract.
type WithdrawOracle struct {
	WithdrawOracleCaller     // Read-only binding to the contract
	WithdrawOracleTransactor // Write-only binding to the contract
	WithdrawOracleFilterer   // Log filterer for contract events
}

// WithdrawOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type WithdrawOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WithdrawOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WithdrawOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WithdrawOracleSession struct {
	Contract     *WithdrawOracle   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WithdrawOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WithdrawOracleCallerSession struct {
	Contract *WithdrawOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// WithdrawOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WithdrawOracleTransactorSession struct {
	Contract     *WithdrawOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// WithdrawOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type WithdrawOracleRaw struct {
	Contract *WithdrawOracle // Generic contract binding to access the raw methods on
}

// WithdrawOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WithdrawOracleCallerRaw struct {
	Contract *WithdrawOracleCaller // Generic read-only contract binding to access the raw methods on
}

// WithdrawOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WithdrawOracleTransactorRaw struct {
	Contract *WithdrawOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWithdrawOracle creates a new instance of WithdrawOracle, bound to a specific deployed contract.
func NewWithdrawOracle(address common.Address, backend bind.ContractBackend) (*WithdrawOracle, error) {
	contract, err := bindWithdrawOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WithdrawOracle{WithdrawOracleCaller: WithdrawOracleCaller{contract: contract}, WithdrawOracleTransactor: WithdrawOracleTransactor{contract: contract}, WithdrawOracleFilterer: WithdrawOracleFilterer{contract: contract}}, nil
}

// NewWithdrawOracleCaller creates a new read-only instance of WithdrawOracle, bound to a specific deployed contract.
func NewWithdrawOracleCaller(address common.Address, caller bind.ContractCaller) (*WithdrawOracleCaller, error) {
	contract, err := bindWithdrawOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleCaller{contract: contract}, nil
}

// NewWithdrawOracleTransactor creates a new write-only instance of WithdrawOracle, bound to a specific deployed contract.
func NewWithdrawOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*WithdrawOracleTransactor, error) {
	contract, err := bindWithdrawOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleTransactor{contract: contract}, nil
}

// NewWithdrawOracleFilterer creates a new log filterer instance of WithdrawOracle, bound to a specific deployed contract.
func NewWithdrawOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*WithdrawOracleFilterer, error) {
	contract, err := bindWithdrawOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleFilterer{contract: contract}, nil
}

// bindWithdrawOracle binds a generic wrapper to an already deployed contract.
func bindWithdrawOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WithdrawOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WithdrawOracle *WithdrawOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WithdrawOracle.Contract.WithdrawOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WithdrawOracle *WithdrawOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.WithdrawOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WithdrawOracle *WithdrawOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.WithdrawOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WithdrawOracle *WithdrawOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WithdrawOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WithdrawOracle *WithdrawOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WithdrawOracle *WithdrawOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.contract.Transact(opts, method, params...)
}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCaller) GENESISTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "GENESIS_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleSession) GENESISTIME() (*big.Int, error) {
	return _WithdrawOracle.Contract.GENESISTIME(&_WithdrawOracle.CallOpts)
}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCallerSession) GENESISTIME() (*big.Int, error) {
	return _WithdrawOracle.Contract.GENESISTIME(&_WithdrawOracle.CallOpts)
}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCaller) SECONDSPERSLOT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "SECONDS_PER_SLOT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleSession) SECONDSPERSLOT() (*big.Int, error) {
	return _WithdrawOracle.Contract.SECONDSPERSLOT(&_WithdrawOracle.CallOpts)
}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCallerSession) SECONDSPERSLOT() (*big.Int, error) {
	return _WithdrawOracle.Contract.SECONDSPERSLOT(&_WithdrawOracle.CallOpts)
}

// ClBalances is a free data retrieval call binding the contract method 0x86c0c742.
//
// Solidity: function clBalances() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCaller) ClBalances(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "clBalances")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClBalances is a free data retrieval call binding the contract method 0x86c0c742.
//
// Solidity: function clBalances() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleSession) ClBalances() (*big.Int, error) {
	return _WithdrawOracle.Contract.ClBalances(&_WithdrawOracle.CallOpts)
}

// ClBalances is a free data retrieval call binding the contract method 0x86c0c742.
//
// Solidity: function clBalances() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCallerSession) ClBalances() (*big.Int, error) {
	return _WithdrawOracle.Contract.ClBalances(&_WithdrawOracle.CallOpts)
}

// ClVaultBalance is a free data retrieval call binding the contract method 0x312adfb8.
//
// Solidity: function clVaultBalance() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCaller) ClVaultBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "clVaultBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClVaultBalance is a free data retrieval call binding the contract method 0x312adfb8.
//
// Solidity: function clVaultBalance() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleSession) ClVaultBalance() (*big.Int, error) {
	return _WithdrawOracle.Contract.ClVaultBalance(&_WithdrawOracle.CallOpts)
}

// ClVaultBalance is a free data retrieval call binding the contract method 0x312adfb8.
//
// Solidity: function clVaultBalance() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCallerSession) ClVaultBalance() (*big.Int, error) {
	return _WithdrawOracle.Contract.ClVaultBalance(&_WithdrawOracle.CallOpts)
}

// ClVaultMinSettleLimit is a free data retrieval call binding the contract method 0xbde5a619.
//
// Solidity: function clVaultMinSettleLimit() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCaller) ClVaultMinSettleLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "clVaultMinSettleLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClVaultMinSettleLimit is a free data retrieval call binding the contract method 0xbde5a619.
//
// Solidity: function clVaultMinSettleLimit() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleSession) ClVaultMinSettleLimit() (*big.Int, error) {
	return _WithdrawOracle.Contract.ClVaultMinSettleLimit(&_WithdrawOracle.CallOpts)
}

// ClVaultMinSettleLimit is a free data retrieval call binding the contract method 0xbde5a619.
//
// Solidity: function clVaultMinSettleLimit() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCallerSession) ClVaultMinSettleLimit() (*big.Int, error) {
	return _WithdrawOracle.Contract.ClVaultMinSettleLimit(&_WithdrawOracle.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_WithdrawOracle *WithdrawOracleCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_WithdrawOracle *WithdrawOracleSession) Dao() (common.Address, error) {
	return _WithdrawOracle.Contract.Dao(&_WithdrawOracle.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_WithdrawOracle *WithdrawOracleCallerSession) Dao() (common.Address, error) {
	return _WithdrawOracle.Contract.Dao(&_WithdrawOracle.CallOpts)
}

// ExitRequestLimit is a free data retrieval call binding the contract method 0x1f78be77.
//
// Solidity: function exitRequestLimit() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCaller) ExitRequestLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "exitRequestLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExitRequestLimit is a free data retrieval call binding the contract method 0x1f78be77.
//
// Solidity: function exitRequestLimit() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleSession) ExitRequestLimit() (*big.Int, error) {
	return _WithdrawOracle.Contract.ExitRequestLimit(&_WithdrawOracle.CallOpts)
}

// ExitRequestLimit is a free data retrieval call binding the contract method 0x1f78be77.
//
// Solidity: function exitRequestLimit() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCallerSession) ExitRequestLimit() (*big.Int, error) {
	return _WithdrawOracle.Contract.ExitRequestLimit(&_WithdrawOracle.CallOpts)
}

// GetClBalances is a free data retrieval call binding the contract method 0x0fd520f6.
//
// Solidity: function getClBalances() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCaller) GetClBalances(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "getClBalances")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClBalances is a free data retrieval call binding the contract method 0x0fd520f6.
//
// Solidity: function getClBalances() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleSession) GetClBalances() (*big.Int, error) {
	return _WithdrawOracle.Contract.GetClBalances(&_WithdrawOracle.CallOpts)
}

// GetClBalances is a free data retrieval call binding the contract method 0x0fd520f6.
//
// Solidity: function getClBalances() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCallerSession) GetClBalances() (*big.Int, error) {
	return _WithdrawOracle.Contract.GetClBalances(&_WithdrawOracle.CallOpts)
}

// GetClVaultBalances is a free data retrieval call binding the contract method 0x67fe1305.
//
// Solidity: function getClVaultBalances() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCaller) GetClVaultBalances(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "getClVaultBalances")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClVaultBalances is a free data retrieval call binding the contract method 0x67fe1305.
//
// Solidity: function getClVaultBalances() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleSession) GetClVaultBalances() (*big.Int, error) {
	return _WithdrawOracle.Contract.GetClVaultBalances(&_WithdrawOracle.CallOpts)
}

// GetClVaultBalances is a free data retrieval call binding the contract method 0x67fe1305.
//
// Solidity: function getClVaultBalances() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCallerSession) GetClVaultBalances() (*big.Int, error) {
	return _WithdrawOracle.Contract.GetClVaultBalances(&_WithdrawOracle.CallOpts)
}

// GetConsensusContract is a free data retrieval call binding the contract method 0x8f55b571.
//
// Solidity: function getConsensusContract() view returns(address)
func (_WithdrawOracle *WithdrawOracleCaller) GetConsensusContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "getConsensusContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetConsensusContract is a free data retrieval call binding the contract method 0x8f55b571.
//
// Solidity: function getConsensusContract() view returns(address)
func (_WithdrawOracle *WithdrawOracleSession) GetConsensusContract() (common.Address, error) {
	return _WithdrawOracle.Contract.GetConsensusContract(&_WithdrawOracle.CallOpts)
}

// GetConsensusContract is a free data retrieval call binding the contract method 0x8f55b571.
//
// Solidity: function getConsensusContract() view returns(address)
func (_WithdrawOracle *WithdrawOracleCallerSession) GetConsensusContract() (common.Address, error) {
	return _WithdrawOracle.Contract.GetConsensusContract(&_WithdrawOracle.CallOpts)
}

// GetConsensusReport is a free data retrieval call binding the contract method 0x60d64d38.
//
// Solidity: function getConsensusReport() view returns(bytes32 hash, uint256 refSlot, uint256 processingDeadlineTime, bool processingStarted)
func (_WithdrawOracle *WithdrawOracleCaller) GetConsensusReport(opts *bind.CallOpts) (struct {
	Hash                   [32]byte
	RefSlot                *big.Int
	ProcessingDeadlineTime *big.Int
	ProcessingStarted      bool
}, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "getConsensusReport")

	outstruct := new(struct {
		Hash                   [32]byte
		RefSlot                *big.Int
		ProcessingDeadlineTime *big.Int
		ProcessingStarted      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.RefSlot = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ProcessingDeadlineTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ProcessingStarted = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// GetConsensusReport is a free data retrieval call binding the contract method 0x60d64d38.
//
// Solidity: function getConsensusReport() view returns(bytes32 hash, uint256 refSlot, uint256 processingDeadlineTime, bool processingStarted)
func (_WithdrawOracle *WithdrawOracleSession) GetConsensusReport() (struct {
	Hash                   [32]byte
	RefSlot                *big.Int
	ProcessingDeadlineTime *big.Int
	ProcessingStarted      bool
}, error) {
	return _WithdrawOracle.Contract.GetConsensusReport(&_WithdrawOracle.CallOpts)
}

// GetConsensusReport is a free data retrieval call binding the contract method 0x60d64d38.
//
// Solidity: function getConsensusReport() view returns(bytes32 hash, uint256 refSlot, uint256 processingDeadlineTime, bool processingStarted)
func (_WithdrawOracle *WithdrawOracleCallerSession) GetConsensusReport() (struct {
	Hash                   [32]byte
	RefSlot                *big.Int
	ProcessingDeadlineTime *big.Int
	ProcessingStarted      bool
}, error) {
	return _WithdrawOracle.Contract.GetConsensusReport(&_WithdrawOracle.CallOpts)
}

// GetConsensusVersion is a free data retrieval call binding the contract method 0x5be20425.
//
// Solidity: function getConsensusVersion() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCaller) GetConsensusVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "getConsensusVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConsensusVersion is a free data retrieval call binding the contract method 0x5be20425.
//
// Solidity: function getConsensusVersion() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleSession) GetConsensusVersion() (*big.Int, error) {
	return _WithdrawOracle.Contract.GetConsensusVersion(&_WithdrawOracle.CallOpts)
}

// GetConsensusVersion is a free data retrieval call binding the contract method 0x5be20425.
//
// Solidity: function getConsensusVersion() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCallerSession) GetConsensusVersion() (*big.Int, error) {
	return _WithdrawOracle.Contract.GetConsensusVersion(&_WithdrawOracle.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "getContractVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleSession) GetContractVersion() (*big.Int, error) {
	return _WithdrawOracle.Contract.GetContractVersion(&_WithdrawOracle.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCallerSession) GetContractVersion() (*big.Int, error) {
	return _WithdrawOracle.Contract.GetContractVersion(&_WithdrawOracle.CallOpts)
}

// GetLastClSettleAmount is a free data retrieval call binding the contract method 0xacfc654b.
//
// Solidity: function getLastClSettleAmount() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCaller) GetLastClSettleAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "getLastClSettleAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastClSettleAmount is a free data retrieval call binding the contract method 0xacfc654b.
//
// Solidity: function getLastClSettleAmount() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleSession) GetLastClSettleAmount() (*big.Int, error) {
	return _WithdrawOracle.Contract.GetLastClSettleAmount(&_WithdrawOracle.CallOpts)
}

// GetLastClSettleAmount is a free data retrieval call binding the contract method 0xacfc654b.
//
// Solidity: function getLastClSettleAmount() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCallerSession) GetLastClSettleAmount() (*big.Int, error) {
	return _WithdrawOracle.Contract.GetLastClSettleAmount(&_WithdrawOracle.CallOpts)
}

// GetLastProcessingRefSlot is a free data retrieval call binding the contract method 0x3584d59c.
//
// Solidity: function getLastProcessingRefSlot() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCaller) GetLastProcessingRefSlot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "getLastProcessingRefSlot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastProcessingRefSlot is a free data retrieval call binding the contract method 0x3584d59c.
//
// Solidity: function getLastProcessingRefSlot() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleSession) GetLastProcessingRefSlot() (*big.Int, error) {
	return _WithdrawOracle.Contract.GetLastProcessingRefSlot(&_WithdrawOracle.CallOpts)
}

// GetLastProcessingRefSlot is a free data retrieval call binding the contract method 0x3584d59c.
//
// Solidity: function getLastProcessingRefSlot() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCallerSession) GetLastProcessingRefSlot() (*big.Int, error) {
	return _WithdrawOracle.Contract.GetLastProcessingRefSlot(&_WithdrawOracle.CallOpts)
}

// GetPendingBalances is a free data retrieval call binding the contract method 0xf0b73e11.
//
// Solidity: function getPendingBalances() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCaller) GetPendingBalances(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "getPendingBalances")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingBalances is a free data retrieval call binding the contract method 0xf0b73e11.
//
// Solidity: function getPendingBalances() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleSession) GetPendingBalances() (*big.Int, error) {
	return _WithdrawOracle.Contract.GetPendingBalances(&_WithdrawOracle.CallOpts)
}

// GetPendingBalances is a free data retrieval call binding the contract method 0xf0b73e11.
//
// Solidity: function getPendingBalances() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCallerSession) GetPendingBalances() (*big.Int, error) {
	return _WithdrawOracle.Contract.GetPendingBalances(&_WithdrawOracle.CallOpts)
}

// GetProcessingState is a free data retrieval call binding the contract method 0x8f7797c2.
//
// Solidity: function getProcessingState() view returns((uint256,uint256,bytes32,bool,uint256) result)
func (_WithdrawOracle *WithdrawOracleCaller) GetProcessingState(opts *bind.CallOpts) (WithdrawOracleProcessingState, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "getProcessingState")

	if err != nil {
		return *new(WithdrawOracleProcessingState), err
	}

	out0 := *abi.ConvertType(out[0], new(WithdrawOracleProcessingState)).(*WithdrawOracleProcessingState)

	return out0, err

}

// GetProcessingState is a free data retrieval call binding the contract method 0x8f7797c2.
//
// Solidity: function getProcessingState() view returns((uint256,uint256,bytes32,bool,uint256) result)
func (_WithdrawOracle *WithdrawOracleSession) GetProcessingState() (WithdrawOracleProcessingState, error) {
	return _WithdrawOracle.Contract.GetProcessingState(&_WithdrawOracle.CallOpts)
}

// GetProcessingState is a free data retrieval call binding the contract method 0x8f7797c2.
//
// Solidity: function getProcessingState() view returns((uint256,uint256,bytes32,bool,uint256) result)
func (_WithdrawOracle *WithdrawOracleCallerSession) GetProcessingState() (WithdrawOracleProcessingState, error) {
	return _WithdrawOracle.Contract.GetProcessingState(&_WithdrawOracle.CallOpts)
}

// LastClSettleAmount is a free data retrieval call binding the contract method 0x2f92553e.
//
// Solidity: function lastClSettleAmount() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCaller) LastClSettleAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "lastClSettleAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastClSettleAmount is a free data retrieval call binding the contract method 0x2f92553e.
//
// Solidity: function lastClSettleAmount() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleSession) LastClSettleAmount() (*big.Int, error) {
	return _WithdrawOracle.Contract.LastClSettleAmount(&_WithdrawOracle.CallOpts)
}

// LastClSettleAmount is a free data retrieval call binding the contract method 0x2f92553e.
//
// Solidity: function lastClSettleAmount() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCallerSession) LastClSettleAmount() (*big.Int, error) {
	return _WithdrawOracle.Contract.LastClSettleAmount(&_WithdrawOracle.CallOpts)
}

// LastRefSlot is a free data retrieval call binding the contract method 0x866f051f.
//
// Solidity: function lastRefSlot() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCaller) LastRefSlot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "lastRefSlot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRefSlot is a free data retrieval call binding the contract method 0x866f051f.
//
// Solidity: function lastRefSlot() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleSession) LastRefSlot() (*big.Int, error) {
	return _WithdrawOracle.Contract.LastRefSlot(&_WithdrawOracle.CallOpts)
}

// LastRefSlot is a free data retrieval call binding the contract method 0x866f051f.
//
// Solidity: function lastRefSlot() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCallerSession) LastRefSlot() (*big.Int, error) {
	return _WithdrawOracle.Contract.LastRefSlot(&_WithdrawOracle.CallOpts)
}

// LiquidStakingContractAddress is a free data retrieval call binding the contract method 0x6404a4c7.
//
// Solidity: function liquidStakingContractAddress() view returns(address)
func (_WithdrawOracle *WithdrawOracleCaller) LiquidStakingContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "liquidStakingContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidStakingContractAddress is a free data retrieval call binding the contract method 0x6404a4c7.
//
// Solidity: function liquidStakingContractAddress() view returns(address)
func (_WithdrawOracle *WithdrawOracleSession) LiquidStakingContractAddress() (common.Address, error) {
	return _WithdrawOracle.Contract.LiquidStakingContractAddress(&_WithdrawOracle.CallOpts)
}

// LiquidStakingContractAddress is a free data retrieval call binding the contract method 0x6404a4c7.
//
// Solidity: function liquidStakingContractAddress() view returns(address)
func (_WithdrawOracle *WithdrawOracleCallerSession) LiquidStakingContractAddress() (common.Address, error) {
	return _WithdrawOracle.Contract.LiquidStakingContractAddress(&_WithdrawOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WithdrawOracle *WithdrawOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WithdrawOracle *WithdrawOracleSession) Owner() (common.Address, error) {
	return _WithdrawOracle.Contract.Owner(&_WithdrawOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WithdrawOracle *WithdrawOracleCallerSession) Owner() (common.Address, error) {
	return _WithdrawOracle.Contract.Owner(&_WithdrawOracle.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WithdrawOracle *WithdrawOracleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WithdrawOracle *WithdrawOracleSession) Paused() (bool, error) {
	return _WithdrawOracle.Contract.Paused(&_WithdrawOracle.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WithdrawOracle *WithdrawOracleCallerSession) Paused() (bool, error) {
	return _WithdrawOracle.Contract.Paused(&_WithdrawOracle.CallOpts)
}

// PendingBalances is a free data retrieval call binding the contract method 0xa65d1dbc.
//
// Solidity: function pendingBalances() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCaller) PendingBalances(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "pendingBalances")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingBalances is a free data retrieval call binding the contract method 0xa65d1dbc.
//
// Solidity: function pendingBalances() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleSession) PendingBalances() (*big.Int, error) {
	return _WithdrawOracle.Contract.PendingBalances(&_WithdrawOracle.CallOpts)
}

// PendingBalances is a free data retrieval call binding the contract method 0xa65d1dbc.
//
// Solidity: function pendingBalances() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCallerSession) PendingBalances() (*big.Int, error) {
	return _WithdrawOracle.Contract.PendingBalances(&_WithdrawOracle.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WithdrawOracle *WithdrawOracleCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WithdrawOracle *WithdrawOracleSession) ProxiableUUID() ([32]byte, error) {
	return _WithdrawOracle.Contract.ProxiableUUID(&_WithdrawOracle.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WithdrawOracle *WithdrawOracleCallerSession) ProxiableUUID() ([32]byte, error) {
	return _WithdrawOracle.Contract.ProxiableUUID(&_WithdrawOracle.CallOpts)
}

// TotalBalanceTolerate is a free data retrieval call binding the contract method 0x4bd80f6f.
//
// Solidity: function totalBalanceTolerate() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCaller) TotalBalanceTolerate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "totalBalanceTolerate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBalanceTolerate is a free data retrieval call binding the contract method 0x4bd80f6f.
//
// Solidity: function totalBalanceTolerate() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleSession) TotalBalanceTolerate() (*big.Int, error) {
	return _WithdrawOracle.Contract.TotalBalanceTolerate(&_WithdrawOracle.CallOpts)
}

// TotalBalanceTolerate is a free data retrieval call binding the contract method 0x4bd80f6f.
//
// Solidity: function totalBalanceTolerate() view returns(uint256)
func (_WithdrawOracle *WithdrawOracleCallerSession) TotalBalanceTolerate() (*big.Int, error) {
	return _WithdrawOracle.Contract.TotalBalanceTolerate(&_WithdrawOracle.CallOpts)
}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_WithdrawOracle *WithdrawOracleCaller) VaultManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WithdrawOracle.contract.Call(opts, &out, "vaultManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_WithdrawOracle *WithdrawOracleSession) VaultManager() (common.Address, error) {
	return _WithdrawOracle.Contract.VaultManager(&_WithdrawOracle.CallOpts)
}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_WithdrawOracle *WithdrawOracleCallerSession) VaultManager() (common.Address, error) {
	return _WithdrawOracle.Contract.VaultManager(&_WithdrawOracle.CallOpts)
}

// AddPendingBalances is a paid mutator transaction binding the contract method 0x8e34ead0.
//
// Solidity: function addPendingBalances(uint256 _pendingBalance) returns()
func (_WithdrawOracle *WithdrawOracleTransactor) AddPendingBalances(opts *bind.TransactOpts, _pendingBalance *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "addPendingBalances", _pendingBalance)
}

// AddPendingBalances is a paid mutator transaction binding the contract method 0x8e34ead0.
//
// Solidity: function addPendingBalances(uint256 _pendingBalance) returns()
func (_WithdrawOracle *WithdrawOracleSession) AddPendingBalances(_pendingBalance *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.AddPendingBalances(&_WithdrawOracle.TransactOpts, _pendingBalance)
}

// AddPendingBalances is a paid mutator transaction binding the contract method 0x8e34ead0.
//
// Solidity: function addPendingBalances(uint256 _pendingBalance) returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) AddPendingBalances(_pendingBalance *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.AddPendingBalances(&_WithdrawOracle.TransactOpts, _pendingBalance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc366aaf3.
//
// Solidity: function initialize(uint256 secondsPerSlot, uint256 genesisTime, address consensusContract, uint256 consensusVersion, uint256 lastProcessingRefSlot, address _dao, uint256 _exitRequestLimit, uint256 _clVaultMinSettleLimit, uint256 _clBalance, uint256 _pendingBalance) returns()
func (_WithdrawOracle *WithdrawOracleTransactor) Initialize(opts *bind.TransactOpts, secondsPerSlot *big.Int, genesisTime *big.Int, consensusContract common.Address, consensusVersion *big.Int, lastProcessingRefSlot *big.Int, _dao common.Address, _exitRequestLimit *big.Int, _clVaultMinSettleLimit *big.Int, _clBalance *big.Int, _pendingBalance *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "initialize", secondsPerSlot, genesisTime, consensusContract, consensusVersion, lastProcessingRefSlot, _dao, _exitRequestLimit, _clVaultMinSettleLimit, _clBalance, _pendingBalance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc366aaf3.
//
// Solidity: function initialize(uint256 secondsPerSlot, uint256 genesisTime, address consensusContract, uint256 consensusVersion, uint256 lastProcessingRefSlot, address _dao, uint256 _exitRequestLimit, uint256 _clVaultMinSettleLimit, uint256 _clBalance, uint256 _pendingBalance) returns()
func (_WithdrawOracle *WithdrawOracleSession) Initialize(secondsPerSlot *big.Int, genesisTime *big.Int, consensusContract common.Address, consensusVersion *big.Int, lastProcessingRefSlot *big.Int, _dao common.Address, _exitRequestLimit *big.Int, _clVaultMinSettleLimit *big.Int, _clBalance *big.Int, _pendingBalance *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.Initialize(&_WithdrawOracle.TransactOpts, secondsPerSlot, genesisTime, consensusContract, consensusVersion, lastProcessingRefSlot, _dao, _exitRequestLimit, _clVaultMinSettleLimit, _clBalance, _pendingBalance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc366aaf3.
//
// Solidity: function initialize(uint256 secondsPerSlot, uint256 genesisTime, address consensusContract, uint256 consensusVersion, uint256 lastProcessingRefSlot, address _dao, uint256 _exitRequestLimit, uint256 _clVaultMinSettleLimit, uint256 _clBalance, uint256 _pendingBalance) returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) Initialize(secondsPerSlot *big.Int, genesisTime *big.Int, consensusContract common.Address, consensusVersion *big.Int, lastProcessingRefSlot *big.Int, _dao common.Address, _exitRequestLimit *big.Int, _clVaultMinSettleLimit *big.Int, _clBalance *big.Int, _pendingBalance *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.Initialize(&_WithdrawOracle.TransactOpts, secondsPerSlot, genesisTime, consensusContract, consensusVersion, lastProcessingRefSlot, _dao, _exitRequestLimit, _clVaultMinSettleLimit, _clBalance, _pendingBalance)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0xccd8af22.
//
// Solidity: function initializeV2(address _consensus, uint256 _lastProcessingRefSlot) returns()
func (_WithdrawOracle *WithdrawOracleTransactor) InitializeV2(opts *bind.TransactOpts, _consensus common.Address, _lastProcessingRefSlot *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "initializeV2", _consensus, _lastProcessingRefSlot)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0xccd8af22.
//
// Solidity: function initializeV2(address _consensus, uint256 _lastProcessingRefSlot) returns()
func (_WithdrawOracle *WithdrawOracleSession) InitializeV2(_consensus common.Address, _lastProcessingRefSlot *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.InitializeV2(&_WithdrawOracle.TransactOpts, _consensus, _lastProcessingRefSlot)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0xccd8af22.
//
// Solidity: function initializeV2(address _consensus, uint256 _lastProcessingRefSlot) returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) InitializeV2(_consensus common.Address, _lastProcessingRefSlot *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.InitializeV2(&_WithdrawOracle.TransactOpts, _consensus, _lastProcessingRefSlot)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WithdrawOracle *WithdrawOracleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WithdrawOracle *WithdrawOracleSession) Pause() (*types.Transaction, error) {
	return _WithdrawOracle.Contract.Pause(&_WithdrawOracle.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) Pause() (*types.Transaction, error) {
	return _WithdrawOracle.Contract.Pause(&_WithdrawOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WithdrawOracle *WithdrawOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WithdrawOracle *WithdrawOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _WithdrawOracle.Contract.RenounceOwnership(&_WithdrawOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WithdrawOracle.Contract.RenounceOwnership(&_WithdrawOracle.TransactOpts)
}

// SetClVaultMinSettleLimit is a paid mutator transaction binding the contract method 0xc14694e3.
//
// Solidity: function setClVaultMinSettleLimit(uint256 _clVaultMinSettleLimit) returns()
func (_WithdrawOracle *WithdrawOracleTransactor) SetClVaultMinSettleLimit(opts *bind.TransactOpts, _clVaultMinSettleLimit *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "setClVaultMinSettleLimit", _clVaultMinSettleLimit)
}

// SetClVaultMinSettleLimit is a paid mutator transaction binding the contract method 0xc14694e3.
//
// Solidity: function setClVaultMinSettleLimit(uint256 _clVaultMinSettleLimit) returns()
func (_WithdrawOracle *WithdrawOracleSession) SetClVaultMinSettleLimit(_clVaultMinSettleLimit *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SetClVaultMinSettleLimit(&_WithdrawOracle.TransactOpts, _clVaultMinSettleLimit)
}

// SetClVaultMinSettleLimit is a paid mutator transaction binding the contract method 0xc14694e3.
//
// Solidity: function setClVaultMinSettleLimit(uint256 _clVaultMinSettleLimit) returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) SetClVaultMinSettleLimit(_clVaultMinSettleLimit *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SetClVaultMinSettleLimit(&_WithdrawOracle.TransactOpts, _clVaultMinSettleLimit)
}

// SetConsensusContract is a paid mutator transaction binding the contract method 0xc469c307.
//
// Solidity: function setConsensusContract(address addr) returns()
func (_WithdrawOracle *WithdrawOracleTransactor) SetConsensusContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "setConsensusContract", addr)
}

// SetConsensusContract is a paid mutator transaction binding the contract method 0xc469c307.
//
// Solidity: function setConsensusContract(address addr) returns()
func (_WithdrawOracle *WithdrawOracleSession) SetConsensusContract(addr common.Address) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SetConsensusContract(&_WithdrawOracle.TransactOpts, addr)
}

// SetConsensusContract is a paid mutator transaction binding the contract method 0xc469c307.
//
// Solidity: function setConsensusContract(address addr) returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) SetConsensusContract(addr common.Address) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SetConsensusContract(&_WithdrawOracle.TransactOpts, addr)
}

// SetConsensusVersion is a paid mutator transaction binding the contract method 0x8d591474.
//
// Solidity: function setConsensusVersion(uint256 version) returns()
func (_WithdrawOracle *WithdrawOracleTransactor) SetConsensusVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "setConsensusVersion", version)
}

// SetConsensusVersion is a paid mutator transaction binding the contract method 0x8d591474.
//
// Solidity: function setConsensusVersion(uint256 version) returns()
func (_WithdrawOracle *WithdrawOracleSession) SetConsensusVersion(version *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SetConsensusVersion(&_WithdrawOracle.TransactOpts, version)
}

// SetConsensusVersion is a paid mutator transaction binding the contract method 0x8d591474.
//
// Solidity: function setConsensusVersion(uint256 version) returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) SetConsensusVersion(version *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SetConsensusVersion(&_WithdrawOracle.TransactOpts, version)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_WithdrawOracle *WithdrawOracleTransactor) SetDaoAddress(opts *bind.TransactOpts, _dao common.Address) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "setDaoAddress", _dao)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_WithdrawOracle *WithdrawOracleSession) SetDaoAddress(_dao common.Address) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SetDaoAddress(&_WithdrawOracle.TransactOpts, _dao)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) SetDaoAddress(_dao common.Address) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SetDaoAddress(&_WithdrawOracle.TransactOpts, _dao)
}

// SetExitRequestLimit is a paid mutator transaction binding the contract method 0x3936e707.
//
// Solidity: function setExitRequestLimit(uint256 _exitRequestLimit) returns()
func (_WithdrawOracle *WithdrawOracleTransactor) SetExitRequestLimit(opts *bind.TransactOpts, _exitRequestLimit *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "setExitRequestLimit", _exitRequestLimit)
}

// SetExitRequestLimit is a paid mutator transaction binding the contract method 0x3936e707.
//
// Solidity: function setExitRequestLimit(uint256 _exitRequestLimit) returns()
func (_WithdrawOracle *WithdrawOracleSession) SetExitRequestLimit(_exitRequestLimit *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SetExitRequestLimit(&_WithdrawOracle.TransactOpts, _exitRequestLimit)
}

// SetExitRequestLimit is a paid mutator transaction binding the contract method 0x3936e707.
//
// Solidity: function setExitRequestLimit(uint256 _exitRequestLimit) returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) SetExitRequestLimit(_exitRequestLimit *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SetExitRequestLimit(&_WithdrawOracle.TransactOpts, _exitRequestLimit)
}

// SetLiquidStaking is a paid mutator transaction binding the contract method 0x08211be5.
//
// Solidity: function setLiquidStaking(address _liquidStakingContractAddress) returns()
func (_WithdrawOracle *WithdrawOracleTransactor) SetLiquidStaking(opts *bind.TransactOpts, _liquidStakingContractAddress common.Address) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "setLiquidStaking", _liquidStakingContractAddress)
}

// SetLiquidStaking is a paid mutator transaction binding the contract method 0x08211be5.
//
// Solidity: function setLiquidStaking(address _liquidStakingContractAddress) returns()
func (_WithdrawOracle *WithdrawOracleSession) SetLiquidStaking(_liquidStakingContractAddress common.Address) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SetLiquidStaking(&_WithdrawOracle.TransactOpts, _liquidStakingContractAddress)
}

// SetLiquidStaking is a paid mutator transaction binding the contract method 0x08211be5.
//
// Solidity: function setLiquidStaking(address _liquidStakingContractAddress) returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) SetLiquidStaking(_liquidStakingContractAddress common.Address) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SetLiquidStaking(&_WithdrawOracle.TransactOpts, _liquidStakingContractAddress)
}

// SetTotalBalanceTolerate is a paid mutator transaction binding the contract method 0x43532f48.
//
// Solidity: function setTotalBalanceTolerate(uint256 _totalBalanceTolerate) returns()
func (_WithdrawOracle *WithdrawOracleTransactor) SetTotalBalanceTolerate(opts *bind.TransactOpts, _totalBalanceTolerate *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "setTotalBalanceTolerate", _totalBalanceTolerate)
}

// SetTotalBalanceTolerate is a paid mutator transaction binding the contract method 0x43532f48.
//
// Solidity: function setTotalBalanceTolerate(uint256 _totalBalanceTolerate) returns()
func (_WithdrawOracle *WithdrawOracleSession) SetTotalBalanceTolerate(_totalBalanceTolerate *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SetTotalBalanceTolerate(&_WithdrawOracle.TransactOpts, _totalBalanceTolerate)
}

// SetTotalBalanceTolerate is a paid mutator transaction binding the contract method 0x43532f48.
//
// Solidity: function setTotalBalanceTolerate(uint256 _totalBalanceTolerate) returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) SetTotalBalanceTolerate(_totalBalanceTolerate *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SetTotalBalanceTolerate(&_WithdrawOracle.TransactOpts, _totalBalanceTolerate)
}

// SetVaultManager is a paid mutator transaction binding the contract method 0xb543503e.
//
// Solidity: function setVaultManager(address _vaultManagerContractAddress) returns()
func (_WithdrawOracle *WithdrawOracleTransactor) SetVaultManager(opts *bind.TransactOpts, _vaultManagerContractAddress common.Address) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "setVaultManager", _vaultManagerContractAddress)
}

// SetVaultManager is a paid mutator transaction binding the contract method 0xb543503e.
//
// Solidity: function setVaultManager(address _vaultManagerContractAddress) returns()
func (_WithdrawOracle *WithdrawOracleSession) SetVaultManager(_vaultManagerContractAddress common.Address) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SetVaultManager(&_WithdrawOracle.TransactOpts, _vaultManagerContractAddress)
}

// SetVaultManager is a paid mutator transaction binding the contract method 0xb543503e.
//
// Solidity: function setVaultManager(address _vaultManagerContractAddress) returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) SetVaultManager(_vaultManagerContractAddress common.Address) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SetVaultManager(&_WithdrawOracle.TransactOpts, _vaultManagerContractAddress)
}

// SubmitConsensusReport is a paid mutator transaction binding the contract method 0x7bcc792b.
//
// Solidity: function submitConsensusReport(bytes32 reportHash, uint256 refSlot, uint256 deadline, uint256 _moduleId) returns()
func (_WithdrawOracle *WithdrawOracleTransactor) SubmitConsensusReport(opts *bind.TransactOpts, reportHash [32]byte, refSlot *big.Int, deadline *big.Int, _moduleId *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "submitConsensusReport", reportHash, refSlot, deadline, _moduleId)
}

// SubmitConsensusReport is a paid mutator transaction binding the contract method 0x7bcc792b.
//
// Solidity: function submitConsensusReport(bytes32 reportHash, uint256 refSlot, uint256 deadline, uint256 _moduleId) returns()
func (_WithdrawOracle *WithdrawOracleSession) SubmitConsensusReport(reportHash [32]byte, refSlot *big.Int, deadline *big.Int, _moduleId *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SubmitConsensusReport(&_WithdrawOracle.TransactOpts, reportHash, refSlot, deadline, _moduleId)
}

// SubmitConsensusReport is a paid mutator transaction binding the contract method 0x7bcc792b.
//
// Solidity: function submitConsensusReport(bytes32 reportHash, uint256 refSlot, uint256 deadline, uint256 _moduleId) returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) SubmitConsensusReport(reportHash [32]byte, refSlot *big.Int, deadline *big.Int, _moduleId *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SubmitConsensusReport(&_WithdrawOracle.TransactOpts, reportHash, refSlot, deadline, _moduleId)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0x5588b8c8.
//
// Solidity: function submitReportData((uint256,uint256,uint256,uint256,uint256,uint256,(uint64,uint96,uint96)[],(uint64,uint96,uint96)[],uint256[],uint256[]) data, uint256 _contractVersion, uint256 _moduleId) returns()
func (_WithdrawOracle *WithdrawOracleTransactor) SubmitReportData(opts *bind.TransactOpts, data WithdrawOracleReportData, _contractVersion *big.Int, _moduleId *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "submitReportData", data, _contractVersion, _moduleId)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0x5588b8c8.
//
// Solidity: function submitReportData((uint256,uint256,uint256,uint256,uint256,uint256,(uint64,uint96,uint96)[],(uint64,uint96,uint96)[],uint256[],uint256[]) data, uint256 _contractVersion, uint256 _moduleId) returns()
func (_WithdrawOracle *WithdrawOracleSession) SubmitReportData(data WithdrawOracleReportData, _contractVersion *big.Int, _moduleId *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SubmitReportData(&_WithdrawOracle.TransactOpts, data, _contractVersion, _moduleId)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0x5588b8c8.
//
// Solidity: function submitReportData((uint256,uint256,uint256,uint256,uint256,uint256,(uint64,uint96,uint96)[],(uint64,uint96,uint96)[],uint256[],uint256[]) data, uint256 _contractVersion, uint256 _moduleId) returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) SubmitReportData(data WithdrawOracleReportData, _contractVersion *big.Int, _moduleId *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.SubmitReportData(&_WithdrawOracle.TransactOpts, data, _contractVersion, _moduleId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WithdrawOracle *WithdrawOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WithdrawOracle *WithdrawOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.TransferOwnership(&_WithdrawOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.TransferOwnership(&_WithdrawOracle.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WithdrawOracle *WithdrawOracleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WithdrawOracle *WithdrawOracleSession) Unpause() (*types.Transaction, error) {
	return _WithdrawOracle.Contract.Unpause(&_WithdrawOracle.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) Unpause() (*types.Transaction, error) {
	return _WithdrawOracle.Contract.Unpause(&_WithdrawOracle.TransactOpts)
}

// UpdateContractVersion is a paid mutator transaction binding the contract method 0x294bb79a.
//
// Solidity: function updateContractVersion(uint256 version) returns()
func (_WithdrawOracle *WithdrawOracleTransactor) UpdateContractVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "updateContractVersion", version)
}

// UpdateContractVersion is a paid mutator transaction binding the contract method 0x294bb79a.
//
// Solidity: function updateContractVersion(uint256 version) returns()
func (_WithdrawOracle *WithdrawOracleSession) UpdateContractVersion(version *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.UpdateContractVersion(&_WithdrawOracle.TransactOpts, version)
}

// UpdateContractVersion is a paid mutator transaction binding the contract method 0x294bb79a.
//
// Solidity: function updateContractVersion(uint256 version) returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) UpdateContractVersion(version *big.Int) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.UpdateContractVersion(&_WithdrawOracle.TransactOpts, version)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WithdrawOracle *WithdrawOracleTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WithdrawOracle *WithdrawOracleSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.UpgradeTo(&_WithdrawOracle.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.UpgradeTo(&_WithdrawOracle.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WithdrawOracle *WithdrawOracleTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WithdrawOracle.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WithdrawOracle *WithdrawOracleSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.UpgradeToAndCall(&_WithdrawOracle.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WithdrawOracle *WithdrawOracleTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WithdrawOracle.Contract.UpgradeToAndCall(&_WithdrawOracle.TransactOpts, newImplementation, data)
}

// WithdrawOracleAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the WithdrawOracle contract.
type WithdrawOracleAdminChangedIterator struct {
	Event *WithdrawOracleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleAdminChanged represents a AdminChanged event raised by the WithdrawOracle contract.
type WithdrawOracleAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*WithdrawOracleAdminChangedIterator, error) {

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleAdminChangedIterator{contract: _WithdrawOracle.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *WithdrawOracleAdminChanged) (event.Subscription, error) {

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleAdminChanged)
				if err := _WithdrawOracle.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseAdminChanged(log types.Log) (*WithdrawOracleAdminChanged, error) {
	event := new(WithdrawOracleAdminChanged)
	if err := _WithdrawOracle.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the WithdrawOracle contract.
type WithdrawOracleBeaconUpgradedIterator struct {
	Event *WithdrawOracleBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleBeaconUpgraded represents a BeaconUpgraded event raised by the WithdrawOracle contract.
type WithdrawOracleBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*WithdrawOracleBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleBeaconUpgradedIterator{contract: _WithdrawOracle.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *WithdrawOracleBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleBeaconUpgraded)
				if err := _WithdrawOracle.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseBeaconUpgraded(log types.Log) (*WithdrawOracleBeaconUpgraded, error) {
	event := new(WithdrawOracleBeaconUpgraded)
	if err := _WithdrawOracle.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleConsensusHashContractSetIterator is returned from FilterConsensusHashContractSet and is used to iterate over the raw logs and unpacked data for ConsensusHashContractSet events raised by the WithdrawOracle contract.
type WithdrawOracleConsensusHashContractSetIterator struct {
	Event *WithdrawOracleConsensusHashContractSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleConsensusHashContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleConsensusHashContractSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleConsensusHashContractSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleConsensusHashContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleConsensusHashContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleConsensusHashContractSet represents a ConsensusHashContractSet event raised by the WithdrawOracle contract.
type WithdrawOracleConsensusHashContractSet struct {
	Addr     common.Address
	PrevAddr common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConsensusHashContractSet is a free log retrieval operation binding the contract event 0x25421480fb7f52d18947876279a213696b58d7e0e5416ce5e2c9f9942661c34c.
//
// Solidity: event ConsensusHashContractSet(address indexed addr, address indexed prevAddr)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterConsensusHashContractSet(opts *bind.FilterOpts, addr []common.Address, prevAddr []common.Address) (*WithdrawOracleConsensusHashContractSetIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var prevAddrRule []interface{}
	for _, prevAddrItem := range prevAddr {
		prevAddrRule = append(prevAddrRule, prevAddrItem)
	}

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "ConsensusHashContractSet", addrRule, prevAddrRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleConsensusHashContractSetIterator{contract: _WithdrawOracle.contract, event: "ConsensusHashContractSet", logs: logs, sub: sub}, nil
}

// WatchConsensusHashContractSet is a free log subscription operation binding the contract event 0x25421480fb7f52d18947876279a213696b58d7e0e5416ce5e2c9f9942661c34c.
//
// Solidity: event ConsensusHashContractSet(address indexed addr, address indexed prevAddr)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchConsensusHashContractSet(opts *bind.WatchOpts, sink chan<- *WithdrawOracleConsensusHashContractSet, addr []common.Address, prevAddr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var prevAddrRule []interface{}
	for _, prevAddrItem := range prevAddr {
		prevAddrRule = append(prevAddrRule, prevAddrItem)
	}

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "ConsensusHashContractSet", addrRule, prevAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleConsensusHashContractSet)
				if err := _WithdrawOracle.contract.UnpackLog(event, "ConsensusHashContractSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConsensusHashContractSet is a log parse operation binding the contract event 0x25421480fb7f52d18947876279a213696b58d7e0e5416ce5e2c9f9942661c34c.
//
// Solidity: event ConsensusHashContractSet(address indexed addr, address indexed prevAddr)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseConsensusHashContractSet(log types.Log) (*WithdrawOracleConsensusHashContractSet, error) {
	event := new(WithdrawOracleConsensusHashContractSet)
	if err := _WithdrawOracle.contract.UnpackLog(event, "ConsensusHashContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleConsensusVersionSetIterator is returned from FilterConsensusVersionSet and is used to iterate over the raw logs and unpacked data for ConsensusVersionSet events raised by the WithdrawOracle contract.
type WithdrawOracleConsensusVersionSetIterator struct {
	Event *WithdrawOracleConsensusVersionSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleConsensusVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleConsensusVersionSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleConsensusVersionSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleConsensusVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleConsensusVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleConsensusVersionSet represents a ConsensusVersionSet event raised by the WithdrawOracle contract.
type WithdrawOracleConsensusVersionSet struct {
	Version     *big.Int
	PrevVersion *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsensusVersionSet is a free log retrieval operation binding the contract event 0xfa5304972d4ec3e3207f0bbf91155a49d0dfa62488f9529403a2a49e4b29a895.
//
// Solidity: event ConsensusVersionSet(uint256 indexed version, uint256 indexed prevVersion)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterConsensusVersionSet(opts *bind.FilterOpts, version []*big.Int, prevVersion []*big.Int) (*WithdrawOracleConsensusVersionSetIterator, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var prevVersionRule []interface{}
	for _, prevVersionItem := range prevVersion {
		prevVersionRule = append(prevVersionRule, prevVersionItem)
	}

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "ConsensusVersionSet", versionRule, prevVersionRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleConsensusVersionSetIterator{contract: _WithdrawOracle.contract, event: "ConsensusVersionSet", logs: logs, sub: sub}, nil
}

// WatchConsensusVersionSet is a free log subscription operation binding the contract event 0xfa5304972d4ec3e3207f0bbf91155a49d0dfa62488f9529403a2a49e4b29a895.
//
// Solidity: event ConsensusVersionSet(uint256 indexed version, uint256 indexed prevVersion)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchConsensusVersionSet(opts *bind.WatchOpts, sink chan<- *WithdrawOracleConsensusVersionSet, version []*big.Int, prevVersion []*big.Int) (event.Subscription, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var prevVersionRule []interface{}
	for _, prevVersionItem := range prevVersion {
		prevVersionRule = append(prevVersionRule, prevVersionItem)
	}

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "ConsensusVersionSet", versionRule, prevVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleConsensusVersionSet)
				if err := _WithdrawOracle.contract.UnpackLog(event, "ConsensusVersionSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConsensusVersionSet is a log parse operation binding the contract event 0xfa5304972d4ec3e3207f0bbf91155a49d0dfa62488f9529403a2a49e4b29a895.
//
// Solidity: event ConsensusVersionSet(uint256 indexed version, uint256 indexed prevVersion)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseConsensusVersionSet(log types.Log) (*WithdrawOracleConsensusVersionSet, error) {
	event := new(WithdrawOracleConsensusVersionSet)
	if err := _WithdrawOracle.contract.UnpackLog(event, "ConsensusVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleContractVersionSetIterator is returned from FilterContractVersionSet and is used to iterate over the raw logs and unpacked data for ContractVersionSet events raised by the WithdrawOracle contract.
type WithdrawOracleContractVersionSetIterator struct {
	Event *WithdrawOracleContractVersionSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleContractVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleContractVersionSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleContractVersionSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleContractVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleContractVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleContractVersionSet represents a ContractVersionSet event raised by the WithdrawOracle contract.
type WithdrawOracleContractVersionSet struct {
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractVersionSet is a free log retrieval operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterContractVersionSet(opts *bind.FilterOpts) (*WithdrawOracleContractVersionSetIterator, error) {

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleContractVersionSetIterator{contract: _WithdrawOracle.contract, event: "ContractVersionSet", logs: logs, sub: sub}, nil
}

// WatchContractVersionSet is a free log subscription operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchContractVersionSet(opts *bind.WatchOpts, sink chan<- *WithdrawOracleContractVersionSet) (event.Subscription, error) {

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleContractVersionSet)
				if err := _WithdrawOracle.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractVersionSet is a log parse operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseContractVersionSet(log types.Log) (*WithdrawOracleContractVersionSet, error) {
	event := new(WithdrawOracleContractVersionSet)
	if err := _WithdrawOracle.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleDaoAddressChangedIterator is returned from FilterDaoAddressChanged and is used to iterate over the raw logs and unpacked data for DaoAddressChanged events raised by the WithdrawOracle contract.
type WithdrawOracleDaoAddressChangedIterator struct {
	Event *WithdrawOracleDaoAddressChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleDaoAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleDaoAddressChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleDaoAddressChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleDaoAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleDaoAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleDaoAddressChanged represents a DaoAddressChanged event raised by the WithdrawOracle contract.
type WithdrawOracleDaoAddressChanged struct {
	OldDao common.Address
	Dao    common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDaoAddressChanged is a free log retrieval operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address _oldDao, address _dao)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterDaoAddressChanged(opts *bind.FilterOpts) (*WithdrawOracleDaoAddressChangedIterator, error) {

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "DaoAddressChanged")
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleDaoAddressChangedIterator{contract: _WithdrawOracle.contract, event: "DaoAddressChanged", logs: logs, sub: sub}, nil
}

// WatchDaoAddressChanged is a free log subscription operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address _oldDao, address _dao)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchDaoAddressChanged(opts *bind.WatchOpts, sink chan<- *WithdrawOracleDaoAddressChanged) (event.Subscription, error) {

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "DaoAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleDaoAddressChanged)
				if err := _WithdrawOracle.contract.UnpackLog(event, "DaoAddressChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDaoAddressChanged is a log parse operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address _oldDao, address _dao)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseDaoAddressChanged(log types.Log) (*WithdrawOracleDaoAddressChanged, error) {
	event := new(WithdrawOracleDaoAddressChanged)
	if err := _WithdrawOracle.contract.UnpackLog(event, "DaoAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WithdrawOracle contract.
type WithdrawOracleInitializedIterator struct {
	Event *WithdrawOracleInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleInitialized represents a Initialized event raised by the WithdrawOracle contract.
type WithdrawOracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*WithdrawOracleInitializedIterator, error) {

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleInitializedIterator{contract: _WithdrawOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WithdrawOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleInitialized)
				if err := _WithdrawOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseInitialized(log types.Log) (*WithdrawOracleInitialized, error) {
	event := new(WithdrawOracleInitialized)
	if err := _WithdrawOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleLiquidStakingChangedIterator is returned from FilterLiquidStakingChanged and is used to iterate over the raw logs and unpacked data for LiquidStakingChanged events raised by the WithdrawOracle contract.
type WithdrawOracleLiquidStakingChangedIterator struct {
	Event *WithdrawOracleLiquidStakingChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleLiquidStakingChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleLiquidStakingChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleLiquidStakingChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleLiquidStakingChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleLiquidStakingChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleLiquidStakingChanged represents a LiquidStakingChanged event raised by the WithdrawOracle contract.
type WithdrawOracleLiquidStakingChanged struct {
	OldLiq common.Address
	NewLiq common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLiquidStakingChanged is a free log retrieval operation binding the contract event 0x516e369f70685d2905d214a9e8567098c02a0e00f5a55bd30baca6b51d446cef.
//
// Solidity: event LiquidStakingChanged(address oldLiq, address newLiq)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterLiquidStakingChanged(opts *bind.FilterOpts) (*WithdrawOracleLiquidStakingChangedIterator, error) {

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "LiquidStakingChanged")
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleLiquidStakingChangedIterator{contract: _WithdrawOracle.contract, event: "LiquidStakingChanged", logs: logs, sub: sub}, nil
}

// WatchLiquidStakingChanged is a free log subscription operation binding the contract event 0x516e369f70685d2905d214a9e8567098c02a0e00f5a55bd30baca6b51d446cef.
//
// Solidity: event LiquidStakingChanged(address oldLiq, address newLiq)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchLiquidStakingChanged(opts *bind.WatchOpts, sink chan<- *WithdrawOracleLiquidStakingChanged) (event.Subscription, error) {

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "LiquidStakingChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleLiquidStakingChanged)
				if err := _WithdrawOracle.contract.UnpackLog(event, "LiquidStakingChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidStakingChanged is a log parse operation binding the contract event 0x516e369f70685d2905d214a9e8567098c02a0e00f5a55bd30baca6b51d446cef.
//
// Solidity: event LiquidStakingChanged(address oldLiq, address newLiq)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseLiquidStakingChanged(log types.Log) (*WithdrawOracleLiquidStakingChanged, error) {
	event := new(WithdrawOracleLiquidStakingChanged)
	if err := _WithdrawOracle.contract.UnpackLog(event, "LiquidStakingChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WithdrawOracle contract.
type WithdrawOracleOwnershipTransferredIterator struct {
	Event *WithdrawOracleOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleOwnershipTransferred represents a OwnershipTransferred event raised by the WithdrawOracle contract.
type WithdrawOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WithdrawOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleOwnershipTransferredIterator{contract: _WithdrawOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WithdrawOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleOwnershipTransferred)
				if err := _WithdrawOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseOwnershipTransferred(log types.Log) (*WithdrawOracleOwnershipTransferred, error) {
	event := new(WithdrawOracleOwnershipTransferred)
	if err := _WithdrawOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOraclePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the WithdrawOracle contract.
type WithdrawOraclePausedIterator struct {
	Event *WithdrawOraclePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOraclePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOraclePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOraclePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOraclePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOraclePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOraclePaused represents a Paused event raised by the WithdrawOracle contract.
type WithdrawOraclePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterPaused(opts *bind.FilterOpts) (*WithdrawOraclePausedIterator, error) {

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &WithdrawOraclePausedIterator{contract: _WithdrawOracle.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *WithdrawOraclePaused) (event.Subscription, error) {

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOraclePaused)
				if err := _WithdrawOracle.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WithdrawOracle *WithdrawOracleFilterer) ParsePaused(log types.Log) (*WithdrawOraclePaused, error) {
	event := new(WithdrawOraclePaused)
	if err := _WithdrawOracle.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOraclePendingBalancesAddIterator is returned from FilterPendingBalancesAdd and is used to iterate over the raw logs and unpacked data for PendingBalancesAdd events raised by the WithdrawOracle contract.
type WithdrawOraclePendingBalancesAddIterator struct {
	Event *WithdrawOraclePendingBalancesAdd // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOraclePendingBalancesAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOraclePendingBalancesAdd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOraclePendingBalancesAdd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOraclePendingBalancesAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOraclePendingBalancesAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOraclePendingBalancesAdd represents a PendingBalancesAdd event raised by the WithdrawOracle contract.
type WithdrawOraclePendingBalancesAdd struct {
	AddBalance   *big.Int
	TotalBalance *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPendingBalancesAdd is a free log retrieval operation binding the contract event 0x16d3bf59b512b44c3e8ce0628db2aaf2bc0606f359f3eb5cb438e53903f77bef.
//
// Solidity: event PendingBalancesAdd(uint256 _addBalance, uint256 _totalBalance)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterPendingBalancesAdd(opts *bind.FilterOpts) (*WithdrawOraclePendingBalancesAddIterator, error) {

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "PendingBalancesAdd")
	if err != nil {
		return nil, err
	}
	return &WithdrawOraclePendingBalancesAddIterator{contract: _WithdrawOracle.contract, event: "PendingBalancesAdd", logs: logs, sub: sub}, nil
}

// WatchPendingBalancesAdd is a free log subscription operation binding the contract event 0x16d3bf59b512b44c3e8ce0628db2aaf2bc0606f359f3eb5cb438e53903f77bef.
//
// Solidity: event PendingBalancesAdd(uint256 _addBalance, uint256 _totalBalance)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchPendingBalancesAdd(opts *bind.WatchOpts, sink chan<- *WithdrawOraclePendingBalancesAdd) (event.Subscription, error) {

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "PendingBalancesAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOraclePendingBalancesAdd)
				if err := _WithdrawOracle.contract.UnpackLog(event, "PendingBalancesAdd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePendingBalancesAdd is a log parse operation binding the contract event 0x16d3bf59b512b44c3e8ce0628db2aaf2bc0606f359f3eb5cb438e53903f77bef.
//
// Solidity: event PendingBalancesAdd(uint256 _addBalance, uint256 _totalBalance)
func (_WithdrawOracle *WithdrawOracleFilterer) ParsePendingBalancesAdd(log types.Log) (*WithdrawOraclePendingBalancesAdd, error) {
	event := new(WithdrawOraclePendingBalancesAdd)
	if err := _WithdrawOracle.contract.UnpackLog(event, "PendingBalancesAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOraclePendingBalancesResetIterator is returned from FilterPendingBalancesReset and is used to iterate over the raw logs and unpacked data for PendingBalancesReset events raised by the WithdrawOracle contract.
type WithdrawOraclePendingBalancesResetIterator struct {
	Event *WithdrawOraclePendingBalancesReset // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOraclePendingBalancesResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOraclePendingBalancesReset)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOraclePendingBalancesReset)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOraclePendingBalancesResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOraclePendingBalancesResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOraclePendingBalancesReset represents a PendingBalancesReset event raised by the WithdrawOracle contract.
type WithdrawOraclePendingBalancesReset struct {
	TotalBalance *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPendingBalancesReset is a free log retrieval operation binding the contract event 0x1944dc21941697055aa945cbccb8d3edd04161a51aff2d83089d1a82de3031a7.
//
// Solidity: event PendingBalancesReset(uint256 totalBalance)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterPendingBalancesReset(opts *bind.FilterOpts) (*WithdrawOraclePendingBalancesResetIterator, error) {

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "PendingBalancesReset")
	if err != nil {
		return nil, err
	}
	return &WithdrawOraclePendingBalancesResetIterator{contract: _WithdrawOracle.contract, event: "PendingBalancesReset", logs: logs, sub: sub}, nil
}

// WatchPendingBalancesReset is a free log subscription operation binding the contract event 0x1944dc21941697055aa945cbccb8d3edd04161a51aff2d83089d1a82de3031a7.
//
// Solidity: event PendingBalancesReset(uint256 totalBalance)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchPendingBalancesReset(opts *bind.WatchOpts, sink chan<- *WithdrawOraclePendingBalancesReset) (event.Subscription, error) {

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "PendingBalancesReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOraclePendingBalancesReset)
				if err := _WithdrawOracle.contract.UnpackLog(event, "PendingBalancesReset", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePendingBalancesReset is a log parse operation binding the contract event 0x1944dc21941697055aa945cbccb8d3edd04161a51aff2d83089d1a82de3031a7.
//
// Solidity: event PendingBalancesReset(uint256 totalBalance)
func (_WithdrawOracle *WithdrawOracleFilterer) ParsePendingBalancesReset(log types.Log) (*WithdrawOraclePendingBalancesReset, error) {
	event := new(WithdrawOraclePendingBalancesReset)
	if err := _WithdrawOracle.contract.UnpackLog(event, "PendingBalancesReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleProcessingStartedIterator is returned from FilterProcessingStarted and is used to iterate over the raw logs and unpacked data for ProcessingStarted events raised by the WithdrawOracle contract.
type WithdrawOracleProcessingStartedIterator struct {
	Event *WithdrawOracleProcessingStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleProcessingStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleProcessingStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleProcessingStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleProcessingStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleProcessingStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleProcessingStarted represents a ProcessingStarted event raised by the WithdrawOracle contract.
type WithdrawOracleProcessingStarted struct {
	RefSlot *big.Int
	Hash    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProcessingStarted is a free log retrieval operation binding the contract event 0xf73febded7d4502284718948a3e1d75406151c6326bde069424a584a4f6af87a.
//
// Solidity: event ProcessingStarted(uint256 indexed refSlot, bytes32 hash)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterProcessingStarted(opts *bind.FilterOpts, refSlot []*big.Int) (*WithdrawOracleProcessingStartedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "ProcessingStarted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleProcessingStartedIterator{contract: _WithdrawOracle.contract, event: "ProcessingStarted", logs: logs, sub: sub}, nil
}

// WatchProcessingStarted is a free log subscription operation binding the contract event 0xf73febded7d4502284718948a3e1d75406151c6326bde069424a584a4f6af87a.
//
// Solidity: event ProcessingStarted(uint256 indexed refSlot, bytes32 hash)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchProcessingStarted(opts *bind.WatchOpts, sink chan<- *WithdrawOracleProcessingStarted, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "ProcessingStarted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleProcessingStarted)
				if err := _WithdrawOracle.contract.UnpackLog(event, "ProcessingStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProcessingStarted is a log parse operation binding the contract event 0xf73febded7d4502284718948a3e1d75406151c6326bde069424a584a4f6af87a.
//
// Solidity: event ProcessingStarted(uint256 indexed refSlot, bytes32 hash)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseProcessingStarted(log types.Log) (*WithdrawOracleProcessingStarted, error) {
	event := new(WithdrawOracleProcessingStarted)
	if err := _WithdrawOracle.contract.UnpackLog(event, "ProcessingStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleReportDataSuccessIterator is returned from FilterReportDataSuccess and is used to iterate over the raw logs and unpacked data for ReportDataSuccess events raised by the WithdrawOracle contract.
type WithdrawOracleReportDataSuccessIterator struct {
	Event *WithdrawOracleReportDataSuccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleReportDataSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleReportDataSuccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleReportDataSuccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleReportDataSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleReportDataSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleReportDataSuccess represents a ReportDataSuccess event raised by the WithdrawOracle contract.
type WithdrawOracleReportDataSuccess struct {
	RefSlot           *big.Int
	ReportExitedCount *big.Int
	ClBalance         *big.Int
	ClVaultBalance    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterReportDataSuccess is a free log retrieval operation binding the contract event 0xfd69dfaf079af10c13c5c218efd437e0dd3033e77bae01f84dd829e19d8216e2.
//
// Solidity: event ReportDataSuccess(uint256 indexed refSlot, uint256 reportExitedCount, uint256 clBalance, uint256 clVaultBalance)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterReportDataSuccess(opts *bind.FilterOpts, refSlot []*big.Int) (*WithdrawOracleReportDataSuccessIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "ReportDataSuccess", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleReportDataSuccessIterator{contract: _WithdrawOracle.contract, event: "ReportDataSuccess", logs: logs, sub: sub}, nil
}

// WatchReportDataSuccess is a free log subscription operation binding the contract event 0xfd69dfaf079af10c13c5c218efd437e0dd3033e77bae01f84dd829e19d8216e2.
//
// Solidity: event ReportDataSuccess(uint256 indexed refSlot, uint256 reportExitedCount, uint256 clBalance, uint256 clVaultBalance)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchReportDataSuccess(opts *bind.WatchOpts, sink chan<- *WithdrawOracleReportDataSuccess, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "ReportDataSuccess", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleReportDataSuccess)
				if err := _WithdrawOracle.contract.UnpackLog(event, "ReportDataSuccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReportDataSuccess is a log parse operation binding the contract event 0xfd69dfaf079af10c13c5c218efd437e0dd3033e77bae01f84dd829e19d8216e2.
//
// Solidity: event ReportDataSuccess(uint256 indexed refSlot, uint256 reportExitedCount, uint256 clBalance, uint256 clVaultBalance)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseReportDataSuccess(log types.Log) (*WithdrawOracleReportDataSuccess, error) {
	event := new(WithdrawOracleReportDataSuccess)
	if err := _WithdrawOracle.contract.UnpackLog(event, "ReportDataSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleReportSubmittedIterator is returned from FilterReportSubmitted and is used to iterate over the raw logs and unpacked data for ReportSubmitted events raised by the WithdrawOracle contract.
type WithdrawOracleReportSubmittedIterator struct {
	Event *WithdrawOracleReportSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleReportSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleReportSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleReportSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleReportSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleReportSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleReportSubmitted represents a ReportSubmitted event raised by the WithdrawOracle contract.
type WithdrawOracleReportSubmitted struct {
	RefSlot                *big.Int
	Hash                   [32]byte
	ProcessingDeadlineTime *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterReportSubmitted is a free log retrieval operation binding the contract event 0xaed7d1a7a1831158dcda1e4214f5862f450bd3eb5721a5f322bf8c9fe1790b0a.
//
// Solidity: event ReportSubmitted(uint256 indexed refSlot, bytes32 hash, uint256 processingDeadlineTime)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterReportSubmitted(opts *bind.FilterOpts, refSlot []*big.Int) (*WithdrawOracleReportSubmittedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "ReportSubmitted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleReportSubmittedIterator{contract: _WithdrawOracle.contract, event: "ReportSubmitted", logs: logs, sub: sub}, nil
}

// WatchReportSubmitted is a free log subscription operation binding the contract event 0xaed7d1a7a1831158dcda1e4214f5862f450bd3eb5721a5f322bf8c9fe1790b0a.
//
// Solidity: event ReportSubmitted(uint256 indexed refSlot, bytes32 hash, uint256 processingDeadlineTime)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchReportSubmitted(opts *bind.WatchOpts, sink chan<- *WithdrawOracleReportSubmitted, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "ReportSubmitted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleReportSubmitted)
				if err := _WithdrawOracle.contract.UnpackLog(event, "ReportSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReportSubmitted is a log parse operation binding the contract event 0xaed7d1a7a1831158dcda1e4214f5862f450bd3eb5721a5f322bf8c9fe1790b0a.
//
// Solidity: event ReportSubmitted(uint256 indexed refSlot, bytes32 hash, uint256 processingDeadlineTime)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseReportSubmitted(log types.Log) (*WithdrawOracleReportSubmitted, error) {
	event := new(WithdrawOracleReportSubmitted)
	if err := _WithdrawOracle.contract.UnpackLog(event, "ReportSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the WithdrawOracle contract.
type WithdrawOracleUnpausedIterator struct {
	Event *WithdrawOracleUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleUnpaused represents a Unpaused event raised by the WithdrawOracle contract.
type WithdrawOracleUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterUnpaused(opts *bind.FilterOpts) (*WithdrawOracleUnpausedIterator, error) {

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleUnpausedIterator{contract: _WithdrawOracle.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WithdrawOracleUnpaused) (event.Subscription, error) {

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleUnpaused)
				if err := _WithdrawOracle.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseUnpaused(log types.Log) (*WithdrawOracleUnpaused, error) {
	event := new(WithdrawOracleUnpaused)
	if err := _WithdrawOracle.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleUpdateClVaultMinSettleLimitIterator is returned from FilterUpdateClVaultMinSettleLimit and is used to iterate over the raw logs and unpacked data for UpdateClVaultMinSettleLimit events raised by the WithdrawOracle contract.
type WithdrawOracleUpdateClVaultMinSettleLimitIterator struct {
	Event *WithdrawOracleUpdateClVaultMinSettleLimit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleUpdateClVaultMinSettleLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleUpdateClVaultMinSettleLimit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleUpdateClVaultMinSettleLimit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleUpdateClVaultMinSettleLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleUpdateClVaultMinSettleLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleUpdateClVaultMinSettleLimit represents a UpdateClVaultMinSettleLimit event raised by the WithdrawOracle contract.
type WithdrawOracleUpdateClVaultMinSettleLimit struct {
	ClVaultMinSettleLimit *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUpdateClVaultMinSettleLimit is a free log retrieval operation binding the contract event 0x6765ab460e8139130e86a064a184d24413dd588c41aa2277c433d2d09832f39e.
//
// Solidity: event UpdateClVaultMinSettleLimit(uint256 clVaultMinSettleLimit)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterUpdateClVaultMinSettleLimit(opts *bind.FilterOpts) (*WithdrawOracleUpdateClVaultMinSettleLimitIterator, error) {

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "UpdateClVaultMinSettleLimit")
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleUpdateClVaultMinSettleLimitIterator{contract: _WithdrawOracle.contract, event: "UpdateClVaultMinSettleLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateClVaultMinSettleLimit is a free log subscription operation binding the contract event 0x6765ab460e8139130e86a064a184d24413dd588c41aa2277c433d2d09832f39e.
//
// Solidity: event UpdateClVaultMinSettleLimit(uint256 clVaultMinSettleLimit)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchUpdateClVaultMinSettleLimit(opts *bind.WatchOpts, sink chan<- *WithdrawOracleUpdateClVaultMinSettleLimit) (event.Subscription, error) {

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "UpdateClVaultMinSettleLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleUpdateClVaultMinSettleLimit)
				if err := _WithdrawOracle.contract.UnpackLog(event, "UpdateClVaultMinSettleLimit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateClVaultMinSettleLimit is a log parse operation binding the contract event 0x6765ab460e8139130e86a064a184d24413dd588c41aa2277c433d2d09832f39e.
//
// Solidity: event UpdateClVaultMinSettleLimit(uint256 clVaultMinSettleLimit)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseUpdateClVaultMinSettleLimit(log types.Log) (*WithdrawOracleUpdateClVaultMinSettleLimit, error) {
	event := new(WithdrawOracleUpdateClVaultMinSettleLimit)
	if err := _WithdrawOracle.contract.UnpackLog(event, "UpdateClVaultMinSettleLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleUpdateExitRequestLimitIterator is returned from FilterUpdateExitRequestLimit and is used to iterate over the raw logs and unpacked data for UpdateExitRequestLimit events raised by the WithdrawOracle contract.
type WithdrawOracleUpdateExitRequestLimitIterator struct {
	Event *WithdrawOracleUpdateExitRequestLimit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleUpdateExitRequestLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleUpdateExitRequestLimit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleUpdateExitRequestLimit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleUpdateExitRequestLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleUpdateExitRequestLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleUpdateExitRequestLimit represents a UpdateExitRequestLimit event raised by the WithdrawOracle contract.
type WithdrawOracleUpdateExitRequestLimit struct {
	ExitRequestLimit *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateExitRequestLimit is a free log retrieval operation binding the contract event 0x6c79895328b404c7bf1115a4be7c52a4a974c20b13547a4fe873d5f997638c26.
//
// Solidity: event UpdateExitRequestLimit(uint256 exitRequestLimit)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterUpdateExitRequestLimit(opts *bind.FilterOpts) (*WithdrawOracleUpdateExitRequestLimitIterator, error) {

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "UpdateExitRequestLimit")
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleUpdateExitRequestLimitIterator{contract: _WithdrawOracle.contract, event: "UpdateExitRequestLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateExitRequestLimit is a free log subscription operation binding the contract event 0x6c79895328b404c7bf1115a4be7c52a4a974c20b13547a4fe873d5f997638c26.
//
// Solidity: event UpdateExitRequestLimit(uint256 exitRequestLimit)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchUpdateExitRequestLimit(opts *bind.WatchOpts, sink chan<- *WithdrawOracleUpdateExitRequestLimit) (event.Subscription, error) {

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "UpdateExitRequestLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleUpdateExitRequestLimit)
				if err := _WithdrawOracle.contract.UnpackLog(event, "UpdateExitRequestLimit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateExitRequestLimit is a log parse operation binding the contract event 0x6c79895328b404c7bf1115a4be7c52a4a974c20b13547a4fe873d5f997638c26.
//
// Solidity: event UpdateExitRequestLimit(uint256 exitRequestLimit)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseUpdateExitRequestLimit(log types.Log) (*WithdrawOracleUpdateExitRequestLimit, error) {
	event := new(WithdrawOracleUpdateExitRequestLimit)
	if err := _WithdrawOracle.contract.UnpackLog(event, "UpdateExitRequestLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleUpdateTotalBalanceTolerateIterator is returned from FilterUpdateTotalBalanceTolerate and is used to iterate over the raw logs and unpacked data for UpdateTotalBalanceTolerate events raised by the WithdrawOracle contract.
type WithdrawOracleUpdateTotalBalanceTolerateIterator struct {
	Event *WithdrawOracleUpdateTotalBalanceTolerate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleUpdateTotalBalanceTolerateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleUpdateTotalBalanceTolerate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleUpdateTotalBalanceTolerate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleUpdateTotalBalanceTolerateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleUpdateTotalBalanceTolerateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleUpdateTotalBalanceTolerate represents a UpdateTotalBalanceTolerate event raised by the WithdrawOracle contract.
type WithdrawOracleUpdateTotalBalanceTolerate struct {
	Old                  *big.Int
	TotalBalanceTolerate *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateTotalBalanceTolerate is a free log retrieval operation binding the contract event 0x72fe86693295f4b3370c792dcd34a475525a0877cdac916ab36dbcbb3bce6219.
//
// Solidity: event UpdateTotalBalanceTolerate(uint256 old, uint256 totalBalanceTolerate)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterUpdateTotalBalanceTolerate(opts *bind.FilterOpts) (*WithdrawOracleUpdateTotalBalanceTolerateIterator, error) {

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "UpdateTotalBalanceTolerate")
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleUpdateTotalBalanceTolerateIterator{contract: _WithdrawOracle.contract, event: "UpdateTotalBalanceTolerate", logs: logs, sub: sub}, nil
}

// WatchUpdateTotalBalanceTolerate is a free log subscription operation binding the contract event 0x72fe86693295f4b3370c792dcd34a475525a0877cdac916ab36dbcbb3bce6219.
//
// Solidity: event UpdateTotalBalanceTolerate(uint256 old, uint256 totalBalanceTolerate)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchUpdateTotalBalanceTolerate(opts *bind.WatchOpts, sink chan<- *WithdrawOracleUpdateTotalBalanceTolerate) (event.Subscription, error) {

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "UpdateTotalBalanceTolerate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleUpdateTotalBalanceTolerate)
				if err := _WithdrawOracle.contract.UnpackLog(event, "UpdateTotalBalanceTolerate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateTotalBalanceTolerate is a log parse operation binding the contract event 0x72fe86693295f4b3370c792dcd34a475525a0877cdac916ab36dbcbb3bce6219.
//
// Solidity: event UpdateTotalBalanceTolerate(uint256 old, uint256 totalBalanceTolerate)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseUpdateTotalBalanceTolerate(log types.Log) (*WithdrawOracleUpdateTotalBalanceTolerate, error) {
	event := new(WithdrawOracleUpdateTotalBalanceTolerate)
	if err := _WithdrawOracle.contract.UnpackLog(event, "UpdateTotalBalanceTolerate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the WithdrawOracle contract.
type WithdrawOracleUpgradedIterator struct {
	Event *WithdrawOracleUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleUpgraded represents a Upgraded event raised by the WithdrawOracle contract.
type WithdrawOracleUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*WithdrawOracleUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleUpgradedIterator{contract: _WithdrawOracle.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *WithdrawOracleUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleUpgraded)
				if err := _WithdrawOracle.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseUpgraded(log types.Log) (*WithdrawOracleUpgraded, error) {
	event := new(WithdrawOracleUpgraded)
	if err := _WithdrawOracle.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleVaultManagerChangedIterator is returned from FilterVaultManagerChanged and is used to iterate over the raw logs and unpacked data for VaultManagerChanged events raised by the WithdrawOracle contract.
type WithdrawOracleVaultManagerChangedIterator struct {
	Event *WithdrawOracleVaultManagerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleVaultManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleVaultManagerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleVaultManagerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleVaultManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleVaultManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleVaultManagerChanged represents a VaultManagerChanged event raised by the WithdrawOracle contract.
type WithdrawOracleVaultManagerChanged struct {
	OldVaultManager common.Address
	NewVaultManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVaultManagerChanged is a free log retrieval operation binding the contract event 0x9e8e45be62c510326288c84df8d83e42380c150d3181f5fb5e7a1be5958ad67a.
//
// Solidity: event VaultManagerChanged(address oldVaultManager, address newVaultManager)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterVaultManagerChanged(opts *bind.FilterOpts) (*WithdrawOracleVaultManagerChangedIterator, error) {

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "VaultManagerChanged")
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleVaultManagerChangedIterator{contract: _WithdrawOracle.contract, event: "VaultManagerChanged", logs: logs, sub: sub}, nil
}

// WatchVaultManagerChanged is a free log subscription operation binding the contract event 0x9e8e45be62c510326288c84df8d83e42380c150d3181f5fb5e7a1be5958ad67a.
//
// Solidity: event VaultManagerChanged(address oldVaultManager, address newVaultManager)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchVaultManagerChanged(opts *bind.WatchOpts, sink chan<- *WithdrawOracleVaultManagerChanged) (event.Subscription, error) {

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "VaultManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleVaultManagerChanged)
				if err := _WithdrawOracle.contract.UnpackLog(event, "VaultManagerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVaultManagerChanged is a log parse operation binding the contract event 0x9e8e45be62c510326288c84df8d83e42380c150d3181f5fb5e7a1be5958ad67a.
//
// Solidity: event VaultManagerChanged(address oldVaultManager, address newVaultManager)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseVaultManagerChanged(log types.Log) (*WithdrawOracleVaultManagerChanged, error) {
	event := new(WithdrawOracleVaultManagerChanged)
	if err := _WithdrawOracle.contract.UnpackLog(event, "VaultManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleWarnDataIncompleteProcessingIterator is returned from FilterWarnDataIncompleteProcessing and is used to iterate over the raw logs and unpacked data for WarnDataIncompleteProcessing events raised by the WithdrawOracle contract.
type WithdrawOracleWarnDataIncompleteProcessingIterator struct {
	Event *WithdrawOracleWarnDataIncompleteProcessing // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleWarnDataIncompleteProcessingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleWarnDataIncompleteProcessing)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleWarnDataIncompleteProcessing)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleWarnDataIncompleteProcessingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleWarnDataIncompleteProcessingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleWarnDataIncompleteProcessing represents a WarnDataIncompleteProcessing event raised by the WithdrawOracle contract.
type WithdrawOracleWarnDataIncompleteProcessing struct {
	RefSlot           *big.Int
	ExitRequestLimit  *big.Int
	ReportExitedCount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWarnDataIncompleteProcessing is a free log retrieval operation binding the contract event 0xefc67aab43195093a8d8ed25d52281d96de480748ece2787888c586e8e1e79b4.
//
// Solidity: event WarnDataIncompleteProcessing(uint256 indexed refSlot, uint256 exitRequestLimit, uint256 reportExitedCount)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterWarnDataIncompleteProcessing(opts *bind.FilterOpts, refSlot []*big.Int) (*WithdrawOracleWarnDataIncompleteProcessingIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "WarnDataIncompleteProcessing", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleWarnDataIncompleteProcessingIterator{contract: _WithdrawOracle.contract, event: "WarnDataIncompleteProcessing", logs: logs, sub: sub}, nil
}

// WatchWarnDataIncompleteProcessing is a free log subscription operation binding the contract event 0xefc67aab43195093a8d8ed25d52281d96de480748ece2787888c586e8e1e79b4.
//
// Solidity: event WarnDataIncompleteProcessing(uint256 indexed refSlot, uint256 exitRequestLimit, uint256 reportExitedCount)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchWarnDataIncompleteProcessing(opts *bind.WatchOpts, sink chan<- *WithdrawOracleWarnDataIncompleteProcessing, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "WarnDataIncompleteProcessing", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleWarnDataIncompleteProcessing)
				if err := _WithdrawOracle.contract.UnpackLog(event, "WarnDataIncompleteProcessing", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWarnDataIncompleteProcessing is a log parse operation binding the contract event 0xefc67aab43195093a8d8ed25d52281d96de480748ece2787888c586e8e1e79b4.
//
// Solidity: event WarnDataIncompleteProcessing(uint256 indexed refSlot, uint256 exitRequestLimit, uint256 reportExitedCount)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseWarnDataIncompleteProcessing(log types.Log) (*WithdrawOracleWarnDataIncompleteProcessing, error) {
	event := new(WithdrawOracleWarnDataIncompleteProcessing)
	if err := _WithdrawOracle.contract.UnpackLog(event, "WarnDataIncompleteProcessing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawOracleWarnProcessingMissedIterator is returned from FilterWarnProcessingMissed and is used to iterate over the raw logs and unpacked data for WarnProcessingMissed events raised by the WithdrawOracle contract.
type WithdrawOracleWarnProcessingMissedIterator struct {
	Event *WithdrawOracleWarnProcessingMissed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawOracleWarnProcessingMissedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawOracleWarnProcessingMissed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawOracleWarnProcessingMissed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawOracleWarnProcessingMissedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawOracleWarnProcessingMissedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawOracleWarnProcessingMissed represents a WarnProcessingMissed event raised by the WithdrawOracle contract.
type WithdrawOracleWarnProcessingMissed struct {
	RefSlot *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWarnProcessingMissed is a free log retrieval operation binding the contract event 0x800b849c8bf80718cf786c99d1091c079fe2c5e420a3ba7ba9b0ef8179ef2c38.
//
// Solidity: event WarnProcessingMissed(uint256 indexed refSlot)
func (_WithdrawOracle *WithdrawOracleFilterer) FilterWarnProcessingMissed(opts *bind.FilterOpts, refSlot []*big.Int) (*WithdrawOracleWarnProcessingMissedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _WithdrawOracle.contract.FilterLogs(opts, "WarnProcessingMissed", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawOracleWarnProcessingMissedIterator{contract: _WithdrawOracle.contract, event: "WarnProcessingMissed", logs: logs, sub: sub}, nil
}

// WatchWarnProcessingMissed is a free log subscription operation binding the contract event 0x800b849c8bf80718cf786c99d1091c079fe2c5e420a3ba7ba9b0ef8179ef2c38.
//
// Solidity: event WarnProcessingMissed(uint256 indexed refSlot)
func (_WithdrawOracle *WithdrawOracleFilterer) WatchWarnProcessingMissed(opts *bind.WatchOpts, sink chan<- *WithdrawOracleWarnProcessingMissed, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _WithdrawOracle.contract.WatchLogs(opts, "WarnProcessingMissed", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawOracleWarnProcessingMissed)
				if err := _WithdrawOracle.contract.UnpackLog(event, "WarnProcessingMissed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWarnProcessingMissed is a log parse operation binding the contract event 0x800b849c8bf80718cf786c99d1091c079fe2c5e420a3ba7ba9b0ef8179ef2c38.
//
// Solidity: event WarnProcessingMissed(uint256 indexed refSlot)
func (_WithdrawOracle *WithdrawOracleFilterer) ParseWarnProcessingMissed(log types.Log) (*WithdrawOracleWarnProcessingMissed, error) {
	event := new(WithdrawOracleWarnProcessingMissed)
	if err := _WithdrawOracle.contract.UnpackLog(event, "WarnProcessingMissed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracle

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
	ReportExitedCount          *big.Int
	WithdrawInfos              []WithdrawInfo
	ExitValidatorInfos         []ExitValidatorInfo
	DelayedExitTokenIds        []*big.Int
	LargeExitDelayedRequestIds []*big.Int
}

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AddressCannotBeSame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ArgumentOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClVaultBalanceNotMinSettleLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClVaultMinSettleLimitNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DaoCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExitRequestLimitNotZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingRefSlot\",\"type\":\"uint256\"}],\"name\":\"InitialRefSlotCannotBeLessThanProcessingOne\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContractVersionIncrement\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestsData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestsDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroContractVersionOnInit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyConsensusContractCanSubmitReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ProcessingDeadlineMissed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefSlotAlreadyProcessing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevRefSlot\",\"type\":\"uint256\"}],\"name\":\"RefSlotCannotDecrease\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingRefSlot\",\"type\":\"uint256\"}],\"name\":\"RefSlotMustBeGreaterThanProcessingOne\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedChainConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedVersion\",\"type\":\"uint256\"}],\"name\":\"UnexpectedConsensusVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"UnexpectedContractVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"consensusHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receivedHash\",\"type\":\"bytes32\"}],\"name\":\"UnexpectedDataHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"consensusRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataRefSlot\",\"type\":\"uint256\"}],\"name\":\"UnexpectedRefSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedRequestsDataLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"format\",\"type\":\"uint256\"}],\"name\":\"UnsupportedRequestsDataFormat\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ValidatorReportedExited\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionCannotBeSame\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevAddr\",\"type\":\"address\"}],\"name\":\"ConsensusHashContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prevVersion\",\"type\":\"uint256\"}],\"name\":\"ConsensusVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"ContractVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldDao\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"DaoAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_before\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_after\",\"type\":\"address\"}],\"name\":\"LiquidStakingChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_addBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalBalance\",\"type\":\"uint256\"}],\"name\":\"PendingBalancesAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalBalance\",\"type\":\"uint256\"}],\"name\":\"PendingBalancesReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ProcessingStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reportExitedCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clVaultBalance\",\"type\":\"uint256\"}],\"name\":\"ReportDataSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"}],\"name\":\"ReportSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clVaultMinSettleLimit\",\"type\":\"uint256\"}],\"name\":\"UpdateClVaultMinSettleLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitRequestLimit\",\"type\":\"uint256\"}],\"name\":\"UpdateExitRequestLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_before\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_after\",\"type\":\"address\"}],\"name\":\"VaultManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitRequestLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reportExitedCount\",\"type\":\"uint256\"}],\"name\":\"WarnDataIncompleteProcessing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"}],\"name\":\"WarnProcessingMissed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GENESIS_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_PER_SLOT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pendingBalance\",\"type\":\"uint256\"}],\"name\":\"addPendingBalances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clVaultBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clVaultMinSettleLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exitRequestLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClVaultBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusReport\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"processingStarted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastProcessingRefSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProcessingState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currentFrameRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"dataSubmitted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"reportExitedCount\",\"type\":\"uint256\"}],\"internalType\":\"structWithdrawOracle.ProcessingState\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"secondsPerSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"genesisTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consensusContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"consensusVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastProcessingRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_exitRequestLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_clVaultMinSettleLimit\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidStakingContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_clVaultMinSettleLimit\",\"type\":\"uint256\"}],\"name\":\"setClVaultMinSettleLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setConsensusContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"setConsensusVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"setDaoAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_exitRequestLimit\",\"type\":\"uint256\"}],\"name\":\"setExitRequestLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidStakingContractAddress\",\"type\":\"address\"}],\"name\":\"setLiquidStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultManagerContractAddress\",\"type\":\"address\"}],\"name\":\"setVaultManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reportHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"submitConsensusReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"consensusVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"clBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"clVaultBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reportExitedCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"clReward\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"clCapital\",\"type\":\"uint96\"}],\"internalType\":\"structWithdrawInfo[]\",\"name\":\"withdrawInfos\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"exitTokenId\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"exitBlockNumber\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"slashAmount\",\"type\":\"uint96\"}],\"internalType\":\"structExitValidatorInfo[]\",\"name\":\"exitValidatorInfos\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"delayedExitTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"largeExitDelayedRequestIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structWithdrawOracle.ReportData\",\"name\":\"data\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"contractVersion\",\"type\":\"uint256\"}],\"name\":\"submitReportData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint256)
func (_Oracle *OracleCaller) GENESISTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "GENESIS_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint256)
func (_Oracle *OracleSession) GENESISTIME() (*big.Int, error) {
	return _Oracle.Contract.GENESISTIME(&_Oracle.CallOpts)
}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint256)
func (_Oracle *OracleCallerSession) GENESISTIME() (*big.Int, error) {
	return _Oracle.Contract.GENESISTIME(&_Oracle.CallOpts)
}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint256)
func (_Oracle *OracleCaller) SECONDSPERSLOT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "SECONDS_PER_SLOT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint256)
func (_Oracle *OracleSession) SECONDSPERSLOT() (*big.Int, error) {
	return _Oracle.Contract.SECONDSPERSLOT(&_Oracle.CallOpts)
}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint256)
func (_Oracle *OracleCallerSession) SECONDSPERSLOT() (*big.Int, error) {
	return _Oracle.Contract.SECONDSPERSLOT(&_Oracle.CallOpts)
}

// ClBalances is a free data retrieval call binding the contract method 0x86c0c742.
//
// Solidity: function clBalances() view returns(uint256)
func (_Oracle *OracleCaller) ClBalances(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "clBalances")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClBalances is a free data retrieval call binding the contract method 0x86c0c742.
//
// Solidity: function clBalances() view returns(uint256)
func (_Oracle *OracleSession) ClBalances() (*big.Int, error) {
	return _Oracle.Contract.ClBalances(&_Oracle.CallOpts)
}

// ClBalances is a free data retrieval call binding the contract method 0x86c0c742.
//
// Solidity: function clBalances() view returns(uint256)
func (_Oracle *OracleCallerSession) ClBalances() (*big.Int, error) {
	return _Oracle.Contract.ClBalances(&_Oracle.CallOpts)
}

// ClVaultBalance is a free data retrieval call binding the contract method 0x312adfb8.
//
// Solidity: function clVaultBalance() view returns(uint256)
func (_Oracle *OracleCaller) ClVaultBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "clVaultBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClVaultBalance is a free data retrieval call binding the contract method 0x312adfb8.
//
// Solidity: function clVaultBalance() view returns(uint256)
func (_Oracle *OracleSession) ClVaultBalance() (*big.Int, error) {
	return _Oracle.Contract.ClVaultBalance(&_Oracle.CallOpts)
}

// ClVaultBalance is a free data retrieval call binding the contract method 0x312adfb8.
//
// Solidity: function clVaultBalance() view returns(uint256)
func (_Oracle *OracleCallerSession) ClVaultBalance() (*big.Int, error) {
	return _Oracle.Contract.ClVaultBalance(&_Oracle.CallOpts)
}

// ClVaultMinSettleLimit is a free data retrieval call binding the contract method 0xbde5a619.
//
// Solidity: function clVaultMinSettleLimit() view returns(uint256)
func (_Oracle *OracleCaller) ClVaultMinSettleLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "clVaultMinSettleLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClVaultMinSettleLimit is a free data retrieval call binding the contract method 0xbde5a619.
//
// Solidity: function clVaultMinSettleLimit() view returns(uint256)
func (_Oracle *OracleSession) ClVaultMinSettleLimit() (*big.Int, error) {
	return _Oracle.Contract.ClVaultMinSettleLimit(&_Oracle.CallOpts)
}

// ClVaultMinSettleLimit is a free data retrieval call binding the contract method 0xbde5a619.
//
// Solidity: function clVaultMinSettleLimit() view returns(uint256)
func (_Oracle *OracleCallerSession) ClVaultMinSettleLimit() (*big.Int, error) {
	return _Oracle.Contract.ClVaultMinSettleLimit(&_Oracle.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Oracle *OracleCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Oracle *OracleSession) Dao() (common.Address, error) {
	return _Oracle.Contract.Dao(&_Oracle.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Oracle *OracleCallerSession) Dao() (common.Address, error) {
	return _Oracle.Contract.Dao(&_Oracle.CallOpts)
}

// ExitRequestLimit is a free data retrieval call binding the contract method 0x1f78be77.
//
// Solidity: function exitRequestLimit() view returns(uint256)
func (_Oracle *OracleCaller) ExitRequestLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "exitRequestLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExitRequestLimit is a free data retrieval call binding the contract method 0x1f78be77.
//
// Solidity: function exitRequestLimit() view returns(uint256)
func (_Oracle *OracleSession) ExitRequestLimit() (*big.Int, error) {
	return _Oracle.Contract.ExitRequestLimit(&_Oracle.CallOpts)
}

// ExitRequestLimit is a free data retrieval call binding the contract method 0x1f78be77.
//
// Solidity: function exitRequestLimit() view returns(uint256)
func (_Oracle *OracleCallerSession) ExitRequestLimit() (*big.Int, error) {
	return _Oracle.Contract.ExitRequestLimit(&_Oracle.CallOpts)
}

// GetClBalances is a free data retrieval call binding the contract method 0x0fd520f6.
//
// Solidity: function getClBalances() view returns(uint256)
func (_Oracle *OracleCaller) GetClBalances(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getClBalances")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClBalances is a free data retrieval call binding the contract method 0x0fd520f6.
//
// Solidity: function getClBalances() view returns(uint256)
func (_Oracle *OracleSession) GetClBalances() (*big.Int, error) {
	return _Oracle.Contract.GetClBalances(&_Oracle.CallOpts)
}

// GetClBalances is a free data retrieval call binding the contract method 0x0fd520f6.
//
// Solidity: function getClBalances() view returns(uint256)
func (_Oracle *OracleCallerSession) GetClBalances() (*big.Int, error) {
	return _Oracle.Contract.GetClBalances(&_Oracle.CallOpts)
}

// GetClVaultBalances is a free data retrieval call binding the contract method 0x67fe1305.
//
// Solidity: function getClVaultBalances() view returns(uint256)
func (_Oracle *OracleCaller) GetClVaultBalances(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getClVaultBalances")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClVaultBalances is a free data retrieval call binding the contract method 0x67fe1305.
//
// Solidity: function getClVaultBalances() view returns(uint256)
func (_Oracle *OracleSession) GetClVaultBalances() (*big.Int, error) {
	return _Oracle.Contract.GetClVaultBalances(&_Oracle.CallOpts)
}

// GetClVaultBalances is a free data retrieval call binding the contract method 0x67fe1305.
//
// Solidity: function getClVaultBalances() view returns(uint256)
func (_Oracle *OracleCallerSession) GetClVaultBalances() (*big.Int, error) {
	return _Oracle.Contract.GetClVaultBalances(&_Oracle.CallOpts)
}

// GetConsensusContract is a free data retrieval call binding the contract method 0x8f55b571.
//
// Solidity: function getConsensusContract() view returns(address)
func (_Oracle *OracleCaller) GetConsensusContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getConsensusContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetConsensusContract is a free data retrieval call binding the contract method 0x8f55b571.
//
// Solidity: function getConsensusContract() view returns(address)
func (_Oracle *OracleSession) GetConsensusContract() (common.Address, error) {
	return _Oracle.Contract.GetConsensusContract(&_Oracle.CallOpts)
}

// GetConsensusContract is a free data retrieval call binding the contract method 0x8f55b571.
//
// Solidity: function getConsensusContract() view returns(address)
func (_Oracle *OracleCallerSession) GetConsensusContract() (common.Address, error) {
	return _Oracle.Contract.GetConsensusContract(&_Oracle.CallOpts)
}

// GetConsensusReport is a free data retrieval call binding the contract method 0x60d64d38.
//
// Solidity: function getConsensusReport() view returns(bytes32 hash, uint256 refSlot, uint256 processingDeadlineTime, bool processingStarted)
func (_Oracle *OracleCaller) GetConsensusReport(opts *bind.CallOpts) (struct {
	Hash                   [32]byte
	RefSlot                *big.Int
	ProcessingDeadlineTime *big.Int
	ProcessingStarted      bool
}, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getConsensusReport")

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
func (_Oracle *OracleSession) GetConsensusReport() (struct {
	Hash                   [32]byte
	RefSlot                *big.Int
	ProcessingDeadlineTime *big.Int
	ProcessingStarted      bool
}, error) {
	return _Oracle.Contract.GetConsensusReport(&_Oracle.CallOpts)
}

// GetConsensusReport is a free data retrieval call binding the contract method 0x60d64d38.
//
// Solidity: function getConsensusReport() view returns(bytes32 hash, uint256 refSlot, uint256 processingDeadlineTime, bool processingStarted)
func (_Oracle *OracleCallerSession) GetConsensusReport() (struct {
	Hash                   [32]byte
	RefSlot                *big.Int
	ProcessingDeadlineTime *big.Int
	ProcessingStarted      bool
}, error) {
	return _Oracle.Contract.GetConsensusReport(&_Oracle.CallOpts)
}

// GetConsensusVersion is a free data retrieval call binding the contract method 0x5be20425.
//
// Solidity: function getConsensusVersion() view returns(uint256)
func (_Oracle *OracleCaller) GetConsensusVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getConsensusVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConsensusVersion is a free data retrieval call binding the contract method 0x5be20425.
//
// Solidity: function getConsensusVersion() view returns(uint256)
func (_Oracle *OracleSession) GetConsensusVersion() (*big.Int, error) {
	return _Oracle.Contract.GetConsensusVersion(&_Oracle.CallOpts)
}

// GetConsensusVersion is a free data retrieval call binding the contract method 0x5be20425.
//
// Solidity: function getConsensusVersion() view returns(uint256)
func (_Oracle *OracleCallerSession) GetConsensusVersion() (*big.Int, error) {
	return _Oracle.Contract.GetConsensusVersion(&_Oracle.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Oracle *OracleCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getContractVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Oracle *OracleSession) GetContractVersion() (*big.Int, error) {
	return _Oracle.Contract.GetContractVersion(&_Oracle.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Oracle *OracleCallerSession) GetContractVersion() (*big.Int, error) {
	return _Oracle.Contract.GetContractVersion(&_Oracle.CallOpts)
}

// GetLastProcessingRefSlot is a free data retrieval call binding the contract method 0x3584d59c.
//
// Solidity: function getLastProcessingRefSlot() view returns(uint256)
func (_Oracle *OracleCaller) GetLastProcessingRefSlot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getLastProcessingRefSlot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastProcessingRefSlot is a free data retrieval call binding the contract method 0x3584d59c.
//
// Solidity: function getLastProcessingRefSlot() view returns(uint256)
func (_Oracle *OracleSession) GetLastProcessingRefSlot() (*big.Int, error) {
	return _Oracle.Contract.GetLastProcessingRefSlot(&_Oracle.CallOpts)
}

// GetLastProcessingRefSlot is a free data retrieval call binding the contract method 0x3584d59c.
//
// Solidity: function getLastProcessingRefSlot() view returns(uint256)
func (_Oracle *OracleCallerSession) GetLastProcessingRefSlot() (*big.Int, error) {
	return _Oracle.Contract.GetLastProcessingRefSlot(&_Oracle.CallOpts)
}

// GetPendingBalances is a free data retrieval call binding the contract method 0xf0b73e11.
//
// Solidity: function getPendingBalances() view returns(uint256)
func (_Oracle *OracleCaller) GetPendingBalances(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getPendingBalances")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingBalances is a free data retrieval call binding the contract method 0xf0b73e11.
//
// Solidity: function getPendingBalances() view returns(uint256)
func (_Oracle *OracleSession) GetPendingBalances() (*big.Int, error) {
	return _Oracle.Contract.GetPendingBalances(&_Oracle.CallOpts)
}

// GetPendingBalances is a free data retrieval call binding the contract method 0xf0b73e11.
//
// Solidity: function getPendingBalances() view returns(uint256)
func (_Oracle *OracleCallerSession) GetPendingBalances() (*big.Int, error) {
	return _Oracle.Contract.GetPendingBalances(&_Oracle.CallOpts)
}

// GetProcessingState is a free data retrieval call binding the contract method 0x8f7797c2.
//
// Solidity: function getProcessingState() view returns((uint256,uint256,bytes32,bool,uint256) result)
func (_Oracle *OracleCaller) GetProcessingState(opts *bind.CallOpts) (WithdrawOracleProcessingState, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getProcessingState")

	if err != nil {
		return *new(WithdrawOracleProcessingState), err
	}

	out0 := *abi.ConvertType(out[0], new(WithdrawOracleProcessingState)).(*WithdrawOracleProcessingState)

	return out0, err

}

// GetProcessingState is a free data retrieval call binding the contract method 0x8f7797c2.
//
// Solidity: function getProcessingState() view returns((uint256,uint256,bytes32,bool,uint256) result)
func (_Oracle *OracleSession) GetProcessingState() (WithdrawOracleProcessingState, error) {
	return _Oracle.Contract.GetProcessingState(&_Oracle.CallOpts)
}

// GetProcessingState is a free data retrieval call binding the contract method 0x8f7797c2.
//
// Solidity: function getProcessingState() view returns((uint256,uint256,bytes32,bool,uint256) result)
func (_Oracle *OracleCallerSession) GetProcessingState() (WithdrawOracleProcessingState, error) {
	return _Oracle.Contract.GetProcessingState(&_Oracle.CallOpts)
}

// LiquidStakingContractAddress is a free data retrieval call binding the contract method 0x6404a4c7.
//
// Solidity: function liquidStakingContractAddress() view returns(address)
func (_Oracle *OracleCaller) LiquidStakingContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "liquidStakingContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidStakingContractAddress is a free data retrieval call binding the contract method 0x6404a4c7.
//
// Solidity: function liquidStakingContractAddress() view returns(address)
func (_Oracle *OracleSession) LiquidStakingContractAddress() (common.Address, error) {
	return _Oracle.Contract.LiquidStakingContractAddress(&_Oracle.CallOpts)
}

// LiquidStakingContractAddress is a free data retrieval call binding the contract method 0x6404a4c7.
//
// Solidity: function liquidStakingContractAddress() view returns(address)
func (_Oracle *OracleCallerSession) LiquidStakingContractAddress() (common.Address, error) {
	return _Oracle.Contract.LiquidStakingContractAddress(&_Oracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oracle *OracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oracle *OracleSession) Owner() (common.Address, error) {
	return _Oracle.Contract.Owner(&_Oracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oracle *OracleCallerSession) Owner() (common.Address, error) {
	return _Oracle.Contract.Owner(&_Oracle.CallOpts)
}

// PendingBalances is a free data retrieval call binding the contract method 0xa65d1dbc.
//
// Solidity: function pendingBalances() view returns(uint256)
func (_Oracle *OracleCaller) PendingBalances(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "pendingBalances")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingBalances is a free data retrieval call binding the contract method 0xa65d1dbc.
//
// Solidity: function pendingBalances() view returns(uint256)
func (_Oracle *OracleSession) PendingBalances() (*big.Int, error) {
	return _Oracle.Contract.PendingBalances(&_Oracle.CallOpts)
}

// PendingBalances is a free data retrieval call binding the contract method 0xa65d1dbc.
//
// Solidity: function pendingBalances() view returns(uint256)
func (_Oracle *OracleCallerSession) PendingBalances() (*big.Int, error) {
	return _Oracle.Contract.PendingBalances(&_Oracle.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Oracle *OracleCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Oracle *OracleSession) ProxiableUUID() ([32]byte, error) {
	return _Oracle.Contract.ProxiableUUID(&_Oracle.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Oracle *OracleCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Oracle.Contract.ProxiableUUID(&_Oracle.CallOpts)
}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_Oracle *OracleCaller) VaultManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "vaultManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_Oracle *OracleSession) VaultManager() (common.Address, error) {
	return _Oracle.Contract.VaultManager(&_Oracle.CallOpts)
}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_Oracle *OracleCallerSession) VaultManager() (common.Address, error) {
	return _Oracle.Contract.VaultManager(&_Oracle.CallOpts)
}

// AddPendingBalances is a paid mutator transaction binding the contract method 0x8e34ead0.
//
// Solidity: function addPendingBalances(uint256 _pendingBalance) returns()
func (_Oracle *OracleTransactor) AddPendingBalances(opts *bind.TransactOpts, _pendingBalance *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "addPendingBalances", _pendingBalance)
}

// AddPendingBalances is a paid mutator transaction binding the contract method 0x8e34ead0.
//
// Solidity: function addPendingBalances(uint256 _pendingBalance) returns()
func (_Oracle *OracleSession) AddPendingBalances(_pendingBalance *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.AddPendingBalances(&_Oracle.TransactOpts, _pendingBalance)
}

// AddPendingBalances is a paid mutator transaction binding the contract method 0x8e34ead0.
//
// Solidity: function addPendingBalances(uint256 _pendingBalance) returns()
func (_Oracle *OracleTransactorSession) AddPendingBalances(_pendingBalance *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.AddPendingBalances(&_Oracle.TransactOpts, _pendingBalance)
}

// Initialize is a paid mutator transaction binding the contract method 0xafeda769.
//
// Solidity: function initialize(uint256 secondsPerSlot, uint256 genesisTime, address consensusContract, uint256 consensusVersion, uint256 lastProcessingRefSlot, address _dao, uint256 _exitRequestLimit, uint256 _clVaultMinSettleLimit) returns()
func (_Oracle *OracleTransactor) Initialize(opts *bind.TransactOpts, secondsPerSlot *big.Int, genesisTime *big.Int, consensusContract common.Address, consensusVersion *big.Int, lastProcessingRefSlot *big.Int, _dao common.Address, _exitRequestLimit *big.Int, _clVaultMinSettleLimit *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "initialize", secondsPerSlot, genesisTime, consensusContract, consensusVersion, lastProcessingRefSlot, _dao, _exitRequestLimit, _clVaultMinSettleLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0xafeda769.
//
// Solidity: function initialize(uint256 secondsPerSlot, uint256 genesisTime, address consensusContract, uint256 consensusVersion, uint256 lastProcessingRefSlot, address _dao, uint256 _exitRequestLimit, uint256 _clVaultMinSettleLimit) returns()
func (_Oracle *OracleSession) Initialize(secondsPerSlot *big.Int, genesisTime *big.Int, consensusContract common.Address, consensusVersion *big.Int, lastProcessingRefSlot *big.Int, _dao common.Address, _exitRequestLimit *big.Int, _clVaultMinSettleLimit *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.Initialize(&_Oracle.TransactOpts, secondsPerSlot, genesisTime, consensusContract, consensusVersion, lastProcessingRefSlot, _dao, _exitRequestLimit, _clVaultMinSettleLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0xafeda769.
//
// Solidity: function initialize(uint256 secondsPerSlot, uint256 genesisTime, address consensusContract, uint256 consensusVersion, uint256 lastProcessingRefSlot, address _dao, uint256 _exitRequestLimit, uint256 _clVaultMinSettleLimit) returns()
func (_Oracle *OracleTransactorSession) Initialize(secondsPerSlot *big.Int, genesisTime *big.Int, consensusContract common.Address, consensusVersion *big.Int, lastProcessingRefSlot *big.Int, _dao common.Address, _exitRequestLimit *big.Int, _clVaultMinSettleLimit *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.Initialize(&_Oracle.TransactOpts, secondsPerSlot, genesisTime, consensusContract, consensusVersion, lastProcessingRefSlot, _dao, _exitRequestLimit, _clVaultMinSettleLimit)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oracle *OracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oracle *OracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oracle.Contract.RenounceOwnership(&_Oracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oracle *OracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oracle.Contract.RenounceOwnership(&_Oracle.TransactOpts)
}

// SetClVaultMinSettleLimit is a paid mutator transaction binding the contract method 0xc14694e3.
//
// Solidity: function setClVaultMinSettleLimit(uint256 _clVaultMinSettleLimit) returns()
func (_Oracle *OracleTransactor) SetClVaultMinSettleLimit(opts *bind.TransactOpts, _clVaultMinSettleLimit *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setClVaultMinSettleLimit", _clVaultMinSettleLimit)
}

// SetClVaultMinSettleLimit is a paid mutator transaction binding the contract method 0xc14694e3.
//
// Solidity: function setClVaultMinSettleLimit(uint256 _clVaultMinSettleLimit) returns()
func (_Oracle *OracleSession) SetClVaultMinSettleLimit(_clVaultMinSettleLimit *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetClVaultMinSettleLimit(&_Oracle.TransactOpts, _clVaultMinSettleLimit)
}

// SetClVaultMinSettleLimit is a paid mutator transaction binding the contract method 0xc14694e3.
//
// Solidity: function setClVaultMinSettleLimit(uint256 _clVaultMinSettleLimit) returns()
func (_Oracle *OracleTransactorSession) SetClVaultMinSettleLimit(_clVaultMinSettleLimit *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetClVaultMinSettleLimit(&_Oracle.TransactOpts, _clVaultMinSettleLimit)
}

// SetConsensusContract is a paid mutator transaction binding the contract method 0xc469c307.
//
// Solidity: function setConsensusContract(address addr) returns()
func (_Oracle *OracleTransactor) SetConsensusContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setConsensusContract", addr)
}

// SetConsensusContract is a paid mutator transaction binding the contract method 0xc469c307.
//
// Solidity: function setConsensusContract(address addr) returns()
func (_Oracle *OracleSession) SetConsensusContract(addr common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetConsensusContract(&_Oracle.TransactOpts, addr)
}

// SetConsensusContract is a paid mutator transaction binding the contract method 0xc469c307.
//
// Solidity: function setConsensusContract(address addr) returns()
func (_Oracle *OracleTransactorSession) SetConsensusContract(addr common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetConsensusContract(&_Oracle.TransactOpts, addr)
}

// SetConsensusVersion is a paid mutator transaction binding the contract method 0x8d591474.
//
// Solidity: function setConsensusVersion(uint256 version) returns()
func (_Oracle *OracleTransactor) SetConsensusVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setConsensusVersion", version)
}

// SetConsensusVersion is a paid mutator transaction binding the contract method 0x8d591474.
//
// Solidity: function setConsensusVersion(uint256 version) returns()
func (_Oracle *OracleSession) SetConsensusVersion(version *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetConsensusVersion(&_Oracle.TransactOpts, version)
}

// SetConsensusVersion is a paid mutator transaction binding the contract method 0x8d591474.
//
// Solidity: function setConsensusVersion(uint256 version) returns()
func (_Oracle *OracleTransactorSession) SetConsensusVersion(version *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetConsensusVersion(&_Oracle.TransactOpts, version)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_Oracle *OracleTransactor) SetDaoAddress(opts *bind.TransactOpts, _dao common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setDaoAddress", _dao)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_Oracle *OracleSession) SetDaoAddress(_dao common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetDaoAddress(&_Oracle.TransactOpts, _dao)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_Oracle *OracleTransactorSession) SetDaoAddress(_dao common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetDaoAddress(&_Oracle.TransactOpts, _dao)
}

// SetExitRequestLimit is a paid mutator transaction binding the contract method 0x3936e707.
//
// Solidity: function setExitRequestLimit(uint256 _exitRequestLimit) returns()
func (_Oracle *OracleTransactor) SetExitRequestLimit(opts *bind.TransactOpts, _exitRequestLimit *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setExitRequestLimit", _exitRequestLimit)
}

// SetExitRequestLimit is a paid mutator transaction binding the contract method 0x3936e707.
//
// Solidity: function setExitRequestLimit(uint256 _exitRequestLimit) returns()
func (_Oracle *OracleSession) SetExitRequestLimit(_exitRequestLimit *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetExitRequestLimit(&_Oracle.TransactOpts, _exitRequestLimit)
}

// SetExitRequestLimit is a paid mutator transaction binding the contract method 0x3936e707.
//
// Solidity: function setExitRequestLimit(uint256 _exitRequestLimit) returns()
func (_Oracle *OracleTransactorSession) SetExitRequestLimit(_exitRequestLimit *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetExitRequestLimit(&_Oracle.TransactOpts, _exitRequestLimit)
}

// SetLiquidStaking is a paid mutator transaction binding the contract method 0x08211be5.
//
// Solidity: function setLiquidStaking(address _liquidStakingContractAddress) returns()
func (_Oracle *OracleTransactor) SetLiquidStaking(opts *bind.TransactOpts, _liquidStakingContractAddress common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setLiquidStaking", _liquidStakingContractAddress)
}

// SetLiquidStaking is a paid mutator transaction binding the contract method 0x08211be5.
//
// Solidity: function setLiquidStaking(address _liquidStakingContractAddress) returns()
func (_Oracle *OracleSession) SetLiquidStaking(_liquidStakingContractAddress common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetLiquidStaking(&_Oracle.TransactOpts, _liquidStakingContractAddress)
}

// SetLiquidStaking is a paid mutator transaction binding the contract method 0x08211be5.
//
// Solidity: function setLiquidStaking(address _liquidStakingContractAddress) returns()
func (_Oracle *OracleTransactorSession) SetLiquidStaking(_liquidStakingContractAddress common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetLiquidStaking(&_Oracle.TransactOpts, _liquidStakingContractAddress)
}

// SetVaultManager is a paid mutator transaction binding the contract method 0xb543503e.
//
// Solidity: function setVaultManager(address _vaultManagerContractAddress) returns()
func (_Oracle *OracleTransactor) SetVaultManager(opts *bind.TransactOpts, _vaultManagerContractAddress common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setVaultManager", _vaultManagerContractAddress)
}

// SetVaultManager is a paid mutator transaction binding the contract method 0xb543503e.
//
// Solidity: function setVaultManager(address _vaultManagerContractAddress) returns()
func (_Oracle *OracleSession) SetVaultManager(_vaultManagerContractAddress common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetVaultManager(&_Oracle.TransactOpts, _vaultManagerContractAddress)
}

// SetVaultManager is a paid mutator transaction binding the contract method 0xb543503e.
//
// Solidity: function setVaultManager(address _vaultManagerContractAddress) returns()
func (_Oracle *OracleTransactorSession) SetVaultManager(_vaultManagerContractAddress common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetVaultManager(&_Oracle.TransactOpts, _vaultManagerContractAddress)
}

// SubmitConsensusReport is a paid mutator transaction binding the contract method 0x063f36ad.
//
// Solidity: function submitConsensusReport(bytes32 reportHash, uint256 refSlot, uint256 deadline) returns()
func (_Oracle *OracleTransactor) SubmitConsensusReport(opts *bind.TransactOpts, reportHash [32]byte, refSlot *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "submitConsensusReport", reportHash, refSlot, deadline)
}

// SubmitConsensusReport is a paid mutator transaction binding the contract method 0x063f36ad.
//
// Solidity: function submitConsensusReport(bytes32 reportHash, uint256 refSlot, uint256 deadline) returns()
func (_Oracle *OracleSession) SubmitConsensusReport(reportHash [32]byte, refSlot *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SubmitConsensusReport(&_Oracle.TransactOpts, reportHash, refSlot, deadline)
}

// SubmitConsensusReport is a paid mutator transaction binding the contract method 0x063f36ad.
//
// Solidity: function submitConsensusReport(bytes32 reportHash, uint256 refSlot, uint256 deadline) returns()
func (_Oracle *OracleTransactorSession) SubmitConsensusReport(reportHash [32]byte, refSlot *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SubmitConsensusReport(&_Oracle.TransactOpts, reportHash, refSlot, deadline)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0x9599f6cd.
//
// Solidity: function submitReportData((uint256,uint256,uint256,uint256,uint256,(uint64,uint96,uint96)[],(uint64,uint96,uint96)[],uint256[],uint256[]) data, uint256 contractVersion) returns()
func (_Oracle *OracleTransactor) SubmitReportData(opts *bind.TransactOpts, data WithdrawOracleReportData, contractVersion *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "submitReportData", data, contractVersion)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0x9599f6cd.
//
// Solidity: function submitReportData((uint256,uint256,uint256,uint256,uint256,(uint64,uint96,uint96)[],(uint64,uint96,uint96)[],uint256[],uint256[]) data, uint256 contractVersion) returns()
func (_Oracle *OracleSession) SubmitReportData(data WithdrawOracleReportData, contractVersion *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SubmitReportData(&_Oracle.TransactOpts, data, contractVersion)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0x9599f6cd.
//
// Solidity: function submitReportData((uint256,uint256,uint256,uint256,uint256,(uint64,uint96,uint96)[],(uint64,uint96,uint96)[],uint256[],uint256[]) data, uint256 contractVersion) returns()
func (_Oracle *OracleTransactorSession) SubmitReportData(data WithdrawOracleReportData, contractVersion *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SubmitReportData(&_Oracle.TransactOpts, data, contractVersion)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oracle *OracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oracle *OracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.TransferOwnership(&_Oracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oracle *OracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.TransferOwnership(&_Oracle.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Oracle *OracleTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Oracle *OracleSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.UpgradeTo(&_Oracle.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Oracle *OracleTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.UpgradeTo(&_Oracle.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Oracle *OracleTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Oracle *OracleSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Oracle.Contract.UpgradeToAndCall(&_Oracle.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Oracle *OracleTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Oracle.Contract.UpgradeToAndCall(&_Oracle.TransactOpts, newImplementation, data)
}

// OracleAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Oracle contract.
type OracleAdminChangedIterator struct {
	Event *OracleAdminChanged // Event containing the contract specifics and raw log

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
func (it *OracleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleAdminChanged)
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
		it.Event = new(OracleAdminChanged)
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
func (it *OracleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleAdminChanged represents a AdminChanged event raised by the Oracle contract.
type OracleAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Oracle *OracleFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*OracleAdminChangedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &OracleAdminChangedIterator{contract: _Oracle.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Oracle *OracleFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *OracleAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleAdminChanged)
				if err := _Oracle.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseAdminChanged(log types.Log) (*OracleAdminChanged, error) {
	event := new(OracleAdminChanged)
	if err := _Oracle.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Oracle contract.
type OracleBeaconUpgradedIterator struct {
	Event *OracleBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *OracleBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleBeaconUpgraded)
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
		it.Event = new(OracleBeaconUpgraded)
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
func (it *OracleBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleBeaconUpgraded represents a BeaconUpgraded event raised by the Oracle contract.
type OracleBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Oracle *OracleFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*OracleBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &OracleBeaconUpgradedIterator{contract: _Oracle.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Oracle *OracleFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *OracleBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleBeaconUpgraded)
				if err := _Oracle.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseBeaconUpgraded(log types.Log) (*OracleBeaconUpgraded, error) {
	event := new(OracleBeaconUpgraded)
	if err := _Oracle.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleConsensusHashContractSetIterator is returned from FilterConsensusHashContractSet and is used to iterate over the raw logs and unpacked data for ConsensusHashContractSet events raised by the Oracle contract.
type OracleConsensusHashContractSetIterator struct {
	Event *OracleConsensusHashContractSet // Event containing the contract specifics and raw log

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
func (it *OracleConsensusHashContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleConsensusHashContractSet)
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
		it.Event = new(OracleConsensusHashContractSet)
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
func (it *OracleConsensusHashContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleConsensusHashContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleConsensusHashContractSet represents a ConsensusHashContractSet event raised by the Oracle contract.
type OracleConsensusHashContractSet struct {
	Addr     common.Address
	PrevAddr common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConsensusHashContractSet is a free log retrieval operation binding the contract event 0x25421480fb7f52d18947876279a213696b58d7e0e5416ce5e2c9f9942661c34c.
//
// Solidity: event ConsensusHashContractSet(address indexed addr, address indexed prevAddr)
func (_Oracle *OracleFilterer) FilterConsensusHashContractSet(opts *bind.FilterOpts, addr []common.Address, prevAddr []common.Address) (*OracleConsensusHashContractSetIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var prevAddrRule []interface{}
	for _, prevAddrItem := range prevAddr {
		prevAddrRule = append(prevAddrRule, prevAddrItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "ConsensusHashContractSet", addrRule, prevAddrRule)
	if err != nil {
		return nil, err
	}
	return &OracleConsensusHashContractSetIterator{contract: _Oracle.contract, event: "ConsensusHashContractSet", logs: logs, sub: sub}, nil
}

// WatchConsensusHashContractSet is a free log subscription operation binding the contract event 0x25421480fb7f52d18947876279a213696b58d7e0e5416ce5e2c9f9942661c34c.
//
// Solidity: event ConsensusHashContractSet(address indexed addr, address indexed prevAddr)
func (_Oracle *OracleFilterer) WatchConsensusHashContractSet(opts *bind.WatchOpts, sink chan<- *OracleConsensusHashContractSet, addr []common.Address, prevAddr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var prevAddrRule []interface{}
	for _, prevAddrItem := range prevAddr {
		prevAddrRule = append(prevAddrRule, prevAddrItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "ConsensusHashContractSet", addrRule, prevAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleConsensusHashContractSet)
				if err := _Oracle.contract.UnpackLog(event, "ConsensusHashContractSet", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseConsensusHashContractSet(log types.Log) (*OracleConsensusHashContractSet, error) {
	event := new(OracleConsensusHashContractSet)
	if err := _Oracle.contract.UnpackLog(event, "ConsensusHashContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleConsensusVersionSetIterator is returned from FilterConsensusVersionSet and is used to iterate over the raw logs and unpacked data for ConsensusVersionSet events raised by the Oracle contract.
type OracleConsensusVersionSetIterator struct {
	Event *OracleConsensusVersionSet // Event containing the contract specifics and raw log

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
func (it *OracleConsensusVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleConsensusVersionSet)
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
		it.Event = new(OracleConsensusVersionSet)
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
func (it *OracleConsensusVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleConsensusVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleConsensusVersionSet represents a ConsensusVersionSet event raised by the Oracle contract.
type OracleConsensusVersionSet struct {
	Version     *big.Int
	PrevVersion *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsensusVersionSet is a free log retrieval operation binding the contract event 0xfa5304972d4ec3e3207f0bbf91155a49d0dfa62488f9529403a2a49e4b29a895.
//
// Solidity: event ConsensusVersionSet(uint256 indexed version, uint256 indexed prevVersion)
func (_Oracle *OracleFilterer) FilterConsensusVersionSet(opts *bind.FilterOpts, version []*big.Int, prevVersion []*big.Int) (*OracleConsensusVersionSetIterator, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var prevVersionRule []interface{}
	for _, prevVersionItem := range prevVersion {
		prevVersionRule = append(prevVersionRule, prevVersionItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "ConsensusVersionSet", versionRule, prevVersionRule)
	if err != nil {
		return nil, err
	}
	return &OracleConsensusVersionSetIterator{contract: _Oracle.contract, event: "ConsensusVersionSet", logs: logs, sub: sub}, nil
}

// WatchConsensusVersionSet is a free log subscription operation binding the contract event 0xfa5304972d4ec3e3207f0bbf91155a49d0dfa62488f9529403a2a49e4b29a895.
//
// Solidity: event ConsensusVersionSet(uint256 indexed version, uint256 indexed prevVersion)
func (_Oracle *OracleFilterer) WatchConsensusVersionSet(opts *bind.WatchOpts, sink chan<- *OracleConsensusVersionSet, version []*big.Int, prevVersion []*big.Int) (event.Subscription, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var prevVersionRule []interface{}
	for _, prevVersionItem := range prevVersion {
		prevVersionRule = append(prevVersionRule, prevVersionItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "ConsensusVersionSet", versionRule, prevVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleConsensusVersionSet)
				if err := _Oracle.contract.UnpackLog(event, "ConsensusVersionSet", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseConsensusVersionSet(log types.Log) (*OracleConsensusVersionSet, error) {
	event := new(OracleConsensusVersionSet)
	if err := _Oracle.contract.UnpackLog(event, "ConsensusVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleContractVersionSetIterator is returned from FilterContractVersionSet and is used to iterate over the raw logs and unpacked data for ContractVersionSet events raised by the Oracle contract.
type OracleContractVersionSetIterator struct {
	Event *OracleContractVersionSet // Event containing the contract specifics and raw log

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
func (it *OracleContractVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractVersionSet)
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
		it.Event = new(OracleContractVersionSet)
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
func (it *OracleContractVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractVersionSet represents a ContractVersionSet event raised by the Oracle contract.
type OracleContractVersionSet struct {
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractVersionSet is a free log retrieval operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_Oracle *OracleFilterer) FilterContractVersionSet(opts *bind.FilterOpts) (*OracleContractVersionSetIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return &OracleContractVersionSetIterator{contract: _Oracle.contract, event: "ContractVersionSet", logs: logs, sub: sub}, nil
}

// WatchContractVersionSet is a free log subscription operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_Oracle *OracleFilterer) WatchContractVersionSet(opts *bind.WatchOpts, sink chan<- *OracleContractVersionSet) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractVersionSet)
				if err := _Oracle.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseContractVersionSet(log types.Log) (*OracleContractVersionSet, error) {
	event := new(OracleContractVersionSet)
	if err := _Oracle.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleDaoAddressChangedIterator is returned from FilterDaoAddressChanged and is used to iterate over the raw logs and unpacked data for DaoAddressChanged events raised by the Oracle contract.
type OracleDaoAddressChangedIterator struct {
	Event *OracleDaoAddressChanged // Event containing the contract specifics and raw log

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
func (it *OracleDaoAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleDaoAddressChanged)
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
		it.Event = new(OracleDaoAddressChanged)
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
func (it *OracleDaoAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleDaoAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleDaoAddressChanged represents a DaoAddressChanged event raised by the Oracle contract.
type OracleDaoAddressChanged struct {
	OldDao common.Address
	Dao    common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDaoAddressChanged is a free log retrieval operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address _oldDao, address _dao)
func (_Oracle *OracleFilterer) FilterDaoAddressChanged(opts *bind.FilterOpts) (*OracleDaoAddressChangedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "DaoAddressChanged")
	if err != nil {
		return nil, err
	}
	return &OracleDaoAddressChangedIterator{contract: _Oracle.contract, event: "DaoAddressChanged", logs: logs, sub: sub}, nil
}

// WatchDaoAddressChanged is a free log subscription operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address _oldDao, address _dao)
func (_Oracle *OracleFilterer) WatchDaoAddressChanged(opts *bind.WatchOpts, sink chan<- *OracleDaoAddressChanged) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "DaoAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleDaoAddressChanged)
				if err := _Oracle.contract.UnpackLog(event, "DaoAddressChanged", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseDaoAddressChanged(log types.Log) (*OracleDaoAddressChanged, error) {
	event := new(OracleDaoAddressChanged)
	if err := _Oracle.contract.UnpackLog(event, "DaoAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Oracle contract.
type OracleInitializedIterator struct {
	Event *OracleInitialized // Event containing the contract specifics and raw log

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
func (it *OracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleInitialized)
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
		it.Event = new(OracleInitialized)
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
func (it *OracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleInitialized represents a Initialized event raised by the Oracle contract.
type OracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Oracle *OracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*OracleInitializedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OracleInitializedIterator{contract: _Oracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Oracle *OracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OracleInitialized) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleInitialized)
				if err := _Oracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseInitialized(log types.Log) (*OracleInitialized, error) {
	event := new(OracleInitialized)
	if err := _Oracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleLiquidStakingChangedIterator is returned from FilterLiquidStakingChanged and is used to iterate over the raw logs and unpacked data for LiquidStakingChanged events raised by the Oracle contract.
type OracleLiquidStakingChangedIterator struct {
	Event *OracleLiquidStakingChanged // Event containing the contract specifics and raw log

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
func (it *OracleLiquidStakingChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleLiquidStakingChanged)
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
		it.Event = new(OracleLiquidStakingChanged)
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
func (it *OracleLiquidStakingChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleLiquidStakingChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleLiquidStakingChanged represents a LiquidStakingChanged event raised by the Oracle contract.
type OracleLiquidStakingChanged struct {
	Before common.Address
	After  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLiquidStakingChanged is a free log retrieval operation binding the contract event 0x516e369f70685d2905d214a9e8567098c02a0e00f5a55bd30baca6b51d446cef.
//
// Solidity: event LiquidStakingChanged(address _before, address _after)
func (_Oracle *OracleFilterer) FilterLiquidStakingChanged(opts *bind.FilterOpts) (*OracleLiquidStakingChangedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "LiquidStakingChanged")
	if err != nil {
		return nil, err
	}
	return &OracleLiquidStakingChangedIterator{contract: _Oracle.contract, event: "LiquidStakingChanged", logs: logs, sub: sub}, nil
}

// WatchLiquidStakingChanged is a free log subscription operation binding the contract event 0x516e369f70685d2905d214a9e8567098c02a0e00f5a55bd30baca6b51d446cef.
//
// Solidity: event LiquidStakingChanged(address _before, address _after)
func (_Oracle *OracleFilterer) WatchLiquidStakingChanged(opts *bind.WatchOpts, sink chan<- *OracleLiquidStakingChanged) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "LiquidStakingChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleLiquidStakingChanged)
				if err := _Oracle.contract.UnpackLog(event, "LiquidStakingChanged", log); err != nil {
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
// Solidity: event LiquidStakingChanged(address _before, address _after)
func (_Oracle *OracleFilterer) ParseLiquidStakingChanged(log types.Log) (*OracleLiquidStakingChanged, error) {
	event := new(OracleLiquidStakingChanged)
	if err := _Oracle.contract.UnpackLog(event, "LiquidStakingChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Oracle contract.
type OracleOwnershipTransferredIterator struct {
	Event *OracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleOwnershipTransferred)
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
		it.Event = new(OracleOwnershipTransferred)
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
func (it *OracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleOwnershipTransferred represents a OwnershipTransferred event raised by the Oracle contract.
type OracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oracle *OracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OracleOwnershipTransferredIterator{contract: _Oracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oracle *OracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleOwnershipTransferred)
				if err := _Oracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseOwnershipTransferred(log types.Log) (*OracleOwnershipTransferred, error) {
	event := new(OracleOwnershipTransferred)
	if err := _Oracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OraclePendingBalancesAddIterator is returned from FilterPendingBalancesAdd and is used to iterate over the raw logs and unpacked data for PendingBalancesAdd events raised by the Oracle contract.
type OraclePendingBalancesAddIterator struct {
	Event *OraclePendingBalancesAdd // Event containing the contract specifics and raw log

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
func (it *OraclePendingBalancesAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OraclePendingBalancesAdd)
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
		it.Event = new(OraclePendingBalancesAdd)
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
func (it *OraclePendingBalancesAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OraclePendingBalancesAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OraclePendingBalancesAdd represents a PendingBalancesAdd event raised by the Oracle contract.
type OraclePendingBalancesAdd struct {
	AddBalance   *big.Int
	TotalBalance *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPendingBalancesAdd is a free log retrieval operation binding the contract event 0x16d3bf59b512b44c3e8ce0628db2aaf2bc0606f359f3eb5cb438e53903f77bef.
//
// Solidity: event PendingBalancesAdd(uint256 _addBalance, uint256 _totalBalance)
func (_Oracle *OracleFilterer) FilterPendingBalancesAdd(opts *bind.FilterOpts) (*OraclePendingBalancesAddIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "PendingBalancesAdd")
	if err != nil {
		return nil, err
	}
	return &OraclePendingBalancesAddIterator{contract: _Oracle.contract, event: "PendingBalancesAdd", logs: logs, sub: sub}, nil
}

// WatchPendingBalancesAdd is a free log subscription operation binding the contract event 0x16d3bf59b512b44c3e8ce0628db2aaf2bc0606f359f3eb5cb438e53903f77bef.
//
// Solidity: event PendingBalancesAdd(uint256 _addBalance, uint256 _totalBalance)
func (_Oracle *OracleFilterer) WatchPendingBalancesAdd(opts *bind.WatchOpts, sink chan<- *OraclePendingBalancesAdd) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "PendingBalancesAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OraclePendingBalancesAdd)
				if err := _Oracle.contract.UnpackLog(event, "PendingBalancesAdd", log); err != nil {
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
func (_Oracle *OracleFilterer) ParsePendingBalancesAdd(log types.Log) (*OraclePendingBalancesAdd, error) {
	event := new(OraclePendingBalancesAdd)
	if err := _Oracle.contract.UnpackLog(event, "PendingBalancesAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OraclePendingBalancesResetIterator is returned from FilterPendingBalancesReset and is used to iterate over the raw logs and unpacked data for PendingBalancesReset events raised by the Oracle contract.
type OraclePendingBalancesResetIterator struct {
	Event *OraclePendingBalancesReset // Event containing the contract specifics and raw log

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
func (it *OraclePendingBalancesResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OraclePendingBalancesReset)
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
		it.Event = new(OraclePendingBalancesReset)
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
func (it *OraclePendingBalancesResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OraclePendingBalancesResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OraclePendingBalancesReset represents a PendingBalancesReset event raised by the Oracle contract.
type OraclePendingBalancesReset struct {
	TotalBalance *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPendingBalancesReset is a free log retrieval operation binding the contract event 0x1944dc21941697055aa945cbccb8d3edd04161a51aff2d83089d1a82de3031a7.
//
// Solidity: event PendingBalancesReset(uint256 _totalBalance)
func (_Oracle *OracleFilterer) FilterPendingBalancesReset(opts *bind.FilterOpts) (*OraclePendingBalancesResetIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "PendingBalancesReset")
	if err != nil {
		return nil, err
	}
	return &OraclePendingBalancesResetIterator{contract: _Oracle.contract, event: "PendingBalancesReset", logs: logs, sub: sub}, nil
}

// WatchPendingBalancesReset is a free log subscription operation binding the contract event 0x1944dc21941697055aa945cbccb8d3edd04161a51aff2d83089d1a82de3031a7.
//
// Solidity: event PendingBalancesReset(uint256 _totalBalance)
func (_Oracle *OracleFilterer) WatchPendingBalancesReset(opts *bind.WatchOpts, sink chan<- *OraclePendingBalancesReset) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "PendingBalancesReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OraclePendingBalancesReset)
				if err := _Oracle.contract.UnpackLog(event, "PendingBalancesReset", log); err != nil {
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
// Solidity: event PendingBalancesReset(uint256 _totalBalance)
func (_Oracle *OracleFilterer) ParsePendingBalancesReset(log types.Log) (*OraclePendingBalancesReset, error) {
	event := new(OraclePendingBalancesReset)
	if err := _Oracle.contract.UnpackLog(event, "PendingBalancesReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleProcessingStartedIterator is returned from FilterProcessingStarted and is used to iterate over the raw logs and unpacked data for ProcessingStarted events raised by the Oracle contract.
type OracleProcessingStartedIterator struct {
	Event *OracleProcessingStarted // Event containing the contract specifics and raw log

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
func (it *OracleProcessingStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProcessingStarted)
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
		it.Event = new(OracleProcessingStarted)
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
func (it *OracleProcessingStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProcessingStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProcessingStarted represents a ProcessingStarted event raised by the Oracle contract.
type OracleProcessingStarted struct {
	RefSlot *big.Int
	Hash    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProcessingStarted is a free log retrieval operation binding the contract event 0xf73febded7d4502284718948a3e1d75406151c6326bde069424a584a4f6af87a.
//
// Solidity: event ProcessingStarted(uint256 indexed refSlot, bytes32 hash)
func (_Oracle *OracleFilterer) FilterProcessingStarted(opts *bind.FilterOpts, refSlot []*big.Int) (*OracleProcessingStartedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "ProcessingStarted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &OracleProcessingStartedIterator{contract: _Oracle.contract, event: "ProcessingStarted", logs: logs, sub: sub}, nil
}

// WatchProcessingStarted is a free log subscription operation binding the contract event 0xf73febded7d4502284718948a3e1d75406151c6326bde069424a584a4f6af87a.
//
// Solidity: event ProcessingStarted(uint256 indexed refSlot, bytes32 hash)
func (_Oracle *OracleFilterer) WatchProcessingStarted(opts *bind.WatchOpts, sink chan<- *OracleProcessingStarted, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "ProcessingStarted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProcessingStarted)
				if err := _Oracle.contract.UnpackLog(event, "ProcessingStarted", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseProcessingStarted(log types.Log) (*OracleProcessingStarted, error) {
	event := new(OracleProcessingStarted)
	if err := _Oracle.contract.UnpackLog(event, "ProcessingStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleReportDataSuccessIterator is returned from FilterReportDataSuccess and is used to iterate over the raw logs and unpacked data for ReportDataSuccess events raised by the Oracle contract.
type OracleReportDataSuccessIterator struct {
	Event *OracleReportDataSuccess // Event containing the contract specifics and raw log

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
func (it *OracleReportDataSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleReportDataSuccess)
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
		it.Event = new(OracleReportDataSuccess)
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
func (it *OracleReportDataSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleReportDataSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleReportDataSuccess represents a ReportDataSuccess event raised by the Oracle contract.
type OracleReportDataSuccess struct {
	RefSlot           *big.Int
	ReportExitedCount *big.Int
	ClBalance         *big.Int
	ClVaultBalance    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterReportDataSuccess is a free log retrieval operation binding the contract event 0xfd69dfaf079af10c13c5c218efd437e0dd3033e77bae01f84dd829e19d8216e2.
//
// Solidity: event ReportDataSuccess(uint256 indexed refSlot, uint256 reportExitedCount, uint256 clBalance, uint256 clVaultBalance)
func (_Oracle *OracleFilterer) FilterReportDataSuccess(opts *bind.FilterOpts, refSlot []*big.Int) (*OracleReportDataSuccessIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "ReportDataSuccess", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &OracleReportDataSuccessIterator{contract: _Oracle.contract, event: "ReportDataSuccess", logs: logs, sub: sub}, nil
}

// WatchReportDataSuccess is a free log subscription operation binding the contract event 0xfd69dfaf079af10c13c5c218efd437e0dd3033e77bae01f84dd829e19d8216e2.
//
// Solidity: event ReportDataSuccess(uint256 indexed refSlot, uint256 reportExitedCount, uint256 clBalance, uint256 clVaultBalance)
func (_Oracle *OracleFilterer) WatchReportDataSuccess(opts *bind.WatchOpts, sink chan<- *OracleReportDataSuccess, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "ReportDataSuccess", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleReportDataSuccess)
				if err := _Oracle.contract.UnpackLog(event, "ReportDataSuccess", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseReportDataSuccess(log types.Log) (*OracleReportDataSuccess, error) {
	event := new(OracleReportDataSuccess)
	if err := _Oracle.contract.UnpackLog(event, "ReportDataSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleReportSubmittedIterator is returned from FilterReportSubmitted and is used to iterate over the raw logs and unpacked data for ReportSubmitted events raised by the Oracle contract.
type OracleReportSubmittedIterator struct {
	Event *OracleReportSubmitted // Event containing the contract specifics and raw log

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
func (it *OracleReportSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleReportSubmitted)
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
		it.Event = new(OracleReportSubmitted)
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
func (it *OracleReportSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleReportSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleReportSubmitted represents a ReportSubmitted event raised by the Oracle contract.
type OracleReportSubmitted struct {
	RefSlot                *big.Int
	Hash                   [32]byte
	ProcessingDeadlineTime *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterReportSubmitted is a free log retrieval operation binding the contract event 0xaed7d1a7a1831158dcda1e4214f5862f450bd3eb5721a5f322bf8c9fe1790b0a.
//
// Solidity: event ReportSubmitted(uint256 indexed refSlot, bytes32 hash, uint256 processingDeadlineTime)
func (_Oracle *OracleFilterer) FilterReportSubmitted(opts *bind.FilterOpts, refSlot []*big.Int) (*OracleReportSubmittedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "ReportSubmitted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &OracleReportSubmittedIterator{contract: _Oracle.contract, event: "ReportSubmitted", logs: logs, sub: sub}, nil
}

// WatchReportSubmitted is a free log subscription operation binding the contract event 0xaed7d1a7a1831158dcda1e4214f5862f450bd3eb5721a5f322bf8c9fe1790b0a.
//
// Solidity: event ReportSubmitted(uint256 indexed refSlot, bytes32 hash, uint256 processingDeadlineTime)
func (_Oracle *OracleFilterer) WatchReportSubmitted(opts *bind.WatchOpts, sink chan<- *OracleReportSubmitted, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "ReportSubmitted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleReportSubmitted)
				if err := _Oracle.contract.UnpackLog(event, "ReportSubmitted", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseReportSubmitted(log types.Log) (*OracleReportSubmitted, error) {
	event := new(OracleReportSubmitted)
	if err := _Oracle.contract.UnpackLog(event, "ReportSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleUpdateClVaultMinSettleLimitIterator is returned from FilterUpdateClVaultMinSettleLimit and is used to iterate over the raw logs and unpacked data for UpdateClVaultMinSettleLimit events raised by the Oracle contract.
type OracleUpdateClVaultMinSettleLimitIterator struct {
	Event *OracleUpdateClVaultMinSettleLimit // Event containing the contract specifics and raw log

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
func (it *OracleUpdateClVaultMinSettleLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleUpdateClVaultMinSettleLimit)
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
		it.Event = new(OracleUpdateClVaultMinSettleLimit)
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
func (it *OracleUpdateClVaultMinSettleLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleUpdateClVaultMinSettleLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleUpdateClVaultMinSettleLimit represents a UpdateClVaultMinSettleLimit event raised by the Oracle contract.
type OracleUpdateClVaultMinSettleLimit struct {
	ClVaultMinSettleLimit *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUpdateClVaultMinSettleLimit is a free log retrieval operation binding the contract event 0x6765ab460e8139130e86a064a184d24413dd588c41aa2277c433d2d09832f39e.
//
// Solidity: event UpdateClVaultMinSettleLimit(uint256 clVaultMinSettleLimit)
func (_Oracle *OracleFilterer) FilterUpdateClVaultMinSettleLimit(opts *bind.FilterOpts) (*OracleUpdateClVaultMinSettleLimitIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "UpdateClVaultMinSettleLimit")
	if err != nil {
		return nil, err
	}
	return &OracleUpdateClVaultMinSettleLimitIterator{contract: _Oracle.contract, event: "UpdateClVaultMinSettleLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateClVaultMinSettleLimit is a free log subscription operation binding the contract event 0x6765ab460e8139130e86a064a184d24413dd588c41aa2277c433d2d09832f39e.
//
// Solidity: event UpdateClVaultMinSettleLimit(uint256 clVaultMinSettleLimit)
func (_Oracle *OracleFilterer) WatchUpdateClVaultMinSettleLimit(opts *bind.WatchOpts, sink chan<- *OracleUpdateClVaultMinSettleLimit) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "UpdateClVaultMinSettleLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleUpdateClVaultMinSettleLimit)
				if err := _Oracle.contract.UnpackLog(event, "UpdateClVaultMinSettleLimit", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseUpdateClVaultMinSettleLimit(log types.Log) (*OracleUpdateClVaultMinSettleLimit, error) {
	event := new(OracleUpdateClVaultMinSettleLimit)
	if err := _Oracle.contract.UnpackLog(event, "UpdateClVaultMinSettleLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleUpdateExitRequestLimitIterator is returned from FilterUpdateExitRequestLimit and is used to iterate over the raw logs and unpacked data for UpdateExitRequestLimit events raised by the Oracle contract.
type OracleUpdateExitRequestLimitIterator struct {
	Event *OracleUpdateExitRequestLimit // Event containing the contract specifics and raw log

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
func (it *OracleUpdateExitRequestLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleUpdateExitRequestLimit)
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
		it.Event = new(OracleUpdateExitRequestLimit)
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
func (it *OracleUpdateExitRequestLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleUpdateExitRequestLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleUpdateExitRequestLimit represents a UpdateExitRequestLimit event raised by the Oracle contract.
type OracleUpdateExitRequestLimit struct {
	ExitRequestLimit *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateExitRequestLimit is a free log retrieval operation binding the contract event 0x6c79895328b404c7bf1115a4be7c52a4a974c20b13547a4fe873d5f997638c26.
//
// Solidity: event UpdateExitRequestLimit(uint256 exitRequestLimit)
func (_Oracle *OracleFilterer) FilterUpdateExitRequestLimit(opts *bind.FilterOpts) (*OracleUpdateExitRequestLimitIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "UpdateExitRequestLimit")
	if err != nil {
		return nil, err
	}
	return &OracleUpdateExitRequestLimitIterator{contract: _Oracle.contract, event: "UpdateExitRequestLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateExitRequestLimit is a free log subscription operation binding the contract event 0x6c79895328b404c7bf1115a4be7c52a4a974c20b13547a4fe873d5f997638c26.
//
// Solidity: event UpdateExitRequestLimit(uint256 exitRequestLimit)
func (_Oracle *OracleFilterer) WatchUpdateExitRequestLimit(opts *bind.WatchOpts, sink chan<- *OracleUpdateExitRequestLimit) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "UpdateExitRequestLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleUpdateExitRequestLimit)
				if err := _Oracle.contract.UnpackLog(event, "UpdateExitRequestLimit", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseUpdateExitRequestLimit(log types.Log) (*OracleUpdateExitRequestLimit, error) {
	event := new(OracleUpdateExitRequestLimit)
	if err := _Oracle.contract.UnpackLog(event, "UpdateExitRequestLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Oracle contract.
type OracleUpgradedIterator struct {
	Event *OracleUpgraded // Event containing the contract specifics and raw log

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
func (it *OracleUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleUpgraded)
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
		it.Event = new(OracleUpgraded)
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
func (it *OracleUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleUpgraded represents a Upgraded event raised by the Oracle contract.
type OracleUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Oracle *OracleFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*OracleUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &OracleUpgradedIterator{contract: _Oracle.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Oracle *OracleFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *OracleUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleUpgraded)
				if err := _Oracle.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseUpgraded(log types.Log) (*OracleUpgraded, error) {
	event := new(OracleUpgraded)
	if err := _Oracle.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleVaultManagerChangedIterator is returned from FilterVaultManagerChanged and is used to iterate over the raw logs and unpacked data for VaultManagerChanged events raised by the Oracle contract.
type OracleVaultManagerChangedIterator struct {
	Event *OracleVaultManagerChanged // Event containing the contract specifics and raw log

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
func (it *OracleVaultManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleVaultManagerChanged)
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
		it.Event = new(OracleVaultManagerChanged)
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
func (it *OracleVaultManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleVaultManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleVaultManagerChanged represents a VaultManagerChanged event raised by the Oracle contract.
type OracleVaultManagerChanged struct {
	Before common.Address
	After  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVaultManagerChanged is a free log retrieval operation binding the contract event 0x9e8e45be62c510326288c84df8d83e42380c150d3181f5fb5e7a1be5958ad67a.
//
// Solidity: event VaultManagerChanged(address _before, address _after)
func (_Oracle *OracleFilterer) FilterVaultManagerChanged(opts *bind.FilterOpts) (*OracleVaultManagerChangedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "VaultManagerChanged")
	if err != nil {
		return nil, err
	}
	return &OracleVaultManagerChangedIterator{contract: _Oracle.contract, event: "VaultManagerChanged", logs: logs, sub: sub}, nil
}

// WatchVaultManagerChanged is a free log subscription operation binding the contract event 0x9e8e45be62c510326288c84df8d83e42380c150d3181f5fb5e7a1be5958ad67a.
//
// Solidity: event VaultManagerChanged(address _before, address _after)
func (_Oracle *OracleFilterer) WatchVaultManagerChanged(opts *bind.WatchOpts, sink chan<- *OracleVaultManagerChanged) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "VaultManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleVaultManagerChanged)
				if err := _Oracle.contract.UnpackLog(event, "VaultManagerChanged", log); err != nil {
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
// Solidity: event VaultManagerChanged(address _before, address _after)
func (_Oracle *OracleFilterer) ParseVaultManagerChanged(log types.Log) (*OracleVaultManagerChanged, error) {
	event := new(OracleVaultManagerChanged)
	if err := _Oracle.contract.UnpackLog(event, "VaultManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleWarnDataIncompleteProcessingIterator is returned from FilterWarnDataIncompleteProcessing and is used to iterate over the raw logs and unpacked data for WarnDataIncompleteProcessing events raised by the Oracle contract.
type OracleWarnDataIncompleteProcessingIterator struct {
	Event *OracleWarnDataIncompleteProcessing // Event containing the contract specifics and raw log

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
func (it *OracleWarnDataIncompleteProcessingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleWarnDataIncompleteProcessing)
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
		it.Event = new(OracleWarnDataIncompleteProcessing)
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
func (it *OracleWarnDataIncompleteProcessingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleWarnDataIncompleteProcessingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleWarnDataIncompleteProcessing represents a WarnDataIncompleteProcessing event raised by the Oracle contract.
type OracleWarnDataIncompleteProcessing struct {
	RefSlot           *big.Int
	ExitRequestLimit  *big.Int
	ReportExitedCount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWarnDataIncompleteProcessing is a free log retrieval operation binding the contract event 0xefc67aab43195093a8d8ed25d52281d96de480748ece2787888c586e8e1e79b4.
//
// Solidity: event WarnDataIncompleteProcessing(uint256 indexed refSlot, uint256 exitRequestLimit, uint256 reportExitedCount)
func (_Oracle *OracleFilterer) FilterWarnDataIncompleteProcessing(opts *bind.FilterOpts, refSlot []*big.Int) (*OracleWarnDataIncompleteProcessingIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "WarnDataIncompleteProcessing", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &OracleWarnDataIncompleteProcessingIterator{contract: _Oracle.contract, event: "WarnDataIncompleteProcessing", logs: logs, sub: sub}, nil
}

// WatchWarnDataIncompleteProcessing is a free log subscription operation binding the contract event 0xefc67aab43195093a8d8ed25d52281d96de480748ece2787888c586e8e1e79b4.
//
// Solidity: event WarnDataIncompleteProcessing(uint256 indexed refSlot, uint256 exitRequestLimit, uint256 reportExitedCount)
func (_Oracle *OracleFilterer) WatchWarnDataIncompleteProcessing(opts *bind.WatchOpts, sink chan<- *OracleWarnDataIncompleteProcessing, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "WarnDataIncompleteProcessing", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleWarnDataIncompleteProcessing)
				if err := _Oracle.contract.UnpackLog(event, "WarnDataIncompleteProcessing", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseWarnDataIncompleteProcessing(log types.Log) (*OracleWarnDataIncompleteProcessing, error) {
	event := new(OracleWarnDataIncompleteProcessing)
	if err := _Oracle.contract.UnpackLog(event, "WarnDataIncompleteProcessing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleWarnProcessingMissedIterator is returned from FilterWarnProcessingMissed and is used to iterate over the raw logs and unpacked data for WarnProcessingMissed events raised by the Oracle contract.
type OracleWarnProcessingMissedIterator struct {
	Event *OracleWarnProcessingMissed // Event containing the contract specifics and raw log

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
func (it *OracleWarnProcessingMissedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleWarnProcessingMissed)
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
		it.Event = new(OracleWarnProcessingMissed)
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
func (it *OracleWarnProcessingMissedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleWarnProcessingMissedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleWarnProcessingMissed represents a WarnProcessingMissed event raised by the Oracle contract.
type OracleWarnProcessingMissed struct {
	RefSlot *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWarnProcessingMissed is a free log retrieval operation binding the contract event 0x800b849c8bf80718cf786c99d1091c079fe2c5e420a3ba7ba9b0ef8179ef2c38.
//
// Solidity: event WarnProcessingMissed(uint256 indexed refSlot)
func (_Oracle *OracleFilterer) FilterWarnProcessingMissed(opts *bind.FilterOpts, refSlot []*big.Int) (*OracleWarnProcessingMissedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "WarnProcessingMissed", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &OracleWarnProcessingMissedIterator{contract: _Oracle.contract, event: "WarnProcessingMissed", logs: logs, sub: sub}, nil
}

// WatchWarnProcessingMissed is a free log subscription operation binding the contract event 0x800b849c8bf80718cf786c99d1091c079fe2c5e420a3ba7ba9b0ef8179ef2c38.
//
// Solidity: event WarnProcessingMissed(uint256 indexed refSlot)
func (_Oracle *OracleFilterer) WatchWarnProcessingMissed(opts *bind.WatchOpts, sink chan<- *OracleWarnProcessingMissed, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "WarnProcessingMissed", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleWarnProcessingMissed)
				if err := _Oracle.contract.UnpackLog(event, "WarnProcessingMissed", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseWarnProcessingMissed(log types.Log) (*OracleWarnProcessingMissed, error) {
	event := new(OracleWarnProcessingMissed)
	if err := _Oracle.contract.UnpackLog(event, "WarnProcessingMissed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

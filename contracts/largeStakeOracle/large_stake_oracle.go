// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package largeStakeOracle

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

// CLStakingExitInfo is an auto generated low-level Go binding around an user-defined struct.
type CLStakingExitInfo struct {
	StakingId *big.Int
	Pubkey    []byte
}

// CLStakingSlashInfo is an auto generated low-level Go binding around an user-defined struct.
type CLStakingSlashInfo struct {
	StakingId   *big.Int
	SlashAmount *big.Int
	Pubkey      []byte
}

// LargeStakeOracleProcessingState is an auto generated low-level Go binding around an user-defined struct.
type LargeStakeOracleProcessingState struct {
	CurrentFrameRefSlot    *big.Int
	ProcessingDeadlineTime *big.Int
	DataHash               [32]byte
	DataSubmitted          bool
}

// LargeStakeOracleReportData is an auto generated low-level Go binding around an user-defined struct.
type LargeStakeOracleReportData struct {
	ConsensusVersion    *big.Int
	RefSlot             *big.Int
	ClStakingExitInfos  []CLStakingExitInfo
	ClStakingSlashInfos []CLStakingSlashInfo
}

// LargeStakeOracleMetaData contains all meta data concerning the LargeStakeOracle contract.
var LargeStakeOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AddressCannotBeSame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DaoCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExitLimitNotZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingRefSlot\",\"type\":\"uint256\"}],\"name\":\"InitialRefSlotCannotBeLessThanProcessingOne\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContractVersionIncrement\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ModuleIdIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ModuleIdNotEqual\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroContractVersionOnInit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyConsensusContractCanSubmitReport\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverExitLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ProcessingDeadlineMissed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefSlotAlreadyProcessing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevRefSlot\",\"type\":\"uint256\"}],\"name\":\"RefSlotCannotDecrease\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingRefSlot\",\"type\":\"uint256\"}],\"name\":\"RefSlotMustBeGreaterThanProcessingOne\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReportDataIsEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedChainConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedVersion\",\"type\":\"uint256\"}],\"name\":\"UnexpectedConsensusVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"UnexpectedContractVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"consensusHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receivedHash\",\"type\":\"bytes32\"}],\"name\":\"UnexpectedDataHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"consensusRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataRefSlot\",\"type\":\"uint256\"}],\"name\":\"UnexpectedRefSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionCannotBeSame\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevAddr\",\"type\":\"address\"}],\"name\":\"ConsensusHashContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prevVersion\",\"type\":\"uint256\"}],\"name\":\"ConsensusVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"ContractVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldDao\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"DaoAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldLargeStake\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLargeStake\",\"type\":\"address\"}],\"name\":\"LargeStakeContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"consensusVersion\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"stakingId\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structCLStakingExitInfo[]\",\"name\":\"cLStakingExitInfos\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"stakingId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"slashAmount\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structCLStakingSlashInfo[]\",\"name\":\"cLStakingSlashInfos\",\"type\":\"tuple[]\"}],\"name\":\"ReportSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldExitLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newExitLimit\",\"type\":\"uint256\"}],\"name\":\"UpdateExitLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"}],\"name\":\"WarnProcessingMissed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GENESIS_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_PER_SLOT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exitLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusReport\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"processingStarted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastProcessingRefSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProcessingState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currentFrameRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"dataSubmitted\",\"type\":\"bool\"}],\"internalType\":\"structLargeStakeOracle.ProcessingState\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"secondsPerSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"genesisTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consensusContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"consensusVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastProcessingRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_largeStakeContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"largeStakeContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setConsensusContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"setConsensusVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"setDaoAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_exitLimit\",\"type\":\"uint256\"}],\"name\":\"setExitLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_largeStakeContractAddress\",\"type\":\"address\"}],\"name\":\"setLargeStakeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reportHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_moduleId\",\"type\":\"uint256\"}],\"name\":\"submitConsensusReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"consensusVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"stakingId\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"internalType\":\"structCLStakingExitInfo[]\",\"name\":\"clStakingExitInfos\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"stakingId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"slashAmount\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"internalType\":\"structCLStakingSlashInfo[]\",\"name\":\"clStakingSlashInfos\",\"type\":\"tuple[]\"}],\"internalType\":\"structLargeStakeOracle.ReportData\",\"name\":\"data\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_contractVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_moduleId\",\"type\":\"uint256\"}],\"name\":\"submitReportData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateContractVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// LargeStakeOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use LargeStakeOracleMetaData.ABI instead.
var LargeStakeOracleABI = LargeStakeOracleMetaData.ABI

// LargeStakeOracle is an auto generated Go binding around an Ethereum contract.
type LargeStakeOracle struct {
	LargeStakeOracleCaller     // Read-only binding to the contract
	LargeStakeOracleTransactor // Write-only binding to the contract
	LargeStakeOracleFilterer   // Log filterer for contract events
}

// LargeStakeOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type LargeStakeOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LargeStakeOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LargeStakeOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LargeStakeOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LargeStakeOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LargeStakeOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LargeStakeOracleSession struct {
	Contract     *LargeStakeOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LargeStakeOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LargeStakeOracleCallerSession struct {
	Contract *LargeStakeOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LargeStakeOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LargeStakeOracleTransactorSession struct {
	Contract     *LargeStakeOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LargeStakeOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type LargeStakeOracleRaw struct {
	Contract *LargeStakeOracle // Generic contract binding to access the raw methods on
}

// LargeStakeOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LargeStakeOracleCallerRaw struct {
	Contract *LargeStakeOracleCaller // Generic read-only contract binding to access the raw methods on
}

// LargeStakeOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LargeStakeOracleTransactorRaw struct {
	Contract *LargeStakeOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLargeStakeOracle creates a new instance of LargeStakeOracle, bound to a specific deployed contract.
func NewLargeStakeOracle(address common.Address, backend bind.ContractBackend) (*LargeStakeOracle, error) {
	contract, err := bindLargeStakeOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LargeStakeOracle{LargeStakeOracleCaller: LargeStakeOracleCaller{contract: contract}, LargeStakeOracleTransactor: LargeStakeOracleTransactor{contract: contract}, LargeStakeOracleFilterer: LargeStakeOracleFilterer{contract: contract}}, nil
}

// NewLargeStakeOracleCaller creates a new read-only instance of LargeStakeOracle, bound to a specific deployed contract.
func NewLargeStakeOracleCaller(address common.Address, caller bind.ContractCaller) (*LargeStakeOracleCaller, error) {
	contract, err := bindLargeStakeOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LargeStakeOracleCaller{contract: contract}, nil
}

// NewLargeStakeOracleTransactor creates a new write-only instance of LargeStakeOracle, bound to a specific deployed contract.
func NewLargeStakeOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*LargeStakeOracleTransactor, error) {
	contract, err := bindLargeStakeOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LargeStakeOracleTransactor{contract: contract}, nil
}

// NewLargeStakeOracleFilterer creates a new log filterer instance of LargeStakeOracle, bound to a specific deployed contract.
func NewLargeStakeOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*LargeStakeOracleFilterer, error) {
	contract, err := bindLargeStakeOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LargeStakeOracleFilterer{contract: contract}, nil
}

// bindLargeStakeOracle binds a generic wrapper to an already deployed contract.
func bindLargeStakeOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LargeStakeOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LargeStakeOracle *LargeStakeOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LargeStakeOracle.Contract.LargeStakeOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LargeStakeOracle *LargeStakeOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.LargeStakeOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LargeStakeOracle *LargeStakeOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.LargeStakeOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LargeStakeOracle *LargeStakeOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LargeStakeOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LargeStakeOracle *LargeStakeOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LargeStakeOracle *LargeStakeOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.contract.Transact(opts, method, params...)
}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint256)
func (_LargeStakeOracle *LargeStakeOracleCaller) GENESISTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LargeStakeOracle.contract.Call(opts, &out, "GENESIS_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint256)
func (_LargeStakeOracle *LargeStakeOracleSession) GENESISTIME() (*big.Int, error) {
	return _LargeStakeOracle.Contract.GENESISTIME(&_LargeStakeOracle.CallOpts)
}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint256)
func (_LargeStakeOracle *LargeStakeOracleCallerSession) GENESISTIME() (*big.Int, error) {
	return _LargeStakeOracle.Contract.GENESISTIME(&_LargeStakeOracle.CallOpts)
}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint256)
func (_LargeStakeOracle *LargeStakeOracleCaller) SECONDSPERSLOT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LargeStakeOracle.contract.Call(opts, &out, "SECONDS_PER_SLOT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint256)
func (_LargeStakeOracle *LargeStakeOracleSession) SECONDSPERSLOT() (*big.Int, error) {
	return _LargeStakeOracle.Contract.SECONDSPERSLOT(&_LargeStakeOracle.CallOpts)
}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint256)
func (_LargeStakeOracle *LargeStakeOracleCallerSession) SECONDSPERSLOT() (*big.Int, error) {
	return _LargeStakeOracle.Contract.SECONDSPERSLOT(&_LargeStakeOracle.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_LargeStakeOracle *LargeStakeOracleCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LargeStakeOracle.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_LargeStakeOracle *LargeStakeOracleSession) Dao() (common.Address, error) {
	return _LargeStakeOracle.Contract.Dao(&_LargeStakeOracle.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_LargeStakeOracle *LargeStakeOracleCallerSession) Dao() (common.Address, error) {
	return _LargeStakeOracle.Contract.Dao(&_LargeStakeOracle.CallOpts)
}

// ExitLimit is a free data retrieval call binding the contract method 0x44ed9f66.
//
// Solidity: function exitLimit() view returns(uint256)
func (_LargeStakeOracle *LargeStakeOracleCaller) ExitLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LargeStakeOracle.contract.Call(opts, &out, "exitLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExitLimit is a free data retrieval call binding the contract method 0x44ed9f66.
//
// Solidity: function exitLimit() view returns(uint256)
func (_LargeStakeOracle *LargeStakeOracleSession) ExitLimit() (*big.Int, error) {
	return _LargeStakeOracle.Contract.ExitLimit(&_LargeStakeOracle.CallOpts)
}

// ExitLimit is a free data retrieval call binding the contract method 0x44ed9f66.
//
// Solidity: function exitLimit() view returns(uint256)
func (_LargeStakeOracle *LargeStakeOracleCallerSession) ExitLimit() (*big.Int, error) {
	return _LargeStakeOracle.Contract.ExitLimit(&_LargeStakeOracle.CallOpts)
}

// GetConsensusContract is a free data retrieval call binding the contract method 0x8f55b571.
//
// Solidity: function getConsensusContract() view returns(address)
func (_LargeStakeOracle *LargeStakeOracleCaller) GetConsensusContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LargeStakeOracle.contract.Call(opts, &out, "getConsensusContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetConsensusContract is a free data retrieval call binding the contract method 0x8f55b571.
//
// Solidity: function getConsensusContract() view returns(address)
func (_LargeStakeOracle *LargeStakeOracleSession) GetConsensusContract() (common.Address, error) {
	return _LargeStakeOracle.Contract.GetConsensusContract(&_LargeStakeOracle.CallOpts)
}

// GetConsensusContract is a free data retrieval call binding the contract method 0x8f55b571.
//
// Solidity: function getConsensusContract() view returns(address)
func (_LargeStakeOracle *LargeStakeOracleCallerSession) GetConsensusContract() (common.Address, error) {
	return _LargeStakeOracle.Contract.GetConsensusContract(&_LargeStakeOracle.CallOpts)
}

// GetConsensusReport is a free data retrieval call binding the contract method 0x60d64d38.
//
// Solidity: function getConsensusReport() view returns(bytes32 hash, uint256 refSlot, uint256 processingDeadlineTime, bool processingStarted)
func (_LargeStakeOracle *LargeStakeOracleCaller) GetConsensusReport(opts *bind.CallOpts) (struct {
	Hash                   [32]byte
	RefSlot                *big.Int
	ProcessingDeadlineTime *big.Int
	ProcessingStarted      bool
}, error) {
	var out []interface{}
	err := _LargeStakeOracle.contract.Call(opts, &out, "getConsensusReport")

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
func (_LargeStakeOracle *LargeStakeOracleSession) GetConsensusReport() (struct {
	Hash                   [32]byte
	RefSlot                *big.Int
	ProcessingDeadlineTime *big.Int
	ProcessingStarted      bool
}, error) {
	return _LargeStakeOracle.Contract.GetConsensusReport(&_LargeStakeOracle.CallOpts)
}

// GetConsensusReport is a free data retrieval call binding the contract method 0x60d64d38.
//
// Solidity: function getConsensusReport() view returns(bytes32 hash, uint256 refSlot, uint256 processingDeadlineTime, bool processingStarted)
func (_LargeStakeOracle *LargeStakeOracleCallerSession) GetConsensusReport() (struct {
	Hash                   [32]byte
	RefSlot                *big.Int
	ProcessingDeadlineTime *big.Int
	ProcessingStarted      bool
}, error) {
	return _LargeStakeOracle.Contract.GetConsensusReport(&_LargeStakeOracle.CallOpts)
}

// GetConsensusVersion is a free data retrieval call binding the contract method 0x5be20425.
//
// Solidity: function getConsensusVersion() view returns(uint256)
func (_LargeStakeOracle *LargeStakeOracleCaller) GetConsensusVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LargeStakeOracle.contract.Call(opts, &out, "getConsensusVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConsensusVersion is a free data retrieval call binding the contract method 0x5be20425.
//
// Solidity: function getConsensusVersion() view returns(uint256)
func (_LargeStakeOracle *LargeStakeOracleSession) GetConsensusVersion() (*big.Int, error) {
	return _LargeStakeOracle.Contract.GetConsensusVersion(&_LargeStakeOracle.CallOpts)
}

// GetConsensusVersion is a free data retrieval call binding the contract method 0x5be20425.
//
// Solidity: function getConsensusVersion() view returns(uint256)
func (_LargeStakeOracle *LargeStakeOracleCallerSession) GetConsensusVersion() (*big.Int, error) {
	return _LargeStakeOracle.Contract.GetConsensusVersion(&_LargeStakeOracle.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_LargeStakeOracle *LargeStakeOracleCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LargeStakeOracle.contract.Call(opts, &out, "getContractVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_LargeStakeOracle *LargeStakeOracleSession) GetContractVersion() (*big.Int, error) {
	return _LargeStakeOracle.Contract.GetContractVersion(&_LargeStakeOracle.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_LargeStakeOracle *LargeStakeOracleCallerSession) GetContractVersion() (*big.Int, error) {
	return _LargeStakeOracle.Contract.GetContractVersion(&_LargeStakeOracle.CallOpts)
}

// GetLastProcessingRefSlot is a free data retrieval call binding the contract method 0x3584d59c.
//
// Solidity: function getLastProcessingRefSlot() view returns(uint256)
func (_LargeStakeOracle *LargeStakeOracleCaller) GetLastProcessingRefSlot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LargeStakeOracle.contract.Call(opts, &out, "getLastProcessingRefSlot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastProcessingRefSlot is a free data retrieval call binding the contract method 0x3584d59c.
//
// Solidity: function getLastProcessingRefSlot() view returns(uint256)
func (_LargeStakeOracle *LargeStakeOracleSession) GetLastProcessingRefSlot() (*big.Int, error) {
	return _LargeStakeOracle.Contract.GetLastProcessingRefSlot(&_LargeStakeOracle.CallOpts)
}

// GetLastProcessingRefSlot is a free data retrieval call binding the contract method 0x3584d59c.
//
// Solidity: function getLastProcessingRefSlot() view returns(uint256)
func (_LargeStakeOracle *LargeStakeOracleCallerSession) GetLastProcessingRefSlot() (*big.Int, error) {
	return _LargeStakeOracle.Contract.GetLastProcessingRefSlot(&_LargeStakeOracle.CallOpts)
}

// GetProcessingState is a free data retrieval call binding the contract method 0x8f7797c2.
//
// Solidity: function getProcessingState() view returns((uint256,uint256,bytes32,bool) result)
func (_LargeStakeOracle *LargeStakeOracleCaller) GetProcessingState(opts *bind.CallOpts) (LargeStakeOracleProcessingState, error) {
	var out []interface{}
	err := _LargeStakeOracle.contract.Call(opts, &out, "getProcessingState")

	if err != nil {
		return *new(LargeStakeOracleProcessingState), err
	}

	out0 := *abi.ConvertType(out[0], new(LargeStakeOracleProcessingState)).(*LargeStakeOracleProcessingState)

	return out0, err

}

// GetProcessingState is a free data retrieval call binding the contract method 0x8f7797c2.
//
// Solidity: function getProcessingState() view returns((uint256,uint256,bytes32,bool) result)
func (_LargeStakeOracle *LargeStakeOracleSession) GetProcessingState() (LargeStakeOracleProcessingState, error) {
	return _LargeStakeOracle.Contract.GetProcessingState(&_LargeStakeOracle.CallOpts)
}

// GetProcessingState is a free data retrieval call binding the contract method 0x8f7797c2.
//
// Solidity: function getProcessingState() view returns((uint256,uint256,bytes32,bool) result)
func (_LargeStakeOracle *LargeStakeOracleCallerSession) GetProcessingState() (LargeStakeOracleProcessingState, error) {
	return _LargeStakeOracle.Contract.GetProcessingState(&_LargeStakeOracle.CallOpts)
}

// LargeStakeContract is a free data retrieval call binding the contract method 0xf7f14113.
//
// Solidity: function largeStakeContract() view returns(address)
func (_LargeStakeOracle *LargeStakeOracleCaller) LargeStakeContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LargeStakeOracle.contract.Call(opts, &out, "largeStakeContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LargeStakeContract is a free data retrieval call binding the contract method 0xf7f14113.
//
// Solidity: function largeStakeContract() view returns(address)
func (_LargeStakeOracle *LargeStakeOracleSession) LargeStakeContract() (common.Address, error) {
	return _LargeStakeOracle.Contract.LargeStakeContract(&_LargeStakeOracle.CallOpts)
}

// LargeStakeContract is a free data retrieval call binding the contract method 0xf7f14113.
//
// Solidity: function largeStakeContract() view returns(address)
func (_LargeStakeOracle *LargeStakeOracleCallerSession) LargeStakeContract() (common.Address, error) {
	return _LargeStakeOracle.Contract.LargeStakeContract(&_LargeStakeOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LargeStakeOracle *LargeStakeOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LargeStakeOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LargeStakeOracle *LargeStakeOracleSession) Owner() (common.Address, error) {
	return _LargeStakeOracle.Contract.Owner(&_LargeStakeOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LargeStakeOracle *LargeStakeOracleCallerSession) Owner() (common.Address, error) {
	return _LargeStakeOracle.Contract.Owner(&_LargeStakeOracle.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LargeStakeOracle *LargeStakeOracleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LargeStakeOracle.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LargeStakeOracle *LargeStakeOracleSession) Paused() (bool, error) {
	return _LargeStakeOracle.Contract.Paused(&_LargeStakeOracle.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LargeStakeOracle *LargeStakeOracleCallerSession) Paused() (bool, error) {
	return _LargeStakeOracle.Contract.Paused(&_LargeStakeOracle.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LargeStakeOracle *LargeStakeOracleCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LargeStakeOracle.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LargeStakeOracle *LargeStakeOracleSession) ProxiableUUID() ([32]byte, error) {
	return _LargeStakeOracle.Contract.ProxiableUUID(&_LargeStakeOracle.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LargeStakeOracle *LargeStakeOracleCallerSession) ProxiableUUID() ([32]byte, error) {
	return _LargeStakeOracle.Contract.ProxiableUUID(&_LargeStakeOracle.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x47a92421.
//
// Solidity: function initialize(uint256 secondsPerSlot, uint256 genesisTime, address consensusContract, uint256 consensusVersion, uint256 lastProcessingRefSlot, address _dao, address _largeStakeContract) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactor) Initialize(opts *bind.TransactOpts, secondsPerSlot *big.Int, genesisTime *big.Int, consensusContract common.Address, consensusVersion *big.Int, lastProcessingRefSlot *big.Int, _dao common.Address, _largeStakeContract common.Address) (*types.Transaction, error) {
	return _LargeStakeOracle.contract.Transact(opts, "initialize", secondsPerSlot, genesisTime, consensusContract, consensusVersion, lastProcessingRefSlot, _dao, _largeStakeContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x47a92421.
//
// Solidity: function initialize(uint256 secondsPerSlot, uint256 genesisTime, address consensusContract, uint256 consensusVersion, uint256 lastProcessingRefSlot, address _dao, address _largeStakeContract) returns()
func (_LargeStakeOracle *LargeStakeOracleSession) Initialize(secondsPerSlot *big.Int, genesisTime *big.Int, consensusContract common.Address, consensusVersion *big.Int, lastProcessingRefSlot *big.Int, _dao common.Address, _largeStakeContract common.Address) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.Initialize(&_LargeStakeOracle.TransactOpts, secondsPerSlot, genesisTime, consensusContract, consensusVersion, lastProcessingRefSlot, _dao, _largeStakeContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x47a92421.
//
// Solidity: function initialize(uint256 secondsPerSlot, uint256 genesisTime, address consensusContract, uint256 consensusVersion, uint256 lastProcessingRefSlot, address _dao, address _largeStakeContract) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactorSession) Initialize(secondsPerSlot *big.Int, genesisTime *big.Int, consensusContract common.Address, consensusVersion *big.Int, lastProcessingRefSlot *big.Int, _dao common.Address, _largeStakeContract common.Address) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.Initialize(&_LargeStakeOracle.TransactOpts, secondsPerSlot, genesisTime, consensusContract, consensusVersion, lastProcessingRefSlot, _dao, _largeStakeContract)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LargeStakeOracle *LargeStakeOracleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LargeStakeOracle.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LargeStakeOracle *LargeStakeOracleSession) Pause() (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.Pause(&_LargeStakeOracle.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LargeStakeOracle *LargeStakeOracleTransactorSession) Pause() (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.Pause(&_LargeStakeOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LargeStakeOracle *LargeStakeOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LargeStakeOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LargeStakeOracle *LargeStakeOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.RenounceOwnership(&_LargeStakeOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LargeStakeOracle *LargeStakeOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.RenounceOwnership(&_LargeStakeOracle.TransactOpts)
}

// SetConsensusContract is a paid mutator transaction binding the contract method 0xc469c307.
//
// Solidity: function setConsensusContract(address addr) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactor) SetConsensusContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _LargeStakeOracle.contract.Transact(opts, "setConsensusContract", addr)
}

// SetConsensusContract is a paid mutator transaction binding the contract method 0xc469c307.
//
// Solidity: function setConsensusContract(address addr) returns()
func (_LargeStakeOracle *LargeStakeOracleSession) SetConsensusContract(addr common.Address) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.SetConsensusContract(&_LargeStakeOracle.TransactOpts, addr)
}

// SetConsensusContract is a paid mutator transaction binding the contract method 0xc469c307.
//
// Solidity: function setConsensusContract(address addr) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactorSession) SetConsensusContract(addr common.Address) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.SetConsensusContract(&_LargeStakeOracle.TransactOpts, addr)
}

// SetConsensusVersion is a paid mutator transaction binding the contract method 0x8d591474.
//
// Solidity: function setConsensusVersion(uint256 version) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactor) SetConsensusVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _LargeStakeOracle.contract.Transact(opts, "setConsensusVersion", version)
}

// SetConsensusVersion is a paid mutator transaction binding the contract method 0x8d591474.
//
// Solidity: function setConsensusVersion(uint256 version) returns()
func (_LargeStakeOracle *LargeStakeOracleSession) SetConsensusVersion(version *big.Int) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.SetConsensusVersion(&_LargeStakeOracle.TransactOpts, version)
}

// SetConsensusVersion is a paid mutator transaction binding the contract method 0x8d591474.
//
// Solidity: function setConsensusVersion(uint256 version) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactorSession) SetConsensusVersion(version *big.Int) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.SetConsensusVersion(&_LargeStakeOracle.TransactOpts, version)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactor) SetDaoAddress(opts *bind.TransactOpts, _dao common.Address) (*types.Transaction, error) {
	return _LargeStakeOracle.contract.Transact(opts, "setDaoAddress", _dao)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_LargeStakeOracle *LargeStakeOracleSession) SetDaoAddress(_dao common.Address) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.SetDaoAddress(&_LargeStakeOracle.TransactOpts, _dao)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactorSession) SetDaoAddress(_dao common.Address) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.SetDaoAddress(&_LargeStakeOracle.TransactOpts, _dao)
}

// SetExitLimit is a paid mutator transaction binding the contract method 0x7f5fcec3.
//
// Solidity: function setExitLimit(uint256 _exitLimit) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactor) SetExitLimit(opts *bind.TransactOpts, _exitLimit *big.Int) (*types.Transaction, error) {
	return _LargeStakeOracle.contract.Transact(opts, "setExitLimit", _exitLimit)
}

// SetExitLimit is a paid mutator transaction binding the contract method 0x7f5fcec3.
//
// Solidity: function setExitLimit(uint256 _exitLimit) returns()
func (_LargeStakeOracle *LargeStakeOracleSession) SetExitLimit(_exitLimit *big.Int) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.SetExitLimit(&_LargeStakeOracle.TransactOpts, _exitLimit)
}

// SetExitLimit is a paid mutator transaction binding the contract method 0x7f5fcec3.
//
// Solidity: function setExitLimit(uint256 _exitLimit) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactorSession) SetExitLimit(_exitLimit *big.Int) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.SetExitLimit(&_LargeStakeOracle.TransactOpts, _exitLimit)
}

// SetLargeStakeContract is a paid mutator transaction binding the contract method 0x7bb37b6e.
//
// Solidity: function setLargeStakeContract(address _largeStakeContractAddress) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactor) SetLargeStakeContract(opts *bind.TransactOpts, _largeStakeContractAddress common.Address) (*types.Transaction, error) {
	return _LargeStakeOracle.contract.Transact(opts, "setLargeStakeContract", _largeStakeContractAddress)
}

// SetLargeStakeContract is a paid mutator transaction binding the contract method 0x7bb37b6e.
//
// Solidity: function setLargeStakeContract(address _largeStakeContractAddress) returns()
func (_LargeStakeOracle *LargeStakeOracleSession) SetLargeStakeContract(_largeStakeContractAddress common.Address) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.SetLargeStakeContract(&_LargeStakeOracle.TransactOpts, _largeStakeContractAddress)
}

// SetLargeStakeContract is a paid mutator transaction binding the contract method 0x7bb37b6e.
//
// Solidity: function setLargeStakeContract(address _largeStakeContractAddress) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactorSession) SetLargeStakeContract(_largeStakeContractAddress common.Address) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.SetLargeStakeContract(&_LargeStakeOracle.TransactOpts, _largeStakeContractAddress)
}

// SubmitConsensusReport is a paid mutator transaction binding the contract method 0x7bcc792b.
//
// Solidity: function submitConsensusReport(bytes32 reportHash, uint256 refSlot, uint256 deadline, uint256 _moduleId) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactor) SubmitConsensusReport(opts *bind.TransactOpts, reportHash [32]byte, refSlot *big.Int, deadline *big.Int, _moduleId *big.Int) (*types.Transaction, error) {
	return _LargeStakeOracle.contract.Transact(opts, "submitConsensusReport", reportHash, refSlot, deadline, _moduleId)
}

// SubmitConsensusReport is a paid mutator transaction binding the contract method 0x7bcc792b.
//
// Solidity: function submitConsensusReport(bytes32 reportHash, uint256 refSlot, uint256 deadline, uint256 _moduleId) returns()
func (_LargeStakeOracle *LargeStakeOracleSession) SubmitConsensusReport(reportHash [32]byte, refSlot *big.Int, deadline *big.Int, _moduleId *big.Int) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.SubmitConsensusReport(&_LargeStakeOracle.TransactOpts, reportHash, refSlot, deadline, _moduleId)
}

// SubmitConsensusReport is a paid mutator transaction binding the contract method 0x7bcc792b.
//
// Solidity: function submitConsensusReport(bytes32 reportHash, uint256 refSlot, uint256 deadline, uint256 _moduleId) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactorSession) SubmitConsensusReport(reportHash [32]byte, refSlot *big.Int, deadline *big.Int, _moduleId *big.Int) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.SubmitConsensusReport(&_LargeStakeOracle.TransactOpts, reportHash, refSlot, deadline, _moduleId)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0x9b3b3100.
//
// Solidity: function submitReportData((uint256,uint256,(uint128,bytes)[],(uint128,uint128,bytes)[]) data, uint256 _contractVersion, uint256 _moduleId) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactor) SubmitReportData(opts *bind.TransactOpts, data LargeStakeOracleReportData, _contractVersion *big.Int, _moduleId *big.Int) (*types.Transaction, error) {
	return _LargeStakeOracle.contract.Transact(opts, "submitReportData", data, _contractVersion, _moduleId)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0x9b3b3100.
//
// Solidity: function submitReportData((uint256,uint256,(uint128,bytes)[],(uint128,uint128,bytes)[]) data, uint256 _contractVersion, uint256 _moduleId) returns()
func (_LargeStakeOracle *LargeStakeOracleSession) SubmitReportData(data LargeStakeOracleReportData, _contractVersion *big.Int, _moduleId *big.Int) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.SubmitReportData(&_LargeStakeOracle.TransactOpts, data, _contractVersion, _moduleId)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0x9b3b3100.
//
// Solidity: function submitReportData((uint256,uint256,(uint128,bytes)[],(uint128,uint128,bytes)[]) data, uint256 _contractVersion, uint256 _moduleId) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactorSession) SubmitReportData(data LargeStakeOracleReportData, _contractVersion *big.Int, _moduleId *big.Int) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.SubmitReportData(&_LargeStakeOracle.TransactOpts, data, _contractVersion, _moduleId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LargeStakeOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LargeStakeOracle *LargeStakeOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.TransferOwnership(&_LargeStakeOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.TransferOwnership(&_LargeStakeOracle.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LargeStakeOracle *LargeStakeOracleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LargeStakeOracle.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LargeStakeOracle *LargeStakeOracleSession) Unpause() (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.Unpause(&_LargeStakeOracle.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LargeStakeOracle *LargeStakeOracleTransactorSession) Unpause() (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.Unpause(&_LargeStakeOracle.TransactOpts)
}

// UpdateContractVersion is a paid mutator transaction binding the contract method 0x294bb79a.
//
// Solidity: function updateContractVersion(uint256 version) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactor) UpdateContractVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _LargeStakeOracle.contract.Transact(opts, "updateContractVersion", version)
}

// UpdateContractVersion is a paid mutator transaction binding the contract method 0x294bb79a.
//
// Solidity: function updateContractVersion(uint256 version) returns()
func (_LargeStakeOracle *LargeStakeOracleSession) UpdateContractVersion(version *big.Int) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.UpdateContractVersion(&_LargeStakeOracle.TransactOpts, version)
}

// UpdateContractVersion is a paid mutator transaction binding the contract method 0x294bb79a.
//
// Solidity: function updateContractVersion(uint256 version) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactorSession) UpdateContractVersion(version *big.Int) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.UpdateContractVersion(&_LargeStakeOracle.TransactOpts, version)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _LargeStakeOracle.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LargeStakeOracle *LargeStakeOracleSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.UpgradeTo(&_LargeStakeOracle.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LargeStakeOracle *LargeStakeOracleTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.UpgradeTo(&_LargeStakeOracle.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LargeStakeOracle *LargeStakeOracleTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LargeStakeOracle.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LargeStakeOracle *LargeStakeOracleSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.UpgradeToAndCall(&_LargeStakeOracle.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LargeStakeOracle *LargeStakeOracleTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LargeStakeOracle.Contract.UpgradeToAndCall(&_LargeStakeOracle.TransactOpts, newImplementation, data)
}

// LargeStakeOracleAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the LargeStakeOracle contract.
type LargeStakeOracleAdminChangedIterator struct {
	Event *LargeStakeOracleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LargeStakeOracleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakeOracleAdminChanged)
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
		it.Event = new(LargeStakeOracleAdminChanged)
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
func (it *LargeStakeOracleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakeOracleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakeOracleAdminChanged represents a AdminChanged event raised by the LargeStakeOracle contract.
type LargeStakeOracleAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LargeStakeOracle *LargeStakeOracleFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*LargeStakeOracleAdminChangedIterator, error) {

	logs, sub, err := _LargeStakeOracle.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &LargeStakeOracleAdminChangedIterator{contract: _LargeStakeOracle.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LargeStakeOracle *LargeStakeOracleFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *LargeStakeOracleAdminChanged) (event.Subscription, error) {

	logs, sub, err := _LargeStakeOracle.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakeOracleAdminChanged)
				if err := _LargeStakeOracle.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_LargeStakeOracle *LargeStakeOracleFilterer) ParseAdminChanged(log types.Log) (*LargeStakeOracleAdminChanged, error) {
	event := new(LargeStakeOracleAdminChanged)
	if err := _LargeStakeOracle.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakeOracleBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the LargeStakeOracle contract.
type LargeStakeOracleBeaconUpgradedIterator struct {
	Event *LargeStakeOracleBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *LargeStakeOracleBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakeOracleBeaconUpgraded)
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
		it.Event = new(LargeStakeOracleBeaconUpgraded)
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
func (it *LargeStakeOracleBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakeOracleBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakeOracleBeaconUpgraded represents a BeaconUpgraded event raised by the LargeStakeOracle contract.
type LargeStakeOracleBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LargeStakeOracle *LargeStakeOracleFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*LargeStakeOracleBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _LargeStakeOracle.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &LargeStakeOracleBeaconUpgradedIterator{contract: _LargeStakeOracle.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LargeStakeOracle *LargeStakeOracleFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *LargeStakeOracleBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _LargeStakeOracle.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakeOracleBeaconUpgraded)
				if err := _LargeStakeOracle.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_LargeStakeOracle *LargeStakeOracleFilterer) ParseBeaconUpgraded(log types.Log) (*LargeStakeOracleBeaconUpgraded, error) {
	event := new(LargeStakeOracleBeaconUpgraded)
	if err := _LargeStakeOracle.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakeOracleConsensusHashContractSetIterator is returned from FilterConsensusHashContractSet and is used to iterate over the raw logs and unpacked data for ConsensusHashContractSet events raised by the LargeStakeOracle contract.
type LargeStakeOracleConsensusHashContractSetIterator struct {
	Event *LargeStakeOracleConsensusHashContractSet // Event containing the contract specifics and raw log

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
func (it *LargeStakeOracleConsensusHashContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakeOracleConsensusHashContractSet)
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
		it.Event = new(LargeStakeOracleConsensusHashContractSet)
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
func (it *LargeStakeOracleConsensusHashContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakeOracleConsensusHashContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakeOracleConsensusHashContractSet represents a ConsensusHashContractSet event raised by the LargeStakeOracle contract.
type LargeStakeOracleConsensusHashContractSet struct {
	Addr     common.Address
	PrevAddr common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConsensusHashContractSet is a free log retrieval operation binding the contract event 0x25421480fb7f52d18947876279a213696b58d7e0e5416ce5e2c9f9942661c34c.
//
// Solidity: event ConsensusHashContractSet(address indexed addr, address indexed prevAddr)
func (_LargeStakeOracle *LargeStakeOracleFilterer) FilterConsensusHashContractSet(opts *bind.FilterOpts, addr []common.Address, prevAddr []common.Address) (*LargeStakeOracleConsensusHashContractSetIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var prevAddrRule []interface{}
	for _, prevAddrItem := range prevAddr {
		prevAddrRule = append(prevAddrRule, prevAddrItem)
	}

	logs, sub, err := _LargeStakeOracle.contract.FilterLogs(opts, "ConsensusHashContractSet", addrRule, prevAddrRule)
	if err != nil {
		return nil, err
	}
	return &LargeStakeOracleConsensusHashContractSetIterator{contract: _LargeStakeOracle.contract, event: "ConsensusHashContractSet", logs: logs, sub: sub}, nil
}

// WatchConsensusHashContractSet is a free log subscription operation binding the contract event 0x25421480fb7f52d18947876279a213696b58d7e0e5416ce5e2c9f9942661c34c.
//
// Solidity: event ConsensusHashContractSet(address indexed addr, address indexed prevAddr)
func (_LargeStakeOracle *LargeStakeOracleFilterer) WatchConsensusHashContractSet(opts *bind.WatchOpts, sink chan<- *LargeStakeOracleConsensusHashContractSet, addr []common.Address, prevAddr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var prevAddrRule []interface{}
	for _, prevAddrItem := range prevAddr {
		prevAddrRule = append(prevAddrRule, prevAddrItem)
	}

	logs, sub, err := _LargeStakeOracle.contract.WatchLogs(opts, "ConsensusHashContractSet", addrRule, prevAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakeOracleConsensusHashContractSet)
				if err := _LargeStakeOracle.contract.UnpackLog(event, "ConsensusHashContractSet", log); err != nil {
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
func (_LargeStakeOracle *LargeStakeOracleFilterer) ParseConsensusHashContractSet(log types.Log) (*LargeStakeOracleConsensusHashContractSet, error) {
	event := new(LargeStakeOracleConsensusHashContractSet)
	if err := _LargeStakeOracle.contract.UnpackLog(event, "ConsensusHashContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakeOracleConsensusVersionSetIterator is returned from FilterConsensusVersionSet and is used to iterate over the raw logs and unpacked data for ConsensusVersionSet events raised by the LargeStakeOracle contract.
type LargeStakeOracleConsensusVersionSetIterator struct {
	Event *LargeStakeOracleConsensusVersionSet // Event containing the contract specifics and raw log

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
func (it *LargeStakeOracleConsensusVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakeOracleConsensusVersionSet)
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
		it.Event = new(LargeStakeOracleConsensusVersionSet)
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
func (it *LargeStakeOracleConsensusVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakeOracleConsensusVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakeOracleConsensusVersionSet represents a ConsensusVersionSet event raised by the LargeStakeOracle contract.
type LargeStakeOracleConsensusVersionSet struct {
	Version     *big.Int
	PrevVersion *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsensusVersionSet is a free log retrieval operation binding the contract event 0xfa5304972d4ec3e3207f0bbf91155a49d0dfa62488f9529403a2a49e4b29a895.
//
// Solidity: event ConsensusVersionSet(uint256 indexed version, uint256 indexed prevVersion)
func (_LargeStakeOracle *LargeStakeOracleFilterer) FilterConsensusVersionSet(opts *bind.FilterOpts, version []*big.Int, prevVersion []*big.Int) (*LargeStakeOracleConsensusVersionSetIterator, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var prevVersionRule []interface{}
	for _, prevVersionItem := range prevVersion {
		prevVersionRule = append(prevVersionRule, prevVersionItem)
	}

	logs, sub, err := _LargeStakeOracle.contract.FilterLogs(opts, "ConsensusVersionSet", versionRule, prevVersionRule)
	if err != nil {
		return nil, err
	}
	return &LargeStakeOracleConsensusVersionSetIterator{contract: _LargeStakeOracle.contract, event: "ConsensusVersionSet", logs: logs, sub: sub}, nil
}

// WatchConsensusVersionSet is a free log subscription operation binding the contract event 0xfa5304972d4ec3e3207f0bbf91155a49d0dfa62488f9529403a2a49e4b29a895.
//
// Solidity: event ConsensusVersionSet(uint256 indexed version, uint256 indexed prevVersion)
func (_LargeStakeOracle *LargeStakeOracleFilterer) WatchConsensusVersionSet(opts *bind.WatchOpts, sink chan<- *LargeStakeOracleConsensusVersionSet, version []*big.Int, prevVersion []*big.Int) (event.Subscription, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var prevVersionRule []interface{}
	for _, prevVersionItem := range prevVersion {
		prevVersionRule = append(prevVersionRule, prevVersionItem)
	}

	logs, sub, err := _LargeStakeOracle.contract.WatchLogs(opts, "ConsensusVersionSet", versionRule, prevVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakeOracleConsensusVersionSet)
				if err := _LargeStakeOracle.contract.UnpackLog(event, "ConsensusVersionSet", log); err != nil {
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
func (_LargeStakeOracle *LargeStakeOracleFilterer) ParseConsensusVersionSet(log types.Log) (*LargeStakeOracleConsensusVersionSet, error) {
	event := new(LargeStakeOracleConsensusVersionSet)
	if err := _LargeStakeOracle.contract.UnpackLog(event, "ConsensusVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakeOracleContractVersionSetIterator is returned from FilterContractVersionSet and is used to iterate over the raw logs and unpacked data for ContractVersionSet events raised by the LargeStakeOracle contract.
type LargeStakeOracleContractVersionSetIterator struct {
	Event *LargeStakeOracleContractVersionSet // Event containing the contract specifics and raw log

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
func (it *LargeStakeOracleContractVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakeOracleContractVersionSet)
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
		it.Event = new(LargeStakeOracleContractVersionSet)
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
func (it *LargeStakeOracleContractVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakeOracleContractVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakeOracleContractVersionSet represents a ContractVersionSet event raised by the LargeStakeOracle contract.
type LargeStakeOracleContractVersionSet struct {
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractVersionSet is a free log retrieval operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_LargeStakeOracle *LargeStakeOracleFilterer) FilterContractVersionSet(opts *bind.FilterOpts) (*LargeStakeOracleContractVersionSetIterator, error) {

	logs, sub, err := _LargeStakeOracle.contract.FilterLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return &LargeStakeOracleContractVersionSetIterator{contract: _LargeStakeOracle.contract, event: "ContractVersionSet", logs: logs, sub: sub}, nil
}

// WatchContractVersionSet is a free log subscription operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_LargeStakeOracle *LargeStakeOracleFilterer) WatchContractVersionSet(opts *bind.WatchOpts, sink chan<- *LargeStakeOracleContractVersionSet) (event.Subscription, error) {

	logs, sub, err := _LargeStakeOracle.contract.WatchLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakeOracleContractVersionSet)
				if err := _LargeStakeOracle.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
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
func (_LargeStakeOracle *LargeStakeOracleFilterer) ParseContractVersionSet(log types.Log) (*LargeStakeOracleContractVersionSet, error) {
	event := new(LargeStakeOracleContractVersionSet)
	if err := _LargeStakeOracle.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakeOracleDaoAddressChangedIterator is returned from FilterDaoAddressChanged and is used to iterate over the raw logs and unpacked data for DaoAddressChanged events raised by the LargeStakeOracle contract.
type LargeStakeOracleDaoAddressChangedIterator struct {
	Event *LargeStakeOracleDaoAddressChanged // Event containing the contract specifics and raw log

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
func (it *LargeStakeOracleDaoAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakeOracleDaoAddressChanged)
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
		it.Event = new(LargeStakeOracleDaoAddressChanged)
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
func (it *LargeStakeOracleDaoAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakeOracleDaoAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakeOracleDaoAddressChanged represents a DaoAddressChanged event raised by the LargeStakeOracle contract.
type LargeStakeOracleDaoAddressChanged struct {
	OldDao common.Address
	Dao    common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDaoAddressChanged is a free log retrieval operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address _oldDao, address _dao)
func (_LargeStakeOracle *LargeStakeOracleFilterer) FilterDaoAddressChanged(opts *bind.FilterOpts) (*LargeStakeOracleDaoAddressChangedIterator, error) {

	logs, sub, err := _LargeStakeOracle.contract.FilterLogs(opts, "DaoAddressChanged")
	if err != nil {
		return nil, err
	}
	return &LargeStakeOracleDaoAddressChangedIterator{contract: _LargeStakeOracle.contract, event: "DaoAddressChanged", logs: logs, sub: sub}, nil
}

// WatchDaoAddressChanged is a free log subscription operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address _oldDao, address _dao)
func (_LargeStakeOracle *LargeStakeOracleFilterer) WatchDaoAddressChanged(opts *bind.WatchOpts, sink chan<- *LargeStakeOracleDaoAddressChanged) (event.Subscription, error) {

	logs, sub, err := _LargeStakeOracle.contract.WatchLogs(opts, "DaoAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakeOracleDaoAddressChanged)
				if err := _LargeStakeOracle.contract.UnpackLog(event, "DaoAddressChanged", log); err != nil {
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
func (_LargeStakeOracle *LargeStakeOracleFilterer) ParseDaoAddressChanged(log types.Log) (*LargeStakeOracleDaoAddressChanged, error) {
	event := new(LargeStakeOracleDaoAddressChanged)
	if err := _LargeStakeOracle.contract.UnpackLog(event, "DaoAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakeOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LargeStakeOracle contract.
type LargeStakeOracleInitializedIterator struct {
	Event *LargeStakeOracleInitialized // Event containing the contract specifics and raw log

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
func (it *LargeStakeOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakeOracleInitialized)
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
		it.Event = new(LargeStakeOracleInitialized)
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
func (it *LargeStakeOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakeOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakeOracleInitialized represents a Initialized event raised by the LargeStakeOracle contract.
type LargeStakeOracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LargeStakeOracle *LargeStakeOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*LargeStakeOracleInitializedIterator, error) {

	logs, sub, err := _LargeStakeOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LargeStakeOracleInitializedIterator{contract: _LargeStakeOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LargeStakeOracle *LargeStakeOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LargeStakeOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _LargeStakeOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakeOracleInitialized)
				if err := _LargeStakeOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LargeStakeOracle *LargeStakeOracleFilterer) ParseInitialized(log types.Log) (*LargeStakeOracleInitialized, error) {
	event := new(LargeStakeOracleInitialized)
	if err := _LargeStakeOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakeOracleLargeStakeContractChangedIterator is returned from FilterLargeStakeContractChanged and is used to iterate over the raw logs and unpacked data for LargeStakeContractChanged events raised by the LargeStakeOracle contract.
type LargeStakeOracleLargeStakeContractChangedIterator struct {
	Event *LargeStakeOracleLargeStakeContractChanged // Event containing the contract specifics and raw log

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
func (it *LargeStakeOracleLargeStakeContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakeOracleLargeStakeContractChanged)
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
		it.Event = new(LargeStakeOracleLargeStakeContractChanged)
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
func (it *LargeStakeOracleLargeStakeContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakeOracleLargeStakeContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakeOracleLargeStakeContractChanged represents a LargeStakeContractChanged event raised by the LargeStakeOracle contract.
type LargeStakeOracleLargeStakeContractChanged struct {
	OldLargeStake common.Address
	NewLargeStake common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLargeStakeContractChanged is a free log retrieval operation binding the contract event 0x4488ff66d0bb7e9dd13642f8c726690f8656f04a5dcf78e5df63d4902e6cbbab.
//
// Solidity: event LargeStakeContractChanged(address oldLargeStake, address newLargeStake)
func (_LargeStakeOracle *LargeStakeOracleFilterer) FilterLargeStakeContractChanged(opts *bind.FilterOpts) (*LargeStakeOracleLargeStakeContractChangedIterator, error) {

	logs, sub, err := _LargeStakeOracle.contract.FilterLogs(opts, "LargeStakeContractChanged")
	if err != nil {
		return nil, err
	}
	return &LargeStakeOracleLargeStakeContractChangedIterator{contract: _LargeStakeOracle.contract, event: "LargeStakeContractChanged", logs: logs, sub: sub}, nil
}

// WatchLargeStakeContractChanged is a free log subscription operation binding the contract event 0x4488ff66d0bb7e9dd13642f8c726690f8656f04a5dcf78e5df63d4902e6cbbab.
//
// Solidity: event LargeStakeContractChanged(address oldLargeStake, address newLargeStake)
func (_LargeStakeOracle *LargeStakeOracleFilterer) WatchLargeStakeContractChanged(opts *bind.WatchOpts, sink chan<- *LargeStakeOracleLargeStakeContractChanged) (event.Subscription, error) {

	logs, sub, err := _LargeStakeOracle.contract.WatchLogs(opts, "LargeStakeContractChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakeOracleLargeStakeContractChanged)
				if err := _LargeStakeOracle.contract.UnpackLog(event, "LargeStakeContractChanged", log); err != nil {
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

// ParseLargeStakeContractChanged is a log parse operation binding the contract event 0x4488ff66d0bb7e9dd13642f8c726690f8656f04a5dcf78e5df63d4902e6cbbab.
//
// Solidity: event LargeStakeContractChanged(address oldLargeStake, address newLargeStake)
func (_LargeStakeOracle *LargeStakeOracleFilterer) ParseLargeStakeContractChanged(log types.Log) (*LargeStakeOracleLargeStakeContractChanged, error) {
	event := new(LargeStakeOracleLargeStakeContractChanged)
	if err := _LargeStakeOracle.contract.UnpackLog(event, "LargeStakeContractChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakeOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LargeStakeOracle contract.
type LargeStakeOracleOwnershipTransferredIterator struct {
	Event *LargeStakeOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LargeStakeOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakeOracleOwnershipTransferred)
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
		it.Event = new(LargeStakeOracleOwnershipTransferred)
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
func (it *LargeStakeOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakeOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakeOracleOwnershipTransferred represents a OwnershipTransferred event raised by the LargeStakeOracle contract.
type LargeStakeOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LargeStakeOracle *LargeStakeOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LargeStakeOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LargeStakeOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LargeStakeOracleOwnershipTransferredIterator{contract: _LargeStakeOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LargeStakeOracle *LargeStakeOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LargeStakeOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LargeStakeOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakeOracleOwnershipTransferred)
				if err := _LargeStakeOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LargeStakeOracle *LargeStakeOracleFilterer) ParseOwnershipTransferred(log types.Log) (*LargeStakeOracleOwnershipTransferred, error) {
	event := new(LargeStakeOracleOwnershipTransferred)
	if err := _LargeStakeOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakeOraclePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LargeStakeOracle contract.
type LargeStakeOraclePausedIterator struct {
	Event *LargeStakeOraclePaused // Event containing the contract specifics and raw log

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
func (it *LargeStakeOraclePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakeOraclePaused)
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
		it.Event = new(LargeStakeOraclePaused)
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
func (it *LargeStakeOraclePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakeOraclePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakeOraclePaused represents a Paused event raised by the LargeStakeOracle contract.
type LargeStakeOraclePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LargeStakeOracle *LargeStakeOracleFilterer) FilterPaused(opts *bind.FilterOpts) (*LargeStakeOraclePausedIterator, error) {

	logs, sub, err := _LargeStakeOracle.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LargeStakeOraclePausedIterator{contract: _LargeStakeOracle.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LargeStakeOracle *LargeStakeOracleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LargeStakeOraclePaused) (event.Subscription, error) {

	logs, sub, err := _LargeStakeOracle.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakeOraclePaused)
				if err := _LargeStakeOracle.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_LargeStakeOracle *LargeStakeOracleFilterer) ParsePaused(log types.Log) (*LargeStakeOraclePaused, error) {
	event := new(LargeStakeOraclePaused)
	if err := _LargeStakeOracle.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakeOracleReportSuccessIterator is returned from FilterReportSuccess and is used to iterate over the raw logs and unpacked data for ReportSuccess events raised by the LargeStakeOracle contract.
type LargeStakeOracleReportSuccessIterator struct {
	Event *LargeStakeOracleReportSuccess // Event containing the contract specifics and raw log

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
func (it *LargeStakeOracleReportSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakeOracleReportSuccess)
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
		it.Event = new(LargeStakeOracleReportSuccess)
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
func (it *LargeStakeOracleReportSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakeOracleReportSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakeOracleReportSuccess represents a ReportSuccess event raised by the LargeStakeOracle contract.
type LargeStakeOracleReportSuccess struct {
	RefSlot             *big.Int
	ConsensusVersion    *big.Int
	CLStakingExitInfos  []CLStakingExitInfo
	CLStakingSlashInfos []CLStakingSlashInfo
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReportSuccess is a free log retrieval operation binding the contract event 0x9ba241db180bbf54506e4868524fbbabc941126dd1d68cdf6ddaac5886441c3a.
//
// Solidity: event ReportSuccess(uint256 indexed refSlot, uint256 consensusVersion, (uint128,bytes)[] cLStakingExitInfos, (uint128,uint128,bytes)[] cLStakingSlashInfos)
func (_LargeStakeOracle *LargeStakeOracleFilterer) FilterReportSuccess(opts *bind.FilterOpts, refSlot []*big.Int) (*LargeStakeOracleReportSuccessIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _LargeStakeOracle.contract.FilterLogs(opts, "ReportSuccess", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &LargeStakeOracleReportSuccessIterator{contract: _LargeStakeOracle.contract, event: "ReportSuccess", logs: logs, sub: sub}, nil
}

// WatchReportSuccess is a free log subscription operation binding the contract event 0x9ba241db180bbf54506e4868524fbbabc941126dd1d68cdf6ddaac5886441c3a.
//
// Solidity: event ReportSuccess(uint256 indexed refSlot, uint256 consensusVersion, (uint128,bytes)[] cLStakingExitInfos, (uint128,uint128,bytes)[] cLStakingSlashInfos)
func (_LargeStakeOracle *LargeStakeOracleFilterer) WatchReportSuccess(opts *bind.WatchOpts, sink chan<- *LargeStakeOracleReportSuccess, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _LargeStakeOracle.contract.WatchLogs(opts, "ReportSuccess", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakeOracleReportSuccess)
				if err := _LargeStakeOracle.contract.UnpackLog(event, "ReportSuccess", log); err != nil {
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

// ParseReportSuccess is a log parse operation binding the contract event 0x9ba241db180bbf54506e4868524fbbabc941126dd1d68cdf6ddaac5886441c3a.
//
// Solidity: event ReportSuccess(uint256 indexed refSlot, uint256 consensusVersion, (uint128,bytes)[] cLStakingExitInfos, (uint128,uint128,bytes)[] cLStakingSlashInfos)
func (_LargeStakeOracle *LargeStakeOracleFilterer) ParseReportSuccess(log types.Log) (*LargeStakeOracleReportSuccess, error) {
	event := new(LargeStakeOracleReportSuccess)
	if err := _LargeStakeOracle.contract.UnpackLog(event, "ReportSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakeOracleUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LargeStakeOracle contract.
type LargeStakeOracleUnpausedIterator struct {
	Event *LargeStakeOracleUnpaused // Event containing the contract specifics and raw log

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
func (it *LargeStakeOracleUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakeOracleUnpaused)
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
		it.Event = new(LargeStakeOracleUnpaused)
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
func (it *LargeStakeOracleUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakeOracleUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakeOracleUnpaused represents a Unpaused event raised by the LargeStakeOracle contract.
type LargeStakeOracleUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LargeStakeOracle *LargeStakeOracleFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LargeStakeOracleUnpausedIterator, error) {

	logs, sub, err := _LargeStakeOracle.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LargeStakeOracleUnpausedIterator{contract: _LargeStakeOracle.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LargeStakeOracle *LargeStakeOracleFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LargeStakeOracleUnpaused) (event.Subscription, error) {

	logs, sub, err := _LargeStakeOracle.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakeOracleUnpaused)
				if err := _LargeStakeOracle.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_LargeStakeOracle *LargeStakeOracleFilterer) ParseUnpaused(log types.Log) (*LargeStakeOracleUnpaused, error) {
	event := new(LargeStakeOracleUnpaused)
	if err := _LargeStakeOracle.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakeOracleUpdateExitLimitIterator is returned from FilterUpdateExitLimit and is used to iterate over the raw logs and unpacked data for UpdateExitLimit events raised by the LargeStakeOracle contract.
type LargeStakeOracleUpdateExitLimitIterator struct {
	Event *LargeStakeOracleUpdateExitLimit // Event containing the contract specifics and raw log

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
func (it *LargeStakeOracleUpdateExitLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakeOracleUpdateExitLimit)
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
		it.Event = new(LargeStakeOracleUpdateExitLimit)
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
func (it *LargeStakeOracleUpdateExitLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakeOracleUpdateExitLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakeOracleUpdateExitLimit represents a UpdateExitLimit event raised by the LargeStakeOracle contract.
type LargeStakeOracleUpdateExitLimit struct {
	OldExitLimit *big.Int
	NewExitLimit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateExitLimit is a free log retrieval operation binding the contract event 0x9680447bbf45ec5a1fdc9fdbf98468d2a8b74e804efb12ad76a81f3c928a51e5.
//
// Solidity: event UpdateExitLimit(uint256 oldExitLimit, uint256 newExitLimit)
func (_LargeStakeOracle *LargeStakeOracleFilterer) FilterUpdateExitLimit(opts *bind.FilterOpts) (*LargeStakeOracleUpdateExitLimitIterator, error) {

	logs, sub, err := _LargeStakeOracle.contract.FilterLogs(opts, "UpdateExitLimit")
	if err != nil {
		return nil, err
	}
	return &LargeStakeOracleUpdateExitLimitIterator{contract: _LargeStakeOracle.contract, event: "UpdateExitLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateExitLimit is a free log subscription operation binding the contract event 0x9680447bbf45ec5a1fdc9fdbf98468d2a8b74e804efb12ad76a81f3c928a51e5.
//
// Solidity: event UpdateExitLimit(uint256 oldExitLimit, uint256 newExitLimit)
func (_LargeStakeOracle *LargeStakeOracleFilterer) WatchUpdateExitLimit(opts *bind.WatchOpts, sink chan<- *LargeStakeOracleUpdateExitLimit) (event.Subscription, error) {

	logs, sub, err := _LargeStakeOracle.contract.WatchLogs(opts, "UpdateExitLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakeOracleUpdateExitLimit)
				if err := _LargeStakeOracle.contract.UnpackLog(event, "UpdateExitLimit", log); err != nil {
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

// ParseUpdateExitLimit is a log parse operation binding the contract event 0x9680447bbf45ec5a1fdc9fdbf98468d2a8b74e804efb12ad76a81f3c928a51e5.
//
// Solidity: event UpdateExitLimit(uint256 oldExitLimit, uint256 newExitLimit)
func (_LargeStakeOracle *LargeStakeOracleFilterer) ParseUpdateExitLimit(log types.Log) (*LargeStakeOracleUpdateExitLimit, error) {
	event := new(LargeStakeOracleUpdateExitLimit)
	if err := _LargeStakeOracle.contract.UnpackLog(event, "UpdateExitLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakeOracleUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the LargeStakeOracle contract.
type LargeStakeOracleUpgradedIterator struct {
	Event *LargeStakeOracleUpgraded // Event containing the contract specifics and raw log

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
func (it *LargeStakeOracleUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakeOracleUpgraded)
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
		it.Event = new(LargeStakeOracleUpgraded)
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
func (it *LargeStakeOracleUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakeOracleUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakeOracleUpgraded represents a Upgraded event raised by the LargeStakeOracle contract.
type LargeStakeOracleUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LargeStakeOracle *LargeStakeOracleFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LargeStakeOracleUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LargeStakeOracle.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LargeStakeOracleUpgradedIterator{contract: _LargeStakeOracle.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LargeStakeOracle *LargeStakeOracleFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LargeStakeOracleUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LargeStakeOracle.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakeOracleUpgraded)
				if err := _LargeStakeOracle.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_LargeStakeOracle *LargeStakeOracleFilterer) ParseUpgraded(log types.Log) (*LargeStakeOracleUpgraded, error) {
	event := new(LargeStakeOracleUpgraded)
	if err := _LargeStakeOracle.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakeOracleWarnProcessingMissedIterator is returned from FilterWarnProcessingMissed and is used to iterate over the raw logs and unpacked data for WarnProcessingMissed events raised by the LargeStakeOracle contract.
type LargeStakeOracleWarnProcessingMissedIterator struct {
	Event *LargeStakeOracleWarnProcessingMissed // Event containing the contract specifics and raw log

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
func (it *LargeStakeOracleWarnProcessingMissedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakeOracleWarnProcessingMissed)
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
		it.Event = new(LargeStakeOracleWarnProcessingMissed)
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
func (it *LargeStakeOracleWarnProcessingMissedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakeOracleWarnProcessingMissedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakeOracleWarnProcessingMissed represents a WarnProcessingMissed event raised by the LargeStakeOracle contract.
type LargeStakeOracleWarnProcessingMissed struct {
	RefSlot *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWarnProcessingMissed is a free log retrieval operation binding the contract event 0x800b849c8bf80718cf786c99d1091c079fe2c5e420a3ba7ba9b0ef8179ef2c38.
//
// Solidity: event WarnProcessingMissed(uint256 indexed refSlot)
func (_LargeStakeOracle *LargeStakeOracleFilterer) FilterWarnProcessingMissed(opts *bind.FilterOpts, refSlot []*big.Int) (*LargeStakeOracleWarnProcessingMissedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _LargeStakeOracle.contract.FilterLogs(opts, "WarnProcessingMissed", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &LargeStakeOracleWarnProcessingMissedIterator{contract: _LargeStakeOracle.contract, event: "WarnProcessingMissed", logs: logs, sub: sub}, nil
}

// WatchWarnProcessingMissed is a free log subscription operation binding the contract event 0x800b849c8bf80718cf786c99d1091c079fe2c5e420a3ba7ba9b0ef8179ef2c38.
//
// Solidity: event WarnProcessingMissed(uint256 indexed refSlot)
func (_LargeStakeOracle *LargeStakeOracleFilterer) WatchWarnProcessingMissed(opts *bind.WatchOpts, sink chan<- *LargeStakeOracleWarnProcessingMissed, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _LargeStakeOracle.contract.WatchLogs(opts, "WarnProcessingMissed", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakeOracleWarnProcessingMissed)
				if err := _LargeStakeOracle.contract.UnpackLog(event, "WarnProcessingMissed", log); err != nil {
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
func (_LargeStakeOracle *LargeStakeOracleFilterer) ParseWarnProcessingMissed(log types.Log) (*LargeStakeOracleWarnProcessingMissed, error) {
	event := new(LargeStakeOracleWarnProcessingMissed)
	if err := _LargeStakeOracle.contract.UnpackLog(event, "WarnProcessingMissed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

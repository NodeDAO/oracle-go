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

// HashConsensusMemberConsensusState is an auto generated low-level Go binding around an user-defined struct.
type HashConsensusMemberConsensusState struct {
	CurrentFrameRefSlot         *big.Int
	CurrentFrameConsensusReport [32]byte
	IsMember                    bool
	IsFastLane                  bool
	CanReport                   bool
	LastMemberReportRefSlot     *big.Int
	CurrentFrameMemberReport    [32]byte
}

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConsensusReportAlreadyProcessing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DaoCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateMember\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateReport\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReport\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EpochsPerFrameCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FastLanePeriodCannotBeLongerThanFrame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitialEpochAlreadyArrived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitialEpochIsYetToArrive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitialEpochRefSlotCannotBeEarlierThanProcessingSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewProcessorCannotBeTheSame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonFastLaneMemberCannotReportWithinFastLaneInterval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonMember\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NumericOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedQuorum\",\"type\":\"uint256\"}],\"name\":\"QuorumTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReportProcessorCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"UnexpectedConsensusVersion\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"report\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"support\",\"type\":\"uint256\"}],\"name\":\"ConsensusReached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldDao\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"DaoAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fastLaneLengthSlots\",\"type\":\"uint256\"}],\"name\":\"FastLaneConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newInitialEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newEpochsPerFrame\",\"type\":\"uint256\"}],\"name\":\"FrameConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalMembers\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuorum\",\"type\":\"uint256\"}],\"name\":\"MemberAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalMembers\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuorum\",\"type\":\"uint256\"}],\"name\":\"MemberRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuorum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalMembers\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevQuorum\",\"type\":\"uint256\"}],\"name\":\"QuorumSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"processor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevProcessor\",\"type\":\"address\"}],\"name\":\"ReportProcessorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"report\",\"type\":\"bytes32\"}],\"name\":\"ReportReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quorum\",\"type\":\"uint256\"}],\"name\":\"addMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableConsensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"slotsPerEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondsPerSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"genesisTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"consensusReport\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isReportProcessing\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getConsensusStateForMember\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currentFrameRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"currentFrameConsensusReport\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isMember\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isFastLane\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canReport\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"lastMemberReportRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"currentFrameMemberReport\",\"type\":\"bytes32\"}],\"internalType\":\"structHashConsensus.MemberConsensusState\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentFrame\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reportProcessingDeadlineSlot\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFastLaneMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"lastReportedRefSlots\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFrameConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"initialEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochsPerFrame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fastLaneLengthSlots\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitialRefSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getIsFastLaneMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getIsMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"lastReportedRefSlots\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQuorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportProcessor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportVariants\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"variants\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"support\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotsPerEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondsPerSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"genesisTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochsPerFrame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fastLaneLengthSlots\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reportProcessor\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quorum\",\"type\":\"uint256\"}],\"name\":\"removeMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"setDaoAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fastLaneLengthSlots\",\"type\":\"uint256\"}],\"name\":\"setFastLaneLengthSlots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochsPerFrame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fastLaneLengthSlots\",\"type\":\"uint256\"}],\"name\":\"setFrameConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quorum\",\"type\":\"uint256\"}],\"name\":\"setQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newProcessor\",\"type\":\"address\"}],\"name\":\"setReportProcessor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"report\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"consensusVersion\",\"type\":\"uint256\"}],\"name\":\"submitReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialEpoch\",\"type\":\"uint256\"}],\"name\":\"updateInitialEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
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

// GetChainConfig is a free data retrieval call binding the contract method 0x606c0c94.
//
// Solidity: function getChainConfig() view returns(uint256 slotsPerEpoch, uint256 secondsPerSlot, uint256 genesisTime)
func (_Oracle *OracleCaller) GetChainConfig(opts *bind.CallOpts) (struct {
	SlotsPerEpoch  *big.Int
	SecondsPerSlot *big.Int
	GenesisTime    *big.Int
}, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getChainConfig")

	outstruct := new(struct {
		SlotsPerEpoch  *big.Int
		SecondsPerSlot *big.Int
		GenesisTime    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SlotsPerEpoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerSlot = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.GenesisTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetChainConfig is a free data retrieval call binding the contract method 0x606c0c94.
//
// Solidity: function getChainConfig() view returns(uint256 slotsPerEpoch, uint256 secondsPerSlot, uint256 genesisTime)
func (_Oracle *OracleSession) GetChainConfig() (struct {
	SlotsPerEpoch  *big.Int
	SecondsPerSlot *big.Int
	GenesisTime    *big.Int
}, error) {
	return _Oracle.Contract.GetChainConfig(&_Oracle.CallOpts)
}

// GetChainConfig is a free data retrieval call binding the contract method 0x606c0c94.
//
// Solidity: function getChainConfig() view returns(uint256 slotsPerEpoch, uint256 secondsPerSlot, uint256 genesisTime)
func (_Oracle *OracleCallerSession) GetChainConfig() (struct {
	SlotsPerEpoch  *big.Int
	SecondsPerSlot *big.Int
	GenesisTime    *big.Int
}, error) {
	return _Oracle.Contract.GetChainConfig(&_Oracle.CallOpts)
}

// GetConsensusState is a free data retrieval call binding the contract method 0xea87627d.
//
// Solidity: function getConsensusState() view returns(uint256 refSlot, bytes32 consensusReport, bool isReportProcessing)
func (_Oracle *OracleCaller) GetConsensusState(opts *bind.CallOpts) (struct {
	RefSlot            *big.Int
	ConsensusReport    [32]byte
	IsReportProcessing bool
}, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getConsensusState")

	outstruct := new(struct {
		RefSlot            *big.Int
		ConsensusReport    [32]byte
		IsReportProcessing bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RefSlot = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ConsensusReport = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.IsReportProcessing = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetConsensusState is a free data retrieval call binding the contract method 0xea87627d.
//
// Solidity: function getConsensusState() view returns(uint256 refSlot, bytes32 consensusReport, bool isReportProcessing)
func (_Oracle *OracleSession) GetConsensusState() (struct {
	RefSlot            *big.Int
	ConsensusReport    [32]byte
	IsReportProcessing bool
}, error) {
	return _Oracle.Contract.GetConsensusState(&_Oracle.CallOpts)
}

// GetConsensusState is a free data retrieval call binding the contract method 0xea87627d.
//
// Solidity: function getConsensusState() view returns(uint256 refSlot, bytes32 consensusReport, bool isReportProcessing)
func (_Oracle *OracleCallerSession) GetConsensusState() (struct {
	RefSlot            *big.Int
	ConsensusReport    [32]byte
	IsReportProcessing bool
}, error) {
	return _Oracle.Contract.GetConsensusState(&_Oracle.CallOpts)
}

// GetConsensusStateForMember is a free data retrieval call binding the contract method 0x60e61801.
//
// Solidity: function getConsensusStateForMember(address addr) view returns((uint256,bytes32,bool,bool,bool,uint256,bytes32) result)
func (_Oracle *OracleCaller) GetConsensusStateForMember(opts *bind.CallOpts, addr common.Address) (HashConsensusMemberConsensusState, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getConsensusStateForMember", addr)

	if err != nil {
		return *new(HashConsensusMemberConsensusState), err
	}

	out0 := *abi.ConvertType(out[0], new(HashConsensusMemberConsensusState)).(*HashConsensusMemberConsensusState)

	return out0, err

}

// GetConsensusStateForMember is a free data retrieval call binding the contract method 0x60e61801.
//
// Solidity: function getConsensusStateForMember(address addr) view returns((uint256,bytes32,bool,bool,bool,uint256,bytes32) result)
func (_Oracle *OracleSession) GetConsensusStateForMember(addr common.Address) (HashConsensusMemberConsensusState, error) {
	return _Oracle.Contract.GetConsensusStateForMember(&_Oracle.CallOpts, addr)
}

// GetConsensusStateForMember is a free data retrieval call binding the contract method 0x60e61801.
//
// Solidity: function getConsensusStateForMember(address addr) view returns((uint256,bytes32,bool,bool,bool,uint256,bytes32) result)
func (_Oracle *OracleCallerSession) GetConsensusStateForMember(addr common.Address) (HashConsensusMemberConsensusState, error) {
	return _Oracle.Contract.GetConsensusStateForMember(&_Oracle.CallOpts, addr)
}

// GetCurrentFrame is a free data retrieval call binding the contract method 0x72f79b13.
//
// Solidity: function getCurrentFrame() view returns(uint256 refSlot, uint256 reportProcessingDeadlineSlot)
func (_Oracle *OracleCaller) GetCurrentFrame(opts *bind.CallOpts) (struct {
	RefSlot                      *big.Int
	ReportProcessingDeadlineSlot *big.Int
}, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getCurrentFrame")

	outstruct := new(struct {
		RefSlot                      *big.Int
		ReportProcessingDeadlineSlot *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RefSlot = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReportProcessingDeadlineSlot = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCurrentFrame is a free data retrieval call binding the contract method 0x72f79b13.
//
// Solidity: function getCurrentFrame() view returns(uint256 refSlot, uint256 reportProcessingDeadlineSlot)
func (_Oracle *OracleSession) GetCurrentFrame() (struct {
	RefSlot                      *big.Int
	ReportProcessingDeadlineSlot *big.Int
}, error) {
	return _Oracle.Contract.GetCurrentFrame(&_Oracle.CallOpts)
}

// GetCurrentFrame is a free data retrieval call binding the contract method 0x72f79b13.
//
// Solidity: function getCurrentFrame() view returns(uint256 refSlot, uint256 reportProcessingDeadlineSlot)
func (_Oracle *OracleCallerSession) GetCurrentFrame() (struct {
	RefSlot                      *big.Int
	ReportProcessingDeadlineSlot *big.Int
}, error) {
	return _Oracle.Contract.GetCurrentFrame(&_Oracle.CallOpts)
}

// GetFastLaneMembers is a free data retrieval call binding the contract method 0x433ab1f3.
//
// Solidity: function getFastLaneMembers() view returns(address[] addresses, uint256[] lastReportedRefSlots)
func (_Oracle *OracleCaller) GetFastLaneMembers(opts *bind.CallOpts) (struct {
	Addresses            []common.Address
	LastReportedRefSlots []*big.Int
}, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getFastLaneMembers")

	outstruct := new(struct {
		Addresses            []common.Address
		LastReportedRefSlots []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addresses = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.LastReportedRefSlots = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetFastLaneMembers is a free data retrieval call binding the contract method 0x433ab1f3.
//
// Solidity: function getFastLaneMembers() view returns(address[] addresses, uint256[] lastReportedRefSlots)
func (_Oracle *OracleSession) GetFastLaneMembers() (struct {
	Addresses            []common.Address
	LastReportedRefSlots []*big.Int
}, error) {
	return _Oracle.Contract.GetFastLaneMembers(&_Oracle.CallOpts)
}

// GetFastLaneMembers is a free data retrieval call binding the contract method 0x433ab1f3.
//
// Solidity: function getFastLaneMembers() view returns(address[] addresses, uint256[] lastReportedRefSlots)
func (_Oracle *OracleCallerSession) GetFastLaneMembers() (struct {
	Addresses            []common.Address
	LastReportedRefSlots []*big.Int
}, error) {
	return _Oracle.Contract.GetFastLaneMembers(&_Oracle.CallOpts)
}

// GetFrameConfig is a free data retrieval call binding the contract method 0x6fb1bf66.
//
// Solidity: function getFrameConfig() view returns(uint256 initialEpoch, uint256 epochsPerFrame, uint256 fastLaneLengthSlots)
func (_Oracle *OracleCaller) GetFrameConfig(opts *bind.CallOpts) (struct {
	InitialEpoch        *big.Int
	EpochsPerFrame      *big.Int
	FastLaneLengthSlots *big.Int
}, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getFrameConfig")

	outstruct := new(struct {
		InitialEpoch        *big.Int
		EpochsPerFrame      *big.Int
		FastLaneLengthSlots *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InitialEpoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EpochsPerFrame = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FastLaneLengthSlots = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetFrameConfig is a free data retrieval call binding the contract method 0x6fb1bf66.
//
// Solidity: function getFrameConfig() view returns(uint256 initialEpoch, uint256 epochsPerFrame, uint256 fastLaneLengthSlots)
func (_Oracle *OracleSession) GetFrameConfig() (struct {
	InitialEpoch        *big.Int
	EpochsPerFrame      *big.Int
	FastLaneLengthSlots *big.Int
}, error) {
	return _Oracle.Contract.GetFrameConfig(&_Oracle.CallOpts)
}

// GetFrameConfig is a free data retrieval call binding the contract method 0x6fb1bf66.
//
// Solidity: function getFrameConfig() view returns(uint256 initialEpoch, uint256 epochsPerFrame, uint256 fastLaneLengthSlots)
func (_Oracle *OracleCallerSession) GetFrameConfig() (struct {
	InitialEpoch        *big.Int
	EpochsPerFrame      *big.Int
	FastLaneLengthSlots *big.Int
}, error) {
	return _Oracle.Contract.GetFrameConfig(&_Oracle.CallOpts)
}

// GetInitialRefSlot is a free data retrieval call binding the contract method 0x6095012f.
//
// Solidity: function getInitialRefSlot() view returns(uint256)
func (_Oracle *OracleCaller) GetInitialRefSlot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getInitialRefSlot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitialRefSlot is a free data retrieval call binding the contract method 0x6095012f.
//
// Solidity: function getInitialRefSlot() view returns(uint256)
func (_Oracle *OracleSession) GetInitialRefSlot() (*big.Int, error) {
	return _Oracle.Contract.GetInitialRefSlot(&_Oracle.CallOpts)
}

// GetInitialRefSlot is a free data retrieval call binding the contract method 0x6095012f.
//
// Solidity: function getInitialRefSlot() view returns(uint256)
func (_Oracle *OracleCallerSession) GetInitialRefSlot() (*big.Int, error) {
	return _Oracle.Contract.GetInitialRefSlot(&_Oracle.CallOpts)
}

// GetIsFastLaneMember is a free data retrieval call binding the contract method 0x20b4d751.
//
// Solidity: function getIsFastLaneMember(address addr) view returns(bool)
func (_Oracle *OracleCaller) GetIsFastLaneMember(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getIsFastLaneMember", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsFastLaneMember is a free data retrieval call binding the contract method 0x20b4d751.
//
// Solidity: function getIsFastLaneMember(address addr) view returns(bool)
func (_Oracle *OracleSession) GetIsFastLaneMember(addr common.Address) (bool, error) {
	return _Oracle.Contract.GetIsFastLaneMember(&_Oracle.CallOpts, addr)
}

// GetIsFastLaneMember is a free data retrieval call binding the contract method 0x20b4d751.
//
// Solidity: function getIsFastLaneMember(address addr) view returns(bool)
func (_Oracle *OracleCallerSession) GetIsFastLaneMember(addr common.Address) (bool, error) {
	return _Oracle.Contract.GetIsFastLaneMember(&_Oracle.CallOpts, addr)
}

// GetIsMember is a free data retrieval call binding the contract method 0x1951c037.
//
// Solidity: function getIsMember(address addr) view returns(bool)
func (_Oracle *OracleCaller) GetIsMember(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getIsMember", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsMember is a free data retrieval call binding the contract method 0x1951c037.
//
// Solidity: function getIsMember(address addr) view returns(bool)
func (_Oracle *OracleSession) GetIsMember(addr common.Address) (bool, error) {
	return _Oracle.Contract.GetIsMember(&_Oracle.CallOpts, addr)
}

// GetIsMember is a free data retrieval call binding the contract method 0x1951c037.
//
// Solidity: function getIsMember(address addr) view returns(bool)
func (_Oracle *OracleCallerSession) GetIsMember(addr common.Address) (bool, error) {
	return _Oracle.Contract.GetIsMember(&_Oracle.CallOpts, addr)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[] addresses, uint256[] lastReportedRefSlots)
func (_Oracle *OracleCaller) GetMembers(opts *bind.CallOpts) (struct {
	Addresses            []common.Address
	LastReportedRefSlots []*big.Int
}, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getMembers")

	outstruct := new(struct {
		Addresses            []common.Address
		LastReportedRefSlots []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addresses = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.LastReportedRefSlots = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[] addresses, uint256[] lastReportedRefSlots)
func (_Oracle *OracleSession) GetMembers() (struct {
	Addresses            []common.Address
	LastReportedRefSlots []*big.Int
}, error) {
	return _Oracle.Contract.GetMembers(&_Oracle.CallOpts)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[] addresses, uint256[] lastReportedRefSlots)
func (_Oracle *OracleCallerSession) GetMembers() (struct {
	Addresses            []common.Address
	LastReportedRefSlots []*big.Int
}, error) {
	return _Oracle.Contract.GetMembers(&_Oracle.CallOpts)
}

// GetQuorum is a free data retrieval call binding the contract method 0xc26c12eb.
//
// Solidity: function getQuorum() view returns(uint256)
func (_Oracle *OracleCaller) GetQuorum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getQuorum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorum is a free data retrieval call binding the contract method 0xc26c12eb.
//
// Solidity: function getQuorum() view returns(uint256)
func (_Oracle *OracleSession) GetQuorum() (*big.Int, error) {
	return _Oracle.Contract.GetQuorum(&_Oracle.CallOpts)
}

// GetQuorum is a free data retrieval call binding the contract method 0xc26c12eb.
//
// Solidity: function getQuorum() view returns(uint256)
func (_Oracle *OracleCallerSession) GetQuorum() (*big.Int, error) {
	return _Oracle.Contract.GetQuorum(&_Oracle.CallOpts)
}

// GetReportProcessor is a free data retrieval call binding the contract method 0x6d058268.
//
// Solidity: function getReportProcessor() view returns(address)
func (_Oracle *OracleCaller) GetReportProcessor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getReportProcessor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReportProcessor is a free data retrieval call binding the contract method 0x6d058268.
//
// Solidity: function getReportProcessor() view returns(address)
func (_Oracle *OracleSession) GetReportProcessor() (common.Address, error) {
	return _Oracle.Contract.GetReportProcessor(&_Oracle.CallOpts)
}

// GetReportProcessor is a free data retrieval call binding the contract method 0x6d058268.
//
// Solidity: function getReportProcessor() view returns(address)
func (_Oracle *OracleCallerSession) GetReportProcessor() (common.Address, error) {
	return _Oracle.Contract.GetReportProcessor(&_Oracle.CallOpts)
}

// GetReportVariants is a free data retrieval call binding the contract method 0x2fd2d750.
//
// Solidity: function getReportVariants() view returns(bytes32[] variants, uint256[] support)
func (_Oracle *OracleCaller) GetReportVariants(opts *bind.CallOpts) (struct {
	Variants [][32]byte
	Support  []*big.Int
}, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getReportVariants")

	outstruct := new(struct {
		Variants [][32]byte
		Support  []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Variants = *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	outstruct.Support = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetReportVariants is a free data retrieval call binding the contract method 0x2fd2d750.
//
// Solidity: function getReportVariants() view returns(bytes32[] variants, uint256[] support)
func (_Oracle *OracleSession) GetReportVariants() (struct {
	Variants [][32]byte
	Support  []*big.Int
}, error) {
	return _Oracle.Contract.GetReportVariants(&_Oracle.CallOpts)
}

// GetReportVariants is a free data retrieval call binding the contract method 0x2fd2d750.
//
// Solidity: function getReportVariants() view returns(bytes32[] variants, uint256[] support)
func (_Oracle *OracleCallerSession) GetReportVariants() (struct {
	Variants [][32]byte
	Support  []*big.Int
}, error) {
	return _Oracle.Contract.GetReportVariants(&_Oracle.CallOpts)
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

// AddMember is a paid mutator transaction binding the contract method 0x98041ea3.
//
// Solidity: function addMember(address addr, uint256 quorum) returns()
func (_Oracle *OracleTransactor) AddMember(opts *bind.TransactOpts, addr common.Address, quorum *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "addMember", addr, quorum)
}

// AddMember is a paid mutator transaction binding the contract method 0x98041ea3.
//
// Solidity: function addMember(address addr, uint256 quorum) returns()
func (_Oracle *OracleSession) AddMember(addr common.Address, quorum *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.AddMember(&_Oracle.TransactOpts, addr, quorum)
}

// AddMember is a paid mutator transaction binding the contract method 0x98041ea3.
//
// Solidity: function addMember(address addr, uint256 quorum) returns()
func (_Oracle *OracleTransactorSession) AddMember(addr common.Address, quorum *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.AddMember(&_Oracle.TransactOpts, addr, quorum)
}

// DisableConsensus is a paid mutator transaction binding the contract method 0xad231cb2.
//
// Solidity: function disableConsensus() returns()
func (_Oracle *OracleTransactor) DisableConsensus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "disableConsensus")
}

// DisableConsensus is a paid mutator transaction binding the contract method 0xad231cb2.
//
// Solidity: function disableConsensus() returns()
func (_Oracle *OracleSession) DisableConsensus() (*types.Transaction, error) {
	return _Oracle.Contract.DisableConsensus(&_Oracle.TransactOpts)
}

// DisableConsensus is a paid mutator transaction binding the contract method 0xad231cb2.
//
// Solidity: function disableConsensus() returns()
func (_Oracle *OracleTransactorSession) DisableConsensus() (*types.Transaction, error) {
	return _Oracle.Contract.DisableConsensus(&_Oracle.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x1f85d8bc.
//
// Solidity: function initialize(uint256 slotsPerEpoch, uint256 secondsPerSlot, uint256 genesisTime, uint256 epochsPerFrame, uint256 fastLaneLengthSlots, address _dao, address reportProcessor) returns()
func (_Oracle *OracleTransactor) Initialize(opts *bind.TransactOpts, slotsPerEpoch *big.Int, secondsPerSlot *big.Int, genesisTime *big.Int, epochsPerFrame *big.Int, fastLaneLengthSlots *big.Int, _dao common.Address, reportProcessor common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "initialize", slotsPerEpoch, secondsPerSlot, genesisTime, epochsPerFrame, fastLaneLengthSlots, _dao, reportProcessor)
}

// Initialize is a paid mutator transaction binding the contract method 0x1f85d8bc.
//
// Solidity: function initialize(uint256 slotsPerEpoch, uint256 secondsPerSlot, uint256 genesisTime, uint256 epochsPerFrame, uint256 fastLaneLengthSlots, address _dao, address reportProcessor) returns()
func (_Oracle *OracleSession) Initialize(slotsPerEpoch *big.Int, secondsPerSlot *big.Int, genesisTime *big.Int, epochsPerFrame *big.Int, fastLaneLengthSlots *big.Int, _dao common.Address, reportProcessor common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.Initialize(&_Oracle.TransactOpts, slotsPerEpoch, secondsPerSlot, genesisTime, epochsPerFrame, fastLaneLengthSlots, _dao, reportProcessor)
}

// Initialize is a paid mutator transaction binding the contract method 0x1f85d8bc.
//
// Solidity: function initialize(uint256 slotsPerEpoch, uint256 secondsPerSlot, uint256 genesisTime, uint256 epochsPerFrame, uint256 fastLaneLengthSlots, address _dao, address reportProcessor) returns()
func (_Oracle *OracleTransactorSession) Initialize(slotsPerEpoch *big.Int, secondsPerSlot *big.Int, genesisTime *big.Int, epochsPerFrame *big.Int, fastLaneLengthSlots *big.Int, _dao common.Address, reportProcessor common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.Initialize(&_Oracle.TransactOpts, slotsPerEpoch, secondsPerSlot, genesisTime, epochsPerFrame, fastLaneLengthSlots, _dao, reportProcessor)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x16f6f03e.
//
// Solidity: function removeMember(address addr, uint256 quorum) returns()
func (_Oracle *OracleTransactor) RemoveMember(opts *bind.TransactOpts, addr common.Address, quorum *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "removeMember", addr, quorum)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x16f6f03e.
//
// Solidity: function removeMember(address addr, uint256 quorum) returns()
func (_Oracle *OracleSession) RemoveMember(addr common.Address, quorum *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.RemoveMember(&_Oracle.TransactOpts, addr, quorum)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x16f6f03e.
//
// Solidity: function removeMember(address addr, uint256 quorum) returns()
func (_Oracle *OracleTransactorSession) RemoveMember(addr common.Address, quorum *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.RemoveMember(&_Oracle.TransactOpts, addr, quorum)
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

// SetFastLaneLengthSlots is a paid mutator transaction binding the contract method 0x99229f58.
//
// Solidity: function setFastLaneLengthSlots(uint256 fastLaneLengthSlots) returns()
func (_Oracle *OracleTransactor) SetFastLaneLengthSlots(opts *bind.TransactOpts, fastLaneLengthSlots *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setFastLaneLengthSlots", fastLaneLengthSlots)
}

// SetFastLaneLengthSlots is a paid mutator transaction binding the contract method 0x99229f58.
//
// Solidity: function setFastLaneLengthSlots(uint256 fastLaneLengthSlots) returns()
func (_Oracle *OracleSession) SetFastLaneLengthSlots(fastLaneLengthSlots *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetFastLaneLengthSlots(&_Oracle.TransactOpts, fastLaneLengthSlots)
}

// SetFastLaneLengthSlots is a paid mutator transaction binding the contract method 0x99229f58.
//
// Solidity: function setFastLaneLengthSlots(uint256 fastLaneLengthSlots) returns()
func (_Oracle *OracleTransactorSession) SetFastLaneLengthSlots(fastLaneLengthSlots *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetFastLaneLengthSlots(&_Oracle.TransactOpts, fastLaneLengthSlots)
}

// SetFrameConfig is a paid mutator transaction binding the contract method 0x34aa6753.
//
// Solidity: function setFrameConfig(uint256 epochsPerFrame, uint256 fastLaneLengthSlots) returns()
func (_Oracle *OracleTransactor) SetFrameConfig(opts *bind.TransactOpts, epochsPerFrame *big.Int, fastLaneLengthSlots *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setFrameConfig", epochsPerFrame, fastLaneLengthSlots)
}

// SetFrameConfig is a paid mutator transaction binding the contract method 0x34aa6753.
//
// Solidity: function setFrameConfig(uint256 epochsPerFrame, uint256 fastLaneLengthSlots) returns()
func (_Oracle *OracleSession) SetFrameConfig(epochsPerFrame *big.Int, fastLaneLengthSlots *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetFrameConfig(&_Oracle.TransactOpts, epochsPerFrame, fastLaneLengthSlots)
}

// SetFrameConfig is a paid mutator transaction binding the contract method 0x34aa6753.
//
// Solidity: function setFrameConfig(uint256 epochsPerFrame, uint256 fastLaneLengthSlots) returns()
func (_Oracle *OracleTransactorSession) SetFrameConfig(epochsPerFrame *big.Int, fastLaneLengthSlots *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetFrameConfig(&_Oracle.TransactOpts, epochsPerFrame, fastLaneLengthSlots)
}

// SetQuorum is a paid mutator transaction binding the contract method 0xc1ba4e59.
//
// Solidity: function setQuorum(uint256 quorum) returns()
func (_Oracle *OracleTransactor) SetQuorum(opts *bind.TransactOpts, quorum *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setQuorum", quorum)
}

// SetQuorum is a paid mutator transaction binding the contract method 0xc1ba4e59.
//
// Solidity: function setQuorum(uint256 quorum) returns()
func (_Oracle *OracleSession) SetQuorum(quorum *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetQuorum(&_Oracle.TransactOpts, quorum)
}

// SetQuorum is a paid mutator transaction binding the contract method 0xc1ba4e59.
//
// Solidity: function setQuorum(uint256 quorum) returns()
func (_Oracle *OracleTransactorSession) SetQuorum(quorum *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetQuorum(&_Oracle.TransactOpts, quorum)
}

// SetReportProcessor is a paid mutator transaction binding the contract method 0xe76cd4e0.
//
// Solidity: function setReportProcessor(address newProcessor) returns()
func (_Oracle *OracleTransactor) SetReportProcessor(opts *bind.TransactOpts, newProcessor common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setReportProcessor", newProcessor)
}

// SetReportProcessor is a paid mutator transaction binding the contract method 0xe76cd4e0.
//
// Solidity: function setReportProcessor(address newProcessor) returns()
func (_Oracle *OracleSession) SetReportProcessor(newProcessor common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetReportProcessor(&_Oracle.TransactOpts, newProcessor)
}

// SetReportProcessor is a paid mutator transaction binding the contract method 0xe76cd4e0.
//
// Solidity: function setReportProcessor(address newProcessor) returns()
func (_Oracle *OracleTransactorSession) SetReportProcessor(newProcessor common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetReportProcessor(&_Oracle.TransactOpts, newProcessor)
}

// SubmitReport is a paid mutator transaction binding the contract method 0xe33a8d39.
//
// Solidity: function submitReport(uint256 slot, bytes32 report, uint256 consensusVersion) returns()
func (_Oracle *OracleTransactor) SubmitReport(opts *bind.TransactOpts, slot *big.Int, report [32]byte, consensusVersion *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "submitReport", slot, report, consensusVersion)
}

// SubmitReport is a paid mutator transaction binding the contract method 0xe33a8d39.
//
// Solidity: function submitReport(uint256 slot, bytes32 report, uint256 consensusVersion) returns()
func (_Oracle *OracleSession) SubmitReport(slot *big.Int, report [32]byte, consensusVersion *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SubmitReport(&_Oracle.TransactOpts, slot, report, consensusVersion)
}

// SubmitReport is a paid mutator transaction binding the contract method 0xe33a8d39.
//
// Solidity: function submitReport(uint256 slot, bytes32 report, uint256 consensusVersion) returns()
func (_Oracle *OracleTransactorSession) SubmitReport(slot *big.Int, report [32]byte, consensusVersion *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SubmitReport(&_Oracle.TransactOpts, slot, report, consensusVersion)
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

// UpdateInitialEpoch is a paid mutator transaction binding the contract method 0x323a41f6.
//
// Solidity: function updateInitialEpoch(uint256 initialEpoch) returns()
func (_Oracle *OracleTransactor) UpdateInitialEpoch(opts *bind.TransactOpts, initialEpoch *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "updateInitialEpoch", initialEpoch)
}

// UpdateInitialEpoch is a paid mutator transaction binding the contract method 0x323a41f6.
//
// Solidity: function updateInitialEpoch(uint256 initialEpoch) returns()
func (_Oracle *OracleSession) UpdateInitialEpoch(initialEpoch *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.UpdateInitialEpoch(&_Oracle.TransactOpts, initialEpoch)
}

// UpdateInitialEpoch is a paid mutator transaction binding the contract method 0x323a41f6.
//
// Solidity: function updateInitialEpoch(uint256 initialEpoch) returns()
func (_Oracle *OracleTransactorSession) UpdateInitialEpoch(initialEpoch *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.UpdateInitialEpoch(&_Oracle.TransactOpts, initialEpoch)
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

// OracleConsensusReachedIterator is returned from FilterConsensusReached and is used to iterate over the raw logs and unpacked data for ConsensusReached events raised by the Oracle contract.
type OracleConsensusReachedIterator struct {
	Event *OracleConsensusReached // Event containing the contract specifics and raw log

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
func (it *OracleConsensusReachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleConsensusReached)
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
		it.Event = new(OracleConsensusReached)
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
func (it *OracleConsensusReachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleConsensusReachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleConsensusReached represents a ConsensusReached event raised by the Oracle contract.
type OracleConsensusReached struct {
	RefSlot *big.Int
	Report  [32]byte
	Support *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConsensusReached is a free log retrieval operation binding the contract event 0x2b6bc782c916fa763822f1e50c6db0f95dade36d6541a8a4cbe070735b8b226d.
//
// Solidity: event ConsensusReached(uint256 indexed refSlot, bytes32 report, uint256 support)
func (_Oracle *OracleFilterer) FilterConsensusReached(opts *bind.FilterOpts, refSlot []*big.Int) (*OracleConsensusReachedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "ConsensusReached", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &OracleConsensusReachedIterator{contract: _Oracle.contract, event: "ConsensusReached", logs: logs, sub: sub}, nil
}

// WatchConsensusReached is a free log subscription operation binding the contract event 0x2b6bc782c916fa763822f1e50c6db0f95dade36d6541a8a4cbe070735b8b226d.
//
// Solidity: event ConsensusReached(uint256 indexed refSlot, bytes32 report, uint256 support)
func (_Oracle *OracleFilterer) WatchConsensusReached(opts *bind.WatchOpts, sink chan<- *OracleConsensusReached, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "ConsensusReached", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleConsensusReached)
				if err := _Oracle.contract.UnpackLog(event, "ConsensusReached", log); err != nil {
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

// ParseConsensusReached is a log parse operation binding the contract event 0x2b6bc782c916fa763822f1e50c6db0f95dade36d6541a8a4cbe070735b8b226d.
//
// Solidity: event ConsensusReached(uint256 indexed refSlot, bytes32 report, uint256 support)
func (_Oracle *OracleFilterer) ParseConsensusReached(log types.Log) (*OracleConsensusReached, error) {
	event := new(OracleConsensusReached)
	if err := _Oracle.contract.UnpackLog(event, "ConsensusReached", log); err != nil {
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

// OracleFastLaneConfigSetIterator is returned from FilterFastLaneConfigSet and is used to iterate over the raw logs and unpacked data for FastLaneConfigSet events raised by the Oracle contract.
type OracleFastLaneConfigSetIterator struct {
	Event *OracleFastLaneConfigSet // Event containing the contract specifics and raw log

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
func (it *OracleFastLaneConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleFastLaneConfigSet)
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
		it.Event = new(OracleFastLaneConfigSet)
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
func (it *OracleFastLaneConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleFastLaneConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleFastLaneConfigSet represents a FastLaneConfigSet event raised by the Oracle contract.
type OracleFastLaneConfigSet struct {
	FastLaneLengthSlots *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterFastLaneConfigSet is a free log retrieval operation binding the contract event 0xab8b22776606cc75c47792d32af7e63ed9ca74e85c9780a7fc7994fdbd6fde2b.
//
// Solidity: event FastLaneConfigSet(uint256 fastLaneLengthSlots)
func (_Oracle *OracleFilterer) FilterFastLaneConfigSet(opts *bind.FilterOpts) (*OracleFastLaneConfigSetIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "FastLaneConfigSet")
	if err != nil {
		return nil, err
	}
	return &OracleFastLaneConfigSetIterator{contract: _Oracle.contract, event: "FastLaneConfigSet", logs: logs, sub: sub}, nil
}

// WatchFastLaneConfigSet is a free log subscription operation binding the contract event 0xab8b22776606cc75c47792d32af7e63ed9ca74e85c9780a7fc7994fdbd6fde2b.
//
// Solidity: event FastLaneConfigSet(uint256 fastLaneLengthSlots)
func (_Oracle *OracleFilterer) WatchFastLaneConfigSet(opts *bind.WatchOpts, sink chan<- *OracleFastLaneConfigSet) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "FastLaneConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleFastLaneConfigSet)
				if err := _Oracle.contract.UnpackLog(event, "FastLaneConfigSet", log); err != nil {
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

// ParseFastLaneConfigSet is a log parse operation binding the contract event 0xab8b22776606cc75c47792d32af7e63ed9ca74e85c9780a7fc7994fdbd6fde2b.
//
// Solidity: event FastLaneConfigSet(uint256 fastLaneLengthSlots)
func (_Oracle *OracleFilterer) ParseFastLaneConfigSet(log types.Log) (*OracleFastLaneConfigSet, error) {
	event := new(OracleFastLaneConfigSet)
	if err := _Oracle.contract.UnpackLog(event, "FastLaneConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleFrameConfigSetIterator is returned from FilterFrameConfigSet and is used to iterate over the raw logs and unpacked data for FrameConfigSet events raised by the Oracle contract.
type OracleFrameConfigSetIterator struct {
	Event *OracleFrameConfigSet // Event containing the contract specifics and raw log

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
func (it *OracleFrameConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleFrameConfigSet)
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
		it.Event = new(OracleFrameConfigSet)
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
func (it *OracleFrameConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleFrameConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleFrameConfigSet represents a FrameConfigSet event raised by the Oracle contract.
type OracleFrameConfigSet struct {
	NewInitialEpoch   *big.Int
	NewEpochsPerFrame *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFrameConfigSet is a free log retrieval operation binding the contract event 0xe343afa5219eaf28c50ce9cd658acd69cbe28b34fa773eb3a523e28007f64afc.
//
// Solidity: event FrameConfigSet(uint256 newInitialEpoch, uint256 newEpochsPerFrame)
func (_Oracle *OracleFilterer) FilterFrameConfigSet(opts *bind.FilterOpts) (*OracleFrameConfigSetIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "FrameConfigSet")
	if err != nil {
		return nil, err
	}
	return &OracleFrameConfigSetIterator{contract: _Oracle.contract, event: "FrameConfigSet", logs: logs, sub: sub}, nil
}

// WatchFrameConfigSet is a free log subscription operation binding the contract event 0xe343afa5219eaf28c50ce9cd658acd69cbe28b34fa773eb3a523e28007f64afc.
//
// Solidity: event FrameConfigSet(uint256 newInitialEpoch, uint256 newEpochsPerFrame)
func (_Oracle *OracleFilterer) WatchFrameConfigSet(opts *bind.WatchOpts, sink chan<- *OracleFrameConfigSet) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "FrameConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleFrameConfigSet)
				if err := _Oracle.contract.UnpackLog(event, "FrameConfigSet", log); err != nil {
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

// ParseFrameConfigSet is a log parse operation binding the contract event 0xe343afa5219eaf28c50ce9cd658acd69cbe28b34fa773eb3a523e28007f64afc.
//
// Solidity: event FrameConfigSet(uint256 newInitialEpoch, uint256 newEpochsPerFrame)
func (_Oracle *OracleFilterer) ParseFrameConfigSet(log types.Log) (*OracleFrameConfigSet, error) {
	event := new(OracleFrameConfigSet)
	if err := _Oracle.contract.UnpackLog(event, "FrameConfigSet", log); err != nil {
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

// OracleMemberAddedIterator is returned from FilterMemberAdded and is used to iterate over the raw logs and unpacked data for MemberAdded events raised by the Oracle contract.
type OracleMemberAddedIterator struct {
	Event *OracleMemberAdded // Event containing the contract specifics and raw log

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
func (it *OracleMemberAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleMemberAdded)
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
		it.Event = new(OracleMemberAdded)
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
func (it *OracleMemberAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleMemberAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleMemberAdded represents a MemberAdded event raised by the Oracle contract.
type OracleMemberAdded struct {
	Addr            common.Address
	NewTotalMembers *big.Int
	NewQuorum       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMemberAdded is a free log retrieval operation binding the contract event 0xe17e0e2cd88e2144dd54f3d823c30d4569092bcac1aabaec1129883e9cc12d2e.
//
// Solidity: event MemberAdded(address indexed addr, uint256 newTotalMembers, uint256 newQuorum)
func (_Oracle *OracleFilterer) FilterMemberAdded(opts *bind.FilterOpts, addr []common.Address) (*OracleMemberAddedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "MemberAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return &OracleMemberAddedIterator{contract: _Oracle.contract, event: "MemberAdded", logs: logs, sub: sub}, nil
}

// WatchMemberAdded is a free log subscription operation binding the contract event 0xe17e0e2cd88e2144dd54f3d823c30d4569092bcac1aabaec1129883e9cc12d2e.
//
// Solidity: event MemberAdded(address indexed addr, uint256 newTotalMembers, uint256 newQuorum)
func (_Oracle *OracleFilterer) WatchMemberAdded(opts *bind.WatchOpts, sink chan<- *OracleMemberAdded, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "MemberAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleMemberAdded)
				if err := _Oracle.contract.UnpackLog(event, "MemberAdded", log); err != nil {
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

// ParseMemberAdded is a log parse operation binding the contract event 0xe17e0e2cd88e2144dd54f3d823c30d4569092bcac1aabaec1129883e9cc12d2e.
//
// Solidity: event MemberAdded(address indexed addr, uint256 newTotalMembers, uint256 newQuorum)
func (_Oracle *OracleFilterer) ParseMemberAdded(log types.Log) (*OracleMemberAdded, error) {
	event := new(OracleMemberAdded)
	if err := _Oracle.contract.UnpackLog(event, "MemberAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleMemberRemovedIterator is returned from FilterMemberRemoved and is used to iterate over the raw logs and unpacked data for MemberRemoved events raised by the Oracle contract.
type OracleMemberRemovedIterator struct {
	Event *OracleMemberRemoved // Event containing the contract specifics and raw log

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
func (it *OracleMemberRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleMemberRemoved)
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
		it.Event = new(OracleMemberRemoved)
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
func (it *OracleMemberRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleMemberRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleMemberRemoved represents a MemberRemoved event raised by the Oracle contract.
type OracleMemberRemoved struct {
	Addr            common.Address
	NewTotalMembers *big.Int
	NewQuorum       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMemberRemoved is a free log retrieval operation binding the contract event 0xa182730913550d27dc6c5813fad297cb0785871bec3d0152c5650e59c5d39d60.
//
// Solidity: event MemberRemoved(address indexed addr, uint256 newTotalMembers, uint256 newQuorum)
func (_Oracle *OracleFilterer) FilterMemberRemoved(opts *bind.FilterOpts, addr []common.Address) (*OracleMemberRemovedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "MemberRemoved", addrRule)
	if err != nil {
		return nil, err
	}
	return &OracleMemberRemovedIterator{contract: _Oracle.contract, event: "MemberRemoved", logs: logs, sub: sub}, nil
}

// WatchMemberRemoved is a free log subscription operation binding the contract event 0xa182730913550d27dc6c5813fad297cb0785871bec3d0152c5650e59c5d39d60.
//
// Solidity: event MemberRemoved(address indexed addr, uint256 newTotalMembers, uint256 newQuorum)
func (_Oracle *OracleFilterer) WatchMemberRemoved(opts *bind.WatchOpts, sink chan<- *OracleMemberRemoved, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "MemberRemoved", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleMemberRemoved)
				if err := _Oracle.contract.UnpackLog(event, "MemberRemoved", log); err != nil {
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

// ParseMemberRemoved is a log parse operation binding the contract event 0xa182730913550d27dc6c5813fad297cb0785871bec3d0152c5650e59c5d39d60.
//
// Solidity: event MemberRemoved(address indexed addr, uint256 newTotalMembers, uint256 newQuorum)
func (_Oracle *OracleFilterer) ParseMemberRemoved(log types.Log) (*OracleMemberRemoved, error) {
	event := new(OracleMemberRemoved)
	if err := _Oracle.contract.UnpackLog(event, "MemberRemoved", log); err != nil {
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

// OracleQuorumSetIterator is returned from FilterQuorumSet and is used to iterate over the raw logs and unpacked data for QuorumSet events raised by the Oracle contract.
type OracleQuorumSetIterator struct {
	Event *OracleQuorumSet // Event containing the contract specifics and raw log

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
func (it *OracleQuorumSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleQuorumSet)
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
		it.Event = new(OracleQuorumSet)
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
func (it *OracleQuorumSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleQuorumSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleQuorumSet represents a QuorumSet event raised by the Oracle contract.
type OracleQuorumSet struct {
	NewQuorum    *big.Int
	TotalMembers *big.Int
	PrevQuorum   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuorumSet is a free log retrieval operation binding the contract event 0x9f40cfd22fe91777c78f252bd21a710f3fb007dc2f321876891e7644ba0ae175.
//
// Solidity: event QuorumSet(uint256 newQuorum, uint256 totalMembers, uint256 prevQuorum)
func (_Oracle *OracleFilterer) FilterQuorumSet(opts *bind.FilterOpts) (*OracleQuorumSetIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "QuorumSet")
	if err != nil {
		return nil, err
	}
	return &OracleQuorumSetIterator{contract: _Oracle.contract, event: "QuorumSet", logs: logs, sub: sub}, nil
}

// WatchQuorumSet is a free log subscription operation binding the contract event 0x9f40cfd22fe91777c78f252bd21a710f3fb007dc2f321876891e7644ba0ae175.
//
// Solidity: event QuorumSet(uint256 newQuorum, uint256 totalMembers, uint256 prevQuorum)
func (_Oracle *OracleFilterer) WatchQuorumSet(opts *bind.WatchOpts, sink chan<- *OracleQuorumSet) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "QuorumSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleQuorumSet)
				if err := _Oracle.contract.UnpackLog(event, "QuorumSet", log); err != nil {
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

// ParseQuorumSet is a log parse operation binding the contract event 0x9f40cfd22fe91777c78f252bd21a710f3fb007dc2f321876891e7644ba0ae175.
//
// Solidity: event QuorumSet(uint256 newQuorum, uint256 totalMembers, uint256 prevQuorum)
func (_Oracle *OracleFilterer) ParseQuorumSet(log types.Log) (*OracleQuorumSet, error) {
	event := new(OracleQuorumSet)
	if err := _Oracle.contract.UnpackLog(event, "QuorumSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleReportProcessorSetIterator is returned from FilterReportProcessorSet and is used to iterate over the raw logs and unpacked data for ReportProcessorSet events raised by the Oracle contract.
type OracleReportProcessorSetIterator struct {
	Event *OracleReportProcessorSet // Event containing the contract specifics and raw log

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
func (it *OracleReportProcessorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleReportProcessorSet)
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
		it.Event = new(OracleReportProcessorSet)
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
func (it *OracleReportProcessorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleReportProcessorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleReportProcessorSet represents a ReportProcessorSet event raised by the Oracle contract.
type OracleReportProcessorSet struct {
	Processor     common.Address
	PrevProcessor common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReportProcessorSet is a free log retrieval operation binding the contract event 0x3b59429457a41af89ea682ac9ed8abb8e99eb5c7d3363d5eedfc6bff6271a81e.
//
// Solidity: event ReportProcessorSet(address indexed processor, address indexed prevProcessor)
func (_Oracle *OracleFilterer) FilterReportProcessorSet(opts *bind.FilterOpts, processor []common.Address, prevProcessor []common.Address) (*OracleReportProcessorSetIterator, error) {

	var processorRule []interface{}
	for _, processorItem := range processor {
		processorRule = append(processorRule, processorItem)
	}
	var prevProcessorRule []interface{}
	for _, prevProcessorItem := range prevProcessor {
		prevProcessorRule = append(prevProcessorRule, prevProcessorItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "ReportProcessorSet", processorRule, prevProcessorRule)
	if err != nil {
		return nil, err
	}
	return &OracleReportProcessorSetIterator{contract: _Oracle.contract, event: "ReportProcessorSet", logs: logs, sub: sub}, nil
}

// WatchReportProcessorSet is a free log subscription operation binding the contract event 0x3b59429457a41af89ea682ac9ed8abb8e99eb5c7d3363d5eedfc6bff6271a81e.
//
// Solidity: event ReportProcessorSet(address indexed processor, address indexed prevProcessor)
func (_Oracle *OracleFilterer) WatchReportProcessorSet(opts *bind.WatchOpts, sink chan<- *OracleReportProcessorSet, processor []common.Address, prevProcessor []common.Address) (event.Subscription, error) {

	var processorRule []interface{}
	for _, processorItem := range processor {
		processorRule = append(processorRule, processorItem)
	}
	var prevProcessorRule []interface{}
	for _, prevProcessorItem := range prevProcessor {
		prevProcessorRule = append(prevProcessorRule, prevProcessorItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "ReportProcessorSet", processorRule, prevProcessorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleReportProcessorSet)
				if err := _Oracle.contract.UnpackLog(event, "ReportProcessorSet", log); err != nil {
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

// ParseReportProcessorSet is a log parse operation binding the contract event 0x3b59429457a41af89ea682ac9ed8abb8e99eb5c7d3363d5eedfc6bff6271a81e.
//
// Solidity: event ReportProcessorSet(address indexed processor, address indexed prevProcessor)
func (_Oracle *OracleFilterer) ParseReportProcessorSet(log types.Log) (*OracleReportProcessorSet, error) {
	event := new(OracleReportProcessorSet)
	if err := _Oracle.contract.UnpackLog(event, "ReportProcessorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleReportReceivedIterator is returned from FilterReportReceived and is used to iterate over the raw logs and unpacked data for ReportReceived events raised by the Oracle contract.
type OracleReportReceivedIterator struct {
	Event *OracleReportReceived // Event containing the contract specifics and raw log

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
func (it *OracleReportReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleReportReceived)
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
		it.Event = new(OracleReportReceived)
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
func (it *OracleReportReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleReportReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleReportReceived represents a ReportReceived event raised by the Oracle contract.
type OracleReportReceived struct {
	RefSlot *big.Int
	Member  common.Address
	Report  [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReportReceived is a free log retrieval operation binding the contract event 0x92f77576dabd7bad26f75c36abb3021b5bbb66a3e5688570a0355daddd417488.
//
// Solidity: event ReportReceived(uint256 indexed refSlot, address indexed member, bytes32 report)
func (_Oracle *OracleFilterer) FilterReportReceived(opts *bind.FilterOpts, refSlot []*big.Int, member []common.Address) (*OracleReportReceivedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}
	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "ReportReceived", refSlotRule, memberRule)
	if err != nil {
		return nil, err
	}
	return &OracleReportReceivedIterator{contract: _Oracle.contract, event: "ReportReceived", logs: logs, sub: sub}, nil
}

// WatchReportReceived is a free log subscription operation binding the contract event 0x92f77576dabd7bad26f75c36abb3021b5bbb66a3e5688570a0355daddd417488.
//
// Solidity: event ReportReceived(uint256 indexed refSlot, address indexed member, bytes32 report)
func (_Oracle *OracleFilterer) WatchReportReceived(opts *bind.WatchOpts, sink chan<- *OracleReportReceived, refSlot []*big.Int, member []common.Address) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}
	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "ReportReceived", refSlotRule, memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleReportReceived)
				if err := _Oracle.contract.UnpackLog(event, "ReportReceived", log); err != nil {
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

// ParseReportReceived is a log parse operation binding the contract event 0x92f77576dabd7bad26f75c36abb3021b5bbb66a3e5688570a0355daddd417488.
//
// Solidity: event ReportReceived(uint256 indexed refSlot, address indexed member, bytes32 report)
func (_Oracle *OracleFilterer) ParseReportReceived(log types.Log) (*OracleReportReceived, error) {
	event := new(OracleReportReceived)
	if err := _Oracle.contract.UnpackLog(event, "ReportReceived", log); err != nil {
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

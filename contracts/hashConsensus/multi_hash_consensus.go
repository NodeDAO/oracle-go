// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hashConsensus

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

// MultiHashConsensusMemberConsensusState is an auto generated low-level Go binding around an user-defined struct.
type MultiHashConsensusMemberConsensusState struct {
	CurrentFrameRefSlot         *big.Int
	CurrentFrameConsensusReport [][32]byte
	IsMember                    bool
	IsFastLane                  bool
	CanReport                   bool
	LastMemberReportRefSlot     *big.Int
	CurrentFrameMemberReport    [][32]byte
}

// HashConsensusMetaData contains all meta data concerning the HashConsensus contract.
var HashConsensusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConsensusReportAlreadyProcessing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DaoCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateMember\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateReport\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateReportProcessor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReport\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EpochsPerFrameCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FastLanePeriodCannotBeLongerThanFrame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitialEpochAlreadyArrived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitialEpochIsYetToArrive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitialEpochRefSlotCannotBeEarlierThanProcessingSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonFastLaneMemberCannotReportWithinFastLaneInterval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonMember\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NumericOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedQuorum\",\"type\":\"uint256\"}],\"name\":\"QuorumTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReportLenNotEqualReportProcessorsLen\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReportProcessorCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reportProcessor\",\"type\":\"address\"}],\"name\":\"ReportProcessorExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reportProcessor\",\"type\":\"address\"}],\"name\":\"ReportProcessorNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"UnexpectedConsensusVersion\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"report\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"support\",\"type\":\"uint256\"}],\"name\":\"ConsensusReached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldDao\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"DaoAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fastLaneLengthSlots\",\"type\":\"uint256\"}],\"name\":\"FastLaneConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newInitialEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newEpochsPerFrame\",\"type\":\"uint256\"}],\"name\":\"FrameConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalMembers\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuorum\",\"type\":\"uint256\"}],\"name\":\"MemberAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalMembers\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuorum\",\"type\":\"uint256\"}],\"name\":\"MemberRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuorum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalMembers\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevQuorum\",\"type\":\"uint256\"}],\"name\":\"QuorumSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"processor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"moduleId\",\"type\":\"uint256\"}],\"name\":\"ReportProcessorAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldProcessor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newProcessor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"moduleId\",\"type\":\"uint256\"}],\"name\":\"ReportProcessorUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"report\",\"type\":\"bytes32[]\"}],\"name\":\"ReportReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quorum\",\"type\":\"uint256\"}],\"name\":\"addMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newProcessor\",\"type\":\"address\"}],\"name\":\"addReportProcessor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableConsensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"slotsPerEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondsPerSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"genesisTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"consensusReport\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getConsensusStateForMember\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currentFrameRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"currentFrameConsensusReport\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool\",\"name\":\"isMember\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isFastLane\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canReport\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"lastMemberReportRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"currentFrameMemberReport\",\"type\":\"bytes32[]\"}],\"internalType\":\"structMultiHashConsensus.MemberConsensusState\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentFrame\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reportProcessingDeadlineSlot\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFastLaneMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"lastReportedRefSlots\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFrameConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"initialEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochsPerFrame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fastLaneLengthSlots\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitialRefSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getIsFastLaneMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getIsMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reportProcessor\",\"type\":\"address\"}],\"name\":\"getIsReportProcessing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getIsReportProcessor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"lastReportedRefSlots\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQuorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reportProcessor\",\"type\":\"address\"}],\"name\":\"getReportModuleId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportProcessors\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportVariants\",\"outputs\":[{\"internalType\":\"bytes32[][]\",\"name\":\"variants\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[]\",\"name\":\"support\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotsPerEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondsPerSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"genesisTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochsPerFrame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fastLaneLengthSlots\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reportProcessor\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quorum\",\"type\":\"uint256\"}],\"name\":\"removeMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"setDaoAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fastLaneLengthSlots\",\"type\":\"uint256\"}],\"name\":\"setFastLaneLengthSlots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochsPerFrame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fastLaneLengthSlots\",\"type\":\"uint256\"}],\"name\":\"setFrameConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quorum\",\"type\":\"uint256\"}],\"name\":\"setQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"report\",\"type\":\"bytes32[]\"}],\"name\":\"submitReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialEpoch\",\"type\":\"uint256\"}],\"name\":\"updateInitialEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldProcessor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newProcessor\",\"type\":\"address\"}],\"name\":\"updateReportProcessor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// HashConsensusABI is the input ABI used to generate the binding from.
// Deprecated: Use HashConsensusMetaData.ABI instead.
var HashConsensusABI = HashConsensusMetaData.ABI

// HashConsensus is an auto generated Go binding around an Ethereum contract.
type HashConsensus struct {
	HashConsensusCaller     // Read-only binding to the contract
	HashConsensusTransactor // Write-only binding to the contract
	HashConsensusFilterer   // Log filterer for contract events
}

// HashConsensusCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashConsensusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashConsensusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashConsensusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashConsensusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashConsensusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashConsensusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashConsensusSession struct {
	Contract     *HashConsensus    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashConsensusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashConsensusCallerSession struct {
	Contract *HashConsensusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// HashConsensusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashConsensusTransactorSession struct {
	Contract     *HashConsensusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// HashConsensusRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashConsensusRaw struct {
	Contract *HashConsensus // Generic contract binding to access the raw methods on
}

// HashConsensusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashConsensusCallerRaw struct {
	Contract *HashConsensusCaller // Generic read-only contract binding to access the raw methods on
}

// HashConsensusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashConsensusTransactorRaw struct {
	Contract *HashConsensusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashConsensus creates a new instance of HashConsensus, bound to a specific deployed contract.
func NewHashConsensus(address common.Address, backend bind.ContractBackend) (*HashConsensus, error) {
	contract, err := bindHashConsensus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashConsensus{HashConsensusCaller: HashConsensusCaller{contract: contract}, HashConsensusTransactor: HashConsensusTransactor{contract: contract}, HashConsensusFilterer: HashConsensusFilterer{contract: contract}}, nil
}

// NewHashConsensusCaller creates a new read-only instance of HashConsensus, bound to a specific deployed contract.
func NewHashConsensusCaller(address common.Address, caller bind.ContractCaller) (*HashConsensusCaller, error) {
	contract, err := bindHashConsensus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashConsensusCaller{contract: contract}, nil
}

// NewHashConsensusTransactor creates a new write-only instance of HashConsensus, bound to a specific deployed contract.
func NewHashConsensusTransactor(address common.Address, transactor bind.ContractTransactor) (*HashConsensusTransactor, error) {
	contract, err := bindHashConsensus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashConsensusTransactor{contract: contract}, nil
}

// NewHashConsensusFilterer creates a new log filterer instance of HashConsensus, bound to a specific deployed contract.
func NewHashConsensusFilterer(address common.Address, filterer bind.ContractFilterer) (*HashConsensusFilterer, error) {
	contract, err := bindHashConsensus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashConsensusFilterer{contract: contract}, nil
}

// bindHashConsensus binds a generic wrapper to an already deployed contract.
func bindHashConsensus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HashConsensusMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashConsensus *HashConsensusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashConsensus.Contract.HashConsensusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashConsensus *HashConsensusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashConsensus.Contract.HashConsensusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashConsensus *HashConsensusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashConsensus.Contract.HashConsensusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashConsensus *HashConsensusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashConsensus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashConsensus *HashConsensusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashConsensus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashConsensus *HashConsensusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashConsensus.Contract.contract.Transact(opts, method, params...)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_HashConsensus *HashConsensusCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_HashConsensus *HashConsensusSession) Dao() (common.Address, error) {
	return _HashConsensus.Contract.Dao(&_HashConsensus.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_HashConsensus *HashConsensusCallerSession) Dao() (common.Address, error) {
	return _HashConsensus.Contract.Dao(&_HashConsensus.CallOpts)
}

// GetChainConfig is a free data retrieval call binding the contract method 0x606c0c94.
//
// Solidity: function getChainConfig() view returns(uint256 slotsPerEpoch, uint256 secondsPerSlot, uint256 genesisTime)
func (_HashConsensus *HashConsensusCaller) GetChainConfig(opts *bind.CallOpts) (struct {
	SlotsPerEpoch  *big.Int
	SecondsPerSlot *big.Int
	GenesisTime    *big.Int
}, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "getChainConfig")

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
func (_HashConsensus *HashConsensusSession) GetChainConfig() (struct {
	SlotsPerEpoch  *big.Int
	SecondsPerSlot *big.Int
	GenesisTime    *big.Int
}, error) {
	return _HashConsensus.Contract.GetChainConfig(&_HashConsensus.CallOpts)
}

// GetChainConfig is a free data retrieval call binding the contract method 0x606c0c94.
//
// Solidity: function getChainConfig() view returns(uint256 slotsPerEpoch, uint256 secondsPerSlot, uint256 genesisTime)
func (_HashConsensus *HashConsensusCallerSession) GetChainConfig() (struct {
	SlotsPerEpoch  *big.Int
	SecondsPerSlot *big.Int
	GenesisTime    *big.Int
}, error) {
	return _HashConsensus.Contract.GetChainConfig(&_HashConsensus.CallOpts)
}

// GetConsensusState is a free data retrieval call binding the contract method 0xea87627d.
//
// Solidity: function getConsensusState() view returns(uint256 refSlot, bytes32[] consensusReport)
func (_HashConsensus *HashConsensusCaller) GetConsensusState(opts *bind.CallOpts) (struct {
	RefSlot         *big.Int
	ConsensusReport [][32]byte
}, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "getConsensusState")

	outstruct := new(struct {
		RefSlot         *big.Int
		ConsensusReport [][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RefSlot = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ConsensusReport = *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

// GetConsensusState is a free data retrieval call binding the contract method 0xea87627d.
//
// Solidity: function getConsensusState() view returns(uint256 refSlot, bytes32[] consensusReport)
func (_HashConsensus *HashConsensusSession) GetConsensusState() (struct {
	RefSlot         *big.Int
	ConsensusReport [][32]byte
}, error) {
	return _HashConsensus.Contract.GetConsensusState(&_HashConsensus.CallOpts)
}

// GetConsensusState is a free data retrieval call binding the contract method 0xea87627d.
//
// Solidity: function getConsensusState() view returns(uint256 refSlot, bytes32[] consensusReport)
func (_HashConsensus *HashConsensusCallerSession) GetConsensusState() (struct {
	RefSlot         *big.Int
	ConsensusReport [][32]byte
}, error) {
	return _HashConsensus.Contract.GetConsensusState(&_HashConsensus.CallOpts)
}

// GetConsensusStateForMember is a free data retrieval call binding the contract method 0x60e61801.
//
// Solidity: function getConsensusStateForMember(address addr) view returns((uint256,bytes32[],bool,bool,bool,uint256,bytes32[]) result)
func (_HashConsensus *HashConsensusCaller) GetConsensusStateForMember(opts *bind.CallOpts, addr common.Address) (MultiHashConsensusMemberConsensusState, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "getConsensusStateForMember", addr)

	if err != nil {
		return *new(MultiHashConsensusMemberConsensusState), err
	}

	out0 := *abi.ConvertType(out[0], new(MultiHashConsensusMemberConsensusState)).(*MultiHashConsensusMemberConsensusState)

	return out0, err

}

// GetConsensusStateForMember is a free data retrieval call binding the contract method 0x60e61801.
//
// Solidity: function getConsensusStateForMember(address addr) view returns((uint256,bytes32[],bool,bool,bool,uint256,bytes32[]) result)
func (_HashConsensus *HashConsensusSession) GetConsensusStateForMember(addr common.Address) (MultiHashConsensusMemberConsensusState, error) {
	return _HashConsensus.Contract.GetConsensusStateForMember(&_HashConsensus.CallOpts, addr)
}

// GetConsensusStateForMember is a free data retrieval call binding the contract method 0x60e61801.
//
// Solidity: function getConsensusStateForMember(address addr) view returns((uint256,bytes32[],bool,bool,bool,uint256,bytes32[]) result)
func (_HashConsensus *HashConsensusCallerSession) GetConsensusStateForMember(addr common.Address) (MultiHashConsensusMemberConsensusState, error) {
	return _HashConsensus.Contract.GetConsensusStateForMember(&_HashConsensus.CallOpts, addr)
}

// GetCurrentFrame is a free data retrieval call binding the contract method 0x72f79b13.
//
// Solidity: function getCurrentFrame() view returns(uint256 refSlot, uint256 reportProcessingDeadlineSlot)
func (_HashConsensus *HashConsensusCaller) GetCurrentFrame(opts *bind.CallOpts) (struct {
	RefSlot                      *big.Int
	ReportProcessingDeadlineSlot *big.Int
}, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "getCurrentFrame")

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
func (_HashConsensus *HashConsensusSession) GetCurrentFrame() (struct {
	RefSlot                      *big.Int
	ReportProcessingDeadlineSlot *big.Int
}, error) {
	return _HashConsensus.Contract.GetCurrentFrame(&_HashConsensus.CallOpts)
}

// GetCurrentFrame is a free data retrieval call binding the contract method 0x72f79b13.
//
// Solidity: function getCurrentFrame() view returns(uint256 refSlot, uint256 reportProcessingDeadlineSlot)
func (_HashConsensus *HashConsensusCallerSession) GetCurrentFrame() (struct {
	RefSlot                      *big.Int
	ReportProcessingDeadlineSlot *big.Int
}, error) {
	return _HashConsensus.Contract.GetCurrentFrame(&_HashConsensus.CallOpts)
}

// GetFastLaneMembers is a free data retrieval call binding the contract method 0x433ab1f3.
//
// Solidity: function getFastLaneMembers() view returns(address[] addresses, uint256[] lastReportedRefSlots)
func (_HashConsensus *HashConsensusCaller) GetFastLaneMembers(opts *bind.CallOpts) (struct {
	Addresses            []common.Address
	LastReportedRefSlots []*big.Int
}, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "getFastLaneMembers")

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
func (_HashConsensus *HashConsensusSession) GetFastLaneMembers() (struct {
	Addresses            []common.Address
	LastReportedRefSlots []*big.Int
}, error) {
	return _HashConsensus.Contract.GetFastLaneMembers(&_HashConsensus.CallOpts)
}

// GetFastLaneMembers is a free data retrieval call binding the contract method 0x433ab1f3.
//
// Solidity: function getFastLaneMembers() view returns(address[] addresses, uint256[] lastReportedRefSlots)
func (_HashConsensus *HashConsensusCallerSession) GetFastLaneMembers() (struct {
	Addresses            []common.Address
	LastReportedRefSlots []*big.Int
}, error) {
	return _HashConsensus.Contract.GetFastLaneMembers(&_HashConsensus.CallOpts)
}

// GetFrameConfig is a free data retrieval call binding the contract method 0x6fb1bf66.
//
// Solidity: function getFrameConfig() view returns(uint256 initialEpoch, uint256 epochsPerFrame, uint256 fastLaneLengthSlots)
func (_HashConsensus *HashConsensusCaller) GetFrameConfig(opts *bind.CallOpts) (struct {
	InitialEpoch        *big.Int
	EpochsPerFrame      *big.Int
	FastLaneLengthSlots *big.Int
}, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "getFrameConfig")

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
func (_HashConsensus *HashConsensusSession) GetFrameConfig() (struct {
	InitialEpoch        *big.Int
	EpochsPerFrame      *big.Int
	FastLaneLengthSlots *big.Int
}, error) {
	return _HashConsensus.Contract.GetFrameConfig(&_HashConsensus.CallOpts)
}

// GetFrameConfig is a free data retrieval call binding the contract method 0x6fb1bf66.
//
// Solidity: function getFrameConfig() view returns(uint256 initialEpoch, uint256 epochsPerFrame, uint256 fastLaneLengthSlots)
func (_HashConsensus *HashConsensusCallerSession) GetFrameConfig() (struct {
	InitialEpoch        *big.Int
	EpochsPerFrame      *big.Int
	FastLaneLengthSlots *big.Int
}, error) {
	return _HashConsensus.Contract.GetFrameConfig(&_HashConsensus.CallOpts)
}

// GetInitialRefSlot is a free data retrieval call binding the contract method 0x6095012f.
//
// Solidity: function getInitialRefSlot() view returns(uint256)
func (_HashConsensus *HashConsensusCaller) GetInitialRefSlot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "getInitialRefSlot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitialRefSlot is a free data retrieval call binding the contract method 0x6095012f.
//
// Solidity: function getInitialRefSlot() view returns(uint256)
func (_HashConsensus *HashConsensusSession) GetInitialRefSlot() (*big.Int, error) {
	return _HashConsensus.Contract.GetInitialRefSlot(&_HashConsensus.CallOpts)
}

// GetInitialRefSlot is a free data retrieval call binding the contract method 0x6095012f.
//
// Solidity: function getInitialRefSlot() view returns(uint256)
func (_HashConsensus *HashConsensusCallerSession) GetInitialRefSlot() (*big.Int, error) {
	return _HashConsensus.Contract.GetInitialRefSlot(&_HashConsensus.CallOpts)
}

// GetIsFastLaneMember is a free data retrieval call binding the contract method 0x20b4d751.
//
// Solidity: function getIsFastLaneMember(address addr) view returns(bool)
func (_HashConsensus *HashConsensusCaller) GetIsFastLaneMember(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "getIsFastLaneMember", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsFastLaneMember is a free data retrieval call binding the contract method 0x20b4d751.
//
// Solidity: function getIsFastLaneMember(address addr) view returns(bool)
func (_HashConsensus *HashConsensusSession) GetIsFastLaneMember(addr common.Address) (bool, error) {
	return _HashConsensus.Contract.GetIsFastLaneMember(&_HashConsensus.CallOpts, addr)
}

// GetIsFastLaneMember is a free data retrieval call binding the contract method 0x20b4d751.
//
// Solidity: function getIsFastLaneMember(address addr) view returns(bool)
func (_HashConsensus *HashConsensusCallerSession) GetIsFastLaneMember(addr common.Address) (bool, error) {
	return _HashConsensus.Contract.GetIsFastLaneMember(&_HashConsensus.CallOpts, addr)
}

// GetIsMember is a free data retrieval call binding the contract method 0x1951c037.
//
// Solidity: function getIsMember(address addr) view returns(bool)
func (_HashConsensus *HashConsensusCaller) GetIsMember(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "getIsMember", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsMember is a free data retrieval call binding the contract method 0x1951c037.
//
// Solidity: function getIsMember(address addr) view returns(bool)
func (_HashConsensus *HashConsensusSession) GetIsMember(addr common.Address) (bool, error) {
	return _HashConsensus.Contract.GetIsMember(&_HashConsensus.CallOpts, addr)
}

// GetIsMember is a free data retrieval call binding the contract method 0x1951c037.
//
// Solidity: function getIsMember(address addr) view returns(bool)
func (_HashConsensus *HashConsensusCallerSession) GetIsMember(addr common.Address) (bool, error) {
	return _HashConsensus.Contract.GetIsMember(&_HashConsensus.CallOpts, addr)
}

// GetIsReportProcessing is a free data retrieval call binding the contract method 0x948073c6.
//
// Solidity: function getIsReportProcessing(address _reportProcessor) view returns(bool)
func (_HashConsensus *HashConsensusCaller) GetIsReportProcessing(opts *bind.CallOpts, _reportProcessor common.Address) (bool, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "getIsReportProcessing", _reportProcessor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsReportProcessing is a free data retrieval call binding the contract method 0x948073c6.
//
// Solidity: function getIsReportProcessing(address _reportProcessor) view returns(bool)
func (_HashConsensus *HashConsensusSession) GetIsReportProcessing(_reportProcessor common.Address) (bool, error) {
	return _HashConsensus.Contract.GetIsReportProcessing(&_HashConsensus.CallOpts, _reportProcessor)
}

// GetIsReportProcessing is a free data retrieval call binding the contract method 0x948073c6.
//
// Solidity: function getIsReportProcessing(address _reportProcessor) view returns(bool)
func (_HashConsensus *HashConsensusCallerSession) GetIsReportProcessing(_reportProcessor common.Address) (bool, error) {
	return _HashConsensus.Contract.GetIsReportProcessing(&_HashConsensus.CallOpts, _reportProcessor)
}

// GetIsReportProcessor is a free data retrieval call binding the contract method 0x2117f9a9.
//
// Solidity: function getIsReportProcessor(address addr) view returns(bool)
func (_HashConsensus *HashConsensusCaller) GetIsReportProcessor(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "getIsReportProcessor", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsReportProcessor is a free data retrieval call binding the contract method 0x2117f9a9.
//
// Solidity: function getIsReportProcessor(address addr) view returns(bool)
func (_HashConsensus *HashConsensusSession) GetIsReportProcessor(addr common.Address) (bool, error) {
	return _HashConsensus.Contract.GetIsReportProcessor(&_HashConsensus.CallOpts, addr)
}

// GetIsReportProcessor is a free data retrieval call binding the contract method 0x2117f9a9.
//
// Solidity: function getIsReportProcessor(address addr) view returns(bool)
func (_HashConsensus *HashConsensusCallerSession) GetIsReportProcessor(addr common.Address) (bool, error) {
	return _HashConsensus.Contract.GetIsReportProcessor(&_HashConsensus.CallOpts, addr)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[] addresses, uint256[] lastReportedRefSlots)
func (_HashConsensus *HashConsensusCaller) GetMembers(opts *bind.CallOpts) (struct {
	Addresses            []common.Address
	LastReportedRefSlots []*big.Int
}, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "getMembers")

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
func (_HashConsensus *HashConsensusSession) GetMembers() (struct {
	Addresses            []common.Address
	LastReportedRefSlots []*big.Int
}, error) {
	return _HashConsensus.Contract.GetMembers(&_HashConsensus.CallOpts)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[] addresses, uint256[] lastReportedRefSlots)
func (_HashConsensus *HashConsensusCallerSession) GetMembers() (struct {
	Addresses            []common.Address
	LastReportedRefSlots []*big.Int
}, error) {
	return _HashConsensus.Contract.GetMembers(&_HashConsensus.CallOpts)
}

// GetQuorum is a free data retrieval call binding the contract method 0xc26c12eb.
//
// Solidity: function getQuorum() view returns(uint256)
func (_HashConsensus *HashConsensusCaller) GetQuorum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "getQuorum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorum is a free data retrieval call binding the contract method 0xc26c12eb.
//
// Solidity: function getQuorum() view returns(uint256)
func (_HashConsensus *HashConsensusSession) GetQuorum() (*big.Int, error) {
	return _HashConsensus.Contract.GetQuorum(&_HashConsensus.CallOpts)
}

// GetQuorum is a free data retrieval call binding the contract method 0xc26c12eb.
//
// Solidity: function getQuorum() view returns(uint256)
func (_HashConsensus *HashConsensusCallerSession) GetQuorum() (*big.Int, error) {
	return _HashConsensus.Contract.GetQuorum(&_HashConsensus.CallOpts)
}

// GetReportModuleId is a free data retrieval call binding the contract method 0x1300f68d.
//
// Solidity: function getReportModuleId(address reportProcessor) view returns(uint256)
func (_HashConsensus *HashConsensusCaller) GetReportModuleId(opts *bind.CallOpts, reportProcessor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "getReportModuleId", reportProcessor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReportModuleId is a free data retrieval call binding the contract method 0x1300f68d.
//
// Solidity: function getReportModuleId(address reportProcessor) view returns(uint256)
func (_HashConsensus *HashConsensusSession) GetReportModuleId(reportProcessor common.Address) (*big.Int, error) {
	return _HashConsensus.Contract.GetReportModuleId(&_HashConsensus.CallOpts, reportProcessor)
}

// GetReportModuleId is a free data retrieval call binding the contract method 0x1300f68d.
//
// Solidity: function getReportModuleId(address reportProcessor) view returns(uint256)
func (_HashConsensus *HashConsensusCallerSession) GetReportModuleId(reportProcessor common.Address) (*big.Int, error) {
	return _HashConsensus.Contract.GetReportModuleId(&_HashConsensus.CallOpts, reportProcessor)
}

// GetReportProcessors is a free data retrieval call binding the contract method 0xb6ac0119.
//
// Solidity: function getReportProcessors() view returns(address[])
func (_HashConsensus *HashConsensusCaller) GetReportProcessors(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "getReportProcessors")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReportProcessors is a free data retrieval call binding the contract method 0xb6ac0119.
//
// Solidity: function getReportProcessors() view returns(address[])
func (_HashConsensus *HashConsensusSession) GetReportProcessors() ([]common.Address, error) {
	return _HashConsensus.Contract.GetReportProcessors(&_HashConsensus.CallOpts)
}

// GetReportProcessors is a free data retrieval call binding the contract method 0xb6ac0119.
//
// Solidity: function getReportProcessors() view returns(address[])
func (_HashConsensus *HashConsensusCallerSession) GetReportProcessors() ([]common.Address, error) {
	return _HashConsensus.Contract.GetReportProcessors(&_HashConsensus.CallOpts)
}

// GetReportVariants is a free data retrieval call binding the contract method 0x2fd2d750.
//
// Solidity: function getReportVariants() view returns(bytes32[][] variants, uint256[] support)
func (_HashConsensus *HashConsensusCaller) GetReportVariants(opts *bind.CallOpts) (struct {
	Variants [][][32]byte
	Support  []*big.Int
}, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "getReportVariants")

	outstruct := new(struct {
		Variants [][][32]byte
		Support  []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Variants = *abi.ConvertType(out[0], new([][][32]byte)).(*[][][32]byte)
	outstruct.Support = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetReportVariants is a free data retrieval call binding the contract method 0x2fd2d750.
//
// Solidity: function getReportVariants() view returns(bytes32[][] variants, uint256[] support)
func (_HashConsensus *HashConsensusSession) GetReportVariants() (struct {
	Variants [][][32]byte
	Support  []*big.Int
}, error) {
	return _HashConsensus.Contract.GetReportVariants(&_HashConsensus.CallOpts)
}

// GetReportVariants is a free data retrieval call binding the contract method 0x2fd2d750.
//
// Solidity: function getReportVariants() view returns(bytes32[][] variants, uint256[] support)
func (_HashConsensus *HashConsensusCallerSession) GetReportVariants() (struct {
	Variants [][][32]byte
	Support  []*big.Int
}, error) {
	return _HashConsensus.Contract.GetReportVariants(&_HashConsensus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HashConsensus *HashConsensusCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HashConsensus *HashConsensusSession) Owner() (common.Address, error) {
	return _HashConsensus.Contract.Owner(&_HashConsensus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HashConsensus *HashConsensusCallerSession) Owner() (common.Address, error) {
	return _HashConsensus.Contract.Owner(&_HashConsensus.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_HashConsensus *HashConsensusCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HashConsensus.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_HashConsensus *HashConsensusSession) ProxiableUUID() ([32]byte, error) {
	return _HashConsensus.Contract.ProxiableUUID(&_HashConsensus.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_HashConsensus *HashConsensusCallerSession) ProxiableUUID() ([32]byte, error) {
	return _HashConsensus.Contract.ProxiableUUID(&_HashConsensus.CallOpts)
}

// AddMember is a paid mutator transaction binding the contract method 0x98041ea3.
//
// Solidity: function addMember(address addr, uint256 quorum) returns()
func (_HashConsensus *HashConsensusTransactor) AddMember(opts *bind.TransactOpts, addr common.Address, quorum *big.Int) (*types.Transaction, error) {
	return _HashConsensus.contract.Transact(opts, "addMember", addr, quorum)
}

// AddMember is a paid mutator transaction binding the contract method 0x98041ea3.
//
// Solidity: function addMember(address addr, uint256 quorum) returns()
func (_HashConsensus *HashConsensusSession) AddMember(addr common.Address, quorum *big.Int) (*types.Transaction, error) {
	return _HashConsensus.Contract.AddMember(&_HashConsensus.TransactOpts, addr, quorum)
}

// AddMember is a paid mutator transaction binding the contract method 0x98041ea3.
//
// Solidity: function addMember(address addr, uint256 quorum) returns()
func (_HashConsensus *HashConsensusTransactorSession) AddMember(addr common.Address, quorum *big.Int) (*types.Transaction, error) {
	return _HashConsensus.Contract.AddMember(&_HashConsensus.TransactOpts, addr, quorum)
}

// AddReportProcessor is a paid mutator transaction binding the contract method 0x546cd29c.
//
// Solidity: function addReportProcessor(address newProcessor) returns()
func (_HashConsensus *HashConsensusTransactor) AddReportProcessor(opts *bind.TransactOpts, newProcessor common.Address) (*types.Transaction, error) {
	return _HashConsensus.contract.Transact(opts, "addReportProcessor", newProcessor)
}

// AddReportProcessor is a paid mutator transaction binding the contract method 0x546cd29c.
//
// Solidity: function addReportProcessor(address newProcessor) returns()
func (_HashConsensus *HashConsensusSession) AddReportProcessor(newProcessor common.Address) (*types.Transaction, error) {
	return _HashConsensus.Contract.AddReportProcessor(&_HashConsensus.TransactOpts, newProcessor)
}

// AddReportProcessor is a paid mutator transaction binding the contract method 0x546cd29c.
//
// Solidity: function addReportProcessor(address newProcessor) returns()
func (_HashConsensus *HashConsensusTransactorSession) AddReportProcessor(newProcessor common.Address) (*types.Transaction, error) {
	return _HashConsensus.Contract.AddReportProcessor(&_HashConsensus.TransactOpts, newProcessor)
}

// DisableConsensus is a paid mutator transaction binding the contract method 0xad231cb2.
//
// Solidity: function disableConsensus() returns()
func (_HashConsensus *HashConsensusTransactor) DisableConsensus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashConsensus.contract.Transact(opts, "disableConsensus")
}

// DisableConsensus is a paid mutator transaction binding the contract method 0xad231cb2.
//
// Solidity: function disableConsensus() returns()
func (_HashConsensus *HashConsensusSession) DisableConsensus() (*types.Transaction, error) {
	return _HashConsensus.Contract.DisableConsensus(&_HashConsensus.TransactOpts)
}

// DisableConsensus is a paid mutator transaction binding the contract method 0xad231cb2.
//
// Solidity: function disableConsensus() returns()
func (_HashConsensus *HashConsensusTransactorSession) DisableConsensus() (*types.Transaction, error) {
	return _HashConsensus.Contract.DisableConsensus(&_HashConsensus.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x1f85d8bc.
//
// Solidity: function initialize(uint256 slotsPerEpoch, uint256 secondsPerSlot, uint256 genesisTime, uint256 epochsPerFrame, uint256 fastLaneLengthSlots, address _dao, address reportProcessor) returns()
func (_HashConsensus *HashConsensusTransactor) Initialize(opts *bind.TransactOpts, slotsPerEpoch *big.Int, secondsPerSlot *big.Int, genesisTime *big.Int, epochsPerFrame *big.Int, fastLaneLengthSlots *big.Int, _dao common.Address, reportProcessor common.Address) (*types.Transaction, error) {
	return _HashConsensus.contract.Transact(opts, "initialize", slotsPerEpoch, secondsPerSlot, genesisTime, epochsPerFrame, fastLaneLengthSlots, _dao, reportProcessor)
}

// Initialize is a paid mutator transaction binding the contract method 0x1f85d8bc.
//
// Solidity: function initialize(uint256 slotsPerEpoch, uint256 secondsPerSlot, uint256 genesisTime, uint256 epochsPerFrame, uint256 fastLaneLengthSlots, address _dao, address reportProcessor) returns()
func (_HashConsensus *HashConsensusSession) Initialize(slotsPerEpoch *big.Int, secondsPerSlot *big.Int, genesisTime *big.Int, epochsPerFrame *big.Int, fastLaneLengthSlots *big.Int, _dao common.Address, reportProcessor common.Address) (*types.Transaction, error) {
	return _HashConsensus.Contract.Initialize(&_HashConsensus.TransactOpts, slotsPerEpoch, secondsPerSlot, genesisTime, epochsPerFrame, fastLaneLengthSlots, _dao, reportProcessor)
}

// Initialize is a paid mutator transaction binding the contract method 0x1f85d8bc.
//
// Solidity: function initialize(uint256 slotsPerEpoch, uint256 secondsPerSlot, uint256 genesisTime, uint256 epochsPerFrame, uint256 fastLaneLengthSlots, address _dao, address reportProcessor) returns()
func (_HashConsensus *HashConsensusTransactorSession) Initialize(slotsPerEpoch *big.Int, secondsPerSlot *big.Int, genesisTime *big.Int, epochsPerFrame *big.Int, fastLaneLengthSlots *big.Int, _dao common.Address, reportProcessor common.Address) (*types.Transaction, error) {
	return _HashConsensus.Contract.Initialize(&_HashConsensus.TransactOpts, slotsPerEpoch, secondsPerSlot, genesisTime, epochsPerFrame, fastLaneLengthSlots, _dao, reportProcessor)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x16f6f03e.
//
// Solidity: function removeMember(address addr, uint256 quorum) returns()
func (_HashConsensus *HashConsensusTransactor) RemoveMember(opts *bind.TransactOpts, addr common.Address, quorum *big.Int) (*types.Transaction, error) {
	return _HashConsensus.contract.Transact(opts, "removeMember", addr, quorum)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x16f6f03e.
//
// Solidity: function removeMember(address addr, uint256 quorum) returns()
func (_HashConsensus *HashConsensusSession) RemoveMember(addr common.Address, quorum *big.Int) (*types.Transaction, error) {
	return _HashConsensus.Contract.RemoveMember(&_HashConsensus.TransactOpts, addr, quorum)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x16f6f03e.
//
// Solidity: function removeMember(address addr, uint256 quorum) returns()
func (_HashConsensus *HashConsensusTransactorSession) RemoveMember(addr common.Address, quorum *big.Int) (*types.Transaction, error) {
	return _HashConsensus.Contract.RemoveMember(&_HashConsensus.TransactOpts, addr, quorum)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HashConsensus *HashConsensusTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashConsensus.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HashConsensus *HashConsensusSession) RenounceOwnership() (*types.Transaction, error) {
	return _HashConsensus.Contract.RenounceOwnership(&_HashConsensus.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HashConsensus *HashConsensusTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _HashConsensus.Contract.RenounceOwnership(&_HashConsensus.TransactOpts)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_HashConsensus *HashConsensusTransactor) SetDaoAddress(opts *bind.TransactOpts, _dao common.Address) (*types.Transaction, error) {
	return _HashConsensus.contract.Transact(opts, "setDaoAddress", _dao)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_HashConsensus *HashConsensusSession) SetDaoAddress(_dao common.Address) (*types.Transaction, error) {
	return _HashConsensus.Contract.SetDaoAddress(&_HashConsensus.TransactOpts, _dao)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_HashConsensus *HashConsensusTransactorSession) SetDaoAddress(_dao common.Address) (*types.Transaction, error) {
	return _HashConsensus.Contract.SetDaoAddress(&_HashConsensus.TransactOpts, _dao)
}

// SetFastLaneLengthSlots is a paid mutator transaction binding the contract method 0x99229f58.
//
// Solidity: function setFastLaneLengthSlots(uint256 fastLaneLengthSlots) returns()
func (_HashConsensus *HashConsensusTransactor) SetFastLaneLengthSlots(opts *bind.TransactOpts, fastLaneLengthSlots *big.Int) (*types.Transaction, error) {
	return _HashConsensus.contract.Transact(opts, "setFastLaneLengthSlots", fastLaneLengthSlots)
}

// SetFastLaneLengthSlots is a paid mutator transaction binding the contract method 0x99229f58.
//
// Solidity: function setFastLaneLengthSlots(uint256 fastLaneLengthSlots) returns()
func (_HashConsensus *HashConsensusSession) SetFastLaneLengthSlots(fastLaneLengthSlots *big.Int) (*types.Transaction, error) {
	return _HashConsensus.Contract.SetFastLaneLengthSlots(&_HashConsensus.TransactOpts, fastLaneLengthSlots)
}

// SetFastLaneLengthSlots is a paid mutator transaction binding the contract method 0x99229f58.
//
// Solidity: function setFastLaneLengthSlots(uint256 fastLaneLengthSlots) returns()
func (_HashConsensus *HashConsensusTransactorSession) SetFastLaneLengthSlots(fastLaneLengthSlots *big.Int) (*types.Transaction, error) {
	return _HashConsensus.Contract.SetFastLaneLengthSlots(&_HashConsensus.TransactOpts, fastLaneLengthSlots)
}

// SetFrameConfig is a paid mutator transaction binding the contract method 0x34aa6753.
//
// Solidity: function setFrameConfig(uint256 epochsPerFrame, uint256 fastLaneLengthSlots) returns()
func (_HashConsensus *HashConsensusTransactor) SetFrameConfig(opts *bind.TransactOpts, epochsPerFrame *big.Int, fastLaneLengthSlots *big.Int) (*types.Transaction, error) {
	return _HashConsensus.contract.Transact(opts, "setFrameConfig", epochsPerFrame, fastLaneLengthSlots)
}

// SetFrameConfig is a paid mutator transaction binding the contract method 0x34aa6753.
//
// Solidity: function setFrameConfig(uint256 epochsPerFrame, uint256 fastLaneLengthSlots) returns()
func (_HashConsensus *HashConsensusSession) SetFrameConfig(epochsPerFrame *big.Int, fastLaneLengthSlots *big.Int) (*types.Transaction, error) {
	return _HashConsensus.Contract.SetFrameConfig(&_HashConsensus.TransactOpts, epochsPerFrame, fastLaneLengthSlots)
}

// SetFrameConfig is a paid mutator transaction binding the contract method 0x34aa6753.
//
// Solidity: function setFrameConfig(uint256 epochsPerFrame, uint256 fastLaneLengthSlots) returns()
func (_HashConsensus *HashConsensusTransactorSession) SetFrameConfig(epochsPerFrame *big.Int, fastLaneLengthSlots *big.Int) (*types.Transaction, error) {
	return _HashConsensus.Contract.SetFrameConfig(&_HashConsensus.TransactOpts, epochsPerFrame, fastLaneLengthSlots)
}

// SetQuorum is a paid mutator transaction binding the contract method 0xc1ba4e59.
//
// Solidity: function setQuorum(uint256 quorum) returns()
func (_HashConsensus *HashConsensusTransactor) SetQuorum(opts *bind.TransactOpts, quorum *big.Int) (*types.Transaction, error) {
	return _HashConsensus.contract.Transact(opts, "setQuorum", quorum)
}

// SetQuorum is a paid mutator transaction binding the contract method 0xc1ba4e59.
//
// Solidity: function setQuorum(uint256 quorum) returns()
func (_HashConsensus *HashConsensusSession) SetQuorum(quorum *big.Int) (*types.Transaction, error) {
	return _HashConsensus.Contract.SetQuorum(&_HashConsensus.TransactOpts, quorum)
}

// SetQuorum is a paid mutator transaction binding the contract method 0xc1ba4e59.
//
// Solidity: function setQuorum(uint256 quorum) returns()
func (_HashConsensus *HashConsensusTransactorSession) SetQuorum(quorum *big.Int) (*types.Transaction, error) {
	return _HashConsensus.Contract.SetQuorum(&_HashConsensus.TransactOpts, quorum)
}

// SubmitReport is a paid mutator transaction binding the contract method 0x56e2ca0a.
//
// Solidity: function submitReport(uint256 slot, bytes32[] report) returns()
func (_HashConsensus *HashConsensusTransactor) SubmitReport(opts *bind.TransactOpts, slot *big.Int, report [][32]byte) (*types.Transaction, error) {
	return _HashConsensus.contract.Transact(opts, "submitReport", slot, report)
}

// SubmitReport is a paid mutator transaction binding the contract method 0x56e2ca0a.
//
// Solidity: function submitReport(uint256 slot, bytes32[] report) returns()
func (_HashConsensus *HashConsensusSession) SubmitReport(slot *big.Int, report [][32]byte) (*types.Transaction, error) {
	return _HashConsensus.Contract.SubmitReport(&_HashConsensus.TransactOpts, slot, report)
}

// SubmitReport is a paid mutator transaction binding the contract method 0x56e2ca0a.
//
// Solidity: function submitReport(uint256 slot, bytes32[] report) returns()
func (_HashConsensus *HashConsensusTransactorSession) SubmitReport(slot *big.Int, report [][32]byte) (*types.Transaction, error) {
	return _HashConsensus.Contract.SubmitReport(&_HashConsensus.TransactOpts, slot, report)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HashConsensus *HashConsensusTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HashConsensus.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HashConsensus *HashConsensusSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HashConsensus.Contract.TransferOwnership(&_HashConsensus.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HashConsensus *HashConsensusTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HashConsensus.Contract.TransferOwnership(&_HashConsensus.TransactOpts, newOwner)
}

// UpdateInitialEpoch is a paid mutator transaction binding the contract method 0x323a41f6.
//
// Solidity: function updateInitialEpoch(uint256 initialEpoch) returns()
func (_HashConsensus *HashConsensusTransactor) UpdateInitialEpoch(opts *bind.TransactOpts, initialEpoch *big.Int) (*types.Transaction, error) {
	return _HashConsensus.contract.Transact(opts, "updateInitialEpoch", initialEpoch)
}

// UpdateInitialEpoch is a paid mutator transaction binding the contract method 0x323a41f6.
//
// Solidity: function updateInitialEpoch(uint256 initialEpoch) returns()
func (_HashConsensus *HashConsensusSession) UpdateInitialEpoch(initialEpoch *big.Int) (*types.Transaction, error) {
	return _HashConsensus.Contract.UpdateInitialEpoch(&_HashConsensus.TransactOpts, initialEpoch)
}

// UpdateInitialEpoch is a paid mutator transaction binding the contract method 0x323a41f6.
//
// Solidity: function updateInitialEpoch(uint256 initialEpoch) returns()
func (_HashConsensus *HashConsensusTransactorSession) UpdateInitialEpoch(initialEpoch *big.Int) (*types.Transaction, error) {
	return _HashConsensus.Contract.UpdateInitialEpoch(&_HashConsensus.TransactOpts, initialEpoch)
}

// UpdateReportProcessor is a paid mutator transaction binding the contract method 0x6ed70c43.
//
// Solidity: function updateReportProcessor(address oldProcessor, address newProcessor) returns()
func (_HashConsensus *HashConsensusTransactor) UpdateReportProcessor(opts *bind.TransactOpts, oldProcessor common.Address, newProcessor common.Address) (*types.Transaction, error) {
	return _HashConsensus.contract.Transact(opts, "updateReportProcessor", oldProcessor, newProcessor)
}

// UpdateReportProcessor is a paid mutator transaction binding the contract method 0x6ed70c43.
//
// Solidity: function updateReportProcessor(address oldProcessor, address newProcessor) returns()
func (_HashConsensus *HashConsensusSession) UpdateReportProcessor(oldProcessor common.Address, newProcessor common.Address) (*types.Transaction, error) {
	return _HashConsensus.Contract.UpdateReportProcessor(&_HashConsensus.TransactOpts, oldProcessor, newProcessor)
}

// UpdateReportProcessor is a paid mutator transaction binding the contract method 0x6ed70c43.
//
// Solidity: function updateReportProcessor(address oldProcessor, address newProcessor) returns()
func (_HashConsensus *HashConsensusTransactorSession) UpdateReportProcessor(oldProcessor common.Address, newProcessor common.Address) (*types.Transaction, error) {
	return _HashConsensus.Contract.UpdateReportProcessor(&_HashConsensus.TransactOpts, oldProcessor, newProcessor)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_HashConsensus *HashConsensusTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _HashConsensus.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_HashConsensus *HashConsensusSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _HashConsensus.Contract.UpgradeTo(&_HashConsensus.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_HashConsensus *HashConsensusTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _HashConsensus.Contract.UpgradeTo(&_HashConsensus.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_HashConsensus *HashConsensusTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _HashConsensus.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_HashConsensus *HashConsensusSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _HashConsensus.Contract.UpgradeToAndCall(&_HashConsensus.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_HashConsensus *HashConsensusTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _HashConsensus.Contract.UpgradeToAndCall(&_HashConsensus.TransactOpts, newImplementation, data)
}

// HashConsensusAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the HashConsensus contract.
type HashConsensusAdminChangedIterator struct {
	Event *HashConsensusAdminChanged // Event containing the contract specifics and raw log

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
func (it *HashConsensusAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashConsensusAdminChanged)
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
		it.Event = new(HashConsensusAdminChanged)
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
func (it *HashConsensusAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashConsensusAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashConsensusAdminChanged represents a AdminChanged event raised by the HashConsensus contract.
type HashConsensusAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_HashConsensus *HashConsensusFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*HashConsensusAdminChangedIterator, error) {

	logs, sub, err := _HashConsensus.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &HashConsensusAdminChangedIterator{contract: _HashConsensus.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_HashConsensus *HashConsensusFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *HashConsensusAdminChanged) (event.Subscription, error) {

	logs, sub, err := _HashConsensus.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashConsensusAdminChanged)
				if err := _HashConsensus.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_HashConsensus *HashConsensusFilterer) ParseAdminChanged(log types.Log) (*HashConsensusAdminChanged, error) {
	event := new(HashConsensusAdminChanged)
	if err := _HashConsensus.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashConsensusBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the HashConsensus contract.
type HashConsensusBeaconUpgradedIterator struct {
	Event *HashConsensusBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *HashConsensusBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashConsensusBeaconUpgraded)
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
		it.Event = new(HashConsensusBeaconUpgraded)
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
func (it *HashConsensusBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashConsensusBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashConsensusBeaconUpgraded represents a BeaconUpgraded event raised by the HashConsensus contract.
type HashConsensusBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_HashConsensus *HashConsensusFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*HashConsensusBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _HashConsensus.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &HashConsensusBeaconUpgradedIterator{contract: _HashConsensus.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_HashConsensus *HashConsensusFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *HashConsensusBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _HashConsensus.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashConsensusBeaconUpgraded)
				if err := _HashConsensus.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_HashConsensus *HashConsensusFilterer) ParseBeaconUpgraded(log types.Log) (*HashConsensusBeaconUpgraded, error) {
	event := new(HashConsensusBeaconUpgraded)
	if err := _HashConsensus.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashConsensusConsensusReachedIterator is returned from FilterConsensusReached and is used to iterate over the raw logs and unpacked data for ConsensusReached events raised by the HashConsensus contract.
type HashConsensusConsensusReachedIterator struct {
	Event *HashConsensusConsensusReached // Event containing the contract specifics and raw log

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
func (it *HashConsensusConsensusReachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashConsensusConsensusReached)
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
		it.Event = new(HashConsensusConsensusReached)
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
func (it *HashConsensusConsensusReachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashConsensusConsensusReachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashConsensusConsensusReached represents a ConsensusReached event raised by the HashConsensus contract.
type HashConsensusConsensusReached struct {
	RefSlot *big.Int
	Report  [][32]byte
	Support *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConsensusReached is a free log retrieval operation binding the contract event 0xbc6c822d332d57d2feaf07942537eceae55d981169f087793da8adae1e38466b.
//
// Solidity: event ConsensusReached(uint256 indexed refSlot, bytes32[] report, uint256 support)
func (_HashConsensus *HashConsensusFilterer) FilterConsensusReached(opts *bind.FilterOpts, refSlot []*big.Int) (*HashConsensusConsensusReachedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _HashConsensus.contract.FilterLogs(opts, "ConsensusReached", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &HashConsensusConsensusReachedIterator{contract: _HashConsensus.contract, event: "ConsensusReached", logs: logs, sub: sub}, nil
}

// WatchConsensusReached is a free log subscription operation binding the contract event 0xbc6c822d332d57d2feaf07942537eceae55d981169f087793da8adae1e38466b.
//
// Solidity: event ConsensusReached(uint256 indexed refSlot, bytes32[] report, uint256 support)
func (_HashConsensus *HashConsensusFilterer) WatchConsensusReached(opts *bind.WatchOpts, sink chan<- *HashConsensusConsensusReached, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _HashConsensus.contract.WatchLogs(opts, "ConsensusReached", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashConsensusConsensusReached)
				if err := _HashConsensus.contract.UnpackLog(event, "ConsensusReached", log); err != nil {
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

// ParseConsensusReached is a log parse operation binding the contract event 0xbc6c822d332d57d2feaf07942537eceae55d981169f087793da8adae1e38466b.
//
// Solidity: event ConsensusReached(uint256 indexed refSlot, bytes32[] report, uint256 support)
func (_HashConsensus *HashConsensusFilterer) ParseConsensusReached(log types.Log) (*HashConsensusConsensusReached, error) {
	event := new(HashConsensusConsensusReached)
	if err := _HashConsensus.contract.UnpackLog(event, "ConsensusReached", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashConsensusDaoAddressChangedIterator is returned from FilterDaoAddressChanged and is used to iterate over the raw logs and unpacked data for DaoAddressChanged events raised by the HashConsensus contract.
type HashConsensusDaoAddressChangedIterator struct {
	Event *HashConsensusDaoAddressChanged // Event containing the contract specifics and raw log

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
func (it *HashConsensusDaoAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashConsensusDaoAddressChanged)
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
		it.Event = new(HashConsensusDaoAddressChanged)
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
func (it *HashConsensusDaoAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashConsensusDaoAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashConsensusDaoAddressChanged represents a DaoAddressChanged event raised by the HashConsensus contract.
type HashConsensusDaoAddressChanged struct {
	OldDao common.Address
	Dao    common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDaoAddressChanged is a free log retrieval operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address _oldDao, address _dao)
func (_HashConsensus *HashConsensusFilterer) FilterDaoAddressChanged(opts *bind.FilterOpts) (*HashConsensusDaoAddressChangedIterator, error) {

	logs, sub, err := _HashConsensus.contract.FilterLogs(opts, "DaoAddressChanged")
	if err != nil {
		return nil, err
	}
	return &HashConsensusDaoAddressChangedIterator{contract: _HashConsensus.contract, event: "DaoAddressChanged", logs: logs, sub: sub}, nil
}

// WatchDaoAddressChanged is a free log subscription operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address _oldDao, address _dao)
func (_HashConsensus *HashConsensusFilterer) WatchDaoAddressChanged(opts *bind.WatchOpts, sink chan<- *HashConsensusDaoAddressChanged) (event.Subscription, error) {

	logs, sub, err := _HashConsensus.contract.WatchLogs(opts, "DaoAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashConsensusDaoAddressChanged)
				if err := _HashConsensus.contract.UnpackLog(event, "DaoAddressChanged", log); err != nil {
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
func (_HashConsensus *HashConsensusFilterer) ParseDaoAddressChanged(log types.Log) (*HashConsensusDaoAddressChanged, error) {
	event := new(HashConsensusDaoAddressChanged)
	if err := _HashConsensus.contract.UnpackLog(event, "DaoAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashConsensusFastLaneConfigSetIterator is returned from FilterFastLaneConfigSet and is used to iterate over the raw logs and unpacked data for FastLaneConfigSet events raised by the HashConsensus contract.
type HashConsensusFastLaneConfigSetIterator struct {
	Event *HashConsensusFastLaneConfigSet // Event containing the contract specifics and raw log

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
func (it *HashConsensusFastLaneConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashConsensusFastLaneConfigSet)
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
		it.Event = new(HashConsensusFastLaneConfigSet)
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
func (it *HashConsensusFastLaneConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashConsensusFastLaneConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashConsensusFastLaneConfigSet represents a FastLaneConfigSet event raised by the HashConsensus contract.
type HashConsensusFastLaneConfigSet struct {
	FastLaneLengthSlots *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterFastLaneConfigSet is a free log retrieval operation binding the contract event 0xab8b22776606cc75c47792d32af7e63ed9ca74e85c9780a7fc7994fdbd6fde2b.
//
// Solidity: event FastLaneConfigSet(uint256 fastLaneLengthSlots)
func (_HashConsensus *HashConsensusFilterer) FilterFastLaneConfigSet(opts *bind.FilterOpts) (*HashConsensusFastLaneConfigSetIterator, error) {

	logs, sub, err := _HashConsensus.contract.FilterLogs(opts, "FastLaneConfigSet")
	if err != nil {
		return nil, err
	}
	return &HashConsensusFastLaneConfigSetIterator{contract: _HashConsensus.contract, event: "FastLaneConfigSet", logs: logs, sub: sub}, nil
}

// WatchFastLaneConfigSet is a free log subscription operation binding the contract event 0xab8b22776606cc75c47792d32af7e63ed9ca74e85c9780a7fc7994fdbd6fde2b.
//
// Solidity: event FastLaneConfigSet(uint256 fastLaneLengthSlots)
func (_HashConsensus *HashConsensusFilterer) WatchFastLaneConfigSet(opts *bind.WatchOpts, sink chan<- *HashConsensusFastLaneConfigSet) (event.Subscription, error) {

	logs, sub, err := _HashConsensus.contract.WatchLogs(opts, "FastLaneConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashConsensusFastLaneConfigSet)
				if err := _HashConsensus.contract.UnpackLog(event, "FastLaneConfigSet", log); err != nil {
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
func (_HashConsensus *HashConsensusFilterer) ParseFastLaneConfigSet(log types.Log) (*HashConsensusFastLaneConfigSet, error) {
	event := new(HashConsensusFastLaneConfigSet)
	if err := _HashConsensus.contract.UnpackLog(event, "FastLaneConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashConsensusFrameConfigSetIterator is returned from FilterFrameConfigSet and is used to iterate over the raw logs and unpacked data for FrameConfigSet events raised by the HashConsensus contract.
type HashConsensusFrameConfigSetIterator struct {
	Event *HashConsensusFrameConfigSet // Event containing the contract specifics and raw log

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
func (it *HashConsensusFrameConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashConsensusFrameConfigSet)
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
		it.Event = new(HashConsensusFrameConfigSet)
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
func (it *HashConsensusFrameConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashConsensusFrameConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashConsensusFrameConfigSet represents a FrameConfigSet event raised by the HashConsensus contract.
type HashConsensusFrameConfigSet struct {
	NewInitialEpoch   *big.Int
	NewEpochsPerFrame *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFrameConfigSet is a free log retrieval operation binding the contract event 0xe343afa5219eaf28c50ce9cd658acd69cbe28b34fa773eb3a523e28007f64afc.
//
// Solidity: event FrameConfigSet(uint256 newInitialEpoch, uint256 newEpochsPerFrame)
func (_HashConsensus *HashConsensusFilterer) FilterFrameConfigSet(opts *bind.FilterOpts) (*HashConsensusFrameConfigSetIterator, error) {

	logs, sub, err := _HashConsensus.contract.FilterLogs(opts, "FrameConfigSet")
	if err != nil {
		return nil, err
	}
	return &HashConsensusFrameConfigSetIterator{contract: _HashConsensus.contract, event: "FrameConfigSet", logs: logs, sub: sub}, nil
}

// WatchFrameConfigSet is a free log subscription operation binding the contract event 0xe343afa5219eaf28c50ce9cd658acd69cbe28b34fa773eb3a523e28007f64afc.
//
// Solidity: event FrameConfigSet(uint256 newInitialEpoch, uint256 newEpochsPerFrame)
func (_HashConsensus *HashConsensusFilterer) WatchFrameConfigSet(opts *bind.WatchOpts, sink chan<- *HashConsensusFrameConfigSet) (event.Subscription, error) {

	logs, sub, err := _HashConsensus.contract.WatchLogs(opts, "FrameConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashConsensusFrameConfigSet)
				if err := _HashConsensus.contract.UnpackLog(event, "FrameConfigSet", log); err != nil {
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
func (_HashConsensus *HashConsensusFilterer) ParseFrameConfigSet(log types.Log) (*HashConsensusFrameConfigSet, error) {
	event := new(HashConsensusFrameConfigSet)
	if err := _HashConsensus.contract.UnpackLog(event, "FrameConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashConsensusInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the HashConsensus contract.
type HashConsensusInitializedIterator struct {
	Event *HashConsensusInitialized // Event containing the contract specifics and raw log

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
func (it *HashConsensusInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashConsensusInitialized)
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
		it.Event = new(HashConsensusInitialized)
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
func (it *HashConsensusInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashConsensusInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashConsensusInitialized represents a Initialized event raised by the HashConsensus contract.
type HashConsensusInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HashConsensus *HashConsensusFilterer) FilterInitialized(opts *bind.FilterOpts) (*HashConsensusInitializedIterator, error) {

	logs, sub, err := _HashConsensus.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &HashConsensusInitializedIterator{contract: _HashConsensus.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HashConsensus *HashConsensusFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *HashConsensusInitialized) (event.Subscription, error) {

	logs, sub, err := _HashConsensus.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashConsensusInitialized)
				if err := _HashConsensus.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_HashConsensus *HashConsensusFilterer) ParseInitialized(log types.Log) (*HashConsensusInitialized, error) {
	event := new(HashConsensusInitialized)
	if err := _HashConsensus.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashConsensusMemberAddedIterator is returned from FilterMemberAdded and is used to iterate over the raw logs and unpacked data for MemberAdded events raised by the HashConsensus contract.
type HashConsensusMemberAddedIterator struct {
	Event *HashConsensusMemberAdded // Event containing the contract specifics and raw log

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
func (it *HashConsensusMemberAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashConsensusMemberAdded)
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
		it.Event = new(HashConsensusMemberAdded)
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
func (it *HashConsensusMemberAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashConsensusMemberAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashConsensusMemberAdded represents a MemberAdded event raised by the HashConsensus contract.
type HashConsensusMemberAdded struct {
	Addr            common.Address
	NewTotalMembers *big.Int
	NewQuorum       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMemberAdded is a free log retrieval operation binding the contract event 0xe17e0e2cd88e2144dd54f3d823c30d4569092bcac1aabaec1129883e9cc12d2e.
//
// Solidity: event MemberAdded(address indexed addr, uint256 newTotalMembers, uint256 newQuorum)
func (_HashConsensus *HashConsensusFilterer) FilterMemberAdded(opts *bind.FilterOpts, addr []common.Address) (*HashConsensusMemberAddedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _HashConsensus.contract.FilterLogs(opts, "MemberAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return &HashConsensusMemberAddedIterator{contract: _HashConsensus.contract, event: "MemberAdded", logs: logs, sub: sub}, nil
}

// WatchMemberAdded is a free log subscription operation binding the contract event 0xe17e0e2cd88e2144dd54f3d823c30d4569092bcac1aabaec1129883e9cc12d2e.
//
// Solidity: event MemberAdded(address indexed addr, uint256 newTotalMembers, uint256 newQuorum)
func (_HashConsensus *HashConsensusFilterer) WatchMemberAdded(opts *bind.WatchOpts, sink chan<- *HashConsensusMemberAdded, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _HashConsensus.contract.WatchLogs(opts, "MemberAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashConsensusMemberAdded)
				if err := _HashConsensus.contract.UnpackLog(event, "MemberAdded", log); err != nil {
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
func (_HashConsensus *HashConsensusFilterer) ParseMemberAdded(log types.Log) (*HashConsensusMemberAdded, error) {
	event := new(HashConsensusMemberAdded)
	if err := _HashConsensus.contract.UnpackLog(event, "MemberAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashConsensusMemberRemovedIterator is returned from FilterMemberRemoved and is used to iterate over the raw logs and unpacked data for MemberRemoved events raised by the HashConsensus contract.
type HashConsensusMemberRemovedIterator struct {
	Event *HashConsensusMemberRemoved // Event containing the contract specifics and raw log

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
func (it *HashConsensusMemberRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashConsensusMemberRemoved)
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
		it.Event = new(HashConsensusMemberRemoved)
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
func (it *HashConsensusMemberRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashConsensusMemberRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashConsensusMemberRemoved represents a MemberRemoved event raised by the HashConsensus contract.
type HashConsensusMemberRemoved struct {
	Addr            common.Address
	NewTotalMembers *big.Int
	NewQuorum       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMemberRemoved is a free log retrieval operation binding the contract event 0xa182730913550d27dc6c5813fad297cb0785871bec3d0152c5650e59c5d39d60.
//
// Solidity: event MemberRemoved(address indexed addr, uint256 newTotalMembers, uint256 newQuorum)
func (_HashConsensus *HashConsensusFilterer) FilterMemberRemoved(opts *bind.FilterOpts, addr []common.Address) (*HashConsensusMemberRemovedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _HashConsensus.contract.FilterLogs(opts, "MemberRemoved", addrRule)
	if err != nil {
		return nil, err
	}
	return &HashConsensusMemberRemovedIterator{contract: _HashConsensus.contract, event: "MemberRemoved", logs: logs, sub: sub}, nil
}

// WatchMemberRemoved is a free log subscription operation binding the contract event 0xa182730913550d27dc6c5813fad297cb0785871bec3d0152c5650e59c5d39d60.
//
// Solidity: event MemberRemoved(address indexed addr, uint256 newTotalMembers, uint256 newQuorum)
func (_HashConsensus *HashConsensusFilterer) WatchMemberRemoved(opts *bind.WatchOpts, sink chan<- *HashConsensusMemberRemoved, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _HashConsensus.contract.WatchLogs(opts, "MemberRemoved", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashConsensusMemberRemoved)
				if err := _HashConsensus.contract.UnpackLog(event, "MemberRemoved", log); err != nil {
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
func (_HashConsensus *HashConsensusFilterer) ParseMemberRemoved(log types.Log) (*HashConsensusMemberRemoved, error) {
	event := new(HashConsensusMemberRemoved)
	if err := _HashConsensus.contract.UnpackLog(event, "MemberRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashConsensusOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HashConsensus contract.
type HashConsensusOwnershipTransferredIterator struct {
	Event *HashConsensusOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HashConsensusOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashConsensusOwnershipTransferred)
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
		it.Event = new(HashConsensusOwnershipTransferred)
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
func (it *HashConsensusOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashConsensusOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashConsensusOwnershipTransferred represents a OwnershipTransferred event raised by the HashConsensus contract.
type HashConsensusOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HashConsensus *HashConsensusFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HashConsensusOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HashConsensus.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HashConsensusOwnershipTransferredIterator{contract: _HashConsensus.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HashConsensus *HashConsensusFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HashConsensusOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HashConsensus.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashConsensusOwnershipTransferred)
				if err := _HashConsensus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_HashConsensus *HashConsensusFilterer) ParseOwnershipTransferred(log types.Log) (*HashConsensusOwnershipTransferred, error) {
	event := new(HashConsensusOwnershipTransferred)
	if err := _HashConsensus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashConsensusQuorumSetIterator is returned from FilterQuorumSet and is used to iterate over the raw logs and unpacked data for QuorumSet events raised by the HashConsensus contract.
type HashConsensusQuorumSetIterator struct {
	Event *HashConsensusQuorumSet // Event containing the contract specifics and raw log

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
func (it *HashConsensusQuorumSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashConsensusQuorumSet)
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
		it.Event = new(HashConsensusQuorumSet)
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
func (it *HashConsensusQuorumSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashConsensusQuorumSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashConsensusQuorumSet represents a QuorumSet event raised by the HashConsensus contract.
type HashConsensusQuorumSet struct {
	NewQuorum    *big.Int
	TotalMembers *big.Int
	PrevQuorum   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuorumSet is a free log retrieval operation binding the contract event 0x9f40cfd22fe91777c78f252bd21a710f3fb007dc2f321876891e7644ba0ae175.
//
// Solidity: event QuorumSet(uint256 newQuorum, uint256 totalMembers, uint256 prevQuorum)
func (_HashConsensus *HashConsensusFilterer) FilterQuorumSet(opts *bind.FilterOpts) (*HashConsensusQuorumSetIterator, error) {

	logs, sub, err := _HashConsensus.contract.FilterLogs(opts, "QuorumSet")
	if err != nil {
		return nil, err
	}
	return &HashConsensusQuorumSetIterator{contract: _HashConsensus.contract, event: "QuorumSet", logs: logs, sub: sub}, nil
}

// WatchQuorumSet is a free log subscription operation binding the contract event 0x9f40cfd22fe91777c78f252bd21a710f3fb007dc2f321876891e7644ba0ae175.
//
// Solidity: event QuorumSet(uint256 newQuorum, uint256 totalMembers, uint256 prevQuorum)
func (_HashConsensus *HashConsensusFilterer) WatchQuorumSet(opts *bind.WatchOpts, sink chan<- *HashConsensusQuorumSet) (event.Subscription, error) {

	logs, sub, err := _HashConsensus.contract.WatchLogs(opts, "QuorumSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashConsensusQuorumSet)
				if err := _HashConsensus.contract.UnpackLog(event, "QuorumSet", log); err != nil {
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
func (_HashConsensus *HashConsensusFilterer) ParseQuorumSet(log types.Log) (*HashConsensusQuorumSet, error) {
	event := new(HashConsensusQuorumSet)
	if err := _HashConsensus.contract.UnpackLog(event, "QuorumSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashConsensusReportProcessorAddIterator is returned from FilterReportProcessorAdd and is used to iterate over the raw logs and unpacked data for ReportProcessorAdd events raised by the HashConsensus contract.
type HashConsensusReportProcessorAddIterator struct {
	Event *HashConsensusReportProcessorAdd // Event containing the contract specifics and raw log

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
func (it *HashConsensusReportProcessorAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashConsensusReportProcessorAdd)
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
		it.Event = new(HashConsensusReportProcessorAdd)
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
func (it *HashConsensusReportProcessorAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashConsensusReportProcessorAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashConsensusReportProcessorAdd represents a ReportProcessorAdd event raised by the HashConsensus contract.
type HashConsensusReportProcessorAdd struct {
	Processor common.Address
	ModuleId  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReportProcessorAdd is a free log retrieval operation binding the contract event 0xef288ced2ba315becacd439f4f3e8d2efa3805f608270e99011e309affd946fb.
//
// Solidity: event ReportProcessorAdd(address indexed processor, uint256 indexed moduleId)
func (_HashConsensus *HashConsensusFilterer) FilterReportProcessorAdd(opts *bind.FilterOpts, processor []common.Address, moduleId []*big.Int) (*HashConsensusReportProcessorAddIterator, error) {

	var processorRule []interface{}
	for _, processorItem := range processor {
		processorRule = append(processorRule, processorItem)
	}
	var moduleIdRule []interface{}
	for _, moduleIdItem := range moduleId {
		moduleIdRule = append(moduleIdRule, moduleIdItem)
	}

	logs, sub, err := _HashConsensus.contract.FilterLogs(opts, "ReportProcessorAdd", processorRule, moduleIdRule)
	if err != nil {
		return nil, err
	}
	return &HashConsensusReportProcessorAddIterator{contract: _HashConsensus.contract, event: "ReportProcessorAdd", logs: logs, sub: sub}, nil
}

// WatchReportProcessorAdd is a free log subscription operation binding the contract event 0xef288ced2ba315becacd439f4f3e8d2efa3805f608270e99011e309affd946fb.
//
// Solidity: event ReportProcessorAdd(address indexed processor, uint256 indexed moduleId)
func (_HashConsensus *HashConsensusFilterer) WatchReportProcessorAdd(opts *bind.WatchOpts, sink chan<- *HashConsensusReportProcessorAdd, processor []common.Address, moduleId []*big.Int) (event.Subscription, error) {

	var processorRule []interface{}
	for _, processorItem := range processor {
		processorRule = append(processorRule, processorItem)
	}
	var moduleIdRule []interface{}
	for _, moduleIdItem := range moduleId {
		moduleIdRule = append(moduleIdRule, moduleIdItem)
	}

	logs, sub, err := _HashConsensus.contract.WatchLogs(opts, "ReportProcessorAdd", processorRule, moduleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashConsensusReportProcessorAdd)
				if err := _HashConsensus.contract.UnpackLog(event, "ReportProcessorAdd", log); err != nil {
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

// ParseReportProcessorAdd is a log parse operation binding the contract event 0xef288ced2ba315becacd439f4f3e8d2efa3805f608270e99011e309affd946fb.
//
// Solidity: event ReportProcessorAdd(address indexed processor, uint256 indexed moduleId)
func (_HashConsensus *HashConsensusFilterer) ParseReportProcessorAdd(log types.Log) (*HashConsensusReportProcessorAdd, error) {
	event := new(HashConsensusReportProcessorAdd)
	if err := _HashConsensus.contract.UnpackLog(event, "ReportProcessorAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashConsensusReportProcessorUpdateIterator is returned from FilterReportProcessorUpdate and is used to iterate over the raw logs and unpacked data for ReportProcessorUpdate events raised by the HashConsensus contract.
type HashConsensusReportProcessorUpdateIterator struct {
	Event *HashConsensusReportProcessorUpdate // Event containing the contract specifics and raw log

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
func (it *HashConsensusReportProcessorUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashConsensusReportProcessorUpdate)
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
		it.Event = new(HashConsensusReportProcessorUpdate)
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
func (it *HashConsensusReportProcessorUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashConsensusReportProcessorUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashConsensusReportProcessorUpdate represents a ReportProcessorUpdate event raised by the HashConsensus contract.
type HashConsensusReportProcessorUpdate struct {
	OldProcessor common.Address
	NewProcessor common.Address
	ModuleId     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterReportProcessorUpdate is a free log retrieval operation binding the contract event 0xff55bf1635066aaf80ee2cf988e114a888d58206217a33cf9b798c01d1f89790.
//
// Solidity: event ReportProcessorUpdate(address indexed oldProcessor, address indexed newProcessor, uint256 indexed moduleId)
func (_HashConsensus *HashConsensusFilterer) FilterReportProcessorUpdate(opts *bind.FilterOpts, oldProcessor []common.Address, newProcessor []common.Address, moduleId []*big.Int) (*HashConsensusReportProcessorUpdateIterator, error) {

	var oldProcessorRule []interface{}
	for _, oldProcessorItem := range oldProcessor {
		oldProcessorRule = append(oldProcessorRule, oldProcessorItem)
	}
	var newProcessorRule []interface{}
	for _, newProcessorItem := range newProcessor {
		newProcessorRule = append(newProcessorRule, newProcessorItem)
	}
	var moduleIdRule []interface{}
	for _, moduleIdItem := range moduleId {
		moduleIdRule = append(moduleIdRule, moduleIdItem)
	}

	logs, sub, err := _HashConsensus.contract.FilterLogs(opts, "ReportProcessorUpdate", oldProcessorRule, newProcessorRule, moduleIdRule)
	if err != nil {
		return nil, err
	}
	return &HashConsensusReportProcessorUpdateIterator{contract: _HashConsensus.contract, event: "ReportProcessorUpdate", logs: logs, sub: sub}, nil
}

// WatchReportProcessorUpdate is a free log subscription operation binding the contract event 0xff55bf1635066aaf80ee2cf988e114a888d58206217a33cf9b798c01d1f89790.
//
// Solidity: event ReportProcessorUpdate(address indexed oldProcessor, address indexed newProcessor, uint256 indexed moduleId)
func (_HashConsensus *HashConsensusFilterer) WatchReportProcessorUpdate(opts *bind.WatchOpts, sink chan<- *HashConsensusReportProcessorUpdate, oldProcessor []common.Address, newProcessor []common.Address, moduleId []*big.Int) (event.Subscription, error) {

	var oldProcessorRule []interface{}
	for _, oldProcessorItem := range oldProcessor {
		oldProcessorRule = append(oldProcessorRule, oldProcessorItem)
	}
	var newProcessorRule []interface{}
	for _, newProcessorItem := range newProcessor {
		newProcessorRule = append(newProcessorRule, newProcessorItem)
	}
	var moduleIdRule []interface{}
	for _, moduleIdItem := range moduleId {
		moduleIdRule = append(moduleIdRule, moduleIdItem)
	}

	logs, sub, err := _HashConsensus.contract.WatchLogs(opts, "ReportProcessorUpdate", oldProcessorRule, newProcessorRule, moduleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashConsensusReportProcessorUpdate)
				if err := _HashConsensus.contract.UnpackLog(event, "ReportProcessorUpdate", log); err != nil {
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

// ParseReportProcessorUpdate is a log parse operation binding the contract event 0xff55bf1635066aaf80ee2cf988e114a888d58206217a33cf9b798c01d1f89790.
//
// Solidity: event ReportProcessorUpdate(address indexed oldProcessor, address indexed newProcessor, uint256 indexed moduleId)
func (_HashConsensus *HashConsensusFilterer) ParseReportProcessorUpdate(log types.Log) (*HashConsensusReportProcessorUpdate, error) {
	event := new(HashConsensusReportProcessorUpdate)
	if err := _HashConsensus.contract.UnpackLog(event, "ReportProcessorUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashConsensusReportReceivedIterator is returned from FilterReportReceived and is used to iterate over the raw logs and unpacked data for ReportReceived events raised by the HashConsensus contract.
type HashConsensusReportReceivedIterator struct {
	Event *HashConsensusReportReceived // Event containing the contract specifics and raw log

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
func (it *HashConsensusReportReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashConsensusReportReceived)
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
		it.Event = new(HashConsensusReportReceived)
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
func (it *HashConsensusReportReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashConsensusReportReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashConsensusReportReceived represents a ReportReceived event raised by the HashConsensus contract.
type HashConsensusReportReceived struct {
	RefSlot *big.Int
	Member  common.Address
	Report  [][32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReportReceived is a free log retrieval operation binding the contract event 0xd6f9dc41f1a10bb63a22fd32e563b3c0b1b6c79d95c41a404a86f6fc16939a34.
//
// Solidity: event ReportReceived(uint256 indexed refSlot, address indexed member, bytes32[] report)
func (_HashConsensus *HashConsensusFilterer) FilterReportReceived(opts *bind.FilterOpts, refSlot []*big.Int, member []common.Address) (*HashConsensusReportReceivedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}
	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _HashConsensus.contract.FilterLogs(opts, "ReportReceived", refSlotRule, memberRule)
	if err != nil {
		return nil, err
	}
	return &HashConsensusReportReceivedIterator{contract: _HashConsensus.contract, event: "ReportReceived", logs: logs, sub: sub}, nil
}

// WatchReportReceived is a free log subscription operation binding the contract event 0xd6f9dc41f1a10bb63a22fd32e563b3c0b1b6c79d95c41a404a86f6fc16939a34.
//
// Solidity: event ReportReceived(uint256 indexed refSlot, address indexed member, bytes32[] report)
func (_HashConsensus *HashConsensusFilterer) WatchReportReceived(opts *bind.WatchOpts, sink chan<- *HashConsensusReportReceived, refSlot []*big.Int, member []common.Address) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}
	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _HashConsensus.contract.WatchLogs(opts, "ReportReceived", refSlotRule, memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashConsensusReportReceived)
				if err := _HashConsensus.contract.UnpackLog(event, "ReportReceived", log); err != nil {
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

// ParseReportReceived is a log parse operation binding the contract event 0xd6f9dc41f1a10bb63a22fd32e563b3c0b1b6c79d95c41a404a86f6fc16939a34.
//
// Solidity: event ReportReceived(uint256 indexed refSlot, address indexed member, bytes32[] report)
func (_HashConsensus *HashConsensusFilterer) ParseReportReceived(log types.Log) (*HashConsensusReportReceived, error) {
	event := new(HashConsensusReportReceived)
	if err := _HashConsensus.contract.UnpackLog(event, "ReportReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashConsensusUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the HashConsensus contract.
type HashConsensusUpgradedIterator struct {
	Event *HashConsensusUpgraded // Event containing the contract specifics and raw log

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
func (it *HashConsensusUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashConsensusUpgraded)
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
		it.Event = new(HashConsensusUpgraded)
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
func (it *HashConsensusUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashConsensusUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashConsensusUpgraded represents a Upgraded event raised by the HashConsensus contract.
type HashConsensusUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_HashConsensus *HashConsensusFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*HashConsensusUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _HashConsensus.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &HashConsensusUpgradedIterator{contract: _HashConsensus.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_HashConsensus *HashConsensusFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *HashConsensusUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _HashConsensus.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashConsensusUpgraded)
				if err := _HashConsensus.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_HashConsensus *HashConsensusFilterer) ParseUpgraded(log types.Log) (*HashConsensusUpgraded, error) {
	event := new(HashConsensusUpgraded)
	if err := _HashConsensus.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

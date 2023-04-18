// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package operatorSlash

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

// OperatorSlashMetaData contains all meta data concerning the OperatorSlash contract.
var OperatorSlashMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ExcessivePenaltyAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSlashNeeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ArrearsReceiveOfSlash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalCompensated\",\"type\":\"uint256\"}],\"name\":\"CompensatedClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldDao\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"DaoAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldLiquidStakingContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_liquidStakingContractAddress\",\"type\":\"address\"}],\"name\":\"LiquidStakingChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldNodeOperatorRegistryContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistryContract\",\"type\":\"address\"}],\"name\":\"NodeOperatorRegistryContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldSlashAmountPerBlockPerValidator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_slashAmountPerBlockPerValidator\",\"type\":\"uint256\"}],\"name\":\"SlashAmountPerBlockPerValidatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_slashAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_requirAmounts\",\"type\":\"uint256\"}],\"name\":\"SlashReceive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldVaultManagerContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_vaultManagerContract\",\"type\":\"address\"}],\"name\":\"VaultManagerContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldWithdrawalRequestContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_withdrawalRequestContractAddress\",\"type\":\"address\"}],\"name\":\"WithdrawalRequestContractSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"claimCompensated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedExitSlashStandard\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidStakingAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nVNFTContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_withdrawalRequestContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultManagerContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_delayedExitSlashStandard\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"largeExitDelayedSlashRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidStakingContract\",\"outputs\":[{\"internalType\":\"contractILiquidStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftExitDelayedSlashRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftHasCompensated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftWillCompensated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeOperatorRegistryContract\",\"outputs\":[{\"internalType\":\"contractINodeOperatorsRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorCompensatedIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorSlashArrears\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"setDaoAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidStakingContractAddress\",\"type\":\"address\"}],\"name\":\"setLiquidStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistryContract\",\"type\":\"address\"}],\"name\":\"setNodeOperatorRegistryContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashAmountPerBlockPerValidator\",\"type\":\"uint256\"}],\"name\":\"setSlashAmountPerBlockPerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultManagerContract\",\"type\":\"address\"}],\"name\":\"setVaultManagerContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_withdrawalRequestContractAddress\",\"type\":\"address\"}],\"name\":\"setWithdrawalRequestContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashAmountPerBlockPerValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"slashArrearsReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_nftExitDelayedTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_largeExitDelayedRequestIds\",\"type\":\"uint256[]\"}],\"name\":\"slashOfExitDelayed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_exitTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"slashOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_exitTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_slashAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_requireAmounts\",\"type\":\"uint256[]\"}],\"name\":\"slashReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vNFTContract\",\"outputs\":[{\"internalType\":\"contractIVNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultManagerContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalRequestContract\",\"outputs\":[{\"internalType\":\"contractIWithdrawalRequest\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OperatorSlashABI is the input ABI used to generate the binding from.
// Deprecated: Use OperatorSlashMetaData.ABI instead.
var OperatorSlashABI = OperatorSlashMetaData.ABI

// OperatorSlash is an auto generated Go binding around an Ethereum contract.
type OperatorSlash struct {
	OperatorSlashCaller     // Read-only binding to the contract
	OperatorSlashTransactor // Write-only binding to the contract
	OperatorSlashFilterer   // Log filterer for contract events
}

// OperatorSlashCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperatorSlashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorSlashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperatorSlashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorSlashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OperatorSlashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorSlashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperatorSlashSession struct {
	Contract     *OperatorSlash    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OperatorSlashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperatorSlashCallerSession struct {
	Contract *OperatorSlashCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OperatorSlashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperatorSlashTransactorSession struct {
	Contract     *OperatorSlashTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OperatorSlashRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperatorSlashRaw struct {
	Contract *OperatorSlash // Generic contract binding to access the raw methods on
}

// OperatorSlashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperatorSlashCallerRaw struct {
	Contract *OperatorSlashCaller // Generic read-only contract binding to access the raw methods on
}

// OperatorSlashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperatorSlashTransactorRaw struct {
	Contract *OperatorSlashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperatorSlash creates a new instance of OperatorSlash, bound to a specific deployed contract.
func NewOperatorSlash(address common.Address, backend bind.ContractBackend) (*OperatorSlash, error) {
	contract, err := bindOperatorSlash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OperatorSlash{OperatorSlashCaller: OperatorSlashCaller{contract: contract}, OperatorSlashTransactor: OperatorSlashTransactor{contract: contract}, OperatorSlashFilterer: OperatorSlashFilterer{contract: contract}}, nil
}

// NewOperatorSlashCaller creates a new read-only instance of OperatorSlash, bound to a specific deployed contract.
func NewOperatorSlashCaller(address common.Address, caller bind.ContractCaller) (*OperatorSlashCaller, error) {
	contract, err := bindOperatorSlash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorSlashCaller{contract: contract}, nil
}

// NewOperatorSlashTransactor creates a new write-only instance of OperatorSlash, bound to a specific deployed contract.
func NewOperatorSlashTransactor(address common.Address, transactor bind.ContractTransactor) (*OperatorSlashTransactor, error) {
	contract, err := bindOperatorSlash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorSlashTransactor{contract: contract}, nil
}

// NewOperatorSlashFilterer creates a new log filterer instance of OperatorSlash, bound to a specific deployed contract.
func NewOperatorSlashFilterer(address common.Address, filterer bind.ContractFilterer) (*OperatorSlashFilterer, error) {
	contract, err := bindOperatorSlash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OperatorSlashFilterer{contract: contract}, nil
}

// bindOperatorSlash binds a generic wrapper to an already deployed contract.
func bindOperatorSlash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OperatorSlashMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorSlash *OperatorSlashRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorSlash.Contract.OperatorSlashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorSlash *OperatorSlashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorSlash.Contract.OperatorSlashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorSlash *OperatorSlashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorSlash.Contract.OperatorSlashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorSlash *OperatorSlashCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorSlash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorSlash *OperatorSlashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorSlash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorSlash *OperatorSlashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorSlash.Contract.contract.Transact(opts, method, params...)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_OperatorSlash *OperatorSlashCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorSlash.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_OperatorSlash *OperatorSlashSession) Dao() (common.Address, error) {
	return _OperatorSlash.Contract.Dao(&_OperatorSlash.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_OperatorSlash *OperatorSlashCallerSession) Dao() (common.Address, error) {
	return _OperatorSlash.Contract.Dao(&_OperatorSlash.CallOpts)
}

// DelayedExitSlashStandard is a free data retrieval call binding the contract method 0xfa2d7ea5.
//
// Solidity: function delayedExitSlashStandard() view returns(uint256)
func (_OperatorSlash *OperatorSlashCaller) DelayedExitSlashStandard(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OperatorSlash.contract.Call(opts, &out, "delayedExitSlashStandard")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayedExitSlashStandard is a free data retrieval call binding the contract method 0xfa2d7ea5.
//
// Solidity: function delayedExitSlashStandard() view returns(uint256)
func (_OperatorSlash *OperatorSlashSession) DelayedExitSlashStandard() (*big.Int, error) {
	return _OperatorSlash.Contract.DelayedExitSlashStandard(&_OperatorSlash.CallOpts)
}

// DelayedExitSlashStandard is a free data retrieval call binding the contract method 0xfa2d7ea5.
//
// Solidity: function delayedExitSlashStandard() view returns(uint256)
func (_OperatorSlash *OperatorSlashCallerSession) DelayedExitSlashStandard() (*big.Int, error) {
	return _OperatorSlash.Contract.DelayedExitSlashStandard(&_OperatorSlash.CallOpts)
}

// LargeExitDelayedSlashRecords is a free data retrieval call binding the contract method 0x9d4b4fd5.
//
// Solidity: function largeExitDelayedSlashRecords(uint256 ) view returns(uint256)
func (_OperatorSlash *OperatorSlashCaller) LargeExitDelayedSlashRecords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OperatorSlash.contract.Call(opts, &out, "largeExitDelayedSlashRecords", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LargeExitDelayedSlashRecords is a free data retrieval call binding the contract method 0x9d4b4fd5.
//
// Solidity: function largeExitDelayedSlashRecords(uint256 ) view returns(uint256)
func (_OperatorSlash *OperatorSlashSession) LargeExitDelayedSlashRecords(arg0 *big.Int) (*big.Int, error) {
	return _OperatorSlash.Contract.LargeExitDelayedSlashRecords(&_OperatorSlash.CallOpts, arg0)
}

// LargeExitDelayedSlashRecords is a free data retrieval call binding the contract method 0x9d4b4fd5.
//
// Solidity: function largeExitDelayedSlashRecords(uint256 ) view returns(uint256)
func (_OperatorSlash *OperatorSlashCallerSession) LargeExitDelayedSlashRecords(arg0 *big.Int) (*big.Int, error) {
	return _OperatorSlash.Contract.LargeExitDelayedSlashRecords(&_OperatorSlash.CallOpts, arg0)
}

// LiquidStakingContract is a free data retrieval call binding the contract method 0xbdcaa355.
//
// Solidity: function liquidStakingContract() view returns(address)
func (_OperatorSlash *OperatorSlashCaller) LiquidStakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorSlash.contract.Call(opts, &out, "liquidStakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidStakingContract is a free data retrieval call binding the contract method 0xbdcaa355.
//
// Solidity: function liquidStakingContract() view returns(address)
func (_OperatorSlash *OperatorSlashSession) LiquidStakingContract() (common.Address, error) {
	return _OperatorSlash.Contract.LiquidStakingContract(&_OperatorSlash.CallOpts)
}

// LiquidStakingContract is a free data retrieval call binding the contract method 0xbdcaa355.
//
// Solidity: function liquidStakingContract() view returns(address)
func (_OperatorSlash *OperatorSlashCallerSession) LiquidStakingContract() (common.Address, error) {
	return _OperatorSlash.Contract.LiquidStakingContract(&_OperatorSlash.CallOpts)
}

// NftExitDelayedSlashRecords is a free data retrieval call binding the contract method 0x1d3fb6bd.
//
// Solidity: function nftExitDelayedSlashRecords(uint256 ) view returns(uint256)
func (_OperatorSlash *OperatorSlashCaller) NftExitDelayedSlashRecords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OperatorSlash.contract.Call(opts, &out, "nftExitDelayedSlashRecords", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftExitDelayedSlashRecords is a free data retrieval call binding the contract method 0x1d3fb6bd.
//
// Solidity: function nftExitDelayedSlashRecords(uint256 ) view returns(uint256)
func (_OperatorSlash *OperatorSlashSession) NftExitDelayedSlashRecords(arg0 *big.Int) (*big.Int, error) {
	return _OperatorSlash.Contract.NftExitDelayedSlashRecords(&_OperatorSlash.CallOpts, arg0)
}

// NftExitDelayedSlashRecords is a free data retrieval call binding the contract method 0x1d3fb6bd.
//
// Solidity: function nftExitDelayedSlashRecords(uint256 ) view returns(uint256)
func (_OperatorSlash *OperatorSlashCallerSession) NftExitDelayedSlashRecords(arg0 *big.Int) (*big.Int, error) {
	return _OperatorSlash.Contract.NftExitDelayedSlashRecords(&_OperatorSlash.CallOpts, arg0)
}

// NftHasCompensated is a free data retrieval call binding the contract method 0xa122faa4.
//
// Solidity: function nftHasCompensated(uint256 ) view returns(uint256)
func (_OperatorSlash *OperatorSlashCaller) NftHasCompensated(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OperatorSlash.contract.Call(opts, &out, "nftHasCompensated", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftHasCompensated is a free data retrieval call binding the contract method 0xa122faa4.
//
// Solidity: function nftHasCompensated(uint256 ) view returns(uint256)
func (_OperatorSlash *OperatorSlashSession) NftHasCompensated(arg0 *big.Int) (*big.Int, error) {
	return _OperatorSlash.Contract.NftHasCompensated(&_OperatorSlash.CallOpts, arg0)
}

// NftHasCompensated is a free data retrieval call binding the contract method 0xa122faa4.
//
// Solidity: function nftHasCompensated(uint256 ) view returns(uint256)
func (_OperatorSlash *OperatorSlashCallerSession) NftHasCompensated(arg0 *big.Int) (*big.Int, error) {
	return _OperatorSlash.Contract.NftHasCompensated(&_OperatorSlash.CallOpts, arg0)
}

// NftWillCompensated is a free data retrieval call binding the contract method 0xe0b64f27.
//
// Solidity: function nftWillCompensated(uint256 ) view returns(uint256)
func (_OperatorSlash *OperatorSlashCaller) NftWillCompensated(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OperatorSlash.contract.Call(opts, &out, "nftWillCompensated", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftWillCompensated is a free data retrieval call binding the contract method 0xe0b64f27.
//
// Solidity: function nftWillCompensated(uint256 ) view returns(uint256)
func (_OperatorSlash *OperatorSlashSession) NftWillCompensated(arg0 *big.Int) (*big.Int, error) {
	return _OperatorSlash.Contract.NftWillCompensated(&_OperatorSlash.CallOpts, arg0)
}

// NftWillCompensated is a free data retrieval call binding the contract method 0xe0b64f27.
//
// Solidity: function nftWillCompensated(uint256 ) view returns(uint256)
func (_OperatorSlash *OperatorSlashCallerSession) NftWillCompensated(arg0 *big.Int) (*big.Int, error) {
	return _OperatorSlash.Contract.NftWillCompensated(&_OperatorSlash.CallOpts, arg0)
}

// NodeOperatorRegistryContract is a free data retrieval call binding the contract method 0x9a42e0ba.
//
// Solidity: function nodeOperatorRegistryContract() view returns(address)
func (_OperatorSlash *OperatorSlashCaller) NodeOperatorRegistryContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorSlash.contract.Call(opts, &out, "nodeOperatorRegistryContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeOperatorRegistryContract is a free data retrieval call binding the contract method 0x9a42e0ba.
//
// Solidity: function nodeOperatorRegistryContract() view returns(address)
func (_OperatorSlash *OperatorSlashSession) NodeOperatorRegistryContract() (common.Address, error) {
	return _OperatorSlash.Contract.NodeOperatorRegistryContract(&_OperatorSlash.CallOpts)
}

// NodeOperatorRegistryContract is a free data retrieval call binding the contract method 0x9a42e0ba.
//
// Solidity: function nodeOperatorRegistryContract() view returns(address)
func (_OperatorSlash *OperatorSlashCallerSession) NodeOperatorRegistryContract() (common.Address, error) {
	return _OperatorSlash.Contract.NodeOperatorRegistryContract(&_OperatorSlash.CallOpts)
}

// OperatorCompensatedIndex is a free data retrieval call binding the contract method 0x2661005e.
//
// Solidity: function operatorCompensatedIndex() view returns(uint256)
func (_OperatorSlash *OperatorSlashCaller) OperatorCompensatedIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OperatorSlash.contract.Call(opts, &out, "operatorCompensatedIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorCompensatedIndex is a free data retrieval call binding the contract method 0x2661005e.
//
// Solidity: function operatorCompensatedIndex() view returns(uint256)
func (_OperatorSlash *OperatorSlashSession) OperatorCompensatedIndex() (*big.Int, error) {
	return _OperatorSlash.Contract.OperatorCompensatedIndex(&_OperatorSlash.CallOpts)
}

// OperatorCompensatedIndex is a free data retrieval call binding the contract method 0x2661005e.
//
// Solidity: function operatorCompensatedIndex() view returns(uint256)
func (_OperatorSlash *OperatorSlashCallerSession) OperatorCompensatedIndex() (*big.Int, error) {
	return _OperatorSlash.Contract.OperatorCompensatedIndex(&_OperatorSlash.CallOpts)
}

// OperatorSlashArrears is a free data retrieval call binding the contract method 0xa7d18eba.
//
// Solidity: function operatorSlashArrears(uint256 , uint256 ) view returns(uint256)
func (_OperatorSlash *OperatorSlashCaller) OperatorSlashArrears(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OperatorSlash.contract.Call(opts, &out, "operatorSlashArrears", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorSlashArrears is a free data retrieval call binding the contract method 0xa7d18eba.
//
// Solidity: function operatorSlashArrears(uint256 , uint256 ) view returns(uint256)
func (_OperatorSlash *OperatorSlashSession) OperatorSlashArrears(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _OperatorSlash.Contract.OperatorSlashArrears(&_OperatorSlash.CallOpts, arg0, arg1)
}

// OperatorSlashArrears is a free data retrieval call binding the contract method 0xa7d18eba.
//
// Solidity: function operatorSlashArrears(uint256 , uint256 ) view returns(uint256)
func (_OperatorSlash *OperatorSlashCallerSession) OperatorSlashArrears(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _OperatorSlash.Contract.OperatorSlashArrears(&_OperatorSlash.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OperatorSlash *OperatorSlashCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorSlash.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OperatorSlash *OperatorSlashSession) Owner() (common.Address, error) {
	return _OperatorSlash.Contract.Owner(&_OperatorSlash.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OperatorSlash *OperatorSlashCallerSession) Owner() (common.Address, error) {
	return _OperatorSlash.Contract.Owner(&_OperatorSlash.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_OperatorSlash *OperatorSlashCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OperatorSlash.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_OperatorSlash *OperatorSlashSession) ProxiableUUID() ([32]byte, error) {
	return _OperatorSlash.Contract.ProxiableUUID(&_OperatorSlash.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_OperatorSlash *OperatorSlashCallerSession) ProxiableUUID() ([32]byte, error) {
	return _OperatorSlash.Contract.ProxiableUUID(&_OperatorSlash.CallOpts)
}

// SlashAmountPerBlockPerValidator is a free data retrieval call binding the contract method 0xaf256032.
//
// Solidity: function slashAmountPerBlockPerValidator() view returns(uint256)
func (_OperatorSlash *OperatorSlashCaller) SlashAmountPerBlockPerValidator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OperatorSlash.contract.Call(opts, &out, "slashAmountPerBlockPerValidator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashAmountPerBlockPerValidator is a free data retrieval call binding the contract method 0xaf256032.
//
// Solidity: function slashAmountPerBlockPerValidator() view returns(uint256)
func (_OperatorSlash *OperatorSlashSession) SlashAmountPerBlockPerValidator() (*big.Int, error) {
	return _OperatorSlash.Contract.SlashAmountPerBlockPerValidator(&_OperatorSlash.CallOpts)
}

// SlashAmountPerBlockPerValidator is a free data retrieval call binding the contract method 0xaf256032.
//
// Solidity: function slashAmountPerBlockPerValidator() view returns(uint256)
func (_OperatorSlash *OperatorSlashCallerSession) SlashAmountPerBlockPerValidator() (*big.Int, error) {
	return _OperatorSlash.Contract.SlashAmountPerBlockPerValidator(&_OperatorSlash.CallOpts)
}

// VNFTContract is a free data retrieval call binding the contract method 0xfc03411f.
//
// Solidity: function vNFTContract() view returns(address)
func (_OperatorSlash *OperatorSlashCaller) VNFTContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorSlash.contract.Call(opts, &out, "vNFTContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VNFTContract is a free data retrieval call binding the contract method 0xfc03411f.
//
// Solidity: function vNFTContract() view returns(address)
func (_OperatorSlash *OperatorSlashSession) VNFTContract() (common.Address, error) {
	return _OperatorSlash.Contract.VNFTContract(&_OperatorSlash.CallOpts)
}

// VNFTContract is a free data retrieval call binding the contract method 0xfc03411f.
//
// Solidity: function vNFTContract() view returns(address)
func (_OperatorSlash *OperatorSlashCallerSession) VNFTContract() (common.Address, error) {
	return _OperatorSlash.Contract.VNFTContract(&_OperatorSlash.CallOpts)
}

// VaultManagerContractAddress is a free data retrieval call binding the contract method 0xb3ee9d6b.
//
// Solidity: function vaultManagerContractAddress() view returns(address)
func (_OperatorSlash *OperatorSlashCaller) VaultManagerContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorSlash.contract.Call(opts, &out, "vaultManagerContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultManagerContractAddress is a free data retrieval call binding the contract method 0xb3ee9d6b.
//
// Solidity: function vaultManagerContractAddress() view returns(address)
func (_OperatorSlash *OperatorSlashSession) VaultManagerContractAddress() (common.Address, error) {
	return _OperatorSlash.Contract.VaultManagerContractAddress(&_OperatorSlash.CallOpts)
}

// VaultManagerContractAddress is a free data retrieval call binding the contract method 0xb3ee9d6b.
//
// Solidity: function vaultManagerContractAddress() view returns(address)
func (_OperatorSlash *OperatorSlashCallerSession) VaultManagerContractAddress() (common.Address, error) {
	return _OperatorSlash.Contract.VaultManagerContractAddress(&_OperatorSlash.CallOpts)
}

// WithdrawalRequestContract is a free data retrieval call binding the contract method 0x45dcb639.
//
// Solidity: function withdrawalRequestContract() view returns(address)
func (_OperatorSlash *OperatorSlashCaller) WithdrawalRequestContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorSlash.contract.Call(opts, &out, "withdrawalRequestContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawalRequestContract is a free data retrieval call binding the contract method 0x45dcb639.
//
// Solidity: function withdrawalRequestContract() view returns(address)
func (_OperatorSlash *OperatorSlashSession) WithdrawalRequestContract() (common.Address, error) {
	return _OperatorSlash.Contract.WithdrawalRequestContract(&_OperatorSlash.CallOpts)
}

// WithdrawalRequestContract is a free data retrieval call binding the contract method 0x45dcb639.
//
// Solidity: function withdrawalRequestContract() view returns(address)
func (_OperatorSlash *OperatorSlashCallerSession) WithdrawalRequestContract() (common.Address, error) {
	return _OperatorSlash.Contract.WithdrawalRequestContract(&_OperatorSlash.CallOpts)
}

// ClaimCompensated is a paid mutator transaction binding the contract method 0x5b81dc37.
//
// Solidity: function claimCompensated(uint256[] _tokenIds, address _owner) returns(uint256)
func (_OperatorSlash *OperatorSlashTransactor) ClaimCompensated(opts *bind.TransactOpts, _tokenIds []*big.Int, _owner common.Address) (*types.Transaction, error) {
	return _OperatorSlash.contract.Transact(opts, "claimCompensated", _tokenIds, _owner)
}

// ClaimCompensated is a paid mutator transaction binding the contract method 0x5b81dc37.
//
// Solidity: function claimCompensated(uint256[] _tokenIds, address _owner) returns(uint256)
func (_OperatorSlash *OperatorSlashSession) ClaimCompensated(_tokenIds []*big.Int, _owner common.Address) (*types.Transaction, error) {
	return _OperatorSlash.Contract.ClaimCompensated(&_OperatorSlash.TransactOpts, _tokenIds, _owner)
}

// ClaimCompensated is a paid mutator transaction binding the contract method 0x5b81dc37.
//
// Solidity: function claimCompensated(uint256[] _tokenIds, address _owner) returns(uint256)
func (_OperatorSlash *OperatorSlashTransactorSession) ClaimCompensated(_tokenIds []*big.Int, _owner common.Address) (*types.Transaction, error) {
	return _OperatorSlash.Contract.ClaimCompensated(&_OperatorSlash.TransactOpts, _tokenIds, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x1460e390.
//
// Solidity: function initialize(address _dao, address _liquidStakingAddress, address _nVNFTContractAddress, address _nodeOperatorRegistryAddress, address _withdrawalRequestContractAddress, address _vaultManagerContractAddress, uint256 _delayedExitSlashStandard) returns()
func (_OperatorSlash *OperatorSlashTransactor) Initialize(opts *bind.TransactOpts, _dao common.Address, _liquidStakingAddress common.Address, _nVNFTContractAddress common.Address, _nodeOperatorRegistryAddress common.Address, _withdrawalRequestContractAddress common.Address, _vaultManagerContractAddress common.Address, _delayedExitSlashStandard *big.Int) (*types.Transaction, error) {
	return _OperatorSlash.contract.Transact(opts, "initialize", _dao, _liquidStakingAddress, _nVNFTContractAddress, _nodeOperatorRegistryAddress, _withdrawalRequestContractAddress, _vaultManagerContractAddress, _delayedExitSlashStandard)
}

// Initialize is a paid mutator transaction binding the contract method 0x1460e390.
//
// Solidity: function initialize(address _dao, address _liquidStakingAddress, address _nVNFTContractAddress, address _nodeOperatorRegistryAddress, address _withdrawalRequestContractAddress, address _vaultManagerContractAddress, uint256 _delayedExitSlashStandard) returns()
func (_OperatorSlash *OperatorSlashSession) Initialize(_dao common.Address, _liquidStakingAddress common.Address, _nVNFTContractAddress common.Address, _nodeOperatorRegistryAddress common.Address, _withdrawalRequestContractAddress common.Address, _vaultManagerContractAddress common.Address, _delayedExitSlashStandard *big.Int) (*types.Transaction, error) {
	return _OperatorSlash.Contract.Initialize(&_OperatorSlash.TransactOpts, _dao, _liquidStakingAddress, _nVNFTContractAddress, _nodeOperatorRegistryAddress, _withdrawalRequestContractAddress, _vaultManagerContractAddress, _delayedExitSlashStandard)
}

// Initialize is a paid mutator transaction binding the contract method 0x1460e390.
//
// Solidity: function initialize(address _dao, address _liquidStakingAddress, address _nVNFTContractAddress, address _nodeOperatorRegistryAddress, address _withdrawalRequestContractAddress, address _vaultManagerContractAddress, uint256 _delayedExitSlashStandard) returns()
func (_OperatorSlash *OperatorSlashTransactorSession) Initialize(_dao common.Address, _liquidStakingAddress common.Address, _nVNFTContractAddress common.Address, _nodeOperatorRegistryAddress common.Address, _withdrawalRequestContractAddress common.Address, _vaultManagerContractAddress common.Address, _delayedExitSlashStandard *big.Int) (*types.Transaction, error) {
	return _OperatorSlash.Contract.Initialize(&_OperatorSlash.TransactOpts, _dao, _liquidStakingAddress, _nVNFTContractAddress, _nodeOperatorRegistryAddress, _withdrawalRequestContractAddress, _vaultManagerContractAddress, _delayedExitSlashStandard)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OperatorSlash *OperatorSlashTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorSlash.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OperatorSlash *OperatorSlashSession) RenounceOwnership() (*types.Transaction, error) {
	return _OperatorSlash.Contract.RenounceOwnership(&_OperatorSlash.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OperatorSlash *OperatorSlashTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OperatorSlash.Contract.RenounceOwnership(&_OperatorSlash.TransactOpts)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_OperatorSlash *OperatorSlashTransactor) SetDaoAddress(opts *bind.TransactOpts, _dao common.Address) (*types.Transaction, error) {
	return _OperatorSlash.contract.Transact(opts, "setDaoAddress", _dao)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_OperatorSlash *OperatorSlashSession) SetDaoAddress(_dao common.Address) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SetDaoAddress(&_OperatorSlash.TransactOpts, _dao)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_OperatorSlash *OperatorSlashTransactorSession) SetDaoAddress(_dao common.Address) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SetDaoAddress(&_OperatorSlash.TransactOpts, _dao)
}

// SetLiquidStaking is a paid mutator transaction binding the contract method 0x08211be5.
//
// Solidity: function setLiquidStaking(address _liquidStakingContractAddress) returns()
func (_OperatorSlash *OperatorSlashTransactor) SetLiquidStaking(opts *bind.TransactOpts, _liquidStakingContractAddress common.Address) (*types.Transaction, error) {
	return _OperatorSlash.contract.Transact(opts, "setLiquidStaking", _liquidStakingContractAddress)
}

// SetLiquidStaking is a paid mutator transaction binding the contract method 0x08211be5.
//
// Solidity: function setLiquidStaking(address _liquidStakingContractAddress) returns()
func (_OperatorSlash *OperatorSlashSession) SetLiquidStaking(_liquidStakingContractAddress common.Address) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SetLiquidStaking(&_OperatorSlash.TransactOpts, _liquidStakingContractAddress)
}

// SetLiquidStaking is a paid mutator transaction binding the contract method 0x08211be5.
//
// Solidity: function setLiquidStaking(address _liquidStakingContractAddress) returns()
func (_OperatorSlash *OperatorSlashTransactorSession) SetLiquidStaking(_liquidStakingContractAddress common.Address) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SetLiquidStaking(&_OperatorSlash.TransactOpts, _liquidStakingContractAddress)
}

// SetNodeOperatorRegistryContract is a paid mutator transaction binding the contract method 0xcb23473e.
//
// Solidity: function setNodeOperatorRegistryContract(address _nodeOperatorRegistryContract) returns()
func (_OperatorSlash *OperatorSlashTransactor) SetNodeOperatorRegistryContract(opts *bind.TransactOpts, _nodeOperatorRegistryContract common.Address) (*types.Transaction, error) {
	return _OperatorSlash.contract.Transact(opts, "setNodeOperatorRegistryContract", _nodeOperatorRegistryContract)
}

// SetNodeOperatorRegistryContract is a paid mutator transaction binding the contract method 0xcb23473e.
//
// Solidity: function setNodeOperatorRegistryContract(address _nodeOperatorRegistryContract) returns()
func (_OperatorSlash *OperatorSlashSession) SetNodeOperatorRegistryContract(_nodeOperatorRegistryContract common.Address) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SetNodeOperatorRegistryContract(&_OperatorSlash.TransactOpts, _nodeOperatorRegistryContract)
}

// SetNodeOperatorRegistryContract is a paid mutator transaction binding the contract method 0xcb23473e.
//
// Solidity: function setNodeOperatorRegistryContract(address _nodeOperatorRegistryContract) returns()
func (_OperatorSlash *OperatorSlashTransactorSession) SetNodeOperatorRegistryContract(_nodeOperatorRegistryContract common.Address) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SetNodeOperatorRegistryContract(&_OperatorSlash.TransactOpts, _nodeOperatorRegistryContract)
}

// SetSlashAmountPerBlockPerValidator is a paid mutator transaction binding the contract method 0x9b65c6de.
//
// Solidity: function setSlashAmountPerBlockPerValidator(uint256 _slashAmountPerBlockPerValidator) returns()
func (_OperatorSlash *OperatorSlashTransactor) SetSlashAmountPerBlockPerValidator(opts *bind.TransactOpts, _slashAmountPerBlockPerValidator *big.Int) (*types.Transaction, error) {
	return _OperatorSlash.contract.Transact(opts, "setSlashAmountPerBlockPerValidator", _slashAmountPerBlockPerValidator)
}

// SetSlashAmountPerBlockPerValidator is a paid mutator transaction binding the contract method 0x9b65c6de.
//
// Solidity: function setSlashAmountPerBlockPerValidator(uint256 _slashAmountPerBlockPerValidator) returns()
func (_OperatorSlash *OperatorSlashSession) SetSlashAmountPerBlockPerValidator(_slashAmountPerBlockPerValidator *big.Int) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SetSlashAmountPerBlockPerValidator(&_OperatorSlash.TransactOpts, _slashAmountPerBlockPerValidator)
}

// SetSlashAmountPerBlockPerValidator is a paid mutator transaction binding the contract method 0x9b65c6de.
//
// Solidity: function setSlashAmountPerBlockPerValidator(uint256 _slashAmountPerBlockPerValidator) returns()
func (_OperatorSlash *OperatorSlashTransactorSession) SetSlashAmountPerBlockPerValidator(_slashAmountPerBlockPerValidator *big.Int) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SetSlashAmountPerBlockPerValidator(&_OperatorSlash.TransactOpts, _slashAmountPerBlockPerValidator)
}

// SetVaultManagerContract is a paid mutator transaction binding the contract method 0xbe14652a.
//
// Solidity: function setVaultManagerContract(address _vaultManagerContract) returns()
func (_OperatorSlash *OperatorSlashTransactor) SetVaultManagerContract(opts *bind.TransactOpts, _vaultManagerContract common.Address) (*types.Transaction, error) {
	return _OperatorSlash.contract.Transact(opts, "setVaultManagerContract", _vaultManagerContract)
}

// SetVaultManagerContract is a paid mutator transaction binding the contract method 0xbe14652a.
//
// Solidity: function setVaultManagerContract(address _vaultManagerContract) returns()
func (_OperatorSlash *OperatorSlashSession) SetVaultManagerContract(_vaultManagerContract common.Address) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SetVaultManagerContract(&_OperatorSlash.TransactOpts, _vaultManagerContract)
}

// SetVaultManagerContract is a paid mutator transaction binding the contract method 0xbe14652a.
//
// Solidity: function setVaultManagerContract(address _vaultManagerContract) returns()
func (_OperatorSlash *OperatorSlashTransactorSession) SetVaultManagerContract(_vaultManagerContract common.Address) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SetVaultManagerContract(&_OperatorSlash.TransactOpts, _vaultManagerContract)
}

// SetWithdrawalRequestContract is a paid mutator transaction binding the contract method 0x4408fb50.
//
// Solidity: function setWithdrawalRequestContract(address _withdrawalRequestContractAddress) returns()
func (_OperatorSlash *OperatorSlashTransactor) SetWithdrawalRequestContract(opts *bind.TransactOpts, _withdrawalRequestContractAddress common.Address) (*types.Transaction, error) {
	return _OperatorSlash.contract.Transact(opts, "setWithdrawalRequestContract", _withdrawalRequestContractAddress)
}

// SetWithdrawalRequestContract is a paid mutator transaction binding the contract method 0x4408fb50.
//
// Solidity: function setWithdrawalRequestContract(address _withdrawalRequestContractAddress) returns()
func (_OperatorSlash *OperatorSlashSession) SetWithdrawalRequestContract(_withdrawalRequestContractAddress common.Address) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SetWithdrawalRequestContract(&_OperatorSlash.TransactOpts, _withdrawalRequestContractAddress)
}

// SetWithdrawalRequestContract is a paid mutator transaction binding the contract method 0x4408fb50.
//
// Solidity: function setWithdrawalRequestContract(address _withdrawalRequestContractAddress) returns()
func (_OperatorSlash *OperatorSlashTransactorSession) SetWithdrawalRequestContract(_withdrawalRequestContractAddress common.Address) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SetWithdrawalRequestContract(&_OperatorSlash.TransactOpts, _withdrawalRequestContractAddress)
}

// SlashArrearsReceive is a paid mutator transaction binding the contract method 0xe72544c9.
//
// Solidity: function slashArrearsReceive(uint256 _operatorId, uint256 _amount) payable returns()
func (_OperatorSlash *OperatorSlashTransactor) SlashArrearsReceive(opts *bind.TransactOpts, _operatorId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _OperatorSlash.contract.Transact(opts, "slashArrearsReceive", _operatorId, _amount)
}

// SlashArrearsReceive is a paid mutator transaction binding the contract method 0xe72544c9.
//
// Solidity: function slashArrearsReceive(uint256 _operatorId, uint256 _amount) payable returns()
func (_OperatorSlash *OperatorSlashSession) SlashArrearsReceive(_operatorId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SlashArrearsReceive(&_OperatorSlash.TransactOpts, _operatorId, _amount)
}

// SlashArrearsReceive is a paid mutator transaction binding the contract method 0xe72544c9.
//
// Solidity: function slashArrearsReceive(uint256 _operatorId, uint256 _amount) payable returns()
func (_OperatorSlash *OperatorSlashTransactorSession) SlashArrearsReceive(_operatorId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SlashArrearsReceive(&_OperatorSlash.TransactOpts, _operatorId, _amount)
}

// SlashOfExitDelayed is a paid mutator transaction binding the contract method 0xb21a6370.
//
// Solidity: function slashOfExitDelayed(uint256[] _nftExitDelayedTokenIds, uint256[] _largeExitDelayedRequestIds) returns()
func (_OperatorSlash *OperatorSlashTransactor) SlashOfExitDelayed(opts *bind.TransactOpts, _nftExitDelayedTokenIds []*big.Int, _largeExitDelayedRequestIds []*big.Int) (*types.Transaction, error) {
	return _OperatorSlash.contract.Transact(opts, "slashOfExitDelayed", _nftExitDelayedTokenIds, _largeExitDelayedRequestIds)
}

// SlashOfExitDelayed is a paid mutator transaction binding the contract method 0xb21a6370.
//
// Solidity: function slashOfExitDelayed(uint256[] _nftExitDelayedTokenIds, uint256[] _largeExitDelayedRequestIds) returns()
func (_OperatorSlash *OperatorSlashSession) SlashOfExitDelayed(_nftExitDelayedTokenIds []*big.Int, _largeExitDelayedRequestIds []*big.Int) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SlashOfExitDelayed(&_OperatorSlash.TransactOpts, _nftExitDelayedTokenIds, _largeExitDelayedRequestIds)
}

// SlashOfExitDelayed is a paid mutator transaction binding the contract method 0xb21a6370.
//
// Solidity: function slashOfExitDelayed(uint256[] _nftExitDelayedTokenIds, uint256[] _largeExitDelayedRequestIds) returns()
func (_OperatorSlash *OperatorSlashTransactorSession) SlashOfExitDelayed(_nftExitDelayedTokenIds []*big.Int, _largeExitDelayedRequestIds []*big.Int) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SlashOfExitDelayed(&_OperatorSlash.TransactOpts, _nftExitDelayedTokenIds, _largeExitDelayedRequestIds)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x538d7ab0.
//
// Solidity: function slashOperator(uint256[] _exitTokenIds, uint256[] _amounts) returns()
func (_OperatorSlash *OperatorSlashTransactor) SlashOperator(opts *bind.TransactOpts, _exitTokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _OperatorSlash.contract.Transact(opts, "slashOperator", _exitTokenIds, _amounts)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x538d7ab0.
//
// Solidity: function slashOperator(uint256[] _exitTokenIds, uint256[] _amounts) returns()
func (_OperatorSlash *OperatorSlashSession) SlashOperator(_exitTokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SlashOperator(&_OperatorSlash.TransactOpts, _exitTokenIds, _amounts)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x538d7ab0.
//
// Solidity: function slashOperator(uint256[] _exitTokenIds, uint256[] _amounts) returns()
func (_OperatorSlash *OperatorSlashTransactorSession) SlashOperator(_exitTokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SlashOperator(&_OperatorSlash.TransactOpts, _exitTokenIds, _amounts)
}

// SlashReceive is a paid mutator transaction binding the contract method 0xc0849498.
//
// Solidity: function slashReceive(uint256[] _exitTokenIds, uint256[] _slashAmounts, uint256[] _requireAmounts) payable returns()
func (_OperatorSlash *OperatorSlashTransactor) SlashReceive(opts *bind.TransactOpts, _exitTokenIds []*big.Int, _slashAmounts []*big.Int, _requireAmounts []*big.Int) (*types.Transaction, error) {
	return _OperatorSlash.contract.Transact(opts, "slashReceive", _exitTokenIds, _slashAmounts, _requireAmounts)
}

// SlashReceive is a paid mutator transaction binding the contract method 0xc0849498.
//
// Solidity: function slashReceive(uint256[] _exitTokenIds, uint256[] _slashAmounts, uint256[] _requireAmounts) payable returns()
func (_OperatorSlash *OperatorSlashSession) SlashReceive(_exitTokenIds []*big.Int, _slashAmounts []*big.Int, _requireAmounts []*big.Int) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SlashReceive(&_OperatorSlash.TransactOpts, _exitTokenIds, _slashAmounts, _requireAmounts)
}

// SlashReceive is a paid mutator transaction binding the contract method 0xc0849498.
//
// Solidity: function slashReceive(uint256[] _exitTokenIds, uint256[] _slashAmounts, uint256[] _requireAmounts) payable returns()
func (_OperatorSlash *OperatorSlashTransactorSession) SlashReceive(_exitTokenIds []*big.Int, _slashAmounts []*big.Int, _requireAmounts []*big.Int) (*types.Transaction, error) {
	return _OperatorSlash.Contract.SlashReceive(&_OperatorSlash.TransactOpts, _exitTokenIds, _slashAmounts, _requireAmounts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OperatorSlash *OperatorSlashTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OperatorSlash.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OperatorSlash *OperatorSlashSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OperatorSlash.Contract.TransferOwnership(&_OperatorSlash.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OperatorSlash *OperatorSlashTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OperatorSlash.Contract.TransferOwnership(&_OperatorSlash.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_OperatorSlash *OperatorSlashTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _OperatorSlash.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_OperatorSlash *OperatorSlashSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _OperatorSlash.Contract.UpgradeTo(&_OperatorSlash.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_OperatorSlash *OperatorSlashTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _OperatorSlash.Contract.UpgradeTo(&_OperatorSlash.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_OperatorSlash *OperatorSlashTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _OperatorSlash.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_OperatorSlash *OperatorSlashSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _OperatorSlash.Contract.UpgradeToAndCall(&_OperatorSlash.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_OperatorSlash *OperatorSlashTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _OperatorSlash.Contract.UpgradeToAndCall(&_OperatorSlash.TransactOpts, newImplementation, data)
}

// OperatorSlashAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the OperatorSlash contract.
type OperatorSlashAdminChangedIterator struct {
	Event *OperatorSlashAdminChanged // Event containing the contract specifics and raw log

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
func (it *OperatorSlashAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSlashAdminChanged)
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
		it.Event = new(OperatorSlashAdminChanged)
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
func (it *OperatorSlashAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSlashAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSlashAdminChanged represents a AdminChanged event raised by the OperatorSlash contract.
type OperatorSlashAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_OperatorSlash *OperatorSlashFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*OperatorSlashAdminChangedIterator, error) {

	logs, sub, err := _OperatorSlash.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &OperatorSlashAdminChangedIterator{contract: _OperatorSlash.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_OperatorSlash *OperatorSlashFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *OperatorSlashAdminChanged) (event.Subscription, error) {

	logs, sub, err := _OperatorSlash.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSlashAdminChanged)
				if err := _OperatorSlash.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_OperatorSlash *OperatorSlashFilterer) ParseAdminChanged(log types.Log) (*OperatorSlashAdminChanged, error) {
	event := new(OperatorSlashAdminChanged)
	if err := _OperatorSlash.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorSlashArrearsReceiveOfSlashIterator is returned from FilterArrearsReceiveOfSlash and is used to iterate over the raw logs and unpacked data for ArrearsReceiveOfSlash events raised by the OperatorSlash contract.
type OperatorSlashArrearsReceiveOfSlashIterator struct {
	Event *OperatorSlashArrearsReceiveOfSlash // Event containing the contract specifics and raw log

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
func (it *OperatorSlashArrearsReceiveOfSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSlashArrearsReceiveOfSlash)
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
		it.Event = new(OperatorSlashArrearsReceiveOfSlash)
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
func (it *OperatorSlashArrearsReceiveOfSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSlashArrearsReceiveOfSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSlashArrearsReceiveOfSlash represents a ArrearsReceiveOfSlash event raised by the OperatorSlash contract.
type OperatorSlashArrearsReceiveOfSlash struct {
	OperatorId *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterArrearsReceiveOfSlash is a free log retrieval operation binding the contract event 0xa2478ef76ec5108b836ec0b59cc344bf613e964e82254a6c2e22a548e9d0fe75.
//
// Solidity: event ArrearsReceiveOfSlash(uint256 _operatorId, uint256 _amount)
func (_OperatorSlash *OperatorSlashFilterer) FilterArrearsReceiveOfSlash(opts *bind.FilterOpts) (*OperatorSlashArrearsReceiveOfSlashIterator, error) {

	logs, sub, err := _OperatorSlash.contract.FilterLogs(opts, "ArrearsReceiveOfSlash")
	if err != nil {
		return nil, err
	}
	return &OperatorSlashArrearsReceiveOfSlashIterator{contract: _OperatorSlash.contract, event: "ArrearsReceiveOfSlash", logs: logs, sub: sub}, nil
}

// WatchArrearsReceiveOfSlash is a free log subscription operation binding the contract event 0xa2478ef76ec5108b836ec0b59cc344bf613e964e82254a6c2e22a548e9d0fe75.
//
// Solidity: event ArrearsReceiveOfSlash(uint256 _operatorId, uint256 _amount)
func (_OperatorSlash *OperatorSlashFilterer) WatchArrearsReceiveOfSlash(opts *bind.WatchOpts, sink chan<- *OperatorSlashArrearsReceiveOfSlash) (event.Subscription, error) {

	logs, sub, err := _OperatorSlash.contract.WatchLogs(opts, "ArrearsReceiveOfSlash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSlashArrearsReceiveOfSlash)
				if err := _OperatorSlash.contract.UnpackLog(event, "ArrearsReceiveOfSlash", log); err != nil {
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

// ParseArrearsReceiveOfSlash is a log parse operation binding the contract event 0xa2478ef76ec5108b836ec0b59cc344bf613e964e82254a6c2e22a548e9d0fe75.
//
// Solidity: event ArrearsReceiveOfSlash(uint256 _operatorId, uint256 _amount)
func (_OperatorSlash *OperatorSlashFilterer) ParseArrearsReceiveOfSlash(log types.Log) (*OperatorSlashArrearsReceiveOfSlash, error) {
	event := new(OperatorSlashArrearsReceiveOfSlash)
	if err := _OperatorSlash.contract.UnpackLog(event, "ArrearsReceiveOfSlash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorSlashBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the OperatorSlash contract.
type OperatorSlashBeaconUpgradedIterator struct {
	Event *OperatorSlashBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *OperatorSlashBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSlashBeaconUpgraded)
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
		it.Event = new(OperatorSlashBeaconUpgraded)
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
func (it *OperatorSlashBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSlashBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSlashBeaconUpgraded represents a BeaconUpgraded event raised by the OperatorSlash contract.
type OperatorSlashBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_OperatorSlash *OperatorSlashFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*OperatorSlashBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _OperatorSlash.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &OperatorSlashBeaconUpgradedIterator{contract: _OperatorSlash.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_OperatorSlash *OperatorSlashFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *OperatorSlashBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _OperatorSlash.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSlashBeaconUpgraded)
				if err := _OperatorSlash.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_OperatorSlash *OperatorSlashFilterer) ParseBeaconUpgraded(log types.Log) (*OperatorSlashBeaconUpgraded, error) {
	event := new(OperatorSlashBeaconUpgraded)
	if err := _OperatorSlash.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorSlashCompensatedClaimedIterator is returned from FilterCompensatedClaimed and is used to iterate over the raw logs and unpacked data for CompensatedClaimed events raised by the OperatorSlash contract.
type OperatorSlashCompensatedClaimedIterator struct {
	Event *OperatorSlashCompensatedClaimed // Event containing the contract specifics and raw log

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
func (it *OperatorSlashCompensatedClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSlashCompensatedClaimed)
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
		it.Event = new(OperatorSlashCompensatedClaimed)
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
func (it *OperatorSlashCompensatedClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSlashCompensatedClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSlashCompensatedClaimed represents a CompensatedClaimed event raised by the OperatorSlash contract.
type OperatorSlashCompensatedClaimed struct {
	Owner            common.Address
	TotalCompensated *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCompensatedClaimed is a free log retrieval operation binding the contract event 0x94adb2a84514113270a81b0c2e53d7f91cd818356e1fa82635116aa75722ca46.
//
// Solidity: event CompensatedClaimed(address _owner, uint256 _totalCompensated)
func (_OperatorSlash *OperatorSlashFilterer) FilterCompensatedClaimed(opts *bind.FilterOpts) (*OperatorSlashCompensatedClaimedIterator, error) {

	logs, sub, err := _OperatorSlash.contract.FilterLogs(opts, "CompensatedClaimed")
	if err != nil {
		return nil, err
	}
	return &OperatorSlashCompensatedClaimedIterator{contract: _OperatorSlash.contract, event: "CompensatedClaimed", logs: logs, sub: sub}, nil
}

// WatchCompensatedClaimed is a free log subscription operation binding the contract event 0x94adb2a84514113270a81b0c2e53d7f91cd818356e1fa82635116aa75722ca46.
//
// Solidity: event CompensatedClaimed(address _owner, uint256 _totalCompensated)
func (_OperatorSlash *OperatorSlashFilterer) WatchCompensatedClaimed(opts *bind.WatchOpts, sink chan<- *OperatorSlashCompensatedClaimed) (event.Subscription, error) {

	logs, sub, err := _OperatorSlash.contract.WatchLogs(opts, "CompensatedClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSlashCompensatedClaimed)
				if err := _OperatorSlash.contract.UnpackLog(event, "CompensatedClaimed", log); err != nil {
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

// ParseCompensatedClaimed is a log parse operation binding the contract event 0x94adb2a84514113270a81b0c2e53d7f91cd818356e1fa82635116aa75722ca46.
//
// Solidity: event CompensatedClaimed(address _owner, uint256 _totalCompensated)
func (_OperatorSlash *OperatorSlashFilterer) ParseCompensatedClaimed(log types.Log) (*OperatorSlashCompensatedClaimed, error) {
	event := new(OperatorSlashCompensatedClaimed)
	if err := _OperatorSlash.contract.UnpackLog(event, "CompensatedClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorSlashDaoAddressChangedIterator is returned from FilterDaoAddressChanged and is used to iterate over the raw logs and unpacked data for DaoAddressChanged events raised by the OperatorSlash contract.
type OperatorSlashDaoAddressChangedIterator struct {
	Event *OperatorSlashDaoAddressChanged // Event containing the contract specifics and raw log

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
func (it *OperatorSlashDaoAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSlashDaoAddressChanged)
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
		it.Event = new(OperatorSlashDaoAddressChanged)
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
func (it *OperatorSlashDaoAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSlashDaoAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSlashDaoAddressChanged represents a DaoAddressChanged event raised by the OperatorSlash contract.
type OperatorSlashDaoAddressChanged struct {
	OldDao common.Address
	Dao    common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDaoAddressChanged is a free log retrieval operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address oldDao, address _dao)
func (_OperatorSlash *OperatorSlashFilterer) FilterDaoAddressChanged(opts *bind.FilterOpts) (*OperatorSlashDaoAddressChangedIterator, error) {

	logs, sub, err := _OperatorSlash.contract.FilterLogs(opts, "DaoAddressChanged")
	if err != nil {
		return nil, err
	}
	return &OperatorSlashDaoAddressChangedIterator{contract: _OperatorSlash.contract, event: "DaoAddressChanged", logs: logs, sub: sub}, nil
}

// WatchDaoAddressChanged is a free log subscription operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address oldDao, address _dao)
func (_OperatorSlash *OperatorSlashFilterer) WatchDaoAddressChanged(opts *bind.WatchOpts, sink chan<- *OperatorSlashDaoAddressChanged) (event.Subscription, error) {

	logs, sub, err := _OperatorSlash.contract.WatchLogs(opts, "DaoAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSlashDaoAddressChanged)
				if err := _OperatorSlash.contract.UnpackLog(event, "DaoAddressChanged", log); err != nil {
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
// Solidity: event DaoAddressChanged(address oldDao, address _dao)
func (_OperatorSlash *OperatorSlashFilterer) ParseDaoAddressChanged(log types.Log) (*OperatorSlashDaoAddressChanged, error) {
	event := new(OperatorSlashDaoAddressChanged)
	if err := _OperatorSlash.contract.UnpackLog(event, "DaoAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorSlashInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OperatorSlash contract.
type OperatorSlashInitializedIterator struct {
	Event *OperatorSlashInitialized // Event containing the contract specifics and raw log

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
func (it *OperatorSlashInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSlashInitialized)
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
		it.Event = new(OperatorSlashInitialized)
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
func (it *OperatorSlashInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSlashInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSlashInitialized represents a Initialized event raised by the OperatorSlash contract.
type OperatorSlashInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OperatorSlash *OperatorSlashFilterer) FilterInitialized(opts *bind.FilterOpts) (*OperatorSlashInitializedIterator, error) {

	logs, sub, err := _OperatorSlash.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OperatorSlashInitializedIterator{contract: _OperatorSlash.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OperatorSlash *OperatorSlashFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OperatorSlashInitialized) (event.Subscription, error) {

	logs, sub, err := _OperatorSlash.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSlashInitialized)
				if err := _OperatorSlash.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_OperatorSlash *OperatorSlashFilterer) ParseInitialized(log types.Log) (*OperatorSlashInitialized, error) {
	event := new(OperatorSlashInitialized)
	if err := _OperatorSlash.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorSlashLiquidStakingChangedIterator is returned from FilterLiquidStakingChanged and is used to iterate over the raw logs and unpacked data for LiquidStakingChanged events raised by the OperatorSlash contract.
type OperatorSlashLiquidStakingChangedIterator struct {
	Event *OperatorSlashLiquidStakingChanged // Event containing the contract specifics and raw log

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
func (it *OperatorSlashLiquidStakingChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSlashLiquidStakingChanged)
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
		it.Event = new(OperatorSlashLiquidStakingChanged)
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
func (it *OperatorSlashLiquidStakingChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSlashLiquidStakingChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSlashLiquidStakingChanged represents a LiquidStakingChanged event raised by the OperatorSlash contract.
type OperatorSlashLiquidStakingChanged struct {
	OldLiquidStakingContract     common.Address
	LiquidStakingContractAddress common.Address
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterLiquidStakingChanged is a free log retrieval operation binding the contract event 0x516e369f70685d2905d214a9e8567098c02a0e00f5a55bd30baca6b51d446cef.
//
// Solidity: event LiquidStakingChanged(address _oldLiquidStakingContract, address _liquidStakingContractAddress)
func (_OperatorSlash *OperatorSlashFilterer) FilterLiquidStakingChanged(opts *bind.FilterOpts) (*OperatorSlashLiquidStakingChangedIterator, error) {

	logs, sub, err := _OperatorSlash.contract.FilterLogs(opts, "LiquidStakingChanged")
	if err != nil {
		return nil, err
	}
	return &OperatorSlashLiquidStakingChangedIterator{contract: _OperatorSlash.contract, event: "LiquidStakingChanged", logs: logs, sub: sub}, nil
}

// WatchLiquidStakingChanged is a free log subscription operation binding the contract event 0x516e369f70685d2905d214a9e8567098c02a0e00f5a55bd30baca6b51d446cef.
//
// Solidity: event LiquidStakingChanged(address _oldLiquidStakingContract, address _liquidStakingContractAddress)
func (_OperatorSlash *OperatorSlashFilterer) WatchLiquidStakingChanged(opts *bind.WatchOpts, sink chan<- *OperatorSlashLiquidStakingChanged) (event.Subscription, error) {

	logs, sub, err := _OperatorSlash.contract.WatchLogs(opts, "LiquidStakingChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSlashLiquidStakingChanged)
				if err := _OperatorSlash.contract.UnpackLog(event, "LiquidStakingChanged", log); err != nil {
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
// Solidity: event LiquidStakingChanged(address _oldLiquidStakingContract, address _liquidStakingContractAddress)
func (_OperatorSlash *OperatorSlashFilterer) ParseLiquidStakingChanged(log types.Log) (*OperatorSlashLiquidStakingChanged, error) {
	event := new(OperatorSlashLiquidStakingChanged)
	if err := _OperatorSlash.contract.UnpackLog(event, "LiquidStakingChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorSlashNodeOperatorRegistryContractSetIterator is returned from FilterNodeOperatorRegistryContractSet and is used to iterate over the raw logs and unpacked data for NodeOperatorRegistryContractSet events raised by the OperatorSlash contract.
type OperatorSlashNodeOperatorRegistryContractSetIterator struct {
	Event *OperatorSlashNodeOperatorRegistryContractSet // Event containing the contract specifics and raw log

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
func (it *OperatorSlashNodeOperatorRegistryContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSlashNodeOperatorRegistryContractSet)
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
		it.Event = new(OperatorSlashNodeOperatorRegistryContractSet)
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
func (it *OperatorSlashNodeOperatorRegistryContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSlashNodeOperatorRegistryContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSlashNodeOperatorRegistryContractSet represents a NodeOperatorRegistryContractSet event raised by the OperatorSlash contract.
type OperatorSlashNodeOperatorRegistryContractSet struct {
	OldNodeOperatorRegistryContract common.Address
	NodeOperatorRegistryContract    common.Address
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorRegistryContractSet is a free log retrieval operation binding the contract event 0x2aa578b9d95064e7e90ca0af5e42ca5499f5e90bd32c4e401df52a686ac6993d.
//
// Solidity: event NodeOperatorRegistryContractSet(address _oldNodeOperatorRegistryContract, address _nodeOperatorRegistryContract)
func (_OperatorSlash *OperatorSlashFilterer) FilterNodeOperatorRegistryContractSet(opts *bind.FilterOpts) (*OperatorSlashNodeOperatorRegistryContractSetIterator, error) {

	logs, sub, err := _OperatorSlash.contract.FilterLogs(opts, "NodeOperatorRegistryContractSet")
	if err != nil {
		return nil, err
	}
	return &OperatorSlashNodeOperatorRegistryContractSetIterator{contract: _OperatorSlash.contract, event: "NodeOperatorRegistryContractSet", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorRegistryContractSet is a free log subscription operation binding the contract event 0x2aa578b9d95064e7e90ca0af5e42ca5499f5e90bd32c4e401df52a686ac6993d.
//
// Solidity: event NodeOperatorRegistryContractSet(address _oldNodeOperatorRegistryContract, address _nodeOperatorRegistryContract)
func (_OperatorSlash *OperatorSlashFilterer) WatchNodeOperatorRegistryContractSet(opts *bind.WatchOpts, sink chan<- *OperatorSlashNodeOperatorRegistryContractSet) (event.Subscription, error) {

	logs, sub, err := _OperatorSlash.contract.WatchLogs(opts, "NodeOperatorRegistryContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSlashNodeOperatorRegistryContractSet)
				if err := _OperatorSlash.contract.UnpackLog(event, "NodeOperatorRegistryContractSet", log); err != nil {
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

// ParseNodeOperatorRegistryContractSet is a log parse operation binding the contract event 0x2aa578b9d95064e7e90ca0af5e42ca5499f5e90bd32c4e401df52a686ac6993d.
//
// Solidity: event NodeOperatorRegistryContractSet(address _oldNodeOperatorRegistryContract, address _nodeOperatorRegistryContract)
func (_OperatorSlash *OperatorSlashFilterer) ParseNodeOperatorRegistryContractSet(log types.Log) (*OperatorSlashNodeOperatorRegistryContractSet, error) {
	event := new(OperatorSlashNodeOperatorRegistryContractSet)
	if err := _OperatorSlash.contract.UnpackLog(event, "NodeOperatorRegistryContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorSlashOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OperatorSlash contract.
type OperatorSlashOwnershipTransferredIterator struct {
	Event *OperatorSlashOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OperatorSlashOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSlashOwnershipTransferred)
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
		it.Event = new(OperatorSlashOwnershipTransferred)
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
func (it *OperatorSlashOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSlashOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSlashOwnershipTransferred represents a OwnershipTransferred event raised by the OperatorSlash contract.
type OperatorSlashOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OperatorSlash *OperatorSlashFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OperatorSlashOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OperatorSlash.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OperatorSlashOwnershipTransferredIterator{contract: _OperatorSlash.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OperatorSlash *OperatorSlashFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OperatorSlashOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OperatorSlash.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSlashOwnershipTransferred)
				if err := _OperatorSlash.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OperatorSlash *OperatorSlashFilterer) ParseOwnershipTransferred(log types.Log) (*OperatorSlashOwnershipTransferred, error) {
	event := new(OperatorSlashOwnershipTransferred)
	if err := _OperatorSlash.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorSlashSlashAmountPerBlockPerValidatorSetIterator is returned from FilterSlashAmountPerBlockPerValidatorSet and is used to iterate over the raw logs and unpacked data for SlashAmountPerBlockPerValidatorSet events raised by the OperatorSlash contract.
type OperatorSlashSlashAmountPerBlockPerValidatorSetIterator struct {
	Event *OperatorSlashSlashAmountPerBlockPerValidatorSet // Event containing the contract specifics and raw log

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
func (it *OperatorSlashSlashAmountPerBlockPerValidatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSlashSlashAmountPerBlockPerValidatorSet)
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
		it.Event = new(OperatorSlashSlashAmountPerBlockPerValidatorSet)
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
func (it *OperatorSlashSlashAmountPerBlockPerValidatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSlashSlashAmountPerBlockPerValidatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSlashSlashAmountPerBlockPerValidatorSet represents a SlashAmountPerBlockPerValidatorSet event raised by the OperatorSlash contract.
type OperatorSlashSlashAmountPerBlockPerValidatorSet struct {
	OldSlashAmountPerBlockPerValidator *big.Int
	SlashAmountPerBlockPerValidator    *big.Int
	Raw                                types.Log // Blockchain specific contextual infos
}

// FilterSlashAmountPerBlockPerValidatorSet is a free log retrieval operation binding the contract event 0x7383bf4f05c7c9288c15921cd69246945e6c98e886ec3ca6cd347f23f7871a31.
//
// Solidity: event SlashAmountPerBlockPerValidatorSet(uint256 _oldSlashAmountPerBlockPerValidator, uint256 _slashAmountPerBlockPerValidator)
func (_OperatorSlash *OperatorSlashFilterer) FilterSlashAmountPerBlockPerValidatorSet(opts *bind.FilterOpts) (*OperatorSlashSlashAmountPerBlockPerValidatorSetIterator, error) {

	logs, sub, err := _OperatorSlash.contract.FilterLogs(opts, "SlashAmountPerBlockPerValidatorSet")
	if err != nil {
		return nil, err
	}
	return &OperatorSlashSlashAmountPerBlockPerValidatorSetIterator{contract: _OperatorSlash.contract, event: "SlashAmountPerBlockPerValidatorSet", logs: logs, sub: sub}, nil
}

// WatchSlashAmountPerBlockPerValidatorSet is a free log subscription operation binding the contract event 0x7383bf4f05c7c9288c15921cd69246945e6c98e886ec3ca6cd347f23f7871a31.
//
// Solidity: event SlashAmountPerBlockPerValidatorSet(uint256 _oldSlashAmountPerBlockPerValidator, uint256 _slashAmountPerBlockPerValidator)
func (_OperatorSlash *OperatorSlashFilterer) WatchSlashAmountPerBlockPerValidatorSet(opts *bind.WatchOpts, sink chan<- *OperatorSlashSlashAmountPerBlockPerValidatorSet) (event.Subscription, error) {

	logs, sub, err := _OperatorSlash.contract.WatchLogs(opts, "SlashAmountPerBlockPerValidatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSlashSlashAmountPerBlockPerValidatorSet)
				if err := _OperatorSlash.contract.UnpackLog(event, "SlashAmountPerBlockPerValidatorSet", log); err != nil {
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

// ParseSlashAmountPerBlockPerValidatorSet is a log parse operation binding the contract event 0x7383bf4f05c7c9288c15921cd69246945e6c98e886ec3ca6cd347f23f7871a31.
//
// Solidity: event SlashAmountPerBlockPerValidatorSet(uint256 _oldSlashAmountPerBlockPerValidator, uint256 _slashAmountPerBlockPerValidator)
func (_OperatorSlash *OperatorSlashFilterer) ParseSlashAmountPerBlockPerValidatorSet(log types.Log) (*OperatorSlashSlashAmountPerBlockPerValidatorSet, error) {
	event := new(OperatorSlashSlashAmountPerBlockPerValidatorSet)
	if err := _OperatorSlash.contract.UnpackLog(event, "SlashAmountPerBlockPerValidatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorSlashSlashReceiveIterator is returned from FilterSlashReceive and is used to iterate over the raw logs and unpacked data for SlashReceive events raised by the OperatorSlash contract.
type OperatorSlashSlashReceiveIterator struct {
	Event *OperatorSlashSlashReceive // Event containing the contract specifics and raw log

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
func (it *OperatorSlashSlashReceiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSlashSlashReceive)
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
		it.Event = new(OperatorSlashSlashReceive)
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
func (it *OperatorSlashSlashReceiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSlashSlashReceiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSlashSlashReceive represents a SlashReceive event raised by the OperatorSlash contract.
type OperatorSlashSlashReceive struct {
	OperatorId    *big.Int
	TokenId       *big.Int
	SlashAmount   *big.Int
	RequirAmounts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSlashReceive is a free log retrieval operation binding the contract event 0xeeba612bf8aaf11b642d483da9db251b732afe21a8035197457492c293eee3f3.
//
// Solidity: event SlashReceive(uint256 _operatorId, uint256 tokenId, uint256 _slashAmount, uint256 _requirAmounts)
func (_OperatorSlash *OperatorSlashFilterer) FilterSlashReceive(opts *bind.FilterOpts) (*OperatorSlashSlashReceiveIterator, error) {

	logs, sub, err := _OperatorSlash.contract.FilterLogs(opts, "SlashReceive")
	if err != nil {
		return nil, err
	}
	return &OperatorSlashSlashReceiveIterator{contract: _OperatorSlash.contract, event: "SlashReceive", logs: logs, sub: sub}, nil
}

// WatchSlashReceive is a free log subscription operation binding the contract event 0xeeba612bf8aaf11b642d483da9db251b732afe21a8035197457492c293eee3f3.
//
// Solidity: event SlashReceive(uint256 _operatorId, uint256 tokenId, uint256 _slashAmount, uint256 _requirAmounts)
func (_OperatorSlash *OperatorSlashFilterer) WatchSlashReceive(opts *bind.WatchOpts, sink chan<- *OperatorSlashSlashReceive) (event.Subscription, error) {

	logs, sub, err := _OperatorSlash.contract.WatchLogs(opts, "SlashReceive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSlashSlashReceive)
				if err := _OperatorSlash.contract.UnpackLog(event, "SlashReceive", log); err != nil {
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

// ParseSlashReceive is a log parse operation binding the contract event 0xeeba612bf8aaf11b642d483da9db251b732afe21a8035197457492c293eee3f3.
//
// Solidity: event SlashReceive(uint256 _operatorId, uint256 tokenId, uint256 _slashAmount, uint256 _requirAmounts)
func (_OperatorSlash *OperatorSlashFilterer) ParseSlashReceive(log types.Log) (*OperatorSlashSlashReceive, error) {
	event := new(OperatorSlashSlashReceive)
	if err := _OperatorSlash.contract.UnpackLog(event, "SlashReceive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorSlashUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the OperatorSlash contract.
type OperatorSlashUpgradedIterator struct {
	Event *OperatorSlashUpgraded // Event containing the contract specifics and raw log

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
func (it *OperatorSlashUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSlashUpgraded)
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
		it.Event = new(OperatorSlashUpgraded)
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
func (it *OperatorSlashUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSlashUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSlashUpgraded represents a Upgraded event raised by the OperatorSlash contract.
type OperatorSlashUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_OperatorSlash *OperatorSlashFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*OperatorSlashUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _OperatorSlash.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &OperatorSlashUpgradedIterator{contract: _OperatorSlash.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_OperatorSlash *OperatorSlashFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *OperatorSlashUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _OperatorSlash.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSlashUpgraded)
				if err := _OperatorSlash.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_OperatorSlash *OperatorSlashFilterer) ParseUpgraded(log types.Log) (*OperatorSlashUpgraded, error) {
	event := new(OperatorSlashUpgraded)
	if err := _OperatorSlash.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorSlashVaultManagerContractSetIterator is returned from FilterVaultManagerContractSet and is used to iterate over the raw logs and unpacked data for VaultManagerContractSet events raised by the OperatorSlash contract.
type OperatorSlashVaultManagerContractSetIterator struct {
	Event *OperatorSlashVaultManagerContractSet // Event containing the contract specifics and raw log

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
func (it *OperatorSlashVaultManagerContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSlashVaultManagerContractSet)
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
		it.Event = new(OperatorSlashVaultManagerContractSet)
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
func (it *OperatorSlashVaultManagerContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSlashVaultManagerContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSlashVaultManagerContractSet represents a VaultManagerContractSet event raised by the OperatorSlash contract.
type OperatorSlashVaultManagerContractSet struct {
	OldVaultManagerContractAddress common.Address
	VaultManagerContract           common.Address
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterVaultManagerContractSet is a free log retrieval operation binding the contract event 0x136260758ef216be6f30b5244361f089faf99890f23864c0a63e2d2def24963f.
//
// Solidity: event VaultManagerContractSet(address _oldVaultManagerContractAddress, address _vaultManagerContract)
func (_OperatorSlash *OperatorSlashFilterer) FilterVaultManagerContractSet(opts *bind.FilterOpts) (*OperatorSlashVaultManagerContractSetIterator, error) {

	logs, sub, err := _OperatorSlash.contract.FilterLogs(opts, "VaultManagerContractSet")
	if err != nil {
		return nil, err
	}
	return &OperatorSlashVaultManagerContractSetIterator{contract: _OperatorSlash.contract, event: "VaultManagerContractSet", logs: logs, sub: sub}, nil
}

// WatchVaultManagerContractSet is a free log subscription operation binding the contract event 0x136260758ef216be6f30b5244361f089faf99890f23864c0a63e2d2def24963f.
//
// Solidity: event VaultManagerContractSet(address _oldVaultManagerContractAddress, address _vaultManagerContract)
func (_OperatorSlash *OperatorSlashFilterer) WatchVaultManagerContractSet(opts *bind.WatchOpts, sink chan<- *OperatorSlashVaultManagerContractSet) (event.Subscription, error) {

	logs, sub, err := _OperatorSlash.contract.WatchLogs(opts, "VaultManagerContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSlashVaultManagerContractSet)
				if err := _OperatorSlash.contract.UnpackLog(event, "VaultManagerContractSet", log); err != nil {
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

// ParseVaultManagerContractSet is a log parse operation binding the contract event 0x136260758ef216be6f30b5244361f089faf99890f23864c0a63e2d2def24963f.
//
// Solidity: event VaultManagerContractSet(address _oldVaultManagerContractAddress, address _vaultManagerContract)
func (_OperatorSlash *OperatorSlashFilterer) ParseVaultManagerContractSet(log types.Log) (*OperatorSlashVaultManagerContractSet, error) {
	event := new(OperatorSlashVaultManagerContractSet)
	if err := _OperatorSlash.contract.UnpackLog(event, "VaultManagerContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorSlashWithdrawalRequestContractSetIterator is returned from FilterWithdrawalRequestContractSet and is used to iterate over the raw logs and unpacked data for WithdrawalRequestContractSet events raised by the OperatorSlash contract.
type OperatorSlashWithdrawalRequestContractSetIterator struct {
	Event *OperatorSlashWithdrawalRequestContractSet // Event containing the contract specifics and raw log

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
func (it *OperatorSlashWithdrawalRequestContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSlashWithdrawalRequestContractSet)
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
		it.Event = new(OperatorSlashWithdrawalRequestContractSet)
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
func (it *OperatorSlashWithdrawalRequestContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSlashWithdrawalRequestContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSlashWithdrawalRequestContractSet represents a WithdrawalRequestContractSet event raised by the OperatorSlash contract.
type OperatorSlashWithdrawalRequestContractSet struct {
	OldWithdrawalRequestContract     common.Address
	WithdrawalRequestContractAddress common.Address
	Raw                              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequestContractSet is a free log retrieval operation binding the contract event 0xb3b3e321ffd1930a33d425b4d1453792a16fca40d763a14c9fc90005360d0598.
//
// Solidity: event WithdrawalRequestContractSet(address _oldWithdrawalRequestContract, address _withdrawalRequestContractAddress)
func (_OperatorSlash *OperatorSlashFilterer) FilterWithdrawalRequestContractSet(opts *bind.FilterOpts) (*OperatorSlashWithdrawalRequestContractSetIterator, error) {

	logs, sub, err := _OperatorSlash.contract.FilterLogs(opts, "WithdrawalRequestContractSet")
	if err != nil {
		return nil, err
	}
	return &OperatorSlashWithdrawalRequestContractSetIterator{contract: _OperatorSlash.contract, event: "WithdrawalRequestContractSet", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequestContractSet is a free log subscription operation binding the contract event 0xb3b3e321ffd1930a33d425b4d1453792a16fca40d763a14c9fc90005360d0598.
//
// Solidity: event WithdrawalRequestContractSet(address _oldWithdrawalRequestContract, address _withdrawalRequestContractAddress)
func (_OperatorSlash *OperatorSlashFilterer) WatchWithdrawalRequestContractSet(opts *bind.WatchOpts, sink chan<- *OperatorSlashWithdrawalRequestContractSet) (event.Subscription, error) {

	logs, sub, err := _OperatorSlash.contract.WatchLogs(opts, "WithdrawalRequestContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSlashWithdrawalRequestContractSet)
				if err := _OperatorSlash.contract.UnpackLog(event, "WithdrawalRequestContractSet", log); err != nil {
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

// ParseWithdrawalRequestContractSet is a log parse operation binding the contract event 0xb3b3e321ffd1930a33d425b4d1453792a16fca40d763a14c9fc90005360d0598.
//
// Solidity: event WithdrawalRequestContractSet(address _oldWithdrawalRequestContract, address _withdrawalRequestContractAddress)
func (_OperatorSlash *OperatorSlashFilterer) ParseWithdrawalRequestContractSet(log types.Log) (*OperatorSlashWithdrawalRequestContractSet, error) {
	event := new(OperatorSlashWithdrawalRequestContractSet)
	if err := _OperatorSlash.contract.UnpackLog(event, "WithdrawalRequestContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

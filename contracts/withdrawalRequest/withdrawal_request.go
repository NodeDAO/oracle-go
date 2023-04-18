// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package withdrawalRequest

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

// WithdrawalRequestWithdrawalInfo is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalRequestWithdrawalInfo struct {
	OperatorId         *big.Int
	WithdrawHeight     *big.Int
	WithdrawNethAmount *big.Int
	WithdrawExchange   *big.Int
	ClaimEthAmount     *big.Int
	Owner              common.Address
	IsClaim            bool
}

// WithdrawalRequestMetaData contains all meta data concerning the WithdrawalRequest contract.
var WithdrawalRequestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyUnstake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NethTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldDao\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"DaoAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPendingEthAmount\",\"type\":\"uint256\"}],\"name\":\"LargeWithdrawalsClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalNethAmount\",\"type\":\"uint256\"}],\"name\":\"LargeWithdrawalsRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldLiquidStakingContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_liquidStakingContractAddress\",\"type\":\"address\"}],\"name\":\"LiquidStakingChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NftUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldNodeOperatorRegistryContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistryContract\",\"type\":\"address\"}],\"name\":\"NodeOperatorRegistryContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldVaultManagerContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_vaultManagerContract\",\"type\":\"address\"}],\"name\":\"VaultManagerContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalsReceive\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_NETH_WITHDRAWAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_NETH_WITHDRAWAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_requestIds\",\"type\":\"uint256[]\"}],\"name\":\"claimLargeWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getNftUnstakeBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getOperatorLargeWithdrawalPendingInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPendingClaimedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getUserUnstakeButOperatorNoExitNfs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getWithdrawalOfOperator\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawNethAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawExchange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimEthAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isClaim\",\"type\":\"bool\"}],\"internalType\":\"structWithdrawalRequest.WithdrawalInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getWithdrawalOfRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getWithdrawalRequestIdOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidStakingAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nVNFTContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nETHContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultManagerContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidStakingContract\",\"outputs\":[{\"internalType\":\"contractILiquidStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nETHContract\",\"outputs\":[{\"internalType\":\"contractINETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeOperatorRegistryContract\",\"outputs\":[{\"internalType\":\"contractINodeOperatorsRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorPendingEthRequestAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"receiveWithdrawals\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"requestLargeWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"setDaoAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidStakingContractAddress\",\"type\":\"address\"}],\"name\":\"setLiquidStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistryContract\",\"type\":\"address\"}],\"name\":\"setNodeOperatorRegistryContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultManagerContract\",\"type\":\"address\"}],\"name\":\"setVaultManagerContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"unstakeNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vNFTContract\",\"outputs\":[{\"internalType\":\"contractIVNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultManagerContract\",\"outputs\":[{\"internalType\":\"contractIVaultManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WithdrawalRequestABI is the input ABI used to generate the binding from.
// Deprecated: Use WithdrawalRequestMetaData.ABI instead.
var WithdrawalRequestABI = WithdrawalRequestMetaData.ABI

// WithdrawalRequest is an auto generated Go binding around an Ethereum contract.
type WithdrawalRequest struct {
	WithdrawalRequestCaller     // Read-only binding to the contract
	WithdrawalRequestTransactor // Write-only binding to the contract
	WithdrawalRequestFilterer   // Log filterer for contract events
}

// WithdrawalRequestCaller is an auto generated read-only Go binding around an Ethereum contract.
type WithdrawalRequestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalRequestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WithdrawalRequestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalRequestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WithdrawalRequestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalRequestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WithdrawalRequestSession struct {
	Contract     *WithdrawalRequest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WithdrawalRequestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WithdrawalRequestCallerSession struct {
	Contract *WithdrawalRequestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// WithdrawalRequestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WithdrawalRequestTransactorSession struct {
	Contract     *WithdrawalRequestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// WithdrawalRequestRaw is an auto generated low-level Go binding around an Ethereum contract.
type WithdrawalRequestRaw struct {
	Contract *WithdrawalRequest // Generic contract binding to access the raw methods on
}

// WithdrawalRequestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WithdrawalRequestCallerRaw struct {
	Contract *WithdrawalRequestCaller // Generic read-only contract binding to access the raw methods on
}

// WithdrawalRequestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WithdrawalRequestTransactorRaw struct {
	Contract *WithdrawalRequestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWithdrawalRequest creates a new instance of WithdrawalRequest, bound to a specific deployed contract.
func NewWithdrawalRequest(address common.Address, backend bind.ContractBackend) (*WithdrawalRequest, error) {
	contract, err := bindWithdrawalRequest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequest{WithdrawalRequestCaller: WithdrawalRequestCaller{contract: contract}, WithdrawalRequestTransactor: WithdrawalRequestTransactor{contract: contract}, WithdrawalRequestFilterer: WithdrawalRequestFilterer{contract: contract}}, nil
}

// NewWithdrawalRequestCaller creates a new read-only instance of WithdrawalRequest, bound to a specific deployed contract.
func NewWithdrawalRequestCaller(address common.Address, caller bind.ContractCaller) (*WithdrawalRequestCaller, error) {
	contract, err := bindWithdrawalRequest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequestCaller{contract: contract}, nil
}

// NewWithdrawalRequestTransactor creates a new write-only instance of WithdrawalRequest, bound to a specific deployed contract.
func NewWithdrawalRequestTransactor(address common.Address, transactor bind.ContractTransactor) (*WithdrawalRequestTransactor, error) {
	contract, err := bindWithdrawalRequest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequestTransactor{contract: contract}, nil
}

// NewWithdrawalRequestFilterer creates a new log filterer instance of WithdrawalRequest, bound to a specific deployed contract.
func NewWithdrawalRequestFilterer(address common.Address, filterer bind.ContractFilterer) (*WithdrawalRequestFilterer, error) {
	contract, err := bindWithdrawalRequest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequestFilterer{contract: contract}, nil
}

// bindWithdrawalRequest binds a generic wrapper to an already deployed contract.
func bindWithdrawalRequest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WithdrawalRequestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WithdrawalRequest *WithdrawalRequestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WithdrawalRequest.Contract.WithdrawalRequestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WithdrawalRequest *WithdrawalRequestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.WithdrawalRequestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WithdrawalRequest *WithdrawalRequestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.WithdrawalRequestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WithdrawalRequest *WithdrawalRequestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WithdrawalRequest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WithdrawalRequest *WithdrawalRequestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WithdrawalRequest *WithdrawalRequestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.contract.Transact(opts, method, params...)
}

// MAXNETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x63823474.
//
// Solidity: function MAX_NETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_WithdrawalRequest *WithdrawalRequestCaller) MAXNETHWITHDRAWALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "MAX_NETH_WITHDRAWAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXNETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x63823474.
//
// Solidity: function MAX_NETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_WithdrawalRequest *WithdrawalRequestSession) MAXNETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _WithdrawalRequest.Contract.MAXNETHWITHDRAWALAMOUNT(&_WithdrawalRequest.CallOpts)
}

// MAXNETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x63823474.
//
// Solidity: function MAX_NETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_WithdrawalRequest *WithdrawalRequestCallerSession) MAXNETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _WithdrawalRequest.Contract.MAXNETHWITHDRAWALAMOUNT(&_WithdrawalRequest.CallOpts)
}

// MINNETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x24c5cba3.
//
// Solidity: function MIN_NETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_WithdrawalRequest *WithdrawalRequestCaller) MINNETHWITHDRAWALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "MIN_NETH_WITHDRAWAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINNETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x24c5cba3.
//
// Solidity: function MIN_NETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_WithdrawalRequest *WithdrawalRequestSession) MINNETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _WithdrawalRequest.Contract.MINNETHWITHDRAWALAMOUNT(&_WithdrawalRequest.CallOpts)
}

// MINNETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x24c5cba3.
//
// Solidity: function MIN_NETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_WithdrawalRequest *WithdrawalRequestCallerSession) MINNETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _WithdrawalRequest.Contract.MINNETHWITHDRAWALAMOUNT(&_WithdrawalRequest.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestSession) Dao() (common.Address, error) {
	return _WithdrawalRequest.Contract.Dao(&_WithdrawalRequest.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestCallerSession) Dao() (common.Address, error) {
	return _WithdrawalRequest.Contract.Dao(&_WithdrawalRequest.CallOpts)
}

// GetNftUnstakeBlockNumber is a free data retrieval call binding the contract method 0x41a8e1c6.
//
// Solidity: function getNftUnstakeBlockNumber(uint256 _tokenId) view returns(uint256)
func (_WithdrawalRequest *WithdrawalRequestCaller) GetNftUnstakeBlockNumber(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "getNftUnstakeBlockNumber", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNftUnstakeBlockNumber is a free data retrieval call binding the contract method 0x41a8e1c6.
//
// Solidity: function getNftUnstakeBlockNumber(uint256 _tokenId) view returns(uint256)
func (_WithdrawalRequest *WithdrawalRequestSession) GetNftUnstakeBlockNumber(_tokenId *big.Int) (*big.Int, error) {
	return _WithdrawalRequest.Contract.GetNftUnstakeBlockNumber(&_WithdrawalRequest.CallOpts, _tokenId)
}

// GetNftUnstakeBlockNumber is a free data retrieval call binding the contract method 0x41a8e1c6.
//
// Solidity: function getNftUnstakeBlockNumber(uint256 _tokenId) view returns(uint256)
func (_WithdrawalRequest *WithdrawalRequestCallerSession) GetNftUnstakeBlockNumber(_tokenId *big.Int) (*big.Int, error) {
	return _WithdrawalRequest.Contract.GetNftUnstakeBlockNumber(&_WithdrawalRequest.CallOpts, _tokenId)
}

// GetOperatorLargeWithdrawalPendingInfo is a free data retrieval call binding the contract method 0x99405f33.
//
// Solidity: function getOperatorLargeWithdrawalPendingInfo(uint256 _operatorId) view returns(uint256, uint256)
func (_WithdrawalRequest *WithdrawalRequestCaller) GetOperatorLargeWithdrawalPendingInfo(opts *bind.CallOpts, _operatorId *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "getOperatorLargeWithdrawalPendingInfo", _operatorId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetOperatorLargeWithdrawalPendingInfo is a free data retrieval call binding the contract method 0x99405f33.
//
// Solidity: function getOperatorLargeWithdrawalPendingInfo(uint256 _operatorId) view returns(uint256, uint256)
func (_WithdrawalRequest *WithdrawalRequestSession) GetOperatorLargeWithdrawalPendingInfo(_operatorId *big.Int) (*big.Int, *big.Int, error) {
	return _WithdrawalRequest.Contract.GetOperatorLargeWithdrawalPendingInfo(&_WithdrawalRequest.CallOpts, _operatorId)
}

// GetOperatorLargeWithdrawalPendingInfo is a free data retrieval call binding the contract method 0x99405f33.
//
// Solidity: function getOperatorLargeWithdrawalPendingInfo(uint256 _operatorId) view returns(uint256, uint256)
func (_WithdrawalRequest *WithdrawalRequestCallerSession) GetOperatorLargeWithdrawalPendingInfo(_operatorId *big.Int) (*big.Int, *big.Int, error) {
	return _WithdrawalRequest.Contract.GetOperatorLargeWithdrawalPendingInfo(&_WithdrawalRequest.CallOpts, _operatorId)
}

// GetTotalPendingClaimedAmounts is a free data retrieval call binding the contract method 0xfdc88efd.
//
// Solidity: function getTotalPendingClaimedAmounts() view returns(uint256)
func (_WithdrawalRequest *WithdrawalRequestCaller) GetTotalPendingClaimedAmounts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "getTotalPendingClaimedAmounts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPendingClaimedAmounts is a free data retrieval call binding the contract method 0xfdc88efd.
//
// Solidity: function getTotalPendingClaimedAmounts() view returns(uint256)
func (_WithdrawalRequest *WithdrawalRequestSession) GetTotalPendingClaimedAmounts() (*big.Int, error) {
	return _WithdrawalRequest.Contract.GetTotalPendingClaimedAmounts(&_WithdrawalRequest.CallOpts)
}

// GetTotalPendingClaimedAmounts is a free data retrieval call binding the contract method 0xfdc88efd.
//
// Solidity: function getTotalPendingClaimedAmounts() view returns(uint256)
func (_WithdrawalRequest *WithdrawalRequestCallerSession) GetTotalPendingClaimedAmounts() (*big.Int, error) {
	return _WithdrawalRequest.Contract.GetTotalPendingClaimedAmounts(&_WithdrawalRequest.CallOpts)
}

// GetUserUnstakeButOperatorNoExitNfs is a free data retrieval call binding the contract method 0xf1c28381.
//
// Solidity: function getUserUnstakeButOperatorNoExitNfs(uint256 _operatorId) view returns(uint256[])
func (_WithdrawalRequest *WithdrawalRequestCaller) GetUserUnstakeButOperatorNoExitNfs(opts *bind.CallOpts, _operatorId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "getUserUnstakeButOperatorNoExitNfs", _operatorId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserUnstakeButOperatorNoExitNfs is a free data retrieval call binding the contract method 0xf1c28381.
//
// Solidity: function getUserUnstakeButOperatorNoExitNfs(uint256 _operatorId) view returns(uint256[])
func (_WithdrawalRequest *WithdrawalRequestSession) GetUserUnstakeButOperatorNoExitNfs(_operatorId *big.Int) ([]*big.Int, error) {
	return _WithdrawalRequest.Contract.GetUserUnstakeButOperatorNoExitNfs(&_WithdrawalRequest.CallOpts, _operatorId)
}

// GetUserUnstakeButOperatorNoExitNfs is a free data retrieval call binding the contract method 0xf1c28381.
//
// Solidity: function getUserUnstakeButOperatorNoExitNfs(uint256 _operatorId) view returns(uint256[])
func (_WithdrawalRequest *WithdrawalRequestCallerSession) GetUserUnstakeButOperatorNoExitNfs(_operatorId *big.Int) ([]*big.Int, error) {
	return _WithdrawalRequest.Contract.GetUserUnstakeButOperatorNoExitNfs(&_WithdrawalRequest.CallOpts, _operatorId)
}

// GetWithdrawalOfOperator is a free data retrieval call binding the contract method 0x46a2ca47.
//
// Solidity: function getWithdrawalOfOperator(uint256 _operatorId) view returns((uint256,uint256,uint256,uint256,uint256,address,bool)[])
func (_WithdrawalRequest *WithdrawalRequestCaller) GetWithdrawalOfOperator(opts *bind.CallOpts, _operatorId *big.Int) ([]WithdrawalRequestWithdrawalInfo, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "getWithdrawalOfOperator", _operatorId)

	if err != nil {
		return *new([]WithdrawalRequestWithdrawalInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]WithdrawalRequestWithdrawalInfo)).(*[]WithdrawalRequestWithdrawalInfo)

	return out0, err

}

// GetWithdrawalOfOperator is a free data retrieval call binding the contract method 0x46a2ca47.
//
// Solidity: function getWithdrawalOfOperator(uint256 _operatorId) view returns((uint256,uint256,uint256,uint256,uint256,address,bool)[])
func (_WithdrawalRequest *WithdrawalRequestSession) GetWithdrawalOfOperator(_operatorId *big.Int) ([]WithdrawalRequestWithdrawalInfo, error) {
	return _WithdrawalRequest.Contract.GetWithdrawalOfOperator(&_WithdrawalRequest.CallOpts, _operatorId)
}

// GetWithdrawalOfOperator is a free data retrieval call binding the contract method 0x46a2ca47.
//
// Solidity: function getWithdrawalOfOperator(uint256 _operatorId) view returns((uint256,uint256,uint256,uint256,uint256,address,bool)[])
func (_WithdrawalRequest *WithdrawalRequestCallerSession) GetWithdrawalOfOperator(_operatorId *big.Int) ([]WithdrawalRequestWithdrawalInfo, error) {
	return _WithdrawalRequest.Contract.GetWithdrawalOfOperator(&_WithdrawalRequest.CallOpts, _operatorId)
}

// GetWithdrawalOfRequestId is a free data retrieval call binding the contract method 0xdeb38284.
//
// Solidity: function getWithdrawalOfRequestId(uint256 _requestId) view returns(uint256, uint256, uint256, uint256, uint256, address, bool)
func (_WithdrawalRequest *WithdrawalRequestCaller) GetWithdrawalOfRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, bool, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "getWithdrawalOfRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	out6 := *abi.ConvertType(out[6], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// GetWithdrawalOfRequestId is a free data retrieval call binding the contract method 0xdeb38284.
//
// Solidity: function getWithdrawalOfRequestId(uint256 _requestId) view returns(uint256, uint256, uint256, uint256, uint256, address, bool)
func (_WithdrawalRequest *WithdrawalRequestSession) GetWithdrawalOfRequestId(_requestId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, bool, error) {
	return _WithdrawalRequest.Contract.GetWithdrawalOfRequestId(&_WithdrawalRequest.CallOpts, _requestId)
}

// GetWithdrawalOfRequestId is a free data retrieval call binding the contract method 0xdeb38284.
//
// Solidity: function getWithdrawalOfRequestId(uint256 _requestId) view returns(uint256, uint256, uint256, uint256, uint256, address, bool)
func (_WithdrawalRequest *WithdrawalRequestCallerSession) GetWithdrawalOfRequestId(_requestId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, bool, error) {
	return _WithdrawalRequest.Contract.GetWithdrawalOfRequestId(&_WithdrawalRequest.CallOpts, _requestId)
}

// GetWithdrawalRequestIdOfOwner is a free data retrieval call binding the contract method 0xcdc2b1c7.
//
// Solidity: function getWithdrawalRequestIdOfOwner(address _owner) view returns(uint256[])
func (_WithdrawalRequest *WithdrawalRequestCaller) GetWithdrawalRequestIdOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "getWithdrawalRequestIdOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetWithdrawalRequestIdOfOwner is a free data retrieval call binding the contract method 0xcdc2b1c7.
//
// Solidity: function getWithdrawalRequestIdOfOwner(address _owner) view returns(uint256[])
func (_WithdrawalRequest *WithdrawalRequestSession) GetWithdrawalRequestIdOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _WithdrawalRequest.Contract.GetWithdrawalRequestIdOfOwner(&_WithdrawalRequest.CallOpts, _owner)
}

// GetWithdrawalRequestIdOfOwner is a free data retrieval call binding the contract method 0xcdc2b1c7.
//
// Solidity: function getWithdrawalRequestIdOfOwner(address _owner) view returns(uint256[])
func (_WithdrawalRequest *WithdrawalRequestCallerSession) GetWithdrawalRequestIdOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _WithdrawalRequest.Contract.GetWithdrawalRequestIdOfOwner(&_WithdrawalRequest.CallOpts, _owner)
}

// LiquidStakingContract is a free data retrieval call binding the contract method 0xbdcaa355.
//
// Solidity: function liquidStakingContract() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestCaller) LiquidStakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "liquidStakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidStakingContract is a free data retrieval call binding the contract method 0xbdcaa355.
//
// Solidity: function liquidStakingContract() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestSession) LiquidStakingContract() (common.Address, error) {
	return _WithdrawalRequest.Contract.LiquidStakingContract(&_WithdrawalRequest.CallOpts)
}

// LiquidStakingContract is a free data retrieval call binding the contract method 0xbdcaa355.
//
// Solidity: function liquidStakingContract() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestCallerSession) LiquidStakingContract() (common.Address, error) {
	return _WithdrawalRequest.Contract.LiquidStakingContract(&_WithdrawalRequest.CallOpts)
}

// NETHContract is a free data retrieval call binding the contract method 0xbd97c759.
//
// Solidity: function nETHContract() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestCaller) NETHContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "nETHContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NETHContract is a free data retrieval call binding the contract method 0xbd97c759.
//
// Solidity: function nETHContract() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestSession) NETHContract() (common.Address, error) {
	return _WithdrawalRequest.Contract.NETHContract(&_WithdrawalRequest.CallOpts)
}

// NETHContract is a free data retrieval call binding the contract method 0xbd97c759.
//
// Solidity: function nETHContract() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestCallerSession) NETHContract() (common.Address, error) {
	return _WithdrawalRequest.Contract.NETHContract(&_WithdrawalRequest.CallOpts)
}

// NodeOperatorRegistryContract is a free data retrieval call binding the contract method 0x9a42e0ba.
//
// Solidity: function nodeOperatorRegistryContract() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestCaller) NodeOperatorRegistryContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "nodeOperatorRegistryContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeOperatorRegistryContract is a free data retrieval call binding the contract method 0x9a42e0ba.
//
// Solidity: function nodeOperatorRegistryContract() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestSession) NodeOperatorRegistryContract() (common.Address, error) {
	return _WithdrawalRequest.Contract.NodeOperatorRegistryContract(&_WithdrawalRequest.CallOpts)
}

// NodeOperatorRegistryContract is a free data retrieval call binding the contract method 0x9a42e0ba.
//
// Solidity: function nodeOperatorRegistryContract() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestCallerSession) NodeOperatorRegistryContract() (common.Address, error) {
	return _WithdrawalRequest.Contract.NodeOperatorRegistryContract(&_WithdrawalRequest.CallOpts)
}

// OperatorPendingEthRequestAmount is a free data retrieval call binding the contract method 0x67dd7ffd.
//
// Solidity: function operatorPendingEthRequestAmount(uint256 ) view returns(uint256)
func (_WithdrawalRequest *WithdrawalRequestCaller) OperatorPendingEthRequestAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "operatorPendingEthRequestAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorPendingEthRequestAmount is a free data retrieval call binding the contract method 0x67dd7ffd.
//
// Solidity: function operatorPendingEthRequestAmount(uint256 ) view returns(uint256)
func (_WithdrawalRequest *WithdrawalRequestSession) OperatorPendingEthRequestAmount(arg0 *big.Int) (*big.Int, error) {
	return _WithdrawalRequest.Contract.OperatorPendingEthRequestAmount(&_WithdrawalRequest.CallOpts, arg0)
}

// OperatorPendingEthRequestAmount is a free data retrieval call binding the contract method 0x67dd7ffd.
//
// Solidity: function operatorPendingEthRequestAmount(uint256 ) view returns(uint256)
func (_WithdrawalRequest *WithdrawalRequestCallerSession) OperatorPendingEthRequestAmount(arg0 *big.Int) (*big.Int, error) {
	return _WithdrawalRequest.Contract.OperatorPendingEthRequestAmount(&_WithdrawalRequest.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestSession) Owner() (common.Address, error) {
	return _WithdrawalRequest.Contract.Owner(&_WithdrawalRequest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestCallerSession) Owner() (common.Address, error) {
	return _WithdrawalRequest.Contract.Owner(&_WithdrawalRequest.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WithdrawalRequest *WithdrawalRequestCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WithdrawalRequest *WithdrawalRequestSession) Paused() (bool, error) {
	return _WithdrawalRequest.Contract.Paused(&_WithdrawalRequest.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WithdrawalRequest *WithdrawalRequestCallerSession) Paused() (bool, error) {
	return _WithdrawalRequest.Contract.Paused(&_WithdrawalRequest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WithdrawalRequest *WithdrawalRequestCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WithdrawalRequest *WithdrawalRequestSession) ProxiableUUID() ([32]byte, error) {
	return _WithdrawalRequest.Contract.ProxiableUUID(&_WithdrawalRequest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WithdrawalRequest *WithdrawalRequestCallerSession) ProxiableUUID() ([32]byte, error) {
	return _WithdrawalRequest.Contract.ProxiableUUID(&_WithdrawalRequest.CallOpts)
}

// VNFTContract is a free data retrieval call binding the contract method 0xfc03411f.
//
// Solidity: function vNFTContract() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestCaller) VNFTContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "vNFTContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VNFTContract is a free data retrieval call binding the contract method 0xfc03411f.
//
// Solidity: function vNFTContract() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestSession) VNFTContract() (common.Address, error) {
	return _WithdrawalRequest.Contract.VNFTContract(&_WithdrawalRequest.CallOpts)
}

// VNFTContract is a free data retrieval call binding the contract method 0xfc03411f.
//
// Solidity: function vNFTContract() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestCallerSession) VNFTContract() (common.Address, error) {
	return _WithdrawalRequest.Contract.VNFTContract(&_WithdrawalRequest.CallOpts)
}

// VaultManagerContract is a free data retrieval call binding the contract method 0xedbb92f4.
//
// Solidity: function vaultManagerContract() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestCaller) VaultManagerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WithdrawalRequest.contract.Call(opts, &out, "vaultManagerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultManagerContract is a free data retrieval call binding the contract method 0xedbb92f4.
//
// Solidity: function vaultManagerContract() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestSession) VaultManagerContract() (common.Address, error) {
	return _WithdrawalRequest.Contract.VaultManagerContract(&_WithdrawalRequest.CallOpts)
}

// VaultManagerContract is a free data retrieval call binding the contract method 0xedbb92f4.
//
// Solidity: function vaultManagerContract() view returns(address)
func (_WithdrawalRequest *WithdrawalRequestCallerSession) VaultManagerContract() (common.Address, error) {
	return _WithdrawalRequest.Contract.VaultManagerContract(&_WithdrawalRequest.CallOpts)
}

// ClaimLargeWithdrawals is a paid mutator transaction binding the contract method 0x1231797a.
//
// Solidity: function claimLargeWithdrawals(uint256[] _requestIds) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactor) ClaimLargeWithdrawals(opts *bind.TransactOpts, _requestIds []*big.Int) (*types.Transaction, error) {
	return _WithdrawalRequest.contract.Transact(opts, "claimLargeWithdrawals", _requestIds)
}

// ClaimLargeWithdrawals is a paid mutator transaction binding the contract method 0x1231797a.
//
// Solidity: function claimLargeWithdrawals(uint256[] _requestIds) returns()
func (_WithdrawalRequest *WithdrawalRequestSession) ClaimLargeWithdrawals(_requestIds []*big.Int) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.ClaimLargeWithdrawals(&_WithdrawalRequest.TransactOpts, _requestIds)
}

// ClaimLargeWithdrawals is a paid mutator transaction binding the contract method 0x1231797a.
//
// Solidity: function claimLargeWithdrawals(uint256[] _requestIds) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactorSession) ClaimLargeWithdrawals(_requestIds []*big.Int) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.ClaimLargeWithdrawals(&_WithdrawalRequest.TransactOpts, _requestIds)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _dao, address _liquidStakingAddress, address _nVNFTContractAddress, address _nETHContractAddress, address _nodeOperatorRegistryAddress, address _vaultManagerContract) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactor) Initialize(opts *bind.TransactOpts, _dao common.Address, _liquidStakingAddress common.Address, _nVNFTContractAddress common.Address, _nETHContractAddress common.Address, _nodeOperatorRegistryAddress common.Address, _vaultManagerContract common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.contract.Transact(opts, "initialize", _dao, _liquidStakingAddress, _nVNFTContractAddress, _nETHContractAddress, _nodeOperatorRegistryAddress, _vaultManagerContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _dao, address _liquidStakingAddress, address _nVNFTContractAddress, address _nETHContractAddress, address _nodeOperatorRegistryAddress, address _vaultManagerContract) returns()
func (_WithdrawalRequest *WithdrawalRequestSession) Initialize(_dao common.Address, _liquidStakingAddress common.Address, _nVNFTContractAddress common.Address, _nETHContractAddress common.Address, _nodeOperatorRegistryAddress common.Address, _vaultManagerContract common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.Initialize(&_WithdrawalRequest.TransactOpts, _dao, _liquidStakingAddress, _nVNFTContractAddress, _nETHContractAddress, _nodeOperatorRegistryAddress, _vaultManagerContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _dao, address _liquidStakingAddress, address _nVNFTContractAddress, address _nETHContractAddress, address _nodeOperatorRegistryAddress, address _vaultManagerContract) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactorSession) Initialize(_dao common.Address, _liquidStakingAddress common.Address, _nVNFTContractAddress common.Address, _nETHContractAddress common.Address, _nodeOperatorRegistryAddress common.Address, _vaultManagerContract common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.Initialize(&_WithdrawalRequest.TransactOpts, _dao, _liquidStakingAddress, _nVNFTContractAddress, _nETHContractAddress, _nodeOperatorRegistryAddress, _vaultManagerContract)
}

// ReceiveWithdrawals is a paid mutator transaction binding the contract method 0xd3d2b2cb.
//
// Solidity: function receiveWithdrawals(uint256 _operatorId, uint256 _amount) payable returns()
func (_WithdrawalRequest *WithdrawalRequestTransactor) ReceiveWithdrawals(opts *bind.TransactOpts, _operatorId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _WithdrawalRequest.contract.Transact(opts, "receiveWithdrawals", _operatorId, _amount)
}

// ReceiveWithdrawals is a paid mutator transaction binding the contract method 0xd3d2b2cb.
//
// Solidity: function receiveWithdrawals(uint256 _operatorId, uint256 _amount) payable returns()
func (_WithdrawalRequest *WithdrawalRequestSession) ReceiveWithdrawals(_operatorId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.ReceiveWithdrawals(&_WithdrawalRequest.TransactOpts, _operatorId, _amount)
}

// ReceiveWithdrawals is a paid mutator transaction binding the contract method 0xd3d2b2cb.
//
// Solidity: function receiveWithdrawals(uint256 _operatorId, uint256 _amount) payable returns()
func (_WithdrawalRequest *WithdrawalRequestTransactorSession) ReceiveWithdrawals(_operatorId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.ReceiveWithdrawals(&_WithdrawalRequest.TransactOpts, _operatorId, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WithdrawalRequest *WithdrawalRequestTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawalRequest.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WithdrawalRequest *WithdrawalRequestSession) RenounceOwnership() (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.RenounceOwnership(&_WithdrawalRequest.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WithdrawalRequest *WithdrawalRequestTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.RenounceOwnership(&_WithdrawalRequest.TransactOpts)
}

// RequestLargeWithdrawals is a paid mutator transaction binding the contract method 0x976b364e.
//
// Solidity: function requestLargeWithdrawals(uint256 _operatorId, uint256[] _amounts) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactor) RequestLargeWithdrawals(opts *bind.TransactOpts, _operatorId *big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _WithdrawalRequest.contract.Transact(opts, "requestLargeWithdrawals", _operatorId, _amounts)
}

// RequestLargeWithdrawals is a paid mutator transaction binding the contract method 0x976b364e.
//
// Solidity: function requestLargeWithdrawals(uint256 _operatorId, uint256[] _amounts) returns()
func (_WithdrawalRequest *WithdrawalRequestSession) RequestLargeWithdrawals(_operatorId *big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.RequestLargeWithdrawals(&_WithdrawalRequest.TransactOpts, _operatorId, _amounts)
}

// RequestLargeWithdrawals is a paid mutator transaction binding the contract method 0x976b364e.
//
// Solidity: function requestLargeWithdrawals(uint256 _operatorId, uint256[] _amounts) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactorSession) RequestLargeWithdrawals(_operatorId *big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.RequestLargeWithdrawals(&_WithdrawalRequest.TransactOpts, _operatorId, _amounts)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactor) SetDaoAddress(opts *bind.TransactOpts, _dao common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.contract.Transact(opts, "setDaoAddress", _dao)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_WithdrawalRequest *WithdrawalRequestSession) SetDaoAddress(_dao common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.SetDaoAddress(&_WithdrawalRequest.TransactOpts, _dao)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactorSession) SetDaoAddress(_dao common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.SetDaoAddress(&_WithdrawalRequest.TransactOpts, _dao)
}

// SetLiquidStaking is a paid mutator transaction binding the contract method 0x08211be5.
//
// Solidity: function setLiquidStaking(address _liquidStakingContractAddress) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactor) SetLiquidStaking(opts *bind.TransactOpts, _liquidStakingContractAddress common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.contract.Transact(opts, "setLiquidStaking", _liquidStakingContractAddress)
}

// SetLiquidStaking is a paid mutator transaction binding the contract method 0x08211be5.
//
// Solidity: function setLiquidStaking(address _liquidStakingContractAddress) returns()
func (_WithdrawalRequest *WithdrawalRequestSession) SetLiquidStaking(_liquidStakingContractAddress common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.SetLiquidStaking(&_WithdrawalRequest.TransactOpts, _liquidStakingContractAddress)
}

// SetLiquidStaking is a paid mutator transaction binding the contract method 0x08211be5.
//
// Solidity: function setLiquidStaking(address _liquidStakingContractAddress) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactorSession) SetLiquidStaking(_liquidStakingContractAddress common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.SetLiquidStaking(&_WithdrawalRequest.TransactOpts, _liquidStakingContractAddress)
}

// SetNodeOperatorRegistryContract is a paid mutator transaction binding the contract method 0xcb23473e.
//
// Solidity: function setNodeOperatorRegistryContract(address _nodeOperatorRegistryContract) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactor) SetNodeOperatorRegistryContract(opts *bind.TransactOpts, _nodeOperatorRegistryContract common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.contract.Transact(opts, "setNodeOperatorRegistryContract", _nodeOperatorRegistryContract)
}

// SetNodeOperatorRegistryContract is a paid mutator transaction binding the contract method 0xcb23473e.
//
// Solidity: function setNodeOperatorRegistryContract(address _nodeOperatorRegistryContract) returns()
func (_WithdrawalRequest *WithdrawalRequestSession) SetNodeOperatorRegistryContract(_nodeOperatorRegistryContract common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.SetNodeOperatorRegistryContract(&_WithdrawalRequest.TransactOpts, _nodeOperatorRegistryContract)
}

// SetNodeOperatorRegistryContract is a paid mutator transaction binding the contract method 0xcb23473e.
//
// Solidity: function setNodeOperatorRegistryContract(address _nodeOperatorRegistryContract) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactorSession) SetNodeOperatorRegistryContract(_nodeOperatorRegistryContract common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.SetNodeOperatorRegistryContract(&_WithdrawalRequest.TransactOpts, _nodeOperatorRegistryContract)
}

// SetVaultManagerContract is a paid mutator transaction binding the contract method 0xbe14652a.
//
// Solidity: function setVaultManagerContract(address _vaultManagerContract) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactor) SetVaultManagerContract(opts *bind.TransactOpts, _vaultManagerContract common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.contract.Transact(opts, "setVaultManagerContract", _vaultManagerContract)
}

// SetVaultManagerContract is a paid mutator transaction binding the contract method 0xbe14652a.
//
// Solidity: function setVaultManagerContract(address _vaultManagerContract) returns()
func (_WithdrawalRequest *WithdrawalRequestSession) SetVaultManagerContract(_vaultManagerContract common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.SetVaultManagerContract(&_WithdrawalRequest.TransactOpts, _vaultManagerContract)
}

// SetVaultManagerContract is a paid mutator transaction binding the contract method 0xbe14652a.
//
// Solidity: function setVaultManagerContract(address _vaultManagerContract) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactorSession) SetVaultManagerContract(_vaultManagerContract common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.SetVaultManagerContract(&_WithdrawalRequest.TransactOpts, _vaultManagerContract)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WithdrawalRequest *WithdrawalRequestSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.TransferOwnership(&_WithdrawalRequest.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.TransferOwnership(&_WithdrawalRequest.TransactOpts, newOwner)
}

// UnstakeNFT is a paid mutator transaction binding the contract method 0x8453734d.
//
// Solidity: function unstakeNFT(uint256[] _tokenIds) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactor) UnstakeNFT(opts *bind.TransactOpts, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _WithdrawalRequest.contract.Transact(opts, "unstakeNFT", _tokenIds)
}

// UnstakeNFT is a paid mutator transaction binding the contract method 0x8453734d.
//
// Solidity: function unstakeNFT(uint256[] _tokenIds) returns()
func (_WithdrawalRequest *WithdrawalRequestSession) UnstakeNFT(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.UnstakeNFT(&_WithdrawalRequest.TransactOpts, _tokenIds)
}

// UnstakeNFT is a paid mutator transaction binding the contract method 0x8453734d.
//
// Solidity: function unstakeNFT(uint256[] _tokenIds) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactorSession) UnstakeNFT(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.UnstakeNFT(&_WithdrawalRequest.TransactOpts, _tokenIds)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WithdrawalRequest *WithdrawalRequestSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.UpgradeTo(&_WithdrawalRequest.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WithdrawalRequest *WithdrawalRequestTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.UpgradeTo(&_WithdrawalRequest.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WithdrawalRequest *WithdrawalRequestTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WithdrawalRequest.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WithdrawalRequest *WithdrawalRequestSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.UpgradeToAndCall(&_WithdrawalRequest.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WithdrawalRequest *WithdrawalRequestTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WithdrawalRequest.Contract.UpgradeToAndCall(&_WithdrawalRequest.TransactOpts, newImplementation, data)
}

// WithdrawalRequestAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the WithdrawalRequest contract.
type WithdrawalRequestAdminChangedIterator struct {
	Event *WithdrawalRequestAdminChanged // Event containing the contract specifics and raw log

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
func (it *WithdrawalRequestAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalRequestAdminChanged)
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
		it.Event = new(WithdrawalRequestAdminChanged)
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
func (it *WithdrawalRequestAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalRequestAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalRequestAdminChanged represents a AdminChanged event raised by the WithdrawalRequest contract.
type WithdrawalRequestAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_WithdrawalRequest *WithdrawalRequestFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*WithdrawalRequestAdminChangedIterator, error) {

	logs, sub, err := _WithdrawalRequest.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequestAdminChangedIterator{contract: _WithdrawalRequest.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_WithdrawalRequest *WithdrawalRequestFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *WithdrawalRequestAdminChanged) (event.Subscription, error) {

	logs, sub, err := _WithdrawalRequest.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalRequestAdminChanged)
				if err := _WithdrawalRequest.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_WithdrawalRequest *WithdrawalRequestFilterer) ParseAdminChanged(log types.Log) (*WithdrawalRequestAdminChanged, error) {
	event := new(WithdrawalRequestAdminChanged)
	if err := _WithdrawalRequest.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalRequestBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the WithdrawalRequest contract.
type WithdrawalRequestBeaconUpgradedIterator struct {
	Event *WithdrawalRequestBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *WithdrawalRequestBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalRequestBeaconUpgraded)
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
		it.Event = new(WithdrawalRequestBeaconUpgraded)
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
func (it *WithdrawalRequestBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalRequestBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalRequestBeaconUpgraded represents a BeaconUpgraded event raised by the WithdrawalRequest contract.
type WithdrawalRequestBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_WithdrawalRequest *WithdrawalRequestFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*WithdrawalRequestBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _WithdrawalRequest.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequestBeaconUpgradedIterator{contract: _WithdrawalRequest.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_WithdrawalRequest *WithdrawalRequestFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *WithdrawalRequestBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _WithdrawalRequest.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalRequestBeaconUpgraded)
				if err := _WithdrawalRequest.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_WithdrawalRequest *WithdrawalRequestFilterer) ParseBeaconUpgraded(log types.Log) (*WithdrawalRequestBeaconUpgraded, error) {
	event := new(WithdrawalRequestBeaconUpgraded)
	if err := _WithdrawalRequest.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalRequestDaoAddressChangedIterator is returned from FilterDaoAddressChanged and is used to iterate over the raw logs and unpacked data for DaoAddressChanged events raised by the WithdrawalRequest contract.
type WithdrawalRequestDaoAddressChangedIterator struct {
	Event *WithdrawalRequestDaoAddressChanged // Event containing the contract specifics and raw log

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
func (it *WithdrawalRequestDaoAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalRequestDaoAddressChanged)
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
		it.Event = new(WithdrawalRequestDaoAddressChanged)
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
func (it *WithdrawalRequestDaoAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalRequestDaoAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalRequestDaoAddressChanged represents a DaoAddressChanged event raised by the WithdrawalRequest contract.
type WithdrawalRequestDaoAddressChanged struct {
	OldDao common.Address
	Dao    common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDaoAddressChanged is a free log retrieval operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address oldDao, address _dao)
func (_WithdrawalRequest *WithdrawalRequestFilterer) FilterDaoAddressChanged(opts *bind.FilterOpts) (*WithdrawalRequestDaoAddressChangedIterator, error) {

	logs, sub, err := _WithdrawalRequest.contract.FilterLogs(opts, "DaoAddressChanged")
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequestDaoAddressChangedIterator{contract: _WithdrawalRequest.contract, event: "DaoAddressChanged", logs: logs, sub: sub}, nil
}

// WatchDaoAddressChanged is a free log subscription operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address oldDao, address _dao)
func (_WithdrawalRequest *WithdrawalRequestFilterer) WatchDaoAddressChanged(opts *bind.WatchOpts, sink chan<- *WithdrawalRequestDaoAddressChanged) (event.Subscription, error) {

	logs, sub, err := _WithdrawalRequest.contract.WatchLogs(opts, "DaoAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalRequestDaoAddressChanged)
				if err := _WithdrawalRequest.contract.UnpackLog(event, "DaoAddressChanged", log); err != nil {
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
func (_WithdrawalRequest *WithdrawalRequestFilterer) ParseDaoAddressChanged(log types.Log) (*WithdrawalRequestDaoAddressChanged, error) {
	event := new(WithdrawalRequestDaoAddressChanged)
	if err := _WithdrawalRequest.contract.UnpackLog(event, "DaoAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalRequestInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WithdrawalRequest contract.
type WithdrawalRequestInitializedIterator struct {
	Event *WithdrawalRequestInitialized // Event containing the contract specifics and raw log

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
func (it *WithdrawalRequestInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalRequestInitialized)
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
		it.Event = new(WithdrawalRequestInitialized)
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
func (it *WithdrawalRequestInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalRequestInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalRequestInitialized represents a Initialized event raised by the WithdrawalRequest contract.
type WithdrawalRequestInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WithdrawalRequest *WithdrawalRequestFilterer) FilterInitialized(opts *bind.FilterOpts) (*WithdrawalRequestInitializedIterator, error) {

	logs, sub, err := _WithdrawalRequest.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequestInitializedIterator{contract: _WithdrawalRequest.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WithdrawalRequest *WithdrawalRequestFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WithdrawalRequestInitialized) (event.Subscription, error) {

	logs, sub, err := _WithdrawalRequest.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalRequestInitialized)
				if err := _WithdrawalRequest.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_WithdrawalRequest *WithdrawalRequestFilterer) ParseInitialized(log types.Log) (*WithdrawalRequestInitialized, error) {
	event := new(WithdrawalRequestInitialized)
	if err := _WithdrawalRequest.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalRequestLargeWithdrawalsClaimIterator is returned from FilterLargeWithdrawalsClaim and is used to iterate over the raw logs and unpacked data for LargeWithdrawalsClaim events raised by the WithdrawalRequest contract.
type WithdrawalRequestLargeWithdrawalsClaimIterator struct {
	Event *WithdrawalRequestLargeWithdrawalsClaim // Event containing the contract specifics and raw log

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
func (it *WithdrawalRequestLargeWithdrawalsClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalRequestLargeWithdrawalsClaim)
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
		it.Event = new(WithdrawalRequestLargeWithdrawalsClaim)
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
func (it *WithdrawalRequestLargeWithdrawalsClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalRequestLargeWithdrawalsClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalRequestLargeWithdrawalsClaim represents a LargeWithdrawalsClaim event raised by the WithdrawalRequest contract.
type WithdrawalRequestLargeWithdrawalsClaim struct {
	Sender                common.Address
	TotalPendingEthAmount *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterLargeWithdrawalsClaim is a free log retrieval operation binding the contract event 0xd3612323e05886ab09a8d484fe950e624084dc8836fcb83402517e9f473a1e53.
//
// Solidity: event LargeWithdrawalsClaim(address sender, uint256 totalPendingEthAmount)
func (_WithdrawalRequest *WithdrawalRequestFilterer) FilterLargeWithdrawalsClaim(opts *bind.FilterOpts) (*WithdrawalRequestLargeWithdrawalsClaimIterator, error) {

	logs, sub, err := _WithdrawalRequest.contract.FilterLogs(opts, "LargeWithdrawalsClaim")
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequestLargeWithdrawalsClaimIterator{contract: _WithdrawalRequest.contract, event: "LargeWithdrawalsClaim", logs: logs, sub: sub}, nil
}

// WatchLargeWithdrawalsClaim is a free log subscription operation binding the contract event 0xd3612323e05886ab09a8d484fe950e624084dc8836fcb83402517e9f473a1e53.
//
// Solidity: event LargeWithdrawalsClaim(address sender, uint256 totalPendingEthAmount)
func (_WithdrawalRequest *WithdrawalRequestFilterer) WatchLargeWithdrawalsClaim(opts *bind.WatchOpts, sink chan<- *WithdrawalRequestLargeWithdrawalsClaim) (event.Subscription, error) {

	logs, sub, err := _WithdrawalRequest.contract.WatchLogs(opts, "LargeWithdrawalsClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalRequestLargeWithdrawalsClaim)
				if err := _WithdrawalRequest.contract.UnpackLog(event, "LargeWithdrawalsClaim", log); err != nil {
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

// ParseLargeWithdrawalsClaim is a log parse operation binding the contract event 0xd3612323e05886ab09a8d484fe950e624084dc8836fcb83402517e9f473a1e53.
//
// Solidity: event LargeWithdrawalsClaim(address sender, uint256 totalPendingEthAmount)
func (_WithdrawalRequest *WithdrawalRequestFilterer) ParseLargeWithdrawalsClaim(log types.Log) (*WithdrawalRequestLargeWithdrawalsClaim, error) {
	event := new(WithdrawalRequestLargeWithdrawalsClaim)
	if err := _WithdrawalRequest.contract.UnpackLog(event, "LargeWithdrawalsClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalRequestLargeWithdrawalsRequestIterator is returned from FilterLargeWithdrawalsRequest and is used to iterate over the raw logs and unpacked data for LargeWithdrawalsRequest events raised by the WithdrawalRequest contract.
type WithdrawalRequestLargeWithdrawalsRequestIterator struct {
	Event *WithdrawalRequestLargeWithdrawalsRequest // Event containing the contract specifics and raw log

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
func (it *WithdrawalRequestLargeWithdrawalsRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalRequestLargeWithdrawalsRequest)
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
		it.Event = new(WithdrawalRequestLargeWithdrawalsRequest)
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
func (it *WithdrawalRequestLargeWithdrawalsRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalRequestLargeWithdrawalsRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalRequestLargeWithdrawalsRequest represents a LargeWithdrawalsRequest event raised by the WithdrawalRequest contract.
type WithdrawalRequestLargeWithdrawalsRequest struct {
	OperatorId      *big.Int
	Sender          common.Address
	TotalNethAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLargeWithdrawalsRequest is a free log retrieval operation binding the contract event 0x5e24974b0e489bf247cedb3c7c52d992a262766293923a8e3eba888371495a3f.
//
// Solidity: event LargeWithdrawalsRequest(uint256 _operatorId, address sender, uint256 totalNethAmount)
func (_WithdrawalRequest *WithdrawalRequestFilterer) FilterLargeWithdrawalsRequest(opts *bind.FilterOpts) (*WithdrawalRequestLargeWithdrawalsRequestIterator, error) {

	logs, sub, err := _WithdrawalRequest.contract.FilterLogs(opts, "LargeWithdrawalsRequest")
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequestLargeWithdrawalsRequestIterator{contract: _WithdrawalRequest.contract, event: "LargeWithdrawalsRequest", logs: logs, sub: sub}, nil
}

// WatchLargeWithdrawalsRequest is a free log subscription operation binding the contract event 0x5e24974b0e489bf247cedb3c7c52d992a262766293923a8e3eba888371495a3f.
//
// Solidity: event LargeWithdrawalsRequest(uint256 _operatorId, address sender, uint256 totalNethAmount)
func (_WithdrawalRequest *WithdrawalRequestFilterer) WatchLargeWithdrawalsRequest(opts *bind.WatchOpts, sink chan<- *WithdrawalRequestLargeWithdrawalsRequest) (event.Subscription, error) {

	logs, sub, err := _WithdrawalRequest.contract.WatchLogs(opts, "LargeWithdrawalsRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalRequestLargeWithdrawalsRequest)
				if err := _WithdrawalRequest.contract.UnpackLog(event, "LargeWithdrawalsRequest", log); err != nil {
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

// ParseLargeWithdrawalsRequest is a log parse operation binding the contract event 0x5e24974b0e489bf247cedb3c7c52d992a262766293923a8e3eba888371495a3f.
//
// Solidity: event LargeWithdrawalsRequest(uint256 _operatorId, address sender, uint256 totalNethAmount)
func (_WithdrawalRequest *WithdrawalRequestFilterer) ParseLargeWithdrawalsRequest(log types.Log) (*WithdrawalRequestLargeWithdrawalsRequest, error) {
	event := new(WithdrawalRequestLargeWithdrawalsRequest)
	if err := _WithdrawalRequest.contract.UnpackLog(event, "LargeWithdrawalsRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalRequestLiquidStakingChangedIterator is returned from FilterLiquidStakingChanged and is used to iterate over the raw logs and unpacked data for LiquidStakingChanged events raised by the WithdrawalRequest contract.
type WithdrawalRequestLiquidStakingChangedIterator struct {
	Event *WithdrawalRequestLiquidStakingChanged // Event containing the contract specifics and raw log

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
func (it *WithdrawalRequestLiquidStakingChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalRequestLiquidStakingChanged)
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
		it.Event = new(WithdrawalRequestLiquidStakingChanged)
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
func (it *WithdrawalRequestLiquidStakingChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalRequestLiquidStakingChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalRequestLiquidStakingChanged represents a LiquidStakingChanged event raised by the WithdrawalRequest contract.
type WithdrawalRequestLiquidStakingChanged struct {
	OldLiquidStakingContract     common.Address
	LiquidStakingContractAddress common.Address
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterLiquidStakingChanged is a free log retrieval operation binding the contract event 0x516e369f70685d2905d214a9e8567098c02a0e00f5a55bd30baca6b51d446cef.
//
// Solidity: event LiquidStakingChanged(address _oldLiquidStakingContract, address _liquidStakingContractAddress)
func (_WithdrawalRequest *WithdrawalRequestFilterer) FilterLiquidStakingChanged(opts *bind.FilterOpts) (*WithdrawalRequestLiquidStakingChangedIterator, error) {

	logs, sub, err := _WithdrawalRequest.contract.FilterLogs(opts, "LiquidStakingChanged")
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequestLiquidStakingChangedIterator{contract: _WithdrawalRequest.contract, event: "LiquidStakingChanged", logs: logs, sub: sub}, nil
}

// WatchLiquidStakingChanged is a free log subscription operation binding the contract event 0x516e369f70685d2905d214a9e8567098c02a0e00f5a55bd30baca6b51d446cef.
//
// Solidity: event LiquidStakingChanged(address _oldLiquidStakingContract, address _liquidStakingContractAddress)
func (_WithdrawalRequest *WithdrawalRequestFilterer) WatchLiquidStakingChanged(opts *bind.WatchOpts, sink chan<- *WithdrawalRequestLiquidStakingChanged) (event.Subscription, error) {

	logs, sub, err := _WithdrawalRequest.contract.WatchLogs(opts, "LiquidStakingChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalRequestLiquidStakingChanged)
				if err := _WithdrawalRequest.contract.UnpackLog(event, "LiquidStakingChanged", log); err != nil {
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
func (_WithdrawalRequest *WithdrawalRequestFilterer) ParseLiquidStakingChanged(log types.Log) (*WithdrawalRequestLiquidStakingChanged, error) {
	event := new(WithdrawalRequestLiquidStakingChanged)
	if err := _WithdrawalRequest.contract.UnpackLog(event, "LiquidStakingChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalRequestNftUnstakeIterator is returned from FilterNftUnstake and is used to iterate over the raw logs and unpacked data for NftUnstake events raised by the WithdrawalRequest contract.
type WithdrawalRequestNftUnstakeIterator struct {
	Event *WithdrawalRequestNftUnstake // Event containing the contract specifics and raw log

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
func (it *WithdrawalRequestNftUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalRequestNftUnstake)
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
		it.Event = new(WithdrawalRequestNftUnstake)
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
func (it *WithdrawalRequestNftUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalRequestNftUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalRequestNftUnstake represents a NftUnstake event raised by the WithdrawalRequest contract.
type WithdrawalRequestNftUnstake struct {
	OperatorId *big.Int
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNftUnstake is a free log retrieval operation binding the contract event 0x9187b7a988fc2779224c0f80f4459cbb33f8c5c4be3656b4debaa4e58c19658e.
//
// Solidity: event NftUnstake(uint256 indexed _operatorId, uint256 tokenId)
func (_WithdrawalRequest *WithdrawalRequestFilterer) FilterNftUnstake(opts *bind.FilterOpts, _operatorId []*big.Int) (*WithdrawalRequestNftUnstakeIterator, error) {

	var _operatorIdRule []interface{}
	for _, _operatorIdItem := range _operatorId {
		_operatorIdRule = append(_operatorIdRule, _operatorIdItem)
	}

	logs, sub, err := _WithdrawalRequest.contract.FilterLogs(opts, "NftUnstake", _operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequestNftUnstakeIterator{contract: _WithdrawalRequest.contract, event: "NftUnstake", logs: logs, sub: sub}, nil
}

// WatchNftUnstake is a free log subscription operation binding the contract event 0x9187b7a988fc2779224c0f80f4459cbb33f8c5c4be3656b4debaa4e58c19658e.
//
// Solidity: event NftUnstake(uint256 indexed _operatorId, uint256 tokenId)
func (_WithdrawalRequest *WithdrawalRequestFilterer) WatchNftUnstake(opts *bind.WatchOpts, sink chan<- *WithdrawalRequestNftUnstake, _operatorId []*big.Int) (event.Subscription, error) {

	var _operatorIdRule []interface{}
	for _, _operatorIdItem := range _operatorId {
		_operatorIdRule = append(_operatorIdRule, _operatorIdItem)
	}

	logs, sub, err := _WithdrawalRequest.contract.WatchLogs(opts, "NftUnstake", _operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalRequestNftUnstake)
				if err := _WithdrawalRequest.contract.UnpackLog(event, "NftUnstake", log); err != nil {
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

// ParseNftUnstake is a log parse operation binding the contract event 0x9187b7a988fc2779224c0f80f4459cbb33f8c5c4be3656b4debaa4e58c19658e.
//
// Solidity: event NftUnstake(uint256 indexed _operatorId, uint256 tokenId)
func (_WithdrawalRequest *WithdrawalRequestFilterer) ParseNftUnstake(log types.Log) (*WithdrawalRequestNftUnstake, error) {
	event := new(WithdrawalRequestNftUnstake)
	if err := _WithdrawalRequest.contract.UnpackLog(event, "NftUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalRequestNodeOperatorRegistryContractSetIterator is returned from FilterNodeOperatorRegistryContractSet and is used to iterate over the raw logs and unpacked data for NodeOperatorRegistryContractSet events raised by the WithdrawalRequest contract.
type WithdrawalRequestNodeOperatorRegistryContractSetIterator struct {
	Event *WithdrawalRequestNodeOperatorRegistryContractSet // Event containing the contract specifics and raw log

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
func (it *WithdrawalRequestNodeOperatorRegistryContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalRequestNodeOperatorRegistryContractSet)
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
		it.Event = new(WithdrawalRequestNodeOperatorRegistryContractSet)
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
func (it *WithdrawalRequestNodeOperatorRegistryContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalRequestNodeOperatorRegistryContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalRequestNodeOperatorRegistryContractSet represents a NodeOperatorRegistryContractSet event raised by the WithdrawalRequest contract.
type WithdrawalRequestNodeOperatorRegistryContractSet struct {
	OldNodeOperatorRegistryContract common.Address
	NodeOperatorRegistryContract    common.Address
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorRegistryContractSet is a free log retrieval operation binding the contract event 0x2aa578b9d95064e7e90ca0af5e42ca5499f5e90bd32c4e401df52a686ac6993d.
//
// Solidity: event NodeOperatorRegistryContractSet(address _oldNodeOperatorRegistryContract, address _nodeOperatorRegistryContract)
func (_WithdrawalRequest *WithdrawalRequestFilterer) FilterNodeOperatorRegistryContractSet(opts *bind.FilterOpts) (*WithdrawalRequestNodeOperatorRegistryContractSetIterator, error) {

	logs, sub, err := _WithdrawalRequest.contract.FilterLogs(opts, "NodeOperatorRegistryContractSet")
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequestNodeOperatorRegistryContractSetIterator{contract: _WithdrawalRequest.contract, event: "NodeOperatorRegistryContractSet", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorRegistryContractSet is a free log subscription operation binding the contract event 0x2aa578b9d95064e7e90ca0af5e42ca5499f5e90bd32c4e401df52a686ac6993d.
//
// Solidity: event NodeOperatorRegistryContractSet(address _oldNodeOperatorRegistryContract, address _nodeOperatorRegistryContract)
func (_WithdrawalRequest *WithdrawalRequestFilterer) WatchNodeOperatorRegistryContractSet(opts *bind.WatchOpts, sink chan<- *WithdrawalRequestNodeOperatorRegistryContractSet) (event.Subscription, error) {

	logs, sub, err := _WithdrawalRequest.contract.WatchLogs(opts, "NodeOperatorRegistryContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalRequestNodeOperatorRegistryContractSet)
				if err := _WithdrawalRequest.contract.UnpackLog(event, "NodeOperatorRegistryContractSet", log); err != nil {
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
func (_WithdrawalRequest *WithdrawalRequestFilterer) ParseNodeOperatorRegistryContractSet(log types.Log) (*WithdrawalRequestNodeOperatorRegistryContractSet, error) {
	event := new(WithdrawalRequestNodeOperatorRegistryContractSet)
	if err := _WithdrawalRequest.contract.UnpackLog(event, "NodeOperatorRegistryContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalRequestOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WithdrawalRequest contract.
type WithdrawalRequestOwnershipTransferredIterator struct {
	Event *WithdrawalRequestOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WithdrawalRequestOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalRequestOwnershipTransferred)
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
		it.Event = new(WithdrawalRequestOwnershipTransferred)
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
func (it *WithdrawalRequestOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalRequestOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalRequestOwnershipTransferred represents a OwnershipTransferred event raised by the WithdrawalRequest contract.
type WithdrawalRequestOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WithdrawalRequest *WithdrawalRequestFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WithdrawalRequestOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WithdrawalRequest.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequestOwnershipTransferredIterator{contract: _WithdrawalRequest.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WithdrawalRequest *WithdrawalRequestFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WithdrawalRequestOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WithdrawalRequest.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalRequestOwnershipTransferred)
				if err := _WithdrawalRequest.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WithdrawalRequest *WithdrawalRequestFilterer) ParseOwnershipTransferred(log types.Log) (*WithdrawalRequestOwnershipTransferred, error) {
	event := new(WithdrawalRequestOwnershipTransferred)
	if err := _WithdrawalRequest.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalRequestPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the WithdrawalRequest contract.
type WithdrawalRequestPausedIterator struct {
	Event *WithdrawalRequestPaused // Event containing the contract specifics and raw log

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
func (it *WithdrawalRequestPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalRequestPaused)
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
		it.Event = new(WithdrawalRequestPaused)
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
func (it *WithdrawalRequestPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalRequestPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalRequestPaused represents a Paused event raised by the WithdrawalRequest contract.
type WithdrawalRequestPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WithdrawalRequest *WithdrawalRequestFilterer) FilterPaused(opts *bind.FilterOpts) (*WithdrawalRequestPausedIterator, error) {

	logs, sub, err := _WithdrawalRequest.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequestPausedIterator{contract: _WithdrawalRequest.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WithdrawalRequest *WithdrawalRequestFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *WithdrawalRequestPaused) (event.Subscription, error) {

	logs, sub, err := _WithdrawalRequest.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalRequestPaused)
				if err := _WithdrawalRequest.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_WithdrawalRequest *WithdrawalRequestFilterer) ParsePaused(log types.Log) (*WithdrawalRequestPaused, error) {
	event := new(WithdrawalRequestPaused)
	if err := _WithdrawalRequest.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalRequestUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the WithdrawalRequest contract.
type WithdrawalRequestUnpausedIterator struct {
	Event *WithdrawalRequestUnpaused // Event containing the contract specifics and raw log

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
func (it *WithdrawalRequestUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalRequestUnpaused)
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
		it.Event = new(WithdrawalRequestUnpaused)
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
func (it *WithdrawalRequestUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalRequestUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalRequestUnpaused represents a Unpaused event raised by the WithdrawalRequest contract.
type WithdrawalRequestUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WithdrawalRequest *WithdrawalRequestFilterer) FilterUnpaused(opts *bind.FilterOpts) (*WithdrawalRequestUnpausedIterator, error) {

	logs, sub, err := _WithdrawalRequest.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequestUnpausedIterator{contract: _WithdrawalRequest.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WithdrawalRequest *WithdrawalRequestFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WithdrawalRequestUnpaused) (event.Subscription, error) {

	logs, sub, err := _WithdrawalRequest.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalRequestUnpaused)
				if err := _WithdrawalRequest.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_WithdrawalRequest *WithdrawalRequestFilterer) ParseUnpaused(log types.Log) (*WithdrawalRequestUnpaused, error) {
	event := new(WithdrawalRequestUnpaused)
	if err := _WithdrawalRequest.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalRequestUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the WithdrawalRequest contract.
type WithdrawalRequestUpgradedIterator struct {
	Event *WithdrawalRequestUpgraded // Event containing the contract specifics and raw log

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
func (it *WithdrawalRequestUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalRequestUpgraded)
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
		it.Event = new(WithdrawalRequestUpgraded)
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
func (it *WithdrawalRequestUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalRequestUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalRequestUpgraded represents a Upgraded event raised by the WithdrawalRequest contract.
type WithdrawalRequestUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WithdrawalRequest *WithdrawalRequestFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*WithdrawalRequestUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WithdrawalRequest.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequestUpgradedIterator{contract: _WithdrawalRequest.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WithdrawalRequest *WithdrawalRequestFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *WithdrawalRequestUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WithdrawalRequest.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalRequestUpgraded)
				if err := _WithdrawalRequest.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_WithdrawalRequest *WithdrawalRequestFilterer) ParseUpgraded(log types.Log) (*WithdrawalRequestUpgraded, error) {
	event := new(WithdrawalRequestUpgraded)
	if err := _WithdrawalRequest.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalRequestVaultManagerContractSetIterator is returned from FilterVaultManagerContractSet and is used to iterate over the raw logs and unpacked data for VaultManagerContractSet events raised by the WithdrawalRequest contract.
type WithdrawalRequestVaultManagerContractSetIterator struct {
	Event *WithdrawalRequestVaultManagerContractSet // Event containing the contract specifics and raw log

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
func (it *WithdrawalRequestVaultManagerContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalRequestVaultManagerContractSet)
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
		it.Event = new(WithdrawalRequestVaultManagerContractSet)
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
func (it *WithdrawalRequestVaultManagerContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalRequestVaultManagerContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalRequestVaultManagerContractSet represents a VaultManagerContractSet event raised by the WithdrawalRequest contract.
type WithdrawalRequestVaultManagerContractSet struct {
	OldVaultManagerContractAddress common.Address
	VaultManagerContract           common.Address
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterVaultManagerContractSet is a free log retrieval operation binding the contract event 0x136260758ef216be6f30b5244361f089faf99890f23864c0a63e2d2def24963f.
//
// Solidity: event VaultManagerContractSet(address _oldVaultManagerContractAddress, address _vaultManagerContract)
func (_WithdrawalRequest *WithdrawalRequestFilterer) FilterVaultManagerContractSet(opts *bind.FilterOpts) (*WithdrawalRequestVaultManagerContractSetIterator, error) {

	logs, sub, err := _WithdrawalRequest.contract.FilterLogs(opts, "VaultManagerContractSet")
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequestVaultManagerContractSetIterator{contract: _WithdrawalRequest.contract, event: "VaultManagerContractSet", logs: logs, sub: sub}, nil
}

// WatchVaultManagerContractSet is a free log subscription operation binding the contract event 0x136260758ef216be6f30b5244361f089faf99890f23864c0a63e2d2def24963f.
//
// Solidity: event VaultManagerContractSet(address _oldVaultManagerContractAddress, address _vaultManagerContract)
func (_WithdrawalRequest *WithdrawalRequestFilterer) WatchVaultManagerContractSet(opts *bind.WatchOpts, sink chan<- *WithdrawalRequestVaultManagerContractSet) (event.Subscription, error) {

	logs, sub, err := _WithdrawalRequest.contract.WatchLogs(opts, "VaultManagerContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalRequestVaultManagerContractSet)
				if err := _WithdrawalRequest.contract.UnpackLog(event, "VaultManagerContractSet", log); err != nil {
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
func (_WithdrawalRequest *WithdrawalRequestFilterer) ParseVaultManagerContractSet(log types.Log) (*WithdrawalRequestVaultManagerContractSet, error) {
	event := new(WithdrawalRequestVaultManagerContractSet)
	if err := _WithdrawalRequest.contract.UnpackLog(event, "VaultManagerContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalRequestWithdrawalsReceiveIterator is returned from FilterWithdrawalsReceive and is used to iterate over the raw logs and unpacked data for WithdrawalsReceive events raised by the WithdrawalRequest contract.
type WithdrawalRequestWithdrawalsReceiveIterator struct {
	Event *WithdrawalRequestWithdrawalsReceive // Event containing the contract specifics and raw log

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
func (it *WithdrawalRequestWithdrawalsReceiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalRequestWithdrawalsReceive)
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
		it.Event = new(WithdrawalRequestWithdrawalsReceive)
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
func (it *WithdrawalRequestWithdrawalsReceiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalRequestWithdrawalsReceiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalRequestWithdrawalsReceive represents a WithdrawalsReceive event raised by the WithdrawalRequest contract.
type WithdrawalRequestWithdrawalsReceive struct {
	OperatorId *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalsReceive is a free log retrieval operation binding the contract event 0xc25e068eabde078f0570c7bfa229e7bc1c15807c08e18988ca69c679a1c0ad5c.
//
// Solidity: event WithdrawalsReceive(uint256 _operatorId, uint256 _amount)
func (_WithdrawalRequest *WithdrawalRequestFilterer) FilterWithdrawalsReceive(opts *bind.FilterOpts) (*WithdrawalRequestWithdrawalsReceiveIterator, error) {

	logs, sub, err := _WithdrawalRequest.contract.FilterLogs(opts, "WithdrawalsReceive")
	if err != nil {
		return nil, err
	}
	return &WithdrawalRequestWithdrawalsReceiveIterator{contract: _WithdrawalRequest.contract, event: "WithdrawalsReceive", logs: logs, sub: sub}, nil
}

// WatchWithdrawalsReceive is a free log subscription operation binding the contract event 0xc25e068eabde078f0570c7bfa229e7bc1c15807c08e18988ca69c679a1c0ad5c.
//
// Solidity: event WithdrawalsReceive(uint256 _operatorId, uint256 _amount)
func (_WithdrawalRequest *WithdrawalRequestFilterer) WatchWithdrawalsReceive(opts *bind.WatchOpts, sink chan<- *WithdrawalRequestWithdrawalsReceive) (event.Subscription, error) {

	logs, sub, err := _WithdrawalRequest.contract.WatchLogs(opts, "WithdrawalsReceive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalRequestWithdrawalsReceive)
				if err := _WithdrawalRequest.contract.UnpackLog(event, "WithdrawalsReceive", log); err != nil {
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

// ParseWithdrawalsReceive is a log parse operation binding the contract event 0xc25e068eabde078f0570c7bfa229e7bc1c15807c08e18988ca69c679a1c0ad5c.
//
// Solidity: event WithdrawalsReceive(uint256 _operatorId, uint256 _amount)
func (_WithdrawalRequest *WithdrawalRequestFilterer) ParseWithdrawalsReceive(log types.Log) (*WithdrawalRequestWithdrawalsReceive, error) {
	event := new(WithdrawalRequestWithdrawalsReceive)
	if err := _WithdrawalRequest.contract.UnpackLog(event, "WithdrawalsReceive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

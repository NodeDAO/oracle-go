// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package operator

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

// OperatorMetaData contains all meta data concerning the Operator contract.
var OperatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ControllerAddrUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientMargin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommission\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardRatio\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPermissionPhase\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorAlreadyTrusted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorExitFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorHasArrears\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorHasBlacklisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorHasExited\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorNotBlacklisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorNotTrusted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionlessPhaseStart\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"CommissionRateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldDao\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"DaoAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"DaoClaimRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldDaoVaultAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_daoVaultAddress\",\"type\":\"address\"}],\"name\":\"DaoVaultAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldDefaultOperatorCommission\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_defaultOperatorCommission\",\"type\":\"uint256\"}],\"name\":\"DefaultOperatorCommissionRateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldLargeStakingContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_largeStakingContractAddress\",\"type\":\"address\"}],\"name\":\"LargeStakingChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"LiquidStakingChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"NodeOperatorBlacklistRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"NodeOperatorBlacklistSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorControllerAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"NodeOperatorNameSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_ownerAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorOwnerAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_vaultContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_rewardAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_ratios\",\"type\":\"uint256[]\"}],\"name\":\"NodeOperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_rewardAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_ratios\",\"type\":\"uint256[]\"}],\"name\":\"NodeOperatorRewardAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_trusted\",\"type\":\"bool\"}],\"name\":\"NodeOperatorTrustedRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_trusted\",\"type\":\"bool\"}],\"name\":\"NodeOperatorTrustedSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"OperatorArrearsIncrease\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"OperatorArrearsReduce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"OperatorClaimRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nowVault\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"OperatorQuit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldOperatorSlashContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_operatorSlashContractAddress\",\"type\":\"address\"}],\"name\":\"OperatorSlashContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldVaultContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_vaultContractAddress\",\"type\":\"address\"}],\"name\":\"OperatorVaultContractReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_withdrawAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"OperatorWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"PermissionlessBlockNumberSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"PledgeDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"RegistrationFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_vaultFactoryContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_vaultFactoryContractAddress\",\"type\":\"address\"}],\"name\":\"VaultFactorContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIC_PLEDGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blacklistOperators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"controllerAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultOperatorCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_fullInfo\",\"type\":\"bool\"}],\"name\":\"getNodeOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_trusted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorRewardSetting\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"rewardAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ratios\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorVaultContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"vaultContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeOperatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_operatorIds\",\"type\":\"uint256[]\"}],\"name\":\"getOperatorCommissionRate\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getPledgeInfoOfOperator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_daoVaultAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultFactoryContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nVNFTContractAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultFactoryContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operatorSlashContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_resetVaultOperatorIds\",\"type\":\"uint256[]\"}],\"name\":\"initializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_largeStakingContractAddress\",\"type\":\"address\"}],\"name\":\"initializeV3\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"isBlacklistOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"isQuitOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"isTrustedOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"isTrustedOperatorOfControllerAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"largeStakingContract\",\"outputs\":[{\"internalType\":\"contractILargeStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidStakingContract\",\"outputs\":[{\"internalType\":\"contractILiquidStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorPledgeVaultBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorSlashAmountOwed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorSlashContract\",\"outputs\":[{\"internalType\":\"contractIOperatorSlash\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permissionlessBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"quitOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_rewardAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_ratios\",\"type\":\"uint256[]\"}],\"name\":\"registerOperator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registrationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"removeBlacklistOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"removeTrustedOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_operatorIds\",\"type\":\"uint256[]\"}],\"name\":\"resetOperatorVaultContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"setBlacklistOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setNodeOperatorOwnerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_daoVaultAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidStakingContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operatorSlashContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultFactoryContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_largeStakingContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_defaultOperatorCommission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_registrationFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_permissionlessBlockNumber\",\"type\":\"uint256\"}],\"name\":\"setNodeOperatorRegistrySetting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"setOperatorCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_rewardAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_ratios\",\"type\":\"uint256[]\"}],\"name\":\"setOperatorSetting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"setTrustedOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashType\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_slashIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_operatorIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"slashOfExitDelayed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"trustedControllerAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"usedControllerAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vNFTContract\",\"outputs\":[{\"internalType\":\"contractIVNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultFactoryContract\",\"outputs\":[{\"internalType\":\"contractIELVaultFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OperatorABI is the input ABI used to generate the binding from.
// Deprecated: Use OperatorMetaData.ABI instead.
var OperatorABI = OperatorMetaData.ABI

// Operator is an auto generated Go binding around an Ethereum contract.
type Operator struct {
	OperatorCaller     // Read-only binding to the contract
	OperatorTransactor // Write-only binding to the contract
	OperatorFilterer   // Log filterer for contract events
}

// OperatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OperatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperatorSession struct {
	Contract     *Operator         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OperatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperatorCallerSession struct {
	Contract *OperatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// OperatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperatorTransactorSession struct {
	Contract     *OperatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OperatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperatorRaw struct {
	Contract *Operator // Generic contract binding to access the raw methods on
}

// OperatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperatorCallerRaw struct {
	Contract *OperatorCaller // Generic read-only contract binding to access the raw methods on
}

// OperatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperatorTransactorRaw struct {
	Contract *OperatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperator creates a new instance of Operator, bound to a specific deployed contract.
func NewOperator(address common.Address, backend bind.ContractBackend) (*Operator, error) {
	contract, err := bindOperator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Operator{OperatorCaller: OperatorCaller{contract: contract}, OperatorTransactor: OperatorTransactor{contract: contract}, OperatorFilterer: OperatorFilterer{contract: contract}}, nil
}

// NewOperatorCaller creates a new read-only instance of Operator, bound to a specific deployed contract.
func NewOperatorCaller(address common.Address, caller bind.ContractCaller) (*OperatorCaller, error) {
	contract, err := bindOperator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorCaller{contract: contract}, nil
}

// NewOperatorTransactor creates a new write-only instance of Operator, bound to a specific deployed contract.
func NewOperatorTransactor(address common.Address, transactor bind.ContractTransactor) (*OperatorTransactor, error) {
	contract, err := bindOperator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorTransactor{contract: contract}, nil
}

// NewOperatorFilterer creates a new log filterer instance of Operator, bound to a specific deployed contract.
func NewOperatorFilterer(address common.Address, filterer bind.ContractFilterer) (*OperatorFilterer, error) {
	contract, err := bindOperator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OperatorFilterer{contract: contract}, nil
}

// bindOperator binds a generic wrapper to an already deployed contract.
func bindOperator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OperatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Operator *OperatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Operator.Contract.OperatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Operator *OperatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operator.Contract.OperatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Operator *OperatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Operator.Contract.OperatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Operator *OperatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Operator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Operator *OperatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Operator *OperatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Operator.Contract.contract.Transact(opts, method, params...)
}

// BASICPLEDGE is a free data retrieval call binding the contract method 0xcfb8012a.
//
// Solidity: function BASIC_PLEDGE() view returns(uint256)
func (_Operator *OperatorCaller) BASICPLEDGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "BASIC_PLEDGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASICPLEDGE is a free data retrieval call binding the contract method 0xcfb8012a.
//
// Solidity: function BASIC_PLEDGE() view returns(uint256)
func (_Operator *OperatorSession) BASICPLEDGE() (*big.Int, error) {
	return _Operator.Contract.BASICPLEDGE(&_Operator.CallOpts)
}

// BASICPLEDGE is a free data retrieval call binding the contract method 0xcfb8012a.
//
// Solidity: function BASIC_PLEDGE() view returns(uint256)
func (_Operator *OperatorCallerSession) BASICPLEDGE() (*big.Int, error) {
	return _Operator.Contract.BASICPLEDGE(&_Operator.CallOpts)
}

// BlacklistOperators is a free data retrieval call binding the contract method 0xc997f338.
//
// Solidity: function blacklistOperators(uint256 ) view returns(bool)
func (_Operator *OperatorCaller) BlacklistOperators(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "blacklistOperators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlacklistOperators is a free data retrieval call binding the contract method 0xc997f338.
//
// Solidity: function blacklistOperators(uint256 ) view returns(bool)
func (_Operator *OperatorSession) BlacklistOperators(arg0 *big.Int) (bool, error) {
	return _Operator.Contract.BlacklistOperators(&_Operator.CallOpts, arg0)
}

// BlacklistOperators is a free data retrieval call binding the contract method 0xc997f338.
//
// Solidity: function blacklistOperators(uint256 ) view returns(bool)
func (_Operator *OperatorCallerSession) BlacklistOperators(arg0 *big.Int) (bool, error) {
	return _Operator.Contract.BlacklistOperators(&_Operator.CallOpts, arg0)
}

// ControllerAddress is a free data retrieval call binding the contract method 0x10586937.
//
// Solidity: function controllerAddress(address ) view returns(uint256)
func (_Operator *OperatorCaller) ControllerAddress(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "controllerAddress", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ControllerAddress is a free data retrieval call binding the contract method 0x10586937.
//
// Solidity: function controllerAddress(address ) view returns(uint256)
func (_Operator *OperatorSession) ControllerAddress(arg0 common.Address) (*big.Int, error) {
	return _Operator.Contract.ControllerAddress(&_Operator.CallOpts, arg0)
}

// ControllerAddress is a free data retrieval call binding the contract method 0x10586937.
//
// Solidity: function controllerAddress(address ) view returns(uint256)
func (_Operator *OperatorCallerSession) ControllerAddress(arg0 common.Address) (*big.Int, error) {
	return _Operator.Contract.ControllerAddress(&_Operator.CallOpts, arg0)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Operator *OperatorCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Operator *OperatorSession) Dao() (common.Address, error) {
	return _Operator.Contract.Dao(&_Operator.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Operator *OperatorCallerSession) Dao() (common.Address, error) {
	return _Operator.Contract.Dao(&_Operator.CallOpts)
}

// DaoVaultAddress is a free data retrieval call binding the contract method 0x3d6a3844.
//
// Solidity: function daoVaultAddress() view returns(address)
func (_Operator *OperatorCaller) DaoVaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "daoVaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaoVaultAddress is a free data retrieval call binding the contract method 0x3d6a3844.
//
// Solidity: function daoVaultAddress() view returns(address)
func (_Operator *OperatorSession) DaoVaultAddress() (common.Address, error) {
	return _Operator.Contract.DaoVaultAddress(&_Operator.CallOpts)
}

// DaoVaultAddress is a free data retrieval call binding the contract method 0x3d6a3844.
//
// Solidity: function daoVaultAddress() view returns(address)
func (_Operator *OperatorCallerSession) DaoVaultAddress() (common.Address, error) {
	return _Operator.Contract.DaoVaultAddress(&_Operator.CallOpts)
}

// DefaultOperatorCommission is a free data retrieval call binding the contract method 0xf96b31dc.
//
// Solidity: function defaultOperatorCommission() view returns(uint256)
func (_Operator *OperatorCaller) DefaultOperatorCommission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "defaultOperatorCommission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultOperatorCommission is a free data retrieval call binding the contract method 0xf96b31dc.
//
// Solidity: function defaultOperatorCommission() view returns(uint256)
func (_Operator *OperatorSession) DefaultOperatorCommission() (*big.Int, error) {
	return _Operator.Contract.DefaultOperatorCommission(&_Operator.CallOpts)
}

// DefaultOperatorCommission is a free data retrieval call binding the contract method 0xf96b31dc.
//
// Solidity: function defaultOperatorCommission() view returns(uint256)
func (_Operator *OperatorCallerSession) DefaultOperatorCommission() (*big.Int, error) {
	return _Operator.Contract.DefaultOperatorCommission(&_Operator.CallOpts)
}

// GetNodeOperator is a free data retrieval call binding the contract method 0x9a56983c.
//
// Solidity: function getNodeOperator(uint256 _id, bool _fullInfo) view returns(bool _trusted, string _name, address _owner, address _controllerAddress, address _vaultContractAddress)
func (_Operator *OperatorCaller) GetNodeOperator(opts *bind.CallOpts, _id *big.Int, _fullInfo bool) (struct {
	Trusted              bool
	Name                 string
	Owner                common.Address
	ControllerAddress    common.Address
	VaultContractAddress common.Address
}, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "getNodeOperator", _id, _fullInfo)

	outstruct := new(struct {
		Trusted              bool
		Name                 string
		Owner                common.Address
		ControllerAddress    common.Address
		VaultContractAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Trusted = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ControllerAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.VaultContractAddress = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetNodeOperator is a free data retrieval call binding the contract method 0x9a56983c.
//
// Solidity: function getNodeOperator(uint256 _id, bool _fullInfo) view returns(bool _trusted, string _name, address _owner, address _controllerAddress, address _vaultContractAddress)
func (_Operator *OperatorSession) GetNodeOperator(_id *big.Int, _fullInfo bool) (struct {
	Trusted              bool
	Name                 string
	Owner                common.Address
	ControllerAddress    common.Address
	VaultContractAddress common.Address
}, error) {
	return _Operator.Contract.GetNodeOperator(&_Operator.CallOpts, _id, _fullInfo)
}

// GetNodeOperator is a free data retrieval call binding the contract method 0x9a56983c.
//
// Solidity: function getNodeOperator(uint256 _id, bool _fullInfo) view returns(bool _trusted, string _name, address _owner, address _controllerAddress, address _vaultContractAddress)
func (_Operator *OperatorCallerSession) GetNodeOperator(_id *big.Int, _fullInfo bool) (struct {
	Trusted              bool
	Name                 string
	Owner                common.Address
	ControllerAddress    common.Address
	VaultContractAddress common.Address
}, error) {
	return _Operator.Contract.GetNodeOperator(&_Operator.CallOpts, _id, _fullInfo)
}

// GetNodeOperatorRewardSetting is a free data retrieval call binding the contract method 0x0a61e6c4.
//
// Solidity: function getNodeOperatorRewardSetting(uint256 _operatorId) view returns(address[] rewardAddresses, uint256[] ratios)
func (_Operator *OperatorCaller) GetNodeOperatorRewardSetting(opts *bind.CallOpts, _operatorId *big.Int) (struct {
	RewardAddresses []common.Address
	Ratios          []*big.Int
}, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "getNodeOperatorRewardSetting", _operatorId)

	outstruct := new(struct {
		RewardAddresses []common.Address
		Ratios          []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardAddresses = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Ratios = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetNodeOperatorRewardSetting is a free data retrieval call binding the contract method 0x0a61e6c4.
//
// Solidity: function getNodeOperatorRewardSetting(uint256 _operatorId) view returns(address[] rewardAddresses, uint256[] ratios)
func (_Operator *OperatorSession) GetNodeOperatorRewardSetting(_operatorId *big.Int) (struct {
	RewardAddresses []common.Address
	Ratios          []*big.Int
}, error) {
	return _Operator.Contract.GetNodeOperatorRewardSetting(&_Operator.CallOpts, _operatorId)
}

// GetNodeOperatorRewardSetting is a free data retrieval call binding the contract method 0x0a61e6c4.
//
// Solidity: function getNodeOperatorRewardSetting(uint256 _operatorId) view returns(address[] rewardAddresses, uint256[] ratios)
func (_Operator *OperatorCallerSession) GetNodeOperatorRewardSetting(_operatorId *big.Int) (struct {
	RewardAddresses []common.Address
	Ratios          []*big.Int
}, error) {
	return _Operator.Contract.GetNodeOperatorRewardSetting(&_Operator.CallOpts, _operatorId)
}

// GetNodeOperatorVaultContract is a free data retrieval call binding the contract method 0x8b9135fc.
//
// Solidity: function getNodeOperatorVaultContract(uint256 _id) view returns(address vaultContractAddress)
func (_Operator *OperatorCaller) GetNodeOperatorVaultContract(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "getNodeOperatorVaultContract", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeOperatorVaultContract is a free data retrieval call binding the contract method 0x8b9135fc.
//
// Solidity: function getNodeOperatorVaultContract(uint256 _id) view returns(address vaultContractAddress)
func (_Operator *OperatorSession) GetNodeOperatorVaultContract(_id *big.Int) (common.Address, error) {
	return _Operator.Contract.GetNodeOperatorVaultContract(&_Operator.CallOpts, _id)
}

// GetNodeOperatorVaultContract is a free data retrieval call binding the contract method 0x8b9135fc.
//
// Solidity: function getNodeOperatorVaultContract(uint256 _id) view returns(address vaultContractAddress)
func (_Operator *OperatorCallerSession) GetNodeOperatorVaultContract(_id *big.Int) (common.Address, error) {
	return _Operator.Contract.GetNodeOperatorVaultContract(&_Operator.CallOpts, _id)
}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_Operator *OperatorCaller) GetNodeOperatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "getNodeOperatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_Operator *OperatorSession) GetNodeOperatorsCount() (*big.Int, error) {
	return _Operator.Contract.GetNodeOperatorsCount(&_Operator.CallOpts)
}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_Operator *OperatorCallerSession) GetNodeOperatorsCount() (*big.Int, error) {
	return _Operator.Contract.GetNodeOperatorsCount(&_Operator.CallOpts)
}

// GetOperatorCommissionRate is a free data retrieval call binding the contract method 0x50aee62f.
//
// Solidity: function getOperatorCommissionRate(uint256[] _operatorIds) view returns(uint256[])
func (_Operator *OperatorCaller) GetOperatorCommissionRate(opts *bind.CallOpts, _operatorIds []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "getOperatorCommissionRate", _operatorIds)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOperatorCommissionRate is a free data retrieval call binding the contract method 0x50aee62f.
//
// Solidity: function getOperatorCommissionRate(uint256[] _operatorIds) view returns(uint256[])
func (_Operator *OperatorSession) GetOperatorCommissionRate(_operatorIds []*big.Int) ([]*big.Int, error) {
	return _Operator.Contract.GetOperatorCommissionRate(&_Operator.CallOpts, _operatorIds)
}

// GetOperatorCommissionRate is a free data retrieval call binding the contract method 0x50aee62f.
//
// Solidity: function getOperatorCommissionRate(uint256[] _operatorIds) view returns(uint256[])
func (_Operator *OperatorCallerSession) GetOperatorCommissionRate(_operatorIds []*big.Int) ([]*big.Int, error) {
	return _Operator.Contract.GetOperatorCommissionRate(&_Operator.CallOpts, _operatorIds)
}

// GetPledgeInfoOfOperator is a free data retrieval call binding the contract method 0x2bcc1852.
//
// Solidity: function getPledgeInfoOfOperator(uint256 _operatorId) view returns(uint256, uint256)
func (_Operator *OperatorCaller) GetPledgeInfoOfOperator(opts *bind.CallOpts, _operatorId *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "getPledgeInfoOfOperator", _operatorId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPledgeInfoOfOperator is a free data retrieval call binding the contract method 0x2bcc1852.
//
// Solidity: function getPledgeInfoOfOperator(uint256 _operatorId) view returns(uint256, uint256)
func (_Operator *OperatorSession) GetPledgeInfoOfOperator(_operatorId *big.Int) (*big.Int, *big.Int, error) {
	return _Operator.Contract.GetPledgeInfoOfOperator(&_Operator.CallOpts, _operatorId)
}

// GetPledgeInfoOfOperator is a free data retrieval call binding the contract method 0x2bcc1852.
//
// Solidity: function getPledgeInfoOfOperator(uint256 _operatorId) view returns(uint256, uint256)
func (_Operator *OperatorCallerSession) GetPledgeInfoOfOperator(_operatorId *big.Int) (*big.Int, *big.Int, error) {
	return _Operator.Contract.GetPledgeInfoOfOperator(&_Operator.CallOpts, _operatorId)
}

// IsBlacklistOperator is a free data retrieval call binding the contract method 0x3b781aad.
//
// Solidity: function isBlacklistOperator(uint256 _operatorId) view returns(bool)
func (_Operator *OperatorCaller) IsBlacklistOperator(opts *bind.CallOpts, _operatorId *big.Int) (bool, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "isBlacklistOperator", _operatorId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklistOperator is a free data retrieval call binding the contract method 0x3b781aad.
//
// Solidity: function isBlacklistOperator(uint256 _operatorId) view returns(bool)
func (_Operator *OperatorSession) IsBlacklistOperator(_operatorId *big.Int) (bool, error) {
	return _Operator.Contract.IsBlacklistOperator(&_Operator.CallOpts, _operatorId)
}

// IsBlacklistOperator is a free data retrieval call binding the contract method 0x3b781aad.
//
// Solidity: function isBlacklistOperator(uint256 _operatorId) view returns(bool)
func (_Operator *OperatorCallerSession) IsBlacklistOperator(_operatorId *big.Int) (bool, error) {
	return _Operator.Contract.IsBlacklistOperator(&_Operator.CallOpts, _operatorId)
}

// IsQuitOperator is a free data retrieval call binding the contract method 0x9611d4e2.
//
// Solidity: function isQuitOperator(uint256 _operatorId) view returns(bool)
func (_Operator *OperatorCaller) IsQuitOperator(opts *bind.CallOpts, _operatorId *big.Int) (bool, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "isQuitOperator", _operatorId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsQuitOperator is a free data retrieval call binding the contract method 0x9611d4e2.
//
// Solidity: function isQuitOperator(uint256 _operatorId) view returns(bool)
func (_Operator *OperatorSession) IsQuitOperator(_operatorId *big.Int) (bool, error) {
	return _Operator.Contract.IsQuitOperator(&_Operator.CallOpts, _operatorId)
}

// IsQuitOperator is a free data retrieval call binding the contract method 0x9611d4e2.
//
// Solidity: function isQuitOperator(uint256 _operatorId) view returns(bool)
func (_Operator *OperatorCallerSession) IsQuitOperator(_operatorId *big.Int) (bool, error) {
	return _Operator.Contract.IsQuitOperator(&_Operator.CallOpts, _operatorId)
}

// IsTrustedOperator is a free data retrieval call binding the contract method 0xf2eeb337.
//
// Solidity: function isTrustedOperator(uint256 _operatorId) view returns(bool)
func (_Operator *OperatorCaller) IsTrustedOperator(opts *bind.CallOpts, _operatorId *big.Int) (bool, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "isTrustedOperator", _operatorId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedOperator is a free data retrieval call binding the contract method 0xf2eeb337.
//
// Solidity: function isTrustedOperator(uint256 _operatorId) view returns(bool)
func (_Operator *OperatorSession) IsTrustedOperator(_operatorId *big.Int) (bool, error) {
	return _Operator.Contract.IsTrustedOperator(&_Operator.CallOpts, _operatorId)
}

// IsTrustedOperator is a free data retrieval call binding the contract method 0xf2eeb337.
//
// Solidity: function isTrustedOperator(uint256 _operatorId) view returns(bool)
func (_Operator *OperatorCallerSession) IsTrustedOperator(_operatorId *big.Int) (bool, error) {
	return _Operator.Contract.IsTrustedOperator(&_Operator.CallOpts, _operatorId)
}

// IsTrustedOperatorOfControllerAddress is a free data retrieval call binding the contract method 0x8f5c9ff0.
//
// Solidity: function isTrustedOperatorOfControllerAddress(address _controllerAddress) view returns(uint256)
func (_Operator *OperatorCaller) IsTrustedOperatorOfControllerAddress(opts *bind.CallOpts, _controllerAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "isTrustedOperatorOfControllerAddress", _controllerAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsTrustedOperatorOfControllerAddress is a free data retrieval call binding the contract method 0x8f5c9ff0.
//
// Solidity: function isTrustedOperatorOfControllerAddress(address _controllerAddress) view returns(uint256)
func (_Operator *OperatorSession) IsTrustedOperatorOfControllerAddress(_controllerAddress common.Address) (*big.Int, error) {
	return _Operator.Contract.IsTrustedOperatorOfControllerAddress(&_Operator.CallOpts, _controllerAddress)
}

// IsTrustedOperatorOfControllerAddress is a free data retrieval call binding the contract method 0x8f5c9ff0.
//
// Solidity: function isTrustedOperatorOfControllerAddress(address _controllerAddress) view returns(uint256)
func (_Operator *OperatorCallerSession) IsTrustedOperatorOfControllerAddress(_controllerAddress common.Address) (*big.Int, error) {
	return _Operator.Contract.IsTrustedOperatorOfControllerAddress(&_Operator.CallOpts, _controllerAddress)
}

// LargeStakingContract is a free data retrieval call binding the contract method 0xf08ae442.
//
// Solidity: function largeStakingContract() view returns(address)
func (_Operator *OperatorCaller) LargeStakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "largeStakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LargeStakingContract is a free data retrieval call binding the contract method 0xf08ae442.
//
// Solidity: function largeStakingContract() view returns(address)
func (_Operator *OperatorSession) LargeStakingContract() (common.Address, error) {
	return _Operator.Contract.LargeStakingContract(&_Operator.CallOpts)
}

// LargeStakingContract is a free data retrieval call binding the contract method 0xf08ae442.
//
// Solidity: function largeStakingContract() view returns(address)
func (_Operator *OperatorCallerSession) LargeStakingContract() (common.Address, error) {
	return _Operator.Contract.LargeStakingContract(&_Operator.CallOpts)
}

// LiquidStakingContract is a free data retrieval call binding the contract method 0xbdcaa355.
//
// Solidity: function liquidStakingContract() view returns(address)
func (_Operator *OperatorCaller) LiquidStakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "liquidStakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidStakingContract is a free data retrieval call binding the contract method 0xbdcaa355.
//
// Solidity: function liquidStakingContract() view returns(address)
func (_Operator *OperatorSession) LiquidStakingContract() (common.Address, error) {
	return _Operator.Contract.LiquidStakingContract(&_Operator.CallOpts)
}

// LiquidStakingContract is a free data retrieval call binding the contract method 0xbdcaa355.
//
// Solidity: function liquidStakingContract() view returns(address)
func (_Operator *OperatorCallerSession) LiquidStakingContract() (common.Address, error) {
	return _Operator.Contract.LiquidStakingContract(&_Operator.CallOpts)
}

// OperatorPledgeVaultBalances is a free data retrieval call binding the contract method 0x92064988.
//
// Solidity: function operatorPledgeVaultBalances(uint256 ) view returns(uint256)
func (_Operator *OperatorCaller) OperatorPledgeVaultBalances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "operatorPledgeVaultBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorPledgeVaultBalances is a free data retrieval call binding the contract method 0x92064988.
//
// Solidity: function operatorPledgeVaultBalances(uint256 ) view returns(uint256)
func (_Operator *OperatorSession) OperatorPledgeVaultBalances(arg0 *big.Int) (*big.Int, error) {
	return _Operator.Contract.OperatorPledgeVaultBalances(&_Operator.CallOpts, arg0)
}

// OperatorPledgeVaultBalances is a free data retrieval call binding the contract method 0x92064988.
//
// Solidity: function operatorPledgeVaultBalances(uint256 ) view returns(uint256)
func (_Operator *OperatorCallerSession) OperatorPledgeVaultBalances(arg0 *big.Int) (*big.Int, error) {
	return _Operator.Contract.OperatorPledgeVaultBalances(&_Operator.CallOpts, arg0)
}

// OperatorSlashAmountOwed is a free data retrieval call binding the contract method 0x6f88e169.
//
// Solidity: function operatorSlashAmountOwed(uint256 ) view returns(uint256)
func (_Operator *OperatorCaller) OperatorSlashAmountOwed(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "operatorSlashAmountOwed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorSlashAmountOwed is a free data retrieval call binding the contract method 0x6f88e169.
//
// Solidity: function operatorSlashAmountOwed(uint256 ) view returns(uint256)
func (_Operator *OperatorSession) OperatorSlashAmountOwed(arg0 *big.Int) (*big.Int, error) {
	return _Operator.Contract.OperatorSlashAmountOwed(&_Operator.CallOpts, arg0)
}

// OperatorSlashAmountOwed is a free data retrieval call binding the contract method 0x6f88e169.
//
// Solidity: function operatorSlashAmountOwed(uint256 ) view returns(uint256)
func (_Operator *OperatorCallerSession) OperatorSlashAmountOwed(arg0 *big.Int) (*big.Int, error) {
	return _Operator.Contract.OperatorSlashAmountOwed(&_Operator.CallOpts, arg0)
}

// OperatorSlashContract is a free data retrieval call binding the contract method 0x0c2559fd.
//
// Solidity: function operatorSlashContract() view returns(address)
func (_Operator *OperatorCaller) OperatorSlashContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "operatorSlashContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorSlashContract is a free data retrieval call binding the contract method 0x0c2559fd.
//
// Solidity: function operatorSlashContract() view returns(address)
func (_Operator *OperatorSession) OperatorSlashContract() (common.Address, error) {
	return _Operator.Contract.OperatorSlashContract(&_Operator.CallOpts)
}

// OperatorSlashContract is a free data retrieval call binding the contract method 0x0c2559fd.
//
// Solidity: function operatorSlashContract() view returns(address)
func (_Operator *OperatorCallerSession) OperatorSlashContract() (common.Address, error) {
	return _Operator.Contract.OperatorSlashContract(&_Operator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Operator *OperatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Operator *OperatorSession) Owner() (common.Address, error) {
	return _Operator.Contract.Owner(&_Operator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Operator *OperatorCallerSession) Owner() (common.Address, error) {
	return _Operator.Contract.Owner(&_Operator.CallOpts)
}

// PermissionlessBlockNumber is a free data retrieval call binding the contract method 0xced1f826.
//
// Solidity: function permissionlessBlockNumber() view returns(uint256)
func (_Operator *OperatorCaller) PermissionlessBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "permissionlessBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PermissionlessBlockNumber is a free data retrieval call binding the contract method 0xced1f826.
//
// Solidity: function permissionlessBlockNumber() view returns(uint256)
func (_Operator *OperatorSession) PermissionlessBlockNumber() (*big.Int, error) {
	return _Operator.Contract.PermissionlessBlockNumber(&_Operator.CallOpts)
}

// PermissionlessBlockNumber is a free data retrieval call binding the contract method 0xced1f826.
//
// Solidity: function permissionlessBlockNumber() view returns(uint256)
func (_Operator *OperatorCallerSession) PermissionlessBlockNumber() (*big.Int, error) {
	return _Operator.Contract.PermissionlessBlockNumber(&_Operator.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Operator *OperatorCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Operator *OperatorSession) ProxiableUUID() ([32]byte, error) {
	return _Operator.Contract.ProxiableUUID(&_Operator.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Operator *OperatorCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Operator.Contract.ProxiableUUID(&_Operator.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() view returns(uint256)
func (_Operator *OperatorCaller) RegistrationFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "registrationFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() view returns(uint256)
func (_Operator *OperatorSession) RegistrationFee() (*big.Int, error) {
	return _Operator.Contract.RegistrationFee(&_Operator.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() view returns(uint256)
func (_Operator *OperatorCallerSession) RegistrationFee() (*big.Int, error) {
	return _Operator.Contract.RegistrationFee(&_Operator.CallOpts)
}

// TrustedControllerAddress is a free data retrieval call binding the contract method 0xd83e2ed3.
//
// Solidity: function trustedControllerAddress(address ) view returns(uint256)
func (_Operator *OperatorCaller) TrustedControllerAddress(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "trustedControllerAddress", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TrustedControllerAddress is a free data retrieval call binding the contract method 0xd83e2ed3.
//
// Solidity: function trustedControllerAddress(address ) view returns(uint256)
func (_Operator *OperatorSession) TrustedControllerAddress(arg0 common.Address) (*big.Int, error) {
	return _Operator.Contract.TrustedControllerAddress(&_Operator.CallOpts, arg0)
}

// TrustedControllerAddress is a free data retrieval call binding the contract method 0xd83e2ed3.
//
// Solidity: function trustedControllerAddress(address ) view returns(uint256)
func (_Operator *OperatorCallerSession) TrustedControllerAddress(arg0 common.Address) (*big.Int, error) {
	return _Operator.Contract.TrustedControllerAddress(&_Operator.CallOpts, arg0)
}

// UsedControllerAddress is a free data retrieval call binding the contract method 0xa0cb5690.
//
// Solidity: function usedControllerAddress(address ) view returns(bool)
func (_Operator *OperatorCaller) UsedControllerAddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "usedControllerAddress", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedControllerAddress is a free data retrieval call binding the contract method 0xa0cb5690.
//
// Solidity: function usedControllerAddress(address ) view returns(bool)
func (_Operator *OperatorSession) UsedControllerAddress(arg0 common.Address) (bool, error) {
	return _Operator.Contract.UsedControllerAddress(&_Operator.CallOpts, arg0)
}

// UsedControllerAddress is a free data retrieval call binding the contract method 0xa0cb5690.
//
// Solidity: function usedControllerAddress(address ) view returns(bool)
func (_Operator *OperatorCallerSession) UsedControllerAddress(arg0 common.Address) (bool, error) {
	return _Operator.Contract.UsedControllerAddress(&_Operator.CallOpts, arg0)
}

// VNFTContract is a free data retrieval call binding the contract method 0xfc03411f.
//
// Solidity: function vNFTContract() view returns(address)
func (_Operator *OperatorCaller) VNFTContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "vNFTContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VNFTContract is a free data retrieval call binding the contract method 0xfc03411f.
//
// Solidity: function vNFTContract() view returns(address)
func (_Operator *OperatorSession) VNFTContract() (common.Address, error) {
	return _Operator.Contract.VNFTContract(&_Operator.CallOpts)
}

// VNFTContract is a free data retrieval call binding the contract method 0xfc03411f.
//
// Solidity: function vNFTContract() view returns(address)
func (_Operator *OperatorCallerSession) VNFTContract() (common.Address, error) {
	return _Operator.Contract.VNFTContract(&_Operator.CallOpts)
}

// VaultFactoryContract is a free data retrieval call binding the contract method 0x6e2d47f6.
//
// Solidity: function vaultFactoryContract() view returns(address)
func (_Operator *OperatorCaller) VaultFactoryContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "vaultFactoryContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultFactoryContract is a free data retrieval call binding the contract method 0x6e2d47f6.
//
// Solidity: function vaultFactoryContract() view returns(address)
func (_Operator *OperatorSession) VaultFactoryContract() (common.Address, error) {
	return _Operator.Contract.VaultFactoryContract(&_Operator.CallOpts)
}

// VaultFactoryContract is a free data retrieval call binding the contract method 0x6e2d47f6.
//
// Solidity: function vaultFactoryContract() view returns(address)
func (_Operator *OperatorCallerSession) VaultFactoryContract() (common.Address, error) {
	return _Operator.Contract.VaultFactoryContract(&_Operator.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _operatorId) payable returns()
func (_Operator *OperatorTransactor) Deposit(opts *bind.TransactOpts, _operatorId *big.Int) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "deposit", _operatorId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _operatorId) payable returns()
func (_Operator *OperatorSession) Deposit(_operatorId *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.Deposit(&_Operator.TransactOpts, _operatorId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _operatorId) payable returns()
func (_Operator *OperatorTransactorSession) Deposit(_operatorId *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.Deposit(&_Operator.TransactOpts, _operatorId)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _dao, address _daoVaultAddress, address _vaultFactoryContractAddress, address _nVNFTContractAddress) returns()
func (_Operator *OperatorTransactor) Initialize(opts *bind.TransactOpts, _dao common.Address, _daoVaultAddress common.Address, _vaultFactoryContractAddress common.Address, _nVNFTContractAddress common.Address) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "initialize", _dao, _daoVaultAddress, _vaultFactoryContractAddress, _nVNFTContractAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _dao, address _daoVaultAddress, address _vaultFactoryContractAddress, address _nVNFTContractAddress) returns()
func (_Operator *OperatorSession) Initialize(_dao common.Address, _daoVaultAddress common.Address, _vaultFactoryContractAddress common.Address, _nVNFTContractAddress common.Address) (*types.Transaction, error) {
	return _Operator.Contract.Initialize(&_Operator.TransactOpts, _dao, _daoVaultAddress, _vaultFactoryContractAddress, _nVNFTContractAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _dao, address _daoVaultAddress, address _vaultFactoryContractAddress, address _nVNFTContractAddress) returns()
func (_Operator *OperatorTransactorSession) Initialize(_dao common.Address, _daoVaultAddress common.Address, _vaultFactoryContractAddress common.Address, _nVNFTContractAddress common.Address) (*types.Transaction, error) {
	return _Operator.Contract.Initialize(&_Operator.TransactOpts, _dao, _daoVaultAddress, _vaultFactoryContractAddress, _nVNFTContractAddress)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0xfb966e9f.
//
// Solidity: function initializeV2(address _vaultFactoryContractAddress, address _operatorSlashContractAddress, uint256[] _resetVaultOperatorIds) returns()
func (_Operator *OperatorTransactor) InitializeV2(opts *bind.TransactOpts, _vaultFactoryContractAddress common.Address, _operatorSlashContractAddress common.Address, _resetVaultOperatorIds []*big.Int) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "initializeV2", _vaultFactoryContractAddress, _operatorSlashContractAddress, _resetVaultOperatorIds)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0xfb966e9f.
//
// Solidity: function initializeV2(address _vaultFactoryContractAddress, address _operatorSlashContractAddress, uint256[] _resetVaultOperatorIds) returns()
func (_Operator *OperatorSession) InitializeV2(_vaultFactoryContractAddress common.Address, _operatorSlashContractAddress common.Address, _resetVaultOperatorIds []*big.Int) (*types.Transaction, error) {
	return _Operator.Contract.InitializeV2(&_Operator.TransactOpts, _vaultFactoryContractAddress, _operatorSlashContractAddress, _resetVaultOperatorIds)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0xfb966e9f.
//
// Solidity: function initializeV2(address _vaultFactoryContractAddress, address _operatorSlashContractAddress, uint256[] _resetVaultOperatorIds) returns()
func (_Operator *OperatorTransactorSession) InitializeV2(_vaultFactoryContractAddress common.Address, _operatorSlashContractAddress common.Address, _resetVaultOperatorIds []*big.Int) (*types.Transaction, error) {
	return _Operator.Contract.InitializeV2(&_Operator.TransactOpts, _vaultFactoryContractAddress, _operatorSlashContractAddress, _resetVaultOperatorIds)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0x3101cfcb.
//
// Solidity: function initializeV3(address _largeStakingContractAddress) returns()
func (_Operator *OperatorTransactor) InitializeV3(opts *bind.TransactOpts, _largeStakingContractAddress common.Address) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "initializeV3", _largeStakingContractAddress)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0x3101cfcb.
//
// Solidity: function initializeV3(address _largeStakingContractAddress) returns()
func (_Operator *OperatorSession) InitializeV3(_largeStakingContractAddress common.Address) (*types.Transaction, error) {
	return _Operator.Contract.InitializeV3(&_Operator.TransactOpts, _largeStakingContractAddress)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0x3101cfcb.
//
// Solidity: function initializeV3(address _largeStakingContractAddress) returns()
func (_Operator *OperatorTransactorSession) InitializeV3(_largeStakingContractAddress common.Address) (*types.Transaction, error) {
	return _Operator.Contract.InitializeV3(&_Operator.TransactOpts, _largeStakingContractAddress)
}

// QuitOperator is a paid mutator transaction binding the contract method 0x14ac7a70.
//
// Solidity: function quitOperator(uint256 _operatorId, address _to) returns()
func (_Operator *OperatorTransactor) QuitOperator(opts *bind.TransactOpts, _operatorId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "quitOperator", _operatorId, _to)
}

// QuitOperator is a paid mutator transaction binding the contract method 0x14ac7a70.
//
// Solidity: function quitOperator(uint256 _operatorId, address _to) returns()
func (_Operator *OperatorSession) QuitOperator(_operatorId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Operator.Contract.QuitOperator(&_Operator.TransactOpts, _operatorId, _to)
}

// QuitOperator is a paid mutator transaction binding the contract method 0x14ac7a70.
//
// Solidity: function quitOperator(uint256 _operatorId, address _to) returns()
func (_Operator *OperatorTransactorSession) QuitOperator(_operatorId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Operator.Contract.QuitOperator(&_Operator.TransactOpts, _operatorId, _to)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x8633dcb0.
//
// Solidity: function registerOperator(string _name, address _controllerAddress, address _owner, address[] _rewardAddresses, uint256[] _ratios) payable returns(uint256 id)
func (_Operator *OperatorTransactor) RegisterOperator(opts *bind.TransactOpts, _name string, _controllerAddress common.Address, _owner common.Address, _rewardAddresses []common.Address, _ratios []*big.Int) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "registerOperator", _name, _controllerAddress, _owner, _rewardAddresses, _ratios)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x8633dcb0.
//
// Solidity: function registerOperator(string _name, address _controllerAddress, address _owner, address[] _rewardAddresses, uint256[] _ratios) payable returns(uint256 id)
func (_Operator *OperatorSession) RegisterOperator(_name string, _controllerAddress common.Address, _owner common.Address, _rewardAddresses []common.Address, _ratios []*big.Int) (*types.Transaction, error) {
	return _Operator.Contract.RegisterOperator(&_Operator.TransactOpts, _name, _controllerAddress, _owner, _rewardAddresses, _ratios)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x8633dcb0.
//
// Solidity: function registerOperator(string _name, address _controllerAddress, address _owner, address[] _rewardAddresses, uint256[] _ratios) payable returns(uint256 id)
func (_Operator *OperatorTransactorSession) RegisterOperator(_name string, _controllerAddress common.Address, _owner common.Address, _rewardAddresses []common.Address, _ratios []*big.Int) (*types.Transaction, error) {
	return _Operator.Contract.RegisterOperator(&_Operator.TransactOpts, _name, _controllerAddress, _owner, _rewardAddresses, _ratios)
}

// RemoveBlacklistOperator is a paid mutator transaction binding the contract method 0x6854c364.
//
// Solidity: function removeBlacklistOperator(uint256 _id) returns()
func (_Operator *OperatorTransactor) RemoveBlacklistOperator(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "removeBlacklistOperator", _id)
}

// RemoveBlacklistOperator is a paid mutator transaction binding the contract method 0x6854c364.
//
// Solidity: function removeBlacklistOperator(uint256 _id) returns()
func (_Operator *OperatorSession) RemoveBlacklistOperator(_id *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.RemoveBlacklistOperator(&_Operator.TransactOpts, _id)
}

// RemoveBlacklistOperator is a paid mutator transaction binding the contract method 0x6854c364.
//
// Solidity: function removeBlacklistOperator(uint256 _id) returns()
func (_Operator *OperatorTransactorSession) RemoveBlacklistOperator(_id *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.RemoveBlacklistOperator(&_Operator.TransactOpts, _id)
}

// RemoveTrustedOperator is a paid mutator transaction binding the contract method 0x964450a3.
//
// Solidity: function removeTrustedOperator(uint256 _id) returns()
func (_Operator *OperatorTransactor) RemoveTrustedOperator(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "removeTrustedOperator", _id)
}

// RemoveTrustedOperator is a paid mutator transaction binding the contract method 0x964450a3.
//
// Solidity: function removeTrustedOperator(uint256 _id) returns()
func (_Operator *OperatorSession) RemoveTrustedOperator(_id *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.RemoveTrustedOperator(&_Operator.TransactOpts, _id)
}

// RemoveTrustedOperator is a paid mutator transaction binding the contract method 0x964450a3.
//
// Solidity: function removeTrustedOperator(uint256 _id) returns()
func (_Operator *OperatorTransactorSession) RemoveTrustedOperator(_id *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.RemoveTrustedOperator(&_Operator.TransactOpts, _id)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Operator *OperatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Operator *OperatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Operator.Contract.RenounceOwnership(&_Operator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Operator *OperatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Operator.Contract.RenounceOwnership(&_Operator.TransactOpts)
}

// ResetOperatorVaultContract is a paid mutator transaction binding the contract method 0xa1f3f228.
//
// Solidity: function resetOperatorVaultContract(uint256[] _operatorIds) returns()
func (_Operator *OperatorTransactor) ResetOperatorVaultContract(opts *bind.TransactOpts, _operatorIds []*big.Int) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "resetOperatorVaultContract", _operatorIds)
}

// ResetOperatorVaultContract is a paid mutator transaction binding the contract method 0xa1f3f228.
//
// Solidity: function resetOperatorVaultContract(uint256[] _operatorIds) returns()
func (_Operator *OperatorSession) ResetOperatorVaultContract(_operatorIds []*big.Int) (*types.Transaction, error) {
	return _Operator.Contract.ResetOperatorVaultContract(&_Operator.TransactOpts, _operatorIds)
}

// ResetOperatorVaultContract is a paid mutator transaction binding the contract method 0xa1f3f228.
//
// Solidity: function resetOperatorVaultContract(uint256[] _operatorIds) returns()
func (_Operator *OperatorTransactorSession) ResetOperatorVaultContract(_operatorIds []*big.Int) (*types.Transaction, error) {
	return _Operator.Contract.ResetOperatorVaultContract(&_Operator.TransactOpts, _operatorIds)
}

// SetBlacklistOperator is a paid mutator transaction binding the contract method 0x467ac78e.
//
// Solidity: function setBlacklistOperator(uint256 _id) returns()
func (_Operator *OperatorTransactor) SetBlacklistOperator(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "setBlacklistOperator", _id)
}

// SetBlacklistOperator is a paid mutator transaction binding the contract method 0x467ac78e.
//
// Solidity: function setBlacklistOperator(uint256 _id) returns()
func (_Operator *OperatorSession) SetBlacklistOperator(_id *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.SetBlacklistOperator(&_Operator.TransactOpts, _id)
}

// SetBlacklistOperator is a paid mutator transaction binding the contract method 0x467ac78e.
//
// Solidity: function setBlacklistOperator(uint256 _id) returns()
func (_Operator *OperatorTransactorSession) SetBlacklistOperator(_id *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.SetBlacklistOperator(&_Operator.TransactOpts, _id)
}

// SetNodeOperatorOwnerAddress is a paid mutator transaction binding the contract method 0x49790fff.
//
// Solidity: function setNodeOperatorOwnerAddress(uint256 _id, address _owner) returns()
func (_Operator *OperatorTransactor) SetNodeOperatorOwnerAddress(opts *bind.TransactOpts, _id *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "setNodeOperatorOwnerAddress", _id, _owner)
}

// SetNodeOperatorOwnerAddress is a paid mutator transaction binding the contract method 0x49790fff.
//
// Solidity: function setNodeOperatorOwnerAddress(uint256 _id, address _owner) returns()
func (_Operator *OperatorSession) SetNodeOperatorOwnerAddress(_id *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Operator.Contract.SetNodeOperatorOwnerAddress(&_Operator.TransactOpts, _id, _owner)
}

// SetNodeOperatorOwnerAddress is a paid mutator transaction binding the contract method 0x49790fff.
//
// Solidity: function setNodeOperatorOwnerAddress(uint256 _id, address _owner) returns()
func (_Operator *OperatorTransactorSession) SetNodeOperatorOwnerAddress(_id *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Operator.Contract.SetNodeOperatorOwnerAddress(&_Operator.TransactOpts, _id, _owner)
}

// SetNodeOperatorRegistrySetting is a paid mutator transaction binding the contract method 0x3ac58bba.
//
// Solidity: function setNodeOperatorRegistrySetting(address _dao, address _daoVaultAddress, address _liquidStakingContractAddress, address _operatorSlashContractAddress, address _vaultFactoryContractAddress, address _largeStakingContractAddress, uint256 _defaultOperatorCommission, uint256 _registrationFee, uint256 _permissionlessBlockNumber) returns()
func (_Operator *OperatorTransactor) SetNodeOperatorRegistrySetting(opts *bind.TransactOpts, _dao common.Address, _daoVaultAddress common.Address, _liquidStakingContractAddress common.Address, _operatorSlashContractAddress common.Address, _vaultFactoryContractAddress common.Address, _largeStakingContractAddress common.Address, _defaultOperatorCommission *big.Int, _registrationFee *big.Int, _permissionlessBlockNumber *big.Int) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "setNodeOperatorRegistrySetting", _dao, _daoVaultAddress, _liquidStakingContractAddress, _operatorSlashContractAddress, _vaultFactoryContractAddress, _largeStakingContractAddress, _defaultOperatorCommission, _registrationFee, _permissionlessBlockNumber)
}

// SetNodeOperatorRegistrySetting is a paid mutator transaction binding the contract method 0x3ac58bba.
//
// Solidity: function setNodeOperatorRegistrySetting(address _dao, address _daoVaultAddress, address _liquidStakingContractAddress, address _operatorSlashContractAddress, address _vaultFactoryContractAddress, address _largeStakingContractAddress, uint256 _defaultOperatorCommission, uint256 _registrationFee, uint256 _permissionlessBlockNumber) returns()
func (_Operator *OperatorSession) SetNodeOperatorRegistrySetting(_dao common.Address, _daoVaultAddress common.Address, _liquidStakingContractAddress common.Address, _operatorSlashContractAddress common.Address, _vaultFactoryContractAddress common.Address, _largeStakingContractAddress common.Address, _defaultOperatorCommission *big.Int, _registrationFee *big.Int, _permissionlessBlockNumber *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.SetNodeOperatorRegistrySetting(&_Operator.TransactOpts, _dao, _daoVaultAddress, _liquidStakingContractAddress, _operatorSlashContractAddress, _vaultFactoryContractAddress, _largeStakingContractAddress, _defaultOperatorCommission, _registrationFee, _permissionlessBlockNumber)
}

// SetNodeOperatorRegistrySetting is a paid mutator transaction binding the contract method 0x3ac58bba.
//
// Solidity: function setNodeOperatorRegistrySetting(address _dao, address _daoVaultAddress, address _liquidStakingContractAddress, address _operatorSlashContractAddress, address _vaultFactoryContractAddress, address _largeStakingContractAddress, uint256 _defaultOperatorCommission, uint256 _registrationFee, uint256 _permissionlessBlockNumber) returns()
func (_Operator *OperatorTransactorSession) SetNodeOperatorRegistrySetting(_dao common.Address, _daoVaultAddress common.Address, _liquidStakingContractAddress common.Address, _operatorSlashContractAddress common.Address, _vaultFactoryContractAddress common.Address, _largeStakingContractAddress common.Address, _defaultOperatorCommission *big.Int, _registrationFee *big.Int, _permissionlessBlockNumber *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.SetNodeOperatorRegistrySetting(&_Operator.TransactOpts, _dao, _daoVaultAddress, _liquidStakingContractAddress, _operatorSlashContractAddress, _vaultFactoryContractAddress, _largeStakingContractAddress, _defaultOperatorCommission, _registrationFee, _permissionlessBlockNumber)
}

// SetOperatorCommissionRate is a paid mutator transaction binding the contract method 0x3a6c1fe6.
//
// Solidity: function setOperatorCommissionRate(uint256 _operatorId, uint256 _rate) returns()
func (_Operator *OperatorTransactor) SetOperatorCommissionRate(opts *bind.TransactOpts, _operatorId *big.Int, _rate *big.Int) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "setOperatorCommissionRate", _operatorId, _rate)
}

// SetOperatorCommissionRate is a paid mutator transaction binding the contract method 0x3a6c1fe6.
//
// Solidity: function setOperatorCommissionRate(uint256 _operatorId, uint256 _rate) returns()
func (_Operator *OperatorSession) SetOperatorCommissionRate(_operatorId *big.Int, _rate *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.SetOperatorCommissionRate(&_Operator.TransactOpts, _operatorId, _rate)
}

// SetOperatorCommissionRate is a paid mutator transaction binding the contract method 0x3a6c1fe6.
//
// Solidity: function setOperatorCommissionRate(uint256 _operatorId, uint256 _rate) returns()
func (_Operator *OperatorTransactorSession) SetOperatorCommissionRate(_operatorId *big.Int, _rate *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.SetOperatorCommissionRate(&_Operator.TransactOpts, _operatorId, _rate)
}

// SetOperatorSetting is a paid mutator transaction binding the contract method 0x6d2cf614.
//
// Solidity: function setOperatorSetting(uint256 _id, string _name, address _controllerAddress, address[] _rewardAddresses, uint256[] _ratios) returns()
func (_Operator *OperatorTransactor) SetOperatorSetting(opts *bind.TransactOpts, _id *big.Int, _name string, _controllerAddress common.Address, _rewardAddresses []common.Address, _ratios []*big.Int) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "setOperatorSetting", _id, _name, _controllerAddress, _rewardAddresses, _ratios)
}

// SetOperatorSetting is a paid mutator transaction binding the contract method 0x6d2cf614.
//
// Solidity: function setOperatorSetting(uint256 _id, string _name, address _controllerAddress, address[] _rewardAddresses, uint256[] _ratios) returns()
func (_Operator *OperatorSession) SetOperatorSetting(_id *big.Int, _name string, _controllerAddress common.Address, _rewardAddresses []common.Address, _ratios []*big.Int) (*types.Transaction, error) {
	return _Operator.Contract.SetOperatorSetting(&_Operator.TransactOpts, _id, _name, _controllerAddress, _rewardAddresses, _ratios)
}

// SetOperatorSetting is a paid mutator transaction binding the contract method 0x6d2cf614.
//
// Solidity: function setOperatorSetting(uint256 _id, string _name, address _controllerAddress, address[] _rewardAddresses, uint256[] _ratios) returns()
func (_Operator *OperatorTransactorSession) SetOperatorSetting(_id *big.Int, _name string, _controllerAddress common.Address, _rewardAddresses []common.Address, _ratios []*big.Int) (*types.Transaction, error) {
	return _Operator.Contract.SetOperatorSetting(&_Operator.TransactOpts, _id, _name, _controllerAddress, _rewardAddresses, _ratios)
}

// SetTrustedOperator is a paid mutator transaction binding the contract method 0x804472b3.
//
// Solidity: function setTrustedOperator(uint256 _id) returns()
func (_Operator *OperatorTransactor) SetTrustedOperator(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "setTrustedOperator", _id)
}

// SetTrustedOperator is a paid mutator transaction binding the contract method 0x804472b3.
//
// Solidity: function setTrustedOperator(uint256 _id) returns()
func (_Operator *OperatorSession) SetTrustedOperator(_id *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.SetTrustedOperator(&_Operator.TransactOpts, _id)
}

// SetTrustedOperator is a paid mutator transaction binding the contract method 0x804472b3.
//
// Solidity: function setTrustedOperator(uint256 _id) returns()
func (_Operator *OperatorTransactorSession) SetTrustedOperator(_id *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.SetTrustedOperator(&_Operator.TransactOpts, _id)
}

// Slash is a paid mutator transaction binding the contract method 0xcccd3bb2.
//
// Solidity: function slash(uint256 _slashType, uint256[] _slashIds, uint256[] _operatorIds, uint256[] _amounts) returns()
func (_Operator *OperatorTransactor) Slash(opts *bind.TransactOpts, _slashType *big.Int, _slashIds []*big.Int, _operatorIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "slash", _slashType, _slashIds, _operatorIds, _amounts)
}

// Slash is a paid mutator transaction binding the contract method 0xcccd3bb2.
//
// Solidity: function slash(uint256 _slashType, uint256[] _slashIds, uint256[] _operatorIds, uint256[] _amounts) returns()
func (_Operator *OperatorSession) Slash(_slashType *big.Int, _slashIds []*big.Int, _operatorIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Operator.Contract.Slash(&_Operator.TransactOpts, _slashType, _slashIds, _operatorIds, _amounts)
}

// Slash is a paid mutator transaction binding the contract method 0xcccd3bb2.
//
// Solidity: function slash(uint256 _slashType, uint256[] _slashIds, uint256[] _operatorIds, uint256[] _amounts) returns()
func (_Operator *OperatorTransactorSession) Slash(_slashType *big.Int, _slashIds []*big.Int, _operatorIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Operator.Contract.Slash(&_Operator.TransactOpts, _slashType, _slashIds, _operatorIds, _amounts)
}

// SlashOfExitDelayed is a paid mutator transaction binding the contract method 0x055bcf32.
//
// Solidity: function slashOfExitDelayed(uint256 _operatorId, uint256 _amount) returns()
func (_Operator *OperatorTransactor) SlashOfExitDelayed(opts *bind.TransactOpts, _operatorId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "slashOfExitDelayed", _operatorId, _amount)
}

// SlashOfExitDelayed is a paid mutator transaction binding the contract method 0x055bcf32.
//
// Solidity: function slashOfExitDelayed(uint256 _operatorId, uint256 _amount) returns()
func (_Operator *OperatorSession) SlashOfExitDelayed(_operatorId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.SlashOfExitDelayed(&_Operator.TransactOpts, _operatorId, _amount)
}

// SlashOfExitDelayed is a paid mutator transaction binding the contract method 0x055bcf32.
//
// Solidity: function slashOfExitDelayed(uint256 _operatorId, uint256 _amount) returns()
func (_Operator *OperatorTransactorSession) SlashOfExitDelayed(_operatorId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.SlashOfExitDelayed(&_Operator.TransactOpts, _operatorId, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Operator *OperatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Operator *OperatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Operator.Contract.TransferOwnership(&_Operator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Operator *OperatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Operator.Contract.TransferOwnership(&_Operator.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Operator *OperatorTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Operator *OperatorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Operator.Contract.UpgradeTo(&_Operator.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Operator *OperatorTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Operator.Contract.UpgradeTo(&_Operator.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Operator *OperatorTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Operator *OperatorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Operator.Contract.UpgradeToAndCall(&_Operator.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Operator *OperatorTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Operator.Contract.UpgradeToAndCall(&_Operator.TransactOpts, newImplementation, data)
}

// WithdrawOperator is a paid mutator transaction binding the contract method 0xf6ba63f8.
//
// Solidity: function withdrawOperator(uint256 _operatorId, uint256 _withdrawAmount, address _to) returns()
func (_Operator *OperatorTransactor) WithdrawOperator(opts *bind.TransactOpts, _operatorId *big.Int, _withdrawAmount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "withdrawOperator", _operatorId, _withdrawAmount, _to)
}

// WithdrawOperator is a paid mutator transaction binding the contract method 0xf6ba63f8.
//
// Solidity: function withdrawOperator(uint256 _operatorId, uint256 _withdrawAmount, address _to) returns()
func (_Operator *OperatorSession) WithdrawOperator(_operatorId *big.Int, _withdrawAmount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Operator.Contract.WithdrawOperator(&_Operator.TransactOpts, _operatorId, _withdrawAmount, _to)
}

// WithdrawOperator is a paid mutator transaction binding the contract method 0xf6ba63f8.
//
// Solidity: function withdrawOperator(uint256 _operatorId, uint256 _withdrawAmount, address _to) returns()
func (_Operator *OperatorTransactorSession) WithdrawOperator(_operatorId *big.Int, _withdrawAmount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Operator.Contract.WithdrawOperator(&_Operator.TransactOpts, _operatorId, _withdrawAmount, _to)
}

// OperatorAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Operator contract.
type OperatorAdminChangedIterator struct {
	Event *OperatorAdminChanged // Event containing the contract specifics and raw log

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
func (it *OperatorAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorAdminChanged)
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
		it.Event = new(OperatorAdminChanged)
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
func (it *OperatorAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorAdminChanged represents a AdminChanged event raised by the Operator contract.
type OperatorAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Operator *OperatorFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*OperatorAdminChangedIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &OperatorAdminChangedIterator{contract: _Operator.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Operator *OperatorFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *OperatorAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorAdminChanged)
				if err := _Operator.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Operator *OperatorFilterer) ParseAdminChanged(log types.Log) (*OperatorAdminChanged, error) {
	event := new(OperatorAdminChanged)
	if err := _Operator.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Operator contract.
type OperatorBeaconUpgradedIterator struct {
	Event *OperatorBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *OperatorBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorBeaconUpgraded)
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
		it.Event = new(OperatorBeaconUpgraded)
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
func (it *OperatorBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorBeaconUpgraded represents a BeaconUpgraded event raised by the Operator contract.
type OperatorBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Operator *OperatorFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*OperatorBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Operator.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &OperatorBeaconUpgradedIterator{contract: _Operator.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Operator *OperatorFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *OperatorBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Operator.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorBeaconUpgraded)
				if err := _Operator.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Operator *OperatorFilterer) ParseBeaconUpgraded(log types.Log) (*OperatorBeaconUpgraded, error) {
	event := new(OperatorBeaconUpgraded)
	if err := _Operator.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorCommissionRateChangedIterator is returned from FilterCommissionRateChanged and is used to iterate over the raw logs and unpacked data for CommissionRateChanged events raised by the Operator contract.
type OperatorCommissionRateChangedIterator struct {
	Event *OperatorCommissionRateChanged // Event containing the contract specifics and raw log

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
func (it *OperatorCommissionRateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorCommissionRateChanged)
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
		it.Event = new(OperatorCommissionRateChanged)
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
func (it *OperatorCommissionRateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorCommissionRateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorCommissionRateChanged represents a CommissionRateChanged event raised by the Operator contract.
type OperatorCommissionRateChanged struct {
	OldRate *big.Int
	Rate    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCommissionRateChanged is a free log retrieval operation binding the contract event 0x74b81a9e0217358c4b0755d3032738dc303e980dde2980905160b1d8e7b68ba6.
//
// Solidity: event CommissionRateChanged(uint256 _oldRate, uint256 _rate)
func (_Operator *OperatorFilterer) FilterCommissionRateChanged(opts *bind.FilterOpts) (*OperatorCommissionRateChangedIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "CommissionRateChanged")
	if err != nil {
		return nil, err
	}
	return &OperatorCommissionRateChangedIterator{contract: _Operator.contract, event: "CommissionRateChanged", logs: logs, sub: sub}, nil
}

// WatchCommissionRateChanged is a free log subscription operation binding the contract event 0x74b81a9e0217358c4b0755d3032738dc303e980dde2980905160b1d8e7b68ba6.
//
// Solidity: event CommissionRateChanged(uint256 _oldRate, uint256 _rate)
func (_Operator *OperatorFilterer) WatchCommissionRateChanged(opts *bind.WatchOpts, sink chan<- *OperatorCommissionRateChanged) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "CommissionRateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorCommissionRateChanged)
				if err := _Operator.contract.UnpackLog(event, "CommissionRateChanged", log); err != nil {
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

// ParseCommissionRateChanged is a log parse operation binding the contract event 0x74b81a9e0217358c4b0755d3032738dc303e980dde2980905160b1d8e7b68ba6.
//
// Solidity: event CommissionRateChanged(uint256 _oldRate, uint256 _rate)
func (_Operator *OperatorFilterer) ParseCommissionRateChanged(log types.Log) (*OperatorCommissionRateChanged, error) {
	event := new(OperatorCommissionRateChanged)
	if err := _Operator.contract.UnpackLog(event, "CommissionRateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorDaoAddressChangedIterator is returned from FilterDaoAddressChanged and is used to iterate over the raw logs and unpacked data for DaoAddressChanged events raised by the Operator contract.
type OperatorDaoAddressChangedIterator struct {
	Event *OperatorDaoAddressChanged // Event containing the contract specifics and raw log

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
func (it *OperatorDaoAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorDaoAddressChanged)
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
		it.Event = new(OperatorDaoAddressChanged)
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
func (it *OperatorDaoAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorDaoAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorDaoAddressChanged represents a DaoAddressChanged event raised by the Operator contract.
type OperatorDaoAddressChanged struct {
	OldDao common.Address
	Dao    common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDaoAddressChanged is a free log retrieval operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address _oldDao, address _dao)
func (_Operator *OperatorFilterer) FilterDaoAddressChanged(opts *bind.FilterOpts) (*OperatorDaoAddressChangedIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "DaoAddressChanged")
	if err != nil {
		return nil, err
	}
	return &OperatorDaoAddressChangedIterator{contract: _Operator.contract, event: "DaoAddressChanged", logs: logs, sub: sub}, nil
}

// WatchDaoAddressChanged is a free log subscription operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address _oldDao, address _dao)
func (_Operator *OperatorFilterer) WatchDaoAddressChanged(opts *bind.WatchOpts, sink chan<- *OperatorDaoAddressChanged) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "DaoAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorDaoAddressChanged)
				if err := _Operator.contract.UnpackLog(event, "DaoAddressChanged", log); err != nil {
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
func (_Operator *OperatorFilterer) ParseDaoAddressChanged(log types.Log) (*OperatorDaoAddressChanged, error) {
	event := new(OperatorDaoAddressChanged)
	if err := _Operator.contract.UnpackLog(event, "DaoAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorDaoClaimRewardsIterator is returned from FilterDaoClaimRewards and is used to iterate over the raw logs and unpacked data for DaoClaimRewards events raised by the Operator contract.
type OperatorDaoClaimRewardsIterator struct {
	Event *OperatorDaoClaimRewards // Event containing the contract specifics and raw log

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
func (it *OperatorDaoClaimRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorDaoClaimRewards)
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
		it.Event = new(OperatorDaoClaimRewards)
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
func (it *OperatorDaoClaimRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorDaoClaimRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorDaoClaimRewards represents a DaoClaimRewards event raised by the Operator contract.
type OperatorDaoClaimRewards struct {
	OperatorId *big.Int
	Rewards    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDaoClaimRewards is a free log retrieval operation binding the contract event 0xd015384993a63fbb67be31c3c4491e03f64fa52369a08927fe6d4cba14286f21.
//
// Solidity: event DaoClaimRewards(uint256 _operatorId, uint256 _rewards)
func (_Operator *OperatorFilterer) FilterDaoClaimRewards(opts *bind.FilterOpts) (*OperatorDaoClaimRewardsIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "DaoClaimRewards")
	if err != nil {
		return nil, err
	}
	return &OperatorDaoClaimRewardsIterator{contract: _Operator.contract, event: "DaoClaimRewards", logs: logs, sub: sub}, nil
}

// WatchDaoClaimRewards is a free log subscription operation binding the contract event 0xd015384993a63fbb67be31c3c4491e03f64fa52369a08927fe6d4cba14286f21.
//
// Solidity: event DaoClaimRewards(uint256 _operatorId, uint256 _rewards)
func (_Operator *OperatorFilterer) WatchDaoClaimRewards(opts *bind.WatchOpts, sink chan<- *OperatorDaoClaimRewards) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "DaoClaimRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorDaoClaimRewards)
				if err := _Operator.contract.UnpackLog(event, "DaoClaimRewards", log); err != nil {
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

// ParseDaoClaimRewards is a log parse operation binding the contract event 0xd015384993a63fbb67be31c3c4491e03f64fa52369a08927fe6d4cba14286f21.
//
// Solidity: event DaoClaimRewards(uint256 _operatorId, uint256 _rewards)
func (_Operator *OperatorFilterer) ParseDaoClaimRewards(log types.Log) (*OperatorDaoClaimRewards, error) {
	event := new(OperatorDaoClaimRewards)
	if err := _Operator.contract.UnpackLog(event, "DaoClaimRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorDaoVaultAddressChangedIterator is returned from FilterDaoVaultAddressChanged and is used to iterate over the raw logs and unpacked data for DaoVaultAddressChanged events raised by the Operator contract.
type OperatorDaoVaultAddressChangedIterator struct {
	Event *OperatorDaoVaultAddressChanged // Event containing the contract specifics and raw log

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
func (it *OperatorDaoVaultAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorDaoVaultAddressChanged)
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
		it.Event = new(OperatorDaoVaultAddressChanged)
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
func (it *OperatorDaoVaultAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorDaoVaultAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorDaoVaultAddressChanged represents a DaoVaultAddressChanged event raised by the Operator contract.
type OperatorDaoVaultAddressChanged struct {
	OldDaoVaultAddress common.Address
	DaoVaultAddress    common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDaoVaultAddressChanged is a free log retrieval operation binding the contract event 0x74f93434acf49508438eb6f219ca22e7e1818b620ccb7acd411c8f520b27b642.
//
// Solidity: event DaoVaultAddressChanged(address _oldDaoVaultAddress, address _daoVaultAddress)
func (_Operator *OperatorFilterer) FilterDaoVaultAddressChanged(opts *bind.FilterOpts) (*OperatorDaoVaultAddressChangedIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "DaoVaultAddressChanged")
	if err != nil {
		return nil, err
	}
	return &OperatorDaoVaultAddressChangedIterator{contract: _Operator.contract, event: "DaoVaultAddressChanged", logs: logs, sub: sub}, nil
}

// WatchDaoVaultAddressChanged is a free log subscription operation binding the contract event 0x74f93434acf49508438eb6f219ca22e7e1818b620ccb7acd411c8f520b27b642.
//
// Solidity: event DaoVaultAddressChanged(address _oldDaoVaultAddress, address _daoVaultAddress)
func (_Operator *OperatorFilterer) WatchDaoVaultAddressChanged(opts *bind.WatchOpts, sink chan<- *OperatorDaoVaultAddressChanged) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "DaoVaultAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorDaoVaultAddressChanged)
				if err := _Operator.contract.UnpackLog(event, "DaoVaultAddressChanged", log); err != nil {
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

// ParseDaoVaultAddressChanged is a log parse operation binding the contract event 0x74f93434acf49508438eb6f219ca22e7e1818b620ccb7acd411c8f520b27b642.
//
// Solidity: event DaoVaultAddressChanged(address _oldDaoVaultAddress, address _daoVaultAddress)
func (_Operator *OperatorFilterer) ParseDaoVaultAddressChanged(log types.Log) (*OperatorDaoVaultAddressChanged, error) {
	event := new(OperatorDaoVaultAddressChanged)
	if err := _Operator.contract.UnpackLog(event, "DaoVaultAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorDefaultOperatorCommissionRateChangedIterator is returned from FilterDefaultOperatorCommissionRateChanged and is used to iterate over the raw logs and unpacked data for DefaultOperatorCommissionRateChanged events raised by the Operator contract.
type OperatorDefaultOperatorCommissionRateChangedIterator struct {
	Event *OperatorDefaultOperatorCommissionRateChanged // Event containing the contract specifics and raw log

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
func (it *OperatorDefaultOperatorCommissionRateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorDefaultOperatorCommissionRateChanged)
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
		it.Event = new(OperatorDefaultOperatorCommissionRateChanged)
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
func (it *OperatorDefaultOperatorCommissionRateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorDefaultOperatorCommissionRateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorDefaultOperatorCommissionRateChanged represents a DefaultOperatorCommissionRateChanged event raised by the Operator contract.
type OperatorDefaultOperatorCommissionRateChanged struct {
	OldDefaultOperatorCommission *big.Int
	DefaultOperatorCommission    *big.Int
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterDefaultOperatorCommissionRateChanged is a free log retrieval operation binding the contract event 0x339c55e5a65dd0e07a470cc3c647937de03f5c538c59a7b55fe6951592ed5523.
//
// Solidity: event DefaultOperatorCommissionRateChanged(uint256 _oldDefaultOperatorCommission, uint256 _defaultOperatorCommission)
func (_Operator *OperatorFilterer) FilterDefaultOperatorCommissionRateChanged(opts *bind.FilterOpts) (*OperatorDefaultOperatorCommissionRateChangedIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "DefaultOperatorCommissionRateChanged")
	if err != nil {
		return nil, err
	}
	return &OperatorDefaultOperatorCommissionRateChangedIterator{contract: _Operator.contract, event: "DefaultOperatorCommissionRateChanged", logs: logs, sub: sub}, nil
}

// WatchDefaultOperatorCommissionRateChanged is a free log subscription operation binding the contract event 0x339c55e5a65dd0e07a470cc3c647937de03f5c538c59a7b55fe6951592ed5523.
//
// Solidity: event DefaultOperatorCommissionRateChanged(uint256 _oldDefaultOperatorCommission, uint256 _defaultOperatorCommission)
func (_Operator *OperatorFilterer) WatchDefaultOperatorCommissionRateChanged(opts *bind.WatchOpts, sink chan<- *OperatorDefaultOperatorCommissionRateChanged) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "DefaultOperatorCommissionRateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorDefaultOperatorCommissionRateChanged)
				if err := _Operator.contract.UnpackLog(event, "DefaultOperatorCommissionRateChanged", log); err != nil {
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

// ParseDefaultOperatorCommissionRateChanged is a log parse operation binding the contract event 0x339c55e5a65dd0e07a470cc3c647937de03f5c538c59a7b55fe6951592ed5523.
//
// Solidity: event DefaultOperatorCommissionRateChanged(uint256 _oldDefaultOperatorCommission, uint256 _defaultOperatorCommission)
func (_Operator *OperatorFilterer) ParseDefaultOperatorCommissionRateChanged(log types.Log) (*OperatorDefaultOperatorCommissionRateChanged, error) {
	event := new(OperatorDefaultOperatorCommissionRateChanged)
	if err := _Operator.contract.UnpackLog(event, "DefaultOperatorCommissionRateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Operator contract.
type OperatorInitializedIterator struct {
	Event *OperatorInitialized // Event containing the contract specifics and raw log

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
func (it *OperatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorInitialized)
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
		it.Event = new(OperatorInitialized)
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
func (it *OperatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorInitialized represents a Initialized event raised by the Operator contract.
type OperatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Operator *OperatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*OperatorInitializedIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OperatorInitializedIterator{contract: _Operator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Operator *OperatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OperatorInitialized) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorInitialized)
				if err := _Operator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Operator *OperatorFilterer) ParseInitialized(log types.Log) (*OperatorInitialized, error) {
	event := new(OperatorInitialized)
	if err := _Operator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorLargeStakingChangedIterator is returned from FilterLargeStakingChanged and is used to iterate over the raw logs and unpacked data for LargeStakingChanged events raised by the Operator contract.
type OperatorLargeStakingChangedIterator struct {
	Event *OperatorLargeStakingChanged // Event containing the contract specifics and raw log

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
func (it *OperatorLargeStakingChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorLargeStakingChanged)
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
		it.Event = new(OperatorLargeStakingChanged)
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
func (it *OperatorLargeStakingChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorLargeStakingChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorLargeStakingChanged represents a LargeStakingChanged event raised by the Operator contract.
type OperatorLargeStakingChanged struct {
	OldLargeStakingContractAddress common.Address
	LargeStakingContractAddress    common.Address
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterLargeStakingChanged is a free log retrieval operation binding the contract event 0x796bfc23a386c56368910c17a5c85907974de0cbafabeaa08ac839b4ed8fe54c.
//
// Solidity: event LargeStakingChanged(address _oldLargeStakingContractAddress, address _largeStakingContractAddress)
func (_Operator *OperatorFilterer) FilterLargeStakingChanged(opts *bind.FilterOpts) (*OperatorLargeStakingChangedIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "LargeStakingChanged")
	if err != nil {
		return nil, err
	}
	return &OperatorLargeStakingChangedIterator{contract: _Operator.contract, event: "LargeStakingChanged", logs: logs, sub: sub}, nil
}

// WatchLargeStakingChanged is a free log subscription operation binding the contract event 0x796bfc23a386c56368910c17a5c85907974de0cbafabeaa08ac839b4ed8fe54c.
//
// Solidity: event LargeStakingChanged(address _oldLargeStakingContractAddress, address _largeStakingContractAddress)
func (_Operator *OperatorFilterer) WatchLargeStakingChanged(opts *bind.WatchOpts, sink chan<- *OperatorLargeStakingChanged) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "LargeStakingChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorLargeStakingChanged)
				if err := _Operator.contract.UnpackLog(event, "LargeStakingChanged", log); err != nil {
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

// ParseLargeStakingChanged is a log parse operation binding the contract event 0x796bfc23a386c56368910c17a5c85907974de0cbafabeaa08ac839b4ed8fe54c.
//
// Solidity: event LargeStakingChanged(address _oldLargeStakingContractAddress, address _largeStakingContractAddress)
func (_Operator *OperatorFilterer) ParseLargeStakingChanged(log types.Log) (*OperatorLargeStakingChanged, error) {
	event := new(OperatorLargeStakingChanged)
	if err := _Operator.contract.UnpackLog(event, "LargeStakingChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorLiquidStakingChangedIterator is returned from FilterLiquidStakingChanged and is used to iterate over the raw logs and unpacked data for LiquidStakingChanged events raised by the Operator contract.
type OperatorLiquidStakingChangedIterator struct {
	Event *OperatorLiquidStakingChanged // Event containing the contract specifics and raw log

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
func (it *OperatorLiquidStakingChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorLiquidStakingChanged)
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
		it.Event = new(OperatorLiquidStakingChanged)
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
func (it *OperatorLiquidStakingChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorLiquidStakingChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorLiquidStakingChanged represents a LiquidStakingChanged event raised by the Operator contract.
type OperatorLiquidStakingChanged struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLiquidStakingChanged is a free log retrieval operation binding the contract event 0x516e369f70685d2905d214a9e8567098c02a0e00f5a55bd30baca6b51d446cef.
//
// Solidity: event LiquidStakingChanged(address _from, address _to)
func (_Operator *OperatorFilterer) FilterLiquidStakingChanged(opts *bind.FilterOpts) (*OperatorLiquidStakingChangedIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "LiquidStakingChanged")
	if err != nil {
		return nil, err
	}
	return &OperatorLiquidStakingChangedIterator{contract: _Operator.contract, event: "LiquidStakingChanged", logs: logs, sub: sub}, nil
}

// WatchLiquidStakingChanged is a free log subscription operation binding the contract event 0x516e369f70685d2905d214a9e8567098c02a0e00f5a55bd30baca6b51d446cef.
//
// Solidity: event LiquidStakingChanged(address _from, address _to)
func (_Operator *OperatorFilterer) WatchLiquidStakingChanged(opts *bind.WatchOpts, sink chan<- *OperatorLiquidStakingChanged) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "LiquidStakingChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorLiquidStakingChanged)
				if err := _Operator.contract.UnpackLog(event, "LiquidStakingChanged", log); err != nil {
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
// Solidity: event LiquidStakingChanged(address _from, address _to)
func (_Operator *OperatorFilterer) ParseLiquidStakingChanged(log types.Log) (*OperatorLiquidStakingChanged, error) {
	event := new(OperatorLiquidStakingChanged)
	if err := _Operator.contract.UnpackLog(event, "LiquidStakingChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorNodeOperatorBlacklistRemoveIterator is returned from FilterNodeOperatorBlacklistRemove and is used to iterate over the raw logs and unpacked data for NodeOperatorBlacklistRemove events raised by the Operator contract.
type OperatorNodeOperatorBlacklistRemoveIterator struct {
	Event *OperatorNodeOperatorBlacklistRemove // Event containing the contract specifics and raw log

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
func (it *OperatorNodeOperatorBlacklistRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorNodeOperatorBlacklistRemove)
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
		it.Event = new(OperatorNodeOperatorBlacklistRemove)
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
func (it *OperatorNodeOperatorBlacklistRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorNodeOperatorBlacklistRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorNodeOperatorBlacklistRemove represents a NodeOperatorBlacklistRemove event raised by the Operator contract.
type OperatorNodeOperatorBlacklistRemove struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorBlacklistRemove is a free log retrieval operation binding the contract event 0xa101e301ec51792caf9bfff02addbb887e91ae6a8717851f42e187a03bfde386.
//
// Solidity: event NodeOperatorBlacklistRemove(uint256 _id)
func (_Operator *OperatorFilterer) FilterNodeOperatorBlacklistRemove(opts *bind.FilterOpts) (*OperatorNodeOperatorBlacklistRemoveIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "NodeOperatorBlacklistRemove")
	if err != nil {
		return nil, err
	}
	return &OperatorNodeOperatorBlacklistRemoveIterator{contract: _Operator.contract, event: "NodeOperatorBlacklistRemove", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorBlacklistRemove is a free log subscription operation binding the contract event 0xa101e301ec51792caf9bfff02addbb887e91ae6a8717851f42e187a03bfde386.
//
// Solidity: event NodeOperatorBlacklistRemove(uint256 _id)
func (_Operator *OperatorFilterer) WatchNodeOperatorBlacklistRemove(opts *bind.WatchOpts, sink chan<- *OperatorNodeOperatorBlacklistRemove) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "NodeOperatorBlacklistRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorNodeOperatorBlacklistRemove)
				if err := _Operator.contract.UnpackLog(event, "NodeOperatorBlacklistRemove", log); err != nil {
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

// ParseNodeOperatorBlacklistRemove is a log parse operation binding the contract event 0xa101e301ec51792caf9bfff02addbb887e91ae6a8717851f42e187a03bfde386.
//
// Solidity: event NodeOperatorBlacklistRemove(uint256 _id)
func (_Operator *OperatorFilterer) ParseNodeOperatorBlacklistRemove(log types.Log) (*OperatorNodeOperatorBlacklistRemove, error) {
	event := new(OperatorNodeOperatorBlacklistRemove)
	if err := _Operator.contract.UnpackLog(event, "NodeOperatorBlacklistRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorNodeOperatorBlacklistSetIterator is returned from FilterNodeOperatorBlacklistSet and is used to iterate over the raw logs and unpacked data for NodeOperatorBlacklistSet events raised by the Operator contract.
type OperatorNodeOperatorBlacklistSetIterator struct {
	Event *OperatorNodeOperatorBlacklistSet // Event containing the contract specifics and raw log

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
func (it *OperatorNodeOperatorBlacklistSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorNodeOperatorBlacklistSet)
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
		it.Event = new(OperatorNodeOperatorBlacklistSet)
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
func (it *OperatorNodeOperatorBlacklistSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorNodeOperatorBlacklistSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorNodeOperatorBlacklistSet represents a NodeOperatorBlacklistSet event raised by the Operator contract.
type OperatorNodeOperatorBlacklistSet struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorBlacklistSet is a free log retrieval operation binding the contract event 0x17d749abbae74182f40195483c16a42aa8a41e883b858307583c49e2d9294893.
//
// Solidity: event NodeOperatorBlacklistSet(uint256 _id)
func (_Operator *OperatorFilterer) FilterNodeOperatorBlacklistSet(opts *bind.FilterOpts) (*OperatorNodeOperatorBlacklistSetIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "NodeOperatorBlacklistSet")
	if err != nil {
		return nil, err
	}
	return &OperatorNodeOperatorBlacklistSetIterator{contract: _Operator.contract, event: "NodeOperatorBlacklistSet", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorBlacklistSet is a free log subscription operation binding the contract event 0x17d749abbae74182f40195483c16a42aa8a41e883b858307583c49e2d9294893.
//
// Solidity: event NodeOperatorBlacklistSet(uint256 _id)
func (_Operator *OperatorFilterer) WatchNodeOperatorBlacklistSet(opts *bind.WatchOpts, sink chan<- *OperatorNodeOperatorBlacklistSet) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "NodeOperatorBlacklistSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorNodeOperatorBlacklistSet)
				if err := _Operator.contract.UnpackLog(event, "NodeOperatorBlacklistSet", log); err != nil {
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

// ParseNodeOperatorBlacklistSet is a log parse operation binding the contract event 0x17d749abbae74182f40195483c16a42aa8a41e883b858307583c49e2d9294893.
//
// Solidity: event NodeOperatorBlacklistSet(uint256 _id)
func (_Operator *OperatorFilterer) ParseNodeOperatorBlacklistSet(log types.Log) (*OperatorNodeOperatorBlacklistSet, error) {
	event := new(OperatorNodeOperatorBlacklistSet)
	if err := _Operator.contract.UnpackLog(event, "NodeOperatorBlacklistSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorNodeOperatorControllerAddressSetIterator is returned from FilterNodeOperatorControllerAddressSet and is used to iterate over the raw logs and unpacked data for NodeOperatorControllerAddressSet events raised by the Operator contract.
type OperatorNodeOperatorControllerAddressSetIterator struct {
	Event *OperatorNodeOperatorControllerAddressSet // Event containing the contract specifics and raw log

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
func (it *OperatorNodeOperatorControllerAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorNodeOperatorControllerAddressSet)
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
		it.Event = new(OperatorNodeOperatorControllerAddressSet)
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
func (it *OperatorNodeOperatorControllerAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorNodeOperatorControllerAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorNodeOperatorControllerAddressSet represents a NodeOperatorControllerAddressSet event raised by the Operator contract.
type OperatorNodeOperatorControllerAddressSet struct {
	Id                *big.Int
	Name              string
	ControllerAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorControllerAddressSet is a free log retrieval operation binding the contract event 0xc8337b181d03c2f662d3f91ed8dfcbc968619bc49c385ffb3f7b410dc315d9d4.
//
// Solidity: event NodeOperatorControllerAddressSet(uint256 _id, string _name, address _controllerAddress)
func (_Operator *OperatorFilterer) FilterNodeOperatorControllerAddressSet(opts *bind.FilterOpts) (*OperatorNodeOperatorControllerAddressSetIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "NodeOperatorControllerAddressSet")
	if err != nil {
		return nil, err
	}
	return &OperatorNodeOperatorControllerAddressSetIterator{contract: _Operator.contract, event: "NodeOperatorControllerAddressSet", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorControllerAddressSet is a free log subscription operation binding the contract event 0xc8337b181d03c2f662d3f91ed8dfcbc968619bc49c385ffb3f7b410dc315d9d4.
//
// Solidity: event NodeOperatorControllerAddressSet(uint256 _id, string _name, address _controllerAddress)
func (_Operator *OperatorFilterer) WatchNodeOperatorControllerAddressSet(opts *bind.WatchOpts, sink chan<- *OperatorNodeOperatorControllerAddressSet) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "NodeOperatorControllerAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorNodeOperatorControllerAddressSet)
				if err := _Operator.contract.UnpackLog(event, "NodeOperatorControllerAddressSet", log); err != nil {
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

// ParseNodeOperatorControllerAddressSet is a log parse operation binding the contract event 0xc8337b181d03c2f662d3f91ed8dfcbc968619bc49c385ffb3f7b410dc315d9d4.
//
// Solidity: event NodeOperatorControllerAddressSet(uint256 _id, string _name, address _controllerAddress)
func (_Operator *OperatorFilterer) ParseNodeOperatorControllerAddressSet(log types.Log) (*OperatorNodeOperatorControllerAddressSet, error) {
	event := new(OperatorNodeOperatorControllerAddressSet)
	if err := _Operator.contract.UnpackLog(event, "NodeOperatorControllerAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorNodeOperatorNameSetIterator is returned from FilterNodeOperatorNameSet and is used to iterate over the raw logs and unpacked data for NodeOperatorNameSet events raised by the Operator contract.
type OperatorNodeOperatorNameSetIterator struct {
	Event *OperatorNodeOperatorNameSet // Event containing the contract specifics and raw log

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
func (it *OperatorNodeOperatorNameSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorNodeOperatorNameSet)
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
		it.Event = new(OperatorNodeOperatorNameSet)
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
func (it *OperatorNodeOperatorNameSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorNodeOperatorNameSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorNodeOperatorNameSet represents a NodeOperatorNameSet event raised by the Operator contract.
type OperatorNodeOperatorNameSet struct {
	Id   *big.Int
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorNameSet is a free log retrieval operation binding the contract event 0xcb16868f4831cc58a28d413f658752a2958bd1f50e94ed6391716b936c48093b.
//
// Solidity: event NodeOperatorNameSet(uint256 _id, string _name)
func (_Operator *OperatorFilterer) FilterNodeOperatorNameSet(opts *bind.FilterOpts) (*OperatorNodeOperatorNameSetIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "NodeOperatorNameSet")
	if err != nil {
		return nil, err
	}
	return &OperatorNodeOperatorNameSetIterator{contract: _Operator.contract, event: "NodeOperatorNameSet", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorNameSet is a free log subscription operation binding the contract event 0xcb16868f4831cc58a28d413f658752a2958bd1f50e94ed6391716b936c48093b.
//
// Solidity: event NodeOperatorNameSet(uint256 _id, string _name)
func (_Operator *OperatorFilterer) WatchNodeOperatorNameSet(opts *bind.WatchOpts, sink chan<- *OperatorNodeOperatorNameSet) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "NodeOperatorNameSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorNodeOperatorNameSet)
				if err := _Operator.contract.UnpackLog(event, "NodeOperatorNameSet", log); err != nil {
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

// ParseNodeOperatorNameSet is a log parse operation binding the contract event 0xcb16868f4831cc58a28d413f658752a2958bd1f50e94ed6391716b936c48093b.
//
// Solidity: event NodeOperatorNameSet(uint256 _id, string _name)
func (_Operator *OperatorFilterer) ParseNodeOperatorNameSet(log types.Log) (*OperatorNodeOperatorNameSet, error) {
	event := new(OperatorNodeOperatorNameSet)
	if err := _Operator.contract.UnpackLog(event, "NodeOperatorNameSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorNodeOperatorOwnerAddressSetIterator is returned from FilterNodeOperatorOwnerAddressSet and is used to iterate over the raw logs and unpacked data for NodeOperatorOwnerAddressSet events raised by the Operator contract.
type OperatorNodeOperatorOwnerAddressSetIterator struct {
	Event *OperatorNodeOperatorOwnerAddressSet // Event containing the contract specifics and raw log

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
func (it *OperatorNodeOperatorOwnerAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorNodeOperatorOwnerAddressSet)
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
		it.Event = new(OperatorNodeOperatorOwnerAddressSet)
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
func (it *OperatorNodeOperatorOwnerAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorNodeOperatorOwnerAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorNodeOperatorOwnerAddressSet represents a NodeOperatorOwnerAddressSet event raised by the Operator contract.
type OperatorNodeOperatorOwnerAddressSet struct {
	Id           *big.Int
	Name         string
	OwnerAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorOwnerAddressSet is a free log retrieval operation binding the contract event 0x7d89e4f8f7e7ed69fd4dd609ed0267a5d082e1307760625946dddbf2a40883b7.
//
// Solidity: event NodeOperatorOwnerAddressSet(uint256 _id, string _name, address _ownerAddress)
func (_Operator *OperatorFilterer) FilterNodeOperatorOwnerAddressSet(opts *bind.FilterOpts) (*OperatorNodeOperatorOwnerAddressSetIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "NodeOperatorOwnerAddressSet")
	if err != nil {
		return nil, err
	}
	return &OperatorNodeOperatorOwnerAddressSetIterator{contract: _Operator.contract, event: "NodeOperatorOwnerAddressSet", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorOwnerAddressSet is a free log subscription operation binding the contract event 0x7d89e4f8f7e7ed69fd4dd609ed0267a5d082e1307760625946dddbf2a40883b7.
//
// Solidity: event NodeOperatorOwnerAddressSet(uint256 _id, string _name, address _ownerAddress)
func (_Operator *OperatorFilterer) WatchNodeOperatorOwnerAddressSet(opts *bind.WatchOpts, sink chan<- *OperatorNodeOperatorOwnerAddressSet) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "NodeOperatorOwnerAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorNodeOperatorOwnerAddressSet)
				if err := _Operator.contract.UnpackLog(event, "NodeOperatorOwnerAddressSet", log); err != nil {
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

// ParseNodeOperatorOwnerAddressSet is a log parse operation binding the contract event 0x7d89e4f8f7e7ed69fd4dd609ed0267a5d082e1307760625946dddbf2a40883b7.
//
// Solidity: event NodeOperatorOwnerAddressSet(uint256 _id, string _name, address _ownerAddress)
func (_Operator *OperatorFilterer) ParseNodeOperatorOwnerAddressSet(log types.Log) (*OperatorNodeOperatorOwnerAddressSet, error) {
	event := new(OperatorNodeOperatorOwnerAddressSet)
	if err := _Operator.contract.UnpackLog(event, "NodeOperatorOwnerAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorNodeOperatorRegisteredIterator is returned from FilterNodeOperatorRegistered and is used to iterate over the raw logs and unpacked data for NodeOperatorRegistered events raised by the Operator contract.
type OperatorNodeOperatorRegisteredIterator struct {
	Event *OperatorNodeOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *OperatorNodeOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorNodeOperatorRegistered)
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
		it.Event = new(OperatorNodeOperatorRegistered)
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
func (it *OperatorNodeOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorNodeOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorNodeOperatorRegistered represents a NodeOperatorRegistered event raised by the Operator contract.
type OperatorNodeOperatorRegistered struct {
	Id                   *big.Int
	Name                 string
	ControllerAddress    common.Address
	VaultContractAddress common.Address
	RewardAddresses      []common.Address
	Ratios               []*big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorRegistered is a free log retrieval operation binding the contract event 0xe4cb1b0e99980eb8118feab9ff9c65775219ab1b75c30de0a11a76652c39edf1.
//
// Solidity: event NodeOperatorRegistered(uint256 _id, string _name, address _controllerAddress, address _vaultContractAddress, address[] _rewardAddresses, uint256[] _ratios)
func (_Operator *OperatorFilterer) FilterNodeOperatorRegistered(opts *bind.FilterOpts) (*OperatorNodeOperatorRegisteredIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "NodeOperatorRegistered")
	if err != nil {
		return nil, err
	}
	return &OperatorNodeOperatorRegisteredIterator{contract: _Operator.contract, event: "NodeOperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorRegistered is a free log subscription operation binding the contract event 0xe4cb1b0e99980eb8118feab9ff9c65775219ab1b75c30de0a11a76652c39edf1.
//
// Solidity: event NodeOperatorRegistered(uint256 _id, string _name, address _controllerAddress, address _vaultContractAddress, address[] _rewardAddresses, uint256[] _ratios)
func (_Operator *OperatorFilterer) WatchNodeOperatorRegistered(opts *bind.WatchOpts, sink chan<- *OperatorNodeOperatorRegistered) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "NodeOperatorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorNodeOperatorRegistered)
				if err := _Operator.contract.UnpackLog(event, "NodeOperatorRegistered", log); err != nil {
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

// ParseNodeOperatorRegistered is a log parse operation binding the contract event 0xe4cb1b0e99980eb8118feab9ff9c65775219ab1b75c30de0a11a76652c39edf1.
//
// Solidity: event NodeOperatorRegistered(uint256 _id, string _name, address _controllerAddress, address _vaultContractAddress, address[] _rewardAddresses, uint256[] _ratios)
func (_Operator *OperatorFilterer) ParseNodeOperatorRegistered(log types.Log) (*OperatorNodeOperatorRegistered, error) {
	event := new(OperatorNodeOperatorRegistered)
	if err := _Operator.contract.UnpackLog(event, "NodeOperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorNodeOperatorRewardAddressSetIterator is returned from FilterNodeOperatorRewardAddressSet and is used to iterate over the raw logs and unpacked data for NodeOperatorRewardAddressSet events raised by the Operator contract.
type OperatorNodeOperatorRewardAddressSetIterator struct {
	Event *OperatorNodeOperatorRewardAddressSet // Event containing the contract specifics and raw log

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
func (it *OperatorNodeOperatorRewardAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorNodeOperatorRewardAddressSet)
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
		it.Event = new(OperatorNodeOperatorRewardAddressSet)
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
func (it *OperatorNodeOperatorRewardAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorNodeOperatorRewardAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorNodeOperatorRewardAddressSet represents a NodeOperatorRewardAddressSet event raised by the Operator contract.
type OperatorNodeOperatorRewardAddressSet struct {
	Id              *big.Int
	RewardAddresses []common.Address
	Ratios          []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorRewardAddressSet is a free log retrieval operation binding the contract event 0x53d07c39db33e0d8b3562e1211bc5a6f67b7581db206c7386641d1482e2422c1.
//
// Solidity: event NodeOperatorRewardAddressSet(uint256 _id, address[] _rewardAddresses, uint256[] _ratios)
func (_Operator *OperatorFilterer) FilterNodeOperatorRewardAddressSet(opts *bind.FilterOpts) (*OperatorNodeOperatorRewardAddressSetIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "NodeOperatorRewardAddressSet")
	if err != nil {
		return nil, err
	}
	return &OperatorNodeOperatorRewardAddressSetIterator{contract: _Operator.contract, event: "NodeOperatorRewardAddressSet", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorRewardAddressSet is a free log subscription operation binding the contract event 0x53d07c39db33e0d8b3562e1211bc5a6f67b7581db206c7386641d1482e2422c1.
//
// Solidity: event NodeOperatorRewardAddressSet(uint256 _id, address[] _rewardAddresses, uint256[] _ratios)
func (_Operator *OperatorFilterer) WatchNodeOperatorRewardAddressSet(opts *bind.WatchOpts, sink chan<- *OperatorNodeOperatorRewardAddressSet) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "NodeOperatorRewardAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorNodeOperatorRewardAddressSet)
				if err := _Operator.contract.UnpackLog(event, "NodeOperatorRewardAddressSet", log); err != nil {
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

// ParseNodeOperatorRewardAddressSet is a log parse operation binding the contract event 0x53d07c39db33e0d8b3562e1211bc5a6f67b7581db206c7386641d1482e2422c1.
//
// Solidity: event NodeOperatorRewardAddressSet(uint256 _id, address[] _rewardAddresses, uint256[] _ratios)
func (_Operator *OperatorFilterer) ParseNodeOperatorRewardAddressSet(log types.Log) (*OperatorNodeOperatorRewardAddressSet, error) {
	event := new(OperatorNodeOperatorRewardAddressSet)
	if err := _Operator.contract.UnpackLog(event, "NodeOperatorRewardAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorNodeOperatorTrustedRemoveIterator is returned from FilterNodeOperatorTrustedRemove and is used to iterate over the raw logs and unpacked data for NodeOperatorTrustedRemove events raised by the Operator contract.
type OperatorNodeOperatorTrustedRemoveIterator struct {
	Event *OperatorNodeOperatorTrustedRemove // Event containing the contract specifics and raw log

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
func (it *OperatorNodeOperatorTrustedRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorNodeOperatorTrustedRemove)
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
		it.Event = new(OperatorNodeOperatorTrustedRemove)
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
func (it *OperatorNodeOperatorTrustedRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorNodeOperatorTrustedRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorNodeOperatorTrustedRemove represents a NodeOperatorTrustedRemove event raised by the Operator contract.
type OperatorNodeOperatorTrustedRemove struct {
	Id      *big.Int
	Name    string
	Trusted bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorTrustedRemove is a free log retrieval operation binding the contract event 0x2caee3bf9bf3b301308a164799ad83038b72182327103fb2de9beb26742b91f0.
//
// Solidity: event NodeOperatorTrustedRemove(uint256 _id, string _name, bool _trusted)
func (_Operator *OperatorFilterer) FilterNodeOperatorTrustedRemove(opts *bind.FilterOpts) (*OperatorNodeOperatorTrustedRemoveIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "NodeOperatorTrustedRemove")
	if err != nil {
		return nil, err
	}
	return &OperatorNodeOperatorTrustedRemoveIterator{contract: _Operator.contract, event: "NodeOperatorTrustedRemove", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorTrustedRemove is a free log subscription operation binding the contract event 0x2caee3bf9bf3b301308a164799ad83038b72182327103fb2de9beb26742b91f0.
//
// Solidity: event NodeOperatorTrustedRemove(uint256 _id, string _name, bool _trusted)
func (_Operator *OperatorFilterer) WatchNodeOperatorTrustedRemove(opts *bind.WatchOpts, sink chan<- *OperatorNodeOperatorTrustedRemove) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "NodeOperatorTrustedRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorNodeOperatorTrustedRemove)
				if err := _Operator.contract.UnpackLog(event, "NodeOperatorTrustedRemove", log); err != nil {
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

// ParseNodeOperatorTrustedRemove is a log parse operation binding the contract event 0x2caee3bf9bf3b301308a164799ad83038b72182327103fb2de9beb26742b91f0.
//
// Solidity: event NodeOperatorTrustedRemove(uint256 _id, string _name, bool _trusted)
func (_Operator *OperatorFilterer) ParseNodeOperatorTrustedRemove(log types.Log) (*OperatorNodeOperatorTrustedRemove, error) {
	event := new(OperatorNodeOperatorTrustedRemove)
	if err := _Operator.contract.UnpackLog(event, "NodeOperatorTrustedRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorNodeOperatorTrustedSetIterator is returned from FilterNodeOperatorTrustedSet and is used to iterate over the raw logs and unpacked data for NodeOperatorTrustedSet events raised by the Operator contract.
type OperatorNodeOperatorTrustedSetIterator struct {
	Event *OperatorNodeOperatorTrustedSet // Event containing the contract specifics and raw log

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
func (it *OperatorNodeOperatorTrustedSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorNodeOperatorTrustedSet)
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
		it.Event = new(OperatorNodeOperatorTrustedSet)
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
func (it *OperatorNodeOperatorTrustedSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorNodeOperatorTrustedSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorNodeOperatorTrustedSet represents a NodeOperatorTrustedSet event raised by the Operator contract.
type OperatorNodeOperatorTrustedSet struct {
	Id      *big.Int
	Name    string
	Trusted bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorTrustedSet is a free log retrieval operation binding the contract event 0x74ae602dfb36b44b653b50a300faa8211da632bf6abc1a0c26f3f1e8c3e033fd.
//
// Solidity: event NodeOperatorTrustedSet(uint256 _id, string _name, bool _trusted)
func (_Operator *OperatorFilterer) FilterNodeOperatorTrustedSet(opts *bind.FilterOpts) (*OperatorNodeOperatorTrustedSetIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "NodeOperatorTrustedSet")
	if err != nil {
		return nil, err
	}
	return &OperatorNodeOperatorTrustedSetIterator{contract: _Operator.contract, event: "NodeOperatorTrustedSet", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorTrustedSet is a free log subscription operation binding the contract event 0x74ae602dfb36b44b653b50a300faa8211da632bf6abc1a0c26f3f1e8c3e033fd.
//
// Solidity: event NodeOperatorTrustedSet(uint256 _id, string _name, bool _trusted)
func (_Operator *OperatorFilterer) WatchNodeOperatorTrustedSet(opts *bind.WatchOpts, sink chan<- *OperatorNodeOperatorTrustedSet) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "NodeOperatorTrustedSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorNodeOperatorTrustedSet)
				if err := _Operator.contract.UnpackLog(event, "NodeOperatorTrustedSet", log); err != nil {
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

// ParseNodeOperatorTrustedSet is a log parse operation binding the contract event 0x74ae602dfb36b44b653b50a300faa8211da632bf6abc1a0c26f3f1e8c3e033fd.
//
// Solidity: event NodeOperatorTrustedSet(uint256 _id, string _name, bool _trusted)
func (_Operator *OperatorFilterer) ParseNodeOperatorTrustedSet(log types.Log) (*OperatorNodeOperatorTrustedSet, error) {
	event := new(OperatorNodeOperatorTrustedSet)
	if err := _Operator.contract.UnpackLog(event, "NodeOperatorTrustedSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorOperatorArrearsIncreaseIterator is returned from FilterOperatorArrearsIncrease and is used to iterate over the raw logs and unpacked data for OperatorArrearsIncrease events raised by the Operator contract.
type OperatorOperatorArrearsIncreaseIterator struct {
	Event *OperatorOperatorArrearsIncrease // Event containing the contract specifics and raw log

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
func (it *OperatorOperatorArrearsIncreaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorOperatorArrearsIncrease)
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
		it.Event = new(OperatorOperatorArrearsIncrease)
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
func (it *OperatorOperatorArrearsIncreaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorOperatorArrearsIncreaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorOperatorArrearsIncrease represents a OperatorArrearsIncrease event raised by the Operator contract.
type OperatorOperatorArrearsIncrease struct {
	OperatorId *big.Int
	Value      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorArrearsIncrease is a free log retrieval operation binding the contract event 0x54a27b8571bb65717906f641f461a1dc2e874d7a374fad2a2cba1deeef69c1cb.
//
// Solidity: event OperatorArrearsIncrease(uint256 _operatorId, uint256 value)
func (_Operator *OperatorFilterer) FilterOperatorArrearsIncrease(opts *bind.FilterOpts) (*OperatorOperatorArrearsIncreaseIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "OperatorArrearsIncrease")
	if err != nil {
		return nil, err
	}
	return &OperatorOperatorArrearsIncreaseIterator{contract: _Operator.contract, event: "OperatorArrearsIncrease", logs: logs, sub: sub}, nil
}

// WatchOperatorArrearsIncrease is a free log subscription operation binding the contract event 0x54a27b8571bb65717906f641f461a1dc2e874d7a374fad2a2cba1deeef69c1cb.
//
// Solidity: event OperatorArrearsIncrease(uint256 _operatorId, uint256 value)
func (_Operator *OperatorFilterer) WatchOperatorArrearsIncrease(opts *bind.WatchOpts, sink chan<- *OperatorOperatorArrearsIncrease) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "OperatorArrearsIncrease")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorOperatorArrearsIncrease)
				if err := _Operator.contract.UnpackLog(event, "OperatorArrearsIncrease", log); err != nil {
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

// ParseOperatorArrearsIncrease is a log parse operation binding the contract event 0x54a27b8571bb65717906f641f461a1dc2e874d7a374fad2a2cba1deeef69c1cb.
//
// Solidity: event OperatorArrearsIncrease(uint256 _operatorId, uint256 value)
func (_Operator *OperatorFilterer) ParseOperatorArrearsIncrease(log types.Log) (*OperatorOperatorArrearsIncrease, error) {
	event := new(OperatorOperatorArrearsIncrease)
	if err := _Operator.contract.UnpackLog(event, "OperatorArrearsIncrease", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorOperatorArrearsReduceIterator is returned from FilterOperatorArrearsReduce and is used to iterate over the raw logs and unpacked data for OperatorArrearsReduce events raised by the Operator contract.
type OperatorOperatorArrearsReduceIterator struct {
	Event *OperatorOperatorArrearsReduce // Event containing the contract specifics and raw log

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
func (it *OperatorOperatorArrearsReduceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorOperatorArrearsReduce)
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
		it.Event = new(OperatorOperatorArrearsReduce)
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
func (it *OperatorOperatorArrearsReduceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorOperatorArrearsReduceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorOperatorArrearsReduce represents a OperatorArrearsReduce event raised by the Operator contract.
type OperatorOperatorArrearsReduce struct {
	OperatorId *big.Int
	Value      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorArrearsReduce is a free log retrieval operation binding the contract event 0x8c2d8ca0a282777673a0646723ec03f18c4af36204576f0692af44e002eb5c7c.
//
// Solidity: event OperatorArrearsReduce(uint256 _operatorId, uint256 value)
func (_Operator *OperatorFilterer) FilterOperatorArrearsReduce(opts *bind.FilterOpts) (*OperatorOperatorArrearsReduceIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "OperatorArrearsReduce")
	if err != nil {
		return nil, err
	}
	return &OperatorOperatorArrearsReduceIterator{contract: _Operator.contract, event: "OperatorArrearsReduce", logs: logs, sub: sub}, nil
}

// WatchOperatorArrearsReduce is a free log subscription operation binding the contract event 0x8c2d8ca0a282777673a0646723ec03f18c4af36204576f0692af44e002eb5c7c.
//
// Solidity: event OperatorArrearsReduce(uint256 _operatorId, uint256 value)
func (_Operator *OperatorFilterer) WatchOperatorArrearsReduce(opts *bind.WatchOpts, sink chan<- *OperatorOperatorArrearsReduce) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "OperatorArrearsReduce")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorOperatorArrearsReduce)
				if err := _Operator.contract.UnpackLog(event, "OperatorArrearsReduce", log); err != nil {
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

// ParseOperatorArrearsReduce is a log parse operation binding the contract event 0x8c2d8ca0a282777673a0646723ec03f18c4af36204576f0692af44e002eb5c7c.
//
// Solidity: event OperatorArrearsReduce(uint256 _operatorId, uint256 value)
func (_Operator *OperatorFilterer) ParseOperatorArrearsReduce(log types.Log) (*OperatorOperatorArrearsReduce, error) {
	event := new(OperatorOperatorArrearsReduce)
	if err := _Operator.contract.UnpackLog(event, "OperatorArrearsReduce", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorOperatorClaimRewardsIterator is returned from FilterOperatorClaimRewards and is used to iterate over the raw logs and unpacked data for OperatorClaimRewards events raised by the Operator contract.
type OperatorOperatorClaimRewardsIterator struct {
	Event *OperatorOperatorClaimRewards // Event containing the contract specifics and raw log

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
func (it *OperatorOperatorClaimRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorOperatorClaimRewards)
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
		it.Event = new(OperatorOperatorClaimRewards)
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
func (it *OperatorOperatorClaimRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorOperatorClaimRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorOperatorClaimRewards represents a OperatorClaimRewards event raised by the Operator contract.
type OperatorOperatorClaimRewards struct {
	OperatorId *big.Int
	Rewards    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorClaimRewards is a free log retrieval operation binding the contract event 0x9af81b3bd245e1f5301f53c691f532ada01ad22ed150ecf59d65eebf084a635f.
//
// Solidity: event OperatorClaimRewards(uint256 _operatorId, uint256 _rewards)
func (_Operator *OperatorFilterer) FilterOperatorClaimRewards(opts *bind.FilterOpts) (*OperatorOperatorClaimRewardsIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "OperatorClaimRewards")
	if err != nil {
		return nil, err
	}
	return &OperatorOperatorClaimRewardsIterator{contract: _Operator.contract, event: "OperatorClaimRewards", logs: logs, sub: sub}, nil
}

// WatchOperatorClaimRewards is a free log subscription operation binding the contract event 0x9af81b3bd245e1f5301f53c691f532ada01ad22ed150ecf59d65eebf084a635f.
//
// Solidity: event OperatorClaimRewards(uint256 _operatorId, uint256 _rewards)
func (_Operator *OperatorFilterer) WatchOperatorClaimRewards(opts *bind.WatchOpts, sink chan<- *OperatorOperatorClaimRewards) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "OperatorClaimRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorOperatorClaimRewards)
				if err := _Operator.contract.UnpackLog(event, "OperatorClaimRewards", log); err != nil {
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

// ParseOperatorClaimRewards is a log parse operation binding the contract event 0x9af81b3bd245e1f5301f53c691f532ada01ad22ed150ecf59d65eebf084a635f.
//
// Solidity: event OperatorClaimRewards(uint256 _operatorId, uint256 _rewards)
func (_Operator *OperatorFilterer) ParseOperatorClaimRewards(log types.Log) (*OperatorOperatorClaimRewards, error) {
	event := new(OperatorOperatorClaimRewards)
	if err := _Operator.contract.UnpackLog(event, "OperatorClaimRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorOperatorQuitIterator is returned from FilterOperatorQuit and is used to iterate over the raw logs and unpacked data for OperatorQuit events raised by the Operator contract.
type OperatorOperatorQuitIterator struct {
	Event *OperatorOperatorQuit // Event containing the contract specifics and raw log

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
func (it *OperatorOperatorQuitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorOperatorQuit)
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
		it.Event = new(OperatorOperatorQuit)
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
func (it *OperatorOperatorQuitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorOperatorQuitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorOperatorQuit represents a OperatorQuit event raised by the Operator contract.
type OperatorOperatorQuit struct {
	OperatorId *big.Int
	NowVault   *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorQuit is a free log retrieval operation binding the contract event 0x97e392587299cf79696405e61997d032d496edd9fad69d96968380eca40c7ce7.
//
// Solidity: event OperatorQuit(uint256 _operatorId, uint256 _nowVault, address _to)
func (_Operator *OperatorFilterer) FilterOperatorQuit(opts *bind.FilterOpts) (*OperatorOperatorQuitIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "OperatorQuit")
	if err != nil {
		return nil, err
	}
	return &OperatorOperatorQuitIterator{contract: _Operator.contract, event: "OperatorQuit", logs: logs, sub: sub}, nil
}

// WatchOperatorQuit is a free log subscription operation binding the contract event 0x97e392587299cf79696405e61997d032d496edd9fad69d96968380eca40c7ce7.
//
// Solidity: event OperatorQuit(uint256 _operatorId, uint256 _nowVault, address _to)
func (_Operator *OperatorFilterer) WatchOperatorQuit(opts *bind.WatchOpts, sink chan<- *OperatorOperatorQuit) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "OperatorQuit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorOperatorQuit)
				if err := _Operator.contract.UnpackLog(event, "OperatorQuit", log); err != nil {
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

// ParseOperatorQuit is a log parse operation binding the contract event 0x97e392587299cf79696405e61997d032d496edd9fad69d96968380eca40c7ce7.
//
// Solidity: event OperatorQuit(uint256 _operatorId, uint256 _nowVault, address _to)
func (_Operator *OperatorFilterer) ParseOperatorQuit(log types.Log) (*OperatorOperatorQuit, error) {
	event := new(OperatorOperatorQuit)
	if err := _Operator.contract.UnpackLog(event, "OperatorQuit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorOperatorSlashContractSetIterator is returned from FilterOperatorSlashContractSet and is used to iterate over the raw logs and unpacked data for OperatorSlashContractSet events raised by the Operator contract.
type OperatorOperatorSlashContractSetIterator struct {
	Event *OperatorOperatorSlashContractSet // Event containing the contract specifics and raw log

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
func (it *OperatorOperatorSlashContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorOperatorSlashContractSet)
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
		it.Event = new(OperatorOperatorSlashContractSet)
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
func (it *OperatorOperatorSlashContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorOperatorSlashContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorOperatorSlashContractSet represents a OperatorSlashContractSet event raised by the Operator contract.
type OperatorOperatorSlashContractSet struct {
	OldOperatorSlashContract     common.Address
	OperatorSlashContractAddress common.Address
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashContractSet is a free log retrieval operation binding the contract event 0x928cdf348ed0c0e57da3069ecf8f21acdf8b26364a2d1334ab766b8576fd34ba.
//
// Solidity: event OperatorSlashContractSet(address _oldOperatorSlashContract, address _operatorSlashContractAddress)
func (_Operator *OperatorFilterer) FilterOperatorSlashContractSet(opts *bind.FilterOpts) (*OperatorOperatorSlashContractSetIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "OperatorSlashContractSet")
	if err != nil {
		return nil, err
	}
	return &OperatorOperatorSlashContractSetIterator{contract: _Operator.contract, event: "OperatorSlashContractSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashContractSet is a free log subscription operation binding the contract event 0x928cdf348ed0c0e57da3069ecf8f21acdf8b26364a2d1334ab766b8576fd34ba.
//
// Solidity: event OperatorSlashContractSet(address _oldOperatorSlashContract, address _operatorSlashContractAddress)
func (_Operator *OperatorFilterer) WatchOperatorSlashContractSet(opts *bind.WatchOpts, sink chan<- *OperatorOperatorSlashContractSet) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "OperatorSlashContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorOperatorSlashContractSet)
				if err := _Operator.contract.UnpackLog(event, "OperatorSlashContractSet", log); err != nil {
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

// ParseOperatorSlashContractSet is a log parse operation binding the contract event 0x928cdf348ed0c0e57da3069ecf8f21acdf8b26364a2d1334ab766b8576fd34ba.
//
// Solidity: event OperatorSlashContractSet(address _oldOperatorSlashContract, address _operatorSlashContractAddress)
func (_Operator *OperatorFilterer) ParseOperatorSlashContractSet(log types.Log) (*OperatorOperatorSlashContractSet, error) {
	event := new(OperatorOperatorSlashContractSet)
	if err := _Operator.contract.UnpackLog(event, "OperatorSlashContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorOperatorVaultContractResetIterator is returned from FilterOperatorVaultContractReset and is used to iterate over the raw logs and unpacked data for OperatorVaultContractReset events raised by the Operator contract.
type OperatorOperatorVaultContractResetIterator struct {
	Event *OperatorOperatorVaultContractReset // Event containing the contract specifics and raw log

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
func (it *OperatorOperatorVaultContractResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorOperatorVaultContractReset)
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
		it.Event = new(OperatorOperatorVaultContractReset)
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
func (it *OperatorOperatorVaultContractResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorOperatorVaultContractResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorOperatorVaultContractReset represents a OperatorVaultContractReset event raised by the Operator contract.
type OperatorOperatorVaultContractReset struct {
	OldVaultContractAddress common.Address
	VaultContractAddress    common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterOperatorVaultContractReset is a free log retrieval operation binding the contract event 0xa44583d9d2c472d8e6428a1e28940263792e3d7ba27051c39527f33d5404dba4.
//
// Solidity: event OperatorVaultContractReset(address _oldVaultContractAddress, address _vaultContractAddress)
func (_Operator *OperatorFilterer) FilterOperatorVaultContractReset(opts *bind.FilterOpts) (*OperatorOperatorVaultContractResetIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "OperatorVaultContractReset")
	if err != nil {
		return nil, err
	}
	return &OperatorOperatorVaultContractResetIterator{contract: _Operator.contract, event: "OperatorVaultContractReset", logs: logs, sub: sub}, nil
}

// WatchOperatorVaultContractReset is a free log subscription operation binding the contract event 0xa44583d9d2c472d8e6428a1e28940263792e3d7ba27051c39527f33d5404dba4.
//
// Solidity: event OperatorVaultContractReset(address _oldVaultContractAddress, address _vaultContractAddress)
func (_Operator *OperatorFilterer) WatchOperatorVaultContractReset(opts *bind.WatchOpts, sink chan<- *OperatorOperatorVaultContractReset) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "OperatorVaultContractReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorOperatorVaultContractReset)
				if err := _Operator.contract.UnpackLog(event, "OperatorVaultContractReset", log); err != nil {
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

// ParseOperatorVaultContractReset is a log parse operation binding the contract event 0xa44583d9d2c472d8e6428a1e28940263792e3d7ba27051c39527f33d5404dba4.
//
// Solidity: event OperatorVaultContractReset(address _oldVaultContractAddress, address _vaultContractAddress)
func (_Operator *OperatorFilterer) ParseOperatorVaultContractReset(log types.Log) (*OperatorOperatorVaultContractReset, error) {
	event := new(OperatorOperatorVaultContractReset)
	if err := _Operator.contract.UnpackLog(event, "OperatorVaultContractReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorOperatorWithdrawIterator is returned from FilterOperatorWithdraw and is used to iterate over the raw logs and unpacked data for OperatorWithdraw events raised by the Operator contract.
type OperatorOperatorWithdrawIterator struct {
	Event *OperatorOperatorWithdraw // Event containing the contract specifics and raw log

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
func (it *OperatorOperatorWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorOperatorWithdraw)
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
		it.Event = new(OperatorOperatorWithdraw)
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
func (it *OperatorOperatorWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorOperatorWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorOperatorWithdraw represents a OperatorWithdraw event raised by the Operator contract.
type OperatorOperatorWithdraw struct {
	OperatorId     *big.Int
	WithdrawAmount *big.Int
	To             common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorWithdraw is a free log retrieval operation binding the contract event 0xaf6e54ef16ab2ee7c0fc473a41807f7cb0745891c6392827fb33943e840b472f.
//
// Solidity: event OperatorWithdraw(uint256 _operatorId, uint256 _withdrawAmount, address _to)
func (_Operator *OperatorFilterer) FilterOperatorWithdraw(opts *bind.FilterOpts) (*OperatorOperatorWithdrawIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "OperatorWithdraw")
	if err != nil {
		return nil, err
	}
	return &OperatorOperatorWithdrawIterator{contract: _Operator.contract, event: "OperatorWithdraw", logs: logs, sub: sub}, nil
}

// WatchOperatorWithdraw is a free log subscription operation binding the contract event 0xaf6e54ef16ab2ee7c0fc473a41807f7cb0745891c6392827fb33943e840b472f.
//
// Solidity: event OperatorWithdraw(uint256 _operatorId, uint256 _withdrawAmount, address _to)
func (_Operator *OperatorFilterer) WatchOperatorWithdraw(opts *bind.WatchOpts, sink chan<- *OperatorOperatorWithdraw) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "OperatorWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorOperatorWithdraw)
				if err := _Operator.contract.UnpackLog(event, "OperatorWithdraw", log); err != nil {
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

// ParseOperatorWithdraw is a log parse operation binding the contract event 0xaf6e54ef16ab2ee7c0fc473a41807f7cb0745891c6392827fb33943e840b472f.
//
// Solidity: event OperatorWithdraw(uint256 _operatorId, uint256 _withdrawAmount, address _to)
func (_Operator *OperatorFilterer) ParseOperatorWithdraw(log types.Log) (*OperatorOperatorWithdraw, error) {
	event := new(OperatorOperatorWithdraw)
	if err := _Operator.contract.UnpackLog(event, "OperatorWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Operator contract.
type OperatorOwnershipTransferredIterator struct {
	Event *OperatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OperatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorOwnershipTransferred)
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
		it.Event = new(OperatorOwnershipTransferred)
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
func (it *OperatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorOwnershipTransferred represents a OwnershipTransferred event raised by the Operator contract.
type OperatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Operator *OperatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OperatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Operator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OperatorOwnershipTransferredIterator{contract: _Operator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Operator *OperatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OperatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Operator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorOwnershipTransferred)
				if err := _Operator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Operator *OperatorFilterer) ParseOwnershipTransferred(log types.Log) (*OperatorOwnershipTransferred, error) {
	event := new(OperatorOwnershipTransferred)
	if err := _Operator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorPermissionlessBlockNumberSetIterator is returned from FilterPermissionlessBlockNumberSet and is used to iterate over the raw logs and unpacked data for PermissionlessBlockNumberSet events raised by the Operator contract.
type OperatorPermissionlessBlockNumberSetIterator struct {
	Event *OperatorPermissionlessBlockNumberSet // Event containing the contract specifics and raw log

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
func (it *OperatorPermissionlessBlockNumberSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorPermissionlessBlockNumberSet)
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
		it.Event = new(OperatorPermissionlessBlockNumberSet)
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
func (it *OperatorPermissionlessBlockNumberSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorPermissionlessBlockNumberSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorPermissionlessBlockNumberSet represents a PermissionlessBlockNumberSet event raised by the Operator contract.
type OperatorPermissionlessBlockNumberSet struct {
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPermissionlessBlockNumberSet is a free log retrieval operation binding the contract event 0xdc6af8e7f8757e67cef03eaa0fe837c22d44cc18b722f618091e9d5277cb8e6a.
//
// Solidity: event PermissionlessBlockNumberSet(uint256 _blockNumber)
func (_Operator *OperatorFilterer) FilterPermissionlessBlockNumberSet(opts *bind.FilterOpts) (*OperatorPermissionlessBlockNumberSetIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "PermissionlessBlockNumberSet")
	if err != nil {
		return nil, err
	}
	return &OperatorPermissionlessBlockNumberSetIterator{contract: _Operator.contract, event: "PermissionlessBlockNumberSet", logs: logs, sub: sub}, nil
}

// WatchPermissionlessBlockNumberSet is a free log subscription operation binding the contract event 0xdc6af8e7f8757e67cef03eaa0fe837c22d44cc18b722f618091e9d5277cb8e6a.
//
// Solidity: event PermissionlessBlockNumberSet(uint256 _blockNumber)
func (_Operator *OperatorFilterer) WatchPermissionlessBlockNumberSet(opts *bind.WatchOpts, sink chan<- *OperatorPermissionlessBlockNumberSet) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "PermissionlessBlockNumberSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorPermissionlessBlockNumberSet)
				if err := _Operator.contract.UnpackLog(event, "PermissionlessBlockNumberSet", log); err != nil {
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

// ParsePermissionlessBlockNumberSet is a log parse operation binding the contract event 0xdc6af8e7f8757e67cef03eaa0fe837c22d44cc18b722f618091e9d5277cb8e6a.
//
// Solidity: event PermissionlessBlockNumberSet(uint256 _blockNumber)
func (_Operator *OperatorFilterer) ParsePermissionlessBlockNumberSet(log types.Log) (*OperatorPermissionlessBlockNumberSet, error) {
	event := new(OperatorPermissionlessBlockNumberSet)
	if err := _Operator.contract.UnpackLog(event, "PermissionlessBlockNumberSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorPledgeDepositedIterator is returned from FilterPledgeDeposited and is used to iterate over the raw logs and unpacked data for PledgeDeposited events raised by the Operator contract.
type OperatorPledgeDepositedIterator struct {
	Event *OperatorPledgeDeposited // Event containing the contract specifics and raw log

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
func (it *OperatorPledgeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorPledgeDeposited)
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
		it.Event = new(OperatorPledgeDeposited)
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
func (it *OperatorPledgeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorPledgeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorPledgeDeposited represents a PledgeDeposited event raised by the Operator contract.
type OperatorPledgeDeposited struct {
	Amount     *big.Int
	OperatorId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPledgeDeposited is a free log retrieval operation binding the contract event 0x9b875040cd8c5e621238661f515f833c53c72ab073a8455b8d207c30e9f46c9a.
//
// Solidity: event PledgeDeposited(uint256 _amount, uint256 _operatorId)
func (_Operator *OperatorFilterer) FilterPledgeDeposited(opts *bind.FilterOpts) (*OperatorPledgeDepositedIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "PledgeDeposited")
	if err != nil {
		return nil, err
	}
	return &OperatorPledgeDepositedIterator{contract: _Operator.contract, event: "PledgeDeposited", logs: logs, sub: sub}, nil
}

// WatchPledgeDeposited is a free log subscription operation binding the contract event 0x9b875040cd8c5e621238661f515f833c53c72ab073a8455b8d207c30e9f46c9a.
//
// Solidity: event PledgeDeposited(uint256 _amount, uint256 _operatorId)
func (_Operator *OperatorFilterer) WatchPledgeDeposited(opts *bind.WatchOpts, sink chan<- *OperatorPledgeDeposited) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "PledgeDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorPledgeDeposited)
				if err := _Operator.contract.UnpackLog(event, "PledgeDeposited", log); err != nil {
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

// ParsePledgeDeposited is a log parse operation binding the contract event 0x9b875040cd8c5e621238661f515f833c53c72ab073a8455b8d207c30e9f46c9a.
//
// Solidity: event PledgeDeposited(uint256 _amount, uint256 _operatorId)
func (_Operator *OperatorFilterer) ParsePledgeDeposited(log types.Log) (*OperatorPledgeDeposited, error) {
	event := new(OperatorPledgeDeposited)
	if err := _Operator.contract.UnpackLog(event, "PledgeDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRegistrationFeeChangedIterator is returned from FilterRegistrationFeeChanged and is used to iterate over the raw logs and unpacked data for RegistrationFeeChanged events raised by the Operator contract.
type OperatorRegistrationFeeChangedIterator struct {
	Event *OperatorRegistrationFeeChanged // Event containing the contract specifics and raw log

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
func (it *OperatorRegistrationFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistrationFeeChanged)
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
		it.Event = new(OperatorRegistrationFeeChanged)
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
func (it *OperatorRegistrationFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistrationFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistrationFeeChanged represents a RegistrationFeeChanged event raised by the Operator contract.
type OperatorRegistrationFeeChanged struct {
	OldFee *big.Int
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRegistrationFeeChanged is a free log retrieval operation binding the contract event 0x72071ceb13e0088702b0aa4b4ca2419810b96f85b26a8d059a60d289c59ea7a9.
//
// Solidity: event RegistrationFeeChanged(uint256 _oldFee, uint256 _fee)
func (_Operator *OperatorFilterer) FilterRegistrationFeeChanged(opts *bind.FilterOpts) (*OperatorRegistrationFeeChangedIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "RegistrationFeeChanged")
	if err != nil {
		return nil, err
	}
	return &OperatorRegistrationFeeChangedIterator{contract: _Operator.contract, event: "RegistrationFeeChanged", logs: logs, sub: sub}, nil
}

// WatchRegistrationFeeChanged is a free log subscription operation binding the contract event 0x72071ceb13e0088702b0aa4b4ca2419810b96f85b26a8d059a60d289c59ea7a9.
//
// Solidity: event RegistrationFeeChanged(uint256 _oldFee, uint256 _fee)
func (_Operator *OperatorFilterer) WatchRegistrationFeeChanged(opts *bind.WatchOpts, sink chan<- *OperatorRegistrationFeeChanged) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "RegistrationFeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistrationFeeChanged)
				if err := _Operator.contract.UnpackLog(event, "RegistrationFeeChanged", log); err != nil {
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

// ParseRegistrationFeeChanged is a log parse operation binding the contract event 0x72071ceb13e0088702b0aa4b4ca2419810b96f85b26a8d059a60d289c59ea7a9.
//
// Solidity: event RegistrationFeeChanged(uint256 _oldFee, uint256 _fee)
func (_Operator *OperatorFilterer) ParseRegistrationFeeChanged(log types.Log) (*OperatorRegistrationFeeChanged, error) {
	event := new(OperatorRegistrationFeeChanged)
	if err := _Operator.contract.UnpackLog(event, "RegistrationFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the Operator contract.
type OperatorSlashedIterator struct {
	Event *OperatorSlashed // Event containing the contract specifics and raw log

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
func (it *OperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSlashed)
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
		it.Event = new(OperatorSlashed)
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
func (it *OperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSlashed represents a Slashed event raised by the Operator contract.
type OperatorSlashed struct {
	OperatorId *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x4f5f38ee30b01a960b4dfdcd520a3ca59c1a664a32dcfe5418ca79b0de6b7236.
//
// Solidity: event Slashed(uint256 _operatorId, uint256 _amount)
func (_Operator *OperatorFilterer) FilterSlashed(opts *bind.FilterOpts) (*OperatorSlashedIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "Slashed")
	if err != nil {
		return nil, err
	}
	return &OperatorSlashedIterator{contract: _Operator.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x4f5f38ee30b01a960b4dfdcd520a3ca59c1a664a32dcfe5418ca79b0de6b7236.
//
// Solidity: event Slashed(uint256 _operatorId, uint256 _amount)
func (_Operator *OperatorFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *OperatorSlashed) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "Slashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSlashed)
				if err := _Operator.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0x4f5f38ee30b01a960b4dfdcd520a3ca59c1a664a32dcfe5418ca79b0de6b7236.
//
// Solidity: event Slashed(uint256 _operatorId, uint256 _amount)
func (_Operator *OperatorFilterer) ParseSlashed(log types.Log) (*OperatorSlashed, error) {
	event := new(OperatorSlashed)
	if err := _Operator.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorTransferredIterator is returned from FilterTransferred and is used to iterate over the raw logs and unpacked data for Transferred events raised by the Operator contract.
type OperatorTransferredIterator struct {
	Event *OperatorTransferred // Event containing the contract specifics and raw log

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
func (it *OperatorTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorTransferred)
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
		it.Event = new(OperatorTransferred)
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
func (it *OperatorTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorTransferred represents a Transferred event raised by the Operator contract.
type OperatorTransferred struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferred is a free log retrieval operation binding the contract event 0xe6d858f14d755446648a6e0c8ab8b5a0f58ccc7920d4c910b0454e4dcd869af0.
//
// Solidity: event Transferred(address _to, uint256 _amount)
func (_Operator *OperatorFilterer) FilterTransferred(opts *bind.FilterOpts) (*OperatorTransferredIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "Transferred")
	if err != nil {
		return nil, err
	}
	return &OperatorTransferredIterator{contract: _Operator.contract, event: "Transferred", logs: logs, sub: sub}, nil
}

// WatchTransferred is a free log subscription operation binding the contract event 0xe6d858f14d755446648a6e0c8ab8b5a0f58ccc7920d4c910b0454e4dcd869af0.
//
// Solidity: event Transferred(address _to, uint256 _amount)
func (_Operator *OperatorFilterer) WatchTransferred(opts *bind.WatchOpts, sink chan<- *OperatorTransferred) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "Transferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorTransferred)
				if err := _Operator.contract.UnpackLog(event, "Transferred", log); err != nil {
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

// ParseTransferred is a log parse operation binding the contract event 0xe6d858f14d755446648a6e0c8ab8b5a0f58ccc7920d4c910b0454e4dcd869af0.
//
// Solidity: event Transferred(address _to, uint256 _amount)
func (_Operator *OperatorFilterer) ParseTransferred(log types.Log) (*OperatorTransferred, error) {
	event := new(OperatorTransferred)
	if err := _Operator.contract.UnpackLog(event, "Transferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Operator contract.
type OperatorUpgradedIterator struct {
	Event *OperatorUpgraded // Event containing the contract specifics and raw log

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
func (it *OperatorUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorUpgraded)
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
		it.Event = new(OperatorUpgraded)
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
func (it *OperatorUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorUpgraded represents a Upgraded event raised by the Operator contract.
type OperatorUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Operator *OperatorFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*OperatorUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Operator.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &OperatorUpgradedIterator{contract: _Operator.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Operator *OperatorFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *OperatorUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Operator.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorUpgraded)
				if err := _Operator.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Operator *OperatorFilterer) ParseUpgraded(log types.Log) (*OperatorUpgraded, error) {
	event := new(OperatorUpgraded)
	if err := _Operator.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorVaultFactorContractSetIterator is returned from FilterVaultFactorContractSet and is used to iterate over the raw logs and unpacked data for VaultFactorContractSet events raised by the Operator contract.
type OperatorVaultFactorContractSetIterator struct {
	Event *OperatorVaultFactorContractSet // Event containing the contract specifics and raw log

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
func (it *OperatorVaultFactorContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorVaultFactorContractSet)
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
		it.Event = new(OperatorVaultFactorContractSet)
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
func (it *OperatorVaultFactorContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorVaultFactorContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorVaultFactorContractSet represents a VaultFactorContractSet event raised by the Operator contract.
type OperatorVaultFactorContractSet struct {
	VaultFactoryContract        common.Address
	VaultFactoryContractAddress common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterVaultFactorContractSet is a free log retrieval operation binding the contract event 0x72766c8a01391f3f73d521b1de59f4a510ad8550ee3130a76a5729d445109f1b.
//
// Solidity: event VaultFactorContractSet(address _vaultFactoryContract, address _vaultFactoryContractAddress)
func (_Operator *OperatorFilterer) FilterVaultFactorContractSet(opts *bind.FilterOpts) (*OperatorVaultFactorContractSetIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "VaultFactorContractSet")
	if err != nil {
		return nil, err
	}
	return &OperatorVaultFactorContractSetIterator{contract: _Operator.contract, event: "VaultFactorContractSet", logs: logs, sub: sub}, nil
}

// WatchVaultFactorContractSet is a free log subscription operation binding the contract event 0x72766c8a01391f3f73d521b1de59f4a510ad8550ee3130a76a5729d445109f1b.
//
// Solidity: event VaultFactorContractSet(address _vaultFactoryContract, address _vaultFactoryContractAddress)
func (_Operator *OperatorFilterer) WatchVaultFactorContractSet(opts *bind.WatchOpts, sink chan<- *OperatorVaultFactorContractSet) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "VaultFactorContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorVaultFactorContractSet)
				if err := _Operator.contract.UnpackLog(event, "VaultFactorContractSet", log); err != nil {
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

// ParseVaultFactorContractSet is a log parse operation binding the contract event 0x72766c8a01391f3f73d521b1de59f4a510ad8550ee3130a76a5729d445109f1b.
//
// Solidity: event VaultFactorContractSet(address _vaultFactoryContract, address _vaultFactoryContractAddress)
func (_Operator *OperatorFilterer) ParseVaultFactorContractSet(log types.Log) (*OperatorVaultFactorContractSet, error) {
	event := new(OperatorVaultFactorContractSet)
	if err := _Operator.contract.UnpackLog(event, "VaultFactorContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Operator contract.
type OperatorWithdrawIterator struct {
	Event *OperatorWithdraw // Event containing the contract specifics and raw log

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
func (it *OperatorWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorWithdraw)
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
		it.Event = new(OperatorWithdraw)
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
func (it *OperatorWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorWithdraw represents a Withdraw event raised by the Operator contract.
type OperatorWithdraw struct {
	Amount     *big.Int
	OperatorId *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x71ef96c43343734b1d843bb85d52ef329f5e9143e4d35827771e3b0dd90c5f84.
//
// Solidity: event Withdraw(uint256 _amount, uint256 _operatorId, address _to)
func (_Operator *OperatorFilterer) FilterWithdraw(opts *bind.FilterOpts) (*OperatorWithdrawIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &OperatorWithdrawIterator{contract: _Operator.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x71ef96c43343734b1d843bb85d52ef329f5e9143e4d35827771e3b0dd90c5f84.
//
// Solidity: event Withdraw(uint256 _amount, uint256 _operatorId, address _to)
func (_Operator *OperatorFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *OperatorWithdraw) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorWithdraw)
				if err := _Operator.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x71ef96c43343734b1d843bb85d52ef329f5e9143e4d35827771e3b0dd90c5f84.
//
// Solidity: event Withdraw(uint256 _amount, uint256 _operatorId, address _to)
func (_Operator *OperatorFilterer) ParseWithdraw(log types.Log) (*OperatorWithdraw, error) {
	event := new(OperatorWithdraw)
	if err := _Operator.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

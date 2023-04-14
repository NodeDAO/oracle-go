// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package liq

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

// LiquidStakingStakeInfo is an auto generated low-level Go binding around an user-defined struct.
type LiquidStakingStakeInfo struct {
	OperatorId *big.Int
	Quota      *big.Int
}

// LiqMetaData contains all meta data concerning the Liq contract.
var LiqMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AssignMustSameOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientMargin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDaoVaultAddr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawalCredentials\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorHasArrears\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorLoanFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequireBlacklistOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequireOperatorTrusted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TotalEthIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnstakeEthNoQuota\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldBeaconOracleContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_beaconOracleContractAddress\",\"type\":\"address\"}],\"name\":\"BeaconOracleContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vaultManagerContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_consensusVaultContract\",\"type\":\"address\"}],\"name\":\"ConsensusVaultContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldDao\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"DaoAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"DaoClaimRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldDaoVaultAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_daoVaultAddress\",\"type\":\"address\"}],\"name\":\"DaoVaultAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_feeRate\",\"type\":\"uint256\"}],\"name\":\"DepositFeeRateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"}],\"name\":\"EthStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amounts\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"EthUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_oldLiquidStakingWithdrawalCredentials\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_liquidStakingWithdrawalCredentials\",\"type\":\"bytes\"}],\"name\":\"LiquidStakingWithdrawalCredentialsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"exitBlockNumbers\",\"type\":\"uint256[]\"}],\"name\":\"NftExitBlockNumberSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"NftStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldNodeOperatorRegistryContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistryContract\",\"type\":\"address\"}],\"name\":\"NodeOperatorRegistryContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_blacklistOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"}],\"name\":\"OperatorAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operatorCanLoanAmounts\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newCanloadAmounts\",\"type\":\"uint256\"}],\"name\":\"OperatorCanLoanAmountsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"OperatorClaimRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"OperatorReinvestClRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"OperatorReinvestElRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"RewardsReceive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"UserClaimRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ValidatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vaultManagerContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_vaultManagerContract\",\"type\":\"address\"}],\"name\":\"VaultManagerContractSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addSlashFundToStakePool\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"assignOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beaconOracleContract\",\"outputs\":[{\"internalType\":\"contractIBeaconOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_operatorIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_rewards\",\"type\":\"uint256[]\"}],\"name\":\"claimRewardsOfDao\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_rewardAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_rewards\",\"type\":\"uint256[]\"}],\"name\":\"claimRewardsOfOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_totalNftRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasHeight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"claimRewardsOfUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"consensusVaultContract\",\"outputs\":[{\"internalType\":\"contractIConsensusVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositContract\",\"outputs\":[{\"internalType\":\"contractIDepositContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"fastUnstakeNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nethAmountIn\",\"type\":\"uint256\"}],\"name\":\"getEthOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethAmountIn\",\"type\":\"uint256\"}],\"name\":\"getNethOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getOperatorNethUnstakePoolAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalEthValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"getUnstakeQuota\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quota\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidStaking.StakeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_daoVaultAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_withdrawalCreds\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistryContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nETHContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nVNFTContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beaconOracleContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositContractAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_operatorIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_users\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_nethAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_consensusVaultContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultManagerContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_withdrawalRequestContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operatorSlashContractAddress\",\"type\":\"address\"}],\"name\":\"initializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalRequestNethAmount\",\"type\":\"uint256\"}],\"name\":\"largeWithdrawalBurnNeth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"largeWithdrawalUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidStakingWithdrawalCredentials\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nETHContract\",\"outputs\":[{\"internalType\":\"contractINETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_exitBlockNumbers\",\"type\":\"uint256[]\"}],\"name\":\"nftExitHandle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeOperatorRegistryContract\",\"outputs\":[{\"internalType\":\"contractINodeOperatorsRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorCanLoanAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorLoadBlockNumbers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorLoanRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorNftPoolBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorPoolBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorSlashContract\",\"outputs\":[{\"internalType\":\"contractIOperatorSlash\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reAssignRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"receiveRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_operatorIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"reinvestClRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_operatorIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"reinvestElRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beaconOracleContractAddress\",\"type\":\"address\"}],\"name\":\"setBeaconOracleContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"setDaoAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_daoVaultAddress\",\"type\":\"address\"}],\"name\":\"setDaoVaultAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeRate\",\"type\":\"uint256\"}],\"name\":\"setDepositFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_liquidStakingWithdrawalCredentials\",\"type\":\"bytes\"}],\"name\":\"setLiquidStakingWithdrawalCredentials\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistryContract\",\"type\":\"address\"}],\"name\":\"setNodeOperatorRegistryContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newCanloadAmounts\",\"type\":\"uint256\"}],\"name\":\"setOperatorCanLoanAmounts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultManagerContract\",\"type\":\"address\"}],\"name\":\"setVaultManagerContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"stakeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"withdrawalCredentialsAddress\",\"type\":\"address\"}],\"name\":\"stakeNFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amounts\",\"type\":\"uint256\"}],\"name\":\"unstakeETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vNFTContract\",\"outputs\":[{\"internalType\":\"contractIVNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultManagerContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalRequestContract\",\"outputs\":[{\"internalType\":\"contractIWithdrawalRequest\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LiqABI is the input ABI used to generate the binding from.
// Deprecated: Use LiqMetaData.ABI instead.
var LiqABI = LiqMetaData.ABI

// Liq is an auto generated Go binding around an Ethereum contract.
type Liq struct {
	LiqCaller     // Read-only binding to the contract
	LiqTransactor // Write-only binding to the contract
	LiqFilterer   // Log filterer for contract events
}

// LiqCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiqCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiqTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiqTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiqFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiqFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiqSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiqSession struct {
	Contract     *Liq              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LiqCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiqCallerSession struct {
	Contract *LiqCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LiqTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiqTransactorSession struct {
	Contract     *LiqTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LiqRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiqRaw struct {
	Contract *Liq // Generic contract binding to access the raw methods on
}

// LiqCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiqCallerRaw struct {
	Contract *LiqCaller // Generic read-only contract binding to access the raw methods on
}

// LiqTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiqTransactorRaw struct {
	Contract *LiqTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiq creates a new instance of Liq, bound to a specific deployed contract.
func NewLiq(address common.Address, backend bind.ContractBackend) (*Liq, error) {
	contract, err := bindLiq(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Liq{LiqCaller: LiqCaller{contract: contract}, LiqTransactor: LiqTransactor{contract: contract}, LiqFilterer: LiqFilterer{contract: contract}}, nil
}

// NewLiqCaller creates a new read-only instance of Liq, bound to a specific deployed contract.
func NewLiqCaller(address common.Address, caller bind.ContractCaller) (*LiqCaller, error) {
	contract, err := bindLiq(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiqCaller{contract: contract}, nil
}

// NewLiqTransactor creates a new write-only instance of Liq, bound to a specific deployed contract.
func NewLiqTransactor(address common.Address, transactor bind.ContractTransactor) (*LiqTransactor, error) {
	contract, err := bindLiq(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiqTransactor{contract: contract}, nil
}

// NewLiqFilterer creates a new log filterer instance of Liq, bound to a specific deployed contract.
func NewLiqFilterer(address common.Address, filterer bind.ContractFilterer) (*LiqFilterer, error) {
	contract, err := bindLiq(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiqFilterer{contract: contract}, nil
}

// bindLiq binds a generic wrapper to an already deployed contract.
func bindLiq(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LiqMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Liq *LiqRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Liq.Contract.LiqCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Liq *LiqRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liq.Contract.LiqTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Liq *LiqRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Liq.Contract.LiqTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Liq *LiqCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Liq.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Liq *LiqTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liq.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Liq *LiqTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Liq.Contract.contract.Transact(opts, method, params...)
}

// BeaconOracleContract is a free data retrieval call binding the contract method 0xe383edb4.
//
// Solidity: function beaconOracleContract() view returns(address)
func (_Liq *LiqCaller) BeaconOracleContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "beaconOracleContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconOracleContract is a free data retrieval call binding the contract method 0xe383edb4.
//
// Solidity: function beaconOracleContract() view returns(address)
func (_Liq *LiqSession) BeaconOracleContract() (common.Address, error) {
	return _Liq.Contract.BeaconOracleContract(&_Liq.CallOpts)
}

// BeaconOracleContract is a free data retrieval call binding the contract method 0xe383edb4.
//
// Solidity: function beaconOracleContract() view returns(address)
func (_Liq *LiqCallerSession) BeaconOracleContract() (common.Address, error) {
	return _Liq.Contract.BeaconOracleContract(&_Liq.CallOpts)
}

// ConsensusVaultContract is a free data retrieval call binding the contract method 0xd6a29dc5.
//
// Solidity: function consensusVaultContract() view returns(address)
func (_Liq *LiqCaller) ConsensusVaultContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "consensusVaultContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConsensusVaultContract is a free data retrieval call binding the contract method 0xd6a29dc5.
//
// Solidity: function consensusVaultContract() view returns(address)
func (_Liq *LiqSession) ConsensusVaultContract() (common.Address, error) {
	return _Liq.Contract.ConsensusVaultContract(&_Liq.CallOpts)
}

// ConsensusVaultContract is a free data retrieval call binding the contract method 0xd6a29dc5.
//
// Solidity: function consensusVaultContract() view returns(address)
func (_Liq *LiqCallerSession) ConsensusVaultContract() (common.Address, error) {
	return _Liq.Contract.ConsensusVaultContract(&_Liq.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Liq *LiqCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Liq *LiqSession) Dao() (common.Address, error) {
	return _Liq.Contract.Dao(&_Liq.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Liq *LiqCallerSession) Dao() (common.Address, error) {
	return _Liq.Contract.Dao(&_Liq.CallOpts)
}

// DaoVaultAddress is a free data retrieval call binding the contract method 0x3d6a3844.
//
// Solidity: function daoVaultAddress() view returns(address)
func (_Liq *LiqCaller) DaoVaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "daoVaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaoVaultAddress is a free data retrieval call binding the contract method 0x3d6a3844.
//
// Solidity: function daoVaultAddress() view returns(address)
func (_Liq *LiqSession) DaoVaultAddress() (common.Address, error) {
	return _Liq.Contract.DaoVaultAddress(&_Liq.CallOpts)
}

// DaoVaultAddress is a free data retrieval call binding the contract method 0x3d6a3844.
//
// Solidity: function daoVaultAddress() view returns(address)
func (_Liq *LiqCallerSession) DaoVaultAddress() (common.Address, error) {
	return _Liq.Contract.DaoVaultAddress(&_Liq.CallOpts)
}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_Liq *LiqCaller) DepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "depositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_Liq *LiqSession) DepositContract() (common.Address, error) {
	return _Liq.Contract.DepositContract(&_Liq.CallOpts)
}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_Liq *LiqCallerSession) DepositContract() (common.Address, error) {
	return _Liq.Contract.DepositContract(&_Liq.CallOpts)
}

// DepositFeeRate is a free data retrieval call binding the contract method 0x852cb9b8.
//
// Solidity: function depositFeeRate() view returns(uint256)
func (_Liq *LiqCaller) DepositFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "depositFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositFeeRate is a free data retrieval call binding the contract method 0x852cb9b8.
//
// Solidity: function depositFeeRate() view returns(uint256)
func (_Liq *LiqSession) DepositFeeRate() (*big.Int, error) {
	return _Liq.Contract.DepositFeeRate(&_Liq.CallOpts)
}

// DepositFeeRate is a free data retrieval call binding the contract method 0x852cb9b8.
//
// Solidity: function depositFeeRate() view returns(uint256)
func (_Liq *LiqCallerSession) DepositFeeRate() (*big.Int, error) {
	return _Liq.Contract.DepositFeeRate(&_Liq.CallOpts)
}

// GetEthOut is a free data retrieval call binding the contract method 0xb2ff2410.
//
// Solidity: function getEthOut(uint256 _nethAmountIn) view returns(uint256)
func (_Liq *LiqCaller) GetEthOut(opts *bind.CallOpts, _nethAmountIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "getEthOut", _nethAmountIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthOut is a free data retrieval call binding the contract method 0xb2ff2410.
//
// Solidity: function getEthOut(uint256 _nethAmountIn) view returns(uint256)
func (_Liq *LiqSession) GetEthOut(_nethAmountIn *big.Int) (*big.Int, error) {
	return _Liq.Contract.GetEthOut(&_Liq.CallOpts, _nethAmountIn)
}

// GetEthOut is a free data retrieval call binding the contract method 0xb2ff2410.
//
// Solidity: function getEthOut(uint256 _nethAmountIn) view returns(uint256)
func (_Liq *LiqCallerSession) GetEthOut(_nethAmountIn *big.Int) (*big.Int, error) {
	return _Liq.Contract.GetEthOut(&_Liq.CallOpts, _nethAmountIn)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns(uint256)
func (_Liq *LiqCaller) GetExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "getExchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns(uint256)
func (_Liq *LiqSession) GetExchangeRate() (*big.Int, error) {
	return _Liq.Contract.GetExchangeRate(&_Liq.CallOpts)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns(uint256)
func (_Liq *LiqCallerSession) GetExchangeRate() (*big.Int, error) {
	return _Liq.Contract.GetExchangeRate(&_Liq.CallOpts)
}

// GetNethOut is a free data retrieval call binding the contract method 0x9ec2d598.
//
// Solidity: function getNethOut(uint256 _ethAmountIn) view returns(uint256)
func (_Liq *LiqCaller) GetNethOut(opts *bind.CallOpts, _ethAmountIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "getNethOut", _ethAmountIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNethOut is a free data retrieval call binding the contract method 0x9ec2d598.
//
// Solidity: function getNethOut(uint256 _ethAmountIn) view returns(uint256)
func (_Liq *LiqSession) GetNethOut(_ethAmountIn *big.Int) (*big.Int, error) {
	return _Liq.Contract.GetNethOut(&_Liq.CallOpts, _ethAmountIn)
}

// GetNethOut is a free data retrieval call binding the contract method 0x9ec2d598.
//
// Solidity: function getNethOut(uint256 _ethAmountIn) view returns(uint256)
func (_Liq *LiqCallerSession) GetNethOut(_ethAmountIn *big.Int) (*big.Int, error) {
	return _Liq.Contract.GetNethOut(&_Liq.CallOpts, _ethAmountIn)
}

// GetOperatorNethUnstakePoolAmounts is a free data retrieval call binding the contract method 0xeef02d3b.
//
// Solidity: function getOperatorNethUnstakePoolAmounts(uint256 _operatorId) view returns(uint256)
func (_Liq *LiqCaller) GetOperatorNethUnstakePoolAmounts(opts *bind.CallOpts, _operatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "getOperatorNethUnstakePoolAmounts", _operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorNethUnstakePoolAmounts is a free data retrieval call binding the contract method 0xeef02d3b.
//
// Solidity: function getOperatorNethUnstakePoolAmounts(uint256 _operatorId) view returns(uint256)
func (_Liq *LiqSession) GetOperatorNethUnstakePoolAmounts(_operatorId *big.Int) (*big.Int, error) {
	return _Liq.Contract.GetOperatorNethUnstakePoolAmounts(&_Liq.CallOpts, _operatorId)
}

// GetOperatorNethUnstakePoolAmounts is a free data retrieval call binding the contract method 0xeef02d3b.
//
// Solidity: function getOperatorNethUnstakePoolAmounts(uint256 _operatorId) view returns(uint256)
func (_Liq *LiqCallerSession) GetOperatorNethUnstakePoolAmounts(_operatorId *big.Int) (*big.Int, error) {
	return _Liq.Contract.GetOperatorNethUnstakePoolAmounts(&_Liq.CallOpts, _operatorId)
}

// GetTotalEthValue is a free data retrieval call binding the contract method 0x4277f693.
//
// Solidity: function getTotalEthValue() view returns(uint256)
func (_Liq *LiqCaller) GetTotalEthValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "getTotalEthValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalEthValue is a free data retrieval call binding the contract method 0x4277f693.
//
// Solidity: function getTotalEthValue() view returns(uint256)
func (_Liq *LiqSession) GetTotalEthValue() (*big.Int, error) {
	return _Liq.Contract.GetTotalEthValue(&_Liq.CallOpts)
}

// GetTotalEthValue is a free data retrieval call binding the contract method 0x4277f693.
//
// Solidity: function getTotalEthValue() view returns(uint256)
func (_Liq *LiqCallerSession) GetTotalEthValue() (*big.Int, error) {
	return _Liq.Contract.GetTotalEthValue(&_Liq.CallOpts)
}

// GetUnstakeQuota is a free data retrieval call binding the contract method 0xbc7b0d20.
//
// Solidity: function getUnstakeQuota(address _from) view returns((uint256,uint256)[])
func (_Liq *LiqCaller) GetUnstakeQuota(opts *bind.CallOpts, _from common.Address) ([]LiquidStakingStakeInfo, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "getUnstakeQuota", _from)

	if err != nil {
		return *new([]LiquidStakingStakeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]LiquidStakingStakeInfo)).(*[]LiquidStakingStakeInfo)

	return out0, err

}

// GetUnstakeQuota is a free data retrieval call binding the contract method 0xbc7b0d20.
//
// Solidity: function getUnstakeQuota(address _from) view returns((uint256,uint256)[])
func (_Liq *LiqSession) GetUnstakeQuota(_from common.Address) ([]LiquidStakingStakeInfo, error) {
	return _Liq.Contract.GetUnstakeQuota(&_Liq.CallOpts, _from)
}

// GetUnstakeQuota is a free data retrieval call binding the contract method 0xbc7b0d20.
//
// Solidity: function getUnstakeQuota(address _from) view returns((uint256,uint256)[])
func (_Liq *LiqCallerSession) GetUnstakeQuota(_from common.Address) ([]LiquidStakingStakeInfo, error) {
	return _Liq.Contract.GetUnstakeQuota(&_Liq.CallOpts, _from)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Liq *LiqCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Liq *LiqSession) IsPaused() (bool, error) {
	return _Liq.Contract.IsPaused(&_Liq.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Liq *LiqCallerSession) IsPaused() (bool, error) {
	return _Liq.Contract.IsPaused(&_Liq.CallOpts)
}

// LiquidStakingWithdrawalCredentials is a free data retrieval call binding the contract method 0x12ebdab4.
//
// Solidity: function liquidStakingWithdrawalCredentials() view returns(bytes)
func (_Liq *LiqCaller) LiquidStakingWithdrawalCredentials(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "liquidStakingWithdrawalCredentials")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// LiquidStakingWithdrawalCredentials is a free data retrieval call binding the contract method 0x12ebdab4.
//
// Solidity: function liquidStakingWithdrawalCredentials() view returns(bytes)
func (_Liq *LiqSession) LiquidStakingWithdrawalCredentials() ([]byte, error) {
	return _Liq.Contract.LiquidStakingWithdrawalCredentials(&_Liq.CallOpts)
}

// LiquidStakingWithdrawalCredentials is a free data retrieval call binding the contract method 0x12ebdab4.
//
// Solidity: function liquidStakingWithdrawalCredentials() view returns(bytes)
func (_Liq *LiqCallerSession) LiquidStakingWithdrawalCredentials() ([]byte, error) {
	return _Liq.Contract.LiquidStakingWithdrawalCredentials(&_Liq.CallOpts)
}

// NETHContract is a free data retrieval call binding the contract method 0xbd97c759.
//
// Solidity: function nETHContract() view returns(address)
func (_Liq *LiqCaller) NETHContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "nETHContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NETHContract is a free data retrieval call binding the contract method 0xbd97c759.
//
// Solidity: function nETHContract() view returns(address)
func (_Liq *LiqSession) NETHContract() (common.Address, error) {
	return _Liq.Contract.NETHContract(&_Liq.CallOpts)
}

// NETHContract is a free data retrieval call binding the contract method 0xbd97c759.
//
// Solidity: function nETHContract() view returns(address)
func (_Liq *LiqCallerSession) NETHContract() (common.Address, error) {
	return _Liq.Contract.NETHContract(&_Liq.CallOpts)
}

// NodeOperatorRegistryContract is a free data retrieval call binding the contract method 0x9a42e0ba.
//
// Solidity: function nodeOperatorRegistryContract() view returns(address)
func (_Liq *LiqCaller) NodeOperatorRegistryContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "nodeOperatorRegistryContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeOperatorRegistryContract is a free data retrieval call binding the contract method 0x9a42e0ba.
//
// Solidity: function nodeOperatorRegistryContract() view returns(address)
func (_Liq *LiqSession) NodeOperatorRegistryContract() (common.Address, error) {
	return _Liq.Contract.NodeOperatorRegistryContract(&_Liq.CallOpts)
}

// NodeOperatorRegistryContract is a free data retrieval call binding the contract method 0x9a42e0ba.
//
// Solidity: function nodeOperatorRegistryContract() view returns(address)
func (_Liq *LiqCallerSession) NodeOperatorRegistryContract() (common.Address, error) {
	return _Liq.Contract.NodeOperatorRegistryContract(&_Liq.CallOpts)
}

// OperatorCanLoanAmounts is a free data retrieval call binding the contract method 0xd4a47d6d.
//
// Solidity: function operatorCanLoanAmounts() view returns(uint256)
func (_Liq *LiqCaller) OperatorCanLoanAmounts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "operatorCanLoanAmounts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorCanLoanAmounts is a free data retrieval call binding the contract method 0xd4a47d6d.
//
// Solidity: function operatorCanLoanAmounts() view returns(uint256)
func (_Liq *LiqSession) OperatorCanLoanAmounts() (*big.Int, error) {
	return _Liq.Contract.OperatorCanLoanAmounts(&_Liq.CallOpts)
}

// OperatorCanLoanAmounts is a free data retrieval call binding the contract method 0xd4a47d6d.
//
// Solidity: function operatorCanLoanAmounts() view returns(uint256)
func (_Liq *LiqCallerSession) OperatorCanLoanAmounts() (*big.Int, error) {
	return _Liq.Contract.OperatorCanLoanAmounts(&_Liq.CallOpts)
}

// OperatorLoadBlockNumbers is a free data retrieval call binding the contract method 0x2f7486f1.
//
// Solidity: function operatorLoadBlockNumbers(uint256 ) view returns(uint256)
func (_Liq *LiqCaller) OperatorLoadBlockNumbers(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "operatorLoadBlockNumbers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorLoadBlockNumbers is a free data retrieval call binding the contract method 0x2f7486f1.
//
// Solidity: function operatorLoadBlockNumbers(uint256 ) view returns(uint256)
func (_Liq *LiqSession) OperatorLoadBlockNumbers(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.OperatorLoadBlockNumbers(&_Liq.CallOpts, arg0)
}

// OperatorLoadBlockNumbers is a free data retrieval call binding the contract method 0x2f7486f1.
//
// Solidity: function operatorLoadBlockNumbers(uint256 ) view returns(uint256)
func (_Liq *LiqCallerSession) OperatorLoadBlockNumbers(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.OperatorLoadBlockNumbers(&_Liq.CallOpts, arg0)
}

// OperatorLoanRecords is a free data retrieval call binding the contract method 0x3eca7213.
//
// Solidity: function operatorLoanRecords(uint256 ) view returns(uint256)
func (_Liq *LiqCaller) OperatorLoanRecords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "operatorLoanRecords", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorLoanRecords is a free data retrieval call binding the contract method 0x3eca7213.
//
// Solidity: function operatorLoanRecords(uint256 ) view returns(uint256)
func (_Liq *LiqSession) OperatorLoanRecords(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.OperatorLoanRecords(&_Liq.CallOpts, arg0)
}

// OperatorLoanRecords is a free data retrieval call binding the contract method 0x3eca7213.
//
// Solidity: function operatorLoanRecords(uint256 ) view returns(uint256)
func (_Liq *LiqCallerSession) OperatorLoanRecords(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.OperatorLoanRecords(&_Liq.CallOpts, arg0)
}

// OperatorNftPoolBalances is a free data retrieval call binding the contract method 0x39007890.
//
// Solidity: function operatorNftPoolBalances(uint256 ) view returns(uint256)
func (_Liq *LiqCaller) OperatorNftPoolBalances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "operatorNftPoolBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorNftPoolBalances is a free data retrieval call binding the contract method 0x39007890.
//
// Solidity: function operatorNftPoolBalances(uint256 ) view returns(uint256)
func (_Liq *LiqSession) OperatorNftPoolBalances(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.OperatorNftPoolBalances(&_Liq.CallOpts, arg0)
}

// OperatorNftPoolBalances is a free data retrieval call binding the contract method 0x39007890.
//
// Solidity: function operatorNftPoolBalances(uint256 ) view returns(uint256)
func (_Liq *LiqCallerSession) OperatorNftPoolBalances(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.OperatorNftPoolBalances(&_Liq.CallOpts, arg0)
}

// OperatorPoolBalances is a free data retrieval call binding the contract method 0x74ee7f86.
//
// Solidity: function operatorPoolBalances(uint256 ) view returns(uint256)
func (_Liq *LiqCaller) OperatorPoolBalances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "operatorPoolBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorPoolBalances is a free data retrieval call binding the contract method 0x74ee7f86.
//
// Solidity: function operatorPoolBalances(uint256 ) view returns(uint256)
func (_Liq *LiqSession) OperatorPoolBalances(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.OperatorPoolBalances(&_Liq.CallOpts, arg0)
}

// OperatorPoolBalances is a free data retrieval call binding the contract method 0x74ee7f86.
//
// Solidity: function operatorPoolBalances(uint256 ) view returns(uint256)
func (_Liq *LiqCallerSession) OperatorPoolBalances(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.OperatorPoolBalances(&_Liq.CallOpts, arg0)
}

// OperatorSlashContract is a free data retrieval call binding the contract method 0x0c2559fd.
//
// Solidity: function operatorSlashContract() view returns(address)
func (_Liq *LiqCaller) OperatorSlashContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "operatorSlashContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorSlashContract is a free data retrieval call binding the contract method 0x0c2559fd.
//
// Solidity: function operatorSlashContract() view returns(address)
func (_Liq *LiqSession) OperatorSlashContract() (common.Address, error) {
	return _Liq.Contract.OperatorSlashContract(&_Liq.CallOpts)
}

// OperatorSlashContract is a free data retrieval call binding the contract method 0x0c2559fd.
//
// Solidity: function operatorSlashContract() view returns(address)
func (_Liq *LiqCallerSession) OperatorSlashContract() (common.Address, error) {
	return _Liq.Contract.OperatorSlashContract(&_Liq.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Liq *LiqCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Liq *LiqSession) Owner() (common.Address, error) {
	return _Liq.Contract.Owner(&_Liq.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Liq *LiqCallerSession) Owner() (common.Address, error) {
	return _Liq.Contract.Owner(&_Liq.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Liq *LiqCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Liq *LiqSession) Paused() (bool, error) {
	return _Liq.Contract.Paused(&_Liq.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Liq *LiqCallerSession) Paused() (bool, error) {
	return _Liq.Contract.Paused(&_Liq.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Liq *LiqCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Liq *LiqSession) ProxiableUUID() ([32]byte, error) {
	return _Liq.Contract.ProxiableUUID(&_Liq.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Liq *LiqCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Liq.Contract.ProxiableUUID(&_Liq.CallOpts)
}

// ReAssignRecords is a free data retrieval call binding the contract method 0x79c5092b.
//
// Solidity: function reAssignRecords(uint256 ) view returns(uint256)
func (_Liq *LiqCaller) ReAssignRecords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "reAssignRecords", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReAssignRecords is a free data retrieval call binding the contract method 0x79c5092b.
//
// Solidity: function reAssignRecords(uint256 ) view returns(uint256)
func (_Liq *LiqSession) ReAssignRecords(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.ReAssignRecords(&_Liq.CallOpts, arg0)
}

// ReAssignRecords is a free data retrieval call binding the contract method 0x79c5092b.
//
// Solidity: function reAssignRecords(uint256 ) view returns(uint256)
func (_Liq *LiqCallerSession) ReAssignRecords(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.ReAssignRecords(&_Liq.CallOpts, arg0)
}

// VNFTContract is a free data retrieval call binding the contract method 0xfc03411f.
//
// Solidity: function vNFTContract() view returns(address)
func (_Liq *LiqCaller) VNFTContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "vNFTContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VNFTContract is a free data retrieval call binding the contract method 0xfc03411f.
//
// Solidity: function vNFTContract() view returns(address)
func (_Liq *LiqSession) VNFTContract() (common.Address, error) {
	return _Liq.Contract.VNFTContract(&_Liq.CallOpts)
}

// VNFTContract is a free data retrieval call binding the contract method 0xfc03411f.
//
// Solidity: function vNFTContract() view returns(address)
func (_Liq *LiqCallerSession) VNFTContract() (common.Address, error) {
	return _Liq.Contract.VNFTContract(&_Liq.CallOpts)
}

// VaultManagerContractAddress is a free data retrieval call binding the contract method 0xb3ee9d6b.
//
// Solidity: function vaultManagerContractAddress() view returns(address)
func (_Liq *LiqCaller) VaultManagerContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "vaultManagerContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultManagerContractAddress is a free data retrieval call binding the contract method 0xb3ee9d6b.
//
// Solidity: function vaultManagerContractAddress() view returns(address)
func (_Liq *LiqSession) VaultManagerContractAddress() (common.Address, error) {
	return _Liq.Contract.VaultManagerContractAddress(&_Liq.CallOpts)
}

// VaultManagerContractAddress is a free data retrieval call binding the contract method 0xb3ee9d6b.
//
// Solidity: function vaultManagerContractAddress() view returns(address)
func (_Liq *LiqCallerSession) VaultManagerContractAddress() (common.Address, error) {
	return _Liq.Contract.VaultManagerContractAddress(&_Liq.CallOpts)
}

// WithdrawalRequestContract is a free data retrieval call binding the contract method 0x45dcb639.
//
// Solidity: function withdrawalRequestContract() view returns(address)
func (_Liq *LiqCaller) WithdrawalRequestContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "withdrawalRequestContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawalRequestContract is a free data retrieval call binding the contract method 0x45dcb639.
//
// Solidity: function withdrawalRequestContract() view returns(address)
func (_Liq *LiqSession) WithdrawalRequestContract() (common.Address, error) {
	return _Liq.Contract.WithdrawalRequestContract(&_Liq.CallOpts)
}

// WithdrawalRequestContract is a free data retrieval call binding the contract method 0x45dcb639.
//
// Solidity: function withdrawalRequestContract() view returns(address)
func (_Liq *LiqCallerSession) WithdrawalRequestContract() (common.Address, error) {
	return _Liq.Contract.WithdrawalRequestContract(&_Liq.CallOpts)
}

// AddSlashFundToStakePool is a paid mutator transaction binding the contract method 0x76b69ca6.
//
// Solidity: function addSlashFundToStakePool(uint256 _operatorId, uint256 _amount) payable returns()
func (_Liq *LiqTransactor) AddSlashFundToStakePool(opts *bind.TransactOpts, _operatorId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "addSlashFundToStakePool", _operatorId, _amount)
}

// AddSlashFundToStakePool is a paid mutator transaction binding the contract method 0x76b69ca6.
//
// Solidity: function addSlashFundToStakePool(uint256 _operatorId, uint256 _amount) payable returns()
func (_Liq *LiqSession) AddSlashFundToStakePool(_operatorId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.AddSlashFundToStakePool(&_Liq.TransactOpts, _operatorId, _amount)
}

// AddSlashFundToStakePool is a paid mutator transaction binding the contract method 0x76b69ca6.
//
// Solidity: function addSlashFundToStakePool(uint256 _operatorId, uint256 _amount) payable returns()
func (_Liq *LiqTransactorSession) AddSlashFundToStakePool(_operatorId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.AddSlashFundToStakePool(&_Liq.TransactOpts, _operatorId, _amount)
}

// AssignOperator is a paid mutator transaction binding the contract method 0xb36e2cce.
//
// Solidity: function assignOperator(uint256 _assignOperatorId, uint256 _operatorId) returns()
func (_Liq *LiqTransactor) AssignOperator(opts *bind.TransactOpts, _assignOperatorId *big.Int, _operatorId *big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "assignOperator", _assignOperatorId, _operatorId)
}

// AssignOperator is a paid mutator transaction binding the contract method 0xb36e2cce.
//
// Solidity: function assignOperator(uint256 _assignOperatorId, uint256 _operatorId) returns()
func (_Liq *LiqSession) AssignOperator(_assignOperatorId *big.Int, _operatorId *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.AssignOperator(&_Liq.TransactOpts, _assignOperatorId, _operatorId)
}

// AssignOperator is a paid mutator transaction binding the contract method 0xb36e2cce.
//
// Solidity: function assignOperator(uint256 _assignOperatorId, uint256 _operatorId) returns()
func (_Liq *LiqTransactorSession) AssignOperator(_assignOperatorId *big.Int, _operatorId *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.AssignOperator(&_Liq.TransactOpts, _assignOperatorId, _operatorId)
}

// ClaimRewardsOfDao is a paid mutator transaction binding the contract method 0x8628974b.
//
// Solidity: function claimRewardsOfDao(uint256[] _operatorIds, uint256[] _rewards) returns()
func (_Liq *LiqTransactor) ClaimRewardsOfDao(opts *bind.TransactOpts, _operatorIds []*big.Int, _rewards []*big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "claimRewardsOfDao", _operatorIds, _rewards)
}

// ClaimRewardsOfDao is a paid mutator transaction binding the contract method 0x8628974b.
//
// Solidity: function claimRewardsOfDao(uint256[] _operatorIds, uint256[] _rewards) returns()
func (_Liq *LiqSession) ClaimRewardsOfDao(_operatorIds []*big.Int, _rewards []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.ClaimRewardsOfDao(&_Liq.TransactOpts, _operatorIds, _rewards)
}

// ClaimRewardsOfDao is a paid mutator transaction binding the contract method 0x8628974b.
//
// Solidity: function claimRewardsOfDao(uint256[] _operatorIds, uint256[] _rewards) returns()
func (_Liq *LiqTransactorSession) ClaimRewardsOfDao(_operatorIds []*big.Int, _rewards []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.ClaimRewardsOfDao(&_Liq.TransactOpts, _operatorIds, _rewards)
}

// ClaimRewardsOfOperator is a paid mutator transaction binding the contract method 0x15cf1a8f.
//
// Solidity: function claimRewardsOfOperator(uint256 _operatorId, address[] _rewardAddresses, uint256[] _rewards) returns()
func (_Liq *LiqTransactor) ClaimRewardsOfOperator(opts *bind.TransactOpts, _operatorId *big.Int, _rewardAddresses []common.Address, _rewards []*big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "claimRewardsOfOperator", _operatorId, _rewardAddresses, _rewards)
}

// ClaimRewardsOfOperator is a paid mutator transaction binding the contract method 0x15cf1a8f.
//
// Solidity: function claimRewardsOfOperator(uint256 _operatorId, address[] _rewardAddresses, uint256[] _rewards) returns()
func (_Liq *LiqSession) ClaimRewardsOfOperator(_operatorId *big.Int, _rewardAddresses []common.Address, _rewards []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.ClaimRewardsOfOperator(&_Liq.TransactOpts, _operatorId, _rewardAddresses, _rewards)
}

// ClaimRewardsOfOperator is a paid mutator transaction binding the contract method 0x15cf1a8f.
//
// Solidity: function claimRewardsOfOperator(uint256 _operatorId, address[] _rewardAddresses, uint256[] _rewards) returns()
func (_Liq *LiqTransactorSession) ClaimRewardsOfOperator(_operatorId *big.Int, _rewardAddresses []common.Address, _rewards []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.ClaimRewardsOfOperator(&_Liq.TransactOpts, _operatorId, _rewardAddresses, _rewards)
}

// ClaimRewardsOfUser is a paid mutator transaction binding the contract method 0x7d75de55.
//
// Solidity: function claimRewardsOfUser(uint256 _operatorId, uint256[] _tokenIds, uint256 _totalNftRewards, uint256 _gasHeight, address _owner) returns()
func (_Liq *LiqTransactor) ClaimRewardsOfUser(opts *bind.TransactOpts, _operatorId *big.Int, _tokenIds []*big.Int, _totalNftRewards *big.Int, _gasHeight *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "claimRewardsOfUser", _operatorId, _tokenIds, _totalNftRewards, _gasHeight, _owner)
}

// ClaimRewardsOfUser is a paid mutator transaction binding the contract method 0x7d75de55.
//
// Solidity: function claimRewardsOfUser(uint256 _operatorId, uint256[] _tokenIds, uint256 _totalNftRewards, uint256 _gasHeight, address _owner) returns()
func (_Liq *LiqSession) ClaimRewardsOfUser(_operatorId *big.Int, _tokenIds []*big.Int, _totalNftRewards *big.Int, _gasHeight *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Liq.Contract.ClaimRewardsOfUser(&_Liq.TransactOpts, _operatorId, _tokenIds, _totalNftRewards, _gasHeight, _owner)
}

// ClaimRewardsOfUser is a paid mutator transaction binding the contract method 0x7d75de55.
//
// Solidity: function claimRewardsOfUser(uint256 _operatorId, uint256[] _tokenIds, uint256 _totalNftRewards, uint256 _gasHeight, address _owner) returns()
func (_Liq *LiqTransactorSession) ClaimRewardsOfUser(_operatorId *big.Int, _tokenIds []*big.Int, _totalNftRewards *big.Int, _gasHeight *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Liq.Contract.ClaimRewardsOfUser(&_Liq.TransactOpts, _operatorId, _tokenIds, _totalNftRewards, _gasHeight, _owner)
}

// FastUnstakeNFT is a paid mutator transaction binding the contract method 0x37a203c8.
//
// Solidity: function fastUnstakeNFT(uint256 _operatorId, uint256 _tokenId, address _to) returns()
func (_Liq *LiqTransactor) FastUnstakeNFT(opts *bind.TransactOpts, _operatorId *big.Int, _tokenId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "fastUnstakeNFT", _operatorId, _tokenId, _to)
}

// FastUnstakeNFT is a paid mutator transaction binding the contract method 0x37a203c8.
//
// Solidity: function fastUnstakeNFT(uint256 _operatorId, uint256 _tokenId, address _to) returns()
func (_Liq *LiqSession) FastUnstakeNFT(_operatorId *big.Int, _tokenId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Liq.Contract.FastUnstakeNFT(&_Liq.TransactOpts, _operatorId, _tokenId, _to)
}

// FastUnstakeNFT is a paid mutator transaction binding the contract method 0x37a203c8.
//
// Solidity: function fastUnstakeNFT(uint256 _operatorId, uint256 _tokenId, address _to) returns()
func (_Liq *LiqTransactorSession) FastUnstakeNFT(_operatorId *big.Int, _tokenId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Liq.Contract.FastUnstakeNFT(&_Liq.TransactOpts, _operatorId, _tokenId, _to)
}

// Initialize is a paid mutator transaction binding the contract method 0x0eda4cf2.
//
// Solidity: function initialize(address _dao, address _daoVaultAddress, bytes _withdrawalCreds, address _nodeOperatorRegistryContractAddress, address _nETHContractAddress, address _nVNFTContractAddress, address _beaconOracleContractAddress, address _depositContractAddress) returns()
func (_Liq *LiqTransactor) Initialize(opts *bind.TransactOpts, _dao common.Address, _daoVaultAddress common.Address, _withdrawalCreds []byte, _nodeOperatorRegistryContractAddress common.Address, _nETHContractAddress common.Address, _nVNFTContractAddress common.Address, _beaconOracleContractAddress common.Address, _depositContractAddress common.Address) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "initialize", _dao, _daoVaultAddress, _withdrawalCreds, _nodeOperatorRegistryContractAddress, _nETHContractAddress, _nVNFTContractAddress, _beaconOracleContractAddress, _depositContractAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x0eda4cf2.
//
// Solidity: function initialize(address _dao, address _daoVaultAddress, bytes _withdrawalCreds, address _nodeOperatorRegistryContractAddress, address _nETHContractAddress, address _nVNFTContractAddress, address _beaconOracleContractAddress, address _depositContractAddress) returns()
func (_Liq *LiqSession) Initialize(_dao common.Address, _daoVaultAddress common.Address, _withdrawalCreds []byte, _nodeOperatorRegistryContractAddress common.Address, _nETHContractAddress common.Address, _nVNFTContractAddress common.Address, _beaconOracleContractAddress common.Address, _depositContractAddress common.Address) (*types.Transaction, error) {
	return _Liq.Contract.Initialize(&_Liq.TransactOpts, _dao, _daoVaultAddress, _withdrawalCreds, _nodeOperatorRegistryContractAddress, _nETHContractAddress, _nVNFTContractAddress, _beaconOracleContractAddress, _depositContractAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x0eda4cf2.
//
// Solidity: function initialize(address _dao, address _daoVaultAddress, bytes _withdrawalCreds, address _nodeOperatorRegistryContractAddress, address _nETHContractAddress, address _nVNFTContractAddress, address _beaconOracleContractAddress, address _depositContractAddress) returns()
func (_Liq *LiqTransactorSession) Initialize(_dao common.Address, _daoVaultAddress common.Address, _withdrawalCreds []byte, _nodeOperatorRegistryContractAddress common.Address, _nETHContractAddress common.Address, _nVNFTContractAddress common.Address, _beaconOracleContractAddress common.Address, _depositContractAddress common.Address) (*types.Transaction, error) {
	return _Liq.Contract.Initialize(&_Liq.TransactOpts, _dao, _daoVaultAddress, _withdrawalCreds, _nodeOperatorRegistryContractAddress, _nETHContractAddress, _nVNFTContractAddress, _beaconOracleContractAddress, _depositContractAddress)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x04517e31.
//
// Solidity: function initializeV2(uint256[] _operatorIds, address[] _users, uint256[] _nethAmounts, address _consensusVaultContractAddress, address _vaultManagerContractAddress, address _withdrawalRequestContractAddress, address _operatorSlashContractAddress) returns()
func (_Liq *LiqTransactor) InitializeV2(opts *bind.TransactOpts, _operatorIds []*big.Int, _users []common.Address, _nethAmounts []*big.Int, _consensusVaultContractAddress common.Address, _vaultManagerContractAddress common.Address, _withdrawalRequestContractAddress common.Address, _operatorSlashContractAddress common.Address) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "initializeV2", _operatorIds, _users, _nethAmounts, _consensusVaultContractAddress, _vaultManagerContractAddress, _withdrawalRequestContractAddress, _operatorSlashContractAddress)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x04517e31.
//
// Solidity: function initializeV2(uint256[] _operatorIds, address[] _users, uint256[] _nethAmounts, address _consensusVaultContractAddress, address _vaultManagerContractAddress, address _withdrawalRequestContractAddress, address _operatorSlashContractAddress) returns()
func (_Liq *LiqSession) InitializeV2(_operatorIds []*big.Int, _users []common.Address, _nethAmounts []*big.Int, _consensusVaultContractAddress common.Address, _vaultManagerContractAddress common.Address, _withdrawalRequestContractAddress common.Address, _operatorSlashContractAddress common.Address) (*types.Transaction, error) {
	return _Liq.Contract.InitializeV2(&_Liq.TransactOpts, _operatorIds, _users, _nethAmounts, _consensusVaultContractAddress, _vaultManagerContractAddress, _withdrawalRequestContractAddress, _operatorSlashContractAddress)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x04517e31.
//
// Solidity: function initializeV2(uint256[] _operatorIds, address[] _users, uint256[] _nethAmounts, address _consensusVaultContractAddress, address _vaultManagerContractAddress, address _withdrawalRequestContractAddress, address _operatorSlashContractAddress) returns()
func (_Liq *LiqTransactorSession) InitializeV2(_operatorIds []*big.Int, _users []common.Address, _nethAmounts []*big.Int, _consensusVaultContractAddress common.Address, _vaultManagerContractAddress common.Address, _withdrawalRequestContractAddress common.Address, _operatorSlashContractAddress common.Address) (*types.Transaction, error) {
	return _Liq.Contract.InitializeV2(&_Liq.TransactOpts, _operatorIds, _users, _nethAmounts, _consensusVaultContractAddress, _vaultManagerContractAddress, _withdrawalRequestContractAddress, _operatorSlashContractAddress)
}

// LargeWithdrawalBurnNeth is a paid mutator transaction binding the contract method 0x47a764d3.
//
// Solidity: function largeWithdrawalBurnNeth(uint256 _totalRequestNethAmount) returns()
func (_Liq *LiqTransactor) LargeWithdrawalBurnNeth(opts *bind.TransactOpts, _totalRequestNethAmount *big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "largeWithdrawalBurnNeth", _totalRequestNethAmount)
}

// LargeWithdrawalBurnNeth is a paid mutator transaction binding the contract method 0x47a764d3.
//
// Solidity: function largeWithdrawalBurnNeth(uint256 _totalRequestNethAmount) returns()
func (_Liq *LiqSession) LargeWithdrawalBurnNeth(_totalRequestNethAmount *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.LargeWithdrawalBurnNeth(&_Liq.TransactOpts, _totalRequestNethAmount)
}

// LargeWithdrawalBurnNeth is a paid mutator transaction binding the contract method 0x47a764d3.
//
// Solidity: function largeWithdrawalBurnNeth(uint256 _totalRequestNethAmount) returns()
func (_Liq *LiqTransactorSession) LargeWithdrawalBurnNeth(_totalRequestNethAmount *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.LargeWithdrawalBurnNeth(&_Liq.TransactOpts, _totalRequestNethAmount)
}

// LargeWithdrawalUnstake is a paid mutator transaction binding the contract method 0x3a548da1.
//
// Solidity: function largeWithdrawalUnstake(uint256 _operatorId, address _from, uint256 _amount) returns()
func (_Liq *LiqTransactor) LargeWithdrawalUnstake(opts *bind.TransactOpts, _operatorId *big.Int, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "largeWithdrawalUnstake", _operatorId, _from, _amount)
}

// LargeWithdrawalUnstake is a paid mutator transaction binding the contract method 0x3a548da1.
//
// Solidity: function largeWithdrawalUnstake(uint256 _operatorId, address _from, uint256 _amount) returns()
func (_Liq *LiqSession) LargeWithdrawalUnstake(_operatorId *big.Int, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.LargeWithdrawalUnstake(&_Liq.TransactOpts, _operatorId, _from, _amount)
}

// LargeWithdrawalUnstake is a paid mutator transaction binding the contract method 0x3a548da1.
//
// Solidity: function largeWithdrawalUnstake(uint256 _operatorId, address _from, uint256 _amount) returns()
func (_Liq *LiqTransactorSession) LargeWithdrawalUnstake(_operatorId *big.Int, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.LargeWithdrawalUnstake(&_Liq.TransactOpts, _operatorId, _from, _amount)
}

// NftExitHandle is a paid mutator transaction binding the contract method 0x5d20c6be.
//
// Solidity: function nftExitHandle(uint256[] _tokenIds, uint256[] _exitBlockNumbers) returns()
func (_Liq *LiqTransactor) NftExitHandle(opts *bind.TransactOpts, _tokenIds []*big.Int, _exitBlockNumbers []*big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "nftExitHandle", _tokenIds, _exitBlockNumbers)
}

// NftExitHandle is a paid mutator transaction binding the contract method 0x5d20c6be.
//
// Solidity: function nftExitHandle(uint256[] _tokenIds, uint256[] _exitBlockNumbers) returns()
func (_Liq *LiqSession) NftExitHandle(_tokenIds []*big.Int, _exitBlockNumbers []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.NftExitHandle(&_Liq.TransactOpts, _tokenIds, _exitBlockNumbers)
}

// NftExitHandle is a paid mutator transaction binding the contract method 0x5d20c6be.
//
// Solidity: function nftExitHandle(uint256[] _tokenIds, uint256[] _exitBlockNumbers) returns()
func (_Liq *LiqTransactorSession) NftExitHandle(_tokenIds []*big.Int, _exitBlockNumbers []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.NftExitHandle(&_Liq.TransactOpts, _tokenIds, _exitBlockNumbers)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Liq *LiqTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Liq *LiqSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Liq.Contract.OnERC721Received(&_Liq.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Liq *LiqTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Liq.Contract.OnERC721Received(&_Liq.TransactOpts, operator, from, tokenId, data)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Liq *LiqTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Liq *LiqSession) Pause() (*types.Transaction, error) {
	return _Liq.Contract.Pause(&_Liq.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Liq *LiqTransactorSession) Pause() (*types.Transaction, error) {
	return _Liq.Contract.Pause(&_Liq.TransactOpts)
}

// ReceiveRewards is a paid mutator transaction binding the contract method 0xd77868da.
//
// Solidity: function receiveRewards(uint256 _rewards) payable returns()
func (_Liq *LiqTransactor) ReceiveRewards(opts *bind.TransactOpts, _rewards *big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "receiveRewards", _rewards)
}

// ReceiveRewards is a paid mutator transaction binding the contract method 0xd77868da.
//
// Solidity: function receiveRewards(uint256 _rewards) payable returns()
func (_Liq *LiqSession) ReceiveRewards(_rewards *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.ReceiveRewards(&_Liq.TransactOpts, _rewards)
}

// ReceiveRewards is a paid mutator transaction binding the contract method 0xd77868da.
//
// Solidity: function receiveRewards(uint256 _rewards) payable returns()
func (_Liq *LiqTransactorSession) ReceiveRewards(_rewards *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.ReceiveRewards(&_Liq.TransactOpts, _rewards)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x193615fe.
//
// Solidity: function registerValidator(bytes[] _pubkeys, bytes[] _signatures, bytes32[] _depositDataRoots) returns()
func (_Liq *LiqTransactor) RegisterValidator(opts *bind.TransactOpts, _pubkeys [][]byte, _signatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "registerValidator", _pubkeys, _signatures, _depositDataRoots)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x193615fe.
//
// Solidity: function registerValidator(bytes[] _pubkeys, bytes[] _signatures, bytes32[] _depositDataRoots) returns()
func (_Liq *LiqSession) RegisterValidator(_pubkeys [][]byte, _signatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _Liq.Contract.RegisterValidator(&_Liq.TransactOpts, _pubkeys, _signatures, _depositDataRoots)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x193615fe.
//
// Solidity: function registerValidator(bytes[] _pubkeys, bytes[] _signatures, bytes32[] _depositDataRoots) returns()
func (_Liq *LiqTransactorSession) RegisterValidator(_pubkeys [][]byte, _signatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _Liq.Contract.RegisterValidator(&_Liq.TransactOpts, _pubkeys, _signatures, _depositDataRoots)
}

// ReinvestClRewards is a paid mutator transaction binding the contract method 0x5564ab6e.
//
// Solidity: function reinvestClRewards(uint256[] _operatorIds, uint256[] _amounts) returns()
func (_Liq *LiqTransactor) ReinvestClRewards(opts *bind.TransactOpts, _operatorIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "reinvestClRewards", _operatorIds, _amounts)
}

// ReinvestClRewards is a paid mutator transaction binding the contract method 0x5564ab6e.
//
// Solidity: function reinvestClRewards(uint256[] _operatorIds, uint256[] _amounts) returns()
func (_Liq *LiqSession) ReinvestClRewards(_operatorIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.ReinvestClRewards(&_Liq.TransactOpts, _operatorIds, _amounts)
}

// ReinvestClRewards is a paid mutator transaction binding the contract method 0x5564ab6e.
//
// Solidity: function reinvestClRewards(uint256[] _operatorIds, uint256[] _amounts) returns()
func (_Liq *LiqTransactorSession) ReinvestClRewards(_operatorIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.ReinvestClRewards(&_Liq.TransactOpts, _operatorIds, _amounts)
}

// ReinvestElRewards is a paid mutator transaction binding the contract method 0x0d020de2.
//
// Solidity: function reinvestElRewards(uint256[] _operatorIds, uint256[] _amounts) returns()
func (_Liq *LiqTransactor) ReinvestElRewards(opts *bind.TransactOpts, _operatorIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "reinvestElRewards", _operatorIds, _amounts)
}

// ReinvestElRewards is a paid mutator transaction binding the contract method 0x0d020de2.
//
// Solidity: function reinvestElRewards(uint256[] _operatorIds, uint256[] _amounts) returns()
func (_Liq *LiqSession) ReinvestElRewards(_operatorIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.ReinvestElRewards(&_Liq.TransactOpts, _operatorIds, _amounts)
}

// ReinvestElRewards is a paid mutator transaction binding the contract method 0x0d020de2.
//
// Solidity: function reinvestElRewards(uint256[] _operatorIds, uint256[] _amounts) returns()
func (_Liq *LiqTransactorSession) ReinvestElRewards(_operatorIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.ReinvestElRewards(&_Liq.TransactOpts, _operatorIds, _amounts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Liq *LiqTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Liq *LiqSession) RenounceOwnership() (*types.Transaction, error) {
	return _Liq.Contract.RenounceOwnership(&_Liq.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Liq *LiqTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Liq.Contract.RenounceOwnership(&_Liq.TransactOpts)
}

// SetBeaconOracleContract is a paid mutator transaction binding the contract method 0xeff71b99.
//
// Solidity: function setBeaconOracleContract(address _beaconOracleContractAddress) returns()
func (_Liq *LiqTransactor) SetBeaconOracleContract(opts *bind.TransactOpts, _beaconOracleContractAddress common.Address) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "setBeaconOracleContract", _beaconOracleContractAddress)
}

// SetBeaconOracleContract is a paid mutator transaction binding the contract method 0xeff71b99.
//
// Solidity: function setBeaconOracleContract(address _beaconOracleContractAddress) returns()
func (_Liq *LiqSession) SetBeaconOracleContract(_beaconOracleContractAddress common.Address) (*types.Transaction, error) {
	return _Liq.Contract.SetBeaconOracleContract(&_Liq.TransactOpts, _beaconOracleContractAddress)
}

// SetBeaconOracleContract is a paid mutator transaction binding the contract method 0xeff71b99.
//
// Solidity: function setBeaconOracleContract(address _beaconOracleContractAddress) returns()
func (_Liq *LiqTransactorSession) SetBeaconOracleContract(_beaconOracleContractAddress common.Address) (*types.Transaction, error) {
	return _Liq.Contract.SetBeaconOracleContract(&_Liq.TransactOpts, _beaconOracleContractAddress)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_Liq *LiqTransactor) SetDaoAddress(opts *bind.TransactOpts, _dao common.Address) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "setDaoAddress", _dao)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_Liq *LiqSession) SetDaoAddress(_dao common.Address) (*types.Transaction, error) {
	return _Liq.Contract.SetDaoAddress(&_Liq.TransactOpts, _dao)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _dao) returns()
func (_Liq *LiqTransactorSession) SetDaoAddress(_dao common.Address) (*types.Transaction, error) {
	return _Liq.Contract.SetDaoAddress(&_Liq.TransactOpts, _dao)
}

// SetDaoVaultAddress is a paid mutator transaction binding the contract method 0xc720d250.
//
// Solidity: function setDaoVaultAddress(address _daoVaultAddress) returns()
func (_Liq *LiqTransactor) SetDaoVaultAddress(opts *bind.TransactOpts, _daoVaultAddress common.Address) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "setDaoVaultAddress", _daoVaultAddress)
}

// SetDaoVaultAddress is a paid mutator transaction binding the contract method 0xc720d250.
//
// Solidity: function setDaoVaultAddress(address _daoVaultAddress) returns()
func (_Liq *LiqSession) SetDaoVaultAddress(_daoVaultAddress common.Address) (*types.Transaction, error) {
	return _Liq.Contract.SetDaoVaultAddress(&_Liq.TransactOpts, _daoVaultAddress)
}

// SetDaoVaultAddress is a paid mutator transaction binding the contract method 0xc720d250.
//
// Solidity: function setDaoVaultAddress(address _daoVaultAddress) returns()
func (_Liq *LiqTransactorSession) SetDaoVaultAddress(_daoVaultAddress common.Address) (*types.Transaction, error) {
	return _Liq.Contract.SetDaoVaultAddress(&_Liq.TransactOpts, _daoVaultAddress)
}

// SetDepositFeeRate is a paid mutator transaction binding the contract method 0x8e691b9a.
//
// Solidity: function setDepositFeeRate(uint256 _feeRate) returns()
func (_Liq *LiqTransactor) SetDepositFeeRate(opts *bind.TransactOpts, _feeRate *big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "setDepositFeeRate", _feeRate)
}

// SetDepositFeeRate is a paid mutator transaction binding the contract method 0x8e691b9a.
//
// Solidity: function setDepositFeeRate(uint256 _feeRate) returns()
func (_Liq *LiqSession) SetDepositFeeRate(_feeRate *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.SetDepositFeeRate(&_Liq.TransactOpts, _feeRate)
}

// SetDepositFeeRate is a paid mutator transaction binding the contract method 0x8e691b9a.
//
// Solidity: function setDepositFeeRate(uint256 _feeRate) returns()
func (_Liq *LiqTransactorSession) SetDepositFeeRate(_feeRate *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.SetDepositFeeRate(&_Liq.TransactOpts, _feeRate)
}

// SetLiquidStakingWithdrawalCredentials is a paid mutator transaction binding the contract method 0x2fa2a2cc.
//
// Solidity: function setLiquidStakingWithdrawalCredentials(bytes _liquidStakingWithdrawalCredentials) returns()
func (_Liq *LiqTransactor) SetLiquidStakingWithdrawalCredentials(opts *bind.TransactOpts, _liquidStakingWithdrawalCredentials []byte) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "setLiquidStakingWithdrawalCredentials", _liquidStakingWithdrawalCredentials)
}

// SetLiquidStakingWithdrawalCredentials is a paid mutator transaction binding the contract method 0x2fa2a2cc.
//
// Solidity: function setLiquidStakingWithdrawalCredentials(bytes _liquidStakingWithdrawalCredentials) returns()
func (_Liq *LiqSession) SetLiquidStakingWithdrawalCredentials(_liquidStakingWithdrawalCredentials []byte) (*types.Transaction, error) {
	return _Liq.Contract.SetLiquidStakingWithdrawalCredentials(&_Liq.TransactOpts, _liquidStakingWithdrawalCredentials)
}

// SetLiquidStakingWithdrawalCredentials is a paid mutator transaction binding the contract method 0x2fa2a2cc.
//
// Solidity: function setLiquidStakingWithdrawalCredentials(bytes _liquidStakingWithdrawalCredentials) returns()
func (_Liq *LiqTransactorSession) SetLiquidStakingWithdrawalCredentials(_liquidStakingWithdrawalCredentials []byte) (*types.Transaction, error) {
	return _Liq.Contract.SetLiquidStakingWithdrawalCredentials(&_Liq.TransactOpts, _liquidStakingWithdrawalCredentials)
}

// SetNodeOperatorRegistryContract is a paid mutator transaction binding the contract method 0xcb23473e.
//
// Solidity: function setNodeOperatorRegistryContract(address _nodeOperatorRegistryContract) returns()
func (_Liq *LiqTransactor) SetNodeOperatorRegistryContract(opts *bind.TransactOpts, _nodeOperatorRegistryContract common.Address) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "setNodeOperatorRegistryContract", _nodeOperatorRegistryContract)
}

// SetNodeOperatorRegistryContract is a paid mutator transaction binding the contract method 0xcb23473e.
//
// Solidity: function setNodeOperatorRegistryContract(address _nodeOperatorRegistryContract) returns()
func (_Liq *LiqSession) SetNodeOperatorRegistryContract(_nodeOperatorRegistryContract common.Address) (*types.Transaction, error) {
	return _Liq.Contract.SetNodeOperatorRegistryContract(&_Liq.TransactOpts, _nodeOperatorRegistryContract)
}

// SetNodeOperatorRegistryContract is a paid mutator transaction binding the contract method 0xcb23473e.
//
// Solidity: function setNodeOperatorRegistryContract(address _nodeOperatorRegistryContract) returns()
func (_Liq *LiqTransactorSession) SetNodeOperatorRegistryContract(_nodeOperatorRegistryContract common.Address) (*types.Transaction, error) {
	return _Liq.Contract.SetNodeOperatorRegistryContract(&_Liq.TransactOpts, _nodeOperatorRegistryContract)
}

// SetOperatorCanLoanAmounts is a paid mutator transaction binding the contract method 0xe59546be.
//
// Solidity: function setOperatorCanLoanAmounts(uint256 _newCanloadAmounts) returns()
func (_Liq *LiqTransactor) SetOperatorCanLoanAmounts(opts *bind.TransactOpts, _newCanloadAmounts *big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "setOperatorCanLoanAmounts", _newCanloadAmounts)
}

// SetOperatorCanLoanAmounts is a paid mutator transaction binding the contract method 0xe59546be.
//
// Solidity: function setOperatorCanLoanAmounts(uint256 _newCanloadAmounts) returns()
func (_Liq *LiqSession) SetOperatorCanLoanAmounts(_newCanloadAmounts *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.SetOperatorCanLoanAmounts(&_Liq.TransactOpts, _newCanloadAmounts)
}

// SetOperatorCanLoanAmounts is a paid mutator transaction binding the contract method 0xe59546be.
//
// Solidity: function setOperatorCanLoanAmounts(uint256 _newCanloadAmounts) returns()
func (_Liq *LiqTransactorSession) SetOperatorCanLoanAmounts(_newCanloadAmounts *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.SetOperatorCanLoanAmounts(&_Liq.TransactOpts, _newCanloadAmounts)
}

// SetVaultManagerContract is a paid mutator transaction binding the contract method 0xbe14652a.
//
// Solidity: function setVaultManagerContract(address _vaultManagerContract) returns()
func (_Liq *LiqTransactor) SetVaultManagerContract(opts *bind.TransactOpts, _vaultManagerContract common.Address) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "setVaultManagerContract", _vaultManagerContract)
}

// SetVaultManagerContract is a paid mutator transaction binding the contract method 0xbe14652a.
//
// Solidity: function setVaultManagerContract(address _vaultManagerContract) returns()
func (_Liq *LiqSession) SetVaultManagerContract(_vaultManagerContract common.Address) (*types.Transaction, error) {
	return _Liq.Contract.SetVaultManagerContract(&_Liq.TransactOpts, _vaultManagerContract)
}

// SetVaultManagerContract is a paid mutator transaction binding the contract method 0xbe14652a.
//
// Solidity: function setVaultManagerContract(address _vaultManagerContract) returns()
func (_Liq *LiqTransactorSession) SetVaultManagerContract(_vaultManagerContract common.Address) (*types.Transaction, error) {
	return _Liq.Contract.SetVaultManagerContract(&_Liq.TransactOpts, _vaultManagerContract)
}

// StakeETH is a paid mutator transaction binding the contract method 0x6ecc20da.
//
// Solidity: function stakeETH(uint256 _operatorId) payable returns()
func (_Liq *LiqTransactor) StakeETH(opts *bind.TransactOpts, _operatorId *big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "stakeETH", _operatorId)
}

// StakeETH is a paid mutator transaction binding the contract method 0x6ecc20da.
//
// Solidity: function stakeETH(uint256 _operatorId) payable returns()
func (_Liq *LiqSession) StakeETH(_operatorId *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.StakeETH(&_Liq.TransactOpts, _operatorId)
}

// StakeETH is a paid mutator transaction binding the contract method 0x6ecc20da.
//
// Solidity: function stakeETH(uint256 _operatorId) payable returns()
func (_Liq *LiqTransactorSession) StakeETH(_operatorId *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.StakeETH(&_Liq.TransactOpts, _operatorId)
}

// StakeNFT is a paid mutator transaction binding the contract method 0x75d1b234.
//
// Solidity: function stakeNFT(uint256 _operatorId, address withdrawalCredentialsAddress) payable returns()
func (_Liq *LiqTransactor) StakeNFT(opts *bind.TransactOpts, _operatorId *big.Int, withdrawalCredentialsAddress common.Address) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "stakeNFT", _operatorId, withdrawalCredentialsAddress)
}

// StakeNFT is a paid mutator transaction binding the contract method 0x75d1b234.
//
// Solidity: function stakeNFT(uint256 _operatorId, address withdrawalCredentialsAddress) payable returns()
func (_Liq *LiqSession) StakeNFT(_operatorId *big.Int, withdrawalCredentialsAddress common.Address) (*types.Transaction, error) {
	return _Liq.Contract.StakeNFT(&_Liq.TransactOpts, _operatorId, withdrawalCredentialsAddress)
}

// StakeNFT is a paid mutator transaction binding the contract method 0x75d1b234.
//
// Solidity: function stakeNFT(uint256 _operatorId, address withdrawalCredentialsAddress) payable returns()
func (_Liq *LiqTransactorSession) StakeNFT(_operatorId *big.Int, withdrawalCredentialsAddress common.Address) (*types.Transaction, error) {
	return _Liq.Contract.StakeNFT(&_Liq.TransactOpts, _operatorId, withdrawalCredentialsAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Liq *LiqTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Liq *LiqSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Liq.Contract.TransferOwnership(&_Liq.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Liq *LiqTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Liq.Contract.TransferOwnership(&_Liq.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Liq *LiqTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Liq *LiqSession) Unpause() (*types.Transaction, error) {
	return _Liq.Contract.Unpause(&_Liq.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Liq *LiqTransactorSession) Unpause() (*types.Transaction, error) {
	return _Liq.Contract.Unpause(&_Liq.TransactOpts)
}

// UnstakeETH is a paid mutator transaction binding the contract method 0xe425353d.
//
// Solidity: function unstakeETH(uint256 _operatorId, uint256 _amounts) returns()
func (_Liq *LiqTransactor) UnstakeETH(opts *bind.TransactOpts, _operatorId *big.Int, _amounts *big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "unstakeETH", _operatorId, _amounts)
}

// UnstakeETH is a paid mutator transaction binding the contract method 0xe425353d.
//
// Solidity: function unstakeETH(uint256 _operatorId, uint256 _amounts) returns()
func (_Liq *LiqSession) UnstakeETH(_operatorId *big.Int, _amounts *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.UnstakeETH(&_Liq.TransactOpts, _operatorId, _amounts)
}

// UnstakeETH is a paid mutator transaction binding the contract method 0xe425353d.
//
// Solidity: function unstakeETH(uint256 _operatorId, uint256 _amounts) returns()
func (_Liq *LiqTransactorSession) UnstakeETH(_operatorId *big.Int, _amounts *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.UnstakeETH(&_Liq.TransactOpts, _operatorId, _amounts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Liq *LiqTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Liq *LiqSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Liq.Contract.UpgradeTo(&_Liq.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Liq *LiqTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Liq.Contract.UpgradeTo(&_Liq.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Liq *LiqTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Liq *LiqSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Liq.Contract.UpgradeToAndCall(&_Liq.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Liq *LiqTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Liq.Contract.UpgradeToAndCall(&_Liq.TransactOpts, newImplementation, data)
}

// LiqAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Liq contract.
type LiqAdminChangedIterator struct {
	Event *LiqAdminChanged // Event containing the contract specifics and raw log

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
func (it *LiqAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqAdminChanged)
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
		it.Event = new(LiqAdminChanged)
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
func (it *LiqAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqAdminChanged represents a AdminChanged event raised by the Liq contract.
type LiqAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Liq *LiqFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*LiqAdminChangedIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &LiqAdminChangedIterator{contract: _Liq.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Liq *LiqFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *LiqAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqAdminChanged)
				if err := _Liq.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Liq *LiqFilterer) ParseAdminChanged(log types.Log) (*LiqAdminChanged, error) {
	event := new(LiqAdminChanged)
	if err := _Liq.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqBeaconOracleContractSetIterator is returned from FilterBeaconOracleContractSet and is used to iterate over the raw logs and unpacked data for BeaconOracleContractSet events raised by the Liq contract.
type LiqBeaconOracleContractSetIterator struct {
	Event *LiqBeaconOracleContractSet // Event containing the contract specifics and raw log

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
func (it *LiqBeaconOracleContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqBeaconOracleContractSet)
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
		it.Event = new(LiqBeaconOracleContractSet)
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
func (it *LiqBeaconOracleContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqBeaconOracleContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqBeaconOracleContractSet represents a BeaconOracleContractSet event raised by the Liq contract.
type LiqBeaconOracleContractSet struct {
	OldBeaconOracleContract     common.Address
	BeaconOracleContractAddress common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterBeaconOracleContractSet is a free log retrieval operation binding the contract event 0x957b5ce8863e1faf1047ac1773f54e9f0598828c9d02178d6705af41d27c4f46.
//
// Solidity: event BeaconOracleContractSet(address _oldBeaconOracleContract, address _beaconOracleContractAddress)
func (_Liq *LiqFilterer) FilterBeaconOracleContractSet(opts *bind.FilterOpts) (*LiqBeaconOracleContractSetIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "BeaconOracleContractSet")
	if err != nil {
		return nil, err
	}
	return &LiqBeaconOracleContractSetIterator{contract: _Liq.contract, event: "BeaconOracleContractSet", logs: logs, sub: sub}, nil
}

// WatchBeaconOracleContractSet is a free log subscription operation binding the contract event 0x957b5ce8863e1faf1047ac1773f54e9f0598828c9d02178d6705af41d27c4f46.
//
// Solidity: event BeaconOracleContractSet(address _oldBeaconOracleContract, address _beaconOracleContractAddress)
func (_Liq *LiqFilterer) WatchBeaconOracleContractSet(opts *bind.WatchOpts, sink chan<- *LiqBeaconOracleContractSet) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "BeaconOracleContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqBeaconOracleContractSet)
				if err := _Liq.contract.UnpackLog(event, "BeaconOracleContractSet", log); err != nil {
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

// ParseBeaconOracleContractSet is a log parse operation binding the contract event 0x957b5ce8863e1faf1047ac1773f54e9f0598828c9d02178d6705af41d27c4f46.
//
// Solidity: event BeaconOracleContractSet(address _oldBeaconOracleContract, address _beaconOracleContractAddress)
func (_Liq *LiqFilterer) ParseBeaconOracleContractSet(log types.Log) (*LiqBeaconOracleContractSet, error) {
	event := new(LiqBeaconOracleContractSet)
	if err := _Liq.contract.UnpackLog(event, "BeaconOracleContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Liq contract.
type LiqBeaconUpgradedIterator struct {
	Event *LiqBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *LiqBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqBeaconUpgraded)
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
		it.Event = new(LiqBeaconUpgraded)
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
func (it *LiqBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqBeaconUpgraded represents a BeaconUpgraded event raised by the Liq contract.
type LiqBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Liq *LiqFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*LiqBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Liq.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &LiqBeaconUpgradedIterator{contract: _Liq.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Liq *LiqFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *LiqBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Liq.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqBeaconUpgraded)
				if err := _Liq.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Liq *LiqFilterer) ParseBeaconUpgraded(log types.Log) (*LiqBeaconUpgraded, error) {
	event := new(LiqBeaconUpgraded)
	if err := _Liq.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqConsensusVaultContractSetIterator is returned from FilterConsensusVaultContractSet and is used to iterate over the raw logs and unpacked data for ConsensusVaultContractSet events raised by the Liq contract.
type LiqConsensusVaultContractSetIterator struct {
	Event *LiqConsensusVaultContractSet // Event containing the contract specifics and raw log

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
func (it *LiqConsensusVaultContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqConsensusVaultContractSet)
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
		it.Event = new(LiqConsensusVaultContractSet)
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
func (it *LiqConsensusVaultContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqConsensusVaultContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqConsensusVaultContractSet represents a ConsensusVaultContractSet event raised by the Liq contract.
type LiqConsensusVaultContractSet struct {
	VaultManagerContractAddress common.Address
	ConsensusVaultContract      common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterConsensusVaultContractSet is a free log retrieval operation binding the contract event 0x955eb996feefc76589abe69cb5c8a3dfdf8cfa4fa9cec1611c9c3e61de5f55bf.
//
// Solidity: event ConsensusVaultContractSet(address vaultManagerContractAddress, address _consensusVaultContract)
func (_Liq *LiqFilterer) FilterConsensusVaultContractSet(opts *bind.FilterOpts) (*LiqConsensusVaultContractSetIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "ConsensusVaultContractSet")
	if err != nil {
		return nil, err
	}
	return &LiqConsensusVaultContractSetIterator{contract: _Liq.contract, event: "ConsensusVaultContractSet", logs: logs, sub: sub}, nil
}

// WatchConsensusVaultContractSet is a free log subscription operation binding the contract event 0x955eb996feefc76589abe69cb5c8a3dfdf8cfa4fa9cec1611c9c3e61de5f55bf.
//
// Solidity: event ConsensusVaultContractSet(address vaultManagerContractAddress, address _consensusVaultContract)
func (_Liq *LiqFilterer) WatchConsensusVaultContractSet(opts *bind.WatchOpts, sink chan<- *LiqConsensusVaultContractSet) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "ConsensusVaultContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqConsensusVaultContractSet)
				if err := _Liq.contract.UnpackLog(event, "ConsensusVaultContractSet", log); err != nil {
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

// ParseConsensusVaultContractSet is a log parse operation binding the contract event 0x955eb996feefc76589abe69cb5c8a3dfdf8cfa4fa9cec1611c9c3e61de5f55bf.
//
// Solidity: event ConsensusVaultContractSet(address vaultManagerContractAddress, address _consensusVaultContract)
func (_Liq *LiqFilterer) ParseConsensusVaultContractSet(log types.Log) (*LiqConsensusVaultContractSet, error) {
	event := new(LiqConsensusVaultContractSet)
	if err := _Liq.contract.UnpackLog(event, "ConsensusVaultContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqDaoAddressChangedIterator is returned from FilterDaoAddressChanged and is used to iterate over the raw logs and unpacked data for DaoAddressChanged events raised by the Liq contract.
type LiqDaoAddressChangedIterator struct {
	Event *LiqDaoAddressChanged // Event containing the contract specifics and raw log

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
func (it *LiqDaoAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqDaoAddressChanged)
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
		it.Event = new(LiqDaoAddressChanged)
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
func (it *LiqDaoAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqDaoAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqDaoAddressChanged represents a DaoAddressChanged event raised by the Liq contract.
type LiqDaoAddressChanged struct {
	OldDao common.Address
	Dao    common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDaoAddressChanged is a free log retrieval operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address _oldDao, address _dao)
func (_Liq *LiqFilterer) FilterDaoAddressChanged(opts *bind.FilterOpts) (*LiqDaoAddressChangedIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "DaoAddressChanged")
	if err != nil {
		return nil, err
	}
	return &LiqDaoAddressChangedIterator{contract: _Liq.contract, event: "DaoAddressChanged", logs: logs, sub: sub}, nil
}

// WatchDaoAddressChanged is a free log subscription operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address _oldDao, address _dao)
func (_Liq *LiqFilterer) WatchDaoAddressChanged(opts *bind.WatchOpts, sink chan<- *LiqDaoAddressChanged) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "DaoAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqDaoAddressChanged)
				if err := _Liq.contract.UnpackLog(event, "DaoAddressChanged", log); err != nil {
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
func (_Liq *LiqFilterer) ParseDaoAddressChanged(log types.Log) (*LiqDaoAddressChanged, error) {
	event := new(LiqDaoAddressChanged)
	if err := _Liq.contract.UnpackLog(event, "DaoAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqDaoClaimRewardsIterator is returned from FilterDaoClaimRewards and is used to iterate over the raw logs and unpacked data for DaoClaimRewards events raised by the Liq contract.
type LiqDaoClaimRewardsIterator struct {
	Event *LiqDaoClaimRewards // Event containing the contract specifics and raw log

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
func (it *LiqDaoClaimRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqDaoClaimRewards)
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
		it.Event = new(LiqDaoClaimRewards)
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
func (it *LiqDaoClaimRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqDaoClaimRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqDaoClaimRewards represents a DaoClaimRewards event raised by the Liq contract.
type LiqDaoClaimRewards struct {
	OperatorId *big.Int
	Rewards    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDaoClaimRewards is a free log retrieval operation binding the contract event 0xd015384993a63fbb67be31c3c4491e03f64fa52369a08927fe6d4cba14286f21.
//
// Solidity: event DaoClaimRewards(uint256 _operatorId, uint256 _rewards)
func (_Liq *LiqFilterer) FilterDaoClaimRewards(opts *bind.FilterOpts) (*LiqDaoClaimRewardsIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "DaoClaimRewards")
	if err != nil {
		return nil, err
	}
	return &LiqDaoClaimRewardsIterator{contract: _Liq.contract, event: "DaoClaimRewards", logs: logs, sub: sub}, nil
}

// WatchDaoClaimRewards is a free log subscription operation binding the contract event 0xd015384993a63fbb67be31c3c4491e03f64fa52369a08927fe6d4cba14286f21.
//
// Solidity: event DaoClaimRewards(uint256 _operatorId, uint256 _rewards)
func (_Liq *LiqFilterer) WatchDaoClaimRewards(opts *bind.WatchOpts, sink chan<- *LiqDaoClaimRewards) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "DaoClaimRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqDaoClaimRewards)
				if err := _Liq.contract.UnpackLog(event, "DaoClaimRewards", log); err != nil {
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
func (_Liq *LiqFilterer) ParseDaoClaimRewards(log types.Log) (*LiqDaoClaimRewards, error) {
	event := new(LiqDaoClaimRewards)
	if err := _Liq.contract.UnpackLog(event, "DaoClaimRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqDaoVaultAddressChangedIterator is returned from FilterDaoVaultAddressChanged and is used to iterate over the raw logs and unpacked data for DaoVaultAddressChanged events raised by the Liq contract.
type LiqDaoVaultAddressChangedIterator struct {
	Event *LiqDaoVaultAddressChanged // Event containing the contract specifics and raw log

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
func (it *LiqDaoVaultAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqDaoVaultAddressChanged)
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
		it.Event = new(LiqDaoVaultAddressChanged)
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
func (it *LiqDaoVaultAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqDaoVaultAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqDaoVaultAddressChanged represents a DaoVaultAddressChanged event raised by the Liq contract.
type LiqDaoVaultAddressChanged struct {
	OldDaoVaultAddress common.Address
	DaoVaultAddress    common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDaoVaultAddressChanged is a free log retrieval operation binding the contract event 0x74f93434acf49508438eb6f219ca22e7e1818b620ccb7acd411c8f520b27b642.
//
// Solidity: event DaoVaultAddressChanged(address _oldDaoVaultAddress, address _daoVaultAddress)
func (_Liq *LiqFilterer) FilterDaoVaultAddressChanged(opts *bind.FilterOpts) (*LiqDaoVaultAddressChangedIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "DaoVaultAddressChanged")
	if err != nil {
		return nil, err
	}
	return &LiqDaoVaultAddressChangedIterator{contract: _Liq.contract, event: "DaoVaultAddressChanged", logs: logs, sub: sub}, nil
}

// WatchDaoVaultAddressChanged is a free log subscription operation binding the contract event 0x74f93434acf49508438eb6f219ca22e7e1818b620ccb7acd411c8f520b27b642.
//
// Solidity: event DaoVaultAddressChanged(address _oldDaoVaultAddress, address _daoVaultAddress)
func (_Liq *LiqFilterer) WatchDaoVaultAddressChanged(opts *bind.WatchOpts, sink chan<- *LiqDaoVaultAddressChanged) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "DaoVaultAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqDaoVaultAddressChanged)
				if err := _Liq.contract.UnpackLog(event, "DaoVaultAddressChanged", log); err != nil {
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
func (_Liq *LiqFilterer) ParseDaoVaultAddressChanged(log types.Log) (*LiqDaoVaultAddressChanged, error) {
	event := new(LiqDaoVaultAddressChanged)
	if err := _Liq.contract.UnpackLog(event, "DaoVaultAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqDepositFeeRateSetIterator is returned from FilterDepositFeeRateSet and is used to iterate over the raw logs and unpacked data for DepositFeeRateSet events raised by the Liq contract.
type LiqDepositFeeRateSetIterator struct {
	Event *LiqDepositFeeRateSet // Event containing the contract specifics and raw log

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
func (it *LiqDepositFeeRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqDepositFeeRateSet)
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
		it.Event = new(LiqDepositFeeRateSet)
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
func (it *LiqDepositFeeRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqDepositFeeRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqDepositFeeRateSet represents a DepositFeeRateSet event raised by the Liq contract.
type LiqDepositFeeRateSet struct {
	OldFeeRate *big.Int
	FeeRate    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositFeeRateSet is a free log retrieval operation binding the contract event 0x90e50637c808ab8723fa9e2137e8972173099a6773142dcf42f0a394f7cef34d.
//
// Solidity: event DepositFeeRateSet(uint256 _oldFeeRate, uint256 _feeRate)
func (_Liq *LiqFilterer) FilterDepositFeeRateSet(opts *bind.FilterOpts) (*LiqDepositFeeRateSetIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "DepositFeeRateSet")
	if err != nil {
		return nil, err
	}
	return &LiqDepositFeeRateSetIterator{contract: _Liq.contract, event: "DepositFeeRateSet", logs: logs, sub: sub}, nil
}

// WatchDepositFeeRateSet is a free log subscription operation binding the contract event 0x90e50637c808ab8723fa9e2137e8972173099a6773142dcf42f0a394f7cef34d.
//
// Solidity: event DepositFeeRateSet(uint256 _oldFeeRate, uint256 _feeRate)
func (_Liq *LiqFilterer) WatchDepositFeeRateSet(opts *bind.WatchOpts, sink chan<- *LiqDepositFeeRateSet) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "DepositFeeRateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqDepositFeeRateSet)
				if err := _Liq.contract.UnpackLog(event, "DepositFeeRateSet", log); err != nil {
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

// ParseDepositFeeRateSet is a log parse operation binding the contract event 0x90e50637c808ab8723fa9e2137e8972173099a6773142dcf42f0a394f7cef34d.
//
// Solidity: event DepositFeeRateSet(uint256 _oldFeeRate, uint256 _feeRate)
func (_Liq *LiqFilterer) ParseDepositFeeRateSet(log types.Log) (*LiqDepositFeeRateSet, error) {
	event := new(LiqDepositFeeRateSet)
	if err := _Liq.contract.UnpackLog(event, "DepositFeeRateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqEthStakeIterator is returned from FilterEthStake and is used to iterate over the raw logs and unpacked data for EthStake events raised by the Liq contract.
type LiqEthStakeIterator struct {
	Event *LiqEthStake // Event containing the contract specifics and raw log

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
func (it *LiqEthStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqEthStake)
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
		it.Event = new(LiqEthStake)
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
func (it *LiqEthStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqEthStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqEthStake represents a EthStake event raised by the Liq contract.
type LiqEthStake struct {
	OperatorId *big.Int
	From       common.Address
	Amount     *big.Int
	AmountOut  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEthStake is a free log retrieval operation binding the contract event 0x1cae59a31c3c6760aa08cb9c351432553e908b8e6f53e7c9ac22715c7d496179.
//
// Solidity: event EthStake(uint256 indexed _operatorId, address indexed _from, uint256 _amount, uint256 _amountOut)
func (_Liq *LiqFilterer) FilterEthStake(opts *bind.FilterOpts, _operatorId []*big.Int, _from []common.Address) (*LiqEthStakeIterator, error) {

	var _operatorIdRule []interface{}
	for _, _operatorIdItem := range _operatorId {
		_operatorIdRule = append(_operatorIdRule, _operatorIdItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Liq.contract.FilterLogs(opts, "EthStake", _operatorIdRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &LiqEthStakeIterator{contract: _Liq.contract, event: "EthStake", logs: logs, sub: sub}, nil
}

// WatchEthStake is a free log subscription operation binding the contract event 0x1cae59a31c3c6760aa08cb9c351432553e908b8e6f53e7c9ac22715c7d496179.
//
// Solidity: event EthStake(uint256 indexed _operatorId, address indexed _from, uint256 _amount, uint256 _amountOut)
func (_Liq *LiqFilterer) WatchEthStake(opts *bind.WatchOpts, sink chan<- *LiqEthStake, _operatorId []*big.Int, _from []common.Address) (event.Subscription, error) {

	var _operatorIdRule []interface{}
	for _, _operatorIdItem := range _operatorId {
		_operatorIdRule = append(_operatorIdRule, _operatorIdItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Liq.contract.WatchLogs(opts, "EthStake", _operatorIdRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqEthStake)
				if err := _Liq.contract.UnpackLog(event, "EthStake", log); err != nil {
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

// ParseEthStake is a log parse operation binding the contract event 0x1cae59a31c3c6760aa08cb9c351432553e908b8e6f53e7c9ac22715c7d496179.
//
// Solidity: event EthStake(uint256 indexed _operatorId, address indexed _from, uint256 _amount, uint256 _amountOut)
func (_Liq *LiqFilterer) ParseEthStake(log types.Log) (*LiqEthStake, error) {
	event := new(LiqEthStake)
	if err := _Liq.contract.UnpackLog(event, "EthStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqEthUnstakeIterator is returned from FilterEthUnstake and is used to iterate over the raw logs and unpacked data for EthUnstake events raised by the Liq contract.
type LiqEthUnstakeIterator struct {
	Event *LiqEthUnstake // Event containing the contract specifics and raw log

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
func (it *LiqEthUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqEthUnstake)
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
		it.Event = new(LiqEthUnstake)
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
func (it *LiqEthUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqEthUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqEthUnstake represents a EthUnstake event raised by the Liq contract.
type LiqEthUnstake struct {
	OperatorId       *big.Int
	TargetOperatorId *big.Int
	Ender            common.Address
	Amounts          *big.Int
	AmountOut        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEthUnstake is a free log retrieval operation binding the contract event 0xc2d18d1ab67a48ae80c3ef1d20c2f2a97201a23db7ca49e5de1edf05610fb003.
//
// Solidity: event EthUnstake(uint256 indexed _operatorId, uint256 targetOperatorId, address ender, uint256 _amounts, uint256 amountOut)
func (_Liq *LiqFilterer) FilterEthUnstake(opts *bind.FilterOpts, _operatorId []*big.Int) (*LiqEthUnstakeIterator, error) {

	var _operatorIdRule []interface{}
	for _, _operatorIdItem := range _operatorId {
		_operatorIdRule = append(_operatorIdRule, _operatorIdItem)
	}

	logs, sub, err := _Liq.contract.FilterLogs(opts, "EthUnstake", _operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &LiqEthUnstakeIterator{contract: _Liq.contract, event: "EthUnstake", logs: logs, sub: sub}, nil
}

// WatchEthUnstake is a free log subscription operation binding the contract event 0xc2d18d1ab67a48ae80c3ef1d20c2f2a97201a23db7ca49e5de1edf05610fb003.
//
// Solidity: event EthUnstake(uint256 indexed _operatorId, uint256 targetOperatorId, address ender, uint256 _amounts, uint256 amountOut)
func (_Liq *LiqFilterer) WatchEthUnstake(opts *bind.WatchOpts, sink chan<- *LiqEthUnstake, _operatorId []*big.Int) (event.Subscription, error) {

	var _operatorIdRule []interface{}
	for _, _operatorIdItem := range _operatorId {
		_operatorIdRule = append(_operatorIdRule, _operatorIdItem)
	}

	logs, sub, err := _Liq.contract.WatchLogs(opts, "EthUnstake", _operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqEthUnstake)
				if err := _Liq.contract.UnpackLog(event, "EthUnstake", log); err != nil {
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

// ParseEthUnstake is a log parse operation binding the contract event 0xc2d18d1ab67a48ae80c3ef1d20c2f2a97201a23db7ca49e5de1edf05610fb003.
//
// Solidity: event EthUnstake(uint256 indexed _operatorId, uint256 targetOperatorId, address ender, uint256 _amounts, uint256 amountOut)
func (_Liq *LiqFilterer) ParseEthUnstake(log types.Log) (*LiqEthUnstake, error) {
	event := new(LiqEthUnstake)
	if err := _Liq.contract.UnpackLog(event, "EthUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Liq contract.
type LiqInitializedIterator struct {
	Event *LiqInitialized // Event containing the contract specifics and raw log

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
func (it *LiqInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqInitialized)
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
		it.Event = new(LiqInitialized)
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
func (it *LiqInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqInitialized represents a Initialized event raised by the Liq contract.
type LiqInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Liq *LiqFilterer) FilterInitialized(opts *bind.FilterOpts) (*LiqInitializedIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LiqInitializedIterator{contract: _Liq.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Liq *LiqFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LiqInitialized) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqInitialized)
				if err := _Liq.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Liq *LiqFilterer) ParseInitialized(log types.Log) (*LiqInitialized, error) {
	event := new(LiqInitialized)
	if err := _Liq.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqLiquidStakingWithdrawalCredentialsSetIterator is returned from FilterLiquidStakingWithdrawalCredentialsSet and is used to iterate over the raw logs and unpacked data for LiquidStakingWithdrawalCredentialsSet events raised by the Liq contract.
type LiqLiquidStakingWithdrawalCredentialsSetIterator struct {
	Event *LiqLiquidStakingWithdrawalCredentialsSet // Event containing the contract specifics and raw log

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
func (it *LiqLiquidStakingWithdrawalCredentialsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqLiquidStakingWithdrawalCredentialsSet)
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
		it.Event = new(LiqLiquidStakingWithdrawalCredentialsSet)
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
func (it *LiqLiquidStakingWithdrawalCredentialsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqLiquidStakingWithdrawalCredentialsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqLiquidStakingWithdrawalCredentialsSet represents a LiquidStakingWithdrawalCredentialsSet event raised by the Liq contract.
type LiqLiquidStakingWithdrawalCredentialsSet struct {
	OldLiquidStakingWithdrawalCredentials []byte
	LiquidStakingWithdrawalCredentials    []byte
	Raw                                   types.Log // Blockchain specific contextual infos
}

// FilterLiquidStakingWithdrawalCredentialsSet is a free log retrieval operation binding the contract event 0x3af3c6f7639df17de134ac1099e1e1d389efdc16cda06cbfbc76142258d5d2fd.
//
// Solidity: event LiquidStakingWithdrawalCredentialsSet(bytes _oldLiquidStakingWithdrawalCredentials, bytes _liquidStakingWithdrawalCredentials)
func (_Liq *LiqFilterer) FilterLiquidStakingWithdrawalCredentialsSet(opts *bind.FilterOpts) (*LiqLiquidStakingWithdrawalCredentialsSetIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "LiquidStakingWithdrawalCredentialsSet")
	if err != nil {
		return nil, err
	}
	return &LiqLiquidStakingWithdrawalCredentialsSetIterator{contract: _Liq.contract, event: "LiquidStakingWithdrawalCredentialsSet", logs: logs, sub: sub}, nil
}

// WatchLiquidStakingWithdrawalCredentialsSet is a free log subscription operation binding the contract event 0x3af3c6f7639df17de134ac1099e1e1d389efdc16cda06cbfbc76142258d5d2fd.
//
// Solidity: event LiquidStakingWithdrawalCredentialsSet(bytes _oldLiquidStakingWithdrawalCredentials, bytes _liquidStakingWithdrawalCredentials)
func (_Liq *LiqFilterer) WatchLiquidStakingWithdrawalCredentialsSet(opts *bind.WatchOpts, sink chan<- *LiqLiquidStakingWithdrawalCredentialsSet) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "LiquidStakingWithdrawalCredentialsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqLiquidStakingWithdrawalCredentialsSet)
				if err := _Liq.contract.UnpackLog(event, "LiquidStakingWithdrawalCredentialsSet", log); err != nil {
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

// ParseLiquidStakingWithdrawalCredentialsSet is a log parse operation binding the contract event 0x3af3c6f7639df17de134ac1099e1e1d389efdc16cda06cbfbc76142258d5d2fd.
//
// Solidity: event LiquidStakingWithdrawalCredentialsSet(bytes _oldLiquidStakingWithdrawalCredentials, bytes _liquidStakingWithdrawalCredentials)
func (_Liq *LiqFilterer) ParseLiquidStakingWithdrawalCredentialsSet(log types.Log) (*LiqLiquidStakingWithdrawalCredentialsSet, error) {
	event := new(LiqLiquidStakingWithdrawalCredentialsSet)
	if err := _Liq.contract.UnpackLog(event, "LiquidStakingWithdrawalCredentialsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqNftExitBlockNumberSetIterator is returned from FilterNftExitBlockNumberSet and is used to iterate over the raw logs and unpacked data for NftExitBlockNumberSet events raised by the Liq contract.
type LiqNftExitBlockNumberSetIterator struct {
	Event *LiqNftExitBlockNumberSet // Event containing the contract specifics and raw log

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
func (it *LiqNftExitBlockNumberSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqNftExitBlockNumberSet)
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
		it.Event = new(LiqNftExitBlockNumberSet)
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
func (it *LiqNftExitBlockNumberSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqNftExitBlockNumberSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqNftExitBlockNumberSet represents a NftExitBlockNumberSet event raised by the Liq contract.
type LiqNftExitBlockNumberSet struct {
	TokenIds         []*big.Int
	ExitBlockNumbers []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNftExitBlockNumberSet is a free log retrieval operation binding the contract event 0x20bcb7d77186ac511956250f550408a100cc464af62c11e5b60c40aa3ec9c42e.
//
// Solidity: event NftExitBlockNumberSet(uint256[] tokenIds, uint256[] exitBlockNumbers)
func (_Liq *LiqFilterer) FilterNftExitBlockNumberSet(opts *bind.FilterOpts) (*LiqNftExitBlockNumberSetIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "NftExitBlockNumberSet")
	if err != nil {
		return nil, err
	}
	return &LiqNftExitBlockNumberSetIterator{contract: _Liq.contract, event: "NftExitBlockNumberSet", logs: logs, sub: sub}, nil
}

// WatchNftExitBlockNumberSet is a free log subscription operation binding the contract event 0x20bcb7d77186ac511956250f550408a100cc464af62c11e5b60c40aa3ec9c42e.
//
// Solidity: event NftExitBlockNumberSet(uint256[] tokenIds, uint256[] exitBlockNumbers)
func (_Liq *LiqFilterer) WatchNftExitBlockNumberSet(opts *bind.WatchOpts, sink chan<- *LiqNftExitBlockNumberSet) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "NftExitBlockNumberSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqNftExitBlockNumberSet)
				if err := _Liq.contract.UnpackLog(event, "NftExitBlockNumberSet", log); err != nil {
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

// ParseNftExitBlockNumberSet is a log parse operation binding the contract event 0x20bcb7d77186ac511956250f550408a100cc464af62c11e5b60c40aa3ec9c42e.
//
// Solidity: event NftExitBlockNumberSet(uint256[] tokenIds, uint256[] exitBlockNumbers)
func (_Liq *LiqFilterer) ParseNftExitBlockNumberSet(log types.Log) (*LiqNftExitBlockNumberSet, error) {
	event := new(LiqNftExitBlockNumberSet)
	if err := _Liq.contract.UnpackLog(event, "NftExitBlockNumberSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqNftStakeIterator is returned from FilterNftStake and is used to iterate over the raw logs and unpacked data for NftStake events raised by the Liq contract.
type LiqNftStakeIterator struct {
	Event *LiqNftStake // Event containing the contract specifics and raw log

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
func (it *LiqNftStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqNftStake)
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
		it.Event = new(LiqNftStake)
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
func (it *LiqNftStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqNftStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqNftStake represents a NftStake event raised by the Liq contract.
type LiqNftStake struct {
	OperatorId *big.Int
	From       common.Address
	Count      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNftStake is a free log retrieval operation binding the contract event 0x7acd64380f0b80d169f87cad8295b93ee2889c7f5c8f861bee0d1e4edb657afb.
//
// Solidity: event NftStake(uint256 indexed _operatorId, address indexed _from, uint256 _count)
func (_Liq *LiqFilterer) FilterNftStake(opts *bind.FilterOpts, _operatorId []*big.Int, _from []common.Address) (*LiqNftStakeIterator, error) {

	var _operatorIdRule []interface{}
	for _, _operatorIdItem := range _operatorId {
		_operatorIdRule = append(_operatorIdRule, _operatorIdItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Liq.contract.FilterLogs(opts, "NftStake", _operatorIdRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &LiqNftStakeIterator{contract: _Liq.contract, event: "NftStake", logs: logs, sub: sub}, nil
}

// WatchNftStake is a free log subscription operation binding the contract event 0x7acd64380f0b80d169f87cad8295b93ee2889c7f5c8f861bee0d1e4edb657afb.
//
// Solidity: event NftStake(uint256 indexed _operatorId, address indexed _from, uint256 _count)
func (_Liq *LiqFilterer) WatchNftStake(opts *bind.WatchOpts, sink chan<- *LiqNftStake, _operatorId []*big.Int, _from []common.Address) (event.Subscription, error) {

	var _operatorIdRule []interface{}
	for _, _operatorIdItem := range _operatorId {
		_operatorIdRule = append(_operatorIdRule, _operatorIdItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Liq.contract.WatchLogs(opts, "NftStake", _operatorIdRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqNftStake)
				if err := _Liq.contract.UnpackLog(event, "NftStake", log); err != nil {
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

// ParseNftStake is a log parse operation binding the contract event 0x7acd64380f0b80d169f87cad8295b93ee2889c7f5c8f861bee0d1e4edb657afb.
//
// Solidity: event NftStake(uint256 indexed _operatorId, address indexed _from, uint256 _count)
func (_Liq *LiqFilterer) ParseNftStake(log types.Log) (*LiqNftStake, error) {
	event := new(LiqNftStake)
	if err := _Liq.contract.UnpackLog(event, "NftStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqNodeOperatorRegistryContractSetIterator is returned from FilterNodeOperatorRegistryContractSet and is used to iterate over the raw logs and unpacked data for NodeOperatorRegistryContractSet events raised by the Liq contract.
type LiqNodeOperatorRegistryContractSetIterator struct {
	Event *LiqNodeOperatorRegistryContractSet // Event containing the contract specifics and raw log

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
func (it *LiqNodeOperatorRegistryContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqNodeOperatorRegistryContractSet)
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
		it.Event = new(LiqNodeOperatorRegistryContractSet)
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
func (it *LiqNodeOperatorRegistryContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqNodeOperatorRegistryContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqNodeOperatorRegistryContractSet represents a NodeOperatorRegistryContractSet event raised by the Liq contract.
type LiqNodeOperatorRegistryContractSet struct {
	OldNodeOperatorRegistryContract common.Address
	NodeOperatorRegistryContract    common.Address
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorRegistryContractSet is a free log retrieval operation binding the contract event 0x2aa578b9d95064e7e90ca0af5e42ca5499f5e90bd32c4e401df52a686ac6993d.
//
// Solidity: event NodeOperatorRegistryContractSet(address _oldNodeOperatorRegistryContract, address _nodeOperatorRegistryContract)
func (_Liq *LiqFilterer) FilterNodeOperatorRegistryContractSet(opts *bind.FilterOpts) (*LiqNodeOperatorRegistryContractSetIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "NodeOperatorRegistryContractSet")
	if err != nil {
		return nil, err
	}
	return &LiqNodeOperatorRegistryContractSetIterator{contract: _Liq.contract, event: "NodeOperatorRegistryContractSet", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorRegistryContractSet is a free log subscription operation binding the contract event 0x2aa578b9d95064e7e90ca0af5e42ca5499f5e90bd32c4e401df52a686ac6993d.
//
// Solidity: event NodeOperatorRegistryContractSet(address _oldNodeOperatorRegistryContract, address _nodeOperatorRegistryContract)
func (_Liq *LiqFilterer) WatchNodeOperatorRegistryContractSet(opts *bind.WatchOpts, sink chan<- *LiqNodeOperatorRegistryContractSet) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "NodeOperatorRegistryContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqNodeOperatorRegistryContractSet)
				if err := _Liq.contract.UnpackLog(event, "NodeOperatorRegistryContractSet", log); err != nil {
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
func (_Liq *LiqFilterer) ParseNodeOperatorRegistryContractSet(log types.Log) (*LiqNodeOperatorRegistryContractSet, error) {
	event := new(LiqNodeOperatorRegistryContractSet)
	if err := _Liq.contract.UnpackLog(event, "NodeOperatorRegistryContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqOperatorAssignedIterator is returned from FilterOperatorAssigned and is used to iterate over the raw logs and unpacked data for OperatorAssigned events raised by the Liq contract.
type LiqOperatorAssignedIterator struct {
	Event *LiqOperatorAssigned // Event containing the contract specifics and raw log

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
func (it *LiqOperatorAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqOperatorAssigned)
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
		it.Event = new(LiqOperatorAssigned)
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
func (it *LiqOperatorAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqOperatorAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqOperatorAssigned represents a OperatorAssigned event raised by the Liq contract.
type LiqOperatorAssigned struct {
	BlacklistOperatorId *big.Int
	OperatorId          *big.Int
	TotalAmount         *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterOperatorAssigned is a free log retrieval operation binding the contract event 0x2fb596e064755189ea64f65d467239f889fbc89a0cc5f7c4ac32cea82547919b.
//
// Solidity: event OperatorAssigned(uint256 indexed _blacklistOperatorId, uint256 _operatorId, uint256 _totalAmount)
func (_Liq *LiqFilterer) FilterOperatorAssigned(opts *bind.FilterOpts, _blacklistOperatorId []*big.Int) (*LiqOperatorAssignedIterator, error) {

	var _blacklistOperatorIdRule []interface{}
	for _, _blacklistOperatorIdItem := range _blacklistOperatorId {
		_blacklistOperatorIdRule = append(_blacklistOperatorIdRule, _blacklistOperatorIdItem)
	}

	logs, sub, err := _Liq.contract.FilterLogs(opts, "OperatorAssigned", _blacklistOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &LiqOperatorAssignedIterator{contract: _Liq.contract, event: "OperatorAssigned", logs: logs, sub: sub}, nil
}

// WatchOperatorAssigned is a free log subscription operation binding the contract event 0x2fb596e064755189ea64f65d467239f889fbc89a0cc5f7c4ac32cea82547919b.
//
// Solidity: event OperatorAssigned(uint256 indexed _blacklistOperatorId, uint256 _operatorId, uint256 _totalAmount)
func (_Liq *LiqFilterer) WatchOperatorAssigned(opts *bind.WatchOpts, sink chan<- *LiqOperatorAssigned, _blacklistOperatorId []*big.Int) (event.Subscription, error) {

	var _blacklistOperatorIdRule []interface{}
	for _, _blacklistOperatorIdItem := range _blacklistOperatorId {
		_blacklistOperatorIdRule = append(_blacklistOperatorIdRule, _blacklistOperatorIdItem)
	}

	logs, sub, err := _Liq.contract.WatchLogs(opts, "OperatorAssigned", _blacklistOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqOperatorAssigned)
				if err := _Liq.contract.UnpackLog(event, "OperatorAssigned", log); err != nil {
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

// ParseOperatorAssigned is a log parse operation binding the contract event 0x2fb596e064755189ea64f65d467239f889fbc89a0cc5f7c4ac32cea82547919b.
//
// Solidity: event OperatorAssigned(uint256 indexed _blacklistOperatorId, uint256 _operatorId, uint256 _totalAmount)
func (_Liq *LiqFilterer) ParseOperatorAssigned(log types.Log) (*LiqOperatorAssigned, error) {
	event := new(LiqOperatorAssigned)
	if err := _Liq.contract.UnpackLog(event, "OperatorAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqOperatorCanLoanAmountsSetIterator is returned from FilterOperatorCanLoanAmountsSet and is used to iterate over the raw logs and unpacked data for OperatorCanLoanAmountsSet events raised by the Liq contract.
type LiqOperatorCanLoanAmountsSetIterator struct {
	Event *LiqOperatorCanLoanAmountsSet // Event containing the contract specifics and raw log

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
func (it *LiqOperatorCanLoanAmountsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqOperatorCanLoanAmountsSet)
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
		it.Event = new(LiqOperatorCanLoanAmountsSet)
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
func (it *LiqOperatorCanLoanAmountsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqOperatorCanLoanAmountsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqOperatorCanLoanAmountsSet represents a OperatorCanLoanAmountsSet event raised by the Liq contract.
type LiqOperatorCanLoanAmountsSet struct {
	OperatorCanLoanAmounts *big.Int
	NewCanloadAmounts      *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterOperatorCanLoanAmountsSet is a free log retrieval operation binding the contract event 0xe4ef8288470b9f1af2ad6df80093b6a682f279693121712af0247f814f2f31a4.
//
// Solidity: event OperatorCanLoanAmountsSet(uint256 operatorCanLoanAmounts, uint256 _newCanloadAmounts)
func (_Liq *LiqFilterer) FilterOperatorCanLoanAmountsSet(opts *bind.FilterOpts) (*LiqOperatorCanLoanAmountsSetIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "OperatorCanLoanAmountsSet")
	if err != nil {
		return nil, err
	}
	return &LiqOperatorCanLoanAmountsSetIterator{contract: _Liq.contract, event: "OperatorCanLoanAmountsSet", logs: logs, sub: sub}, nil
}

// WatchOperatorCanLoanAmountsSet is a free log subscription operation binding the contract event 0xe4ef8288470b9f1af2ad6df80093b6a682f279693121712af0247f814f2f31a4.
//
// Solidity: event OperatorCanLoanAmountsSet(uint256 operatorCanLoanAmounts, uint256 _newCanloadAmounts)
func (_Liq *LiqFilterer) WatchOperatorCanLoanAmountsSet(opts *bind.WatchOpts, sink chan<- *LiqOperatorCanLoanAmountsSet) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "OperatorCanLoanAmountsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqOperatorCanLoanAmountsSet)
				if err := _Liq.contract.UnpackLog(event, "OperatorCanLoanAmountsSet", log); err != nil {
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

// ParseOperatorCanLoanAmountsSet is a log parse operation binding the contract event 0xe4ef8288470b9f1af2ad6df80093b6a682f279693121712af0247f814f2f31a4.
//
// Solidity: event OperatorCanLoanAmountsSet(uint256 operatorCanLoanAmounts, uint256 _newCanloadAmounts)
func (_Liq *LiqFilterer) ParseOperatorCanLoanAmountsSet(log types.Log) (*LiqOperatorCanLoanAmountsSet, error) {
	event := new(LiqOperatorCanLoanAmountsSet)
	if err := _Liq.contract.UnpackLog(event, "OperatorCanLoanAmountsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqOperatorClaimRewardsIterator is returned from FilterOperatorClaimRewards and is used to iterate over the raw logs and unpacked data for OperatorClaimRewards events raised by the Liq contract.
type LiqOperatorClaimRewardsIterator struct {
	Event *LiqOperatorClaimRewards // Event containing the contract specifics and raw log

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
func (it *LiqOperatorClaimRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqOperatorClaimRewards)
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
		it.Event = new(LiqOperatorClaimRewards)
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
func (it *LiqOperatorClaimRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqOperatorClaimRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqOperatorClaimRewards represents a OperatorClaimRewards event raised by the Liq contract.
type LiqOperatorClaimRewards struct {
	OperatorId *big.Int
	Rewards    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorClaimRewards is a free log retrieval operation binding the contract event 0x9af81b3bd245e1f5301f53c691f532ada01ad22ed150ecf59d65eebf084a635f.
//
// Solidity: event OperatorClaimRewards(uint256 _operatorId, uint256 _rewards)
func (_Liq *LiqFilterer) FilterOperatorClaimRewards(opts *bind.FilterOpts) (*LiqOperatorClaimRewardsIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "OperatorClaimRewards")
	if err != nil {
		return nil, err
	}
	return &LiqOperatorClaimRewardsIterator{contract: _Liq.contract, event: "OperatorClaimRewards", logs: logs, sub: sub}, nil
}

// WatchOperatorClaimRewards is a free log subscription operation binding the contract event 0x9af81b3bd245e1f5301f53c691f532ada01ad22ed150ecf59d65eebf084a635f.
//
// Solidity: event OperatorClaimRewards(uint256 _operatorId, uint256 _rewards)
func (_Liq *LiqFilterer) WatchOperatorClaimRewards(opts *bind.WatchOpts, sink chan<- *LiqOperatorClaimRewards) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "OperatorClaimRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqOperatorClaimRewards)
				if err := _Liq.contract.UnpackLog(event, "OperatorClaimRewards", log); err != nil {
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
func (_Liq *LiqFilterer) ParseOperatorClaimRewards(log types.Log) (*LiqOperatorClaimRewards, error) {
	event := new(LiqOperatorClaimRewards)
	if err := _Liq.contract.UnpackLog(event, "OperatorClaimRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqOperatorReinvestClRewardsIterator is returned from FilterOperatorReinvestClRewards and is used to iterate over the raw logs and unpacked data for OperatorReinvestClRewards events raised by the Liq contract.
type LiqOperatorReinvestClRewardsIterator struct {
	Event *LiqOperatorReinvestClRewards // Event containing the contract specifics and raw log

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
func (it *LiqOperatorReinvestClRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqOperatorReinvestClRewards)
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
		it.Event = new(LiqOperatorReinvestClRewards)
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
func (it *LiqOperatorReinvestClRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqOperatorReinvestClRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqOperatorReinvestClRewards represents a OperatorReinvestClRewards event raised by the Liq contract.
type LiqOperatorReinvestClRewards struct {
	OperatorId *big.Int
	Rewards    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorReinvestClRewards is a free log retrieval operation binding the contract event 0x7e0fd5b648afd09ad3f3913bdf766931ac752a7ac7c392b7713362379923b0fe.
//
// Solidity: event OperatorReinvestClRewards(uint256 _operatorId, uint256 _rewards)
func (_Liq *LiqFilterer) FilterOperatorReinvestClRewards(opts *bind.FilterOpts) (*LiqOperatorReinvestClRewardsIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "OperatorReinvestClRewards")
	if err != nil {
		return nil, err
	}
	return &LiqOperatorReinvestClRewardsIterator{contract: _Liq.contract, event: "OperatorReinvestClRewards", logs: logs, sub: sub}, nil
}

// WatchOperatorReinvestClRewards is a free log subscription operation binding the contract event 0x7e0fd5b648afd09ad3f3913bdf766931ac752a7ac7c392b7713362379923b0fe.
//
// Solidity: event OperatorReinvestClRewards(uint256 _operatorId, uint256 _rewards)
func (_Liq *LiqFilterer) WatchOperatorReinvestClRewards(opts *bind.WatchOpts, sink chan<- *LiqOperatorReinvestClRewards) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "OperatorReinvestClRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqOperatorReinvestClRewards)
				if err := _Liq.contract.UnpackLog(event, "OperatorReinvestClRewards", log); err != nil {
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

// ParseOperatorReinvestClRewards is a log parse operation binding the contract event 0x7e0fd5b648afd09ad3f3913bdf766931ac752a7ac7c392b7713362379923b0fe.
//
// Solidity: event OperatorReinvestClRewards(uint256 _operatorId, uint256 _rewards)
func (_Liq *LiqFilterer) ParseOperatorReinvestClRewards(log types.Log) (*LiqOperatorReinvestClRewards, error) {
	event := new(LiqOperatorReinvestClRewards)
	if err := _Liq.contract.UnpackLog(event, "OperatorReinvestClRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqOperatorReinvestElRewardsIterator is returned from FilterOperatorReinvestElRewards and is used to iterate over the raw logs and unpacked data for OperatorReinvestElRewards events raised by the Liq contract.
type LiqOperatorReinvestElRewardsIterator struct {
	Event *LiqOperatorReinvestElRewards // Event containing the contract specifics and raw log

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
func (it *LiqOperatorReinvestElRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqOperatorReinvestElRewards)
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
		it.Event = new(LiqOperatorReinvestElRewards)
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
func (it *LiqOperatorReinvestElRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqOperatorReinvestElRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqOperatorReinvestElRewards represents a OperatorReinvestElRewards event raised by the Liq contract.
type LiqOperatorReinvestElRewards struct {
	OperatorId *big.Int
	Rewards    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorReinvestElRewards is a free log retrieval operation binding the contract event 0x933f5d3652cfb50d9224df3ff1b228450d2917a80b7338cc68a309f5b20c38f6.
//
// Solidity: event OperatorReinvestElRewards(uint256 _operatorId, uint256 _rewards)
func (_Liq *LiqFilterer) FilterOperatorReinvestElRewards(opts *bind.FilterOpts) (*LiqOperatorReinvestElRewardsIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "OperatorReinvestElRewards")
	if err != nil {
		return nil, err
	}
	return &LiqOperatorReinvestElRewardsIterator{contract: _Liq.contract, event: "OperatorReinvestElRewards", logs: logs, sub: sub}, nil
}

// WatchOperatorReinvestElRewards is a free log subscription operation binding the contract event 0x933f5d3652cfb50d9224df3ff1b228450d2917a80b7338cc68a309f5b20c38f6.
//
// Solidity: event OperatorReinvestElRewards(uint256 _operatorId, uint256 _rewards)
func (_Liq *LiqFilterer) WatchOperatorReinvestElRewards(opts *bind.WatchOpts, sink chan<- *LiqOperatorReinvestElRewards) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "OperatorReinvestElRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqOperatorReinvestElRewards)
				if err := _Liq.contract.UnpackLog(event, "OperatorReinvestElRewards", log); err != nil {
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

// ParseOperatorReinvestElRewards is a log parse operation binding the contract event 0x933f5d3652cfb50d9224df3ff1b228450d2917a80b7338cc68a309f5b20c38f6.
//
// Solidity: event OperatorReinvestElRewards(uint256 _operatorId, uint256 _rewards)
func (_Liq *LiqFilterer) ParseOperatorReinvestElRewards(log types.Log) (*LiqOperatorReinvestElRewards, error) {
	event := new(LiqOperatorReinvestElRewards)
	if err := _Liq.contract.UnpackLog(event, "OperatorReinvestElRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Liq contract.
type LiqOwnershipTransferredIterator struct {
	Event *LiqOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LiqOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqOwnershipTransferred)
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
		it.Event = new(LiqOwnershipTransferred)
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
func (it *LiqOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqOwnershipTransferred represents a OwnershipTransferred event raised by the Liq contract.
type LiqOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Liq *LiqFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LiqOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Liq.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LiqOwnershipTransferredIterator{contract: _Liq.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Liq *LiqFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LiqOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Liq.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqOwnershipTransferred)
				if err := _Liq.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Liq *LiqFilterer) ParseOwnershipTransferred(log types.Log) (*LiqOwnershipTransferred, error) {
	event := new(LiqOwnershipTransferred)
	if err := _Liq.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Liq contract.
type LiqPausedIterator struct {
	Event *LiqPaused // Event containing the contract specifics and raw log

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
func (it *LiqPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqPaused)
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
		it.Event = new(LiqPaused)
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
func (it *LiqPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqPaused represents a Paused event raised by the Liq contract.
type LiqPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Liq *LiqFilterer) FilterPaused(opts *bind.FilterOpts) (*LiqPausedIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LiqPausedIterator{contract: _Liq.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Liq *LiqFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LiqPaused) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqPaused)
				if err := _Liq.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Liq *LiqFilterer) ParsePaused(log types.Log) (*LiqPaused, error) {
	event := new(LiqPaused)
	if err := _Liq.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqRewardsReceiveIterator is returned from FilterRewardsReceive and is used to iterate over the raw logs and unpacked data for RewardsReceive events raised by the Liq contract.
type LiqRewardsReceiveIterator struct {
	Event *LiqRewardsReceive // Event containing the contract specifics and raw log

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
func (it *LiqRewardsReceiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqRewardsReceive)
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
		it.Event = new(LiqRewardsReceive)
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
func (it *LiqRewardsReceiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqRewardsReceiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqRewardsReceive represents a RewardsReceive event raised by the Liq contract.
type LiqRewardsReceive struct {
	Rewards *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardsReceive is a free log retrieval operation binding the contract event 0x128bc67045828b9e917cda6b9717921b174e8e2d1f135074c8d45e41d4e69ec6.
//
// Solidity: event RewardsReceive(uint256 _rewards)
func (_Liq *LiqFilterer) FilterRewardsReceive(opts *bind.FilterOpts) (*LiqRewardsReceiveIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "RewardsReceive")
	if err != nil {
		return nil, err
	}
	return &LiqRewardsReceiveIterator{contract: _Liq.contract, event: "RewardsReceive", logs: logs, sub: sub}, nil
}

// WatchRewardsReceive is a free log subscription operation binding the contract event 0x128bc67045828b9e917cda6b9717921b174e8e2d1f135074c8d45e41d4e69ec6.
//
// Solidity: event RewardsReceive(uint256 _rewards)
func (_Liq *LiqFilterer) WatchRewardsReceive(opts *bind.WatchOpts, sink chan<- *LiqRewardsReceive) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "RewardsReceive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqRewardsReceive)
				if err := _Liq.contract.UnpackLog(event, "RewardsReceive", log); err != nil {
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

// ParseRewardsReceive is a log parse operation binding the contract event 0x128bc67045828b9e917cda6b9717921b174e8e2d1f135074c8d45e41d4e69ec6.
//
// Solidity: event RewardsReceive(uint256 _rewards)
func (_Liq *LiqFilterer) ParseRewardsReceive(log types.Log) (*LiqRewardsReceive, error) {
	event := new(LiqRewardsReceive)
	if err := _Liq.contract.UnpackLog(event, "RewardsReceive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqTransferredIterator is returned from FilterTransferred and is used to iterate over the raw logs and unpacked data for Transferred events raised by the Liq contract.
type LiqTransferredIterator struct {
	Event *LiqTransferred // Event containing the contract specifics and raw log

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
func (it *LiqTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqTransferred)
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
		it.Event = new(LiqTransferred)
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
func (it *LiqTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqTransferred represents a Transferred event raised by the Liq contract.
type LiqTransferred struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferred is a free log retrieval operation binding the contract event 0xe6d858f14d755446648a6e0c8ab8b5a0f58ccc7920d4c910b0454e4dcd869af0.
//
// Solidity: event Transferred(address _to, uint256 _amount)
func (_Liq *LiqFilterer) FilterTransferred(opts *bind.FilterOpts) (*LiqTransferredIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "Transferred")
	if err != nil {
		return nil, err
	}
	return &LiqTransferredIterator{contract: _Liq.contract, event: "Transferred", logs: logs, sub: sub}, nil
}

// WatchTransferred is a free log subscription operation binding the contract event 0xe6d858f14d755446648a6e0c8ab8b5a0f58ccc7920d4c910b0454e4dcd869af0.
//
// Solidity: event Transferred(address _to, uint256 _amount)
func (_Liq *LiqFilterer) WatchTransferred(opts *bind.WatchOpts, sink chan<- *LiqTransferred) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "Transferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqTransferred)
				if err := _Liq.contract.UnpackLog(event, "Transferred", log); err != nil {
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
func (_Liq *LiqFilterer) ParseTransferred(log types.Log) (*LiqTransferred, error) {
	event := new(LiqTransferred)
	if err := _Liq.contract.UnpackLog(event, "Transferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Liq contract.
type LiqUnpausedIterator struct {
	Event *LiqUnpaused // Event containing the contract specifics and raw log

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
func (it *LiqUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqUnpaused)
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
		it.Event = new(LiqUnpaused)
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
func (it *LiqUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqUnpaused represents a Unpaused event raised by the Liq contract.
type LiqUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Liq *LiqFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LiqUnpausedIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LiqUnpausedIterator{contract: _Liq.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Liq *LiqFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LiqUnpaused) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqUnpaused)
				if err := _Liq.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Liq *LiqFilterer) ParseUnpaused(log types.Log) (*LiqUnpaused, error) {
	event := new(LiqUnpaused)
	if err := _Liq.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Liq contract.
type LiqUpgradedIterator struct {
	Event *LiqUpgraded // Event containing the contract specifics and raw log

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
func (it *LiqUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqUpgraded)
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
		it.Event = new(LiqUpgraded)
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
func (it *LiqUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqUpgraded represents a Upgraded event raised by the Liq contract.
type LiqUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Liq *LiqFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LiqUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Liq.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LiqUpgradedIterator{contract: _Liq.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Liq *LiqFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LiqUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Liq.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqUpgraded)
				if err := _Liq.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Liq *LiqFilterer) ParseUpgraded(log types.Log) (*LiqUpgraded, error) {
	event := new(LiqUpgraded)
	if err := _Liq.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqUserClaimRewardsIterator is returned from FilterUserClaimRewards and is used to iterate over the raw logs and unpacked data for UserClaimRewards events raised by the Liq contract.
type LiqUserClaimRewardsIterator struct {
	Event *LiqUserClaimRewards // Event containing the contract specifics and raw log

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
func (it *LiqUserClaimRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqUserClaimRewards)
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
		it.Event = new(LiqUserClaimRewards)
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
func (it *LiqUserClaimRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqUserClaimRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqUserClaimRewards represents a UserClaimRewards event raised by the Liq contract.
type LiqUserClaimRewards struct {
	OperatorId *big.Int
	TokenIds   []*big.Int
	Rewards    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUserClaimRewards is a free log retrieval operation binding the contract event 0xe1419955066833fcc44409a448a939482f4bdc3af35be67c59740ea729ff4a0b.
//
// Solidity: event UserClaimRewards(uint256 _operatorId, uint256[] _tokenIds, uint256 _rewards)
func (_Liq *LiqFilterer) FilterUserClaimRewards(opts *bind.FilterOpts) (*LiqUserClaimRewardsIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "UserClaimRewards")
	if err != nil {
		return nil, err
	}
	return &LiqUserClaimRewardsIterator{contract: _Liq.contract, event: "UserClaimRewards", logs: logs, sub: sub}, nil
}

// WatchUserClaimRewards is a free log subscription operation binding the contract event 0xe1419955066833fcc44409a448a939482f4bdc3af35be67c59740ea729ff4a0b.
//
// Solidity: event UserClaimRewards(uint256 _operatorId, uint256[] _tokenIds, uint256 _rewards)
func (_Liq *LiqFilterer) WatchUserClaimRewards(opts *bind.WatchOpts, sink chan<- *LiqUserClaimRewards) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "UserClaimRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqUserClaimRewards)
				if err := _Liq.contract.UnpackLog(event, "UserClaimRewards", log); err != nil {
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

// ParseUserClaimRewards is a log parse operation binding the contract event 0xe1419955066833fcc44409a448a939482f4bdc3af35be67c59740ea729ff4a0b.
//
// Solidity: event UserClaimRewards(uint256 _operatorId, uint256[] _tokenIds, uint256 _rewards)
func (_Liq *LiqFilterer) ParseUserClaimRewards(log types.Log) (*LiqUserClaimRewards, error) {
	event := new(LiqUserClaimRewards)
	if err := _Liq.contract.UnpackLog(event, "UserClaimRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the Liq contract.
type LiqValidatorRegisteredIterator struct {
	Event *LiqValidatorRegistered // Event containing the contract specifics and raw log

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
func (it *LiqValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqValidatorRegistered)
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
		it.Event = new(LiqValidatorRegistered)
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
func (it *LiqValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqValidatorRegistered represents a ValidatorRegistered event raised by the Liq contract.
type LiqValidatorRegistered struct {
	OperatorId *big.Int
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0x700aef92048a2622a7b2403606dfa5c7abaec58382232cdd9ec196907eb44396.
//
// Solidity: event ValidatorRegistered(uint256 indexed _operatorId, uint256 _tokenId)
func (_Liq *LiqFilterer) FilterValidatorRegistered(opts *bind.FilterOpts, _operatorId []*big.Int) (*LiqValidatorRegisteredIterator, error) {

	var _operatorIdRule []interface{}
	for _, _operatorIdItem := range _operatorId {
		_operatorIdRule = append(_operatorIdRule, _operatorIdItem)
	}

	logs, sub, err := _Liq.contract.FilterLogs(opts, "ValidatorRegistered", _operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &LiqValidatorRegisteredIterator{contract: _Liq.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0x700aef92048a2622a7b2403606dfa5c7abaec58382232cdd9ec196907eb44396.
//
// Solidity: event ValidatorRegistered(uint256 indexed _operatorId, uint256 _tokenId)
func (_Liq *LiqFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *LiqValidatorRegistered, _operatorId []*big.Int) (event.Subscription, error) {

	var _operatorIdRule []interface{}
	for _, _operatorIdItem := range _operatorId {
		_operatorIdRule = append(_operatorIdRule, _operatorIdItem)
	}

	logs, sub, err := _Liq.contract.WatchLogs(opts, "ValidatorRegistered", _operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqValidatorRegistered)
				if err := _Liq.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
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

// ParseValidatorRegistered is a log parse operation binding the contract event 0x700aef92048a2622a7b2403606dfa5c7abaec58382232cdd9ec196907eb44396.
//
// Solidity: event ValidatorRegistered(uint256 indexed _operatorId, uint256 _tokenId)
func (_Liq *LiqFilterer) ParseValidatorRegistered(log types.Log) (*LiqValidatorRegistered, error) {
	event := new(LiqValidatorRegistered)
	if err := _Liq.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqVaultManagerContractSetIterator is returned from FilterVaultManagerContractSet and is used to iterate over the raw logs and unpacked data for VaultManagerContractSet events raised by the Liq contract.
type LiqVaultManagerContractSetIterator struct {
	Event *LiqVaultManagerContractSet // Event containing the contract specifics and raw log

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
func (it *LiqVaultManagerContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqVaultManagerContractSet)
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
		it.Event = new(LiqVaultManagerContractSet)
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
func (it *LiqVaultManagerContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqVaultManagerContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqVaultManagerContractSet represents a VaultManagerContractSet event raised by the Liq contract.
type LiqVaultManagerContractSet struct {
	VaultManagerContractAddress common.Address
	VaultManagerContract        common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterVaultManagerContractSet is a free log retrieval operation binding the contract event 0x136260758ef216be6f30b5244361f089faf99890f23864c0a63e2d2def24963f.
//
// Solidity: event VaultManagerContractSet(address vaultManagerContractAddress, address _vaultManagerContract)
func (_Liq *LiqFilterer) FilterVaultManagerContractSet(opts *bind.FilterOpts) (*LiqVaultManagerContractSetIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "VaultManagerContractSet")
	if err != nil {
		return nil, err
	}
	return &LiqVaultManagerContractSetIterator{contract: _Liq.contract, event: "VaultManagerContractSet", logs: logs, sub: sub}, nil
}

// WatchVaultManagerContractSet is a free log subscription operation binding the contract event 0x136260758ef216be6f30b5244361f089faf99890f23864c0a63e2d2def24963f.
//
// Solidity: event VaultManagerContractSet(address vaultManagerContractAddress, address _vaultManagerContract)
func (_Liq *LiqFilterer) WatchVaultManagerContractSet(opts *bind.WatchOpts, sink chan<- *LiqVaultManagerContractSet) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "VaultManagerContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqVaultManagerContractSet)
				if err := _Liq.contract.UnpackLog(event, "VaultManagerContractSet", log); err != nil {
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
// Solidity: event VaultManagerContractSet(address vaultManagerContractAddress, address _vaultManagerContract)
func (_Liq *LiqFilterer) ParseVaultManagerContractSet(log types.Log) (*LiqVaultManagerContractSet, error) {
	event := new(LiqVaultManagerContractSet)
	if err := _Liq.contract.UnpackLog(event, "VaultManagerContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

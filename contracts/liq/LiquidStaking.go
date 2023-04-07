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

// LiquidStakingWithdrawalInfo is an auto generated low-level Go binding around an user-defined struct.
type LiquidStakingWithdrawalInfo struct {
	OperatorId         *big.Int
	WithdrawHeight     *big.Int
	WithdrawNethAmount *big.Int
	WithdrawExchange   *big.Int
	ClaimEthAmount     *big.Int
	Owner              common.Address
	IsClaim            bool
}

// LiqMetaData contains all meta data concerning the Liq contract.
var LiqMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ArrearsReceiveOfSlash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldBeaconOracleContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_beaconOracleContractAddress\",\"type\":\"address\"}],\"name\":\"BeaconOracleContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_blacklistOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"}],\"name\":\"BlacklistOperatorAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldDao\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"DaoAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"DaoClaimRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldDaoVaultAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_daoVaultAddress\",\"type\":\"address\"}],\"name\":\"DaoVaultAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_feeRate\",\"type\":\"uint256\"}],\"name\":\"DepositFeeRateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"}],\"name\":\"EthStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amounts\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"EthUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalNethAmount\",\"type\":\"uint256\"}],\"name\":\"LargeWithdrawalsRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_oldLiquidStakingWithdrawalCredentials\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_liquidStakingWithdrawalCredentials\",\"type\":\"bytes\"}],\"name\":\"LiquidStakingWithdrawalCredentialsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"exitBlockNumbers\",\"type\":\"uint256[]\"}],\"name\":\"NftExitBlockNumberSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"NftStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"}],\"name\":\"NftUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"}],\"name\":\"NftUnwrap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"}],\"name\":\"NftWrap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldNodeOperatorRegistryContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistryContract\",\"type\":\"address\"}],\"name\":\"NodeOperatorRegistryContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"OperatorClaimRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"OperatorReinvestClRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"OperatorReinvestElRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_quitOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"}],\"name\":\"QuitOperatorAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"RewardsReceive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_slashAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_requirAmounts\",\"type\":\"uint256\"}],\"name\":\"SlashReceive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"UserClaimRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ValidatorRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_NETH_WITHDRAWAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_NETH_WITHDRAWAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"assignBlacklistOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_quitOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"assignQuitOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beaconOracleContract\",\"outputs\":[{\"internalType\":\"contractIBeaconOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"requestIds\",\"type\":\"uint256[]\"}],\"name\":\"claimLargeWitdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_operatorIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_rewards\",\"type\":\"uint256[]\"}],\"name\":\"claimRewardsOfDao\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"}],\"name\":\"claimRewardsOfOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasHeight\",\"type\":\"uint256\"}],\"name\":\"claimRewardsOfUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"consensusVaultContract\",\"outputs\":[{\"internalType\":\"contractIConsensusVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedExitSlashStandard\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositContract\",\"outputs\":[{\"internalType\":\"contractIDepositContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nethAmountIn\",\"type\":\"uint256\"}],\"name\":\"getEthOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethAmountIn\",\"type\":\"uint256\"}],\"name\":\"getNethOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalEthValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getUserUnstakeButOperatorNoExitNfs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getWithdrawalOfOperator\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawNethAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawExchange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimEthAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isClaim\",\"type\":\"bool\"}],\"internalType\":\"structLiquidStaking.WithdrawalInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getWithdrawalRequestIdOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_daoVaultAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_withdrawalCreds\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistryContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nETHContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nVNFTContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beaconOracleContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositContractAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"largeExitDelayedSlashRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidStakingWithdrawalCredentials\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nETHContract\",\"outputs\":[{\"internalType\":\"contractINETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftExitDelayedSlashRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_exitBlockNumbers\",\"type\":\"uint256[]\"}],\"name\":\"nftExitHandle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftHasCompensated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftUnstakeBlockNumbers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftWillCompensated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeOperatorRegistryContract\",\"outputs\":[{\"internalType\":\"contractINodeOperatorsRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorCompensatedIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorLoadBlockNumbers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorLoanRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorPendingEthPoolBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorPendingEthRequestAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorPoolBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorSlashArrears\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reAssignRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"receiveRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_operatorIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"reinvestClRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_operatorIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"reinvestElRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"requestLargeWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beaconOracleContractAddress\",\"type\":\"address\"}],\"name\":\"setBeaconOracleContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"setDaoAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_daoVaultAddress\",\"type\":\"address\"}],\"name\":\"setDaoVaultAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeRate\",\"type\":\"uint256\"}],\"name\":\"setDepositFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_liquidStakingWithdrawalCredentials\",\"type\":\"bytes\"}],\"name\":\"setLiquidStakingWithdrawalCredentials\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistryContract\",\"type\":\"address\"}],\"name\":\"setNodeOperatorRegistryContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashAmountPerBlockPerValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"slashArrearsReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_nftExitDelayedTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_largeExitDelayedRequestIds\",\"type\":\"uint256[]\"}],\"name\":\"slashOfExitDelayed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_exitTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"slashOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_exitTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_slashAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_requireAmounts\",\"type\":\"uint256[]\"}],\"name\":\"slashReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"stakeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"withdrawalCredentialsAddress\",\"type\":\"address\"}],\"name\":\"stakeNFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quota\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLockedNethBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amounts\",\"type\":\"uint256\"}],\"name\":\"unstakeETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"unstakeNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vNFTContract\",\"outputs\":[{\"internalType\":\"contractIVNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultManagerContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawalQueues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawNethAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawExchange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimEthAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isClaim\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// MAXNETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x63823474.
//
// Solidity: function MAX_NETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_Liq *LiqCaller) MAXNETHWITHDRAWALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "MAX_NETH_WITHDRAWAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXNETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x63823474.
//
// Solidity: function MAX_NETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_Liq *LiqSession) MAXNETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _Liq.Contract.MAXNETHWITHDRAWALAMOUNT(&_Liq.CallOpts)
}

// MAXNETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x63823474.
//
// Solidity: function MAX_NETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_Liq *LiqCallerSession) MAXNETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _Liq.Contract.MAXNETHWITHDRAWALAMOUNT(&_Liq.CallOpts)
}

// MINNETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x24c5cba3.
//
// Solidity: function MIN_NETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_Liq *LiqCaller) MINNETHWITHDRAWALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "MIN_NETH_WITHDRAWAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINNETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x24c5cba3.
//
// Solidity: function MIN_NETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_Liq *LiqSession) MINNETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _Liq.Contract.MINNETHWITHDRAWALAMOUNT(&_Liq.CallOpts)
}

// MINNETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x24c5cba3.
//
// Solidity: function MIN_NETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_Liq *LiqCallerSession) MINNETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _Liq.Contract.MINNETHWITHDRAWALAMOUNT(&_Liq.CallOpts)
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

// DelayedExitSlashStandard is a free data retrieval call binding the contract method 0xfa2d7ea5.
//
// Solidity: function delayedExitSlashStandard() view returns(uint256)
func (_Liq *LiqCaller) DelayedExitSlashStandard(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "delayedExitSlashStandard")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayedExitSlashStandard is a free data retrieval call binding the contract method 0xfa2d7ea5.
//
// Solidity: function delayedExitSlashStandard() view returns(uint256)
func (_Liq *LiqSession) DelayedExitSlashStandard() (*big.Int, error) {
	return _Liq.Contract.DelayedExitSlashStandard(&_Liq.CallOpts)
}

// DelayedExitSlashStandard is a free data retrieval call binding the contract method 0xfa2d7ea5.
//
// Solidity: function delayedExitSlashStandard() view returns(uint256)
func (_Liq *LiqCallerSession) DelayedExitSlashStandard() (*big.Int, error) {
	return _Liq.Contract.DelayedExitSlashStandard(&_Liq.CallOpts)
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

// GetUserUnstakeButOperatorNoExitNfs is a free data retrieval call binding the contract method 0xf1c28381.
//
// Solidity: function getUserUnstakeButOperatorNoExitNfs(uint256 _operatorId) view returns(uint256[])
func (_Liq *LiqCaller) GetUserUnstakeButOperatorNoExitNfs(opts *bind.CallOpts, _operatorId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "getUserUnstakeButOperatorNoExitNfs", _operatorId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserUnstakeButOperatorNoExitNfs is a free data retrieval call binding the contract method 0xf1c28381.
//
// Solidity: function getUserUnstakeButOperatorNoExitNfs(uint256 _operatorId) view returns(uint256[])
func (_Liq *LiqSession) GetUserUnstakeButOperatorNoExitNfs(_operatorId *big.Int) ([]*big.Int, error) {
	return _Liq.Contract.GetUserUnstakeButOperatorNoExitNfs(&_Liq.CallOpts, _operatorId)
}

// GetUserUnstakeButOperatorNoExitNfs is a free data retrieval call binding the contract method 0xf1c28381.
//
// Solidity: function getUserUnstakeButOperatorNoExitNfs(uint256 _operatorId) view returns(uint256[])
func (_Liq *LiqCallerSession) GetUserUnstakeButOperatorNoExitNfs(_operatorId *big.Int) ([]*big.Int, error) {
	return _Liq.Contract.GetUserUnstakeButOperatorNoExitNfs(&_Liq.CallOpts, _operatorId)
}

// GetWithdrawalOfOperator is a free data retrieval call binding the contract method 0x46a2ca47.
//
// Solidity: function getWithdrawalOfOperator(uint256 _operatorId) view returns((uint256,uint256,uint256,uint256,uint256,address,bool)[])
func (_Liq *LiqCaller) GetWithdrawalOfOperator(opts *bind.CallOpts, _operatorId *big.Int) ([]LiquidStakingWithdrawalInfo, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "getWithdrawalOfOperator", _operatorId)

	if err != nil {
		return *new([]LiquidStakingWithdrawalInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]LiquidStakingWithdrawalInfo)).(*[]LiquidStakingWithdrawalInfo)

	return out0, err

}

// GetWithdrawalOfOperator is a free data retrieval call binding the contract method 0x46a2ca47.
//
// Solidity: function getWithdrawalOfOperator(uint256 _operatorId) view returns((uint256,uint256,uint256,uint256,uint256,address,bool)[])
func (_Liq *LiqSession) GetWithdrawalOfOperator(_operatorId *big.Int) ([]LiquidStakingWithdrawalInfo, error) {
	return _Liq.Contract.GetWithdrawalOfOperator(&_Liq.CallOpts, _operatorId)
}

// GetWithdrawalOfOperator is a free data retrieval call binding the contract method 0x46a2ca47.
//
// Solidity: function getWithdrawalOfOperator(uint256 _operatorId) view returns((uint256,uint256,uint256,uint256,uint256,address,bool)[])
func (_Liq *LiqCallerSession) GetWithdrawalOfOperator(_operatorId *big.Int) ([]LiquidStakingWithdrawalInfo, error) {
	return _Liq.Contract.GetWithdrawalOfOperator(&_Liq.CallOpts, _operatorId)
}

// GetWithdrawalRequestIdOfOwner is a free data retrieval call binding the contract method 0xcdc2b1c7.
//
// Solidity: function getWithdrawalRequestIdOfOwner(address _owner) view returns(uint256[])
func (_Liq *LiqCaller) GetWithdrawalRequestIdOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "getWithdrawalRequestIdOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetWithdrawalRequestIdOfOwner is a free data retrieval call binding the contract method 0xcdc2b1c7.
//
// Solidity: function getWithdrawalRequestIdOfOwner(address _owner) view returns(uint256[])
func (_Liq *LiqSession) GetWithdrawalRequestIdOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _Liq.Contract.GetWithdrawalRequestIdOfOwner(&_Liq.CallOpts, _owner)
}

// GetWithdrawalRequestIdOfOwner is a free data retrieval call binding the contract method 0xcdc2b1c7.
//
// Solidity: function getWithdrawalRequestIdOfOwner(address _owner) view returns(uint256[])
func (_Liq *LiqCallerSession) GetWithdrawalRequestIdOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _Liq.Contract.GetWithdrawalRequestIdOfOwner(&_Liq.CallOpts, _owner)
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

// LargeExitDelayedSlashRecords is a free data retrieval call binding the contract method 0x9d4b4fd5.
//
// Solidity: function largeExitDelayedSlashRecords(uint256 ) view returns(uint256)
func (_Liq *LiqCaller) LargeExitDelayedSlashRecords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "largeExitDelayedSlashRecords", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LargeExitDelayedSlashRecords is a free data retrieval call binding the contract method 0x9d4b4fd5.
//
// Solidity: function largeExitDelayedSlashRecords(uint256 ) view returns(uint256)
func (_Liq *LiqSession) LargeExitDelayedSlashRecords(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.LargeExitDelayedSlashRecords(&_Liq.CallOpts, arg0)
}

// LargeExitDelayedSlashRecords is a free data retrieval call binding the contract method 0x9d4b4fd5.
//
// Solidity: function largeExitDelayedSlashRecords(uint256 ) view returns(uint256)
func (_Liq *LiqCallerSession) LargeExitDelayedSlashRecords(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.LargeExitDelayedSlashRecords(&_Liq.CallOpts, arg0)
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

// NftExitDelayedSlashRecords is a free data retrieval call binding the contract method 0x1d3fb6bd.
//
// Solidity: function nftExitDelayedSlashRecords(uint256 ) view returns(uint256)
func (_Liq *LiqCaller) NftExitDelayedSlashRecords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "nftExitDelayedSlashRecords", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftExitDelayedSlashRecords is a free data retrieval call binding the contract method 0x1d3fb6bd.
//
// Solidity: function nftExitDelayedSlashRecords(uint256 ) view returns(uint256)
func (_Liq *LiqSession) NftExitDelayedSlashRecords(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.NftExitDelayedSlashRecords(&_Liq.CallOpts, arg0)
}

// NftExitDelayedSlashRecords is a free data retrieval call binding the contract method 0x1d3fb6bd.
//
// Solidity: function nftExitDelayedSlashRecords(uint256 ) view returns(uint256)
func (_Liq *LiqCallerSession) NftExitDelayedSlashRecords(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.NftExitDelayedSlashRecords(&_Liq.CallOpts, arg0)
}

// NftHasCompensated is a free data retrieval call binding the contract method 0xa122faa4.
//
// Solidity: function nftHasCompensated(uint256 ) view returns(uint256)
func (_Liq *LiqCaller) NftHasCompensated(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "nftHasCompensated", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftHasCompensated is a free data retrieval call binding the contract method 0xa122faa4.
//
// Solidity: function nftHasCompensated(uint256 ) view returns(uint256)
func (_Liq *LiqSession) NftHasCompensated(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.NftHasCompensated(&_Liq.CallOpts, arg0)
}

// NftHasCompensated is a free data retrieval call binding the contract method 0xa122faa4.
//
// Solidity: function nftHasCompensated(uint256 ) view returns(uint256)
func (_Liq *LiqCallerSession) NftHasCompensated(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.NftHasCompensated(&_Liq.CallOpts, arg0)
}

// NftUnstakeBlockNumbers is a free data retrieval call binding the contract method 0x5f68dbf8.
//
// Solidity: function nftUnstakeBlockNumbers(uint256 ) view returns(uint256)
func (_Liq *LiqCaller) NftUnstakeBlockNumbers(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "nftUnstakeBlockNumbers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftUnstakeBlockNumbers is a free data retrieval call binding the contract method 0x5f68dbf8.
//
// Solidity: function nftUnstakeBlockNumbers(uint256 ) view returns(uint256)
func (_Liq *LiqSession) NftUnstakeBlockNumbers(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.NftUnstakeBlockNumbers(&_Liq.CallOpts, arg0)
}

// NftUnstakeBlockNumbers is a free data retrieval call binding the contract method 0x5f68dbf8.
//
// Solidity: function nftUnstakeBlockNumbers(uint256 ) view returns(uint256)
func (_Liq *LiqCallerSession) NftUnstakeBlockNumbers(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.NftUnstakeBlockNumbers(&_Liq.CallOpts, arg0)
}

// NftWillCompensated is a free data retrieval call binding the contract method 0xe0b64f27.
//
// Solidity: function nftWillCompensated(uint256 ) view returns(uint256)
func (_Liq *LiqCaller) NftWillCompensated(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "nftWillCompensated", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftWillCompensated is a free data retrieval call binding the contract method 0xe0b64f27.
//
// Solidity: function nftWillCompensated(uint256 ) view returns(uint256)
func (_Liq *LiqSession) NftWillCompensated(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.NftWillCompensated(&_Liq.CallOpts, arg0)
}

// NftWillCompensated is a free data retrieval call binding the contract method 0xe0b64f27.
//
// Solidity: function nftWillCompensated(uint256 ) view returns(uint256)
func (_Liq *LiqCallerSession) NftWillCompensated(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.NftWillCompensated(&_Liq.CallOpts, arg0)
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

// OperatorCompensatedIndex is a free data retrieval call binding the contract method 0x2661005e.
//
// Solidity: function operatorCompensatedIndex() view returns(uint256)
func (_Liq *LiqCaller) OperatorCompensatedIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "operatorCompensatedIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorCompensatedIndex is a free data retrieval call binding the contract method 0x2661005e.
//
// Solidity: function operatorCompensatedIndex() view returns(uint256)
func (_Liq *LiqSession) OperatorCompensatedIndex() (*big.Int, error) {
	return _Liq.Contract.OperatorCompensatedIndex(&_Liq.CallOpts)
}

// OperatorCompensatedIndex is a free data retrieval call binding the contract method 0x2661005e.
//
// Solidity: function operatorCompensatedIndex() view returns(uint256)
func (_Liq *LiqCallerSession) OperatorCompensatedIndex() (*big.Int, error) {
	return _Liq.Contract.OperatorCompensatedIndex(&_Liq.CallOpts)
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

// OperatorPendingEthPoolBalance is a free data retrieval call binding the contract method 0x20109b36.
//
// Solidity: function operatorPendingEthPoolBalance(uint256 ) view returns(uint256)
func (_Liq *LiqCaller) OperatorPendingEthPoolBalance(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "operatorPendingEthPoolBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorPendingEthPoolBalance is a free data retrieval call binding the contract method 0x20109b36.
//
// Solidity: function operatorPendingEthPoolBalance(uint256 ) view returns(uint256)
func (_Liq *LiqSession) OperatorPendingEthPoolBalance(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.OperatorPendingEthPoolBalance(&_Liq.CallOpts, arg0)
}

// OperatorPendingEthPoolBalance is a free data retrieval call binding the contract method 0x20109b36.
//
// Solidity: function operatorPendingEthPoolBalance(uint256 ) view returns(uint256)
func (_Liq *LiqCallerSession) OperatorPendingEthPoolBalance(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.OperatorPendingEthPoolBalance(&_Liq.CallOpts, arg0)
}

// OperatorPendingEthRequestAmount is a free data retrieval call binding the contract method 0x67dd7ffd.
//
// Solidity: function operatorPendingEthRequestAmount(uint256 ) view returns(uint256)
func (_Liq *LiqCaller) OperatorPendingEthRequestAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "operatorPendingEthRequestAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorPendingEthRequestAmount is a free data retrieval call binding the contract method 0x67dd7ffd.
//
// Solidity: function operatorPendingEthRequestAmount(uint256 ) view returns(uint256)
func (_Liq *LiqSession) OperatorPendingEthRequestAmount(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.OperatorPendingEthRequestAmount(&_Liq.CallOpts, arg0)
}

// OperatorPendingEthRequestAmount is a free data retrieval call binding the contract method 0x67dd7ffd.
//
// Solidity: function operatorPendingEthRequestAmount(uint256 ) view returns(uint256)
func (_Liq *LiqCallerSession) OperatorPendingEthRequestAmount(arg0 *big.Int) (*big.Int, error) {
	return _Liq.Contract.OperatorPendingEthRequestAmount(&_Liq.CallOpts, arg0)
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

// OperatorSlashArrears is a free data retrieval call binding the contract method 0xa7d18eba.
//
// Solidity: function operatorSlashArrears(uint256 , uint256 ) view returns(uint256)
func (_Liq *LiqCaller) OperatorSlashArrears(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "operatorSlashArrears", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorSlashArrears is a free data retrieval call binding the contract method 0xa7d18eba.
//
// Solidity: function operatorSlashArrears(uint256 , uint256 ) view returns(uint256)
func (_Liq *LiqSession) OperatorSlashArrears(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Liq.Contract.OperatorSlashArrears(&_Liq.CallOpts, arg0, arg1)
}

// OperatorSlashArrears is a free data retrieval call binding the contract method 0xa7d18eba.
//
// Solidity: function operatorSlashArrears(uint256 , uint256 ) view returns(uint256)
func (_Liq *LiqCallerSession) OperatorSlashArrears(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Liq.Contract.OperatorSlashArrears(&_Liq.CallOpts, arg0, arg1)
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

// SlashAmountPerBlockPerValidator is a free data retrieval call binding the contract method 0xaf256032.
//
// Solidity: function slashAmountPerBlockPerValidator() view returns(uint256)
func (_Liq *LiqCaller) SlashAmountPerBlockPerValidator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "slashAmountPerBlockPerValidator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashAmountPerBlockPerValidator is a free data retrieval call binding the contract method 0xaf256032.
//
// Solidity: function slashAmountPerBlockPerValidator() view returns(uint256)
func (_Liq *LiqSession) SlashAmountPerBlockPerValidator() (*big.Int, error) {
	return _Liq.Contract.SlashAmountPerBlockPerValidator(&_Liq.CallOpts)
}

// SlashAmountPerBlockPerValidator is a free data retrieval call binding the contract method 0xaf256032.
//
// Solidity: function slashAmountPerBlockPerValidator() view returns(uint256)
func (_Liq *LiqCallerSession) SlashAmountPerBlockPerValidator() (*big.Int, error) {
	return _Liq.Contract.SlashAmountPerBlockPerValidator(&_Liq.CallOpts)
}

// StakeRecords is a free data retrieval call binding the contract method 0x84660ae6.
//
// Solidity: function stakeRecords(address , uint256 ) view returns(uint256 operatorId, uint256 quota)
func (_Liq *LiqCaller) StakeRecords(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	OperatorId *big.Int
	Quota      *big.Int
}, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "stakeRecords", arg0, arg1)

	outstruct := new(struct {
		OperatorId *big.Int
		Quota      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OperatorId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Quota = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StakeRecords is a free data retrieval call binding the contract method 0x84660ae6.
//
// Solidity: function stakeRecords(address , uint256 ) view returns(uint256 operatorId, uint256 quota)
func (_Liq *LiqSession) StakeRecords(arg0 common.Address, arg1 *big.Int) (struct {
	OperatorId *big.Int
	Quota      *big.Int
}, error) {
	return _Liq.Contract.StakeRecords(&_Liq.CallOpts, arg0, arg1)
}

// StakeRecords is a free data retrieval call binding the contract method 0x84660ae6.
//
// Solidity: function stakeRecords(address , uint256 ) view returns(uint256 operatorId, uint256 quota)
func (_Liq *LiqCallerSession) StakeRecords(arg0 common.Address, arg1 *big.Int) (struct {
	OperatorId *big.Int
	Quota      *big.Int
}, error) {
	return _Liq.Contract.StakeRecords(&_Liq.CallOpts, arg0, arg1)
}

// TotalLockedNethBalance is a free data retrieval call binding the contract method 0xc03c381c.
//
// Solidity: function totalLockedNethBalance() view returns(uint256)
func (_Liq *LiqCaller) TotalLockedNethBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "totalLockedNethBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLockedNethBalance is a free data retrieval call binding the contract method 0xc03c381c.
//
// Solidity: function totalLockedNethBalance() view returns(uint256)
func (_Liq *LiqSession) TotalLockedNethBalance() (*big.Int, error) {
	return _Liq.Contract.TotalLockedNethBalance(&_Liq.CallOpts)
}

// TotalLockedNethBalance is a free data retrieval call binding the contract method 0xc03c381c.
//
// Solidity: function totalLockedNethBalance() view returns(uint256)
func (_Liq *LiqCallerSession) TotalLockedNethBalance() (*big.Int, error) {
	return _Liq.Contract.TotalLockedNethBalance(&_Liq.CallOpts)
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

// WithdrawalQueues is a free data retrieval call binding the contract method 0x7743b6d3.
//
// Solidity: function withdrawalQueues(uint256 ) view returns(uint256 operatorId, uint256 withdrawHeight, uint256 withdrawNethAmount, uint256 withdrawExchange, uint256 claimEthAmount, address owner, bool isClaim)
func (_Liq *LiqCaller) WithdrawalQueues(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OperatorId         *big.Int
	WithdrawHeight     *big.Int
	WithdrawNethAmount *big.Int
	WithdrawExchange   *big.Int
	ClaimEthAmount     *big.Int
	Owner              common.Address
	IsClaim            bool
}, error) {
	var out []interface{}
	err := _Liq.contract.Call(opts, &out, "withdrawalQueues", arg0)

	outstruct := new(struct {
		OperatorId         *big.Int
		WithdrawHeight     *big.Int
		WithdrawNethAmount *big.Int
		WithdrawExchange   *big.Int
		ClaimEthAmount     *big.Int
		Owner              common.Address
		IsClaim            bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OperatorId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WithdrawHeight = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.WithdrawNethAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.WithdrawExchange = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ClaimEthAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.IsClaim = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// WithdrawalQueues is a free data retrieval call binding the contract method 0x7743b6d3.
//
// Solidity: function withdrawalQueues(uint256 ) view returns(uint256 operatorId, uint256 withdrawHeight, uint256 withdrawNethAmount, uint256 withdrawExchange, uint256 claimEthAmount, address owner, bool isClaim)
func (_Liq *LiqSession) WithdrawalQueues(arg0 *big.Int) (struct {
	OperatorId         *big.Int
	WithdrawHeight     *big.Int
	WithdrawNethAmount *big.Int
	WithdrawExchange   *big.Int
	ClaimEthAmount     *big.Int
	Owner              common.Address
	IsClaim            bool
}, error) {
	return _Liq.Contract.WithdrawalQueues(&_Liq.CallOpts, arg0)
}

// WithdrawalQueues is a free data retrieval call binding the contract method 0x7743b6d3.
//
// Solidity: function withdrawalQueues(uint256 ) view returns(uint256 operatorId, uint256 withdrawHeight, uint256 withdrawNethAmount, uint256 withdrawExchange, uint256 claimEthAmount, address owner, bool isClaim)
func (_Liq *LiqCallerSession) WithdrawalQueues(arg0 *big.Int) (struct {
	OperatorId         *big.Int
	WithdrawHeight     *big.Int
	WithdrawNethAmount *big.Int
	WithdrawExchange   *big.Int
	ClaimEthAmount     *big.Int
	Owner              common.Address
	IsClaim            bool
}, error) {
	return _Liq.Contract.WithdrawalQueues(&_Liq.CallOpts, arg0)
}

// AssignBlacklistOperator is a paid mutator transaction binding the contract method 0x31348aa8.
//
// Solidity: function assignBlacklistOperator(uint256 _assignOperatorId, uint256 _operatorId) returns()
func (_Liq *LiqTransactor) AssignBlacklistOperator(opts *bind.TransactOpts, _assignOperatorId *big.Int, _operatorId *big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "assignBlacklistOperator", _assignOperatorId, _operatorId)
}

// AssignBlacklistOperator is a paid mutator transaction binding the contract method 0x31348aa8.
//
// Solidity: function assignBlacklistOperator(uint256 _assignOperatorId, uint256 _operatorId) returns()
func (_Liq *LiqSession) AssignBlacklistOperator(_assignOperatorId *big.Int, _operatorId *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.AssignBlacklistOperator(&_Liq.TransactOpts, _assignOperatorId, _operatorId)
}

// AssignBlacklistOperator is a paid mutator transaction binding the contract method 0x31348aa8.
//
// Solidity: function assignBlacklistOperator(uint256 _assignOperatorId, uint256 _operatorId) returns()
func (_Liq *LiqTransactorSession) AssignBlacklistOperator(_assignOperatorId *big.Int, _operatorId *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.AssignBlacklistOperator(&_Liq.TransactOpts, _assignOperatorId, _operatorId)
}

// AssignQuitOperator is a paid mutator transaction binding the contract method 0x04916b32.
//
// Solidity: function assignQuitOperator(uint256 _quitOperatorId, uint256 _operatorId) returns()
func (_Liq *LiqTransactor) AssignQuitOperator(opts *bind.TransactOpts, _quitOperatorId *big.Int, _operatorId *big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "assignQuitOperator", _quitOperatorId, _operatorId)
}

// AssignQuitOperator is a paid mutator transaction binding the contract method 0x04916b32.
//
// Solidity: function assignQuitOperator(uint256 _quitOperatorId, uint256 _operatorId) returns()
func (_Liq *LiqSession) AssignQuitOperator(_quitOperatorId *big.Int, _operatorId *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.AssignQuitOperator(&_Liq.TransactOpts, _quitOperatorId, _operatorId)
}

// AssignQuitOperator is a paid mutator transaction binding the contract method 0x04916b32.
//
// Solidity: function assignQuitOperator(uint256 _quitOperatorId, uint256 _operatorId) returns()
func (_Liq *LiqTransactorSession) AssignQuitOperator(_quitOperatorId *big.Int, _operatorId *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.AssignQuitOperator(&_Liq.TransactOpts, _quitOperatorId, _operatorId)
}

// ClaimLargeWitdrawals is a paid mutator transaction binding the contract method 0x97ca2c61.
//
// Solidity: function claimLargeWitdrawals(uint256[] requestIds) returns()
func (_Liq *LiqTransactor) ClaimLargeWitdrawals(opts *bind.TransactOpts, requestIds []*big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "claimLargeWitdrawals", requestIds)
}

// ClaimLargeWitdrawals is a paid mutator transaction binding the contract method 0x97ca2c61.
//
// Solidity: function claimLargeWitdrawals(uint256[] requestIds) returns()
func (_Liq *LiqSession) ClaimLargeWitdrawals(requestIds []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.ClaimLargeWitdrawals(&_Liq.TransactOpts, requestIds)
}

// ClaimLargeWitdrawals is a paid mutator transaction binding the contract method 0x97ca2c61.
//
// Solidity: function claimLargeWitdrawals(uint256[] requestIds) returns()
func (_Liq *LiqTransactorSession) ClaimLargeWitdrawals(requestIds []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.ClaimLargeWitdrawals(&_Liq.TransactOpts, requestIds)
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

// ClaimRewardsOfOperator is a paid mutator transaction binding the contract method 0x38fac833.
//
// Solidity: function claimRewardsOfOperator(uint256 _operatorId, uint256 _reward) returns()
func (_Liq *LiqTransactor) ClaimRewardsOfOperator(opts *bind.TransactOpts, _operatorId *big.Int, _reward *big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "claimRewardsOfOperator", _operatorId, _reward)
}

// ClaimRewardsOfOperator is a paid mutator transaction binding the contract method 0x38fac833.
//
// Solidity: function claimRewardsOfOperator(uint256 _operatorId, uint256 _reward) returns()
func (_Liq *LiqSession) ClaimRewardsOfOperator(_operatorId *big.Int, _reward *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.ClaimRewardsOfOperator(&_Liq.TransactOpts, _operatorId, _reward)
}

// ClaimRewardsOfOperator is a paid mutator transaction binding the contract method 0x38fac833.
//
// Solidity: function claimRewardsOfOperator(uint256 _operatorId, uint256 _reward) returns()
func (_Liq *LiqTransactorSession) ClaimRewardsOfOperator(_operatorId *big.Int, _reward *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.ClaimRewardsOfOperator(&_Liq.TransactOpts, _operatorId, _reward)
}

// ClaimRewardsOfUser is a paid mutator transaction binding the contract method 0x538823f4.
//
// Solidity: function claimRewardsOfUser(uint256 _operatorId, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasHeight) returns()
func (_Liq *LiqTransactor) ClaimRewardsOfUser(opts *bind.TransactOpts, _operatorId *big.Int, _tokenIds []*big.Int, _amounts []*big.Int, _gasHeight *big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "claimRewardsOfUser", _operatorId, _tokenIds, _amounts, _gasHeight)
}

// ClaimRewardsOfUser is a paid mutator transaction binding the contract method 0x538823f4.
//
// Solidity: function claimRewardsOfUser(uint256 _operatorId, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasHeight) returns()
func (_Liq *LiqSession) ClaimRewardsOfUser(_operatorId *big.Int, _tokenIds []*big.Int, _amounts []*big.Int, _gasHeight *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.ClaimRewardsOfUser(&_Liq.TransactOpts, _operatorId, _tokenIds, _amounts, _gasHeight)
}

// ClaimRewardsOfUser is a paid mutator transaction binding the contract method 0x538823f4.
//
// Solidity: function claimRewardsOfUser(uint256 _operatorId, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasHeight) returns()
func (_Liq *LiqTransactorSession) ClaimRewardsOfUser(_operatorId *big.Int, _tokenIds []*big.Int, _amounts []*big.Int, _gasHeight *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.ClaimRewardsOfUser(&_Liq.TransactOpts, _operatorId, _tokenIds, _amounts, _gasHeight)
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

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_Liq *LiqTransactor) InitializeV2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "initializeV2")
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_Liq *LiqSession) InitializeV2() (*types.Transaction, error) {
	return _Liq.Contract.InitializeV2(&_Liq.TransactOpts)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_Liq *LiqTransactorSession) InitializeV2() (*types.Transaction, error) {
	return _Liq.Contract.InitializeV2(&_Liq.TransactOpts)
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

// RequestLargeWithdrawals is a paid mutator transaction binding the contract method 0x976b364e.
//
// Solidity: function requestLargeWithdrawals(uint256 _operatorId, uint256[] _amounts) returns()
func (_Liq *LiqTransactor) RequestLargeWithdrawals(opts *bind.TransactOpts, _operatorId *big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "requestLargeWithdrawals", _operatorId, _amounts)
}

// RequestLargeWithdrawals is a paid mutator transaction binding the contract method 0x976b364e.
//
// Solidity: function requestLargeWithdrawals(uint256 _operatorId, uint256[] _amounts) returns()
func (_Liq *LiqSession) RequestLargeWithdrawals(_operatorId *big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.RequestLargeWithdrawals(&_Liq.TransactOpts, _operatorId, _amounts)
}

// RequestLargeWithdrawals is a paid mutator transaction binding the contract method 0x976b364e.
//
// Solidity: function requestLargeWithdrawals(uint256 _operatorId, uint256[] _amounts) returns()
func (_Liq *LiqTransactorSession) RequestLargeWithdrawals(_operatorId *big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.RequestLargeWithdrawals(&_Liq.TransactOpts, _operatorId, _amounts)
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

// SlashArrearsReceive is a paid mutator transaction binding the contract method 0xe72544c9.
//
// Solidity: function slashArrearsReceive(uint256 _operatorId, uint256 _amount) payable returns()
func (_Liq *LiqTransactor) SlashArrearsReceive(opts *bind.TransactOpts, _operatorId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "slashArrearsReceive", _operatorId, _amount)
}

// SlashArrearsReceive is a paid mutator transaction binding the contract method 0xe72544c9.
//
// Solidity: function slashArrearsReceive(uint256 _operatorId, uint256 _amount) payable returns()
func (_Liq *LiqSession) SlashArrearsReceive(_operatorId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.SlashArrearsReceive(&_Liq.TransactOpts, _operatorId, _amount)
}

// SlashArrearsReceive is a paid mutator transaction binding the contract method 0xe72544c9.
//
// Solidity: function slashArrearsReceive(uint256 _operatorId, uint256 _amount) payable returns()
func (_Liq *LiqTransactorSession) SlashArrearsReceive(_operatorId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Liq.Contract.SlashArrearsReceive(&_Liq.TransactOpts, _operatorId, _amount)
}

// SlashOfExitDelayed is a paid mutator transaction binding the contract method 0xb21a6370.
//
// Solidity: function slashOfExitDelayed(uint256[] _nftExitDelayedTokenIds, uint256[] _largeExitDelayedRequestIds) returns()
func (_Liq *LiqTransactor) SlashOfExitDelayed(opts *bind.TransactOpts, _nftExitDelayedTokenIds []*big.Int, _largeExitDelayedRequestIds []*big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "slashOfExitDelayed", _nftExitDelayedTokenIds, _largeExitDelayedRequestIds)
}

// SlashOfExitDelayed is a paid mutator transaction binding the contract method 0xb21a6370.
//
// Solidity: function slashOfExitDelayed(uint256[] _nftExitDelayedTokenIds, uint256[] _largeExitDelayedRequestIds) returns()
func (_Liq *LiqSession) SlashOfExitDelayed(_nftExitDelayedTokenIds []*big.Int, _largeExitDelayedRequestIds []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.SlashOfExitDelayed(&_Liq.TransactOpts, _nftExitDelayedTokenIds, _largeExitDelayedRequestIds)
}

// SlashOfExitDelayed is a paid mutator transaction binding the contract method 0xb21a6370.
//
// Solidity: function slashOfExitDelayed(uint256[] _nftExitDelayedTokenIds, uint256[] _largeExitDelayedRequestIds) returns()
func (_Liq *LiqTransactorSession) SlashOfExitDelayed(_nftExitDelayedTokenIds []*big.Int, _largeExitDelayedRequestIds []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.SlashOfExitDelayed(&_Liq.TransactOpts, _nftExitDelayedTokenIds, _largeExitDelayedRequestIds)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x538d7ab0.
//
// Solidity: function slashOperator(uint256[] _exitTokenIds, uint256[] _amounts) returns()
func (_Liq *LiqTransactor) SlashOperator(opts *bind.TransactOpts, _exitTokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "slashOperator", _exitTokenIds, _amounts)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x538d7ab0.
//
// Solidity: function slashOperator(uint256[] _exitTokenIds, uint256[] _amounts) returns()
func (_Liq *LiqSession) SlashOperator(_exitTokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.SlashOperator(&_Liq.TransactOpts, _exitTokenIds, _amounts)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x538d7ab0.
//
// Solidity: function slashOperator(uint256[] _exitTokenIds, uint256[] _amounts) returns()
func (_Liq *LiqTransactorSession) SlashOperator(_exitTokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.SlashOperator(&_Liq.TransactOpts, _exitTokenIds, _amounts)
}

// SlashReceive is a paid mutator transaction binding the contract method 0xc0849498.
//
// Solidity: function slashReceive(uint256[] _exitTokenIds, uint256[] _slashAmounts, uint256[] _requireAmounts) payable returns()
func (_Liq *LiqTransactor) SlashReceive(opts *bind.TransactOpts, _exitTokenIds []*big.Int, _slashAmounts []*big.Int, _requireAmounts []*big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "slashReceive", _exitTokenIds, _slashAmounts, _requireAmounts)
}

// SlashReceive is a paid mutator transaction binding the contract method 0xc0849498.
//
// Solidity: function slashReceive(uint256[] _exitTokenIds, uint256[] _slashAmounts, uint256[] _requireAmounts) payable returns()
func (_Liq *LiqSession) SlashReceive(_exitTokenIds []*big.Int, _slashAmounts []*big.Int, _requireAmounts []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.SlashReceive(&_Liq.TransactOpts, _exitTokenIds, _slashAmounts, _requireAmounts)
}

// SlashReceive is a paid mutator transaction binding the contract method 0xc0849498.
//
// Solidity: function slashReceive(uint256[] _exitTokenIds, uint256[] _slashAmounts, uint256[] _requireAmounts) payable returns()
func (_Liq *LiqTransactorSession) SlashReceive(_exitTokenIds []*big.Int, _slashAmounts []*big.Int, _requireAmounts []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.SlashReceive(&_Liq.TransactOpts, _exitTokenIds, _slashAmounts, _requireAmounts)
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

// UnstakeNFT is a paid mutator transaction binding the contract method 0x8453734d.
//
// Solidity: function unstakeNFT(uint256[] _tokenIds) returns()
func (_Liq *LiqTransactor) UnstakeNFT(opts *bind.TransactOpts, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Liq.contract.Transact(opts, "unstakeNFT", _tokenIds)
}

// UnstakeNFT is a paid mutator transaction binding the contract method 0x8453734d.
//
// Solidity: function unstakeNFT(uint256[] _tokenIds) returns()
func (_Liq *LiqSession) UnstakeNFT(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.UnstakeNFT(&_Liq.TransactOpts, _tokenIds)
}

// UnstakeNFT is a paid mutator transaction binding the contract method 0x8453734d.
//
// Solidity: function unstakeNFT(uint256[] _tokenIds) returns()
func (_Liq *LiqTransactorSession) UnstakeNFT(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _Liq.Contract.UnstakeNFT(&_Liq.TransactOpts, _tokenIds)
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

// LiqArrearsReceiveOfSlashIterator is returned from FilterArrearsReceiveOfSlash and is used to iterate over the raw logs and unpacked data for ArrearsReceiveOfSlash events raised by the Liq contract.
type LiqArrearsReceiveOfSlashIterator struct {
	Event *LiqArrearsReceiveOfSlash // Event containing the contract specifics and raw log

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
func (it *LiqArrearsReceiveOfSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqArrearsReceiveOfSlash)
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
		it.Event = new(LiqArrearsReceiveOfSlash)
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
func (it *LiqArrearsReceiveOfSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqArrearsReceiveOfSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqArrearsReceiveOfSlash represents a ArrearsReceiveOfSlash event raised by the Liq contract.
type LiqArrearsReceiveOfSlash struct {
	OperatorId *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterArrearsReceiveOfSlash is a free log retrieval operation binding the contract event 0xa2478ef76ec5108b836ec0b59cc344bf613e964e82254a6c2e22a548e9d0fe75.
//
// Solidity: event ArrearsReceiveOfSlash(uint256 _operatorId, uint256 _amount)
func (_Liq *LiqFilterer) FilterArrearsReceiveOfSlash(opts *bind.FilterOpts) (*LiqArrearsReceiveOfSlashIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "ArrearsReceiveOfSlash")
	if err != nil {
		return nil, err
	}
	return &LiqArrearsReceiveOfSlashIterator{contract: _Liq.contract, event: "ArrearsReceiveOfSlash", logs: logs, sub: sub}, nil
}

// WatchArrearsReceiveOfSlash is a free log subscription operation binding the contract event 0xa2478ef76ec5108b836ec0b59cc344bf613e964e82254a6c2e22a548e9d0fe75.
//
// Solidity: event ArrearsReceiveOfSlash(uint256 _operatorId, uint256 _amount)
func (_Liq *LiqFilterer) WatchArrearsReceiveOfSlash(opts *bind.WatchOpts, sink chan<- *LiqArrearsReceiveOfSlash) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "ArrearsReceiveOfSlash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqArrearsReceiveOfSlash)
				if err := _Liq.contract.UnpackLog(event, "ArrearsReceiveOfSlash", log); err != nil {
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
func (_Liq *LiqFilterer) ParseArrearsReceiveOfSlash(log types.Log) (*LiqArrearsReceiveOfSlash, error) {
	event := new(LiqArrearsReceiveOfSlash)
	if err := _Liq.contract.UnpackLog(event, "ArrearsReceiveOfSlash", log); err != nil {
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

// LiqBlacklistOperatorAssignedIterator is returned from FilterBlacklistOperatorAssigned and is used to iterate over the raw logs and unpacked data for BlacklistOperatorAssigned events raised by the Liq contract.
type LiqBlacklistOperatorAssignedIterator struct {
	Event *LiqBlacklistOperatorAssigned // Event containing the contract specifics and raw log

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
func (it *LiqBlacklistOperatorAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqBlacklistOperatorAssigned)
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
		it.Event = new(LiqBlacklistOperatorAssigned)
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
func (it *LiqBlacklistOperatorAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqBlacklistOperatorAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqBlacklistOperatorAssigned represents a BlacklistOperatorAssigned event raised by the Liq contract.
type LiqBlacklistOperatorAssigned struct {
	BlacklistOperatorId *big.Int
	OperatorId          *big.Int
	TotalAmount         *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterBlacklistOperatorAssigned is a free log retrieval operation binding the contract event 0x46e36b61e3d2c27e6dc88d3eb33962629e6989428bd68e322e1175e317e0521c.
//
// Solidity: event BlacklistOperatorAssigned(uint256 _blacklistOperatorId, uint256 _operatorId, uint256 _totalAmount)
func (_Liq *LiqFilterer) FilterBlacklistOperatorAssigned(opts *bind.FilterOpts) (*LiqBlacklistOperatorAssignedIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "BlacklistOperatorAssigned")
	if err != nil {
		return nil, err
	}
	return &LiqBlacklistOperatorAssignedIterator{contract: _Liq.contract, event: "BlacklistOperatorAssigned", logs: logs, sub: sub}, nil
}

// WatchBlacklistOperatorAssigned is a free log subscription operation binding the contract event 0x46e36b61e3d2c27e6dc88d3eb33962629e6989428bd68e322e1175e317e0521c.
//
// Solidity: event BlacklistOperatorAssigned(uint256 _blacklistOperatorId, uint256 _operatorId, uint256 _totalAmount)
func (_Liq *LiqFilterer) WatchBlacklistOperatorAssigned(opts *bind.WatchOpts, sink chan<- *LiqBlacklistOperatorAssigned) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "BlacklistOperatorAssigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqBlacklistOperatorAssigned)
				if err := _Liq.contract.UnpackLog(event, "BlacklistOperatorAssigned", log); err != nil {
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

// ParseBlacklistOperatorAssigned is a log parse operation binding the contract event 0x46e36b61e3d2c27e6dc88d3eb33962629e6989428bd68e322e1175e317e0521c.
//
// Solidity: event BlacklistOperatorAssigned(uint256 _blacklistOperatorId, uint256 _operatorId, uint256 _totalAmount)
func (_Liq *LiqFilterer) ParseBlacklistOperatorAssigned(log types.Log) (*LiqBlacklistOperatorAssigned, error) {
	event := new(LiqBlacklistOperatorAssigned)
	if err := _Liq.contract.UnpackLog(event, "BlacklistOperatorAssigned", log); err != nil {
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
	From      common.Address
	Amount    *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthStake is a free log retrieval operation binding the contract event 0x838d17987e57e587c458220b9b38723c41fbc3f397550b506712960a73ef19f9.
//
// Solidity: event EthStake(address indexed _from, uint256 _amount, uint256 _amountOut)
func (_Liq *LiqFilterer) FilterEthStake(opts *bind.FilterOpts, _from []common.Address) (*LiqEthStakeIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Liq.contract.FilterLogs(opts, "EthStake", _fromRule)
	if err != nil {
		return nil, err
	}
	return &LiqEthStakeIterator{contract: _Liq.contract, event: "EthStake", logs: logs, sub: sub}, nil
}

// WatchEthStake is a free log subscription operation binding the contract event 0x838d17987e57e587c458220b9b38723c41fbc3f397550b506712960a73ef19f9.
//
// Solidity: event EthStake(address indexed _from, uint256 _amount, uint256 _amountOut)
func (_Liq *LiqFilterer) WatchEthStake(opts *bind.WatchOpts, sink chan<- *LiqEthStake, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Liq.contract.WatchLogs(opts, "EthStake", _fromRule)
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

// ParseEthStake is a log parse operation binding the contract event 0x838d17987e57e587c458220b9b38723c41fbc3f397550b506712960a73ef19f9.
//
// Solidity: event EthStake(address indexed _from, uint256 _amount, uint256 _amountOut)
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
// Solidity: event EthUnstake(uint256 _operatorId, uint256 targetOperatorId, address ender, uint256 _amounts, uint256 amountOut)
func (_Liq *LiqFilterer) FilterEthUnstake(opts *bind.FilterOpts) (*LiqEthUnstakeIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "EthUnstake")
	if err != nil {
		return nil, err
	}
	return &LiqEthUnstakeIterator{contract: _Liq.contract, event: "EthUnstake", logs: logs, sub: sub}, nil
}

// WatchEthUnstake is a free log subscription operation binding the contract event 0xc2d18d1ab67a48ae80c3ef1d20c2f2a97201a23db7ca49e5de1edf05610fb003.
//
// Solidity: event EthUnstake(uint256 _operatorId, uint256 targetOperatorId, address ender, uint256 _amounts, uint256 amountOut)
func (_Liq *LiqFilterer) WatchEthUnstake(opts *bind.WatchOpts, sink chan<- *LiqEthUnstake) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "EthUnstake")
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
// Solidity: event EthUnstake(uint256 _operatorId, uint256 targetOperatorId, address ender, uint256 _amounts, uint256 amountOut)
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

// LiqLargeWithdrawalsRequestIterator is returned from FilterLargeWithdrawalsRequest and is used to iterate over the raw logs and unpacked data for LargeWithdrawalsRequest events raised by the Liq contract.
type LiqLargeWithdrawalsRequestIterator struct {
	Event *LiqLargeWithdrawalsRequest // Event containing the contract specifics and raw log

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
func (it *LiqLargeWithdrawalsRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqLargeWithdrawalsRequest)
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
		it.Event = new(LiqLargeWithdrawalsRequest)
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
func (it *LiqLargeWithdrawalsRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqLargeWithdrawalsRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqLargeWithdrawalsRequest represents a LargeWithdrawalsRequest event raised by the Liq contract.
type LiqLargeWithdrawalsRequest struct {
	OperatorId      *big.Int
	Sender          common.Address
	TotalNethAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLargeWithdrawalsRequest is a free log retrieval operation binding the contract event 0x5e24974b0e489bf247cedb3c7c52d992a262766293923a8e3eba888371495a3f.
//
// Solidity: event LargeWithdrawalsRequest(uint256 _operatorId, address sender, uint256 totalNethAmount)
func (_Liq *LiqFilterer) FilterLargeWithdrawalsRequest(opts *bind.FilterOpts) (*LiqLargeWithdrawalsRequestIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "LargeWithdrawalsRequest")
	if err != nil {
		return nil, err
	}
	return &LiqLargeWithdrawalsRequestIterator{contract: _Liq.contract, event: "LargeWithdrawalsRequest", logs: logs, sub: sub}, nil
}

// WatchLargeWithdrawalsRequest is a free log subscription operation binding the contract event 0x5e24974b0e489bf247cedb3c7c52d992a262766293923a8e3eba888371495a3f.
//
// Solidity: event LargeWithdrawalsRequest(uint256 _operatorId, address sender, uint256 totalNethAmount)
func (_Liq *LiqFilterer) WatchLargeWithdrawalsRequest(opts *bind.WatchOpts, sink chan<- *LiqLargeWithdrawalsRequest) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "LargeWithdrawalsRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqLargeWithdrawalsRequest)
				if err := _Liq.contract.UnpackLog(event, "LargeWithdrawalsRequest", log); err != nil {
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
func (_Liq *LiqFilterer) ParseLargeWithdrawalsRequest(log types.Log) (*LiqLargeWithdrawalsRequest, error) {
	event := new(LiqLargeWithdrawalsRequest)
	if err := _Liq.contract.UnpackLog(event, "LargeWithdrawalsRequest", log); err != nil {
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
	From  common.Address
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNftStake is a free log retrieval operation binding the contract event 0xa763c79f5b64c651fefa444e4020953d79c1f9f4dfaf121251f97f5277ec6dd2.
//
// Solidity: event NftStake(address indexed _from, uint256 _count)
func (_Liq *LiqFilterer) FilterNftStake(opts *bind.FilterOpts, _from []common.Address) (*LiqNftStakeIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Liq.contract.FilterLogs(opts, "NftStake", _fromRule)
	if err != nil {
		return nil, err
	}
	return &LiqNftStakeIterator{contract: _Liq.contract, event: "NftStake", logs: logs, sub: sub}, nil
}

// WatchNftStake is a free log subscription operation binding the contract event 0xa763c79f5b64c651fefa444e4020953d79c1f9f4dfaf121251f97f5277ec6dd2.
//
// Solidity: event NftStake(address indexed _from, uint256 _count)
func (_Liq *LiqFilterer) WatchNftStake(opts *bind.WatchOpts, sink chan<- *LiqNftStake, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Liq.contract.WatchLogs(opts, "NftStake", _fromRule)
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

// ParseNftStake is a log parse operation binding the contract event 0xa763c79f5b64c651fefa444e4020953d79c1f9f4dfaf121251f97f5277ec6dd2.
//
// Solidity: event NftStake(address indexed _from, uint256 _count)
func (_Liq *LiqFilterer) ParseNftStake(log types.Log) (*LiqNftStake, error) {
	event := new(LiqNftStake)
	if err := _Liq.contract.UnpackLog(event, "NftStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqNftUnstakeIterator is returned from FilterNftUnstake and is used to iterate over the raw logs and unpacked data for NftUnstake events raised by the Liq contract.
type LiqNftUnstakeIterator struct {
	Event *LiqNftUnstake // Event containing the contract specifics and raw log

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
func (it *LiqNftUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqNftUnstake)
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
		it.Event = new(LiqNftUnstake)
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
func (it *LiqNftUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqNftUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqNftUnstake represents a NftUnstake event raised by the Liq contract.
type LiqNftUnstake struct {
	TokenId    *big.Int
	OperatorId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNftUnstake is a free log retrieval operation binding the contract event 0x9187b7a988fc2779224c0f80f4459cbb33f8c5c4be3656b4debaa4e58c19658e.
//
// Solidity: event NftUnstake(uint256 tokenId, uint256 operatorId)
func (_Liq *LiqFilterer) FilterNftUnstake(opts *bind.FilterOpts) (*LiqNftUnstakeIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "NftUnstake")
	if err != nil {
		return nil, err
	}
	return &LiqNftUnstakeIterator{contract: _Liq.contract, event: "NftUnstake", logs: logs, sub: sub}, nil
}

// WatchNftUnstake is a free log subscription operation binding the contract event 0x9187b7a988fc2779224c0f80f4459cbb33f8c5c4be3656b4debaa4e58c19658e.
//
// Solidity: event NftUnstake(uint256 tokenId, uint256 operatorId)
func (_Liq *LiqFilterer) WatchNftUnstake(opts *bind.WatchOpts, sink chan<- *LiqNftUnstake) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "NftUnstake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqNftUnstake)
				if err := _Liq.contract.UnpackLog(event, "NftUnstake", log); err != nil {
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
// Solidity: event NftUnstake(uint256 tokenId, uint256 operatorId)
func (_Liq *LiqFilterer) ParseNftUnstake(log types.Log) (*LiqNftUnstake, error) {
	event := new(LiqNftUnstake)
	if err := _Liq.contract.UnpackLog(event, "NftUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqNftUnwrapIterator is returned from FilterNftUnwrap and is used to iterate over the raw logs and unpacked data for NftUnwrap events raised by the Liq contract.
type LiqNftUnwrapIterator struct {
	Event *LiqNftUnwrap // Event containing the contract specifics and raw log

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
func (it *LiqNftUnwrapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqNftUnwrap)
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
		it.Event = new(LiqNftUnwrap)
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
func (it *LiqNftUnwrapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqNftUnwrapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqNftUnwrap represents a NftUnwrap event raised by the Liq contract.
type LiqNftUnwrap struct {
	TokenId    *big.Int
	OperatorId *big.Int
	Value      *big.Int
	AmountOut  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNftUnwrap is a free log retrieval operation binding the contract event 0x30935f3959f4a31d19f4c341041713ee7ffc6ee6dd0f9117beb992facbc149d7.
//
// Solidity: event NftUnwrap(uint256 _tokenId, uint256 operatorId, uint256 _value, uint256 _amountOut)
func (_Liq *LiqFilterer) FilterNftUnwrap(opts *bind.FilterOpts) (*LiqNftUnwrapIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "NftUnwrap")
	if err != nil {
		return nil, err
	}
	return &LiqNftUnwrapIterator{contract: _Liq.contract, event: "NftUnwrap", logs: logs, sub: sub}, nil
}

// WatchNftUnwrap is a free log subscription operation binding the contract event 0x30935f3959f4a31d19f4c341041713ee7ffc6ee6dd0f9117beb992facbc149d7.
//
// Solidity: event NftUnwrap(uint256 _tokenId, uint256 operatorId, uint256 _value, uint256 _amountOut)
func (_Liq *LiqFilterer) WatchNftUnwrap(opts *bind.WatchOpts, sink chan<- *LiqNftUnwrap) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "NftUnwrap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqNftUnwrap)
				if err := _Liq.contract.UnpackLog(event, "NftUnwrap", log); err != nil {
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

// ParseNftUnwrap is a log parse operation binding the contract event 0x30935f3959f4a31d19f4c341041713ee7ffc6ee6dd0f9117beb992facbc149d7.
//
// Solidity: event NftUnwrap(uint256 _tokenId, uint256 operatorId, uint256 _value, uint256 _amountOut)
func (_Liq *LiqFilterer) ParseNftUnwrap(log types.Log) (*LiqNftUnwrap, error) {
	event := new(LiqNftUnwrap)
	if err := _Liq.contract.UnpackLog(event, "NftUnwrap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiqNftWrapIterator is returned from FilterNftWrap and is used to iterate over the raw logs and unpacked data for NftWrap events raised by the Liq contract.
type LiqNftWrapIterator struct {
	Event *LiqNftWrap // Event containing the contract specifics and raw log

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
func (it *LiqNftWrapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqNftWrap)
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
		it.Event = new(LiqNftWrap)
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
func (it *LiqNftWrapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqNftWrapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqNftWrap represents a NftWrap event raised by the Liq contract.
type LiqNftWrap struct {
	TokenId    *big.Int
	OperatorId *big.Int
	Value      *big.Int
	AmountOut  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNftWrap is a free log retrieval operation binding the contract event 0x8893882104b7b95694bf51d917dc99f43bad27a00effe34bb555dfe280d14df4.
//
// Solidity: event NftWrap(uint256 _tokenId, uint256 _operatorId, uint256 _value, uint256 _amountOut)
func (_Liq *LiqFilterer) FilterNftWrap(opts *bind.FilterOpts) (*LiqNftWrapIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "NftWrap")
	if err != nil {
		return nil, err
	}
	return &LiqNftWrapIterator{contract: _Liq.contract, event: "NftWrap", logs: logs, sub: sub}, nil
}

// WatchNftWrap is a free log subscription operation binding the contract event 0x8893882104b7b95694bf51d917dc99f43bad27a00effe34bb555dfe280d14df4.
//
// Solidity: event NftWrap(uint256 _tokenId, uint256 _operatorId, uint256 _value, uint256 _amountOut)
func (_Liq *LiqFilterer) WatchNftWrap(opts *bind.WatchOpts, sink chan<- *LiqNftWrap) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "NftWrap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqNftWrap)
				if err := _Liq.contract.UnpackLog(event, "NftWrap", log); err != nil {
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

// ParseNftWrap is a log parse operation binding the contract event 0x8893882104b7b95694bf51d917dc99f43bad27a00effe34bb555dfe280d14df4.
//
// Solidity: event NftWrap(uint256 _tokenId, uint256 _operatorId, uint256 _value, uint256 _amountOut)
func (_Liq *LiqFilterer) ParseNftWrap(log types.Log) (*LiqNftWrap, error) {
	event := new(LiqNftWrap)
	if err := _Liq.contract.UnpackLog(event, "NftWrap", log); err != nil {
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

// LiqQuitOperatorAssignedIterator is returned from FilterQuitOperatorAssigned and is used to iterate over the raw logs and unpacked data for QuitOperatorAssigned events raised by the Liq contract.
type LiqQuitOperatorAssignedIterator struct {
	Event *LiqQuitOperatorAssigned // Event containing the contract specifics and raw log

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
func (it *LiqQuitOperatorAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqQuitOperatorAssigned)
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
		it.Event = new(LiqQuitOperatorAssigned)
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
func (it *LiqQuitOperatorAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqQuitOperatorAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqQuitOperatorAssigned represents a QuitOperatorAssigned event raised by the Liq contract.
type LiqQuitOperatorAssigned struct {
	QuitOperatorId *big.Int
	OperatorId     *big.Int
	TotalAmount    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterQuitOperatorAssigned is a free log retrieval operation binding the contract event 0xca1faa936bb688b7adf43d587c52f47df688f934011c256ad33a8fdeb51d382e.
//
// Solidity: event QuitOperatorAssigned(uint256 _quitOperatorId, uint256 _operatorId, uint256 _totalAmount)
func (_Liq *LiqFilterer) FilterQuitOperatorAssigned(opts *bind.FilterOpts) (*LiqQuitOperatorAssignedIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "QuitOperatorAssigned")
	if err != nil {
		return nil, err
	}
	return &LiqQuitOperatorAssignedIterator{contract: _Liq.contract, event: "QuitOperatorAssigned", logs: logs, sub: sub}, nil
}

// WatchQuitOperatorAssigned is a free log subscription operation binding the contract event 0xca1faa936bb688b7adf43d587c52f47df688f934011c256ad33a8fdeb51d382e.
//
// Solidity: event QuitOperatorAssigned(uint256 _quitOperatorId, uint256 _operatorId, uint256 _totalAmount)
func (_Liq *LiqFilterer) WatchQuitOperatorAssigned(opts *bind.WatchOpts, sink chan<- *LiqQuitOperatorAssigned) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "QuitOperatorAssigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqQuitOperatorAssigned)
				if err := _Liq.contract.UnpackLog(event, "QuitOperatorAssigned", log); err != nil {
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

// ParseQuitOperatorAssigned is a log parse operation binding the contract event 0xca1faa936bb688b7adf43d587c52f47df688f934011c256ad33a8fdeb51d382e.
//
// Solidity: event QuitOperatorAssigned(uint256 _quitOperatorId, uint256 _operatorId, uint256 _totalAmount)
func (_Liq *LiqFilterer) ParseQuitOperatorAssigned(log types.Log) (*LiqQuitOperatorAssigned, error) {
	event := new(LiqQuitOperatorAssigned)
	if err := _Liq.contract.UnpackLog(event, "QuitOperatorAssigned", log); err != nil {
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

// LiqSlashReceiveIterator is returned from FilterSlashReceive and is used to iterate over the raw logs and unpacked data for SlashReceive events raised by the Liq contract.
type LiqSlashReceiveIterator struct {
	Event *LiqSlashReceive // Event containing the contract specifics and raw log

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
func (it *LiqSlashReceiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiqSlashReceive)
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
		it.Event = new(LiqSlashReceive)
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
func (it *LiqSlashReceiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiqSlashReceiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiqSlashReceive represents a SlashReceive event raised by the Liq contract.
type LiqSlashReceive struct {
	OperatorId    *big.Int
	TokenId       *big.Int
	SlashAmount   *big.Int
	RequirAmounts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSlashReceive is a free log retrieval operation binding the contract event 0xeeba612bf8aaf11b642d483da9db251b732afe21a8035197457492c293eee3f3.
//
// Solidity: event SlashReceive(uint256 _operatorId, uint256 tokenId, uint256 _slashAmount, uint256 _requirAmounts)
func (_Liq *LiqFilterer) FilterSlashReceive(opts *bind.FilterOpts) (*LiqSlashReceiveIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "SlashReceive")
	if err != nil {
		return nil, err
	}
	return &LiqSlashReceiveIterator{contract: _Liq.contract, event: "SlashReceive", logs: logs, sub: sub}, nil
}

// WatchSlashReceive is a free log subscription operation binding the contract event 0xeeba612bf8aaf11b642d483da9db251b732afe21a8035197457492c293eee3f3.
//
// Solidity: event SlashReceive(uint256 _operatorId, uint256 tokenId, uint256 _slashAmount, uint256 _requirAmounts)
func (_Liq *LiqFilterer) WatchSlashReceive(opts *bind.WatchOpts, sink chan<- *LiqSlashReceive) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "SlashReceive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiqSlashReceive)
				if err := _Liq.contract.UnpackLog(event, "SlashReceive", log); err != nil {
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
func (_Liq *LiqFilterer) ParseSlashReceive(log types.Log) (*LiqSlashReceive, error) {
	event := new(LiqSlashReceive)
	if err := _Liq.contract.UnpackLog(event, "SlashReceive", log); err != nil {
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
// Solidity: event ValidatorRegistered(uint256 _operatorId, uint256 _tokenId)
func (_Liq *LiqFilterer) FilterValidatorRegistered(opts *bind.FilterOpts) (*LiqValidatorRegisteredIterator, error) {

	logs, sub, err := _Liq.contract.FilterLogs(opts, "ValidatorRegistered")
	if err != nil {
		return nil, err
	}
	return &LiqValidatorRegisteredIterator{contract: _Liq.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0x700aef92048a2622a7b2403606dfa5c7abaec58382232cdd9ec196907eb44396.
//
// Solidity: event ValidatorRegistered(uint256 _operatorId, uint256 _tokenId)
func (_Liq *LiqFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *LiqValidatorRegistered) (event.Subscription, error) {

	logs, sub, err := _Liq.contract.WatchLogs(opts, "ValidatorRegistered")
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
// Solidity: event ValidatorRegistered(uint256 _operatorId, uint256 _tokenId)
func (_Liq *LiqFilterer) ParseValidatorRegistered(log types.Log) (*LiqValidatorRegistered, error) {
	event := new(LiqValidatorRegistered)
	if err := _Liq.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

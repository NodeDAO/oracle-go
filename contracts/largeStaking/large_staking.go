// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package largeStaking

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

// LargeStakingStakingInfo is an auto generated low-level Go binding around an user-defined struct.
type LargeStakingStakingInfo struct {
	IsELRewardSharing    bool
	StakingId            *big.Int
	OperatorId           *big.Int
	StakingAmount        *big.Int
	AlreadyUsedAmount    *big.Int
	UnstakeRequestAmount *big.Int
	UnstakeAmount        *big.Int
	Owner                common.Address
	ElRewardAddr         common.Address
	WithdrawCredentials  [32]byte
}

// LargeStakingMetaData contains all meta data concerning the LargeStaking contract.
var LargeStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DuplicatePubKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientMargin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReport\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardAddr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardRatio\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawalCredentials\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequireOperatorTrusted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SharedRewardPoolNotOpened\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SharedRewardPoolOpened\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stakingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stakeAmounts\",\"type\":\"uint256\"}],\"name\":\"AppendMigretaStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stakingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"AppendStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldConsensusOracleContractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_consensusOracleContractAddr\",\"type\":\"address\"}],\"name\":\"ConsensusOracleChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldDao\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"DaoAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldDaoElCommissionRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_daoElCommissionRate\",\"type\":\"uint256\"}],\"name\":\"DaoELCommissionRateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stakingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_daoVaultAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_daoRewards\",\"type\":\"uint256\"}],\"name\":\"DaoPrivateRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"daoVaultAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_daoRewards\",\"type\":\"uint256\"}],\"name\":\"DaoSharedRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldDaoVaultAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_daoVaultAddress\",\"type\":\"address\"}],\"name\":\"DaoVaultAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldElRewardFactory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_elRewardFactory\",\"type\":\"address\"}],\"name\":\"ELRewardFactoryChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_daoReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_poolReward\",\"type\":\"uint256\"}],\"name\":\"ELShareingRewardSettle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stakingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_daoReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_poolReward\",\"type\":\"uint256\"}],\"name\":\"ElPrivateRewardSettle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stakingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_unstakeAmount\",\"type\":\"uint256\"}],\"name\":\"FastUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_curStakingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_elRewardAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_withdrawCredentials\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_isELRewardSharing\",\"type\":\"bool\"}],\"name\":\"LargeStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stakingIds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorIds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amounts\",\"type\":\"uint256\"}],\"name\":\"LargeStakingSlash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stakingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LargeUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_curStakingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_elRewardAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_withdrawCredentials\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_isELRewardSharing\",\"type\":\"bool\"}],\"name\":\"MigretaStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldMinStakeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_minStakeAmount\",\"type\":\"uint256\"}],\"name\":\"MinStakeAmountChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldNodeOperatorRegistryContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistryAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorsRegistryChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stakingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorRewards\",\"type\":\"uint256\"}],\"name\":\"OperatorPrivateRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_rewardAddresses\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewardAmounts\",\"type\":\"uint256\"}],\"name\":\"OperatorRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorRewards\",\"type\":\"uint256\"}],\"name\":\"OperatorSharedRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldOperatorSlashContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_operatorSlashContract\",\"type\":\"address\"}],\"name\":\"OperatorSlashChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_elRewardPoolAddr\",\"type\":\"address\"}],\"name\":\"SharedRewardPoolStart\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stakingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"UserRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"ValidatorExitReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stakeingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_pubKey\",\"type\":\"bytes\"}],\"name\":\"ValidatorRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_STAKE_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_withdrawCredentials\",\"type\":\"address\"}],\"name\":\"appendLargeStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_withdrawCredentials\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"_pubKeys\",\"type\":\"bytes[]\"}],\"name\":\"appendMigrateStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_stakingIds\",\"type\":\"uint256[]\"}],\"name\":\"claimRewardsOfDao\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_privatePoolStakingIds\",\"type\":\"uint256[]\"}],\"name\":\"claimRewardsOfOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"claimRewardsOfUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"consensusOracleContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoElCommissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"daoPrivateRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"daoSharedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositContract\",\"outputs\":[{\"internalType\":\"contractIDepositContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"eLSharedRewardSettleInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"valuePerSharePoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"elPrivateRewardPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"elRewardFactory\",\"outputs\":[{\"internalType\":\"contractIELRewardFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"elSharedRewardPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getOperatorValidatorCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingId\",\"type\":\"uint256\"}],\"name\":\"getRewardPoolInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardPoolAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getStakingInfoOfOwner\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isELRewardSharing\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"stakingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alreadyUsedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeRequestAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"elRewardAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawCredentials\",\"type\":\"bytes32\"}],\"internalType\":\"structLargeStaking.StakingInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingId\",\"type\":\"uint256\"}],\"name\":\"getValidatorsOfStakingId\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_withdrawCredentials\",\"type\":\"address\"}],\"name\":\"getWithdrawCredentials\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_daoVaultAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operatorSlashContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_consensusOracleContractAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_elRewardFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_elRewardAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_withdrawCredentials\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isELRewardSharing\",\"type\":\"bool\"}],\"name\":\"largeStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"largeStakings\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isELRewardSharing\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"stakingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alreadyUsedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeRequestAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"elRewardAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawCredentials\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"largeUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_elRewardAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_withdrawCredentials\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isELRewardSharing\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"_pubKeys\",\"type\":\"bytes[]\"}],\"name\":\"migrateStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeOperatorRegistryContract\",\"outputs\":[{\"internalType\":\"contractINodeOperatorsRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorPrivateRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorSharedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorSlashContract\",\"outputs\":[{\"internalType\":\"contractIOperatorSlash\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"stakingId\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"internalType\":\"structCLStakingExitInfo[]\",\"name\":\"_clStakingExitInfo\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"stakingId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"slashAmount\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"internalType\":\"structCLStakingSlashInfo[]\",\"name\":\"_clStakingSlashInfo\",\"type\":\"tuple[]\"}],\"name\":\"reportCLStakingData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingId\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"userReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_daoVaultAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_daoElCommissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_MIN_STAKE_AMOUNT\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_consensusOracleContractAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_elRewardFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operatorSlashContract\",\"type\":\"address\"}],\"name\":\"setLargeStakingSetting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingId\",\"type\":\"uint256\"}],\"name\":\"settleElPrivateReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"settleElSharedReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"startupSharedRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLargeStakingCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"unclaimedPrivateRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"unclaimedSharedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"validatorExitReportBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"validatorOfStaking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"validatorRegisterBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"validatorSlashAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"valuePerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LargeStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use LargeStakingMetaData.ABI instead.
var LargeStakingABI = LargeStakingMetaData.ABI

// LargeStaking is an auto generated Go binding around an Ethereum contract.
type LargeStaking struct {
	LargeStakingCaller     // Read-only binding to the contract
	LargeStakingTransactor // Write-only binding to the contract
	LargeStakingFilterer   // Log filterer for contract events
}

// LargeStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type LargeStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LargeStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LargeStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LargeStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LargeStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LargeStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LargeStakingSession struct {
	Contract     *LargeStaking     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LargeStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LargeStakingCallerSession struct {
	Contract *LargeStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LargeStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LargeStakingTransactorSession struct {
	Contract     *LargeStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LargeStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type LargeStakingRaw struct {
	Contract *LargeStaking // Generic contract binding to access the raw methods on
}

// LargeStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LargeStakingCallerRaw struct {
	Contract *LargeStakingCaller // Generic read-only contract binding to access the raw methods on
}

// LargeStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LargeStakingTransactorRaw struct {
	Contract *LargeStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLargeStaking creates a new instance of LargeStaking, bound to a specific deployed contract.
func NewLargeStaking(address common.Address, backend bind.ContractBackend) (*LargeStaking, error) {
	contract, err := bindLargeStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LargeStaking{LargeStakingCaller: LargeStakingCaller{contract: contract}, LargeStakingTransactor: LargeStakingTransactor{contract: contract}, LargeStakingFilterer: LargeStakingFilterer{contract: contract}}, nil
}

// NewLargeStakingCaller creates a new read-only instance of LargeStaking, bound to a specific deployed contract.
func NewLargeStakingCaller(address common.Address, caller bind.ContractCaller) (*LargeStakingCaller, error) {
	contract, err := bindLargeStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LargeStakingCaller{contract: contract}, nil
}

// NewLargeStakingTransactor creates a new write-only instance of LargeStaking, bound to a specific deployed contract.
func NewLargeStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*LargeStakingTransactor, error) {
	contract, err := bindLargeStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LargeStakingTransactor{contract: contract}, nil
}

// NewLargeStakingFilterer creates a new log filterer instance of LargeStaking, bound to a specific deployed contract.
func NewLargeStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*LargeStakingFilterer, error) {
	contract, err := bindLargeStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LargeStakingFilterer{contract: contract}, nil
}

// bindLargeStaking binds a generic wrapper to an already deployed contract.
func bindLargeStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LargeStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LargeStaking *LargeStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LargeStaking.Contract.LargeStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LargeStaking *LargeStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LargeStaking.Contract.LargeStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LargeStaking *LargeStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LargeStaking.Contract.LargeStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LargeStaking *LargeStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LargeStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LargeStaking *LargeStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LargeStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LargeStaking *LargeStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LargeStaking.Contract.contract.Transact(opts, method, params...)
}

// MINSTAKEAMOUNT is a free data retrieval call binding the contract method 0x27ed7188.
//
// Solidity: function MIN_STAKE_AMOUNT() view returns(uint256)
func (_LargeStaking *LargeStakingCaller) MINSTAKEAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "MIN_STAKE_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINSTAKEAMOUNT is a free data retrieval call binding the contract method 0x27ed7188.
//
// Solidity: function MIN_STAKE_AMOUNT() view returns(uint256)
func (_LargeStaking *LargeStakingSession) MINSTAKEAMOUNT() (*big.Int, error) {
	return _LargeStaking.Contract.MINSTAKEAMOUNT(&_LargeStaking.CallOpts)
}

// MINSTAKEAMOUNT is a free data retrieval call binding the contract method 0x27ed7188.
//
// Solidity: function MIN_STAKE_AMOUNT() view returns(uint256)
func (_LargeStaking *LargeStakingCallerSession) MINSTAKEAMOUNT() (*big.Int, error) {
	return _LargeStaking.Contract.MINSTAKEAMOUNT(&_LargeStaking.CallOpts)
}

// ConsensusOracleContractAddr is a free data retrieval call binding the contract method 0xa10e3a1c.
//
// Solidity: function consensusOracleContractAddr() view returns(address)
func (_LargeStaking *LargeStakingCaller) ConsensusOracleContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "consensusOracleContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConsensusOracleContractAddr is a free data retrieval call binding the contract method 0xa10e3a1c.
//
// Solidity: function consensusOracleContractAddr() view returns(address)
func (_LargeStaking *LargeStakingSession) ConsensusOracleContractAddr() (common.Address, error) {
	return _LargeStaking.Contract.ConsensusOracleContractAddr(&_LargeStaking.CallOpts)
}

// ConsensusOracleContractAddr is a free data retrieval call binding the contract method 0xa10e3a1c.
//
// Solidity: function consensusOracleContractAddr() view returns(address)
func (_LargeStaking *LargeStakingCallerSession) ConsensusOracleContractAddr() (common.Address, error) {
	return _LargeStaking.Contract.ConsensusOracleContractAddr(&_LargeStaking.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_LargeStaking *LargeStakingCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_LargeStaking *LargeStakingSession) Dao() (common.Address, error) {
	return _LargeStaking.Contract.Dao(&_LargeStaking.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_LargeStaking *LargeStakingCallerSession) Dao() (common.Address, error) {
	return _LargeStaking.Contract.Dao(&_LargeStaking.CallOpts)
}

// DaoElCommissionRate is a free data retrieval call binding the contract method 0x94564043.
//
// Solidity: function daoElCommissionRate() view returns(uint256)
func (_LargeStaking *LargeStakingCaller) DaoElCommissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "daoElCommissionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DaoElCommissionRate is a free data retrieval call binding the contract method 0x94564043.
//
// Solidity: function daoElCommissionRate() view returns(uint256)
func (_LargeStaking *LargeStakingSession) DaoElCommissionRate() (*big.Int, error) {
	return _LargeStaking.Contract.DaoElCommissionRate(&_LargeStaking.CallOpts)
}

// DaoElCommissionRate is a free data retrieval call binding the contract method 0x94564043.
//
// Solidity: function daoElCommissionRate() view returns(uint256)
func (_LargeStaking *LargeStakingCallerSession) DaoElCommissionRate() (*big.Int, error) {
	return _LargeStaking.Contract.DaoElCommissionRate(&_LargeStaking.CallOpts)
}

// DaoPrivateRewards is a free data retrieval call binding the contract method 0x50fe19ff.
//
// Solidity: function daoPrivateRewards(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingCaller) DaoPrivateRewards(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "daoPrivateRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DaoPrivateRewards is a free data retrieval call binding the contract method 0x50fe19ff.
//
// Solidity: function daoPrivateRewards(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingSession) DaoPrivateRewards(arg0 *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.DaoPrivateRewards(&_LargeStaking.CallOpts, arg0)
}

// DaoPrivateRewards is a free data retrieval call binding the contract method 0x50fe19ff.
//
// Solidity: function daoPrivateRewards(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingCallerSession) DaoPrivateRewards(arg0 *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.DaoPrivateRewards(&_LargeStaking.CallOpts, arg0)
}

// DaoSharedRewards is a free data retrieval call binding the contract method 0xea3899ca.
//
// Solidity: function daoSharedRewards(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingCaller) DaoSharedRewards(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "daoSharedRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DaoSharedRewards is a free data retrieval call binding the contract method 0xea3899ca.
//
// Solidity: function daoSharedRewards(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingSession) DaoSharedRewards(arg0 *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.DaoSharedRewards(&_LargeStaking.CallOpts, arg0)
}

// DaoSharedRewards is a free data retrieval call binding the contract method 0xea3899ca.
//
// Solidity: function daoSharedRewards(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingCallerSession) DaoSharedRewards(arg0 *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.DaoSharedRewards(&_LargeStaking.CallOpts, arg0)
}

// DaoVaultAddress is a free data retrieval call binding the contract method 0x3d6a3844.
//
// Solidity: function daoVaultAddress() view returns(address)
func (_LargeStaking *LargeStakingCaller) DaoVaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "daoVaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaoVaultAddress is a free data retrieval call binding the contract method 0x3d6a3844.
//
// Solidity: function daoVaultAddress() view returns(address)
func (_LargeStaking *LargeStakingSession) DaoVaultAddress() (common.Address, error) {
	return _LargeStaking.Contract.DaoVaultAddress(&_LargeStaking.CallOpts)
}

// DaoVaultAddress is a free data retrieval call binding the contract method 0x3d6a3844.
//
// Solidity: function daoVaultAddress() view returns(address)
func (_LargeStaking *LargeStakingCallerSession) DaoVaultAddress() (common.Address, error) {
	return _LargeStaking.Contract.DaoVaultAddress(&_LargeStaking.CallOpts)
}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_LargeStaking *LargeStakingCaller) DepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "depositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_LargeStaking *LargeStakingSession) DepositContract() (common.Address, error) {
	return _LargeStaking.Contract.DepositContract(&_LargeStaking.CallOpts)
}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_LargeStaking *LargeStakingCallerSession) DepositContract() (common.Address, error) {
	return _LargeStaking.Contract.DepositContract(&_LargeStaking.CallOpts)
}

// ELSharedRewardSettleInfo is a free data retrieval call binding the contract method 0x2cfe6a25.
//
// Solidity: function eLSharedRewardSettleInfo(uint256 ) view returns(uint256 valuePerSharePoint, uint256 rewardBalance)
func (_LargeStaking *LargeStakingCaller) ELSharedRewardSettleInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ValuePerSharePoint *big.Int
	RewardBalance      *big.Int
}, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "eLSharedRewardSettleInfo", arg0)

	outstruct := new(struct {
		ValuePerSharePoint *big.Int
		RewardBalance      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValuePerSharePoint = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ELSharedRewardSettleInfo is a free data retrieval call binding the contract method 0x2cfe6a25.
//
// Solidity: function eLSharedRewardSettleInfo(uint256 ) view returns(uint256 valuePerSharePoint, uint256 rewardBalance)
func (_LargeStaking *LargeStakingSession) ELSharedRewardSettleInfo(arg0 *big.Int) (struct {
	ValuePerSharePoint *big.Int
	RewardBalance      *big.Int
}, error) {
	return _LargeStaking.Contract.ELSharedRewardSettleInfo(&_LargeStaking.CallOpts, arg0)
}

// ELSharedRewardSettleInfo is a free data retrieval call binding the contract method 0x2cfe6a25.
//
// Solidity: function eLSharedRewardSettleInfo(uint256 ) view returns(uint256 valuePerSharePoint, uint256 rewardBalance)
func (_LargeStaking *LargeStakingCallerSession) ELSharedRewardSettleInfo(arg0 *big.Int) (struct {
	ValuePerSharePoint *big.Int
	RewardBalance      *big.Int
}, error) {
	return _LargeStaking.Contract.ELSharedRewardSettleInfo(&_LargeStaking.CallOpts, arg0)
}

// ElPrivateRewardPool is a free data retrieval call binding the contract method 0x1e920634.
//
// Solidity: function elPrivateRewardPool(uint256 ) view returns(address)
func (_LargeStaking *LargeStakingCaller) ElPrivateRewardPool(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "elPrivateRewardPool", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ElPrivateRewardPool is a free data retrieval call binding the contract method 0x1e920634.
//
// Solidity: function elPrivateRewardPool(uint256 ) view returns(address)
func (_LargeStaking *LargeStakingSession) ElPrivateRewardPool(arg0 *big.Int) (common.Address, error) {
	return _LargeStaking.Contract.ElPrivateRewardPool(&_LargeStaking.CallOpts, arg0)
}

// ElPrivateRewardPool is a free data retrieval call binding the contract method 0x1e920634.
//
// Solidity: function elPrivateRewardPool(uint256 ) view returns(address)
func (_LargeStaking *LargeStakingCallerSession) ElPrivateRewardPool(arg0 *big.Int) (common.Address, error) {
	return _LargeStaking.Contract.ElPrivateRewardPool(&_LargeStaking.CallOpts, arg0)
}

// ElRewardFactory is a free data retrieval call binding the contract method 0xa92536ba.
//
// Solidity: function elRewardFactory() view returns(address)
func (_LargeStaking *LargeStakingCaller) ElRewardFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "elRewardFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ElRewardFactory is a free data retrieval call binding the contract method 0xa92536ba.
//
// Solidity: function elRewardFactory() view returns(address)
func (_LargeStaking *LargeStakingSession) ElRewardFactory() (common.Address, error) {
	return _LargeStaking.Contract.ElRewardFactory(&_LargeStaking.CallOpts)
}

// ElRewardFactory is a free data retrieval call binding the contract method 0xa92536ba.
//
// Solidity: function elRewardFactory() view returns(address)
func (_LargeStaking *LargeStakingCallerSession) ElRewardFactory() (common.Address, error) {
	return _LargeStaking.Contract.ElRewardFactory(&_LargeStaking.CallOpts)
}

// ElSharedRewardPool is a free data retrieval call binding the contract method 0xc589290e.
//
// Solidity: function elSharedRewardPool(uint256 ) view returns(address)
func (_LargeStaking *LargeStakingCaller) ElSharedRewardPool(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "elSharedRewardPool", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ElSharedRewardPool is a free data retrieval call binding the contract method 0xc589290e.
//
// Solidity: function elSharedRewardPool(uint256 ) view returns(address)
func (_LargeStaking *LargeStakingSession) ElSharedRewardPool(arg0 *big.Int) (common.Address, error) {
	return _LargeStaking.Contract.ElSharedRewardPool(&_LargeStaking.CallOpts, arg0)
}

// ElSharedRewardPool is a free data retrieval call binding the contract method 0xc589290e.
//
// Solidity: function elSharedRewardPool(uint256 ) view returns(address)
func (_LargeStaking *LargeStakingCallerSession) ElSharedRewardPool(arg0 *big.Int) (common.Address, error) {
	return _LargeStaking.Contract.ElSharedRewardPool(&_LargeStaking.CallOpts, arg0)
}

// GetOperatorValidatorCounts is a free data retrieval call binding the contract method 0x68cce002.
//
// Solidity: function getOperatorValidatorCounts(uint256 _operatorId) view returns(uint256)
func (_LargeStaking *LargeStakingCaller) GetOperatorValidatorCounts(opts *bind.CallOpts, _operatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "getOperatorValidatorCounts", _operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorValidatorCounts is a free data retrieval call binding the contract method 0x68cce002.
//
// Solidity: function getOperatorValidatorCounts(uint256 _operatorId) view returns(uint256)
func (_LargeStaking *LargeStakingSession) GetOperatorValidatorCounts(_operatorId *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.GetOperatorValidatorCounts(&_LargeStaking.CallOpts, _operatorId)
}

// GetOperatorValidatorCounts is a free data retrieval call binding the contract method 0x68cce002.
//
// Solidity: function getOperatorValidatorCounts(uint256 _operatorId) view returns(uint256)
func (_LargeStaking *LargeStakingCallerSession) GetOperatorValidatorCounts(_operatorId *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.GetOperatorValidatorCounts(&_LargeStaking.CallOpts, _operatorId)
}

// GetRewardPoolInfo is a free data retrieval call binding the contract method 0x46ee293e.
//
// Solidity: function getRewardPoolInfo(uint256 _stakingId) view returns(uint256 operatorId, address rewardPoolAddr, uint256 rewards)
func (_LargeStaking *LargeStakingCaller) GetRewardPoolInfo(opts *bind.CallOpts, _stakingId *big.Int) (struct {
	OperatorId     *big.Int
	RewardPoolAddr common.Address
	Rewards        *big.Int
}, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "getRewardPoolInfo", _stakingId)

	outstruct := new(struct {
		OperatorId     *big.Int
		RewardPoolAddr common.Address
		Rewards        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OperatorId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardPoolAddr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Rewards = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRewardPoolInfo is a free data retrieval call binding the contract method 0x46ee293e.
//
// Solidity: function getRewardPoolInfo(uint256 _stakingId) view returns(uint256 operatorId, address rewardPoolAddr, uint256 rewards)
func (_LargeStaking *LargeStakingSession) GetRewardPoolInfo(_stakingId *big.Int) (struct {
	OperatorId     *big.Int
	RewardPoolAddr common.Address
	Rewards        *big.Int
}, error) {
	return _LargeStaking.Contract.GetRewardPoolInfo(&_LargeStaking.CallOpts, _stakingId)
}

// GetRewardPoolInfo is a free data retrieval call binding the contract method 0x46ee293e.
//
// Solidity: function getRewardPoolInfo(uint256 _stakingId) view returns(uint256 operatorId, address rewardPoolAddr, uint256 rewards)
func (_LargeStaking *LargeStakingCallerSession) GetRewardPoolInfo(_stakingId *big.Int) (struct {
	OperatorId     *big.Int
	RewardPoolAddr common.Address
	Rewards        *big.Int
}, error) {
	return _LargeStaking.Contract.GetRewardPoolInfo(&_LargeStaking.CallOpts, _stakingId)
}

// GetStakingInfoOfOwner is a free data retrieval call binding the contract method 0x188ef00f.
//
// Solidity: function getStakingInfoOfOwner(address _owner) view returns((bool,uint256,uint256,uint256,uint256,uint256,uint256,address,address,bytes32)[])
func (_LargeStaking *LargeStakingCaller) GetStakingInfoOfOwner(opts *bind.CallOpts, _owner common.Address) ([]LargeStakingStakingInfo, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "getStakingInfoOfOwner", _owner)

	if err != nil {
		return *new([]LargeStakingStakingInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]LargeStakingStakingInfo)).(*[]LargeStakingStakingInfo)

	return out0, err

}

// GetStakingInfoOfOwner is a free data retrieval call binding the contract method 0x188ef00f.
//
// Solidity: function getStakingInfoOfOwner(address _owner) view returns((bool,uint256,uint256,uint256,uint256,uint256,uint256,address,address,bytes32)[])
func (_LargeStaking *LargeStakingSession) GetStakingInfoOfOwner(_owner common.Address) ([]LargeStakingStakingInfo, error) {
	return _LargeStaking.Contract.GetStakingInfoOfOwner(&_LargeStaking.CallOpts, _owner)
}

// GetStakingInfoOfOwner is a free data retrieval call binding the contract method 0x188ef00f.
//
// Solidity: function getStakingInfoOfOwner(address _owner) view returns((bool,uint256,uint256,uint256,uint256,uint256,uint256,address,address,bytes32)[])
func (_LargeStaking *LargeStakingCallerSession) GetStakingInfoOfOwner(_owner common.Address) ([]LargeStakingStakingInfo, error) {
	return _LargeStaking.Contract.GetStakingInfoOfOwner(&_LargeStaking.CallOpts, _owner)
}

// GetValidatorsOfStakingId is a free data retrieval call binding the contract method 0x7bd1f1fa.
//
// Solidity: function getValidatorsOfStakingId(uint256 _stakingId) view returns(bytes[])
func (_LargeStaking *LargeStakingCaller) GetValidatorsOfStakingId(opts *bind.CallOpts, _stakingId *big.Int) ([][]byte, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "getValidatorsOfStakingId", _stakingId)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetValidatorsOfStakingId is a free data retrieval call binding the contract method 0x7bd1f1fa.
//
// Solidity: function getValidatorsOfStakingId(uint256 _stakingId) view returns(bytes[])
func (_LargeStaking *LargeStakingSession) GetValidatorsOfStakingId(_stakingId *big.Int) ([][]byte, error) {
	return _LargeStaking.Contract.GetValidatorsOfStakingId(&_LargeStaking.CallOpts, _stakingId)
}

// GetValidatorsOfStakingId is a free data retrieval call binding the contract method 0x7bd1f1fa.
//
// Solidity: function getValidatorsOfStakingId(uint256 _stakingId) view returns(bytes[])
func (_LargeStaking *LargeStakingCallerSession) GetValidatorsOfStakingId(_stakingId *big.Int) ([][]byte, error) {
	return _LargeStaking.Contract.GetValidatorsOfStakingId(&_LargeStaking.CallOpts, _stakingId)
}

// GetWithdrawCredentials is a free data retrieval call binding the contract method 0xb24b4b83.
//
// Solidity: function getWithdrawCredentials(address _withdrawCredentials) pure returns(bytes32)
func (_LargeStaking *LargeStakingCaller) GetWithdrawCredentials(opts *bind.CallOpts, _withdrawCredentials common.Address) ([32]byte, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "getWithdrawCredentials", _withdrawCredentials)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetWithdrawCredentials is a free data retrieval call binding the contract method 0xb24b4b83.
//
// Solidity: function getWithdrawCredentials(address _withdrawCredentials) pure returns(bytes32)
func (_LargeStaking *LargeStakingSession) GetWithdrawCredentials(_withdrawCredentials common.Address) ([32]byte, error) {
	return _LargeStaking.Contract.GetWithdrawCredentials(&_LargeStaking.CallOpts, _withdrawCredentials)
}

// GetWithdrawCredentials is a free data retrieval call binding the contract method 0xb24b4b83.
//
// Solidity: function getWithdrawCredentials(address _withdrawCredentials) pure returns(bytes32)
func (_LargeStaking *LargeStakingCallerSession) GetWithdrawCredentials(_withdrawCredentials common.Address) ([32]byte, error) {
	return _LargeStaking.Contract.GetWithdrawCredentials(&_LargeStaking.CallOpts, _withdrawCredentials)
}

// LargeStakings is a free data retrieval call binding the contract method 0x7d54dcf4.
//
// Solidity: function largeStakings(uint256 ) view returns(bool isELRewardSharing, uint256 stakingId, uint256 operatorId, uint256 stakingAmount, uint256 alreadyUsedAmount, uint256 unstakeRequestAmount, uint256 unstakeAmount, address owner, address elRewardAddr, bytes32 withdrawCredentials)
func (_LargeStaking *LargeStakingCaller) LargeStakings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	IsELRewardSharing    bool
	StakingId            *big.Int
	OperatorId           *big.Int
	StakingAmount        *big.Int
	AlreadyUsedAmount    *big.Int
	UnstakeRequestAmount *big.Int
	UnstakeAmount        *big.Int
	Owner                common.Address
	ElRewardAddr         common.Address
	WithdrawCredentials  [32]byte
}, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "largeStakings", arg0)

	outstruct := new(struct {
		IsELRewardSharing    bool
		StakingId            *big.Int
		OperatorId           *big.Int
		StakingAmount        *big.Int
		AlreadyUsedAmount    *big.Int
		UnstakeRequestAmount *big.Int
		UnstakeAmount        *big.Int
		Owner                common.Address
		ElRewardAddr         common.Address
		WithdrawCredentials  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsELRewardSharing = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.StakingId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.OperatorId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StakingAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AlreadyUsedAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.UnstakeRequestAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.UnstakeAmount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.ElRewardAddr = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.WithdrawCredentials = *abi.ConvertType(out[9], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// LargeStakings is a free data retrieval call binding the contract method 0x7d54dcf4.
//
// Solidity: function largeStakings(uint256 ) view returns(bool isELRewardSharing, uint256 stakingId, uint256 operatorId, uint256 stakingAmount, uint256 alreadyUsedAmount, uint256 unstakeRequestAmount, uint256 unstakeAmount, address owner, address elRewardAddr, bytes32 withdrawCredentials)
func (_LargeStaking *LargeStakingSession) LargeStakings(arg0 *big.Int) (struct {
	IsELRewardSharing    bool
	StakingId            *big.Int
	OperatorId           *big.Int
	StakingAmount        *big.Int
	AlreadyUsedAmount    *big.Int
	UnstakeRequestAmount *big.Int
	UnstakeAmount        *big.Int
	Owner                common.Address
	ElRewardAddr         common.Address
	WithdrawCredentials  [32]byte
}, error) {
	return _LargeStaking.Contract.LargeStakings(&_LargeStaking.CallOpts, arg0)
}

// LargeStakings is a free data retrieval call binding the contract method 0x7d54dcf4.
//
// Solidity: function largeStakings(uint256 ) view returns(bool isELRewardSharing, uint256 stakingId, uint256 operatorId, uint256 stakingAmount, uint256 alreadyUsedAmount, uint256 unstakeRequestAmount, uint256 unstakeAmount, address owner, address elRewardAddr, bytes32 withdrawCredentials)
func (_LargeStaking *LargeStakingCallerSession) LargeStakings(arg0 *big.Int) (struct {
	IsELRewardSharing    bool
	StakingId            *big.Int
	OperatorId           *big.Int
	StakingAmount        *big.Int
	AlreadyUsedAmount    *big.Int
	UnstakeRequestAmount *big.Int
	UnstakeAmount        *big.Int
	Owner                common.Address
	ElRewardAddr         common.Address
	WithdrawCredentials  [32]byte
}, error) {
	return _LargeStaking.Contract.LargeStakings(&_LargeStaking.CallOpts, arg0)
}

// NodeOperatorRegistryContract is a free data retrieval call binding the contract method 0x9a42e0ba.
//
// Solidity: function nodeOperatorRegistryContract() view returns(address)
func (_LargeStaking *LargeStakingCaller) NodeOperatorRegistryContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "nodeOperatorRegistryContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeOperatorRegistryContract is a free data retrieval call binding the contract method 0x9a42e0ba.
//
// Solidity: function nodeOperatorRegistryContract() view returns(address)
func (_LargeStaking *LargeStakingSession) NodeOperatorRegistryContract() (common.Address, error) {
	return _LargeStaking.Contract.NodeOperatorRegistryContract(&_LargeStaking.CallOpts)
}

// NodeOperatorRegistryContract is a free data retrieval call binding the contract method 0x9a42e0ba.
//
// Solidity: function nodeOperatorRegistryContract() view returns(address)
func (_LargeStaking *LargeStakingCallerSession) NodeOperatorRegistryContract() (common.Address, error) {
	return _LargeStaking.Contract.NodeOperatorRegistryContract(&_LargeStaking.CallOpts)
}

// OperatorPrivateRewards is a free data retrieval call binding the contract method 0xc2f1cccd.
//
// Solidity: function operatorPrivateRewards(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingCaller) OperatorPrivateRewards(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "operatorPrivateRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorPrivateRewards is a free data retrieval call binding the contract method 0xc2f1cccd.
//
// Solidity: function operatorPrivateRewards(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingSession) OperatorPrivateRewards(arg0 *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.OperatorPrivateRewards(&_LargeStaking.CallOpts, arg0)
}

// OperatorPrivateRewards is a free data retrieval call binding the contract method 0xc2f1cccd.
//
// Solidity: function operatorPrivateRewards(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingCallerSession) OperatorPrivateRewards(arg0 *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.OperatorPrivateRewards(&_LargeStaking.CallOpts, arg0)
}

// OperatorSharedRewards is a free data retrieval call binding the contract method 0xeb234d8e.
//
// Solidity: function operatorSharedRewards(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingCaller) OperatorSharedRewards(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "operatorSharedRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorSharedRewards is a free data retrieval call binding the contract method 0xeb234d8e.
//
// Solidity: function operatorSharedRewards(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingSession) OperatorSharedRewards(arg0 *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.OperatorSharedRewards(&_LargeStaking.CallOpts, arg0)
}

// OperatorSharedRewards is a free data retrieval call binding the contract method 0xeb234d8e.
//
// Solidity: function operatorSharedRewards(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingCallerSession) OperatorSharedRewards(arg0 *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.OperatorSharedRewards(&_LargeStaking.CallOpts, arg0)
}

// OperatorSlashContract is a free data retrieval call binding the contract method 0x0c2559fd.
//
// Solidity: function operatorSlashContract() view returns(address)
func (_LargeStaking *LargeStakingCaller) OperatorSlashContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "operatorSlashContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorSlashContract is a free data retrieval call binding the contract method 0x0c2559fd.
//
// Solidity: function operatorSlashContract() view returns(address)
func (_LargeStaking *LargeStakingSession) OperatorSlashContract() (common.Address, error) {
	return _LargeStaking.Contract.OperatorSlashContract(&_LargeStaking.CallOpts)
}

// OperatorSlashContract is a free data retrieval call binding the contract method 0x0c2559fd.
//
// Solidity: function operatorSlashContract() view returns(address)
func (_LargeStaking *LargeStakingCallerSession) OperatorSlashContract() (common.Address, error) {
	return _LargeStaking.Contract.OperatorSlashContract(&_LargeStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LargeStaking *LargeStakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LargeStaking *LargeStakingSession) Owner() (common.Address, error) {
	return _LargeStaking.Contract.Owner(&_LargeStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LargeStaking *LargeStakingCallerSession) Owner() (common.Address, error) {
	return _LargeStaking.Contract.Owner(&_LargeStaking.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LargeStaking *LargeStakingCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LargeStaking *LargeStakingSession) ProxiableUUID() ([32]byte, error) {
	return _LargeStaking.Contract.ProxiableUUID(&_LargeStaking.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LargeStaking *LargeStakingCallerSession) ProxiableUUID() ([32]byte, error) {
	return _LargeStaking.Contract.ProxiableUUID(&_LargeStaking.CallOpts)
}

// Reward is a free data retrieval call binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 _stakingId) view returns(uint256 userReward)
func (_LargeStaking *LargeStakingCaller) Reward(opts *bind.CallOpts, _stakingId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "reward", _stakingId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reward is a free data retrieval call binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 _stakingId) view returns(uint256 userReward)
func (_LargeStaking *LargeStakingSession) Reward(_stakingId *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.Reward(&_LargeStaking.CallOpts, _stakingId)
}

// Reward is a free data retrieval call binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 _stakingId) view returns(uint256 userReward)
func (_LargeStaking *LargeStakingCallerSession) Reward(_stakingId *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.Reward(&_LargeStaking.CallOpts, _stakingId)
}

// TotalLargeStakingCounts is a free data retrieval call binding the contract method 0x4d457f3b.
//
// Solidity: function totalLargeStakingCounts() view returns(uint256)
func (_LargeStaking *LargeStakingCaller) TotalLargeStakingCounts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "totalLargeStakingCounts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLargeStakingCounts is a free data retrieval call binding the contract method 0x4d457f3b.
//
// Solidity: function totalLargeStakingCounts() view returns(uint256)
func (_LargeStaking *LargeStakingSession) TotalLargeStakingCounts() (*big.Int, error) {
	return _LargeStaking.Contract.TotalLargeStakingCounts(&_LargeStaking.CallOpts)
}

// TotalLargeStakingCounts is a free data retrieval call binding the contract method 0x4d457f3b.
//
// Solidity: function totalLargeStakingCounts() view returns(uint256)
func (_LargeStaking *LargeStakingCallerSession) TotalLargeStakingCounts() (*big.Int, error) {
	return _LargeStaking.Contract.TotalLargeStakingCounts(&_LargeStaking.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x13f2dad0.
//
// Solidity: function totalShares(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingCaller) TotalShares(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "totalShares", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x13f2dad0.
//
// Solidity: function totalShares(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingSession) TotalShares(arg0 *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.TotalShares(&_LargeStaking.CallOpts, arg0)
}

// TotalShares is a free data retrieval call binding the contract method 0x13f2dad0.
//
// Solidity: function totalShares(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingCallerSession) TotalShares(arg0 *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.TotalShares(&_LargeStaking.CallOpts, arg0)
}

// UnclaimedPrivateRewards is a free data retrieval call binding the contract method 0xfb144f15.
//
// Solidity: function unclaimedPrivateRewards(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingCaller) UnclaimedPrivateRewards(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "unclaimedPrivateRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnclaimedPrivateRewards is a free data retrieval call binding the contract method 0xfb144f15.
//
// Solidity: function unclaimedPrivateRewards(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingSession) UnclaimedPrivateRewards(arg0 *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.UnclaimedPrivateRewards(&_LargeStaking.CallOpts, arg0)
}

// UnclaimedPrivateRewards is a free data retrieval call binding the contract method 0xfb144f15.
//
// Solidity: function unclaimedPrivateRewards(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingCallerSession) UnclaimedPrivateRewards(arg0 *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.UnclaimedPrivateRewards(&_LargeStaking.CallOpts, arg0)
}

// UnclaimedSharedRewards is a free data retrieval call binding the contract method 0x1822edc2.
//
// Solidity: function unclaimedSharedRewards(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingCaller) UnclaimedSharedRewards(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "unclaimedSharedRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnclaimedSharedRewards is a free data retrieval call binding the contract method 0x1822edc2.
//
// Solidity: function unclaimedSharedRewards(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingSession) UnclaimedSharedRewards(arg0 *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.UnclaimedSharedRewards(&_LargeStaking.CallOpts, arg0)
}

// UnclaimedSharedRewards is a free data retrieval call binding the contract method 0x1822edc2.
//
// Solidity: function unclaimedSharedRewards(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingCallerSession) UnclaimedSharedRewards(arg0 *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.UnclaimedSharedRewards(&_LargeStaking.CallOpts, arg0)
}

// ValidatorExitReportBlock is a free data retrieval call binding the contract method 0xd8b9c8de.
//
// Solidity: function validatorExitReportBlock(bytes ) view returns(uint256)
func (_LargeStaking *LargeStakingCaller) ValidatorExitReportBlock(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "validatorExitReportBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorExitReportBlock is a free data retrieval call binding the contract method 0xd8b9c8de.
//
// Solidity: function validatorExitReportBlock(bytes ) view returns(uint256)
func (_LargeStaking *LargeStakingSession) ValidatorExitReportBlock(arg0 []byte) (*big.Int, error) {
	return _LargeStaking.Contract.ValidatorExitReportBlock(&_LargeStaking.CallOpts, arg0)
}

// ValidatorExitReportBlock is a free data retrieval call binding the contract method 0xd8b9c8de.
//
// Solidity: function validatorExitReportBlock(bytes ) view returns(uint256)
func (_LargeStaking *LargeStakingCallerSession) ValidatorExitReportBlock(arg0 []byte) (*big.Int, error) {
	return _LargeStaking.Contract.ValidatorExitReportBlock(&_LargeStaking.CallOpts, arg0)
}

// ValidatorOfStaking is a free data retrieval call binding the contract method 0xe9871c91.
//
// Solidity: function validatorOfStaking(bytes ) view returns(uint256)
func (_LargeStaking *LargeStakingCaller) ValidatorOfStaking(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "validatorOfStaking", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorOfStaking is a free data retrieval call binding the contract method 0xe9871c91.
//
// Solidity: function validatorOfStaking(bytes ) view returns(uint256)
func (_LargeStaking *LargeStakingSession) ValidatorOfStaking(arg0 []byte) (*big.Int, error) {
	return _LargeStaking.Contract.ValidatorOfStaking(&_LargeStaking.CallOpts, arg0)
}

// ValidatorOfStaking is a free data retrieval call binding the contract method 0xe9871c91.
//
// Solidity: function validatorOfStaking(bytes ) view returns(uint256)
func (_LargeStaking *LargeStakingCallerSession) ValidatorOfStaking(arg0 []byte) (*big.Int, error) {
	return _LargeStaking.Contract.ValidatorOfStaking(&_LargeStaking.CallOpts, arg0)
}

// ValidatorRegisterBlock is a free data retrieval call binding the contract method 0x4488605d.
//
// Solidity: function validatorRegisterBlock(bytes ) view returns(uint256)
func (_LargeStaking *LargeStakingCaller) ValidatorRegisterBlock(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "validatorRegisterBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorRegisterBlock is a free data retrieval call binding the contract method 0x4488605d.
//
// Solidity: function validatorRegisterBlock(bytes ) view returns(uint256)
func (_LargeStaking *LargeStakingSession) ValidatorRegisterBlock(arg0 []byte) (*big.Int, error) {
	return _LargeStaking.Contract.ValidatorRegisterBlock(&_LargeStaking.CallOpts, arg0)
}

// ValidatorRegisterBlock is a free data retrieval call binding the contract method 0x4488605d.
//
// Solidity: function validatorRegisterBlock(bytes ) view returns(uint256)
func (_LargeStaking *LargeStakingCallerSession) ValidatorRegisterBlock(arg0 []byte) (*big.Int, error) {
	return _LargeStaking.Contract.ValidatorRegisterBlock(&_LargeStaking.CallOpts, arg0)
}

// ValidatorSlashAmount is a free data retrieval call binding the contract method 0x35062a35.
//
// Solidity: function validatorSlashAmount(bytes ) view returns(uint256)
func (_LargeStaking *LargeStakingCaller) ValidatorSlashAmount(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "validatorSlashAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorSlashAmount is a free data retrieval call binding the contract method 0x35062a35.
//
// Solidity: function validatorSlashAmount(bytes ) view returns(uint256)
func (_LargeStaking *LargeStakingSession) ValidatorSlashAmount(arg0 []byte) (*big.Int, error) {
	return _LargeStaking.Contract.ValidatorSlashAmount(&_LargeStaking.CallOpts, arg0)
}

// ValidatorSlashAmount is a free data retrieval call binding the contract method 0x35062a35.
//
// Solidity: function validatorSlashAmount(bytes ) view returns(uint256)
func (_LargeStaking *LargeStakingCallerSession) ValidatorSlashAmount(arg0 []byte) (*big.Int, error) {
	return _LargeStaking.Contract.ValidatorSlashAmount(&_LargeStaking.CallOpts, arg0)
}

// ValuePerShare is a free data retrieval call binding the contract method 0x182324e4.
//
// Solidity: function valuePerShare(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingCaller) ValuePerShare(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LargeStaking.contract.Call(opts, &out, "valuePerShare", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValuePerShare is a free data retrieval call binding the contract method 0x182324e4.
//
// Solidity: function valuePerShare(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingSession) ValuePerShare(arg0 *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.ValuePerShare(&_LargeStaking.CallOpts, arg0)
}

// ValuePerShare is a free data retrieval call binding the contract method 0x182324e4.
//
// Solidity: function valuePerShare(uint256 ) view returns(uint256)
func (_LargeStaking *LargeStakingCallerSession) ValuePerShare(arg0 *big.Int) (*big.Int, error) {
	return _LargeStaking.Contract.ValuePerShare(&_LargeStaking.CallOpts, arg0)
}

// AppendLargeStake is a paid mutator transaction binding the contract method 0x1df9eaa6.
//
// Solidity: function appendLargeStake(uint256 _stakingId, address _owner, address _withdrawCredentials) payable returns()
func (_LargeStaking *LargeStakingTransactor) AppendLargeStake(opts *bind.TransactOpts, _stakingId *big.Int, _owner common.Address, _withdrawCredentials common.Address) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "appendLargeStake", _stakingId, _owner, _withdrawCredentials)
}

// AppendLargeStake is a paid mutator transaction binding the contract method 0x1df9eaa6.
//
// Solidity: function appendLargeStake(uint256 _stakingId, address _owner, address _withdrawCredentials) payable returns()
func (_LargeStaking *LargeStakingSession) AppendLargeStake(_stakingId *big.Int, _owner common.Address, _withdrawCredentials common.Address) (*types.Transaction, error) {
	return _LargeStaking.Contract.AppendLargeStake(&_LargeStaking.TransactOpts, _stakingId, _owner, _withdrawCredentials)
}

// AppendLargeStake is a paid mutator transaction binding the contract method 0x1df9eaa6.
//
// Solidity: function appendLargeStake(uint256 _stakingId, address _owner, address _withdrawCredentials) payable returns()
func (_LargeStaking *LargeStakingTransactorSession) AppendLargeStake(_stakingId *big.Int, _owner common.Address, _withdrawCredentials common.Address) (*types.Transaction, error) {
	return _LargeStaking.Contract.AppendLargeStake(&_LargeStaking.TransactOpts, _stakingId, _owner, _withdrawCredentials)
}

// AppendMigrateStake is a paid mutator transaction binding the contract method 0x177b37ce.
//
// Solidity: function appendMigrateStake(uint256 _stakingId, address _owner, address _withdrawCredentials, bytes[] _pubKeys) returns()
func (_LargeStaking *LargeStakingTransactor) AppendMigrateStake(opts *bind.TransactOpts, _stakingId *big.Int, _owner common.Address, _withdrawCredentials common.Address, _pubKeys [][]byte) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "appendMigrateStake", _stakingId, _owner, _withdrawCredentials, _pubKeys)
}

// AppendMigrateStake is a paid mutator transaction binding the contract method 0x177b37ce.
//
// Solidity: function appendMigrateStake(uint256 _stakingId, address _owner, address _withdrawCredentials, bytes[] _pubKeys) returns()
func (_LargeStaking *LargeStakingSession) AppendMigrateStake(_stakingId *big.Int, _owner common.Address, _withdrawCredentials common.Address, _pubKeys [][]byte) (*types.Transaction, error) {
	return _LargeStaking.Contract.AppendMigrateStake(&_LargeStaking.TransactOpts, _stakingId, _owner, _withdrawCredentials, _pubKeys)
}

// AppendMigrateStake is a paid mutator transaction binding the contract method 0x177b37ce.
//
// Solidity: function appendMigrateStake(uint256 _stakingId, address _owner, address _withdrawCredentials, bytes[] _pubKeys) returns()
func (_LargeStaking *LargeStakingTransactorSession) AppendMigrateStake(_stakingId *big.Int, _owner common.Address, _withdrawCredentials common.Address, _pubKeys [][]byte) (*types.Transaction, error) {
	return _LargeStaking.Contract.AppendMigrateStake(&_LargeStaking.TransactOpts, _stakingId, _owner, _withdrawCredentials, _pubKeys)
}

// ClaimRewardsOfDao is a paid mutator transaction binding the contract method 0x694f0b11.
//
// Solidity: function claimRewardsOfDao(uint256[] _stakingIds) returns()
func (_LargeStaking *LargeStakingTransactor) ClaimRewardsOfDao(opts *bind.TransactOpts, _stakingIds []*big.Int) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "claimRewardsOfDao", _stakingIds)
}

// ClaimRewardsOfDao is a paid mutator transaction binding the contract method 0x694f0b11.
//
// Solidity: function claimRewardsOfDao(uint256[] _stakingIds) returns()
func (_LargeStaking *LargeStakingSession) ClaimRewardsOfDao(_stakingIds []*big.Int) (*types.Transaction, error) {
	return _LargeStaking.Contract.ClaimRewardsOfDao(&_LargeStaking.TransactOpts, _stakingIds)
}

// ClaimRewardsOfDao is a paid mutator transaction binding the contract method 0x694f0b11.
//
// Solidity: function claimRewardsOfDao(uint256[] _stakingIds) returns()
func (_LargeStaking *LargeStakingTransactorSession) ClaimRewardsOfDao(_stakingIds []*big.Int) (*types.Transaction, error) {
	return _LargeStaking.Contract.ClaimRewardsOfDao(&_LargeStaking.TransactOpts, _stakingIds)
}

// ClaimRewardsOfOperator is a paid mutator transaction binding the contract method 0xbb7e7a97.
//
// Solidity: function claimRewardsOfOperator(uint256 _operatorId, uint256[] _privatePoolStakingIds) returns()
func (_LargeStaking *LargeStakingTransactor) ClaimRewardsOfOperator(opts *bind.TransactOpts, _operatorId *big.Int, _privatePoolStakingIds []*big.Int) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "claimRewardsOfOperator", _operatorId, _privatePoolStakingIds)
}

// ClaimRewardsOfOperator is a paid mutator transaction binding the contract method 0xbb7e7a97.
//
// Solidity: function claimRewardsOfOperator(uint256 _operatorId, uint256[] _privatePoolStakingIds) returns()
func (_LargeStaking *LargeStakingSession) ClaimRewardsOfOperator(_operatorId *big.Int, _privatePoolStakingIds []*big.Int) (*types.Transaction, error) {
	return _LargeStaking.Contract.ClaimRewardsOfOperator(&_LargeStaking.TransactOpts, _operatorId, _privatePoolStakingIds)
}

// ClaimRewardsOfOperator is a paid mutator transaction binding the contract method 0xbb7e7a97.
//
// Solidity: function claimRewardsOfOperator(uint256 _operatorId, uint256[] _privatePoolStakingIds) returns()
func (_LargeStaking *LargeStakingTransactorSession) ClaimRewardsOfOperator(_operatorId *big.Int, _privatePoolStakingIds []*big.Int) (*types.Transaction, error) {
	return _LargeStaking.Contract.ClaimRewardsOfOperator(&_LargeStaking.TransactOpts, _operatorId, _privatePoolStakingIds)
}

// ClaimRewardsOfUser is a paid mutator transaction binding the contract method 0x20e369ba.
//
// Solidity: function claimRewardsOfUser(uint256 _stakingId, uint256 rewards) returns()
func (_LargeStaking *LargeStakingTransactor) ClaimRewardsOfUser(opts *bind.TransactOpts, _stakingId *big.Int, rewards *big.Int) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "claimRewardsOfUser", _stakingId, rewards)
}

// ClaimRewardsOfUser is a paid mutator transaction binding the contract method 0x20e369ba.
//
// Solidity: function claimRewardsOfUser(uint256 _stakingId, uint256 rewards) returns()
func (_LargeStaking *LargeStakingSession) ClaimRewardsOfUser(_stakingId *big.Int, rewards *big.Int) (*types.Transaction, error) {
	return _LargeStaking.Contract.ClaimRewardsOfUser(&_LargeStaking.TransactOpts, _stakingId, rewards)
}

// ClaimRewardsOfUser is a paid mutator transaction binding the contract method 0x20e369ba.
//
// Solidity: function claimRewardsOfUser(uint256 _stakingId, uint256 rewards) returns()
func (_LargeStaking *LargeStakingTransactorSession) ClaimRewardsOfUser(_stakingId *big.Int, rewards *big.Int) (*types.Transaction, error) {
	return _LargeStaking.Contract.ClaimRewardsOfUser(&_LargeStaking.TransactOpts, _stakingId, rewards)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _dao, address _daoVaultAddress, address _nodeOperatorRegistryAddress, address _operatorSlashContract, address _consensusOracleContractAddr, address _elRewardFactory, address _depositContract) returns()
func (_LargeStaking *LargeStakingTransactor) Initialize(opts *bind.TransactOpts, _dao common.Address, _daoVaultAddress common.Address, _nodeOperatorRegistryAddress common.Address, _operatorSlashContract common.Address, _consensusOracleContractAddr common.Address, _elRewardFactory common.Address, _depositContract common.Address) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "initialize", _dao, _daoVaultAddress, _nodeOperatorRegistryAddress, _operatorSlashContract, _consensusOracleContractAddr, _elRewardFactory, _depositContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _dao, address _daoVaultAddress, address _nodeOperatorRegistryAddress, address _operatorSlashContract, address _consensusOracleContractAddr, address _elRewardFactory, address _depositContract) returns()
func (_LargeStaking *LargeStakingSession) Initialize(_dao common.Address, _daoVaultAddress common.Address, _nodeOperatorRegistryAddress common.Address, _operatorSlashContract common.Address, _consensusOracleContractAddr common.Address, _elRewardFactory common.Address, _depositContract common.Address) (*types.Transaction, error) {
	return _LargeStaking.Contract.Initialize(&_LargeStaking.TransactOpts, _dao, _daoVaultAddress, _nodeOperatorRegistryAddress, _operatorSlashContract, _consensusOracleContractAddr, _elRewardFactory, _depositContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _dao, address _daoVaultAddress, address _nodeOperatorRegistryAddress, address _operatorSlashContract, address _consensusOracleContractAddr, address _elRewardFactory, address _depositContract) returns()
func (_LargeStaking *LargeStakingTransactorSession) Initialize(_dao common.Address, _daoVaultAddress common.Address, _nodeOperatorRegistryAddress common.Address, _operatorSlashContract common.Address, _consensusOracleContractAddr common.Address, _elRewardFactory common.Address, _depositContract common.Address) (*types.Transaction, error) {
	return _LargeStaking.Contract.Initialize(&_LargeStaking.TransactOpts, _dao, _daoVaultAddress, _nodeOperatorRegistryAddress, _operatorSlashContract, _consensusOracleContractAddr, _elRewardFactory, _depositContract)
}

// LargeStake is a paid mutator transaction binding the contract method 0x1bbe36de.
//
// Solidity: function largeStake(uint256 _operatorId, address _elRewardAddr, address _withdrawCredentials, bool _isELRewardSharing) payable returns()
func (_LargeStaking *LargeStakingTransactor) LargeStake(opts *bind.TransactOpts, _operatorId *big.Int, _elRewardAddr common.Address, _withdrawCredentials common.Address, _isELRewardSharing bool) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "largeStake", _operatorId, _elRewardAddr, _withdrawCredentials, _isELRewardSharing)
}

// LargeStake is a paid mutator transaction binding the contract method 0x1bbe36de.
//
// Solidity: function largeStake(uint256 _operatorId, address _elRewardAddr, address _withdrawCredentials, bool _isELRewardSharing) payable returns()
func (_LargeStaking *LargeStakingSession) LargeStake(_operatorId *big.Int, _elRewardAddr common.Address, _withdrawCredentials common.Address, _isELRewardSharing bool) (*types.Transaction, error) {
	return _LargeStaking.Contract.LargeStake(&_LargeStaking.TransactOpts, _operatorId, _elRewardAddr, _withdrawCredentials, _isELRewardSharing)
}

// LargeStake is a paid mutator transaction binding the contract method 0x1bbe36de.
//
// Solidity: function largeStake(uint256 _operatorId, address _elRewardAddr, address _withdrawCredentials, bool _isELRewardSharing) payable returns()
func (_LargeStaking *LargeStakingTransactorSession) LargeStake(_operatorId *big.Int, _elRewardAddr common.Address, _withdrawCredentials common.Address, _isELRewardSharing bool) (*types.Transaction, error) {
	return _LargeStaking.Contract.LargeStake(&_LargeStaking.TransactOpts, _operatorId, _elRewardAddr, _withdrawCredentials, _isELRewardSharing)
}

// LargeUnstake is a paid mutator transaction binding the contract method 0x1642a5b8.
//
// Solidity: function largeUnstake(uint256 _stakingId, uint256 _amount) returns()
func (_LargeStaking *LargeStakingTransactor) LargeUnstake(opts *bind.TransactOpts, _stakingId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "largeUnstake", _stakingId, _amount)
}

// LargeUnstake is a paid mutator transaction binding the contract method 0x1642a5b8.
//
// Solidity: function largeUnstake(uint256 _stakingId, uint256 _amount) returns()
func (_LargeStaking *LargeStakingSession) LargeUnstake(_stakingId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _LargeStaking.Contract.LargeUnstake(&_LargeStaking.TransactOpts, _stakingId, _amount)
}

// LargeUnstake is a paid mutator transaction binding the contract method 0x1642a5b8.
//
// Solidity: function largeUnstake(uint256 _stakingId, uint256 _amount) returns()
func (_LargeStaking *LargeStakingTransactorSession) LargeUnstake(_stakingId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _LargeStaking.Contract.LargeUnstake(&_LargeStaking.TransactOpts, _stakingId, _amount)
}

// MigrateStake is a paid mutator transaction binding the contract method 0xc49e85a0.
//
// Solidity: function migrateStake(address _owner, address _elRewardAddr, address _withdrawCredentials, bool _isELRewardSharing, bytes[] _pubKeys) returns()
func (_LargeStaking *LargeStakingTransactor) MigrateStake(opts *bind.TransactOpts, _owner common.Address, _elRewardAddr common.Address, _withdrawCredentials common.Address, _isELRewardSharing bool, _pubKeys [][]byte) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "migrateStake", _owner, _elRewardAddr, _withdrawCredentials, _isELRewardSharing, _pubKeys)
}

// MigrateStake is a paid mutator transaction binding the contract method 0xc49e85a0.
//
// Solidity: function migrateStake(address _owner, address _elRewardAddr, address _withdrawCredentials, bool _isELRewardSharing, bytes[] _pubKeys) returns()
func (_LargeStaking *LargeStakingSession) MigrateStake(_owner common.Address, _elRewardAddr common.Address, _withdrawCredentials common.Address, _isELRewardSharing bool, _pubKeys [][]byte) (*types.Transaction, error) {
	return _LargeStaking.Contract.MigrateStake(&_LargeStaking.TransactOpts, _owner, _elRewardAddr, _withdrawCredentials, _isELRewardSharing, _pubKeys)
}

// MigrateStake is a paid mutator transaction binding the contract method 0xc49e85a0.
//
// Solidity: function migrateStake(address _owner, address _elRewardAddr, address _withdrawCredentials, bool _isELRewardSharing, bytes[] _pubKeys) returns()
func (_LargeStaking *LargeStakingTransactorSession) MigrateStake(_owner common.Address, _elRewardAddr common.Address, _withdrawCredentials common.Address, _isELRewardSharing bool, _pubKeys [][]byte) (*types.Transaction, error) {
	return _LargeStaking.Contract.MigrateStake(&_LargeStaking.TransactOpts, _owner, _elRewardAddr, _withdrawCredentials, _isELRewardSharing, _pubKeys)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x6d0ca605.
//
// Solidity: function registerValidator(uint256 _stakingId, bytes[] _pubkeys, bytes[] _signatures, bytes32[] _depositDataRoots) returns()
func (_LargeStaking *LargeStakingTransactor) RegisterValidator(opts *bind.TransactOpts, _stakingId *big.Int, _pubkeys [][]byte, _signatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "registerValidator", _stakingId, _pubkeys, _signatures, _depositDataRoots)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x6d0ca605.
//
// Solidity: function registerValidator(uint256 _stakingId, bytes[] _pubkeys, bytes[] _signatures, bytes32[] _depositDataRoots) returns()
func (_LargeStaking *LargeStakingSession) RegisterValidator(_stakingId *big.Int, _pubkeys [][]byte, _signatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _LargeStaking.Contract.RegisterValidator(&_LargeStaking.TransactOpts, _stakingId, _pubkeys, _signatures, _depositDataRoots)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x6d0ca605.
//
// Solidity: function registerValidator(uint256 _stakingId, bytes[] _pubkeys, bytes[] _signatures, bytes32[] _depositDataRoots) returns()
func (_LargeStaking *LargeStakingTransactorSession) RegisterValidator(_stakingId *big.Int, _pubkeys [][]byte, _signatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _LargeStaking.Contract.RegisterValidator(&_LargeStaking.TransactOpts, _stakingId, _pubkeys, _signatures, _depositDataRoots)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LargeStaking *LargeStakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LargeStaking *LargeStakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _LargeStaking.Contract.RenounceOwnership(&_LargeStaking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LargeStaking *LargeStakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LargeStaking.Contract.RenounceOwnership(&_LargeStaking.TransactOpts)
}

// ReportCLStakingData is a paid mutator transaction binding the contract method 0xa036ec84.
//
// Solidity: function reportCLStakingData((uint128,bytes)[] _clStakingExitInfo, (uint128,uint128,bytes)[] _clStakingSlashInfo) returns()
func (_LargeStaking *LargeStakingTransactor) ReportCLStakingData(opts *bind.TransactOpts, _clStakingExitInfo []CLStakingExitInfo, _clStakingSlashInfo []CLStakingSlashInfo) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "reportCLStakingData", _clStakingExitInfo, _clStakingSlashInfo)
}

// ReportCLStakingData is a paid mutator transaction binding the contract method 0xa036ec84.
//
// Solidity: function reportCLStakingData((uint128,bytes)[] _clStakingExitInfo, (uint128,uint128,bytes)[] _clStakingSlashInfo) returns()
func (_LargeStaking *LargeStakingSession) ReportCLStakingData(_clStakingExitInfo []CLStakingExitInfo, _clStakingSlashInfo []CLStakingSlashInfo) (*types.Transaction, error) {
	return _LargeStaking.Contract.ReportCLStakingData(&_LargeStaking.TransactOpts, _clStakingExitInfo, _clStakingSlashInfo)
}

// ReportCLStakingData is a paid mutator transaction binding the contract method 0xa036ec84.
//
// Solidity: function reportCLStakingData((uint128,bytes)[] _clStakingExitInfo, (uint128,uint128,bytes)[] _clStakingSlashInfo) returns()
func (_LargeStaking *LargeStakingTransactorSession) ReportCLStakingData(_clStakingExitInfo []CLStakingExitInfo, _clStakingSlashInfo []CLStakingSlashInfo) (*types.Transaction, error) {
	return _LargeStaking.Contract.ReportCLStakingData(&_LargeStaking.TransactOpts, _clStakingExitInfo, _clStakingSlashInfo)
}

// SetLargeStakingSetting is a paid mutator transaction binding the contract method 0x3454ea16.
//
// Solidity: function setLargeStakingSetting(address _dao, address _daoVaultAddress, uint256 _daoElCommissionRate, uint256 _MIN_STAKE_AMOUNT, address _nodeOperatorRegistryAddress, address _consensusOracleContractAddr, address _elRewardFactory, address _operatorSlashContract) returns()
func (_LargeStaking *LargeStakingTransactor) SetLargeStakingSetting(opts *bind.TransactOpts, _dao common.Address, _daoVaultAddress common.Address, _daoElCommissionRate *big.Int, _MIN_STAKE_AMOUNT *big.Int, _nodeOperatorRegistryAddress common.Address, _consensusOracleContractAddr common.Address, _elRewardFactory common.Address, _operatorSlashContract common.Address) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "setLargeStakingSetting", _dao, _daoVaultAddress, _daoElCommissionRate, _MIN_STAKE_AMOUNT, _nodeOperatorRegistryAddress, _consensusOracleContractAddr, _elRewardFactory, _operatorSlashContract)
}

// SetLargeStakingSetting is a paid mutator transaction binding the contract method 0x3454ea16.
//
// Solidity: function setLargeStakingSetting(address _dao, address _daoVaultAddress, uint256 _daoElCommissionRate, uint256 _MIN_STAKE_AMOUNT, address _nodeOperatorRegistryAddress, address _consensusOracleContractAddr, address _elRewardFactory, address _operatorSlashContract) returns()
func (_LargeStaking *LargeStakingSession) SetLargeStakingSetting(_dao common.Address, _daoVaultAddress common.Address, _daoElCommissionRate *big.Int, _MIN_STAKE_AMOUNT *big.Int, _nodeOperatorRegistryAddress common.Address, _consensusOracleContractAddr common.Address, _elRewardFactory common.Address, _operatorSlashContract common.Address) (*types.Transaction, error) {
	return _LargeStaking.Contract.SetLargeStakingSetting(&_LargeStaking.TransactOpts, _dao, _daoVaultAddress, _daoElCommissionRate, _MIN_STAKE_AMOUNT, _nodeOperatorRegistryAddress, _consensusOracleContractAddr, _elRewardFactory, _operatorSlashContract)
}

// SetLargeStakingSetting is a paid mutator transaction binding the contract method 0x3454ea16.
//
// Solidity: function setLargeStakingSetting(address _dao, address _daoVaultAddress, uint256 _daoElCommissionRate, uint256 _MIN_STAKE_AMOUNT, address _nodeOperatorRegistryAddress, address _consensusOracleContractAddr, address _elRewardFactory, address _operatorSlashContract) returns()
func (_LargeStaking *LargeStakingTransactorSession) SetLargeStakingSetting(_dao common.Address, _daoVaultAddress common.Address, _daoElCommissionRate *big.Int, _MIN_STAKE_AMOUNT *big.Int, _nodeOperatorRegistryAddress common.Address, _consensusOracleContractAddr common.Address, _elRewardFactory common.Address, _operatorSlashContract common.Address) (*types.Transaction, error) {
	return _LargeStaking.Contract.SetLargeStakingSetting(&_LargeStaking.TransactOpts, _dao, _daoVaultAddress, _daoElCommissionRate, _MIN_STAKE_AMOUNT, _nodeOperatorRegistryAddress, _consensusOracleContractAddr, _elRewardFactory, _operatorSlashContract)
}

// SettleElPrivateReward is a paid mutator transaction binding the contract method 0x1c40b6cf.
//
// Solidity: function settleElPrivateReward(uint256 _stakingId) returns()
func (_LargeStaking *LargeStakingTransactor) SettleElPrivateReward(opts *bind.TransactOpts, _stakingId *big.Int) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "settleElPrivateReward", _stakingId)
}

// SettleElPrivateReward is a paid mutator transaction binding the contract method 0x1c40b6cf.
//
// Solidity: function settleElPrivateReward(uint256 _stakingId) returns()
func (_LargeStaking *LargeStakingSession) SettleElPrivateReward(_stakingId *big.Int) (*types.Transaction, error) {
	return _LargeStaking.Contract.SettleElPrivateReward(&_LargeStaking.TransactOpts, _stakingId)
}

// SettleElPrivateReward is a paid mutator transaction binding the contract method 0x1c40b6cf.
//
// Solidity: function settleElPrivateReward(uint256 _stakingId) returns()
func (_LargeStaking *LargeStakingTransactorSession) SettleElPrivateReward(_stakingId *big.Int) (*types.Transaction, error) {
	return _LargeStaking.Contract.SettleElPrivateReward(&_LargeStaking.TransactOpts, _stakingId)
}

// SettleElSharedReward is a paid mutator transaction binding the contract method 0x65479b8b.
//
// Solidity: function settleElSharedReward(uint256 _operatorId) returns()
func (_LargeStaking *LargeStakingTransactor) SettleElSharedReward(opts *bind.TransactOpts, _operatorId *big.Int) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "settleElSharedReward", _operatorId)
}

// SettleElSharedReward is a paid mutator transaction binding the contract method 0x65479b8b.
//
// Solidity: function settleElSharedReward(uint256 _operatorId) returns()
func (_LargeStaking *LargeStakingSession) SettleElSharedReward(_operatorId *big.Int) (*types.Transaction, error) {
	return _LargeStaking.Contract.SettleElSharedReward(&_LargeStaking.TransactOpts, _operatorId)
}

// SettleElSharedReward is a paid mutator transaction binding the contract method 0x65479b8b.
//
// Solidity: function settleElSharedReward(uint256 _operatorId) returns()
func (_LargeStaking *LargeStakingTransactorSession) SettleElSharedReward(_operatorId *big.Int) (*types.Transaction, error) {
	return _LargeStaking.Contract.SettleElSharedReward(&_LargeStaking.TransactOpts, _operatorId)
}

// StartupSharedRewardPool is a paid mutator transaction binding the contract method 0x6fbe33d5.
//
// Solidity: function startupSharedRewardPool(uint256 _operatorId) returns()
func (_LargeStaking *LargeStakingTransactor) StartupSharedRewardPool(opts *bind.TransactOpts, _operatorId *big.Int) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "startupSharedRewardPool", _operatorId)
}

// StartupSharedRewardPool is a paid mutator transaction binding the contract method 0x6fbe33d5.
//
// Solidity: function startupSharedRewardPool(uint256 _operatorId) returns()
func (_LargeStaking *LargeStakingSession) StartupSharedRewardPool(_operatorId *big.Int) (*types.Transaction, error) {
	return _LargeStaking.Contract.StartupSharedRewardPool(&_LargeStaking.TransactOpts, _operatorId)
}

// StartupSharedRewardPool is a paid mutator transaction binding the contract method 0x6fbe33d5.
//
// Solidity: function startupSharedRewardPool(uint256 _operatorId) returns()
func (_LargeStaking *LargeStakingTransactorSession) StartupSharedRewardPool(_operatorId *big.Int) (*types.Transaction, error) {
	return _LargeStaking.Contract.StartupSharedRewardPool(&_LargeStaking.TransactOpts, _operatorId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LargeStaking *LargeStakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LargeStaking *LargeStakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LargeStaking.Contract.TransferOwnership(&_LargeStaking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LargeStaking *LargeStakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LargeStaking.Contract.TransferOwnership(&_LargeStaking.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LargeStaking *LargeStakingTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LargeStaking *LargeStakingSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LargeStaking.Contract.UpgradeTo(&_LargeStaking.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LargeStaking *LargeStakingTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LargeStaking.Contract.UpgradeTo(&_LargeStaking.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LargeStaking *LargeStakingTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LargeStaking.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LargeStaking *LargeStakingSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LargeStaking.Contract.UpgradeToAndCall(&_LargeStaking.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LargeStaking *LargeStakingTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LargeStaking.Contract.UpgradeToAndCall(&_LargeStaking.TransactOpts, newImplementation, data)
}

// LargeStakingAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the LargeStaking contract.
type LargeStakingAdminChangedIterator struct {
	Event *LargeStakingAdminChanged // Event containing the contract specifics and raw log

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
func (it *LargeStakingAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingAdminChanged)
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
		it.Event = new(LargeStakingAdminChanged)
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
func (it *LargeStakingAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingAdminChanged represents a AdminChanged event raised by the LargeStaking contract.
type LargeStakingAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LargeStaking *LargeStakingFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*LargeStakingAdminChangedIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &LargeStakingAdminChangedIterator{contract: _LargeStaking.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LargeStaking *LargeStakingFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *LargeStakingAdminChanged) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingAdminChanged)
				if err := _LargeStaking.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_LargeStaking *LargeStakingFilterer) ParseAdminChanged(log types.Log) (*LargeStakingAdminChanged, error) {
	event := new(LargeStakingAdminChanged)
	if err := _LargeStaking.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingAppendMigretaStakeIterator is returned from FilterAppendMigretaStake and is used to iterate over the raw logs and unpacked data for AppendMigretaStake events raised by the LargeStaking contract.
type LargeStakingAppendMigretaStakeIterator struct {
	Event *LargeStakingAppendMigretaStake // Event containing the contract specifics and raw log

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
func (it *LargeStakingAppendMigretaStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingAppendMigretaStake)
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
		it.Event = new(LargeStakingAppendMigretaStake)
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
func (it *LargeStakingAppendMigretaStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingAppendMigretaStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingAppendMigretaStake represents a AppendMigretaStake event raised by the LargeStaking contract.
type LargeStakingAppendMigretaStake struct {
	StakingId    *big.Int
	StakeAmounts *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAppendMigretaStake is a free log retrieval operation binding the contract event 0x9b894838c9537381dddefdb88d510d3929cec600826695b625fa3f874e1c96ed.
//
// Solidity: event AppendMigretaStake(uint256 _stakingId, uint256 _stakeAmounts)
func (_LargeStaking *LargeStakingFilterer) FilterAppendMigretaStake(opts *bind.FilterOpts) (*LargeStakingAppendMigretaStakeIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "AppendMigretaStake")
	if err != nil {
		return nil, err
	}
	return &LargeStakingAppendMigretaStakeIterator{contract: _LargeStaking.contract, event: "AppendMigretaStake", logs: logs, sub: sub}, nil
}

// WatchAppendMigretaStake is a free log subscription operation binding the contract event 0x9b894838c9537381dddefdb88d510d3929cec600826695b625fa3f874e1c96ed.
//
// Solidity: event AppendMigretaStake(uint256 _stakingId, uint256 _stakeAmounts)
func (_LargeStaking *LargeStakingFilterer) WatchAppendMigretaStake(opts *bind.WatchOpts, sink chan<- *LargeStakingAppendMigretaStake) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "AppendMigretaStake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingAppendMigretaStake)
				if err := _LargeStaking.contract.UnpackLog(event, "AppendMigretaStake", log); err != nil {
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

// ParseAppendMigretaStake is a log parse operation binding the contract event 0x9b894838c9537381dddefdb88d510d3929cec600826695b625fa3f874e1c96ed.
//
// Solidity: event AppendMigretaStake(uint256 _stakingId, uint256 _stakeAmounts)
func (_LargeStaking *LargeStakingFilterer) ParseAppendMigretaStake(log types.Log) (*LargeStakingAppendMigretaStake, error) {
	event := new(LargeStakingAppendMigretaStake)
	if err := _LargeStaking.contract.UnpackLog(event, "AppendMigretaStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingAppendStakeIterator is returned from FilterAppendStake and is used to iterate over the raw logs and unpacked data for AppendStake events raised by the LargeStaking contract.
type LargeStakingAppendStakeIterator struct {
	Event *LargeStakingAppendStake // Event containing the contract specifics and raw log

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
func (it *LargeStakingAppendStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingAppendStake)
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
		it.Event = new(LargeStakingAppendStake)
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
func (it *LargeStakingAppendStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingAppendStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingAppendStake represents a AppendStake event raised by the LargeStaking contract.
type LargeStakingAppendStake struct {
	StakingId *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAppendStake is a free log retrieval operation binding the contract event 0x6f5e1b7f4680eb279bd51c12508b71101faa95ecc944ca6536f6da27877b4aed.
//
// Solidity: event AppendStake(uint256 _stakingId, uint256 _amount)
func (_LargeStaking *LargeStakingFilterer) FilterAppendStake(opts *bind.FilterOpts) (*LargeStakingAppendStakeIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "AppendStake")
	if err != nil {
		return nil, err
	}
	return &LargeStakingAppendStakeIterator{contract: _LargeStaking.contract, event: "AppendStake", logs: logs, sub: sub}, nil
}

// WatchAppendStake is a free log subscription operation binding the contract event 0x6f5e1b7f4680eb279bd51c12508b71101faa95ecc944ca6536f6da27877b4aed.
//
// Solidity: event AppendStake(uint256 _stakingId, uint256 _amount)
func (_LargeStaking *LargeStakingFilterer) WatchAppendStake(opts *bind.WatchOpts, sink chan<- *LargeStakingAppendStake) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "AppendStake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingAppendStake)
				if err := _LargeStaking.contract.UnpackLog(event, "AppendStake", log); err != nil {
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

// ParseAppendStake is a log parse operation binding the contract event 0x6f5e1b7f4680eb279bd51c12508b71101faa95ecc944ca6536f6da27877b4aed.
//
// Solidity: event AppendStake(uint256 _stakingId, uint256 _amount)
func (_LargeStaking *LargeStakingFilterer) ParseAppendStake(log types.Log) (*LargeStakingAppendStake, error) {
	event := new(LargeStakingAppendStake)
	if err := _LargeStaking.contract.UnpackLog(event, "AppendStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the LargeStaking contract.
type LargeStakingBeaconUpgradedIterator struct {
	Event *LargeStakingBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *LargeStakingBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingBeaconUpgraded)
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
		it.Event = new(LargeStakingBeaconUpgraded)
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
func (it *LargeStakingBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingBeaconUpgraded represents a BeaconUpgraded event raised by the LargeStaking contract.
type LargeStakingBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LargeStaking *LargeStakingFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*LargeStakingBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &LargeStakingBeaconUpgradedIterator{contract: _LargeStaking.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LargeStaking *LargeStakingFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *LargeStakingBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingBeaconUpgraded)
				if err := _LargeStaking.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_LargeStaking *LargeStakingFilterer) ParseBeaconUpgraded(log types.Log) (*LargeStakingBeaconUpgraded, error) {
	event := new(LargeStakingBeaconUpgraded)
	if err := _LargeStaking.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingConsensusOracleChangedIterator is returned from FilterConsensusOracleChanged and is used to iterate over the raw logs and unpacked data for ConsensusOracleChanged events raised by the LargeStaking contract.
type LargeStakingConsensusOracleChangedIterator struct {
	Event *LargeStakingConsensusOracleChanged // Event containing the contract specifics and raw log

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
func (it *LargeStakingConsensusOracleChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingConsensusOracleChanged)
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
		it.Event = new(LargeStakingConsensusOracleChanged)
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
func (it *LargeStakingConsensusOracleChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingConsensusOracleChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingConsensusOracleChanged represents a ConsensusOracleChanged event raised by the LargeStaking contract.
type LargeStakingConsensusOracleChanged struct {
	OldConsensusOracleContractAddr common.Address
	ConsensusOracleContractAddr    common.Address
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterConsensusOracleChanged is a free log retrieval operation binding the contract event 0xb92a521dda1a9cd90eb9401c62c27b3cb0f2f441e4a0c6b5c72e3eebb3d0f2d6.
//
// Solidity: event ConsensusOracleChanged(address _oldConsensusOracleContractAddr, address _consensusOracleContractAddr)
func (_LargeStaking *LargeStakingFilterer) FilterConsensusOracleChanged(opts *bind.FilterOpts) (*LargeStakingConsensusOracleChangedIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "ConsensusOracleChanged")
	if err != nil {
		return nil, err
	}
	return &LargeStakingConsensusOracleChangedIterator{contract: _LargeStaking.contract, event: "ConsensusOracleChanged", logs: logs, sub: sub}, nil
}

// WatchConsensusOracleChanged is a free log subscription operation binding the contract event 0xb92a521dda1a9cd90eb9401c62c27b3cb0f2f441e4a0c6b5c72e3eebb3d0f2d6.
//
// Solidity: event ConsensusOracleChanged(address _oldConsensusOracleContractAddr, address _consensusOracleContractAddr)
func (_LargeStaking *LargeStakingFilterer) WatchConsensusOracleChanged(opts *bind.WatchOpts, sink chan<- *LargeStakingConsensusOracleChanged) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "ConsensusOracleChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingConsensusOracleChanged)
				if err := _LargeStaking.contract.UnpackLog(event, "ConsensusOracleChanged", log); err != nil {
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

// ParseConsensusOracleChanged is a log parse operation binding the contract event 0xb92a521dda1a9cd90eb9401c62c27b3cb0f2f441e4a0c6b5c72e3eebb3d0f2d6.
//
// Solidity: event ConsensusOracleChanged(address _oldConsensusOracleContractAddr, address _consensusOracleContractAddr)
func (_LargeStaking *LargeStakingFilterer) ParseConsensusOracleChanged(log types.Log) (*LargeStakingConsensusOracleChanged, error) {
	event := new(LargeStakingConsensusOracleChanged)
	if err := _LargeStaking.contract.UnpackLog(event, "ConsensusOracleChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingDaoAddressChangedIterator is returned from FilterDaoAddressChanged and is used to iterate over the raw logs and unpacked data for DaoAddressChanged events raised by the LargeStaking contract.
type LargeStakingDaoAddressChangedIterator struct {
	Event *LargeStakingDaoAddressChanged // Event containing the contract specifics and raw log

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
func (it *LargeStakingDaoAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingDaoAddressChanged)
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
		it.Event = new(LargeStakingDaoAddressChanged)
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
func (it *LargeStakingDaoAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingDaoAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingDaoAddressChanged represents a DaoAddressChanged event raised by the LargeStaking contract.
type LargeStakingDaoAddressChanged struct {
	OldDao common.Address
	Dao    common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDaoAddressChanged is a free log retrieval operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address _oldDao, address _dao)
func (_LargeStaking *LargeStakingFilterer) FilterDaoAddressChanged(opts *bind.FilterOpts) (*LargeStakingDaoAddressChangedIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "DaoAddressChanged")
	if err != nil {
		return nil, err
	}
	return &LargeStakingDaoAddressChangedIterator{contract: _LargeStaking.contract, event: "DaoAddressChanged", logs: logs, sub: sub}, nil
}

// WatchDaoAddressChanged is a free log subscription operation binding the contract event 0xd5b3b0e6e0098a82fa04cf04faff9109f98389a10c80f20eb7186b9274168946.
//
// Solidity: event DaoAddressChanged(address _oldDao, address _dao)
func (_LargeStaking *LargeStakingFilterer) WatchDaoAddressChanged(opts *bind.WatchOpts, sink chan<- *LargeStakingDaoAddressChanged) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "DaoAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingDaoAddressChanged)
				if err := _LargeStaking.contract.UnpackLog(event, "DaoAddressChanged", log); err != nil {
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
func (_LargeStaking *LargeStakingFilterer) ParseDaoAddressChanged(log types.Log) (*LargeStakingDaoAddressChanged, error) {
	event := new(LargeStakingDaoAddressChanged)
	if err := _LargeStaking.contract.UnpackLog(event, "DaoAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingDaoELCommissionRateChangedIterator is returned from FilterDaoELCommissionRateChanged and is used to iterate over the raw logs and unpacked data for DaoELCommissionRateChanged events raised by the LargeStaking contract.
type LargeStakingDaoELCommissionRateChangedIterator struct {
	Event *LargeStakingDaoELCommissionRateChanged // Event containing the contract specifics and raw log

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
func (it *LargeStakingDaoELCommissionRateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingDaoELCommissionRateChanged)
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
		it.Event = new(LargeStakingDaoELCommissionRateChanged)
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
func (it *LargeStakingDaoELCommissionRateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingDaoELCommissionRateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingDaoELCommissionRateChanged represents a DaoELCommissionRateChanged event raised by the LargeStaking contract.
type LargeStakingDaoELCommissionRateChanged struct {
	OldDaoElCommissionRate *big.Int
	DaoElCommissionRate    *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterDaoELCommissionRateChanged is a free log retrieval operation binding the contract event 0xf34be367dcf9d61f1d01f06f296eb532d42f925474645b42aa776dd4c83900fc.
//
// Solidity: event DaoELCommissionRateChanged(uint256 _oldDaoElCommissionRate, uint256 _daoElCommissionRate)
func (_LargeStaking *LargeStakingFilterer) FilterDaoELCommissionRateChanged(opts *bind.FilterOpts) (*LargeStakingDaoELCommissionRateChangedIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "DaoELCommissionRateChanged")
	if err != nil {
		return nil, err
	}
	return &LargeStakingDaoELCommissionRateChangedIterator{contract: _LargeStaking.contract, event: "DaoELCommissionRateChanged", logs: logs, sub: sub}, nil
}

// WatchDaoELCommissionRateChanged is a free log subscription operation binding the contract event 0xf34be367dcf9d61f1d01f06f296eb532d42f925474645b42aa776dd4c83900fc.
//
// Solidity: event DaoELCommissionRateChanged(uint256 _oldDaoElCommissionRate, uint256 _daoElCommissionRate)
func (_LargeStaking *LargeStakingFilterer) WatchDaoELCommissionRateChanged(opts *bind.WatchOpts, sink chan<- *LargeStakingDaoELCommissionRateChanged) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "DaoELCommissionRateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingDaoELCommissionRateChanged)
				if err := _LargeStaking.contract.UnpackLog(event, "DaoELCommissionRateChanged", log); err != nil {
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

// ParseDaoELCommissionRateChanged is a log parse operation binding the contract event 0xf34be367dcf9d61f1d01f06f296eb532d42f925474645b42aa776dd4c83900fc.
//
// Solidity: event DaoELCommissionRateChanged(uint256 _oldDaoElCommissionRate, uint256 _daoElCommissionRate)
func (_LargeStaking *LargeStakingFilterer) ParseDaoELCommissionRateChanged(log types.Log) (*LargeStakingDaoELCommissionRateChanged, error) {
	event := new(LargeStakingDaoELCommissionRateChanged)
	if err := _LargeStaking.contract.UnpackLog(event, "DaoELCommissionRateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingDaoPrivateRewardClaimedIterator is returned from FilterDaoPrivateRewardClaimed and is used to iterate over the raw logs and unpacked data for DaoPrivateRewardClaimed events raised by the LargeStaking contract.
type LargeStakingDaoPrivateRewardClaimedIterator struct {
	Event *LargeStakingDaoPrivateRewardClaimed // Event containing the contract specifics and raw log

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
func (it *LargeStakingDaoPrivateRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingDaoPrivateRewardClaimed)
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
		it.Event = new(LargeStakingDaoPrivateRewardClaimed)
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
func (it *LargeStakingDaoPrivateRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingDaoPrivateRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingDaoPrivateRewardClaimed represents a DaoPrivateRewardClaimed event raised by the LargeStaking contract.
type LargeStakingDaoPrivateRewardClaimed struct {
	StakingId       *big.Int
	DaoVaultAddress common.Address
	DaoRewards      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDaoPrivateRewardClaimed is a free log retrieval operation binding the contract event 0xd71df0c78871dfa6de8aca66c8226e204aafcc7ba4ee2e5321407268117e379a.
//
// Solidity: event DaoPrivateRewardClaimed(uint256 _stakingId, address _daoVaultAddress, uint256 _daoRewards)
func (_LargeStaking *LargeStakingFilterer) FilterDaoPrivateRewardClaimed(opts *bind.FilterOpts) (*LargeStakingDaoPrivateRewardClaimedIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "DaoPrivateRewardClaimed")
	if err != nil {
		return nil, err
	}
	return &LargeStakingDaoPrivateRewardClaimedIterator{contract: _LargeStaking.contract, event: "DaoPrivateRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchDaoPrivateRewardClaimed is a free log subscription operation binding the contract event 0xd71df0c78871dfa6de8aca66c8226e204aafcc7ba4ee2e5321407268117e379a.
//
// Solidity: event DaoPrivateRewardClaimed(uint256 _stakingId, address _daoVaultAddress, uint256 _daoRewards)
func (_LargeStaking *LargeStakingFilterer) WatchDaoPrivateRewardClaimed(opts *bind.WatchOpts, sink chan<- *LargeStakingDaoPrivateRewardClaimed) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "DaoPrivateRewardClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingDaoPrivateRewardClaimed)
				if err := _LargeStaking.contract.UnpackLog(event, "DaoPrivateRewardClaimed", log); err != nil {
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

// ParseDaoPrivateRewardClaimed is a log parse operation binding the contract event 0xd71df0c78871dfa6de8aca66c8226e204aafcc7ba4ee2e5321407268117e379a.
//
// Solidity: event DaoPrivateRewardClaimed(uint256 _stakingId, address _daoVaultAddress, uint256 _daoRewards)
func (_LargeStaking *LargeStakingFilterer) ParseDaoPrivateRewardClaimed(log types.Log) (*LargeStakingDaoPrivateRewardClaimed, error) {
	event := new(LargeStakingDaoPrivateRewardClaimed)
	if err := _LargeStaking.contract.UnpackLog(event, "DaoPrivateRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingDaoSharedRewardClaimedIterator is returned from FilterDaoSharedRewardClaimed and is used to iterate over the raw logs and unpacked data for DaoSharedRewardClaimed events raised by the LargeStaking contract.
type LargeStakingDaoSharedRewardClaimedIterator struct {
	Event *LargeStakingDaoSharedRewardClaimed // Event containing the contract specifics and raw log

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
func (it *LargeStakingDaoSharedRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingDaoSharedRewardClaimed)
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
		it.Event = new(LargeStakingDaoSharedRewardClaimed)
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
func (it *LargeStakingDaoSharedRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingDaoSharedRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingDaoSharedRewardClaimed represents a DaoSharedRewardClaimed event raised by the LargeStaking contract.
type LargeStakingDaoSharedRewardClaimed struct {
	OperatorId      *big.Int
	DaoVaultAddress common.Address
	DaoRewards      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDaoSharedRewardClaimed is a free log retrieval operation binding the contract event 0x4cad88612bcdabf2e7f7d86bb479e4d0d65d5bbba006328fd95092f0f901ea0f.
//
// Solidity: event DaoSharedRewardClaimed(uint256 _operatorId, address daoVaultAddress, uint256 _daoRewards)
func (_LargeStaking *LargeStakingFilterer) FilterDaoSharedRewardClaimed(opts *bind.FilterOpts) (*LargeStakingDaoSharedRewardClaimedIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "DaoSharedRewardClaimed")
	if err != nil {
		return nil, err
	}
	return &LargeStakingDaoSharedRewardClaimedIterator{contract: _LargeStaking.contract, event: "DaoSharedRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchDaoSharedRewardClaimed is a free log subscription operation binding the contract event 0x4cad88612bcdabf2e7f7d86bb479e4d0d65d5bbba006328fd95092f0f901ea0f.
//
// Solidity: event DaoSharedRewardClaimed(uint256 _operatorId, address daoVaultAddress, uint256 _daoRewards)
func (_LargeStaking *LargeStakingFilterer) WatchDaoSharedRewardClaimed(opts *bind.WatchOpts, sink chan<- *LargeStakingDaoSharedRewardClaimed) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "DaoSharedRewardClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingDaoSharedRewardClaimed)
				if err := _LargeStaking.contract.UnpackLog(event, "DaoSharedRewardClaimed", log); err != nil {
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

// ParseDaoSharedRewardClaimed is a log parse operation binding the contract event 0x4cad88612bcdabf2e7f7d86bb479e4d0d65d5bbba006328fd95092f0f901ea0f.
//
// Solidity: event DaoSharedRewardClaimed(uint256 _operatorId, address daoVaultAddress, uint256 _daoRewards)
func (_LargeStaking *LargeStakingFilterer) ParseDaoSharedRewardClaimed(log types.Log) (*LargeStakingDaoSharedRewardClaimed, error) {
	event := new(LargeStakingDaoSharedRewardClaimed)
	if err := _LargeStaking.contract.UnpackLog(event, "DaoSharedRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingDaoVaultAddressChangedIterator is returned from FilterDaoVaultAddressChanged and is used to iterate over the raw logs and unpacked data for DaoVaultAddressChanged events raised by the LargeStaking contract.
type LargeStakingDaoVaultAddressChangedIterator struct {
	Event *LargeStakingDaoVaultAddressChanged // Event containing the contract specifics and raw log

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
func (it *LargeStakingDaoVaultAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingDaoVaultAddressChanged)
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
		it.Event = new(LargeStakingDaoVaultAddressChanged)
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
func (it *LargeStakingDaoVaultAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingDaoVaultAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingDaoVaultAddressChanged represents a DaoVaultAddressChanged event raised by the LargeStaking contract.
type LargeStakingDaoVaultAddressChanged struct {
	OldDaoVaultAddress common.Address
	DaoVaultAddress    common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDaoVaultAddressChanged is a free log retrieval operation binding the contract event 0x74f93434acf49508438eb6f219ca22e7e1818b620ccb7acd411c8f520b27b642.
//
// Solidity: event DaoVaultAddressChanged(address _oldDaoVaultAddress, address _daoVaultAddress)
func (_LargeStaking *LargeStakingFilterer) FilterDaoVaultAddressChanged(opts *bind.FilterOpts) (*LargeStakingDaoVaultAddressChangedIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "DaoVaultAddressChanged")
	if err != nil {
		return nil, err
	}
	return &LargeStakingDaoVaultAddressChangedIterator{contract: _LargeStaking.contract, event: "DaoVaultAddressChanged", logs: logs, sub: sub}, nil
}

// WatchDaoVaultAddressChanged is a free log subscription operation binding the contract event 0x74f93434acf49508438eb6f219ca22e7e1818b620ccb7acd411c8f520b27b642.
//
// Solidity: event DaoVaultAddressChanged(address _oldDaoVaultAddress, address _daoVaultAddress)
func (_LargeStaking *LargeStakingFilterer) WatchDaoVaultAddressChanged(opts *bind.WatchOpts, sink chan<- *LargeStakingDaoVaultAddressChanged) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "DaoVaultAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingDaoVaultAddressChanged)
				if err := _LargeStaking.contract.UnpackLog(event, "DaoVaultAddressChanged", log); err != nil {
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
func (_LargeStaking *LargeStakingFilterer) ParseDaoVaultAddressChanged(log types.Log) (*LargeStakingDaoVaultAddressChanged, error) {
	event := new(LargeStakingDaoVaultAddressChanged)
	if err := _LargeStaking.contract.UnpackLog(event, "DaoVaultAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingELRewardFactoryChangedIterator is returned from FilterELRewardFactoryChanged and is used to iterate over the raw logs and unpacked data for ELRewardFactoryChanged events raised by the LargeStaking contract.
type LargeStakingELRewardFactoryChangedIterator struct {
	Event *LargeStakingELRewardFactoryChanged // Event containing the contract specifics and raw log

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
func (it *LargeStakingELRewardFactoryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingELRewardFactoryChanged)
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
		it.Event = new(LargeStakingELRewardFactoryChanged)
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
func (it *LargeStakingELRewardFactoryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingELRewardFactoryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingELRewardFactoryChanged represents a ELRewardFactoryChanged event raised by the LargeStaking contract.
type LargeStakingELRewardFactoryChanged struct {
	OldElRewardFactory common.Address
	ElRewardFactory    common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterELRewardFactoryChanged is a free log retrieval operation binding the contract event 0x99b8b92216a9f1b6dc61530456e6b6534e31be66c79cbc374f522fa0f0e8dafa.
//
// Solidity: event ELRewardFactoryChanged(address _oldElRewardFactory, address _elRewardFactory)
func (_LargeStaking *LargeStakingFilterer) FilterELRewardFactoryChanged(opts *bind.FilterOpts) (*LargeStakingELRewardFactoryChangedIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "ELRewardFactoryChanged")
	if err != nil {
		return nil, err
	}
	return &LargeStakingELRewardFactoryChangedIterator{contract: _LargeStaking.contract, event: "ELRewardFactoryChanged", logs: logs, sub: sub}, nil
}

// WatchELRewardFactoryChanged is a free log subscription operation binding the contract event 0x99b8b92216a9f1b6dc61530456e6b6534e31be66c79cbc374f522fa0f0e8dafa.
//
// Solidity: event ELRewardFactoryChanged(address _oldElRewardFactory, address _elRewardFactory)
func (_LargeStaking *LargeStakingFilterer) WatchELRewardFactoryChanged(opts *bind.WatchOpts, sink chan<- *LargeStakingELRewardFactoryChanged) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "ELRewardFactoryChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingELRewardFactoryChanged)
				if err := _LargeStaking.contract.UnpackLog(event, "ELRewardFactoryChanged", log); err != nil {
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

// ParseELRewardFactoryChanged is a log parse operation binding the contract event 0x99b8b92216a9f1b6dc61530456e6b6534e31be66c79cbc374f522fa0f0e8dafa.
//
// Solidity: event ELRewardFactoryChanged(address _oldElRewardFactory, address _elRewardFactory)
func (_LargeStaking *LargeStakingFilterer) ParseELRewardFactoryChanged(log types.Log) (*LargeStakingELRewardFactoryChanged, error) {
	event := new(LargeStakingELRewardFactoryChanged)
	if err := _LargeStaking.contract.UnpackLog(event, "ELRewardFactoryChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingELShareingRewardSettleIterator is returned from FilterELShareingRewardSettle and is used to iterate over the raw logs and unpacked data for ELShareingRewardSettle events raised by the LargeStaking contract.
type LargeStakingELShareingRewardSettleIterator struct {
	Event *LargeStakingELShareingRewardSettle // Event containing the contract specifics and raw log

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
func (it *LargeStakingELShareingRewardSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingELShareingRewardSettle)
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
		it.Event = new(LargeStakingELShareingRewardSettle)
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
func (it *LargeStakingELShareingRewardSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingELShareingRewardSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingELShareingRewardSettle represents a ELShareingRewardSettle event raised by the LargeStaking contract.
type LargeStakingELShareingRewardSettle struct {
	OperatorId     *big.Int
	DaoReward      *big.Int
	OperatorReward *big.Int
	PoolReward     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterELShareingRewardSettle is a free log retrieval operation binding the contract event 0x5c7564ba828519c8949300d92ff3ca573f105612c557187056cef4aace30debd.
//
// Solidity: event ELShareingRewardSettle(uint256 _operatorId, uint256 _daoReward, uint256 _operatorReward, uint256 _poolReward)
func (_LargeStaking *LargeStakingFilterer) FilterELShareingRewardSettle(opts *bind.FilterOpts) (*LargeStakingELShareingRewardSettleIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "ELShareingRewardSettle")
	if err != nil {
		return nil, err
	}
	return &LargeStakingELShareingRewardSettleIterator{contract: _LargeStaking.contract, event: "ELShareingRewardSettle", logs: logs, sub: sub}, nil
}

// WatchELShareingRewardSettle is a free log subscription operation binding the contract event 0x5c7564ba828519c8949300d92ff3ca573f105612c557187056cef4aace30debd.
//
// Solidity: event ELShareingRewardSettle(uint256 _operatorId, uint256 _daoReward, uint256 _operatorReward, uint256 _poolReward)
func (_LargeStaking *LargeStakingFilterer) WatchELShareingRewardSettle(opts *bind.WatchOpts, sink chan<- *LargeStakingELShareingRewardSettle) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "ELShareingRewardSettle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingELShareingRewardSettle)
				if err := _LargeStaking.contract.UnpackLog(event, "ELShareingRewardSettle", log); err != nil {
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

// ParseELShareingRewardSettle is a log parse operation binding the contract event 0x5c7564ba828519c8949300d92ff3ca573f105612c557187056cef4aace30debd.
//
// Solidity: event ELShareingRewardSettle(uint256 _operatorId, uint256 _daoReward, uint256 _operatorReward, uint256 _poolReward)
func (_LargeStaking *LargeStakingFilterer) ParseELShareingRewardSettle(log types.Log) (*LargeStakingELShareingRewardSettle, error) {
	event := new(LargeStakingELShareingRewardSettle)
	if err := _LargeStaking.contract.UnpackLog(event, "ELShareingRewardSettle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingElPrivateRewardSettleIterator is returned from FilterElPrivateRewardSettle and is used to iterate over the raw logs and unpacked data for ElPrivateRewardSettle events raised by the LargeStaking contract.
type LargeStakingElPrivateRewardSettleIterator struct {
	Event *LargeStakingElPrivateRewardSettle // Event containing the contract specifics and raw log

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
func (it *LargeStakingElPrivateRewardSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingElPrivateRewardSettle)
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
		it.Event = new(LargeStakingElPrivateRewardSettle)
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
func (it *LargeStakingElPrivateRewardSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingElPrivateRewardSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingElPrivateRewardSettle represents a ElPrivateRewardSettle event raised by the LargeStaking contract.
type LargeStakingElPrivateRewardSettle struct {
	StakingId      *big.Int
	OperatorId     *big.Int
	DaoReward      *big.Int
	OperatorReward *big.Int
	PoolReward     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterElPrivateRewardSettle is a free log retrieval operation binding the contract event 0x37820977d3d5e708af6aeea312e14a4e5472dbcd6a4c17c726cad2e8803dbebf.
//
// Solidity: event ElPrivateRewardSettle(uint256 _stakingId, uint256 _operatorId, uint256 _daoReward, uint256 _operatorReward, uint256 _poolReward)
func (_LargeStaking *LargeStakingFilterer) FilterElPrivateRewardSettle(opts *bind.FilterOpts) (*LargeStakingElPrivateRewardSettleIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "ElPrivateRewardSettle")
	if err != nil {
		return nil, err
	}
	return &LargeStakingElPrivateRewardSettleIterator{contract: _LargeStaking.contract, event: "ElPrivateRewardSettle", logs: logs, sub: sub}, nil
}

// WatchElPrivateRewardSettle is a free log subscription operation binding the contract event 0x37820977d3d5e708af6aeea312e14a4e5472dbcd6a4c17c726cad2e8803dbebf.
//
// Solidity: event ElPrivateRewardSettle(uint256 _stakingId, uint256 _operatorId, uint256 _daoReward, uint256 _operatorReward, uint256 _poolReward)
func (_LargeStaking *LargeStakingFilterer) WatchElPrivateRewardSettle(opts *bind.WatchOpts, sink chan<- *LargeStakingElPrivateRewardSettle) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "ElPrivateRewardSettle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingElPrivateRewardSettle)
				if err := _LargeStaking.contract.UnpackLog(event, "ElPrivateRewardSettle", log); err != nil {
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

// ParseElPrivateRewardSettle is a log parse operation binding the contract event 0x37820977d3d5e708af6aeea312e14a4e5472dbcd6a4c17c726cad2e8803dbebf.
//
// Solidity: event ElPrivateRewardSettle(uint256 _stakingId, uint256 _operatorId, uint256 _daoReward, uint256 _operatorReward, uint256 _poolReward)
func (_LargeStaking *LargeStakingFilterer) ParseElPrivateRewardSettle(log types.Log) (*LargeStakingElPrivateRewardSettle, error) {
	event := new(LargeStakingElPrivateRewardSettle)
	if err := _LargeStaking.contract.UnpackLog(event, "ElPrivateRewardSettle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingFastUnstakeIterator is returned from FilterFastUnstake and is used to iterate over the raw logs and unpacked data for FastUnstake events raised by the LargeStaking contract.
type LargeStakingFastUnstakeIterator struct {
	Event *LargeStakingFastUnstake // Event containing the contract specifics and raw log

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
func (it *LargeStakingFastUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingFastUnstake)
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
		it.Event = new(LargeStakingFastUnstake)
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
func (it *LargeStakingFastUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingFastUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingFastUnstake represents a FastUnstake event raised by the LargeStaking contract.
type LargeStakingFastUnstake struct {
	StakingId     *big.Int
	UnstakeAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFastUnstake is a free log retrieval operation binding the contract event 0x54cd8fd499e21e87c30d4dbffbe19a0001791d18fa3dd4912dd507d69915ff2d.
//
// Solidity: event FastUnstake(uint256 _stakingId, uint256 _unstakeAmount)
func (_LargeStaking *LargeStakingFilterer) FilterFastUnstake(opts *bind.FilterOpts) (*LargeStakingFastUnstakeIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "FastUnstake")
	if err != nil {
		return nil, err
	}
	return &LargeStakingFastUnstakeIterator{contract: _LargeStaking.contract, event: "FastUnstake", logs: logs, sub: sub}, nil
}

// WatchFastUnstake is a free log subscription operation binding the contract event 0x54cd8fd499e21e87c30d4dbffbe19a0001791d18fa3dd4912dd507d69915ff2d.
//
// Solidity: event FastUnstake(uint256 _stakingId, uint256 _unstakeAmount)
func (_LargeStaking *LargeStakingFilterer) WatchFastUnstake(opts *bind.WatchOpts, sink chan<- *LargeStakingFastUnstake) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "FastUnstake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingFastUnstake)
				if err := _LargeStaking.contract.UnpackLog(event, "FastUnstake", log); err != nil {
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

// ParseFastUnstake is a log parse operation binding the contract event 0x54cd8fd499e21e87c30d4dbffbe19a0001791d18fa3dd4912dd507d69915ff2d.
//
// Solidity: event FastUnstake(uint256 _stakingId, uint256 _unstakeAmount)
func (_LargeStaking *LargeStakingFilterer) ParseFastUnstake(log types.Log) (*LargeStakingFastUnstake, error) {
	event := new(LargeStakingFastUnstake)
	if err := _LargeStaking.contract.UnpackLog(event, "FastUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LargeStaking contract.
type LargeStakingInitializedIterator struct {
	Event *LargeStakingInitialized // Event containing the contract specifics and raw log

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
func (it *LargeStakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingInitialized)
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
		it.Event = new(LargeStakingInitialized)
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
func (it *LargeStakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingInitialized represents a Initialized event raised by the LargeStaking contract.
type LargeStakingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LargeStaking *LargeStakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*LargeStakingInitializedIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LargeStakingInitializedIterator{contract: _LargeStaking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LargeStaking *LargeStakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LargeStakingInitialized) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingInitialized)
				if err := _LargeStaking.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LargeStaking *LargeStakingFilterer) ParseInitialized(log types.Log) (*LargeStakingInitialized, error) {
	event := new(LargeStakingInitialized)
	if err := _LargeStaking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingLargeStakeIterator is returned from FilterLargeStake and is used to iterate over the raw logs and unpacked data for LargeStake events raised by the LargeStaking contract.
type LargeStakingLargeStakeIterator struct {
	Event *LargeStakingLargeStake // Event containing the contract specifics and raw log

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
func (it *LargeStakingLargeStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingLargeStake)
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
		it.Event = new(LargeStakingLargeStake)
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
func (it *LargeStakingLargeStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingLargeStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingLargeStake represents a LargeStake event raised by the LargeStaking contract.
type LargeStakingLargeStake struct {
	OperatorId          *big.Int
	CurStakingId        *big.Int
	Amount              *big.Int
	Owner               common.Address
	ElRewardAddr        common.Address
	WithdrawCredentials common.Address
	IsELRewardSharing   bool
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLargeStake is a free log retrieval operation binding the contract event 0x7567223d27cdfa09492d4fc8db4b8bf0a75b6c76ee8a15c9cc49bc5977b357c0.
//
// Solidity: event LargeStake(uint256 _operatorId, uint256 _curStakingId, uint256 _amount, address _owner, address _elRewardAddr, address _withdrawCredentials, bool _isELRewardSharing)
func (_LargeStaking *LargeStakingFilterer) FilterLargeStake(opts *bind.FilterOpts) (*LargeStakingLargeStakeIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "LargeStake")
	if err != nil {
		return nil, err
	}
	return &LargeStakingLargeStakeIterator{contract: _LargeStaking.contract, event: "LargeStake", logs: logs, sub: sub}, nil
}

// WatchLargeStake is a free log subscription operation binding the contract event 0x7567223d27cdfa09492d4fc8db4b8bf0a75b6c76ee8a15c9cc49bc5977b357c0.
//
// Solidity: event LargeStake(uint256 _operatorId, uint256 _curStakingId, uint256 _amount, address _owner, address _elRewardAddr, address _withdrawCredentials, bool _isELRewardSharing)
func (_LargeStaking *LargeStakingFilterer) WatchLargeStake(opts *bind.WatchOpts, sink chan<- *LargeStakingLargeStake) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "LargeStake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingLargeStake)
				if err := _LargeStaking.contract.UnpackLog(event, "LargeStake", log); err != nil {
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

// ParseLargeStake is a log parse operation binding the contract event 0x7567223d27cdfa09492d4fc8db4b8bf0a75b6c76ee8a15c9cc49bc5977b357c0.
//
// Solidity: event LargeStake(uint256 _operatorId, uint256 _curStakingId, uint256 _amount, address _owner, address _elRewardAddr, address _withdrawCredentials, bool _isELRewardSharing)
func (_LargeStaking *LargeStakingFilterer) ParseLargeStake(log types.Log) (*LargeStakingLargeStake, error) {
	event := new(LargeStakingLargeStake)
	if err := _LargeStaking.contract.UnpackLog(event, "LargeStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingLargeStakingSlashIterator is returned from FilterLargeStakingSlash and is used to iterate over the raw logs and unpacked data for LargeStakingSlash events raised by the LargeStaking contract.
type LargeStakingLargeStakingSlashIterator struct {
	Event *LargeStakingLargeStakingSlash // Event containing the contract specifics and raw log

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
func (it *LargeStakingLargeStakingSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingLargeStakingSlash)
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
		it.Event = new(LargeStakingLargeStakingSlash)
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
func (it *LargeStakingLargeStakingSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingLargeStakingSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingLargeStakingSlash represents a LargeStakingSlash event raised by the LargeStaking contract.
type LargeStakingLargeStakingSlash struct {
	StakingIds  *big.Int
	OperatorIds *big.Int
	Pubkey      []byte
	Amounts     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLargeStakingSlash is a free log retrieval operation binding the contract event 0x188b7d5e7b342b5b902a9451b2cb773d65c960f0040d8f6ce0bc239ac6087641.
//
// Solidity: event LargeStakingSlash(uint256 _stakingIds, uint256 _operatorIds, bytes _pubkey, uint256 _amounts)
func (_LargeStaking *LargeStakingFilterer) FilterLargeStakingSlash(opts *bind.FilterOpts) (*LargeStakingLargeStakingSlashIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "LargeStakingSlash")
	if err != nil {
		return nil, err
	}
	return &LargeStakingLargeStakingSlashIterator{contract: _LargeStaking.contract, event: "LargeStakingSlash", logs: logs, sub: sub}, nil
}

// WatchLargeStakingSlash is a free log subscription operation binding the contract event 0x188b7d5e7b342b5b902a9451b2cb773d65c960f0040d8f6ce0bc239ac6087641.
//
// Solidity: event LargeStakingSlash(uint256 _stakingIds, uint256 _operatorIds, bytes _pubkey, uint256 _amounts)
func (_LargeStaking *LargeStakingFilterer) WatchLargeStakingSlash(opts *bind.WatchOpts, sink chan<- *LargeStakingLargeStakingSlash) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "LargeStakingSlash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingLargeStakingSlash)
				if err := _LargeStaking.contract.UnpackLog(event, "LargeStakingSlash", log); err != nil {
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

// ParseLargeStakingSlash is a log parse operation binding the contract event 0x188b7d5e7b342b5b902a9451b2cb773d65c960f0040d8f6ce0bc239ac6087641.
//
// Solidity: event LargeStakingSlash(uint256 _stakingIds, uint256 _operatorIds, bytes _pubkey, uint256 _amounts)
func (_LargeStaking *LargeStakingFilterer) ParseLargeStakingSlash(log types.Log) (*LargeStakingLargeStakingSlash, error) {
	event := new(LargeStakingLargeStakingSlash)
	if err := _LargeStaking.contract.UnpackLog(event, "LargeStakingSlash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingLargeUnstakeIterator is returned from FilterLargeUnstake and is used to iterate over the raw logs and unpacked data for LargeUnstake events raised by the LargeStaking contract.
type LargeStakingLargeUnstakeIterator struct {
	Event *LargeStakingLargeUnstake // Event containing the contract specifics and raw log

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
func (it *LargeStakingLargeUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingLargeUnstake)
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
		it.Event = new(LargeStakingLargeUnstake)
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
func (it *LargeStakingLargeUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingLargeUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingLargeUnstake represents a LargeUnstake event raised by the LargeStaking contract.
type LargeStakingLargeUnstake struct {
	StakingId *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLargeUnstake is a free log retrieval operation binding the contract event 0xb9be9207a343d7218aad56bfb9a9dfe97c57d260f539ef1ec6875cf4170a42ac.
//
// Solidity: event LargeUnstake(uint256 _stakingId, uint256 _amount)
func (_LargeStaking *LargeStakingFilterer) FilterLargeUnstake(opts *bind.FilterOpts) (*LargeStakingLargeUnstakeIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "LargeUnstake")
	if err != nil {
		return nil, err
	}
	return &LargeStakingLargeUnstakeIterator{contract: _LargeStaking.contract, event: "LargeUnstake", logs: logs, sub: sub}, nil
}

// WatchLargeUnstake is a free log subscription operation binding the contract event 0xb9be9207a343d7218aad56bfb9a9dfe97c57d260f539ef1ec6875cf4170a42ac.
//
// Solidity: event LargeUnstake(uint256 _stakingId, uint256 _amount)
func (_LargeStaking *LargeStakingFilterer) WatchLargeUnstake(opts *bind.WatchOpts, sink chan<- *LargeStakingLargeUnstake) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "LargeUnstake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingLargeUnstake)
				if err := _LargeStaking.contract.UnpackLog(event, "LargeUnstake", log); err != nil {
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

// ParseLargeUnstake is a log parse operation binding the contract event 0xb9be9207a343d7218aad56bfb9a9dfe97c57d260f539ef1ec6875cf4170a42ac.
//
// Solidity: event LargeUnstake(uint256 _stakingId, uint256 _amount)
func (_LargeStaking *LargeStakingFilterer) ParseLargeUnstake(log types.Log) (*LargeStakingLargeUnstake, error) {
	event := new(LargeStakingLargeUnstake)
	if err := _LargeStaking.contract.UnpackLog(event, "LargeUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingMigretaStakeIterator is returned from FilterMigretaStake and is used to iterate over the raw logs and unpacked data for MigretaStake events raised by the LargeStaking contract.
type LargeStakingMigretaStakeIterator struct {
	Event *LargeStakingMigretaStake // Event containing the contract specifics and raw log

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
func (it *LargeStakingMigretaStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingMigretaStake)
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
		it.Event = new(LargeStakingMigretaStake)
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
func (it *LargeStakingMigretaStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingMigretaStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingMigretaStake represents a MigretaStake event raised by the LargeStaking contract.
type LargeStakingMigretaStake struct {
	OperatorId          *big.Int
	CurStakingId        *big.Int
	Amount              *big.Int
	Owner               common.Address
	ElRewardAddr        common.Address
	WithdrawCredentials common.Address
	IsELRewardSharing   bool
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMigretaStake is a free log retrieval operation binding the contract event 0x92819bcbc0f10468f0586226ea0807f64d52875dff93b601fa14e9189f552874.
//
// Solidity: event MigretaStake(uint256 _operatorId, uint256 _curStakingId, uint256 _amount, address _owner, address _elRewardAddr, address _withdrawCredentials, bool _isELRewardSharing)
func (_LargeStaking *LargeStakingFilterer) FilterMigretaStake(opts *bind.FilterOpts) (*LargeStakingMigretaStakeIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "MigretaStake")
	if err != nil {
		return nil, err
	}
	return &LargeStakingMigretaStakeIterator{contract: _LargeStaking.contract, event: "MigretaStake", logs: logs, sub: sub}, nil
}

// WatchMigretaStake is a free log subscription operation binding the contract event 0x92819bcbc0f10468f0586226ea0807f64d52875dff93b601fa14e9189f552874.
//
// Solidity: event MigretaStake(uint256 _operatorId, uint256 _curStakingId, uint256 _amount, address _owner, address _elRewardAddr, address _withdrawCredentials, bool _isELRewardSharing)
func (_LargeStaking *LargeStakingFilterer) WatchMigretaStake(opts *bind.WatchOpts, sink chan<- *LargeStakingMigretaStake) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "MigretaStake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingMigretaStake)
				if err := _LargeStaking.contract.UnpackLog(event, "MigretaStake", log); err != nil {
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

// ParseMigretaStake is a log parse operation binding the contract event 0x92819bcbc0f10468f0586226ea0807f64d52875dff93b601fa14e9189f552874.
//
// Solidity: event MigretaStake(uint256 _operatorId, uint256 _curStakingId, uint256 _amount, address _owner, address _elRewardAddr, address _withdrawCredentials, bool _isELRewardSharing)
func (_LargeStaking *LargeStakingFilterer) ParseMigretaStake(log types.Log) (*LargeStakingMigretaStake, error) {
	event := new(LargeStakingMigretaStake)
	if err := _LargeStaking.contract.UnpackLog(event, "MigretaStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingMinStakeAmountChangeIterator is returned from FilterMinStakeAmountChange and is used to iterate over the raw logs and unpacked data for MinStakeAmountChange events raised by the LargeStaking contract.
type LargeStakingMinStakeAmountChangeIterator struct {
	Event *LargeStakingMinStakeAmountChange // Event containing the contract specifics and raw log

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
func (it *LargeStakingMinStakeAmountChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingMinStakeAmountChange)
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
		it.Event = new(LargeStakingMinStakeAmountChange)
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
func (it *LargeStakingMinStakeAmountChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingMinStakeAmountChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingMinStakeAmountChange represents a MinStakeAmountChange event raised by the LargeStaking contract.
type LargeStakingMinStakeAmountChange struct {
	OldMinStakeAmount *big.Int
	MinStakeAmount    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMinStakeAmountChange is a free log retrieval operation binding the contract event 0x59993a665fbc5ab65fc6d986fb877821545f1dbfe0d7bc73ef1db7953734a904.
//
// Solidity: event MinStakeAmountChange(uint256 _oldMinStakeAmount, uint256 _minStakeAmount)
func (_LargeStaking *LargeStakingFilterer) FilterMinStakeAmountChange(opts *bind.FilterOpts) (*LargeStakingMinStakeAmountChangeIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "MinStakeAmountChange")
	if err != nil {
		return nil, err
	}
	return &LargeStakingMinStakeAmountChangeIterator{contract: _LargeStaking.contract, event: "MinStakeAmountChange", logs: logs, sub: sub}, nil
}

// WatchMinStakeAmountChange is a free log subscription operation binding the contract event 0x59993a665fbc5ab65fc6d986fb877821545f1dbfe0d7bc73ef1db7953734a904.
//
// Solidity: event MinStakeAmountChange(uint256 _oldMinStakeAmount, uint256 _minStakeAmount)
func (_LargeStaking *LargeStakingFilterer) WatchMinStakeAmountChange(opts *bind.WatchOpts, sink chan<- *LargeStakingMinStakeAmountChange) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "MinStakeAmountChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingMinStakeAmountChange)
				if err := _LargeStaking.contract.UnpackLog(event, "MinStakeAmountChange", log); err != nil {
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

// ParseMinStakeAmountChange is a log parse operation binding the contract event 0x59993a665fbc5ab65fc6d986fb877821545f1dbfe0d7bc73ef1db7953734a904.
//
// Solidity: event MinStakeAmountChange(uint256 _oldMinStakeAmount, uint256 _minStakeAmount)
func (_LargeStaking *LargeStakingFilterer) ParseMinStakeAmountChange(log types.Log) (*LargeStakingMinStakeAmountChange, error) {
	event := new(LargeStakingMinStakeAmountChange)
	if err := _LargeStaking.contract.UnpackLog(event, "MinStakeAmountChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingNodeOperatorsRegistryChangedIterator is returned from FilterNodeOperatorsRegistryChanged and is used to iterate over the raw logs and unpacked data for NodeOperatorsRegistryChanged events raised by the LargeStaking contract.
type LargeStakingNodeOperatorsRegistryChangedIterator struct {
	Event *LargeStakingNodeOperatorsRegistryChanged // Event containing the contract specifics and raw log

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
func (it *LargeStakingNodeOperatorsRegistryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingNodeOperatorsRegistryChanged)
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
		it.Event = new(LargeStakingNodeOperatorsRegistryChanged)
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
func (it *LargeStakingNodeOperatorsRegistryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingNodeOperatorsRegistryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingNodeOperatorsRegistryChanged represents a NodeOperatorsRegistryChanged event raised by the LargeStaking contract.
type LargeStakingNodeOperatorsRegistryChanged struct {
	OldNodeOperatorRegistryContract common.Address
	NodeOperatorRegistryAddress     common.Address
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorsRegistryChanged is a free log retrieval operation binding the contract event 0x59d27bf723582840cd845203a8055dd24a721e04a47b7d9d602804464416f93c.
//
// Solidity: event NodeOperatorsRegistryChanged(address _oldNodeOperatorRegistryContract, address _nodeOperatorRegistryAddress)
func (_LargeStaking *LargeStakingFilterer) FilterNodeOperatorsRegistryChanged(opts *bind.FilterOpts) (*LargeStakingNodeOperatorsRegistryChangedIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "NodeOperatorsRegistryChanged")
	if err != nil {
		return nil, err
	}
	return &LargeStakingNodeOperatorsRegistryChangedIterator{contract: _LargeStaking.contract, event: "NodeOperatorsRegistryChanged", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorsRegistryChanged is a free log subscription operation binding the contract event 0x59d27bf723582840cd845203a8055dd24a721e04a47b7d9d602804464416f93c.
//
// Solidity: event NodeOperatorsRegistryChanged(address _oldNodeOperatorRegistryContract, address _nodeOperatorRegistryAddress)
func (_LargeStaking *LargeStakingFilterer) WatchNodeOperatorsRegistryChanged(opts *bind.WatchOpts, sink chan<- *LargeStakingNodeOperatorsRegistryChanged) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "NodeOperatorsRegistryChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingNodeOperatorsRegistryChanged)
				if err := _LargeStaking.contract.UnpackLog(event, "NodeOperatorsRegistryChanged", log); err != nil {
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

// ParseNodeOperatorsRegistryChanged is a log parse operation binding the contract event 0x59d27bf723582840cd845203a8055dd24a721e04a47b7d9d602804464416f93c.
//
// Solidity: event NodeOperatorsRegistryChanged(address _oldNodeOperatorRegistryContract, address _nodeOperatorRegistryAddress)
func (_LargeStaking *LargeStakingFilterer) ParseNodeOperatorsRegistryChanged(log types.Log) (*LargeStakingNodeOperatorsRegistryChanged, error) {
	event := new(LargeStakingNodeOperatorsRegistryChanged)
	if err := _LargeStaking.contract.UnpackLog(event, "NodeOperatorsRegistryChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingOperatorPrivateRewardClaimedIterator is returned from FilterOperatorPrivateRewardClaimed and is used to iterate over the raw logs and unpacked data for OperatorPrivateRewardClaimed events raised by the LargeStaking contract.
type LargeStakingOperatorPrivateRewardClaimedIterator struct {
	Event *LargeStakingOperatorPrivateRewardClaimed // Event containing the contract specifics and raw log

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
func (it *LargeStakingOperatorPrivateRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingOperatorPrivateRewardClaimed)
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
		it.Event = new(LargeStakingOperatorPrivateRewardClaimed)
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
func (it *LargeStakingOperatorPrivateRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingOperatorPrivateRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingOperatorPrivateRewardClaimed represents a OperatorPrivateRewardClaimed event raised by the LargeStaking contract.
type LargeStakingOperatorPrivateRewardClaimed struct {
	StakingId       *big.Int
	OperatorId      *big.Int
	OperatorRewards *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorPrivateRewardClaimed is a free log retrieval operation binding the contract event 0xe44fbecaa7c5377853203a05bf4ea5592c59d9a41dbf26e36e7323652811cb56.
//
// Solidity: event OperatorPrivateRewardClaimed(uint256 _stakingId, uint256 _operatorId, uint256 _operatorRewards)
func (_LargeStaking *LargeStakingFilterer) FilterOperatorPrivateRewardClaimed(opts *bind.FilterOpts) (*LargeStakingOperatorPrivateRewardClaimedIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "OperatorPrivateRewardClaimed")
	if err != nil {
		return nil, err
	}
	return &LargeStakingOperatorPrivateRewardClaimedIterator{contract: _LargeStaking.contract, event: "OperatorPrivateRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchOperatorPrivateRewardClaimed is a free log subscription operation binding the contract event 0xe44fbecaa7c5377853203a05bf4ea5592c59d9a41dbf26e36e7323652811cb56.
//
// Solidity: event OperatorPrivateRewardClaimed(uint256 _stakingId, uint256 _operatorId, uint256 _operatorRewards)
func (_LargeStaking *LargeStakingFilterer) WatchOperatorPrivateRewardClaimed(opts *bind.WatchOpts, sink chan<- *LargeStakingOperatorPrivateRewardClaimed) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "OperatorPrivateRewardClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingOperatorPrivateRewardClaimed)
				if err := _LargeStaking.contract.UnpackLog(event, "OperatorPrivateRewardClaimed", log); err != nil {
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

// ParseOperatorPrivateRewardClaimed is a log parse operation binding the contract event 0xe44fbecaa7c5377853203a05bf4ea5592c59d9a41dbf26e36e7323652811cb56.
//
// Solidity: event OperatorPrivateRewardClaimed(uint256 _stakingId, uint256 _operatorId, uint256 _operatorRewards)
func (_LargeStaking *LargeStakingFilterer) ParseOperatorPrivateRewardClaimed(log types.Log) (*LargeStakingOperatorPrivateRewardClaimed, error) {
	event := new(LargeStakingOperatorPrivateRewardClaimed)
	if err := _LargeStaking.contract.UnpackLog(event, "OperatorPrivateRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingOperatorRewardClaimedIterator is returned from FilterOperatorRewardClaimed and is used to iterate over the raw logs and unpacked data for OperatorRewardClaimed events raised by the LargeStaking contract.
type LargeStakingOperatorRewardClaimedIterator struct {
	Event *LargeStakingOperatorRewardClaimed // Event containing the contract specifics and raw log

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
func (it *LargeStakingOperatorRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingOperatorRewardClaimed)
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
		it.Event = new(LargeStakingOperatorRewardClaimed)
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
func (it *LargeStakingOperatorRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingOperatorRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingOperatorRewardClaimed represents a OperatorRewardClaimed event raised by the LargeStaking contract.
type LargeStakingOperatorRewardClaimed struct {
	OperatorId      *big.Int
	RewardAddresses common.Address
	RewardAmounts   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorRewardClaimed is a free log retrieval operation binding the contract event 0x9f9012d7cf5ed4a0e865893d0f5bbd35ba0940078e3a64c9db605a29dd6b37d6.
//
// Solidity: event OperatorRewardClaimed(uint256 _operatorId, address _rewardAddresses, uint256 _rewardAmounts)
func (_LargeStaking *LargeStakingFilterer) FilterOperatorRewardClaimed(opts *bind.FilterOpts) (*LargeStakingOperatorRewardClaimedIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "OperatorRewardClaimed")
	if err != nil {
		return nil, err
	}
	return &LargeStakingOperatorRewardClaimedIterator{contract: _LargeStaking.contract, event: "OperatorRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchOperatorRewardClaimed is a free log subscription operation binding the contract event 0x9f9012d7cf5ed4a0e865893d0f5bbd35ba0940078e3a64c9db605a29dd6b37d6.
//
// Solidity: event OperatorRewardClaimed(uint256 _operatorId, address _rewardAddresses, uint256 _rewardAmounts)
func (_LargeStaking *LargeStakingFilterer) WatchOperatorRewardClaimed(opts *bind.WatchOpts, sink chan<- *LargeStakingOperatorRewardClaimed) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "OperatorRewardClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingOperatorRewardClaimed)
				if err := _LargeStaking.contract.UnpackLog(event, "OperatorRewardClaimed", log); err != nil {
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

// ParseOperatorRewardClaimed is a log parse operation binding the contract event 0x9f9012d7cf5ed4a0e865893d0f5bbd35ba0940078e3a64c9db605a29dd6b37d6.
//
// Solidity: event OperatorRewardClaimed(uint256 _operatorId, address _rewardAddresses, uint256 _rewardAmounts)
func (_LargeStaking *LargeStakingFilterer) ParseOperatorRewardClaimed(log types.Log) (*LargeStakingOperatorRewardClaimed, error) {
	event := new(LargeStakingOperatorRewardClaimed)
	if err := _LargeStaking.contract.UnpackLog(event, "OperatorRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingOperatorSharedRewardClaimedIterator is returned from FilterOperatorSharedRewardClaimed and is used to iterate over the raw logs and unpacked data for OperatorSharedRewardClaimed events raised by the LargeStaking contract.
type LargeStakingOperatorSharedRewardClaimedIterator struct {
	Event *LargeStakingOperatorSharedRewardClaimed // Event containing the contract specifics and raw log

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
func (it *LargeStakingOperatorSharedRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingOperatorSharedRewardClaimed)
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
		it.Event = new(LargeStakingOperatorSharedRewardClaimed)
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
func (it *LargeStakingOperatorSharedRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingOperatorSharedRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingOperatorSharedRewardClaimed represents a OperatorSharedRewardClaimed event raised by the LargeStaking contract.
type LargeStakingOperatorSharedRewardClaimed struct {
	OperatorId      *big.Int
	OperatorRewards *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharedRewardClaimed is a free log retrieval operation binding the contract event 0x2a18521abfa690fbb534172ea75c9c99a05aa1539de058ed0ddebf81f965a0a0.
//
// Solidity: event OperatorSharedRewardClaimed(uint256 _operatorId, uint256 _operatorRewards)
func (_LargeStaking *LargeStakingFilterer) FilterOperatorSharedRewardClaimed(opts *bind.FilterOpts) (*LargeStakingOperatorSharedRewardClaimedIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "OperatorSharedRewardClaimed")
	if err != nil {
		return nil, err
	}
	return &LargeStakingOperatorSharedRewardClaimedIterator{contract: _LargeStaking.contract, event: "OperatorSharedRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchOperatorSharedRewardClaimed is a free log subscription operation binding the contract event 0x2a18521abfa690fbb534172ea75c9c99a05aa1539de058ed0ddebf81f965a0a0.
//
// Solidity: event OperatorSharedRewardClaimed(uint256 _operatorId, uint256 _operatorRewards)
func (_LargeStaking *LargeStakingFilterer) WatchOperatorSharedRewardClaimed(opts *bind.WatchOpts, sink chan<- *LargeStakingOperatorSharedRewardClaimed) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "OperatorSharedRewardClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingOperatorSharedRewardClaimed)
				if err := _LargeStaking.contract.UnpackLog(event, "OperatorSharedRewardClaimed", log); err != nil {
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

// ParseOperatorSharedRewardClaimed is a log parse operation binding the contract event 0x2a18521abfa690fbb534172ea75c9c99a05aa1539de058ed0ddebf81f965a0a0.
//
// Solidity: event OperatorSharedRewardClaimed(uint256 _operatorId, uint256 _operatorRewards)
func (_LargeStaking *LargeStakingFilterer) ParseOperatorSharedRewardClaimed(log types.Log) (*LargeStakingOperatorSharedRewardClaimed, error) {
	event := new(LargeStakingOperatorSharedRewardClaimed)
	if err := _LargeStaking.contract.UnpackLog(event, "OperatorSharedRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingOperatorSlashChangedIterator is returned from FilterOperatorSlashChanged and is used to iterate over the raw logs and unpacked data for OperatorSlashChanged events raised by the LargeStaking contract.
type LargeStakingOperatorSlashChangedIterator struct {
	Event *LargeStakingOperatorSlashChanged // Event containing the contract specifics and raw log

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
func (it *LargeStakingOperatorSlashChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingOperatorSlashChanged)
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
		it.Event = new(LargeStakingOperatorSlashChanged)
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
func (it *LargeStakingOperatorSlashChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingOperatorSlashChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingOperatorSlashChanged represents a OperatorSlashChanged event raised by the LargeStaking contract.
type LargeStakingOperatorSlashChanged struct {
	OldOperatorSlashContract common.Address
	OperatorSlashContract    common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashChanged is a free log retrieval operation binding the contract event 0x3b64b88ced21a3f5bf79f1a7c04427dba599390f7dd89550737b43523e2fb0a5.
//
// Solidity: event OperatorSlashChanged(address _oldOperatorSlashContract, address _operatorSlashContract)
func (_LargeStaking *LargeStakingFilterer) FilterOperatorSlashChanged(opts *bind.FilterOpts) (*LargeStakingOperatorSlashChangedIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "OperatorSlashChanged")
	if err != nil {
		return nil, err
	}
	return &LargeStakingOperatorSlashChangedIterator{contract: _LargeStaking.contract, event: "OperatorSlashChanged", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashChanged is a free log subscription operation binding the contract event 0x3b64b88ced21a3f5bf79f1a7c04427dba599390f7dd89550737b43523e2fb0a5.
//
// Solidity: event OperatorSlashChanged(address _oldOperatorSlashContract, address _operatorSlashContract)
func (_LargeStaking *LargeStakingFilterer) WatchOperatorSlashChanged(opts *bind.WatchOpts, sink chan<- *LargeStakingOperatorSlashChanged) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "OperatorSlashChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingOperatorSlashChanged)
				if err := _LargeStaking.contract.UnpackLog(event, "OperatorSlashChanged", log); err != nil {
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

// ParseOperatorSlashChanged is a log parse operation binding the contract event 0x3b64b88ced21a3f5bf79f1a7c04427dba599390f7dd89550737b43523e2fb0a5.
//
// Solidity: event OperatorSlashChanged(address _oldOperatorSlashContract, address _operatorSlashContract)
func (_LargeStaking *LargeStakingFilterer) ParseOperatorSlashChanged(log types.Log) (*LargeStakingOperatorSlashChanged, error) {
	event := new(LargeStakingOperatorSlashChanged)
	if err := _LargeStaking.contract.UnpackLog(event, "OperatorSlashChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LargeStaking contract.
type LargeStakingOwnershipTransferredIterator struct {
	Event *LargeStakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LargeStakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingOwnershipTransferred)
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
		it.Event = new(LargeStakingOwnershipTransferred)
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
func (it *LargeStakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingOwnershipTransferred represents a OwnershipTransferred event raised by the LargeStaking contract.
type LargeStakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LargeStaking *LargeStakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LargeStakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LargeStakingOwnershipTransferredIterator{contract: _LargeStaking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LargeStaking *LargeStakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LargeStakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingOwnershipTransferred)
				if err := _LargeStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LargeStaking *LargeStakingFilterer) ParseOwnershipTransferred(log types.Log) (*LargeStakingOwnershipTransferred, error) {
	event := new(LargeStakingOwnershipTransferred)
	if err := _LargeStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingSharedRewardPoolStartIterator is returned from FilterSharedRewardPoolStart and is used to iterate over the raw logs and unpacked data for SharedRewardPoolStart events raised by the LargeStaking contract.
type LargeStakingSharedRewardPoolStartIterator struct {
	Event *LargeStakingSharedRewardPoolStart // Event containing the contract specifics and raw log

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
func (it *LargeStakingSharedRewardPoolStartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingSharedRewardPoolStart)
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
		it.Event = new(LargeStakingSharedRewardPoolStart)
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
func (it *LargeStakingSharedRewardPoolStartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingSharedRewardPoolStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingSharedRewardPoolStart represents a SharedRewardPoolStart event raised by the LargeStaking contract.
type LargeStakingSharedRewardPoolStart struct {
	OperatorId       *big.Int
	ElRewardPoolAddr common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSharedRewardPoolStart is a free log retrieval operation binding the contract event 0xd9a8b9b119f8c34d13ec8079deba98077519e6277ea4d364d61560e91284d83d.
//
// Solidity: event SharedRewardPoolStart(uint256 _operatorId, address _elRewardPoolAddr)
func (_LargeStaking *LargeStakingFilterer) FilterSharedRewardPoolStart(opts *bind.FilterOpts) (*LargeStakingSharedRewardPoolStartIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "SharedRewardPoolStart")
	if err != nil {
		return nil, err
	}
	return &LargeStakingSharedRewardPoolStartIterator{contract: _LargeStaking.contract, event: "SharedRewardPoolStart", logs: logs, sub: sub}, nil
}

// WatchSharedRewardPoolStart is a free log subscription operation binding the contract event 0xd9a8b9b119f8c34d13ec8079deba98077519e6277ea4d364d61560e91284d83d.
//
// Solidity: event SharedRewardPoolStart(uint256 _operatorId, address _elRewardPoolAddr)
func (_LargeStaking *LargeStakingFilterer) WatchSharedRewardPoolStart(opts *bind.WatchOpts, sink chan<- *LargeStakingSharedRewardPoolStart) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "SharedRewardPoolStart")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingSharedRewardPoolStart)
				if err := _LargeStaking.contract.UnpackLog(event, "SharedRewardPoolStart", log); err != nil {
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

// ParseSharedRewardPoolStart is a log parse operation binding the contract event 0xd9a8b9b119f8c34d13ec8079deba98077519e6277ea4d364d61560e91284d83d.
//
// Solidity: event SharedRewardPoolStart(uint256 _operatorId, address _elRewardPoolAddr)
func (_LargeStaking *LargeStakingFilterer) ParseSharedRewardPoolStart(log types.Log) (*LargeStakingSharedRewardPoolStart, error) {
	event := new(LargeStakingSharedRewardPoolStart)
	if err := _LargeStaking.contract.UnpackLog(event, "SharedRewardPoolStart", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the LargeStaking contract.
type LargeStakingUpgradedIterator struct {
	Event *LargeStakingUpgraded // Event containing the contract specifics and raw log

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
func (it *LargeStakingUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingUpgraded)
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
		it.Event = new(LargeStakingUpgraded)
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
func (it *LargeStakingUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingUpgraded represents a Upgraded event raised by the LargeStaking contract.
type LargeStakingUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LargeStaking *LargeStakingFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LargeStakingUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LargeStakingUpgradedIterator{contract: _LargeStaking.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LargeStaking *LargeStakingFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LargeStakingUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingUpgraded)
				if err := _LargeStaking.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_LargeStaking *LargeStakingFilterer) ParseUpgraded(log types.Log) (*LargeStakingUpgraded, error) {
	event := new(LargeStakingUpgraded)
	if err := _LargeStaking.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingUserRewardClaimedIterator is returned from FilterUserRewardClaimed and is used to iterate over the raw logs and unpacked data for UserRewardClaimed events raised by the LargeStaking contract.
type LargeStakingUserRewardClaimedIterator struct {
	Event *LargeStakingUserRewardClaimed // Event containing the contract specifics and raw log

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
func (it *LargeStakingUserRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingUserRewardClaimed)
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
		it.Event = new(LargeStakingUserRewardClaimed)
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
func (it *LargeStakingUserRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingUserRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingUserRewardClaimed represents a UserRewardClaimed event raised by the LargeStaking contract.
type LargeStakingUserRewardClaimed struct {
	StakingId   *big.Int
	Beneficiary common.Address
	Rewards     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUserRewardClaimed is a free log retrieval operation binding the contract event 0x10b9969d2c05a458029b40aed2ee60918d51f8feb1c2285d6cb1dde5034a12aa.
//
// Solidity: event UserRewardClaimed(uint256 _stakingId, address _beneficiary, uint256 _rewards)
func (_LargeStaking *LargeStakingFilterer) FilterUserRewardClaimed(opts *bind.FilterOpts) (*LargeStakingUserRewardClaimedIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "UserRewardClaimed")
	if err != nil {
		return nil, err
	}
	return &LargeStakingUserRewardClaimedIterator{contract: _LargeStaking.contract, event: "UserRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchUserRewardClaimed is a free log subscription operation binding the contract event 0x10b9969d2c05a458029b40aed2ee60918d51f8feb1c2285d6cb1dde5034a12aa.
//
// Solidity: event UserRewardClaimed(uint256 _stakingId, address _beneficiary, uint256 _rewards)
func (_LargeStaking *LargeStakingFilterer) WatchUserRewardClaimed(opts *bind.WatchOpts, sink chan<- *LargeStakingUserRewardClaimed) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "UserRewardClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingUserRewardClaimed)
				if err := _LargeStaking.contract.UnpackLog(event, "UserRewardClaimed", log); err != nil {
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

// ParseUserRewardClaimed is a log parse operation binding the contract event 0x10b9969d2c05a458029b40aed2ee60918d51f8feb1c2285d6cb1dde5034a12aa.
//
// Solidity: event UserRewardClaimed(uint256 _stakingId, address _beneficiary, uint256 _rewards)
func (_LargeStaking *LargeStakingFilterer) ParseUserRewardClaimed(log types.Log) (*LargeStakingUserRewardClaimed, error) {
	event := new(LargeStakingUserRewardClaimed)
	if err := _LargeStaking.contract.UnpackLog(event, "UserRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingValidatorExitReportIterator is returned from FilterValidatorExitReport and is used to iterate over the raw logs and unpacked data for ValidatorExitReport events raised by the LargeStaking contract.
type LargeStakingValidatorExitReportIterator struct {
	Event *LargeStakingValidatorExitReport // Event containing the contract specifics and raw log

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
func (it *LargeStakingValidatorExitReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingValidatorExitReport)
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
		it.Event = new(LargeStakingValidatorExitReport)
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
func (it *LargeStakingValidatorExitReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingValidatorExitReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingValidatorExitReport represents a ValidatorExitReport event raised by the LargeStaking contract.
type LargeStakingValidatorExitReport struct {
	OperatorId *big.Int
	Pubkey     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorExitReport is a free log retrieval operation binding the contract event 0xce36d5425f5ded39d8f3fdf6993fa1813765eda9036a01723ca907fc617d7b86.
//
// Solidity: event ValidatorExitReport(uint256 _operatorId, bytes _pubkey)
func (_LargeStaking *LargeStakingFilterer) FilterValidatorExitReport(opts *bind.FilterOpts) (*LargeStakingValidatorExitReportIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "ValidatorExitReport")
	if err != nil {
		return nil, err
	}
	return &LargeStakingValidatorExitReportIterator{contract: _LargeStaking.contract, event: "ValidatorExitReport", logs: logs, sub: sub}, nil
}

// WatchValidatorExitReport is a free log subscription operation binding the contract event 0xce36d5425f5ded39d8f3fdf6993fa1813765eda9036a01723ca907fc617d7b86.
//
// Solidity: event ValidatorExitReport(uint256 _operatorId, bytes _pubkey)
func (_LargeStaking *LargeStakingFilterer) WatchValidatorExitReport(opts *bind.WatchOpts, sink chan<- *LargeStakingValidatorExitReport) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "ValidatorExitReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingValidatorExitReport)
				if err := _LargeStaking.contract.UnpackLog(event, "ValidatorExitReport", log); err != nil {
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

// ParseValidatorExitReport is a log parse operation binding the contract event 0xce36d5425f5ded39d8f3fdf6993fa1813765eda9036a01723ca907fc617d7b86.
//
// Solidity: event ValidatorExitReport(uint256 _operatorId, bytes _pubkey)
func (_LargeStaking *LargeStakingFilterer) ParseValidatorExitReport(log types.Log) (*LargeStakingValidatorExitReport, error) {
	event := new(LargeStakingValidatorExitReport)
	if err := _LargeStaking.contract.UnpackLog(event, "ValidatorExitReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LargeStakingValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the LargeStaking contract.
type LargeStakingValidatorRegisteredIterator struct {
	Event *LargeStakingValidatorRegistered // Event containing the contract specifics and raw log

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
func (it *LargeStakingValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LargeStakingValidatorRegistered)
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
		it.Event = new(LargeStakingValidatorRegistered)
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
func (it *LargeStakingValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LargeStakingValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LargeStakingValidatorRegistered represents a ValidatorRegistered event raised by the LargeStaking contract.
type LargeStakingValidatorRegistered struct {
	OperatorId *big.Int
	StakeingId *big.Int
	PubKey     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0xd342ed204414e87cc9a5ba37eea0194e34bdc8a046bea5fdc36a00462499dccf.
//
// Solidity: event ValidatorRegistered(uint256 _operatorId, uint256 _stakeingId, bytes _pubKey)
func (_LargeStaking *LargeStakingFilterer) FilterValidatorRegistered(opts *bind.FilterOpts) (*LargeStakingValidatorRegisteredIterator, error) {

	logs, sub, err := _LargeStaking.contract.FilterLogs(opts, "ValidatorRegistered")
	if err != nil {
		return nil, err
	}
	return &LargeStakingValidatorRegisteredIterator{contract: _LargeStaking.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0xd342ed204414e87cc9a5ba37eea0194e34bdc8a046bea5fdc36a00462499dccf.
//
// Solidity: event ValidatorRegistered(uint256 _operatorId, uint256 _stakeingId, bytes _pubKey)
func (_LargeStaking *LargeStakingFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *LargeStakingValidatorRegistered) (event.Subscription, error) {

	logs, sub, err := _LargeStaking.contract.WatchLogs(opts, "ValidatorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LargeStakingValidatorRegistered)
				if err := _LargeStaking.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
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

// ParseValidatorRegistered is a log parse operation binding the contract event 0xd342ed204414e87cc9a5ba37eea0194e34bdc8a046bea5fdc36a00462499dccf.
//
// Solidity: event ValidatorRegistered(uint256 _operatorId, uint256 _stakeingId, bytes _pubKey)
func (_LargeStaking *LargeStakingFilterer) ParseValidatorRegistered(log types.Log) (*LargeStakingValidatorRegistered, error) {
	event := new(LargeStakingValidatorRegistered)
	if err := _LargeStaking.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

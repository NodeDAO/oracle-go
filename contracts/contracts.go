// desc:
// @author renshiwei
// Date: 2023/4/7 11:38

package contracts

import (
	"fmt"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/config"
	"github.com/NodeDAO/oracle-go/contracts/hashConsensus"
	"github.com/NodeDAO/oracle-go/contracts/largeStakeOracle"
	"github.com/NodeDAO/oracle-go/contracts/largeStaking"
	"github.com/NodeDAO/oracle-go/contracts/liq"
	"github.com/NodeDAO/oracle-go/contracts/operator"
	"github.com/NodeDAO/oracle-go/contracts/operatorSlash"
	"github.com/NodeDAO/oracle-go/contracts/vnft"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/NodeDAO/oracle-go/contracts/withdrawalRequest"
	"github.com/NodeDAO/oracle-go/eth1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strings"
)

//var (
//	MULTICALL_MAINNET_ADDRESS = common.HexToAddress("0xeefba1e63905ef1d7acba5a8513c70307c1ce441")
//	MULTICALL_GOERLI_ADDRESS  = common.HexToAddress("0x77dca2c955b15e9de4dbbcf1246b4b85b651e50e")
//)

type withdrawOracleHelper struct {
	Network  string
	Address  string
	Contract *withdrawOracle.WithdrawOracle
}

type hashConsensusHelper struct {
	Network  string
	Address  string
	Contract *hashConsensus.HashConsensus
}

type vnftHelper struct {
	Network  string
	Address  string
	Contract *vnft.Vnft
}

type liqHelper struct {
	Network  string
	Address  string
	Contract *liq.Liq
}

type nodeOperatorHelper struct {
	Network  string
	Address  string
	Contract *operator.Operator
}

type operatorSlashHelper struct {
	Network  string
	Address  string
	Contract *operatorSlash.OperatorSlash
}

type withdrawalRequestHelper struct {
	Network  string
	Address  string
	Contract *withdrawalRequest.WithdrawalRequest
}

type largeStakingHelper struct {
	Network  string
	Address  string
	Contract *largeStaking.LargeStaking
}

type largeStakeOracleHelper struct {
	Network  string
	Address  string
	Contract *largeStakeOracle.LargeStakeOracle
}

var (
	WithdrawOracleContract    *withdrawOracleHelper
	VnftContract              *vnftHelper
	LiqContract               *liqHelper
	OperatorContract          *nodeOperatorHelper
	OperatorSlashContract     *operatorSlashHelper
	WithdrawalRequestContract *withdrawalRequestHelper
	LargeStakingContract      *largeStakingHelper
	LargeStakeOracleContract  *largeStakeOracleHelper
)

const (
	MAINNET = "mainnet"
	GOERLI  = "goerli"
)

const (
	LIQ_ADDRESS_MAINNET = "0x8103151E2377e78C04a3d2564e20542680ed3096"
	LIQ_ADDRESS_GOERLI  = "0x949AC43bb71F8710B0F1193880b338f0323DeB1a"

	VNFT_ADDRESS_MAINNET = "0x58553F5c5a6AEE89EaBFd42c231A18aB0872700d"
	VNFT_ADDRESS_GOERLI  = "0x3CB42bb75Cf1BcC077010ac1E3d3Be22D13326FA"

	NODE_OPERATOR_ADDRESS_MAINNET = "0x8742178Ac172eC7235E54808d5F327C30A51c492"
	NODE_OPERATOR_ADDRESS_GOERLI  = "0x20C43025E44984375c4dC882bFF2016C6E601f0A"

	WITHDRAW_ORACLE_ADDRESS_MAINNET = "0x2B74f97aDC698b571C2F046673Fd5Cd028673c41"
	WITHDRAW_ORACLE_ADDRESS_GOERLI  = "0x1E726f6111B58e74CCD63d5b659191A49366CaD9"

	CL_VAULT_ADDRESS_MAINNET = "0x4b8Dc35b44296D8D6DCc7aFEBBbe283c997E80Ae"
	CL_VAULT_ADDRESS_GOERLi  = "0x138d5D3C2d7d68bFC653726c8a5E8bA301452202"

	OPERATOR_SLASH_ADDRESS_MAINNET = "0x82c87cC83c9fA09DAdBEBFB8f8b9152Ee6104B5d"
	OPERATOR_SLASH_ADDRESS_GOERLi  = "0x69b11EF441EEb3A7cb2A3d82bC31F90596A7C48d"

	WITHDRAWAL_REQUEST_ADDRESS_MAINNET = "0xE81fC969D14Cad8537ebAFa2a1c478F29d7840FC"
	WITHDRAWAL_REQUEST_ADDRESS_GOERLi  = "0x006e69F509E31c91263C03a744B47c3b03eAC391"

	LARGE_STAKING_ADDRESS_MAINNET = ""
	LARGE_STAKING_ADDRESS_GOERLI  = "0xB71D8903Ae22df40DdDb189AfBcE5e99B23b7077"

	LARGE_STAKE_ORACLE_ADDRESS_MAINNET = ""
	LARGE_STAKE_ORACLE_ADDRESS_GOERLI  = "0xB8E0EE431d78273d7BAefEB0Fb64897626b0B8FA"
)

var network string

func InitContracts() {
	var err error

	if config.Config.Eth.Network == "" {
		network = GOERLI
		logger.Warn("config's network is empty. Using default network is goerli")
	} else {
		network = config.Config.Eth.Network
	}

	WithdrawOracleContract, err = newWithdrawOracle()
	VnftContract, err = newVnft()
	LiqContract, err = newLiq()
	OperatorContract, err = newNodeOperator()
	OperatorSlashContract, err = newOperatorSlash()
	WithdrawalRequestContract, err = newWithdrawalRequest()
	LargeStakingContract, err = newLargeStaking()
	LargeStakeOracleContract, err = newLargeStakeOracle()

	if err != nil {
		panic(fmt.Sprintf("New contract error: %+v", err))
	}

	logger.Info("Loading Contract finsh.", zap.String("Network", network))
}

func GetClVaultAddress() string {
	if strings.ToLower(network) == MAINNET {
		return CL_VAULT_ADDRESS_MAINNET
	} else if strings.ToLower(network) == GOERLI {
		return CL_VAULT_ADDRESS_GOERLi
	}

	return ""
}

func newWithdrawOracle() (*withdrawOracleHelper, error) {
	e := &withdrawOracleHelper{
		Network: network,
	}
	if strings.ToLower(network) == MAINNET {
		e.Address = WITHDRAW_ORACLE_ADDRESS_MAINNET
	} else if strings.ToLower(network) == GOERLI {
		e.Address = WITHDRAW_ORACLE_ADDRESS_GOERLI
	}

	if e.Address == "" {
		return nil, errors.New("withdrawOracleHelper contract address is empty.")
	}

	var err error
	e.Contract, err = withdrawOracle.NewWithdrawOracle(common.HexToAddress(e.Address), eth1.ElClient.Client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to new withdraw Oracle.")
	}
	return e, nil
}

func newVnft() (*vnftHelper, error) {
	e := &vnftHelper{
		Network: network,
	}
	if strings.ToLower(network) == MAINNET {
		e.Address = VNFT_ADDRESS_MAINNET
	} else if strings.ToLower(network) == GOERLI {
		e.Address = VNFT_ADDRESS_GOERLI
	}

	if e.Address == "" {
		return nil, errors.New("vnftHelper contract address is empty.")
	}

	var err error
	e.Contract, err = vnft.NewVnft(common.HexToAddress(e.Address), eth1.ElClient.Client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to new VNFT.")
	}
	return e, nil
}

func newLiq() (*liqHelper, error) {
	e := &liqHelper{
		Network: network,
	}
	if strings.ToLower(network) == MAINNET {
		e.Address = LIQ_ADDRESS_MAINNET
	} else if strings.ToLower(network) == GOERLI {
		e.Address = LIQ_ADDRESS_GOERLI
	}

	if e.Address == "" {
		return nil, errors.New("liqHelper contract address is empty.")
	}

	var err error
	e.Contract, err = liq.NewLiq(common.HexToAddress(e.Address), eth1.ElClient.Client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to new liqHelper.")
	}
	return e, nil
}

func newNodeOperator() (*nodeOperatorHelper, error) {
	e := &nodeOperatorHelper{
		Network: network,
	}
	if strings.ToLower(network) == MAINNET {
		e.Address = NODE_OPERATOR_ADDRESS_MAINNET
	} else if strings.ToLower(network) == GOERLI {
		e.Address = NODE_OPERATOR_ADDRESS_GOERLI
	}

	if e.Address == "" {
		return nil, errors.New("nodeOperatorHelper contract address is empty.")
	}

	var err error
	e.Contract, err = operator.NewOperator(common.HexToAddress(e.Address), eth1.ElClient.Client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to new withdraw Oracle.")
	}
	return e, nil
}

func newOperatorSlash() (*operatorSlashHelper, error) {
	e := &operatorSlashHelper{
		Network: network,
	}
	if strings.ToLower(network) == MAINNET {
		e.Address = OPERATOR_SLASH_ADDRESS_MAINNET
	} else if strings.ToLower(network) == GOERLI {
		e.Address = OPERATOR_SLASH_ADDRESS_GOERLi
	}

	var err error
	e.Contract, err = operatorSlash.NewOperatorSlash(common.HexToAddress(e.Address), eth1.ElClient.Client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to new operatorSlashHelper.")
	}
	return e, nil
}

func newWithdrawalRequest() (*withdrawalRequestHelper, error) {
	e := &withdrawalRequestHelper{
		Network: network,
	}
	if strings.ToLower(network) == MAINNET {
		e.Address = WITHDRAWAL_REQUEST_ADDRESS_MAINNET
	} else if strings.ToLower(network) == GOERLI {
		e.Address = WITHDRAWAL_REQUEST_ADDRESS_GOERLi
	}

	var err error
	e.Contract, err = withdrawalRequest.NewWithdrawalRequest(common.HexToAddress(e.Address), eth1.ElClient.Client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to new withdrawal Request.")
	}
	return e, nil
}

func newLargeStaking() (*largeStakingHelper, error) {
	e := &largeStakingHelper{
		Network: network,
	}
	if strings.ToLower(network) == MAINNET {
		e.Address = LARGE_STAKING_ADDRESS_MAINNET
	} else if strings.ToLower(network) == GOERLI {
		e.Address = LARGE_STAKING_ADDRESS_GOERLI
	}

	var err error
	e.Contract, err = largeStaking.NewLargeStaking(common.HexToAddress(e.Address), eth1.ElClient.Client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to new largeStaking.")
	}
	return e, nil
}

func newLargeStakeOracle() (*largeStakeOracleHelper, error) {
	e := &largeStakeOracleHelper{
		Network: network,
	}
	if strings.ToLower(network) == MAINNET {
		e.Address = LARGE_STAKE_ORACLE_ADDRESS_MAINNET
	} else if strings.ToLower(network) == GOERLI {
		e.Address = LARGE_STAKE_ORACLE_ADDRESS_GOERLI
	}

	var err error
	e.Contract, err = largeStakeOracle.NewLargeStakeOracle(common.HexToAddress(e.Address), eth1.ElClient.Client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to new LargeStakeOracle.")
	}
	return e, nil
}

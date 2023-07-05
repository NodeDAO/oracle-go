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
	"github.com/NodeDAO/oracle-go/contracts/vnft"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
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
	WithdrawOracleContract   *withdrawOracleHelper
	VnftContract             *vnftHelper
	LiqContract              *liqHelper
	OperatorContract         *nodeOperatorHelper
	LargeStakingContract     *largeStakingHelper
	LargeStakeOracleContract *largeStakeOracleHelper
	HashConsensusContract    *hashConsensusHelper
)

const (
	MAINNET = "mainnet"
	GOERLI  = "goerli"
)

const (
	HASH_CONSENSUS_ADDRESS_MAINNET = ""
	HASH_CONSENSUS_ADDRESS_GOERLI  = "0xBF7b3b741052D33ca0f522A0D70589e350d38bb7"

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
	LargeStakingContract, err = newLargeStaking()
	LargeStakeOracleContract, err = newLargeStakeOracle()
	HashConsensusContract, err = newHashConsensus()

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

func newHashConsensus() (*hashConsensusHelper, error) {
	e := &hashConsensusHelper{
		Network: network,
	}
	if strings.ToLower(network) == MAINNET {
		e.Address = HASH_CONSENSUS_ADDRESS_MAINNET
	} else if strings.ToLower(network) == GOERLI {
		e.Address = HASH_CONSENSUS_ADDRESS_GOERLI
	}

	var err error
	e.Contract, err = hashConsensus.NewHashConsensus(common.HexToAddress(e.Address), eth1.ElClient.Client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to new HashConsensus.")
	}
	return e, nil
}

// desc:
// @author renshiwei
// Date: 2023/4/7 11:38

package contracts

import (
	"fmt"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/config"
	"github.com/NodeDAO/oracle-go/contracts/hashConsensus"
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

type WithdrawOracle struct {
	Network  string
	Address  string
	Contract *withdrawOracle.WithdrawOracle
}

type HashConsensus struct {
	Network  string
	Address  string
	Contract *hashConsensus.HashConsensus
}

type Vnft struct {
	Network  string
	Address  string
	Contract *vnft.Vnft
}

type Liq struct {
	Network  string
	Address  string
	Contract *liq.Liq
}

type NodeOperator struct {
	Network  string
	Address  string
	Contract *operator.Operator
}

var (
	WithdrawOracleContract *WithdrawOracle
	VnftContract           *Vnft
	LiqContract            *Liq
	OperatorContract       *NodeOperator
)

const (
	MAINNET = "mainnet"
	GOERLI  = "goerli"
)

const (
	LIQ_ADDRESS_MAINNET = ""
	LIQ_ADDRESS_GOERLI  = ""

	VNFT_ADDRESS_MAINNET = ""
	VNFT_ADDRESS_GOERLI  = ""

	NODE_OPERATOR_ADDRESS_MAINNET = ""
	NODE_OPERATOR_ADDRESS_GOERLI  = ""

	WITHDRAW_ORACLE_ADDRESS_MAINNET = ""
	WITHDRAW_ORACLE_ADDRESS_GOERLI  = ""

	CL_VAULT_ADDRESS_MAINNET = ""
	CL_VAULT_ADDRESS_GOERLi  = ""
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

	WithdrawOracleContract, err = NewWithdrawOracle()
	VnftContract, err = NewVnft()
	LiqContract, err = NewLiq()
	OperatorContract, err = NewNodeOperator()
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

func NewWithdrawOracle() (*WithdrawOracle, error) {
	e := &WithdrawOracle{
		Network: network,
	}
	if strings.ToLower(network) == MAINNET {
		e.Address = WITHDRAW_ORACLE_ADDRESS_MAINNET
	} else if strings.ToLower(network) == GOERLI {
		e.Address = WITHDRAW_ORACLE_ADDRESS_GOERLI
	}

	if e.Address == "" {
		return nil, errors.New("WithdrawOracle contract address is empty.")
	}

	var err error
	fmt.Println(eth1.ElClient.Client)
	e.Contract, err = withdrawOracle.NewWithdrawOracle(common.HexToAddress(e.Address), eth1.ElClient.Client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to new withdraw Oracle.")
	}
	return e, nil
}

func NewVnft() (*Vnft, error) {
	e := &Vnft{
		Network: network,
	}
	if strings.ToLower(network) == MAINNET {
		e.Address = VNFT_ADDRESS_MAINNET
	} else if strings.ToLower(network) == GOERLI {
		e.Address = VNFT_ADDRESS_GOERLI
	}

	if e.Address == "" {
		return nil, errors.New("Vnft contract address is empty.")
	}

	var err error
	e.Contract, err = vnft.NewVnft(common.HexToAddress(e.Address), eth1.ElClient.Client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to new VNFT.")
	}
	return e, nil
}

func NewLiq() (*Liq, error) {
	e := &Liq{
		Network: network,
	}
	if strings.ToLower(network) == MAINNET {
		e.Address = LIQ_ADDRESS_MAINNET
	} else if strings.ToLower(network) == GOERLI {
		e.Address = LIQ_ADDRESS_GOERLI
	}

	if e.Address == "" {
		return nil, errors.New("Liq contract address is empty.")
	}

	var err error
	e.Contract, err = liq.NewLiq(common.HexToAddress(e.Address), eth1.ElClient.Client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to new Liq.")
	}
	return e, nil
}

func NewNodeOperator() (*NodeOperator, error) {
	e := &NodeOperator{
		Network: network,
	}
	if strings.ToLower(network) == MAINNET {
		e.Address = NODE_OPERATOR_ADDRESS_MAINNET
	} else if strings.ToLower(network) == GOERLI {
		e.Address = NODE_OPERATOR_ADDRESS_GOERLI
	}

	if e.Address == "" {
		return nil, errors.New("NodeOperator contract address is empty.")
	}

	var err error
	e.Contract, err = operator.NewOperator(common.HexToAddress(e.Address), eth1.ElClient.Client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to new withdraw Oracle.")
	}
	return e, nil
}

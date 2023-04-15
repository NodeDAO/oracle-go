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

var (
	WithdrawOracleContract    *withdrawOracleHelper
	VnftContract              *vnftHelper
	LiqContract               *liqHelper
	OperatorContract          *nodeOperatorHelper
	OperatorSlashContract     *operatorSlashHelper
	WithdrawalRequestContract *withdrawalRequestHelper
)

const (
	MAINNET = "mainnet"
	GOERLI  = "goerli"
)

const (
	LIQ_ADDRESS_MAINNET = ""
	LIQ_ADDRESS_GOERLI  = "0xa8256fD3A31648D49D0f3551e6e45Db6f5f91d53"

	VNFT_ADDRESS_MAINNET = ""
	VNFT_ADDRESS_GOERLI  = "0xe3CE494D51Cb9806187b5Deca1B4B06c97e52EFc"

	NODE_OPERATOR_ADDRESS_MAINNET = ""
	NODE_OPERATOR_ADDRESS_GOERLI  = "0xD9d87abAd8651e1E69799416AEc54fCCdd1dAAcE"

	WITHDRAW_ORACLE_ADDRESS_MAINNET = ""
	WITHDRAW_ORACLE_ADDRESS_GOERLI  = "0x28fbAe8c1A4c04209eb4452907F85560055bC675"

	CL_VAULT_ADDRESS_MAINNET = ""
	CL_VAULT_ADDRESS_GOERLi  = "0x22e172cb3b7a333d73f321462EEBcadd3f0775a6"

	OPERATOR_SLASH_ADDRESS_MAINNET = ""
	OPERATOR_SLASH_ADDRESS_GOERLi  = "0x0f14e0381bBBc2cF3602262Dbd175f2A8fD9A145"

	WITHDRAWAL_REQUEST_ADDRESS_MAINNET = ""
	WITHDRAWAL_REQUEST_ADDRESS_GOERLi  = "0x60bBd3541Dd43c8bf0aCaAdDc06d12d8486822aB"
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
	OperatorSlashContract, err = NewOperatorSlash()
	WithdrawalRequestContract, err = NewWithdrawalRequest()
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

func NewWithdrawOracle() (*withdrawOracleHelper, error) {
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

func NewVnft() (*vnftHelper, error) {
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

func NewLiq() (*liqHelper, error) {
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

func NewNodeOperator() (*nodeOperatorHelper, error) {
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

func NewOperatorSlash() (*operatorSlashHelper, error) {
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

func NewWithdrawalRequest() (*withdrawalRequestHelper, error) {
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

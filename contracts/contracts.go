// desc:
// @author renshiwei
// Date: 2023/4/7 11:38

package contracts

import (
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/config"
	"github.com/NodeDAO/oracle-go/contracts/hashConsensus"
	"github.com/NodeDAO/oracle-go/contracts/liq"
	"github.com/NodeDAO/oracle-go/contracts/vnft"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/NodeDAO/oracle-go/eth1"
	"github.com/ethereum/go-ethereum/common"
)

var (
	MULTICALL_MAINNET_ADDRESS = common.HexToAddress("0xeefba1e63905ef1d7acba5a8513c70307c1ce441")
	MULTICALL_GOERLI_ADDRESS  = common.HexToAddress("0x77dca2c955b15e9de4dbbcf1246b4b85b651e50e")

	LIQ_ADDRESS             = common.HexToAddress(config.Config.Contracts.Liq)
	VNFT_ADDRESS            = common.HexToAddress(config.Config.Contracts.Vnft)
	HASH_CONSENSUS_ADDRESS  = common.HexToAddress(config.Config.Contracts.HashConsensus)
	WITHDRAW_ORACLE_ADDRESS = common.HexToAddress(config.Config.Contracts.WithdrawOracle)
)

var (
	WithdrawOracleContract *withdrawOracle.WithdrawOracle
	HashConsensusContract  *hashConsensus.HashConsensus
	VnftContract           *vnft.Vnft
	LiqContract            *liq.Liq
)

func init() {
	var err error
	WithdrawOracleContract, err = withdrawOracle.NewWithdrawOracle(WITHDRAW_ORACLE_ADDRESS, eth1.ElClient.Client)
	HashConsensusContract, err = hashConsensus.NewHashConsensus(HASH_CONSENSUS_ADDRESS, eth1.ElClient.Client)
	VnftContract, err = vnft.NewVnft(VNFT_ADDRESS, eth1.ElClient.Client)
	LiqContract, err = liq.NewLiq(LIQ_ADDRESS, eth1.ElClient.Client)
	if err != nil {
		logger.Errorf("New contract error: %+v", err)
		panic(err)
	}
}

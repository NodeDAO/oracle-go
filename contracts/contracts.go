// desc:
// @author renshiwei
// Date: 2023/4/7 11:38

package contracts

import (
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/config/global"
	"github.com/NodeDAO/oracle-go/contracts/hashConsensus"
	"github.com/NodeDAO/oracle-go/contracts/liq"
	"github.com/NodeDAO/oracle-go/contracts/vnft"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/ethereum/go-ethereum/common"
)

var (
	LIQ_ADDRESS             = common.HexToAddress(global.Config.Contracts.Liq)
	VNFT_ADDRESS            = common.HexToAddress(global.Config.Contracts.Vnft)
	HASH_CONSENSUS_ADDRESS  = common.HexToAddress(global.Config.Contracts.HashConsensus)
	WITHDRAW_ORACLE_ADDRESS = common.HexToAddress(global.Config.Contracts.WithdrawOracle)
)

var (
	WithdrawOracleContract *withdrawOracle.WithdrawOracle
	HashConsensusContract  *hashConsensus.HashConsensus
	VnftContract           *vnft.Vnft
	LiqContract            *liq.Liq
)

func init() {
	var err error
	WithdrawOracleContract, err = withdrawOracle.NewWithdrawOracle(WITHDRAW_ORACLE_ADDRESS, global.ElClient.Client)
	HashConsensusContract, err = hashConsensus.NewHashConsensus(HASH_CONSENSUS_ADDRESS, global.ElClient.Client)
	VnftContract, err = vnft.NewVnft(VNFT_ADDRESS, global.ElClient.Client)
	LiqContract, err = liq.NewLiq(LIQ_ADDRESS, global.ElClient.Client)
	if err != nil {
		logger.Errorf("New contract error: %+v", err)
		panic(err)
	}
}

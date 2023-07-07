// description:
// @author renshiwei
// Date: 2023/6/29

package largestake

import (
	"github.com/NodeDAO/oracle-go/contracts/largeStakeOracle"
	consensusApi "github.com/attestantio/go-eth2-client/api/v1"
	"math/big"
)

type LargeStakeHelper struct {
	// params
	refSlot          *big.Int
	consensusVersion *big.Int
	refBlockNumber   *big.Int
}

type LargeStakeValidator struct {
	Validator *consensusApi.Validator
	IsExited  bool
	StakingId *big.Int
}

type LargeStakeReportRes struct {
	IsNeedReport   bool
	ReportData     *largeStakeOracle.LargeStakeOracleReportData
	ReportDataHash [32]byte
}

const (
	DAY_BLOCK_NUMBER     = 60 * 60 * 24 / 12
	TWO_DAY_BLOCK_NUMBER = DAY_BLOCK_NUMBER * 2
)

package withdraw

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NodeDAO/oracle-go/app"
	"github.com/NodeDAO/oracle-go/common/errs"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/config"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/NodeDAO/oracle-go/eth1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"math/big"
	"testing"
)

func TestProcessReport(t *testing.T) {
	var err error
	ctx := context.Background()
	//config.InitConfig("../../conf/config-mainnet.dev.yaml")
	config.InitConfig("../../conf/config-goerli.dev.yaml")
	logger.InitLog("debug", "console")
	app.InitServer(ctx)
	w := new(WithdrawHelper)

	err = w.ProcessReport(ctx)

	if sleepErr, ok := errors.Cause(err).(*errs.SleepError); ok {
		logger.Debug("oracle sleep",
			zap.String("msg", sleepErr.Msg),
			zap.String("sleep time", sleepErr.Sleep.String()),
		)
		err = nil
	}

	require.NoError(t, err)
}

func TestProcessReportLoop(t *testing.T) {
	var err error
	ctx := context.Background()
	config.InitConfig("../../conf/config-mainnet.dev.yaml")
	//config.InitConfig("../../conf/config-goerli.dev.yaml")
	logger.InitLog("debug", "console")
	app.InitServer(ctx)
	w := new(WithdrawHelper)

	for i := 0; i < 50; i++ {
		err = w.ProcessReport(ctx)
	}

	if sleepErr, ok := errors.Cause(err).(*errs.SleepError); ok {
		logger.Debug("oracle sleep",
			zap.String("msg", sleepErr.Msg),
			zap.String("sleep time", sleepErr.Sleep.String()),
		)
		err = nil
	}

	require.NoError(t, err)
}

func TestOperatorReward_SettleAmount(t *testing.T) {
	reportDataStr := `
		{"ConsensusVersion":1,"RefSlot":5442463,"ClBalance":607887079648000000000,"ClVaultBalance":65606791189000000000,"ClSettleAmount":65606791189000000000,"ReportExitedCount":2,"WithdrawInfos":[{"OperatorId":1,"ClReward":20859387634551724140,"ClCapital":32000000000000000000},{"OperatorId":2,"ClReward":10429693817275862068,"ClCapital":0},{"OperatorId":3,"ClReward":0,"ClCapital":0},{"OperatorId":4,"ClReward":1158854868586206896,"ClCapital":0},{"OperatorId":6,"ClReward":1158854868586206896,"ClCapital":0},{"OperatorId":7,"ClReward":0,"ClCapital":0},{"OperatorId":10,"ClReward":0,"ClCapital":0},{"OperatorId":11,"ClReward":0,"ClCapital":0},{"OperatorId":12,"ClReward":0,"ClCapital":0}],"ExitValidatorInfos":[{"ExitTokenId":17,"ExitBlockNumber":8838718,"SlashAmount":0},{"ExitTokenId":18,"ExitBlockNumber":8838482,"SlashAmount":0}],"DelayedExitTokenIds":[],"LargeExitDelayedRequestIds":[]}
	`

	var reportData *withdrawOracle.WithdrawOracleReportData
	err := json.Unmarshal([]byte(reportDataStr), &reportData)
	require.NoError(t, err)

	sumSettle := big.NewInt(0)
	for _, info := range reportData.WithdrawInfos {
		opSettle := new(big.Int).Add(info.ClReward, info.ClCapital)
		sumSettle = new(big.Int).Add(sumSettle, opSettle)
	}

	//if reportData.ClSettleAmount.Cmp(sumSettle) != 0 {
	//	accuracy := new(big.Int).Sub(reportData.ClSettleAmount, sumSettle)
	//	reportData.WithdrawInfos[0].ClReward = new(big.Int).Add(reportData.WithdrawInfos[0].ClReward, accuracy)
	//
	//	sumSettle = new(big.Int).Add(sumSettle, accuracy)
	//}

	require.Equal(t, reportData.ClSettleAmount, sumSettle)
}

//func TestExitValidator(t *testing.T) {
//	ctx := context.Background()
//	config.InitConfig("../../conf/config-dev.yaml")
//	app.InitServer(ctx)
//	// delayedTokenId
//	delayedTokenId := 17
//
//
//	w := new(WithdrawHelper)
//	err := w.ProcessReport(ctx)
//	require.NoError(t, err)
//}

func TestBigInt(t *testing.T) {
	reportExitedCount, _ := new(big.Int).SetString("0", 0)
	reportExitedCount2 := big.NewInt(1)

	fmt.Println(reportExitedCount)
	fmt.Println(reportExitedCount2)

	x, _ := new(big.Int).SetString("11", 0)
	y, _ := new(big.Int).SetString("3", 0)
	h, k := new(big.Int).DivMod(x, y, big.NewInt(0))
	fmt.Println(h)
	fmt.Println(k)

	h1, _ := decimal.NewFromString("20859387634551724140")
	h2, _ := decimal.NewFromString("10429693817275862068")
	h3, _ := decimal.NewFromString("1158854868586206896")
	h4, _ := decimal.NewFromString("1158854868586206896")
	h5, _ := decimal.NewFromString("32000000000000000000")
	s1 := h1.Add(h2).Add(h3).Add(h4).Add(h5)

	s2, _ := decimal.NewFromString("65606791189000000000")

	require.Equal(t, s2.String(), s1.String())
}

func TestNethEx0501(t *testing.T) {
	ether1, _ := decimal.NewFromString("1000000000000000000")

	clBalance, _ := decimal.NewFromString("3520000000000000000000")
	clVaultBalance, _ := decimal.NewFromString("3297310124000000000")
	operatorPoolBalancesSum, _ := decimal.NewFromString("56684526602667774526")
	totalETH := clBalance.Add(clVaultBalance).Add(operatorPoolBalancesSum)
	fmt.Println("totalETH", totalETH.String())

	nETH1, _ := decimal.NewFromString("1000000000000000000")
	totalNETH, _ := decimal.NewFromString("3565401715328792503385")
	ex := nETH1.Mul(totalETH).Div(totalNETH)
	// 1.004089334824513800
	fmt.Println("1 nETH = ", ex.Div(ether1).String(), " ether")
}

func TestNethEx0503(t *testing.T) {
	ether1, _ := decimal.NewFromString("1000000000000000000")

	clBalance, _ := decimal.NewFromString("3552098096645000000000")
	clVaultBalance, _ := decimal.NewFromString("3653139213000000000")
	operatorPoolBalancesSum, _ := decimal.NewFromString("10674703411883184608")
	totalETH := clBalance.Add(clVaultBalance).Add(operatorPoolBalancesSum)
	fmt.Println("totalETH", totalETH.String(), " wei ", totalETH.Div(ether1).String(), " ether")

	nETH1, _ := decimal.NewFromString("1000000000000000000")
	totalNETH, _ := decimal.NewFromString("3551442991856375267186")
	ex := nETH1.Mul(totalETH).Div(totalNETH)
	// 1.0042188337100904
	fmt.Println("1 nETH = ", ex.Div(ether1).String(), " ether")
}

func TestClBalance(t *testing.T) {
	ether1, _ := decimal.NewFromString("1000000000000000000")

	h1, _ := decimal.NewFromString("3552000000000000000000")
	//h2, _ := decimal.NewFromString("3552272701086000000000")
	h2, _ := decimal.NewFromString("3552346631528000000000")
	s1 := h2.Sub(h1)
	// 0.272701086 ether
	fmt.Println(s1.Div(ether1).String())
}

func TestCheckTotalClBalance(t *testing.T) {
	ether1, _ := decimal.NewFromString("1000000000000000000")

	h1, _ := decimal.NewFromString("128006940284000000000")
	h2, _ := decimal.NewFromString("20892684000000000")
	preTotal := h1.Add(h2)

	h3, _ := decimal.NewFromString("128000503467000000000")
	h4, _ := decimal.NewFromString("1028527387000000000")
	curTotal := h3.Add(h4)

	require.Equal(t, true, preTotal.LessThan(curTotal))

	s1 := h2.Sub(h1)
	// 0.272701086 ether
	fmt.Println(s1.Div(ether1).String())
}

func TestBalance(t *testing.T) {
	var err error
	ctx := context.Background()
	config.InitConfig("../../conf/config-mainnet.dev.yaml")
	//config.InitConfig("../../conf/config-goerli.dev.yaml")
	logger.InitLog("debug", "console")

	eth1.ElClient, err = eth1.NewEthClient(ctx, config.Config.Eth.ElAddr)
	require.NoError(t, err)

	balanceAt, err := eth1.ElClient.Client.BalanceAt(ctx, common.HexToAddress("0x4b8Dc35b44296D8D6DCc7aFEBBbe283c997E80Ae"), big.NewInt(18205397))
	require.NoError(t, err)
	fmt.Println("balanceAt", balanceAt.String())
}

package withdraw

import (
	"context"
	"fmt"
	"github.com/NodeDAO/oracle-go/app"
	"github.com/NodeDAO/oracle-go/config"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestProcessReport(t *testing.T) {
	ctx := context.Background()
	config.InitConfig("../../conf/config-dev.yaml")
	app.InitServer(ctx)
	w := new(WithdrawHelper)
	err := w.ProcessReport(ctx)
	require.NoError(t, err)
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

	h1, _ := decimal.NewFromString("491176250068965517")
	h2, _ := decimal.NewFromString("982352500137931034")
	h3, _ := decimal.NewFromString("54575138896551724")
	h4, _ := decimal.NewFromString("54575138896551724")
	s1 := h1.Add(h2).Add(h3).Add(h4).String()
	fmt.Println(s1)
	s2 := h1.Add(h2).Add(h3).Add(h4).Add(decimal.New(29, 0)).String()
	fmt.Println(s2)
}

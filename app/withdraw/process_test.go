package withdraw

import (
	"context"
	"fmt"
	"github.com/NodeDAO/oracle-go/app"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestProcessReport(t *testing.T) {
	ctx := context.Background()
	app.InitServer(ctx, "../../conf/config-dev.yaml")
	w := new(WithdrawHelper)
	err := w.ProcessReport(ctx)
	require.NoError(t, err)
}

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
}

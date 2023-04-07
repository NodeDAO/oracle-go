// desc:
// @author renshiwei
// Date: 2023/4/6 16:48

package eth1

import (
	"context"
	"fmt"
	"github.com/NodeDAO/oracle-go/config"
	"github.com/NodeDAO/oracle-go/config/global"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBalanceAt(t *testing.T) {
	config.InitConfig("../conf/config-dev.yaml")
	elAddr := global.Config.Eth.ElAddr
	background := context.Background()

	ethClient, err := NewEthClient(background, elAddr)
	require.NoError(t, err)

	// 0x75C0c372da875a4Fc78E8A37f58618a6D18904e8
	blockNumber := decimal.NewFromInt(8782889)
	balance, err := ethClient.BalanceAt(background, "0x892e7c8C5E716e17891ABf9395a0de1f2fc84786", blockNumber)
	require.NoError(t, err)
	fmt.Println("balance:", balance)
}

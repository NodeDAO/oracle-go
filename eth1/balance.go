// desc:
// @author renshiwei
// Date: 2023/4/6 16:28

package eth1

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func (e *EthClient) BalanceAt(ctx context.Context, address string, blockNumber decimal.Decimal) (decimal.Decimal, error) {
	balance, err := e.Client.BalanceAt(ctx, common.HexToAddress(address), blockNumber.BigInt())
	if err != nil {
		return decimal.Zero, errors.Wrapf(err, "Fail to get balance address:%s blockNumber:%s", address, blockNumber.String())
	}
	return decimal.NewFromBigInt(balance, 0), nil
}

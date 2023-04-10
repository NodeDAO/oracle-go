// desc:
// @author renshiwei
// Date: 2023/4/6 16:28

package eth1

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"math/big"
)

func (e *EthClient) BalanceAt(ctx context.Context, address string, blockNumber *big.Int) (*big.Int, error) {
	balance, err := e.Client.BalanceAt(ctx, common.HexToAddress(address), blockNumber)
	if err != nil {
		return decimal.Zero.BigInt(), errors.Wrapf(err, "Fail to get balance address:%s blockNumber:%s", address, blockNumber.String())
	}
	return balance, nil
}

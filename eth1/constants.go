// desc:
// @author renshiwei
// Date: 2023/4/7 19:31

package eth1

import (
	"github.com/ethereum/go-ethereum/params"
	"github.com/shopspring/decimal"
	"math/big"
)

func GWEIToWEI(value *big.Int) *big.Int {
	return new(big.Int).Mul(value, big.NewInt(params.GWei))
}

func ETH32() decimal.Decimal {
	eth32, _ := decimal.NewFromString("32000000000000000000")
	return eth32
}

// desc:
// @author renshiwei
// Date: 2023/4/7 19:31

package eth1

import "github.com/shopspring/decimal"

func ETH32() decimal.Decimal {
	eth32, _ := decimal.NewFromString("32000000000000000000")
	return eth32
}

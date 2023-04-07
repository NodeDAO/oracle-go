// desc:
// @author renshiwei
// Date: 2023/4/7 18:23

package eth1

import (
	"github.com/alethio/web3-go/ethrpc"
	"github.com/alethio/web3-go/ethrpc/provider/httprpc"
	"github.com/alethio/web3-multicall-go/multicall"
)

func NewMulticallClient(ethAddr, multicallAddress string) (*multicall.Multicall, error) {
	p, err := httprpc.New(ethAddr)
	if err != nil {

	}

	eth, err := ethrpc.New(p)
	if err != nil {

	}

	mc, err := multicall.New(eth, multicall.ContractAddress(multicallAddress))
	if err != nil {

	}

	return &mc, nil
}

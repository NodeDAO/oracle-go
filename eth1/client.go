// desc:
// @author renshiwei
// Date: 2023/4/6 15:19

package eth1

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type EthClient struct {
	client *ethclient.Client
}

func NewEthClient(ctx context.Context, rpcHost string) (*EthClient, error) {
	elClient, err := getEthClient(ctx, rpcHost)
	if err != nil {
		return nil, errors.Unwrap(err)
	}
	return &EthClient{
		client: elClient,
	}, nil
}

func getEthClient(ctx context.Context, rpcHost string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpcHost)
	if err != nil {
		return nil, errors.Wrap(err, "Fail to connect to Executive layer client")
	}
	return client, nil
}

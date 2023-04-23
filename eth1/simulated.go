package eth1

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/pkg/errors"
	"math/big"
)

// NewSimulatedClient ETH client of constructing simulation transactions
func NewSimulatedClient(chainID *big.Int, walletBalance *big.Int, gasLimit uint64) (*backends.SimulatedBackend, error) {
	opts, err := KeyTransactOpts(chainID)
	if err != nil {
		return nil, errors.Wrapf(err, "")
	}

	address := opts.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: walletBalance,
		},
	}

	simulatedClient := backends.NewSimulatedBackend(genesisAlloc, gasLimit)

	return simulatedClient, nil
}

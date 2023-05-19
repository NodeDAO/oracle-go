package eth1

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPubkeyFromPrivateKey(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	privateKeyBytes := crypto.FromECDSA(privateKey)
	priv := hexutil.Encode(privateKeyBytes)

	address, err := AddressFromPrivateKey(priv)
	require.NoError(t, err)
	fmt.Println(address)

	address2, err := AddressFromPrivateKey(priv[2:])
	require.NoError(t, err)
	require.Equal(t, address, address2)
}

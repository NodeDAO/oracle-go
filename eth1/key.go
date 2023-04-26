// desc:
// @author renshiwei
// Date: 2023/4/11 17:02

package eth1

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"strings"
)

// AddressFromPrivateKey Push out the address from the private key
func AddressFromPrivateKey(pri string) (common.Address, error) {
	if !strings.HasPrefix(pri, "0x") {
		pri = fmt.Sprintf("0x%s", pri)
	}
	priHex, err := hexutil.Decode(pri)
	privateKey, err := crypto.ToECDSA(priHex)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "Failed to ToECDSA")
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("publicKey is not a *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, nil
}

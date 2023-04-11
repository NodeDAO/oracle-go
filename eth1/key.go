// desc:
// @author renshiwei
// Date: 2023/4/11 17:02

package eth1

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

func PubkeyFromPrivateKey(pri string) (common.Address, error) {
	privateKey, err := crypto.ToECDSA([]byte(pri))
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

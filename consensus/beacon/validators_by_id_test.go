// desc:
// @author renshiwei
// Date: 2023/4/7 16:37

package beacon

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestValidatorsByPubKey(t *testing.T) {
	initClAddr()
	b := &BeaconService{
		BaseUrl: clAddr,
		Timeout: 1 * time.Minute,
	}

	validatorMap, err := b.ValidatorsByPubKey(context.Background(), "head", []string{
		"0x8000091c2ae64ee414a54c1cc1fc67dec663408bc636cb86756e0200e41a75c8f86603f104f02c856983d2783116be13",
		"0x80000e851c0f53c3246ff726d7ff7766661ca5e12a07c45c114d208d54f0f8233d4380b2e9aff759d69795d1df905526",
	})
	require.NoError(t, err)

	fmt.Printf("pubkey1:%+v\n", validatorMap["0x8000091c2ae64ee414a54c1cc1fc67dec663408bc636cb86756e0200e41a75c8f86603f104f02c856983d2783116be13"])
	fmt.Printf("pubkey2:%+v\n", validatorMap["0x80000e851c0f53c3246ff726d7ff7766661ca5e12a07c45c114d208d54f0f8233d4380b2e9aff759d69795d1df905526"])
}

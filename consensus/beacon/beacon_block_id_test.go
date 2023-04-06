// desc:
// @author renshiwei
// Date: 2023/4/6 13:06

package beacon

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetBeaconBlock(t *testing.T) {
	initClAddr()
	b := &BeaconService{
		BaseUrl: clAddr,
	}

	beaconBlock, err := b.BeaconBlock(context.Background(), "head")
	require.NoError(t, err)

	marshal, err := json.Marshal(beaconBlock)
	require.NoError(t, err)

	fmt.Printf("beaconBlock:%s\n", string(marshal))
}

func TestGetExecutionPayload(t *testing.T) {
	initClAddr()
	b := &BeaconService{
		BaseUrl: clAddr,
	}

	executionPayload, err := b.ExecutionPayload(context.Background(), "head")
	require.NoError(t, err)

	fmt.Printf("BlockNumber:%s\n", executionPayload.BlockNumber)
	fmt.Printf("Timestamp:%s\n", executionPayload.Timestamp)
	fmt.Printf("BlockHash:%s\n", executionPayload.BlockHash)
}

func TestGetExecutionBlock(t *testing.T) {
	initClAddr()
	b := &BeaconService{
		BaseUrl: clAddr,
	}

	executionBlock, err := b.ExecutionBlock(context.Background(), "5354527")
	require.NoError(t, err)

	fmt.Printf("BlockNumber:%s\n", executionBlock.BlockNumber)
	fmt.Printf("Timestamp:%s\n", executionBlock.Timestamp)
	fmt.Printf("BlockHash:%s\n", executionBlock.BlockHash)
}

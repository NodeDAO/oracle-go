// desc:
// @author renshiwei
// Date: 2023/4/6 13:10

package beacon

import (
	"context"
	"time"
)

// BeaconBlocksProvider is the interface for providing beacon blocks.
type BeaconBlocksProvider interface {
	// BeaconBlock provides the block header of a given block ID.
	BeaconBlock(ctx context.Context, blockID string) (*BeaconBlock, error)
}

// ExecutionInfoProvider is the interface for get ExecutionPayload.
type ExecutionInfoProvider interface {
	// ExecutionPayload provides the block's Execution info of a given block ID.
	ExecutionPayload(ctx context.Context, blockID string) (*ExecutionPayload, error)
}

type BeaconService struct {
	BaseUrl string
	Timeout time.Duration
}

func New(ctx context.Context, addr string, timeout time.Duration) (*BeaconService, error) {
	return &BeaconService{
		BaseUrl: addr,
		Timeout: timeout,
	}, nil
}

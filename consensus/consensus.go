// desc:
// @author renshiwei
// Date: 2023/4/6 20:02

package consensus

import (
	"context"
	"github.com/NodeDAO/oracle-go/consensus/beacon"
	"github.com/NodeDAO/oracle-go/consensus/chaintime/standard"
	eth2client "github.com/attestantio/go-eth2-client"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"time"
)

type Consensus struct {
	ConsensusClient        *eth2client.Service
	ChainTimeService       *standard.Service
	CustomizeBeaconService *beacon.BeaconService

	Timeout time.Duration
}

var ConsensusClient *Consensus

func New(ctx context.Context, addr string, timeout time.Duration) (*Consensus, error) {
	beaconNode, err := beacon.ConnectToBeaconNode(ctx, addr, timeout, true)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	beaconService, err := beacon.New(ctx, addr, timeout)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to connect to beacon service")
	}

	chainTimeService, err := standard.New(ctx,
		standard.WithGenesisTimeProvider(beaconNode.(eth2client.GenesisTimeProvider)),
		standard.WithSpecProvider(beaconNode.(eth2client.SpecProvider)),
		standard.WithLogLevel(zerolog.Disabled),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create chaintime service")
	}

	return &Consensus{
		ConsensusClient:        &beaconNode,
		ChainTimeService:       chainTimeService,
		CustomizeBeaconService: beaconService,
		Timeout:                timeout,
	}, nil

}

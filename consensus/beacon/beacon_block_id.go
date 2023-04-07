// desc:
// beacon-api: @see https://ethereum.github.io/beacon-APIs/#/Beacon/getBlockV2
// /eth/v2/beacon/blocks/{}
//
// struct @see api/v1/capella/blindedbeaconblock.go
// @author renshiwei
// Date: 2023/4/6 10:51

package beacon

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NodeDAO/oracle-go/utils/httptool"
	//"github.com/attestantio/go-eth2-client/api/v1/capella"
	"github.com/pkg/errors"
)

type BeaconBlock struct {
	Version string `json:"version"`
	Data    struct {
		Message struct {
			Slot          string `json:"slot"`
			ProposerIndex string `json:"proposer_index"`
			ParentRoot    string `json:"parent_root"`
			StateRoot     string `json:"state_root"`
			Body          struct {
				RandaoReveal string `json:"randao_reveal"`
				Eth1Data     struct {
					DepositRoot  string `json:"deposit_root"`
					DepositCount string `json:"deposit_count"`
					BlockHash    string `json:"block_hash"`
				} `json:"eth1_data"`
				Graffiti          string        `json:"graffiti"`
				ProposerSlashings []interface{} `json:"proposer_slashings"`
				AttesterSlashings []interface{} `json:"attester_slashings"`
				Attestations      []struct {
					AggregationBits string `json:"aggregation_bits"`
					Data            struct {
						Slot            string `json:"slot"`
						Index           string `json:"index"`
						BeaconBlockRoot string `json:"beacon_block_root"`
						Source          struct {
							Epoch string `json:"epoch"`
							Root  string `json:"root"`
						} `json:"source"`
						Target struct {
							Epoch string `json:"epoch"`
							Root  string `json:"root"`
						} `json:"target"`
					} `json:"data"`
					Signature string `json:"signature"`
				} `json:"attestations"`
				Deposits       []interface{} `json:"deposits"`
				VoluntaryExits []interface{} `json:"voluntary_exits"`
				SyncAggregate  struct {
				} `json:"sync_aggregate"`
				ExecutionPayload *ExecutionPayload `json:"execution_payload"`
			} `json:"body"`
		} `json:"message"`
		Signature string `json:"signature"`
	} `json:"data"`
	ExecutionOptimistic bool `json:"execution_optimistic"`
	Finalized           bool `json:"finalized"`
}

type ExecutionPayload struct {
	ParentHash    string        `json:"parent_hash"`
	FeeRecipient  string        `json:"fee_recipient"`
	StateRoot     string        `json:"state_root"`
	ReceiptsRoot  string        `json:"receipts_root"`
	LogsBloom     string        `json:"logs_bloom"`
	PrevRandao    string        `json:"prev_randao"`
	BlockNumber   string        `json:"block_number"`
	GasLimit      string        `json:"gas_limit"`
	GasUsed       string        `json:"gas_used"`
	Timestamp     string        `json:"timestamp"`
	ExtraData     string        `json:"extra_data"`
	BaseFeePerGas string        `json:"base_fee_per_gas"`
	BlockHash     string        `json:"block_hash"`
	Transactions  []interface{} `json:"transactions"`
}

type ExecutionBlock struct {
	ParentHash    string `json:"parent_hash"`
	FeeRecipient  string `json:"fee_recipient"`
	BlockNumber   string `json:"block_number"`
	GasLimit      string `json:"gas_limit"`
	GasUsed       string `json:"gas_used"`
	Timestamp     string `json:"timestamp"`
	BaseFeePerGas string `json:"base_fee_per_gas"`
	BlockHash     string `json:"block_hash"`
}

func (b *BeaconService) BeaconBlock(ctx context.Context, blockID string) (*BeaconBlock, error) {
	httpTool, err := httptool.New(ctx, b.Timeout)
	if err != nil {
		return nil, errors.Unwrap(err)
	}

	respBodyReader, err := httpTool.GetRequest(ctx, fmt.Sprintf("%s/eth/v2/beacon/blocks/%s", b.BaseUrl, blockID))
	if err != nil {
		return nil, errors.Wrap(err, "failed to request beacon block header")
	}
	if respBodyReader == nil {
		return nil, nil
	}

	var resp *BeaconBlock
	if err := json.NewDecoder(respBodyReader).Decode(&resp); err != nil {
		return nil, errors.Wrap(err, "failed to parse beacon block")
	}

	return resp, nil
}

func (b *BeaconService) ExecutionPayload(ctx context.Context, blockID string) (*ExecutionPayload, error) {
	beaconBlock, err := b.BeaconBlock(ctx, blockID)
	if err != nil {
		return nil, errors.Unwrap(err)
	}

	return beaconBlock.Data.Message.Body.ExecutionPayload, nil
}

func (b *BeaconService) ExecutionBlock(ctx context.Context, blockID string) (*ExecutionBlock, error) {
	executionPayload, err := b.ExecutionPayload(ctx, blockID)
	if err != nil {
		return nil, errors.Unwrap(err)
	}

	return &ExecutionBlock{
		ParentHash:    executionPayload.ParentHash,
		FeeRecipient:  executionPayload.FeeRecipient,
		BlockNumber:   executionPayload.BlockNumber,
		GasLimit:      executionPayload.GasLimit,
		GasUsed:       executionPayload.GasUsed,
		Timestamp:     executionPayload.Timestamp,
		BaseFeePerGas: executionPayload.BaseFeePerGas,
		BlockHash:     executionPayload.BlockHash,
	}, nil
}

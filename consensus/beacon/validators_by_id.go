// desc:
// @author renshiwei
// Date: 2023/4/7 16:23

package beacon

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NodeDAO/oracle-go/utils/httptool"
	consensusApi "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/pkg/errors"
	"strings"
)

type validatorsByPubKeyJSON struct {
	Data []*consensusApi.Validator `json:"data"`
}

func (b *BeaconService) ValidatorsByPubKey(ctx context.Context, stateID string, validatorPubKeys []string) (map[string]*consensusApi.Validator, error) {
	httpTool, err := httptool.New(ctx, b.Timeout)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	url := fmt.Sprintf("%s/eth/v1/beacon/states/%s/validators", b.BaseUrl, stateID)
	if len(validatorPubKeys) != 0 {
		ids := make([]string, len(validatorPubKeys))
		for i := range validatorPubKeys {
			ids[i] = validatorPubKeys[i]
		}
		url = fmt.Sprintf("%s?id=%s", url, strings.Join(ids, ","))
	}
	respBodyReader, err := httpTool.GetRequest(ctx, url)

	if err != nil {
		return nil, errors.Wrap(err, "failed to request beacon block header")
	}
	if respBodyReader == nil {
		return nil, nil
	}

	var validatorsByPubKeyJSON validatorsByPubKeyJSON
	if err := json.NewDecoder(respBodyReader).Decode(&validatorsByPubKeyJSON); err != nil {
		return nil, errors.Wrap(err, "failed to parse validators")
	}
	if validatorsByPubKeyJSON.Data == nil {
		return nil, errors.New("no validators returned")
	}

	res := make(map[string]*consensusApi.Validator)
	for _, validator := range validatorsByPubKeyJSON.Data {
		res[validator.Validator.PublicKey.String()] = validator
	}

	return res, nil
}

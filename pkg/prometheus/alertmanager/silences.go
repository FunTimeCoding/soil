package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
	rawSilence "github.com/prometheus/alertmanager/api/v2/client/silence"
)

func (c *Client) Silences(expired bool) ([]*silence.Silence, error) {
	p := rawSilence.NewGetSilencesParams()
	response, e := c.client.Silence.GetSilences(p)

	if e != nil {
		return nil, e
	}

	var result []*silence.Silence

	for _, s := range silence.NewSlice(response.GetPayload(), c.host) {
		if !expired && s.Expired() {
			continue
		}

		result = append(result, s)
	}

	return silence.Sort(result), nil
}

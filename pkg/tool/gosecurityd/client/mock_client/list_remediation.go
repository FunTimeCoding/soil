package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
)

func (c *Client) ListRemediation(
	_ context.Context,
	p *client.ListRemediationParams,
) (*client.RemediationList, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return nil, c.failure
	}

	if p == nil {
		p = &client.ListRemediationParams{}
	}

	s, e := convert.Selection(p.Show, p.Label, p.Search, nil, 0)

	if e != nil {
		return nil, e
	}

	lowest := constant.CriticalScore

	if p.Score != nil {
		lowest = *p.Score
	}

	result := &client.RemediationList{Items: []client.RemediationItem{}}

	for _, v := range c.remediation {
		if v.TopScore < lowest && !v.Unknown {
			continue
		}

		if narrows(s) {
			kept := []string{}

			for _, h := range v.Hosts {
				if c.chosen(h, s) {
					kept = append(kept, h)
				}
			}

			if len(kept) == 0 {
				continue
			}

			v.Hosts = kept
		}

		result.Items = append(result.Items, v)

		if v.Addressed {
			result.Addressed++
		} else {
			result.Open++
		}
	}

	return result, nil
}

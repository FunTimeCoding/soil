package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
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

	lowest := constant.CriticalScore

	if p != nil && p.Score != nil {
		lowest = *p.Score
	}

	result := &client.RemediationList{Items: []client.RemediationItem{}}

	for _, v := range c.remediation {
		if v.TopScore < lowest && !v.Unknown {
			continue
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

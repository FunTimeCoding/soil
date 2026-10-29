package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
)

func (c *Client) GetAlertsAfter(
	_ context.Context,
	p *client.GetAlertsAfterParams,
) ([]client.AlertsResponse, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	limit := constant.PageDefault

	if p.Limit != nil {
		limit = min(max(*p.Limit, 1), constant.PageMaximum)
	}

	result := make([]client.AlertsResponse, 0, limit)

	for _, a := range c.alerts {
		if a.Identifier > p.Identifier && len(result) < limit {
			result = append(result, a)
		}
	}

	return result, nil
}

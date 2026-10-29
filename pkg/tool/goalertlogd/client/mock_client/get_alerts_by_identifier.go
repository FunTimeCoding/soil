package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"slices"
)

func (c *Client) GetAlertsByIdentifier(
	_ context.Context,
	p *client.GetAlertsByIdentifierParams,
) ([]client.AlertsResponse, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	var wanted []int

	for _, i := range p.Identifier {
		if i > 0 && len(wanted) < constant.PageMaximum {
			wanted = append(wanted, i)
		}
	}

	result := make([]client.AlertsResponse, 0, len(wanted))

	for _, a := range c.alerts {
		if slices.Contains(wanted, a.Identifier) {
			result = append(result, a)
		}
	}

	return result, nil
}

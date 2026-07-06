package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_termination"
)

func (c *Client) TunnelTerminations() ([]*tunnel_termination.Termination, error) {
	result, _, e := c.client.VpnAPI.VpnTunnelTerminationsList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	terminations := make(
		[]*tunnel_termination.Termination,
		0,
		len(result.Results),
	)

	for _, t := range result.Results {
		terminations = append(terminations, tunnel_termination.New(&t))
	}

	return terminations, nil
}

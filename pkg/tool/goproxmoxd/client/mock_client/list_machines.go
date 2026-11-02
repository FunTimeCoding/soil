package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ListMachines(
	_ context.Context,
	p *client.ListMachinesParams,
) ([]client.Machine, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	var result []client.Machine

	for _, m := range c.MachineList {
		if p == nil || p.Node == nil || (m.Node != nil && *m.Node == *p.Node) {
			result = append(result, m)
		}
	}

	return result, nil
}

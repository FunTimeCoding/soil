package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
)

func (c *Client) ListFlags(
	_ context.Context,
	p *client.ListFlagsParams,
) ([]*client.Flag, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return nil, c.failure
	}

	all := p != nil && p.IncludeMitigated != nil && *p.IncludeMitigated
	result := []*client.Flag{}

	for _, f := range c.flags {
		if all || !f.Mitigated {
			result = append(result, f)
		}
	}

	return result, nil
}

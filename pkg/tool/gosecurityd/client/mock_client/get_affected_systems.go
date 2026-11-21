package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/vulnerability/identifier"
)

func (c *Client) GetAffectedSystems(
	_ context.Context,
	name string,
) (*client.AffectedSystems, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return nil, c.failure
	}

	if e := identifier.Check(name); e != nil {
		return nil, e
	}

	if seeded, okay := c.affected[name]; okay {
		return seeded, nil
	}

	return &client.AffectedSystems{
		Identifier: name,
		Systems:    []client.AffectedSystem{},
	}, nil
}

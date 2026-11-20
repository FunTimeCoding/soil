package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"github.com/funtimecoding/soil/pkg/vulnerability/identifier"
)

func (c *Client) GetDebianStatus(
	_ context.Context,
	name string,
	p *client.GetDebianStatusParams,
) (*client.DebianStatus, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return nil, c.failure
	}

	if e := identifier.Check(name); e != nil {
		return nil, e
	}

	seeded, okay := c.statuses[name]

	if !okay {
		return nil, not_found.New(vulnerabilityConstant.Kind, name)
	}

	result := &client.DebianStatus{
		Identifier:      seeded.Identifier,
		TrackerModified: seeded.TrackerModified,
	}

	for _, s := range seeded.Statuses {
		if p == nil || p.Release == nil || s.Release == *p.Release {
			result.Statuses = append(result.Statuses, s)
		}
	}

	return result, nil
}

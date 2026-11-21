package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
)

func (c *Client) GetUpgrade(
	_ context.Context,
	release string,
	source string,
) (*client.UpgradeGroup, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return nil, c.failure
	}

	for _, g := range c.upgrades {
		if g.Release == release && g.Source == source {
			return &g, nil
		}
	}

	return nil, not_found.New(constant.UpgradeKind, join.Space(release, source))
}

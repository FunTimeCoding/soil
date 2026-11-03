package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
)

func (c *Client) groupIndex(name string) (int, error) {
	for i, g := range c.GroupList {
		if g.Name == name {
			return i, nil
		}
	}

	return 0, not_found.New(constant.GroupKind, name)
}

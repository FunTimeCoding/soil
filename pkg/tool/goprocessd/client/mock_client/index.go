package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
)

func (c *Client) index(name string) (int, error) {
	for i, p := range c.ProcessList {
		if p.Name == name {
			return i, nil
		}
	}

	return -1, not_found.New(constant.ProcessKind, name)
}

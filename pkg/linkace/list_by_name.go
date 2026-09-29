package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/linkace/constant"
	"github.com/funtimecoding/soil/pkg/linkace/list"
)

func (c *Client) ListByName(name string) (*list.List, error) {
	all, e := c.Lists()

	if e != nil {
		return nil, e
	}

	for _, l := range all {
		if l.Name == name {
			return l, nil
		}
	}

	return nil, not_found.New(constant.ListSubject, name)
}

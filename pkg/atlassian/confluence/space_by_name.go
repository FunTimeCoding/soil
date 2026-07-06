package confluence

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/space"
)

func (c *Client) SpaceByName(name string) (*space.Space, error) {
	spaces, e := c.Spaces()

	if e != nil {
		return nil, e
	}

	for _, s := range spaces {
		if s.Name == name {
			return s, nil
		}
	}

	return nil, fmt.Errorf(
		"space not found: %s: %w",
		name,
		constant.ErrorNotFound,
	)
}

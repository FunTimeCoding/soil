package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/space"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
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

	return nil, not_found.New("space", name)
}

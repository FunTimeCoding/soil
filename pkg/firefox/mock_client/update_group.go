package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) UpdateGroup(
	groupIdentifier int,
	title string,
	color string,
	collapsed *bool,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e := c.failures.Caller(); e != nil {
		return e
	}

	g, okay := c.groups[groupIdentifier]

	if !okay {
		return not_found.New("group", groupIdentifier)
	}

	if title != "" {
		g.Title = title
	}

	if color != "" {
		g.Color = color
	}

	if collapsed != nil {
		g.Collapsed = *collapsed
	}

	return nil
}

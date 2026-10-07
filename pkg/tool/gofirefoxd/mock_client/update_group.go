package mock_client

import "fmt"

func (c *Client) UpdateGroup(
	groupIdentifier int,
	title string,
	color string,
	collapsed *bool,
) error {
	g, okay := c.groups[groupIdentifier]

	if !okay {
		return fmt.Errorf("group %d not found", groupIdentifier)
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

	c.groups[groupIdentifier] = g

	return nil
}

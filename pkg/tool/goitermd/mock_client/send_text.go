package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) SendText(
	identifier string,
	text string,
) error {
	s, okay := c.screens[identifier]

	if !okay {
		return not_found.Format("session %s not found", identifier)
	}

	s.Lines = append(s.Lines, text)
	c.screens[identifier] = s

	return nil
}

package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) SetTabTitle(
	tabIdentifier string,
	title string,
) error {
	for i, s := range c.sessions {
		if s.TabIdentifier == tabIdentifier {
			c.sessions[i].TabTitle = title

			return nil
		}
	}

	return not_found.Format("tab %s not found", tabIdentifier)
}

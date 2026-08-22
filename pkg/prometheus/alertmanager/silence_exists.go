package alertmanager

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) SilenceExists(name string) (bool, error) {
	_, e := c.SilenceByRule(name)

	if e != nil {
		if not_found.Is(e) {
			return false, nil
		}

		return false, e
	}

	return true, nil
}

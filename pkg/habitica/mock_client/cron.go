package mock_client

import "github.com/funtimecoding/soil/pkg/habitica/cron"

func (c *Client) Cron() (*cron.Cron, error) {
	if c.Failure != nil {
		return cron.Stub(), c.Failure
	}

	return cron.New(), nil
}

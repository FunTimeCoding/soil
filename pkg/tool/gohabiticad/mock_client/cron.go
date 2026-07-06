package mock_client

import "github.com/funtimecoding/soil/pkg/habitica/cron"

func (c *Client) Cron() (*cron.Cron, error) {
	return cron.New(), nil
}

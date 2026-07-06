package mock_client

import "github.com/funtimecoding/soil/pkg/habitica/cron"

func (c *Client) MustCron() *cron.Cron {
	return cron.New()
}

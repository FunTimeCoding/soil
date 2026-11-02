package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Setting(x context.Context) (*gitlab.Settings, error) {
	result, _, e := c.client.Settings.GetSettings(gitlab.WithContext(x))

	return result, wrapError(e)
}

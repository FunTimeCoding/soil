package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) DeletePost(_ *model.Post) error {
	return c.Failure
}

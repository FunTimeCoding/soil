package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Me() (*model.User, error) {
	return c.me, nil
}

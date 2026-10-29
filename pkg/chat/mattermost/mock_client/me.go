package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Me() (*model.User, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.me, nil
}

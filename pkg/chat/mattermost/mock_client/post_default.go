package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) PostDefault(_ string) (*model.Post, error) {
	return nil, c.Failure
}

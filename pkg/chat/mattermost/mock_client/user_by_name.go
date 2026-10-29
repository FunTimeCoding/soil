package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) UserByName(_ string) (*model.User, error) {
	return nil, c.Failure
}

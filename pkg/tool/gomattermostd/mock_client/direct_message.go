package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) DirectMessage(
	_ *model.User,
	_ string,
) (*model.Post, error) {
	return nil, nil
}

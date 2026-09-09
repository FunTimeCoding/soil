package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Users(
	_ int,
	_ int,
) ([]*model.User, error) {
	return nil, nil
}

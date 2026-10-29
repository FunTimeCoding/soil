package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) SearchUsers(
	_ *model.Team,
	_ string,
	_ int,
) ([]*model.User, error) {
	return nil, c.Failure
}

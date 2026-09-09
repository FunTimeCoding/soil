package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Channels(
	_ *model.Team,
	_ *model.User,
) ([]*model.Channel, error) {
	return nil, nil
}

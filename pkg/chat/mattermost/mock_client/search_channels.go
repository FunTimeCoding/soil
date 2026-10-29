package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) SearchChannels(
	_ *model.Team,
	_ string,
) ([]*model.Channel, error) {
	return nil, c.Failure
}

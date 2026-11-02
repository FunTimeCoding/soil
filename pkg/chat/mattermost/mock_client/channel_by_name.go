package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) ChannelByName(
	_ *model.Team,
	name string,
) (*model.Channel, error) {
	return c.TeamChannel(name)
}

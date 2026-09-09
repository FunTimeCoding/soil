package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) MustAllChannels(
	_ int,
	_ int,
) []*model.ChannelWithTeamData {
	return nil
}

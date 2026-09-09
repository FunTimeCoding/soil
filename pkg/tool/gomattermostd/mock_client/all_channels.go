package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) AllChannels(
	_ int,
	_ int,
) ([]*model.ChannelWithTeamData, error) {
	return nil, nil
}

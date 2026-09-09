package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) MustTeamChannel(name string) *model.Channel {
	return &model.Channel{Name: name}
}

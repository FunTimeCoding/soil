package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) ChannelUsers(_ *model.Channel) ([]*model.User, error) {
	return nil, c.Failure
}

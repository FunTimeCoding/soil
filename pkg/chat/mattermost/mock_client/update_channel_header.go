package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) UpdateChannelHeader(
	_ *model.Channel,
	_ string,
	_ string,
	_ string,
) (*model.Channel, error) {
	return nil, c.Failure
}

package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) WipePrefix(
	_ *model.Channel,
	_ string,
) {
}

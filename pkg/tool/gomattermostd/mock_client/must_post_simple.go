package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) MustPostSimple(
	_ *model.Channel,
	_ string,
) *model.Post {
	return nil
}

package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) DefaultChannel() *model.Channel {
	return nil
}

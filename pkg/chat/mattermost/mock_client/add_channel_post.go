package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) AddChannelPost(
	channel string,
	identifier string,
) {
	c.post[identifier] = &model.Post{Id: identifier, ChannelId: channel}
}

package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Channel(identifier string) (*model.Channel, error) {
	result, r, e := c.client.GetChannel(c.context, identifier)

	return result, wrapError(r, e)
}

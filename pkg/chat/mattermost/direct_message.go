package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) DirectMessage(
	u *model.User,
	text string,
) (*model.Post, error) {
	me, e := c.Me()

	if e != nil {
		return nil, e
	}

	channel, _, e := c.client.CreateDirectChannel(c.context, me.Id, u.Id)

	if e != nil {
		return nil, wrapError(e)
	}

	return c.PostSimple(channel, text)
}

package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) React(
	p *model.Post,
	emoji string,
) error {
	me, e := c.Me()

	if e != nil {
		return e
	}

	_, r, f := c.client.SaveReaction(
		c.context,
		&model.Reaction{UserId: me.Id, PostId: p.Id, EmojiName: emoji},
	)

	return wrapError(r, f)
}

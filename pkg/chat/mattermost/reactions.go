package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Reactions(p *model.Post) ([]*model.Reaction, error) {
	result, r, e := c.client.GetReactions(c.context, p.Id)

	return result, wrapError(r, e)
}

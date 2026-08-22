package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Reactions(p *model.Post) ([]*model.Reaction, error) {
	result, _, e := c.client.GetReactions(c.context, p.Id)

	return result, wrapError(e)
}

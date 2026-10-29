package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Post(p *model.Post) (*model.Post, error) {
	result, r, e := c.client.CreatePost(c.context, p)

	return result, wrapError(r, e)
}

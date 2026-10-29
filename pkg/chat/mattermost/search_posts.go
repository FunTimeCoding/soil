package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) SearchPosts(
	t *model.Team,
	terms string,
) (*model.PostList, error) {
	result, r, e := c.client.SearchPosts(c.context, t.Id, terms, false)

	return result, wrapError(r, e)
}

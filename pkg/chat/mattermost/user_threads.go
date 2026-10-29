package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) UserThreads(
	u *model.User,
	t *model.Team,
	o model.GetUserThreadsOpts,
) (*model.Threads, error) {
	result, r, e := c.client.GetUserThreads(c.context, u.Id, t.Id, o)

	return result, wrapError(r, e)
}

package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) UserThreads(
	_ *model.User,
	_ *model.Team,
	_ model.GetUserThreadsOpts,
) (*model.Threads, error) {
	return nil, c.Failure
}

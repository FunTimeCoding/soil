package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Posts(_ *model.Channel) (*model.PostList, error) {
	return nil, c.Failure
}

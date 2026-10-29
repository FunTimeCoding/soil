package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Teams(_ string) ([]*model.Team, error) {
	return nil, c.Failure
}

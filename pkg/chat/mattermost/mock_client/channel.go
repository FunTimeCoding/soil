package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Channel(identifier string) (*model.Channel, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return &model.Channel{Id: identifier}, nil
}

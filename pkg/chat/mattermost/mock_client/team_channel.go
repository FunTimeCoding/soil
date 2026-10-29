package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) TeamChannel(name string) (*model.Channel, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return &model.Channel{Name: name}, nil
}

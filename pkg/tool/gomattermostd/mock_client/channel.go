package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Channel(_ string) (*model.Channel, error) {
	return nil, nil
}

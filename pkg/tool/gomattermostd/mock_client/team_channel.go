package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) TeamChannel(_ string) (*model.Channel, error) {
	return nil, nil
}

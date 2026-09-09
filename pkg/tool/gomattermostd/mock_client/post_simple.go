package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) PostSimple(
	_ *model.Channel,
	_ string,
) (*model.Post, error) {
	return nil, nil
}

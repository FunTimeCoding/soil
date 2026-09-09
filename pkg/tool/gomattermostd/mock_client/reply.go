package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Reply(
	_ *model.Channel,
	_ *model.Post,
	_ string,
) (*model.Post, error) {
	return nil, nil
}

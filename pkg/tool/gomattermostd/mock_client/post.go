package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Post(_ *model.Post) (*model.Post, error) {
	return nil, nil
}

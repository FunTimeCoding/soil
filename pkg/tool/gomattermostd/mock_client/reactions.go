package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Reactions(_ *model.Post) ([]*model.Reaction, error) {
	return nil, nil
}

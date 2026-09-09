package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) React(
	_ *model.Post,
	_ string,
) error {
	return nil
}

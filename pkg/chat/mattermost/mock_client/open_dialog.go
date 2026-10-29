package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) OpenDialog(_ model.OpenDialogRequest) error {
	return c.Failure
}

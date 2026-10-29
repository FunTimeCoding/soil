package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) SubmitDialog(_ model.SubmitDialogRequest) (*model.SubmitDialogResponse, error) {
	return nil, c.Failure
}

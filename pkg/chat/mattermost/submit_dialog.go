package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) SubmitDialog(v model.SubmitDialogRequest) (*model.SubmitDialogResponse, error) {
	result, _, e := c.client.SubmitInteractiveDialog(c.context, v)

	return result, wrapError(e)
}

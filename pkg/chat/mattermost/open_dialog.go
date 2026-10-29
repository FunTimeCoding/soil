package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) OpenDialog(v model.OpenDialogRequest) error {
	r, e := c.client.OpenInteractiveDialog(c.context, v)

	return wrapError(r, e)
}

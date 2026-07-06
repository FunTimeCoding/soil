package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustSubmitDialog(v model.SubmitDialogRequest) *model.SubmitDialogResponse {
	result, e := c.SubmitDialog(v)
	errors.PanicOnError(e)

	return result
}

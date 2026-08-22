package mattermost

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
)

func (c *Client) DeletePost(p *model.Post) error {
	_, e := c.client.DeletePost(c.context, p.Id)

	if f, okay := errors.AsType[*model.AppError](e); okay &&
		f.StatusCode == http.StatusNotFound {
		return not_found.New(constant.MattermostPostField, p.Id)
	}

	return wrapError(e)
}

package mattermost

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
)

func (c *Client) DeletePost(p *model.Post) error {
	_, e := c.client.DeletePost(c.context, p.Id)

	if f, okay := errors.AsType[*model.AppError](e); okay &&
		f.StatusCode == http.StatusNotFound {
		return fmt.Errorf(
			"post not found: %s: %w",
			p.Id,
			constant.ErrorMattermostNotFound,
		)
	}

	return e
}

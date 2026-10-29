package mattermost

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/constant"
)

func (c *Client) Reachable(x context.Context) error {
	_, r, e := c.client.GetMe(x, constant.MattermostEmptyEntityTag)

	return wrapError(r, e)
}

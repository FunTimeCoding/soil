package mattermost

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustRefreshSocket() {
	errors.PanicOnError(c.RefreshSocket(context.Background()))
}

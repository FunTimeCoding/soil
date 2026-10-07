package mattermost

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustRefreshSocket() {
	errors.PanicOnError(c.RefreshSocket())
}

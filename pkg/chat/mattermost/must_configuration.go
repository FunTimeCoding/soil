package mattermost

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustConfiguration() map[string]string {
	result, e := c.Configuration()
	errors.PanicOnError(e)

	return result
}

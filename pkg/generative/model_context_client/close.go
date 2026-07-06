package model_context_client

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) Close() {
	errors.PanicClose(c.session)
}

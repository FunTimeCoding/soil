package opnsense

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/opnsense/request"
)

func (c *Client) MustSetHost(
	identifier string,
	h *request.Host,
) {
	errors.PanicOnError(c.SetHost(identifier, h))
}

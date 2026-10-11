package technitium

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/technitium/response"
)

func (c *Client) MustSetZoneOptions(o *response.ZoneOptions) {
	errors.PanicOnError(c.SetZoneOptions(o))
}

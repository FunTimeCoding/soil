package technitium

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/technitium/response"
)

func (c *Client) MustZoneOptions(zone string) *response.ZoneOptions {
	result, e := c.ZoneOptions(zone)
	errors.PanicOnError(e)

	return result
}

package chroma

import (
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) Databases(t v2.Tenant) []v2.Database {
	result, e := c.client.ListDatabases(c.context, t)
	errors.PanicOnError(e)

	return result
}

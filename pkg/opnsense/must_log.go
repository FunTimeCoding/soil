package opnsense

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/opnsense/log_entry"
)

func (c *Client) MustLog(limit int) []*log_entry.Entry {
	result, e := c.Log(limit)
	errors.PanicOnError(e)

	return result
}

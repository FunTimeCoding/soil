package opnsense

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/opnsense/blocklist"
)

func (c *Client) MustBlocklists(phrase string) []*blocklist.Blocklist {
	result, e := c.Blocklists(phrase)
	errors.PanicOnError(e)

	return result
}

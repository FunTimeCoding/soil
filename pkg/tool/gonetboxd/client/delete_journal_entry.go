package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) DeleteJournalEntry(identifier int32) {
	result, e := c.client.DeleteJournalEntry(c.context, identifier)
	errors.PanicOnError(e)
	web.ReadString(result)
}

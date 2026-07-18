package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListDeviceJournalEntries(
	device string,
	limit int32,
	offset int32,
) string {
	p := &client.ListDeviceJournalEntriesParams{}

	if limit > 0 {
		p.Limit = &limit
	}

	if offset > 0 {
		p.Offset = &offset
	}

	result, e := c.client.ListDeviceJournalEntries(c.context, device, p)
	errors.PanicOnError(e)

	return web.ReadString(result)
}

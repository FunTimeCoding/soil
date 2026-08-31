package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListVirtualJournalEntries(
	machine string,
	limit int32,
	offset int32,
) *response.Response {
	p := &client.ListVirtualJournalEntriesParams{}

	if limit > 0 {
		p.Limit = &limit
	}

	if offset > 0 {
		p.Offset = &offset
	}

	result, e := c.client.ListVirtualJournalEntries(c.context, machine, p)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}

package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) AddVirtualJournalEntry(
	machine string,
	kind string,
	comments string,
) string {
	body := client.AddVirtualJournalEntryJSONRequestBody{
		Comments: comments,
	}

	if kind != "" {
		body.Kind = &kind
	}

	result, e := c.client.AddVirtualJournalEntry(c.context, machine, body)
	errors.PanicOnError(e)

	return web.ReadString(result)
}

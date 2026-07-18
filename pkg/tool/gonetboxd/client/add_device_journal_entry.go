package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) AddDeviceJournalEntry(
	device string,
	kind string,
	comments string,
) string {
	body := client.AddDeviceJournalEntryJSONRequestBody{
		Comments: comments,
	}

	if kind != "" {
		body.Kind = &kind
	}

	result, e := c.client.AddDeviceJournalEntry(c.context, device, body)
	errors.PanicOnError(e)

	return web.ReadString(result)
}

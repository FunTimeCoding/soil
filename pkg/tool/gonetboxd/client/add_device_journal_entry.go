package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) AddDeviceJournalEntry(
	x context.Context,
	device string,
	kind string,
	comments string,
) (client.JournalEntry, error) {
	body := client.AddDeviceJournalEntryJSONRequestBody{Comments: comments}

	if kind != "" {
		body.Kind = &kind
	}

	r, e := c.client.AddDeviceJournalEntryWithResponse(x, device, body)

	if e != nil {
		return client.JournalEntry{}, e
	}

	if r.JSON201 == nil {
		return client.JournalEntry{}, refusal.New(
			"add device journal entry",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}

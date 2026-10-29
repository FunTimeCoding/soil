package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) AddVirtualJournalEntry(
	x context.Context,
	machine string,
	kind string,
	comments string,
) (client.JournalEntry, error) {
	body := client.AddVirtualJournalEntryJSONRequestBody{Comments: comments}

	if kind != "" {
		body.Kind = &kind
	}

	r, e := c.client.AddVirtualJournalEntryWithResponse(x, machine, body)

	if e != nil {
		return client.JournalEntry{}, e
	}

	if r.JSON201 == nil {
		return client.JournalEntry{}, refusal.New(
			"add virtual journal entry",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}

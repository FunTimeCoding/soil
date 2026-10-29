package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) UpdateJournalEntry(
	x context.Context,
	identifier int32,
	kind string,
	comments string,
) (client.JournalEntry, error) {
	body := client.UpdateJournalEntryJSONRequestBody{}

	if comments != "" {
		body.Comments = &comments
	}

	if kind != "" {
		body.Kind = &kind
	}

	r, e := c.client.UpdateJournalEntryWithResponse(x, identifier, body)

	if e != nil {
		return client.JournalEntry{}, e
	}

	if r.JSON200 == nil {
		return client.JournalEntry{}, refusal.New(
			"update journal entry",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}

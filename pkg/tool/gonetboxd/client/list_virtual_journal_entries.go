package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListVirtualJournalEntries(
	x context.Context,
	machine string,
	limit int32,
	offset int32,
) ([]*client.JournalEntry, error) {
	p := &client.ListVirtualJournalEntriesParams{}

	if limit > 0 {
		p.Limit = &limit
	}

	if offset > 0 {
		p.Offset = &offset
	}

	r, e := c.client.ListVirtualJournalEntriesWithResponse(x, machine, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(
			"list virtual journal entries",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}

package client

import "context"

func (c *Client) DeleteJournalEntry(
	x context.Context,
	identifier int32,
) error {
	r, e := c.client.DeleteJournalEntryWithResponse(x, identifier)

	if e != nil {
		return e
	}

	return noContent("delete journal entry", r.HTTPResponse, r.Body)
}

package netbox

import "github.com/funtimecoding/soil/pkg/netbox/refusal"

func (c *Client) DeleteJournalEntry(identifier int32) error {
	_, e := c.client.ExtrasAPI.
		ExtrasJournalEntriesDestroy(c.context, identifier).
		Execute()

	return refusal.Classify(e)
}

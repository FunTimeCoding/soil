package netbox

func (c *Client) DeleteJournalEntry(identifier int32) error {
	_, e := c.client.ExtrasAPI.
		ExtrasJournalEntriesDestroy(c.context, identifier).
		Execute()

	return e
}

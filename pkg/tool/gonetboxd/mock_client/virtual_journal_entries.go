package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/journal_entry"

func (c *Client) VirtualJournalEntries(
	_ string,
	_ int32,
	_ int32,
) ([]*journal_entry.Entry, error) {
	return nil, nil
}

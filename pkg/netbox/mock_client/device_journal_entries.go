package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/journal_entry"

func (c *Client) DeviceJournalEntries(
	_ string,
	_ int32,
	_ int32,
) ([]*journal_entry.Entry, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}

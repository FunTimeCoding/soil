package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/journal_entry"

func (c *Client) AddDeviceJournalEntry(
	_ string,
	_ string,
	_ string,
) (*journal_entry.Entry, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}

package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/journal_entry"
)

func (c *Client) DeviceJournalEntries(
	device string,
	limit int32,
	offset int32,
) ([]*journal_entry.Entry, error) {
	d, e := c.DeviceByName(device)

	if e != nil {
		return nil, e
	}

	return c.journalEntries(
		constant.DeviceAddress,
		d.Identifier,
		limit,
		offset,
	)
}

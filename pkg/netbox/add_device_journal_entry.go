package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/journal_entry"
)

func (c *Client) AddDeviceJournalEntry(
	device string,
	kind string,
	comments string,
) (*journal_entry.Entry, error) {
	d, e := c.DeviceByName(device)

	if e != nil {
		return nil, e
	}

	return c.addJournalEntry(
		constant.DeviceAddress,
		int64(d.Identifier),
		kind,
		comments,
	)
}

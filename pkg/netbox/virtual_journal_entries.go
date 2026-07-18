package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/journal_entry"
)

func (c *Client) VirtualJournalEntries(
	machine string,
	limit int32,
	offset int32,
) ([]*journal_entry.Entry, error) {
	m, e := c.VirtualMachineByName(machine)

	if e != nil {
		return nil, e
	}

	return c.journalEntries(
		constant.VirtualMachineAddress,
		m.Identifier,
		limit,
		offset,
	)
}

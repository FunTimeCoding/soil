package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/journal_entry"
)

func (c *Client) AddVirtualJournalEntry(
	machine string,
	kind string,
	comments string,
) (*journal_entry.Entry, error) {
	m, e := c.VirtualMachineByName(machine)

	if e != nil {
		return nil, e
	}

	return c.addJournalEntry(
		constant.VirtualMachineAddress,
		int64(m.Identifier),
		kind,
		comments,
	)
}

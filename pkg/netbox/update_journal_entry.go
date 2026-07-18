package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/journal_entry"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) UpdateJournalEntry(
	identifier int32,
	kind string,
	comments string,
) (*journal_entry.Entry, error) {
	w := netbox.NewPatchedWritableJournalEntryRequest()

	if comments != "" {
		w.SetComments(comments)
	}

	if kind != "" {
		v, e := netbox.NewJournalEntryKindValueFromValue(kind)

		if e != nil {
			return nil, e
		}

		w.SetKind(*v)
	}

	result, _, e := c.client.ExtrasAPI.
		ExtrasJournalEntriesPartialUpdate(c.context, identifier).
		PatchedWritableJournalEntryRequest(*w).
		Execute()

	if e != nil {
		return nil, e
	}

	return journal_entry.New(result), nil
}

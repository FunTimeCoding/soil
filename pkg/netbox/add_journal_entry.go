package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/journal_entry"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) addJournalEntry(
	objectType string,
	objectIdentifier int64,
	kind string,
	comments string,
) (*journal_entry.Entry, error) {
	w := netbox.NewWritableJournalEntryRequest(
		objectType,
		objectIdentifier,
		comments,
	)

	if kind != "" {
		v, e := netbox.NewJournalEntryKindValueFromValue(kind)

		if e != nil {
			return nil, e
		}

		w.SetKind(*v)
	}

	result, _, e := c.client.ExtrasAPI.
		ExtrasJournalEntriesCreate(c.context).
		WritableJournalEntryRequest(*w).
		Execute()

	if e != nil {
		return nil, e
	}

	return journal_entry.New(result), nil
}

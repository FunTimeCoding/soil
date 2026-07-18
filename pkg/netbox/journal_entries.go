package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/journal_entry"
)

func (c *Client) journalEntries(
	objectType string,
	objectIdentifier int32,
	limit int32,
	offset int32,
) ([]*journal_entry.Entry, error) {
	if limit <= 0 {
		limit = constant.PageLimit
	}

	result, _, e := c.client.ExtrasAPI.
		ExtrasJournalEntriesList(c.context).
		AssignedObjectType(objectType).
		AssignedObjectId([]int32{objectIdentifier}).
		Ordering("-created").
		Limit(limit).
		Offset(offset).
		Execute()

	if e != nil {
		return nil, e
	}

	return journal_entry.NewSlice(result.Results), nil
}

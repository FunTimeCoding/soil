package journal_entry

import (
	"github.com/netbox-community/go-netbox/v4"
	"time"
)

func New(v *netbox.JournalEntry) *Entry {
	result := &Entry{
		Identifier: v.Id,
		Comments:   v.Comments,
	}

	if v.Kind != nil && v.Kind.Value != nil {
		result.Kind = string(*v.Kind.Value)
	}

	if v.Created.IsSet() && v.Created.Get() != nil {
		result.Created = v.Created.Get().Format(time.RFC3339)
	}

	return result
}

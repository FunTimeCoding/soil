package journal_entry

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.JournalEntry) []*Entry {
	var result []*Entry

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}

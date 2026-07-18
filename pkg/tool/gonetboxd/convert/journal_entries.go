package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/journal_entry"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func JournalEntries(v []*journal_entry.Entry) []*server.JournalEntry {
	result := make([]*server.JournalEntry, 0, len(v))

	for _, e := range v {
		result = append(result, JournalEntry(e))
	}

	return result
}

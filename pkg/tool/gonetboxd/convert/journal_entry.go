package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/journal_entry"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func JournalEntry(e *journal_entry.Entry) *server.JournalEntry {
	result := &server.JournalEntry{
		Identifier: e.Identifier,
		Comments:   e.Comments,
	}

	if e.Kind != "" {
		result.Kind = &e.Kind
	}

	if e.Created != "" {
		result.Created = &e.Created
	}

	return result
}

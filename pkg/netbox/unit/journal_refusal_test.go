package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"net/http"
	"testing"
)

func TestAMissingJournalEntryOnUpdateIsNotFound(t *testing.T) {
	_, e := refusingNetbox(
		t,
		http.StatusNotFound,
		`{"detail":"No JournalEntry matches the given query."}`,
	).UpdateJournalEntry(9, "", "charlie")
	assert.True(t, not_found.Is(e))
}

func TestAMissingJournalEntryOnDeleteIsNotFound(t *testing.T) {
	e := refusingNetbox(
		t,
		http.StatusNotFound,
		`{"detail":"No JournalEntry matches the given query."}`,
	).DeleteJournalEntry(9)
	assert.True(t, not_found.Is(e))
}

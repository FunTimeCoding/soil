package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/unit/netbox_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestMissingJournalEntryOnUpdateIsNotFound(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNotFound),
		upstream_server.WithAnswer(
			`{"detail":"No JournalEntry matches the given query."}`,
		),
	)
	_, e := netbox_tester.New(t, s.URL).UpdateJournalEntry(9, "", "charlie")
	assert.True(t, not_found.Is(e))
}

func TestMissingJournalEntryOnDeleteIsNotFound(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNotFound),
		upstream_server.WithAnswer(
			`{"detail":"No JournalEntry matches the given query."}`,
		),
	)
	e := netbox_tester.New(t, s.URL).DeleteJournalEntry(9)
	assert.True(t, not_found.Is(e))
}

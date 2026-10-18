package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"net/http"
	"testing"
)

func TestARefusedPushIsAnError(t *testing.T) {
	e := newIndexer(
		t,
		http.StatusInternalServerError,
		`{"error":"disk full","event_identifier":"bravo"}`,
	).Push("notes", "charlie.md", "# Charlie\n", nil)
	assert.True(t, unexpected.Is(e))
	assert.String(t, "push document status: 500: disk full", e.Error())
}

func TestARefusedListLeavesNoPartialIndex(t *testing.T) {
	existing, e := newIndexer(
		t,
		http.StatusServiceUnavailable,
		`{"error":"restarting"}`,
	).Existing("notes")
	assert.Nil(t, existing)
	assert.String(t, "list documents: restarting", e.Error())
}

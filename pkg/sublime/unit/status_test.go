package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/sublime/unit/sublime_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestUnknownViewIsNotFound(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNotFound),
		upstream_server.WithAnswer(`{"error": "view not found"}`),
	)
	_, e := sublime_tester.New(t, s.URL).View(7)
	assert.True(t, not_found.Is(e))
	assert.String(t, "read view: view not found", e.Error())
}

func TestMissingStringIsValidation(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusBadRequest),
		upstream_server.WithAnswer(
			`{"error": "old_string not found in buffer"}`,
		),
	)
	_, e := sublime_tester.New(t, s.URL).EditView(7, "alfa", "bravo", false)
	assert.True(t, validation.Is(e))
	assert.String(t, "edit view: old_string not found in buffer", e.Error())
}

func TestAmbiguousEditIsConflict(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusConflict),
		upstream_server.WithAnswer(
			`{"error": "old_string matches multiple times; use replace_all"}`,
		),
	)
	_, e := sublime_tester.New(t, s.URL).EditView(7, "alfa", "bravo", false)
	assert.True(t, conflict.Is(e))
	assert.String(
		t,
		"edit view: old_string matches multiple times; use replace_all",
		e.Error(),
	)
}

func TestUnknownStatusIsFault(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusInternalServerError),
		upstream_server.WithAnswer(`{"error": "boom"}`),
	)
	e := sublime_tester.New(t, s.URL).CloseView(7)
	assert.True(t, unexpected.Is(e))
}

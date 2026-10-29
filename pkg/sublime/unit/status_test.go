package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"net/http"
	"testing"
)

func TestAnUnknownViewIsNotFound(t *testing.T) {
	_, e := bridge(
		t,
		http.StatusNotFound,
		`{"error": "view not found"}`,
	).View(7)
	assert.True(t, not_found.Is(e))
	assert.String(t, "read view: view not found", e.Error())
}

func TestAMissingStringIsValidation(t *testing.T) {
	_, e := bridge(
		t,
		http.StatusBadRequest,
		`{"error": "old_string not found in buffer"}`,
	).EditView(7, "alfa", "bravo", false)
	assert.True(t, validation.Is(e))
	assert.String(t, "edit view: old_string not found in buffer", e.Error())
}

func TestAnAmbiguousEditIsConflict(t *testing.T) {
	_, e := bridge(
		t,
		http.StatusConflict,
		`{"error": "old_string matches multiple times; use replace_all"}`,
	).EditView(7, "alfa", "bravo", false)
	assert.True(t, conflict.Is(e))
	assert.String(
		t,
		"edit view: old_string matches multiple times; use replace_all",
		e.Error(),
	)
}

func TestAnUnknownStatusIsAFault(t *testing.T) {
	e := bridge(t, http.StatusInternalServerError, `{"error": "boom"}`).
		CloseView(7)
	assert.True(t, unexpected.Is(e))
}

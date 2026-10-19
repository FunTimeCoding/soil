package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"net/http"
	"testing"
)

func TestAnUnknownKeyIsValidation(t *testing.T) {
	e := bridge(
		t,
		http.StatusUnprocessableEntity,
		`{"error": "unknown key: alfa. valid keys: enter, tab"}`,
	).SendKey("bravo", "alfa")
	assert.True(t, validation.Is(e))
	assert.String(
		t,
		"send key: unknown key: alfa. valid keys: enter, tab",
		e.Error(),
	)
}

func TestAKeyToAnUnknownSessionIsNotFound(t *testing.T) {
	e := bridge(
		t,
		http.StatusNotFound,
		`{"error": "session not found"}`,
	).SendKey("bravo", "enter")
	assert.True(t, not_found.Is(e))
	assert.String(t, "send key: session not found", e.Error())
}

func TestATabWithoutAWindowIsUnavailable(t *testing.T) {
	_, e := bridge(
		t,
		http.StatusServiceUnavailable,
		`{"error": "no window"}`,
	).CreateTab()
	assert.True(t, unavailable.Is(e))
	assert.String(t, "create tab: no window", e.Error())
}

func TestAMalformedRequestIsAFault(t *testing.T) {
	e := bridge(
		t,
		http.StatusBadRequest,
		`{"error": "invalid JSON"}`,
	).SendKey("bravo", "enter")
	assert.True(t, unexpected.Is(e))
	assert.String(t, "send key: 400: invalid JSON", e.Error())
}

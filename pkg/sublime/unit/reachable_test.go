package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"net/http"
	"testing"
)

func TestAnAnsweringPluginIsReachable(t *testing.T) {
	assert.FatalOnError(
		t,
		bridge(t, http.StatusOK, `{"views":[]}`).Reachable(t.Context()),
	)
}

func TestARestartingPluginIsUnavailable(t *testing.T) {
	e := bridge(
		t,
		http.StatusServiceUnavailable,
		`{"error":"Sublime is starting"}`,
	).Reachable(t.Context())
	assert.True(t, unavailable.Is(e))
}

func TestViewsFromARestartingPluginAreUnavailable(t *testing.T) {
	_, e := bridge(
		t,
		http.StatusServiceUnavailable,
		`{"error":"Sublime is starting"}`,
	).Views()
	assert.True(t, unavailable.Is(e))
}

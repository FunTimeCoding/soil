package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/sublime/unit/sublime_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestAnsweringPluginIsReachable(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer(`{"views":[]}`))
	assert.FatalOnError(t, sublime_tester.New(t, s.URL).Reachable(t.Context()))
}

func TestRestartingPluginIsUnavailable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusServiceUnavailable),
		upstream_server.WithAnswer(`{"error":"Sublime is starting"}`),
	)
	e := sublime_tester.New(t, s.URL).Reachable(t.Context())
	assert.True(t, unavailable.Is(e))
}

func TestViewsFromRestartingPluginAreUnavailable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusServiceUnavailable),
		upstream_server.WithAnswer(`{"error":"Sublime is starting"}`),
	)
	_, e := sublime_tester.New(t, s.URL).Views()
	assert.True(t, unavailable.Is(e))
}

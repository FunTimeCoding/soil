package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclaude"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
	"time"
)

func TestAwaitCallsignAsksAgainAfterReleasedHold(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(
			http.StatusNoContent,
			http.StatusNoContent,
			http.StatusNoContent,
			http.StatusNoContent,
			http.StatusOK,
		),
		upstream_server.WithAnswer(`{"callsign":"Ash"}`),
		upstream_server.WithObject(),
	)
	assert.String(
		t,
		"Ash",
		goclaude.AwaitCallsign(
			client_tester.New(t, s.URL),
			"session-1",
			time.Now(),
			time.Millisecond,
		),
	)
}

func TestAwaitCallsignGivesUpAfterRepeatedFailures(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusInternalServerError),
	)
	assert.String(
		t,
		"",
		goclaude.AwaitCallsign(
			client_tester.New(t, s.URL),
			"session-1",
			time.Now(),
			time.Millisecond,
		),
	)
}

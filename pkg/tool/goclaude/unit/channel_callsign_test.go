package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclaude"
	"net/http"
	"testing"
	"time"
)

func TestAwaitCallsignAsksAgainAfterAReleasedHold(t *testing.T) {
	c := stubServer(
		t,
		`{"callsign":"Ash"}`,
		http.StatusNoContent,
		http.StatusNoContent,
		http.StatusNoContent,
		http.StatusNoContent,
	)
	assert.String(
		t,
		"Ash",
		goclaude.AwaitCallsign(c, "session-1", time.Now(), time.Millisecond),
	)
}

func TestAwaitCallsignGivesUpAfterRepeatedFailures(t *testing.T) {
	c := stubServer(
		t,
		`{"callsign":"Ash"}`,
		http.StatusInternalServerError,
		http.StatusInternalServerError,
		http.StatusInternalServerError,
	)
	assert.String(
		t,
		"",
		goclaude.AwaitCallsign(c, "session-1", time.Now(), time.Millisecond),
	)
}

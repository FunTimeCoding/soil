package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/iterm"
	"github.com/funtimecoding/soil/pkg/iterm/face"
	"github.com/funtimecoding/soil/pkg/iterm/mock_client"
	"net/http"
	"reflect"
	"testing"
)

func TestAnsweringBridgeIsReachable(t *testing.T) {
	assert.FatalOnError(
		t,
		bridge(t, http.StatusOK, `{"sessions":[]}`).Reachable(t.Context()),
	)
}

func TestRestartingBridgeIsUnavailable(t *testing.T) {
	e := bridge(
		t,
		http.StatusServiceUnavailable,
		`{"error":"iTerm2 is starting"}`,
	).Reachable(t.Context())
	assert.True(t, unavailable.Is(e))
}

func TestSessionsFromRestartingBridgeAreUnavailable(t *testing.T) {
	_, e := bridge(
		t,
		http.StatusServiceUnavailable,
		`{"error":"iTerm2 is starting"}`,
	).Sessions()
	assert.True(t, unavailable.Is(e))
}

func TestItermClientAndItsMockShareFace(t *testing.T) {
	f := reflect.TypeFor[face.ItermClient]()
	assert.True(t, reflect.TypeFor[*iterm.Client]().Implements(f))
	assert.True(t, reflect.TypeFor[*mock_client.Client]().Implements(f))
}

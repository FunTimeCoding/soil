package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"testing"
)

func TestDeadTokenFailsProbe(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusUnauthorized, "")
	e := newJellyfin(t, s.URL).Reachable(t.Context())
	assert.NotNil(t, e)
	assert.String(t, "/System/Info", last().URL.Path)
}

func TestRestartingJellyfinIsUnavailable(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusServiceUnavailable,
		`{"message":"Jellyfin is starting"}`,
	)
	_, e := newJellyfin(t, s.URL).Libraries()
	assert.True(t, unavailable.Is(e))
}

func TestRefusedRequestIsFault(t *testing.T) {
	s, _ := upstream_tester.New(t, http.StatusForbidden, "")
	_, e := newJellyfin(t, s.URL).Sessions()
	assert.True(t, unexpected.Is(e))
}

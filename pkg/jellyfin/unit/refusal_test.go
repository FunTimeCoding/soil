package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/jellyfin/unit/jellyfin_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestDeadTokenFailsProbe(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusUnauthorized),
	)
	e := jellyfin_tester.New(t, s.URL).Reachable(t.Context())
	assert.NotNil(t, e)
	assert.String(t, "/System/Info", s.Last().URL.Path)
}

func TestRestartingJellyfinIsUnavailable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusServiceUnavailable),
		upstream_server.WithAnswer(`{"message":"Jellyfin is starting"}`),
	)
	_, e := jellyfin_tester.New(t, s.URL).Libraries()
	assert.True(t, unavailable.Is(e))
}

func TestRefusedRequestIsFault(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusForbidden),
	)
	_, e := jellyfin_tester.New(t, s.URL).Sessions()
	assert.True(t, unexpected.Is(e))
}

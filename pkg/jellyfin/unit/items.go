package unit

import (
	"github.com/funtimecoding/soil/pkg/jellyfin"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"testing"
)

func items(
	t *testing.T,
	answer string,
) *jellyfin.Client {
	t.Helper()
	s, _ := upstream_tester.New(t, http.StatusOK, answer)
	host, port := upstream_tester.HostPort(t, s.URL)

	return jellyfin.New(host, port, "alfa-token")
}

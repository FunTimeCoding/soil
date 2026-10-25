package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/unit/web_tester"
	"net/http"
	"testing"
)

func newMountedMux(t *testing.T) *http.ServeMux {
	t.Helper()
	result := http.NewServeMux()
	g := guard.New(result, []string{"alfa"})
	c := newAuthorizationClient(newFakeIssuer(t).URL, memory.New())
	t.Cleanup(c.Close)
	c.Mount(g)
	g.Session("GET /target", web_tester.Serve)

	return result
}

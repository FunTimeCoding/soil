package unit

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
	"net/http"
	"net/http/httptest"
	"testing"
)

func reachableServer(t *testing.T) *httptest.Server {
	t.Helper()
	m := http.NewServeMux()
	m.HandleFunc(
		route.Get(constant.HealthPath),
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			w.WriteHeader(http.StatusNoContent)
		},
	)
	guard.New(m, []string{"alfa"})
	s := httptest.NewServer(m)
	t.Cleanup(s.Close)

	return s
}

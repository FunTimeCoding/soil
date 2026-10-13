package unit

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newHeaderServer(t *testing.T) (*httptest.Server, *string) {
	t.Helper()
	var seen string
	result := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				seen = r.Header.Get(constant.Authorization)
				w.WriteHeader(http.StatusOK)
			},
		),
	)
	t.Cleanup(result.Close)

	return result, &seen
}

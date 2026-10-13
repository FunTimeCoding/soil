package unit

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newBasicAuthServer(t *testing.T) (*httptest.Server, *string) {
	t.Helper()
	var seen string
	result := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				user, password, _ := r.BasicAuth()
				seen = join.Empty(user, constant.Colon, password)
				w.WriteHeader(http.StatusOK)
			},
		),
	)
	t.Cleanup(result.Close)

	return result, &seen
}

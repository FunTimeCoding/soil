package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func answering(
	t *testing.T,
	status int,
	calls *atomic.Int32,
) *client.Client {
	t.Helper()
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				calls.Add(1)
				w.WriteHeader(status)
			},
		),
	)
	t.Cleanup(s.Close)
	result := client.New(s.URL, false, "")
	t.Cleanup(result.Close)

	return result
}

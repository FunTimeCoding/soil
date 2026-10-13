package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func newCountingServer(t *testing.T) (*httptest.Server, *atomic.Int64) {
	t.Helper()
	var opened atomic.Int64
	result := httptest.NewUnstartedServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				_, e := w.Write([]byte("counted"))
				errors.PanicOnError(e)
			},
		),
	)
	result.Config.ConnState = func(
		_ net.Conn,
		c http.ConnState,
	) {
		if c == http.StateNew {
			opened.Add(1)
		}
	}
	result.Start()
	t.Cleanup(result.Close)

	return result, &opened
}

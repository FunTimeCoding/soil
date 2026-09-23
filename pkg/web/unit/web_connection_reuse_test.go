package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	"net"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestConnectionReuse(t *testing.T) {
	var opened atomic.Int64
	s := httptest.NewUnstartedServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, _ *http.Request) {
				_, e := w.Write([]byte("connection-reuse-payload"))
				errors.PanicOnError(e)
			},
		),
	)
	s.Config.ConnState = func(_ net.Conn, c http.ConnState) {
		if c == http.StateNew {
			opened.Add(1)
		}
	}
	s.Start()
	defer s.Close()

	for range 4 {
		assert.String(
			t,
			"connection-reuse-payload",
			web.GetString(web.InsecureClient(), s.URL),
		)
	}

	assert.Integer(t, int64(1), opened.Load())
}

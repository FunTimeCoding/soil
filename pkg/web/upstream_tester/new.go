package upstream_tester

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func New(
	t *testing.T,
	status int,
	answer string,
) (*httptest.Server, func() *http.Request) {
	t.Helper()
	var mutex sync.Mutex
	var last *http.Request
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				sent, f := io.ReadAll(r.Body)
				errors.PanicOnError(f)
				kept := r.Clone(r.Context())
				kept.Body = io.NopCloser(bytes.NewReader(sent))
				mutex.Lock()
				last = kept
				mutex.Unlock()
				w.WriteHeader(status)
				_, e := w.Write([]byte(answer))
				errors.PanicOnError(e)
			},
		),
	)
	t.Cleanup(s.Close)

	return s, func() *http.Request {
		mutex.Lock()
		defer mutex.Unlock()

		return last
	}
}

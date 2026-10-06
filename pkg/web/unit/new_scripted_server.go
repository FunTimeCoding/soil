package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func newScriptedServer(
	t *testing.T,
	body string,
	statuses ...int,
) (*httptest.Server, func() []*http.Request) {
	t.Helper()
	var mutex sync.Mutex
	var seen []*http.Request
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
				seen = append(seen, kept)
				status := statuses[min(len(seen), len(statuses))-1]
				mutex.Unlock()
				w.WriteHeader(status)
				_, e := w.Write([]byte(body))
				errors.PanicOnError(e)
			},
		),
	)
	t.Cleanup(s.Close)

	return s, func() []*http.Request {
		mutex.Lock()
		defer mutex.Unlock()

		return seen
	}
}

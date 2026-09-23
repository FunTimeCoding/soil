package example

import (
	"example/errors"
	"net/http"
	"testing"
)

func newHandle(t *testing.T) *Handle {
	r, _ := http.Get("http://example.com")
	result := &Handle{response: r}
	t.Cleanup(func() { errors.PanicClose(result) })

	return result
}

func UseClosureArranged(t *testing.T) {
	h := newHandle(t)
	h.Ping()
}

func newServer(t *testing.T) *Server {
	r, _ := http.Get("http://example.com")
	result := &Server{response: r}
	t.Cleanup(result.Close)

	return result
}

func wrapServer(t *testing.T) *Server {
	s := newServer(t)
	s.Ping()

	return s
}

func UseArranged(t *testing.T) {
	s := newServer(t)
	s.Ping()
}

func UseWrapped(t *testing.T) {
	s := wrapServer(t)
	s.Ping()
}

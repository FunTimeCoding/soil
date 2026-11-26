package bidi_server

import (
	"github.com/funtimecoding/soil/pkg/browser/types/message"
	"net/http/httptest"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	s := &Server{t: t, answers: make(map[string]*message.Reply)}
	s.server = httptest.NewServer(s)
	t.Cleanup(s.server.Close)

	return s
}

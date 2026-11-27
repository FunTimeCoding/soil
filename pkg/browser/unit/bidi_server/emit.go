package bidi_server

import (
	"context"
	"github.com/coder/websocket/wsjson"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/message"
)

func (s *Server) Emit(
	method string,
	parameters string,
) {
	s.t.Helper()
	s.mutex.Lock()
	connection := s.connection
	s.mutex.Unlock()

	if connection == nil {
		s.t.Fatalf("emit %s: no client connected", method)
	}

	if e := wsjson.Write(
		context.Background(),
		connection,
		message.NewEvent(method, parameters),
	); e != nil {
		s.t.Fatalf("emit %s: %v", method, e)
	}
}

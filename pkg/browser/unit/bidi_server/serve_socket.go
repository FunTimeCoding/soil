package bidi_server

import (
	"context"
	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/funtimecoding/soil/pkg/browser/types/message"
	"net/http"
)

func (s *Server) serveSocket(
	w http.ResponseWriter,
	q *http.Request,
) {
	connection, e := websocket.Accept(w, q, nil)

	if e != nil {
		s.t.Errorf("accept: %v", e)

		return
	}

	s.mutex.Lock()
	s.connection = connection
	s.mutex.Unlock()

	for {
		command := message.NewRequest(0, "", nil)

		if f := wsjson.Read(context.Background(), connection, command); f != nil {
			return
		}

		if g := wsjson.Write(
			context.Background(),
			connection,
			s.reply(command),
		); g != nil {
			return
		}
	}
}

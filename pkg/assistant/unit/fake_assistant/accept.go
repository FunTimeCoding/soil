package fake_assistant

import (
	"github.com/gorilla/websocket"
	"net/http"
)

func (s *Server) accept(
	w http.ResponseWriter,
	q *http.Request,
) {
	signal(s.dialed)
	s.mutex.Lock()

	if s.refusal > 0 {
		s.refusal--
		s.mutex.Unlock()
		w.WriteHeader(http.StatusServiceUnavailable)

		return
	}

	invalid := s.invalid
	s.mutex.Unlock()
	u := websocket.Upgrader{}
	c, e := u.Upgrade(w, q, nil)

	if e != nil {
		return
	}

	s.converse(c, invalid)
}

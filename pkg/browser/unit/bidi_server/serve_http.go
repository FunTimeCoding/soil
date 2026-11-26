package bidi_server

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"net/http"
)

func (s *Server) ServeHTTP(
	w http.ResponseWriter,
	q *http.Request,
) {
	session := join.Slash(
		[]string{constant.SessionRoute, constant.FixtureSession},
	)

	switch {
	case q.URL.Path == constant.StatusRoute:
		writer.Print(w, "%s", `{"value":{"ready":true}}`)
	case q.URL.Path == constant.SessionRoute && q.Method == http.MethodPost:
		writer.Print(
			w,
			`{"value":{"sessionId":%q,"capabilities":{"webSocketUrl":%q}}}`,
			constant.FixtureSession,
			fmt.Sprintf("ws://%s%s", q.Host, session),
		)
	case q.URL.Path == session && q.Method == http.MethodDelete:
		s.mutex.Lock()
		s.deleted = true
		s.mutex.Unlock()
		writer.Print(w, "%s", `{"value":null}`)
	case q.URL.Path == session:
		s.serveSocket(w, q)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

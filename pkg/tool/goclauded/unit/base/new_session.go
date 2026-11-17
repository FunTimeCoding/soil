package base

import "github.com/funtimecoding/soil/pkg/tool/goclauded/unit/session_tester"

func (s *Server) NewSession() *session_tester.Session {
	s.t.Helper()

	return session_tester.New(s.t, s.Server)
}

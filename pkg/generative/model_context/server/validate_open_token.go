package server

import "github.com/funtimecoding/soil/pkg/console"

func (s *Server) validateOpenToken(token string) bool {
	if _, e := s.tokenVerifier().Verify(s.context, token); e != nil {
		console.Format("OIDC validate fail: %v\n", e)

		return false
	}

	return true
}

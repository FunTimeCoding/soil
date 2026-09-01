package server

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (s *Server) validateOpenToken(token string) bool {
	t, e := s.tokenVerifier().Verify(s.context, token)

	if e != nil {
		console.Format("OIDC validate fail: %v\n", e)

		return false
	}

	if false {
		// TODO: Log claims?
		claims := make(map[string]any)
		errors.PanicOnError(t.Claims(&claims))
		console.Format("OIDC claims: %+v\n", claims)
	}

	return true
}

package server

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (s *Server) validateOpenToken(token string) bool {
	t, e := s.tokenVerifier().Verify(s.context, token)

	if e != nil {
		fmt.Printf("OIDC validate fail: %v\n", e)

		return false
	}

	if false {
		// TODO: Log claims?
		claims := make(map[string]any)
		errors.PanicOnError(t.Claims(&claims))
		fmt.Printf("OIDC claims: %+v\n", claims)
	}

	return true
}

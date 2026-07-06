package server

import (
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (s *Server) tokenVerifier() *oidc.IDTokenVerifier {
	s.once.Do(
		func() {
			provider, e := oidc.NewProvider(
				s.context,
				s.authorizationLocator,
			)
			errors.PanicOnError(e)
			s.verifier = provider.Verifier(
				&oidc.Config{ClientID: s.clientIdentifier},
			)
		},
	)

	return s.verifier
}

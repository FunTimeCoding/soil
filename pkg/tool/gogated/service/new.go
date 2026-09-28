package service

import (
	"crypto/rsa"
	"github.com/funtimecoding/soil/pkg/tool/gogated/store"
	"github.com/ory/fosite"
)

func New(
	s *store.Store,
	p fosite.OAuth2Provider,
	signingKey *rsa.PrivateKey,
	keyIdentifier string,
	issuer string,
) *Service {
	return &Service{
		store:         s,
		provider:      p,
		signingKey:    signingKey,
		keyIdentifier: keyIdentifier,
		issuer:        issuer,
	}
}

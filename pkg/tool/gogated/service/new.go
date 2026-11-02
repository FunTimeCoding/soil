package service

import (
	"crypto/rsa"
	"github.com/funtimecoding/soil/pkg/directory/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gogated/store"
)

func New(
	s *store.Store,
	signingKey *rsa.PrivateKey,
	keyIdentifier string,
	secret string,
	issuer string,
	superUserMail string,
	d face.DirectoryClient,
	directoryGuard *peer.Guard,
) *Service {
	return &Service{
		store:          s,
		provider:       provider(secret, issuer, s, signingKey),
		signingKey:     signingKey,
		keyIdentifier:  keyIdentifier,
		issuer:         issuer,
		superUserMail:  superUserMail,
		directory:      d,
		directoryGuard: directoryGuard,
	}
}

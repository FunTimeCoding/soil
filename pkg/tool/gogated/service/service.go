package service

import (
	"crypto/rsa"
	"github.com/funtimecoding/soil/pkg/directory/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gogated/store"
	"github.com/ory/fosite"
)

type Service struct {
	store          *store.Store
	provider       fosite.OAuth2Provider
	signingKey     *rsa.PrivateKey
	keyIdentifier  string
	issuer         string
	superUserMail  string
	directory      face.DirectoryClient
	directoryGuard *peer.Guard
}

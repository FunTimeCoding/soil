package client

import (
	"crypto/cipher"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"sync"
)

type Client struct {
	issuer          string
	identifier      string
	secret          string
	signInPath      string
	callbackLocator string
	seal            cipher.AEAD
	requester       *requester.Requester
	reporter        face.Reporter
	provider        *oidc.Provider
	verifier        *oidc.IDTokenVerifier
	providerMutex   sync.Mutex
}

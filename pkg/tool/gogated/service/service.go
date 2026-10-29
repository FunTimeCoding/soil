package service

import (
	"crypto/rsa"
	"github.com/funtimecoding/soil/pkg/directory/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gogated/store"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/ory/fosite"
	"gorm.io/gorm"
)

type Service struct {
	database      *gorm.DB
	store         *store.Store
	provider      fosite.OAuth2Provider
	signingKey    *rsa.PrivateKey
	keyIdentifier string
	issuer        string
	superUserMail string
	directory      face.Client
	directoryGuard *peer.Guard
	authorization  *client.Client
}

package service

import (
	"context"
	"crypto/rsa"
	"github.com/funtimecoding/soil/pkg/tool/gogated/store"
	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"
	"github.com/ory/fosite/token/jwt"
	"time"
)

func provider(
	secret string,
	issuer string,
	s *store.Store,
	signingKey *rsa.PrivateKey,
) fosite.OAuth2Provider {
	c := &fosite.Config{
		GlobalSecret:          []byte(secret),
		AccessTokenLifespan:   time.Hour,
		RefreshTokenLifespan:  30 * 24 * time.Hour,
		AuthorizeCodeLifespan: 10 * time.Minute,
		IDTokenIssuer:         issuer,
		EnforcePKCE:           true,
		ScopeStrategy:         fosite.WildcardScopeStrategy,
	}
	keyGetter := func(context.Context) (interface{}, error) {
		return signingKey, nil
	}

	return compose.Compose(
		c,
		s,
		&compose.CommonStrategy{
			CoreStrategy: compose.NewOAuth2HMACStrategy(c),
			OpenIDConnectTokenStrategy: compose.NewOpenIDConnectStrategy(
				keyGetter,
				c,
			),
			Signer: &jwt.DefaultSigner{GetPrivateKey: keyGetter},
		},
		compose.OAuth2AuthorizeExplicitFactory,
		compose.OAuth2PKCEFactory,
		compose.OAuth2RefreshTokenGrantFactory,
		compose.OpenIDConnectExplicitFactory,
	)
}

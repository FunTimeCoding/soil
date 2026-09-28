package base

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gogated/server"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/tool/gogated/store"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"
	"github.com/ory/fosite/token/jwt"
	"net/http"
	"testing"
	"time"
)

func New(t *testing.T) *Server {
	t.Helper()
	database := lite.NewMemory()
	migrate.AutoMigrate(database)
	o := store.New(database)
	signingKey, keyIdentifier, e := o.EnsureSigningKey()
	assert.FatalOnError(t, e)
	errors.PanicOnError(
		o.SeedAdministrator(constant.FixtureMail, constant.FixturePassword),
	)
	c := &fosite.Config{
		GlobalSecret:               []byte(constant.FixtureSecret),
		AccessTokenLifespan:        time.Hour,
		RefreshTokenLifespan:       30 * 24 * time.Hour,
		AuthorizeCodeLifespan:      10 * time.Minute,
		IDTokenIssuer:              constant.FixtureIssuer,
		EnforcePKCE:                true,
		ScopeStrategy:              fosite.WildcardScopeStrategy,
		SendDebugMessagesToClients: true,
	}
	keyGetter := func(context.Context) (interface{}, error) {
		return signingKey, nil
	}
	provider := compose.Compose(
		c,
		o,
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
	s := service.New(
		o,
		provider,
		signingKey,
		keyIdentifier,
		constant.FixtureIssuer,
	)
	h := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			server.New(s).Mount(g)
		},
	)
	result := &Server{Store: o, Service: s, Web: h}
	t.Cleanup(result.Close)

	return result
}

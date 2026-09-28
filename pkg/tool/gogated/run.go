package gogated

import (
	"context"
	"github.com/funtimecoding/soil/pkg/directory"
	directoryConstant "github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gogated/option"
	"github.com/funtimecoding/soil/pkg/tool/gogated/server"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/tool/gogated/store"
	"github.com/funtimecoding/soil/pkg/tool/gogated/web"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"
	"github.com/ory/fosite/token/jwt"
	"net/http"
	"time"
)

func Run(
	o *option.Option,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	m := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(m)
	t := store.New(m)
	signingKey, keyIdentifier, e := t.EnsureSigningKey()
	errors.PanicOnError(e)
	errors.PanicOnError(
		t.SeedAdministrator(o.SuperUserMail, o.SuperUserPassword),
	)
	c := &fosite.Config{
		GlobalSecret:          []byte(o.Secret),
		AccessTokenLifespan:   time.Hour,
		RefreshTokenLifespan:  30 * 24 * time.Hour,
		AuthorizeCodeLifespan: 10 * time.Minute,
		IDTokenIssuer:         o.Issuer,
		EnforcePKCE:           true,
		ScopeStrategy:         fosite.WildcardScopeStrategy,
	}
	keyGetter := func(context.Context) (interface{}, error) {
		return signingKey, nil
	}
	provider := compose.Compose(
		c,
		t,
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
	s := service.New(t, provider, signingKey, keyIdentifier, o.Issuer)

	if environment.Exists(directoryConstant.HostEnvironment) {
		s = s.WithDirectory(directory.NewEnvironment())
	}

	errors.PanicOnError(
		s.SeedAdminClient(
			o.AdminClientIdentifier,
			o.AdminClientSecret,
			o.Issuer,
		),
	)
	authorization := client.New(
		o.Issuer,
		o.AdminClientIdentifier,
		o.AdminClientSecret,
		webConstant.SignInPath,
		join.Empty(o.Issuer, webConstant.CallbackPath),
		client.DeriveKey(o.Secret),
	)
	v := server.New(s)
	administration := web.New(s, authorization, o.SuperUserMail)
	go runCleanupLoop(s)
	lifecycle.New(
		l,
		lifecycle.WithServer(
			lifecycleServer.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					Mount(
						v,
						administration,
						s,
						r,
						i.Recorder(),
						o.Version,
						guard.New(m, o.ServiceTokens),
					)
				},
			).WithMiddleware(administration.Recovery(r)),
		),
	).RunUntilSignal()
}

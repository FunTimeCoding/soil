package daemon

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gogated/option"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/tool/gogated/store"
	"github.com/funtimecoding/soil/pkg/tool/gogated/web"
)

func New(
	o *option.Option,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d)
	s := store.New(d)
	signingKey, keyIdentifier, e := s.EnsureSigningKey()
	errors.PanicOnError(e)
	errors.PanicOnError(
		s.SeedAdministrator(o.SuperUserMail, o.SuperUserPassword),
	)
	errors.PanicOnError(
		s.SeedAdminClient(
			o.AdminClientIdentifier,
			o.AdminClientSecret,
			o.Issuer,
		),
	)
	result := &Daemon{
		database:      d,
		authorization: authorizationClient(o).WithReporter(r),
		logger:        l,
		reporter:      r,
	}

	if c := o.Directory(); c != nil {
		result.directory = c
		result.directoryGuard = peer.New(
			constant.DirectoryPeer,
			constant.DirectoryPeer,
			c,
			l,
			r,
		)
	}

	result.service = service.New(
		s,
		signingKey,
		keyIdentifier,
		o.Secret,
		o.Issuer,
		o.SuperUserMail,
		result.directory,
		result.directoryGuard,
	)
	result.web = web.New(result.service, result.authorization, r)

	return result
}

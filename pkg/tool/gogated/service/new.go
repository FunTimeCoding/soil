package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gogated/option"
	"github.com/funtimecoding/soil/pkg/tool/gogated/store"
)

func New(
	o *option.Option,
	l *logger.Logger,
	r face.Reporter,
) *Service {
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
	result := &Service{
		database:      d,
		store:         s,
		provider:      provider(o, s, signingKey),
		signingKey:    signingKey,
		keyIdentifier: keyIdentifier,
		issuer:        o.Issuer,
		superUserMail: o.SuperUserMail,
		authorization: authorizationClient(o).WithReporter(r),
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

	return result
}

package base

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/tool/gogated/store"
)

func (s *Server) Service() *service.Service {
	s.t.Helper()

	if s.service != nil {
		return s.service
	}

	o := s.daemonOption()
	m := s.Database()
	migrate.AutoMigrate(m)
	s.store = store.New(m)
	signingKey, keyIdentifier, e := s.store.EnsureSigningKey()
	assert.FatalOnError(s.t, e)
	assert.FatalOnError(
		s.t,
		s.store.SeedAdministrator(o.SuperUserMail, o.SuperUserPassword),
	)
	assert.FatalOnError(
		s.t,
		s.store.SeedAdminClient(
			o.AdminClientIdentifier,
			o.AdminClientSecret,
			o.Issuer,
		),
	)
	var g *peer.Guard

	if s.directory != nil {
		g = peer.New(
			constant.DirectoryPeer,
			constant.DirectoryPeer,
			s.directory,
			logger.New(context.Background()),
			s.Reporter,
		)
		s.t.Cleanup(g.Close)
	}

	s.service = service.New(
		s.store,
		signingKey,
		keyIdentifier,
		o.Secret,
		o.Issuer,
		o.SuperUserMail,
		s.directory,
		g,
	)

	return s.service
}

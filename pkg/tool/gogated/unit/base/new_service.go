package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/directory/face"
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gogated/option"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/tool/gogated/store"
	"gorm.io/gorm"
	"testing"
)

func NewService(
	t *testing.T,
	o *option.Option,
	m *gorm.DB,
	d face.DirectoryClient,
	l *logger.Logger,
	r soilFace.Reporter,
) *service.Service {
	t.Helper()
	migrate.AutoMigrate(m)
	s := store.New(m)
	signingKey, keyIdentifier, e := s.EnsureSigningKey()
	assert.FatalOnError(t, e)
	assert.FatalOnError(
		t,
		s.SeedAdministrator(o.SuperUserMail, o.SuperUserPassword),
	)
	assert.FatalOnError(
		t,
		s.SeedAdminClient(
			o.AdminClientIdentifier,
			o.AdminClientSecret,
			o.Issuer,
		),
	)
	var g *peer.Guard

	if d != nil {
		g = peer.New(constant.DirectoryPeer, constant.DirectoryPeer, d, l, r)
		t.Cleanup(g.Close)
	}

	return service.New(
		s,
		signingKey,
		keyIdentifier,
		o.Secret,
		o.Issuer,
		o.SuperUserMail,
		d,
		g,
	)
}

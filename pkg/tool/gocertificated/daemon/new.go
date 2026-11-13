package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/option"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/publish"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/web"
)

func New(
	o *option.Certificate,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	f := o.Forge()
	g := peer.New(constant.ForgePeer, "", f, l, r)
	s := service.New(
		store.New(d),
		g,
		publish.New(
			f,
			o.Project,
			o.Branch,
			o.AuthorityDirectory,
			o.SecretAuthority,
			o.SecretPath,
		),
		o.Clock,
	)

	return &Daemon{
		database: d,
		forge:    f,
		guard:    g,
		service:  s,
		web:      web.New(s, r),
	}
}

package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/option"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/publish"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
)

func New(
	o *option.Certificate,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	f := o.Forge()
	project, e := f.ResolveProject(o.Project)
	errors.PanicOnError(e)

	return &Service{
		database: d,
		store:    store.New(d),
		forge:    f,
		guard:    peer.New(constant.ForgePeer, "", f, l, r),
		publisher: publish.New(
			f,
			project,
			o.Branch,
			o.AuthorityDirectory,
			o.SecretAuthority,
			o.SecretPath,
		),
		authorization: authorizationClient(o).WithReporter(r),
	}
}

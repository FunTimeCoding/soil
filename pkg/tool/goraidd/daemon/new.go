package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/option"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/service"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/web"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/worker"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/constant"
)

func New(
	o *option.Raid,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	p := o.Parser()
	result := &Daemon{
		database:      d,
		parser:        p,
		parserGuard:   peer.New(constant.Identity.Name(), "", p, l, r),
		authorization: authorizationClient(o).WithReporter(r),
		logger:        l,
		reporter:      r,
	}
	v := service.New(
		store.New(d),
		p,
		result.parserGuard,
		o.LogCachePath,
		o.ElitePath,
		o.OutputPath,
		l,
	)
	result.service = v
	result.web = web.New(v, result.authorization, r)
	result.watcher = worker.New(
		v.HandleEvent,
		v.WatchedPaths(),
		l,
		r,
		v.SyncLogCache,
		v.EnrichExisting,
		v.Cleanup,
	)

	return result
}

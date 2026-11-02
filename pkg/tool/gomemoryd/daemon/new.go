package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web"
	queryConstant "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/worker"
)

func New(
	o *option.Memory,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	d := connection.New(l, o.LitePath)
	q := o.Query()
	result := &Daemon{
		database:  d,
		query:     q,
		guard:     peer.New(queryConstant.Identity.Name(), q.Address(), q, l, r),
		hiddenTag: o.HiddenTag,
	}
	result.service = service.New(
		store.New(d),
		q,
		result.guard,
		o.HiddenTag,
		o.ReferenceRoot,
		r,
	)
	result.web = web.New(result.service, r)
	result.clock = worker.New(
		result.service.Reconcile,
		l,
		r,
		constant.ReconcileInterval,
	)

	return result
}

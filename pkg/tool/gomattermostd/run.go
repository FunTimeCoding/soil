package gomattermostd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/service"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/worker"
	"net/http"
)

func Run(
	o *option.Mattermost,
	s face.Instrument,
) {
	r := s.Reporter()
	l := logger.New(context.Background())
	v := service.New(o, l, r)
	defer v.Close()
	errors.PanicOnError(v.Wait(context.Background()))
	b := lifecycle.New(l)

	for _, w := range v.Workers() {
		b.WithWorker(w)
	}

	if v.Subscribing() {
		b.WithWorker(worker.NewEvery(v.FlushDue, l, r, constant.FlushTick)).
			WithWorker(worker.NewEvery(v.Purge, l, r, o.PurgeInterval))
	}

	if v.Monitoring() {
		b.WithWorker(worker.NewEvery(v.PollMonitor, l, r, o.Monitor.Interval))
	}

	b.WithServer(
		server.New(
			constant.Identity,
			o.Address,
			func(m *http.ServeMux) {
				Mount(v, r, s.Recorder(), guard.New(m, o.ServiceTokens))
			},
		).WithMiddleware(web.RecoveryMiddleware(r)),
	).RunUntilSignal()
}

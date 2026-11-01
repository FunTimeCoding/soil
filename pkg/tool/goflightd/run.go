package goflightd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/option"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/service"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/worker"
	"net/http"
)

func Run(
	o *option.Flight,
	i face.Instrument,
) {
	r := i.Reporter()
	g := logger.New(context.Background())
	v := service.New(o, g)
	defer v.Close()
	b := lifecycle.New(g).
		WithWorker(worker.New(v.CaptureStream, g, r, constant.StreamPause)).
		WithWorker(
			worker.NewEvery(v.SampleBluetooth, g, r, constant.BluetoothInterval),
		).
		WithWorker(worker.NewEvery(v.Prune, g, r, constant.PruneInterval))

	if v.WirelessAvailable() {
		b.WithWorker(
			worker.NewEvery(v.SampleWireless, g, r, constant.WirelessInterval),
		)
	} else {
		g.Structured(
			"wireless collector disabled",
			"reason",
			"wdutil requires passwordless sudo",
		)
	}

	b.WithServer(
		server.New(
			constant.Identity,
			o.Address,
			func(m *http.ServeMux) {
				Mount(v, r, i.Recorder(), guard.New(m, o.ServiceTokens))
			},
		).WithMiddleware(web.RecoveryMiddleware(r)),
	).RunUntilSignal()
}

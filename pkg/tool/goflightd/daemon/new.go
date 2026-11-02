package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/option"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/service"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
	"github.com/funtimecoding/soil/pkg/worker"
)

func New(
	o *option.Flight,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d)
	v := service.New(
		store.New(d),
		l,
		o.Predicate,
		o.Retention,
		o.Clock,
		o.Command,
		o.Stream,
	)
	result := &Daemon{
		database: d,
		service:  v,
		clocks: []*worker.Worker{
			worker.New(v.CaptureStream, l, r, constant.StreamPause),
			worker.NewEvery(v.SampleBluetooth, l, r, constant.BluetoothInterval),
			worker.NewEvery(v.Prune, l, r, constant.PruneInterval),
		},
	}

	if v.WirelessAvailable() {
		result.clocks = append(
			result.clocks,
			worker.NewEvery(v.SampleWireless, l, r, constant.WirelessInterval),
		)
	} else {
		l.Structured(
			"wireless collector disabled",
			"reason",
			"wdutil requires passwordless sudo",
		)
	}

	return result
}

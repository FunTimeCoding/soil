package gobackrestd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/option"
)

func Run(
	o *option.Exporter,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	m := metric.New()
	d := daemon.New(o, l, r, m)
	defer d.Close()
	c := lifecycle.New(l)
	c.WithWorkers(d.Workers()...)
	c.WithServer(
		server.New(constant.Identity, &o.Metric, r, m.Mount).WithVerbose(
			o.Verbose,
		),
	).RunUntilSignal()
}

package gogithubd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/option"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
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

	for _, w := range d.Workers() {
		c.WithWorker(w)
	}

	c.WithServer(
		server.New(
			constant.Identity,
			o.MetricAddress,
			func(x *http.ServeMux) {
				x.Handle(webConstant.MetricsPath, m.Exporter())
			},
		).WithVerbose(o.Verbose),
	).RunUntilSignal()
}

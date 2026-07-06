package gogithubd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/option"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/worker"
)

func Run(
	o *option.Exporter,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	m := metric.New(0, o.Verbose, nil)
	lifecycle.New(
		l,
		lifecycle.WithVerbose(o.Verbose),
		lifecycle.WithWorker(m),
		lifecycle.WithWorker(
			worker.New(
				github.NewEnvironment(),
				o.Owner,
				o.Interval,
				m.Registry(),
				l,
				r,
			),
		),
	).RunUntilSignal()
}

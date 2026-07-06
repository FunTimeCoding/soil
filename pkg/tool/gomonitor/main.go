package gomonitor

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/bubbletea"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/monitor/check/collect"
	"github.com/funtimecoding/soil/pkg/tool/gomonitor/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitor/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(argument.Connect, false, "Connect to the server")
	a.Boolean(argument.Once, false, "Run once and exit")
	a.Boolean(argument.DryRun, false, "Print sources without executing")
	a.Boolean(argument.Parallel, false, "Run checks in parallel")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Once = a.GetBoolean(argument.Once)
	o.Connect = a.GetBoolean(argument.Connect)
	o.DryRun = a.GetBoolean(argument.DryRun)
	o.Parallel = a.GetBoolean(argument.Parallel)

	if o.Once {
		collect.Check(o.DryRun, o.Parallel)

		return
	}

	bubbletea.Run(monitor.New(o.Connect), true)
}

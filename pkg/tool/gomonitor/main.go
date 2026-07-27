package gomonitor

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
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
	a.Boolean(argumentConstant.Connect, false, "Connect to the server")
	a.Boolean(argumentConstant.Once, false, "Run once and exit")
	a.Boolean(
		argumentConstant.DryRun,
		false,
		"Print sources without executing",
	)
	a.Boolean(argumentConstant.Parallel, false, "Run checks in parallel")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Once = a.GetBoolean(argumentConstant.Once)
	o.Connect = a.GetBoolean(argumentConstant.Connect)
	o.DryRun = a.GetBoolean(argumentConstant.DryRun)
	o.Parallel = a.GetBoolean(argumentConstant.Parallel)

	if o.Once {
		collect.Check(o.DryRun, o.Parallel)

		return
	}

	bubbletea.Run(monitor.New(o.Connect), true)
}

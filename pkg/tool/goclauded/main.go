package goclauded

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/option"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Lite()
	a.Metric()
	a.Parse()
	o := option.New()
	o.Metric.Address = a.MetricAddress()
	o.Read(a)
	o.LitePath = a.GetString(argumentConstant.Lite)
	o.SessionExportDirectory = environment.Required(
		constant.SessionExportDirectoryEnvironment,
	)
	o.Harbor = claude.New().Base()
	o.MonitorUsage = environment.Exists(constant.MonitorUsageEnvironment) &&
		!environment.Exists(constant.UsageEndpointEnvironment)
	o.UtilizationFallback = utilization.Supported() &&
		(environment.Exists(generativeConstant.AnthropicBaseLinkEnvironment) ||
			environment.Exists(constant.UsageEndpointEnvironment))
	Run(o, s)
}

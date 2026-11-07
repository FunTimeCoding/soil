package goatlasd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/option"
	"time"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Metric()
	a.Database()
	a.Duration(argumentConstant.Interval, 5*time.Minute, constant.IntervalUsage)
	a.Duration(
		constant.RetentionArgument,
		7*24*time.Hour,
		constant.RetentionUsage,
	)
	a.Parse()
	o := option.New()
	o.Metric.Address = a.MetricAddress()
	o.Read(a)
	o.LitePath = a.GetString(argumentConstant.Lite)
	o.PostgresLocator = a.GetString(argumentConstant.Postgres)
	o.Interval = a.GetDuration(argumentConstant.Interval)
	o.Retention = a.GetDuration(constant.RetentionArgument)
	peers(o)
	Run(o, s)
}

package gosaltd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/option"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Metric()
	a.Database()
	a.Repository()
	a.CloneDirectory()
	a.CopyModes()
	a.Parse()
	o := option.New()
	o.Read(a)
	o.Metric.Address = a.MetricAddress()
	o.Repository = a.Required(argumentConstant.Repository)
	o.CloneDirectory = a.Required(argumentConstant.CloneDirectory)
	o.CopyModes = a.GetBoolean(argumentConstant.CopyModes)
	o.SaltDirectory = environment.Required(constant.SaltDirectoryEnvironment)
	o.PostgresLocator = a.GetString(argumentConstant.Postgres)
	o.LitePath = a.GetString(argumentConstant.Lite)
	Run(o, s)
}

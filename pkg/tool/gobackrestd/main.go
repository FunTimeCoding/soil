package gobackrestd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/option"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Metric()
	a.Boolean(argumentConstant.Verbose, false, "Verbose output")
	a.Parse()
	o := option.New()
	o.Metric.Address = a.MetricAddress()
	o.Verbose = a.GetBoolean(argumentConstant.Verbose)
	Run(o, s)
}

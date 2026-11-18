package gobatteryd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/option"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Metric()
	a.Boolean(argumentConstant.Verbose, false, "Verbose output")
	a.Inventory()
	a.Parse()
	o := option.New()
	o.Metric.Address = a.MetricAddress()
	o.Verbose = a.GetBoolean(argumentConstant.Verbose)
	o.Inventory = inventory.Load[instance.Instance](
		a.GetString(argumentConstant.Inventory),
	)
	Run(o, s)
}

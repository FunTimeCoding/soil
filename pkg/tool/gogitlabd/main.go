package gogitlabd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/option"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Metric()
	a.Inventory()
	a.Database()
	a.Parse()
	o := option.New()
	o.Read(a)
	o.Metric.Address = a.MetricAddress()
	o.LitePath = a.GetString(argumentConstant.Lite)
	o.PostgresLocator = a.GetString(argumentConstant.Postgres)
	o.Connector = client.NewOptional
	o.Inventory = inventory.Load[instance.Instance](
		a.GetString(argumentConstant.Inventory),
	)
	Run(o, s)
}

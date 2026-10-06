package goalertmanagerd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/inventory"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/option"
	"github.com/funtimecoding/soil/pkg/web"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.String(
		argumentConstant.Inventory,
		constant.Identity.InventoryPath(),
		"Inventory file path",
	)
	a.Parse()
	o := option.New()
	o.Address = a.Address()
	o.ServiceTokens = web.ServiceTokens()
	o.Inventory = inventory.Load(a.GetString(argumentConstant.Inventory))
	Run(o, s)
}

package goprometheusd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/option"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Inventory()
	a.Parse()
	o := option.New()
	o.Read(a)
	o.Inventory = inventory.Load[instance.Instance](
		a.GetString(argumentConstant.Inventory),
	)
	Run(o, s)
}

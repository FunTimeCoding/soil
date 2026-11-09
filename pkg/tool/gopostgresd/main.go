package gopostgresd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/option"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	Arguments(a)
	a.Parse()
	o := option.New()
	o.Read(a)
	o.Inventory = inventory.Load[instance.Instance](
		a.GetString(argumentConstant.Inventory),
	)
	o.BackupDirectory = a.GetString(constant.BackupDirectoryArgument)
	Run(o, s)
}

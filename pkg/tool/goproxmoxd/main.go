package goproxmoxd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/inventory"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	s := instrument.New(constant.Identity, version)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.String(
		argumentConstant.Inventory,
		"goproxmoxd.yaml",
		"Inventory file path",
	)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Address = a.Address()
	o.Inventory = inventory.Load(a.GetString(argumentConstant.Inventory))
	o.Version = version
	Run(o, s)
}

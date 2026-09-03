package gosourced

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/option"
	"github.com/funtimecoding/soil/pkg/web"
	"path/filepath"
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
	defaultInventory := filepath.Join(
		system.Home(),
		".local",
		"share",
		"gosourced",
		"gosourced.yaml",
	)
	a.String(
		argumentConstant.Inventory,
		defaultInventory,
		"Inventory file path",
	)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Address = a.Address()
	o.ServiceTokens = web.ServiceTokens()
	o.Version = version
	o.Inventory = inventory.Load(a.GetString(argumentConstant.Inventory))
	Run(o, s)
}

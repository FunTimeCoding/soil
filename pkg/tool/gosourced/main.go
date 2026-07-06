package gosourced

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/option"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"path/filepath"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Integer(argument.Port, web.ListenPort, web.PortUsage)
	defaultInventory := filepath.Join(
		system.Home(),
		".local",
		"share",
		"gosourced",
		"gosourced.yaml",
	)
	a.String(argument.Inventory, defaultInventory, "Inventory file path")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Port = a.RequiredInteger(argument.Port)
	o.Version = version
	o.Inventory = inventory.Load(a.GetString(argument.Inventory))
	Run(o, r)
}

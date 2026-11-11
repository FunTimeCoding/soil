package goansibled

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/option"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	Arguments(a)
	a.Parse()
	o := option.New()
	o.Read(a)
	o.Metric.Address = a.MetricAddress()
	o.Repository = a.Required(argumentConstant.Repository)
	o.CloneDirectory = a.Required(argumentConstant.CloneDirectory)
	o.CopyModes = a.GetBoolean(argumentConstant.CopyModes)
	o.AnsibleDirectory = environment.Required(
		constant.AnsibleDirectoryEnvironment,
	)
	o.Playbook = a.Slice(argumentConstant.Playbook)
	o.Site = environment.Optional(constant.AnsibleSiteEnvironment)
	o.PostgresLocator = a.GetString(argumentConstant.Postgres)
	o.LitePath = a.GetString(argumentConstant.Lite)

	if environment.Exists(gonetboxdConstant.HostEnvironment) {
		l := locator.Environment(
			gonetboxdConstant.HostEnvironment,
			gonetboxdConstant.PortEnvironment,
			gonetboxdConstant.InsecureEnvironment,
		)
		token := environment.Required(gonetboxdConstant.TokenEnvironment)
		o.Inventory = func() face.Client { return client.New(l, token) }
	}

	Run(o, s)
}

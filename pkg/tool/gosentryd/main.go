package gosentryd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	sentry "github.com/funtimecoding/soil/pkg/errors/sentry/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/option"
	web "github.com/funtimecoding/soil/pkg/web/constant"
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
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Port = a.RequiredInteger(argument.Port)
	o.Organization = environment.Required(sentry.OrganizationEnvironment)
	o.Version = version
	Run(o, r)
}

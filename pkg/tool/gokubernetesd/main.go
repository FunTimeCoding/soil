package gokubernetesd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Boolean(argument.ReadOnly, false, "Disable write operations")
	a.Lite()
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Address = a.Address()
	o.ReadOnly = a.GetBoolean(argument.ReadOnly)
	o.LitePath = a.GetString(argument.Lite)
	o.Version = version
	Run(o, r)
}

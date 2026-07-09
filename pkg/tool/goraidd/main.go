package goraidd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/option"
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
	a.Database()
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Address = a.Address()
	o.PostgresLocator = a.GetString(argument.Postgres)
	o.LitePath = a.GetString(argument.Lite)
	o.LogCachePath = "/srv/arcdps-config"
	o.ElitePath = "/srv/elite-insights"
	o.OutputPath = "/srv/gw2-report"
	Run(o, r)
}

package goflightd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/option"
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
	a.String(
		constant.PredicateFlag,
		constant.StreamPredicate,
		constant.PredicateUsage,
	)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Address = a.Address()
	o.PostgresLocator = a.GetString(argumentConstant.Postgres)
	o.LitePath = a.GetString(argumentConstant.Lite)
	o.Predicate = a.GetString(constant.PredicateFlag)
	o.Version = version
	Run(o, r)
}

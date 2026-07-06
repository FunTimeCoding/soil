package gogithubd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/option"
	"time"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(argument.Verbose, false, "Verbose output")
	a.String(
		argument.Owner,
		environment.Optional(constant.OwnerEnvironment),
		"GitHub owner",
	)
	a.Duration(argument.Interval, 5*time.Minute, "Poll interval")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Owner = a.GetString(argument.Owner)
	o.Interval = a.GetDuration(argument.Interval)
	o.Verbose = a.GetBoolean(argument.Verbose)
	Run(o, r)
}

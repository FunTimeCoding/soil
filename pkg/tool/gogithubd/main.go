package gogithubd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(argumentConstant.Verbose, false, "Verbose output")
	a.String(
		argumentConstant.Owner,
		environment.Optional(constant.OwnerEnvironment),
		"GitHub owner",
	)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Owner = a.GetString(argumentConstant.Owner)
	o.Verbose = a.GetBoolean(argumentConstant.Verbose)
	Run(o, r)
}

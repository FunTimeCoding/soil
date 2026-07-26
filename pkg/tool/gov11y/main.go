package gov11y

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/gov11y/constant"
	"github.com/funtimecoding/soil/pkg/vulnerability/check/vulnerability"
	"github.com/funtimecoding/soil/pkg/vulnerability/check/vulnerability/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.String(argumentConstant.Filter, "", "modules, comma separated")
	a.Boolean(argumentConstant.Verbose, false, "Verbose output")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Verbose = a.GetBoolean(argumentConstant.Verbose)
	o.Filter = a.Slice(argumentConstant.Filter)
	vulnerability.Check(o)
}

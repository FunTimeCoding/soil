package gov11y

import (
	"github.com/funtimecoding/soil/pkg/argument"
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
	a.String(argument.Filter, "", "modules, comma separated")
	a.Boolean(argument.Verbose, false, "Verbose output")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Verbose = a.GetBoolean(argument.Verbose)
	o.Filter = a.Slice(argument.Filter)
	vulnerability.Check(o)
}

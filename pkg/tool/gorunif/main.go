package gorunif

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/gorunif/constant"
	"github.com/funtimecoding/soil/pkg/tool/gorunif/run_if"
	runIfConstant "github.com/funtimecoding/soil/pkg/tool/gorunif/run_if/constant"
	"github.com/funtimecoding/soil/pkg/tool/gorunif/run_if/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.String(
		runIfConstant.Base,
		"",
		"Base commit (default: remote tracking branch)",
	)
	a.String(runIfConstant.Head, "HEAD", "Head commit")
	a.Boolean(
		runIfConstant.Suffix,
		false,
		"Match path as suffix instead of prefix",
	)
	a.Boolean(argument.Verbose, false, "Verbose output")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Verbose = a.GetBoolean(argument.Verbose)
	o.Base = a.GetString(runIfConstant.Base)
	o.Head = a.GetString(runIfConstant.Head)
	o.Suffix = a.GetBoolean(runIfConstant.Suffix)
	o.Pattern = a.RequiredPositional(0, "PATTERN")
	o.Execute = a.RequiredPositional(1, "EXECUTE")
	run_if.Run(o)
}

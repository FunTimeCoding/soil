package gorunif

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	git "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gorunif/constant"
	"github.com/funtimecoding/soil/pkg/tool/gorunif/run_if"
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
		argumentConstant.Base,
		"",
		"Base commit (default: remote tracking branch)",
	)
	a.String(argumentConstant.Head, git.HeadReference, "Head commit")
	a.Boolean(constant.Suffix, false, "Match path as suffix instead of prefix")
	a.Boolean(argumentConstant.Verbose, false, "Verbose output")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Directory = system.WorkDirectory()
	o.Verbose = a.GetBoolean(argumentConstant.Verbose)
	o.Base = a.GetString(argumentConstant.Base)
	o.Head = a.GetString(argumentConstant.Head)
	o.Suffix = a.GetBoolean(constant.Suffix)
	o.Pattern = a.RequiredPositional(0, "PATTERN")
	o.Execute = a.RequiredPositional(1, "EXECUTE")
	run_if.Run(o)
}

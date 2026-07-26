package gofile

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/monitor/check/file"
	"github.com/funtimecoding/soil/pkg/monitor/check/file/option"
	"github.com/funtimecoding/soil/pkg/tool/gofile/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(argumentConstant.Notation, false, "JSON output")
	a.Boolean(argumentConstant.All, false, "Include filtered in output")
	a.Boolean(argumentConstant.Verbose, false, "Verbose output")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Notation = a.GetBoolean(argumentConstant.Notation)
	o.All = a.GetBoolean(argumentConstant.All)
	o.Verbose = a.GetBoolean(argumentConstant.Verbose)
	o.Paths = a.Positionals()
	file.Check(o)
}

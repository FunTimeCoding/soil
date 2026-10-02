package golicense

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/go_mod/check/license"
	"github.com/funtimecoding/soil/pkg/go_mod/check/license/option"
	"github.com/funtimecoding/soil/pkg/tool/golicense/constant"
)

func Main(
	programVersion string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), programVersion).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(argumentConstant.Notation, false, "JSON output")
	a.Boolean(argumentConstant.All, false, "Include filtered in output")
	a.Boolean(argumentConstant.Verbose, false, "Verbose output")
	a.Parse(programVersion, gitHash, buildDate)
	o := option.New()
	o.Notation = a.GetBoolean(argumentConstant.Notation)
	o.All = a.GetBoolean(argumentConstant.All)
	o.Verbose = a.GetBoolean(argumentConstant.Verbose)
	o.Path = a.PositionalFallback(0, library.CurrentDirectory)
	license.Check(o)
}

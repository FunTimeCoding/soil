package goloc

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/measure"
	"github.com/funtimecoding/soil/pkg/measure/option"
	"github.com/funtimecoding/soil/pkg/tool/goloc/constant"
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
	a.Boolean(
		argumentConstant.File,
		false,
		"One row per file instead of per language",
	)
	a.Boolean(argumentConstant.Verbose, false, "List files no language claimed")
	a.String(
		argumentConstant.Skip,
		"",
		"Additional directory or file names to skip, comma separated",
	)
	a.String(
		argumentConstant.Language,
		"",
		"Only count these languages, comma separated",
	)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Notation = a.GetBoolean(argumentConstant.Notation)
	o.ByFile = a.GetBoolean(argumentConstant.File)
	o.Verbose = a.GetBoolean(argumentConstant.Verbose)
	o.Skips = a.Slice(argumentConstant.Skip)
	o.Languages = a.Slice(argumentConstant.Language)
	o.Paths = a.Positionals()

	if len(o.Paths) == 0 {
		o.Paths = []string{library.CurrentDirectory}
	}

	measure.Run(o)
}

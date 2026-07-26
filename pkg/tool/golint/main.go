package golint

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/tool/golint/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(
		argumentConstant.Fix,
		false,
		"Fix concerns that can be fixed",
	)
	a.Boolean(
		argumentConstant.Summary,
		false,
		"Print one line per modified file instead of per-edit detail",
	)
	a.String(
		argumentConstant.Skip,
		"",
		"Directories to skip, comma separated",
	)
	a.Boolean(
		argumentConstant.Verbose,
		false,
		"Verbose output",
	)
	a.Parse(version, gitHash, buildDate)
	lint.Lint(
		a.GetString(argumentConstant.Skip),
		a.GetBoolean(argumentConstant.Verbose),
		a.GetBoolean(argumentConstant.Fix),
		a.GetBoolean(argumentConstant.Summary),
	)
}

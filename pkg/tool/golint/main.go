package golint

import (
	"github.com/funtimecoding/soil/pkg/argument"
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
		argument.Fix,
		false,
		"Fix concerns that can be fixed",
	)
	a.Boolean(
		argument.Summary,
		false,
		"Print one line per modified file instead of per-edit detail",
	)
	a.String(
		argument.Skip,
		"",
		"Directories to skip, comma separated",
	)
	a.Boolean(
		argument.Verbose,
		false,
		"Verbose output",
	)
	a.Parse(version, gitHash, buildDate)
	lint.Lint(
		a.GetString(argument.Skip),
		a.GetBoolean(argument.Verbose),
		a.GetBoolean(argument.Fix),
		a.GetBoolean(argument.Summary),
	)
}

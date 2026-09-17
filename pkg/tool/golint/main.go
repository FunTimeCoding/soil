package golint

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/lint"
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
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
	a.Boolean(argumentConstant.Fix, false, "Fix concerns that can be fixed")
	a.Boolean(
		argumentConstant.Summary,
		false,
		"Print one line per modified file instead of per-edit detail",
	)
	a.Boolean(
		argumentConstant.Census,
		false,
		"List every unchecked markdown reference sorted by reason, with counts",
	)
	a.String(argumentConstant.Skip, "", "Directories to skip, comma separated")
	a.String(
		argumentConstant.Root,
		"",
		"Repository root to lint, discovered upward from the working directory when empty",
	)
	a.String(
		argumentConstant.Configuration,
		environment.Fallback(lintConstant.ConfigurationEnvironment, ""),
		"Repository vocabulary configuration path (private registries)",
	)
	a.Boolean(argumentConstant.Verbose, false, "Verbose output")
	a.Parse(version, gitHash, buildDate)
	root, work := lint.Root(a.GetString(argumentConstant.Root))
	scopes, e := lint.Scopes(root, work, a.Positionals())

	if e != nil {
		system.Exitf(1, "%s\n", e)
	}

	o := option.New(
		a.GetString(argumentConstant.Skip),
		a.GetBoolean(argumentConstant.Verbose),
	)
	o.Scopes = scopes
	o.Configuration = a.GetString(argumentConstant.Configuration)
	o.Census = a.GetBoolean(argumentConstant.Census)
	o.Fix = a.GetBoolean(argumentConstant.Fix)
	o.Summary = a.GetBoolean(argumentConstant.Summary)
	lint.Lint(constant.Identity.Name(), root, o)
}

package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/constant"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean("summary", false, "One line per file")
	a.Boolean("comment", false, "Include the stray_comment lint")
	a.String(
		argumentConstant.Root,
		"",
		"Repository root to analyze, discovered upward from the working directory when empty",
	)
	a.Parse(version, gitHash, buildDate)
	root, work := lint.Root(a.GetString(argumentConstant.Root))
	patterns, e := lint.Patterns(root, work, a.Positionals())

	if e != nil {
		system.Exitf(1, "%s\n", e)
	}

	lint.Header(constant.Identity.Name(), root, lint.PatternDetail(patterns))
	o := option.New()
	o.Root = root
	o.Summary = a.GetBoolean("summary")
	o.Comment = a.GetBoolean("comment")
	o.Patterns = patterns
	Run(o)
}

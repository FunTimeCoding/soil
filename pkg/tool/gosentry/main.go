package gosentry

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/check/issue"
	"github.com/funtimecoding/soil/pkg/errors/sentry/check/issue/option"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/gosentry/constant"
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
		argument.Copyable,
		false,
		"Disable OSC8 links and add a copyable link instead",
	)
	a.Boolean(argument.Notation, false, "JSON output")
	a.Boolean(argument.All, false, "Include filtered in output")
	a.Boolean(argument.Verbose, false, "Verbose output")
	a.String(
		argument.Issue,
		"",
		"Show details for a specific issue (e.g. GO-1B)",
	)
	a.Parse(version, gitHash, buildDate)

	if i := a.GetString(argument.Issue); i != "" {
		showIssue(i)

		return
	}

	p := option.New()
	p.Notation = a.GetBoolean(argument.Notation)
	p.All = a.GetBoolean(argument.All)
	p.Verbose = a.GetBoolean(argument.Verbose)
	p.Copyable = a.GetBoolean(argument.Copyable)
	issue.Check(p)
}

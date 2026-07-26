package gosentry

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
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
		argumentConstant.Copyable,
		false,
		"Disable OSC8 links and add a copyable link instead",
	)
	a.Boolean(argumentConstant.Notation, false, "JSON output")
	a.Boolean(argumentConstant.All, false, "Include filtered in output")
	a.Boolean(argumentConstant.Verbose, false, "Verbose output")
	a.String(
		argumentConstant.Issue,
		"",
		"Show details for a specific issue (e.g. GO-1B)",
	)
	a.Parse(version, gitHash, buildDate)

	if i := a.GetString(argumentConstant.Issue); i != "" {
		showIssue(i)

		return
	}

	p := option.New()
	p.Notation = a.GetBoolean(argumentConstant.Notation)
	p.All = a.GetBoolean(argumentConstant.All)
	p.Verbose = a.GetBoolean(argumentConstant.Verbose)
	p.Copyable = a.GetBoolean(argumentConstant.Copyable)
	issue.Check(p)
}

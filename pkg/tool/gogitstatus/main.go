package gogitstatus

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/git/check/status"
	"github.com/funtimecoding/soil/pkg/git/check/status/option"
	git "github.com/funtimecoding/soil/pkg/git/constant"
	monitor "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gogitstatus/constant"
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
	a.String(
		argumentConstant.Path,
		"",
		"Path to scan for git repositories. If not set, the current work directory will be used.",
	)
	a.Integer(
		argumentConstant.Depth,
		3,
		fmt.Sprintf(
			"Depth to scan for %s. Default is 3.",
			monitor.GoGitStatus.Plural,
		),
	)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Notation = a.GetBoolean(argumentConstant.Notation)
	o.All = a.GetBoolean(argumentConstant.All)
	o.Verbose = a.GetBoolean(argumentConstant.Verbose)
	o.Path = a.GetString(argumentConstant.Path)
	o.Depth = a.GetInteger(argumentConstant.Depth)

	if s := environment.Optional(git.RepositoryRootEnvironment); s != "" {
		o.Path = s
	}

	status.Check(o)
}

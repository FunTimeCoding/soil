package gogitstatus

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/git/check/status"
	"github.com/funtimecoding/soil/pkg/git/check/status/option"
	item "github.com/funtimecoding/soil/pkg/monitor/item/constant"
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
	a.Boolean(argument.Notation, false, "JSON output")
	a.Boolean(argument.All, false, "Include filtered in output")
	a.Boolean(argument.Verbose, false, "Verbose output")
	a.String(
		argument.Path,
		"",
		"Path to scan for git repositories. If not set, the current work directory will be used.",
	)
	a.Integer(
		argument.Depth,
		3,
		fmt.Sprintf(
			"Depth to scan for %s. Default is 3.",
			item.GoGitStatus.Plural,
		),
	)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Notation = a.GetBoolean(argument.Notation)
	o.All = a.GetBoolean(argument.All)
	o.Verbose = a.GetBoolean(argument.Verbose)
	o.Path = a.GetString(argument.Path)
	o.Depth = a.GetInteger(argument.Depth)

	if s := environment.Optional(status.RepositoryRootEnvironment); s != "" {
		o.Path = s
	}

	status.Check(o)
}

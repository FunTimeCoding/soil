package goversion

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/git/check/status"
	"github.com/funtimecoding/soil/pkg/go_mod/check/version"
	"github.com/funtimecoding/soil/pkg/go_mod/check/version/option"
	item "github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"github.com/funtimecoding/soil/pkg/runtime"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goversion/constant"
)

func Main(
	programVersion string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), programVersion).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(argument.Notation, false, "JSON output")
	a.Boolean(argument.All, false, "Include filtered in output")
	a.String(argument.Skip, "", "Skip matches")
	a.Integer(
		argument.Depth,
		3,
		fmt.Sprintf(
			"Depth to scan for %s. Default: 3",
			item.GoVersion.Plural,
		),
	)
	a.Parse(programVersion, gitHash, buildDate)
	o := option.New()
	o.Notation = a.GetBoolean(argument.Notation)
	o.All = a.GetBoolean(argument.All)
	o.Path = a.PositionalFallback(
		0,
		environment.Fallback(
			status.RepositoryRootEnvironment,
			library.CurrentDirectory,
		),
	)
	o.Depth = a.GetInteger(argument.Depth)

	if s := environment.Optional(version.SkipEnvironment); s != "" {
		o.Skip = split.Comma(s)
	}

	if len(o.Skip) == 0 {
		o.Skip = a.Slice(argument.Skip)
	}

	v := runtime.ExecutableVersion()

	if v == nil {
		system.Exitf(1, "could not get Go version\n")

		return
	}

	o.RuntimeVersion = v.String()
	version.Check(o)
}

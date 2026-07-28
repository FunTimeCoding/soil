package goversion

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	git "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/check/version"
	"github.com/funtimecoding/soil/pkg/go_mod/check/version/option"
	go_mod "github.com/funtimecoding/soil/pkg/go_mod/constant"
	monitor "github.com/funtimecoding/soil/pkg/monitor/constant"
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
	a.Boolean(argumentConstant.Notation, false, "JSON output")
	a.Boolean(argumentConstant.All, false, "Include filtered in output")
	a.String(argumentConstant.Skip, "", "Skip matches")
	a.Integer(
		argumentConstant.Depth,
		3,
		fmt.Sprintf(
			"Depth to scan for %s. Default: 3",
			monitor.GoVersion.Plural,
		),
	)
	a.Parse(programVersion, gitHash, buildDate)
	o := option.New()
	o.Notation = a.GetBoolean(argumentConstant.Notation)
	o.All = a.GetBoolean(argumentConstant.All)
	o.Path = a.PositionalFallback(
		0,
		environment.Fallback(
			git.RepositoryRootEnvironment,
			library.CurrentDirectory,
		),
	)
	o.Depth = a.GetInteger(argumentConstant.Depth)

	if s := environment.Optional(go_mod.VersionSkipEnvironment); s != "" {
		o.Skip = split.Comma(s)
	}

	if len(o.Skip) == 0 {
		o.Skip = a.Slice(argumentConstant.Skip)
	}

	v := runtime.ExecutableVersion()

	if v == nil {
		system.Exitf(1, "could not get Go version\n")

		return
	}

	o.RuntimeVersion = v.String()
	version.Check(o)
}

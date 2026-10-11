package gomodule

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	library "github.com/funtimecoding/soil/pkg/constant"
	git "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/check/version"
	"github.com/funtimecoding/soil/pkg/go_mod/check/version/option"
	goModule "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/runtime"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gomodule/constant"
	"github.com/spf13/cobra"
)

func versionCommand(t *terminal.Terminal) *cobra.Command {
	o := option.New()
	var skip string
	result := &cobra.Command{
		Use:   constant.VersionUsage,
		Short: "Check the Go version of every project below a path",
		Args:  cobra.MaximumNArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			o.Path = environment.Fallback(
				git.RepositoryRootEnvironment,
				library.CurrentDirectory,
			)

			if len(arguments) > 0 {
				o.Path = arguments[0]
			}

			if s := environment.Optional(goModule.VersionSkipEnvironment); s != "" {
				o.Skip = split.Comma(s)
			} else if skip != "" {
				o.Skip = split.Comma(skip)
			}

			v := runtime.ExecutableVersion()

			if v == nil {
				t.Exitf("could not get Go version\n")

				return
			}

			o.RuntimeVersion = v.String()
			version.Check(o)
		},
	}
	f := result.Flags()
	f.BoolVar(&o.Notation, argumentConstant.Notation, false, "JSON output")
	f.BoolVar(&o.All, argumentConstant.All, false, "Include filtered in output")
	f.StringVar(&skip, argumentConstant.Skip, "", "Skip matches")
	f.IntVar(
		&o.Depth,
		argumentConstant.Depth,
		constant.DefaultDepth,
		"Depth to scan for Go projects",
	)

	return result
}

package gobuild

import (
	"github.com/funtimecoding/soil/pkg/build"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gobuild/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobuild/option"
	"github.com/spf13/cobra"
	"runtime"
)

func installCommand() *cobra.Command {
	o := option.New()
	result := &cobra.Command{
		Use:   join.Space(constant.InstallCommand, "name..."),
		Short: "Build the host architecture into the binary directory",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(
			c *cobra.Command,
			arguments []string,
		) error {
			o.Names = arguments
			o.Install = true
			o.Targets = []string{
				build.SystemArchitecture(runtime.GOOS, runtime.GOARCH),
			}

			if e := Validate(o); e != nil {
				return e
			}

			c.SilenceUsage = true

			return Run(o)
		},
	}
	moduleFlags(result, o)

	return result
}

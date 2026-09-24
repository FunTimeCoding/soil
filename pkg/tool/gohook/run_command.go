package gohook

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	git "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gohook/configuration"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohook/option"
	"github.com/spf13/cobra"
)

func runCommand() *cobra.Command {
	o := option.New()
	result := &cobra.Command{
		Use:   constant.RunUsage,
		Short: "Run the jobs of one hook whose paths changed",
		Args:  cobra.MinimumNArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			o.Hook = arguments[0]
			o.Arguments = arguments[1:]
			o.Input = standardInput()
			system.ExitOnCode(Run(configuration.Load(root()), o))
		},
	}
	result.Flags().StringVar(
		&o.Base,
		argumentConstant.Base,
		"",
		"Base commit; defaults to the range the hook implies",
	)
	result.Flags().StringVar(
		&o.Head,
		argumentConstant.Head,
		git.HeadReference,
		"Head commit",
	)
	result.Flags().BoolVar(
		&o.Verbose,
		argumentConstant.Verbose,
		false,
		"Print the range, skipped jobs and executed jobs",
	)

	return result
}

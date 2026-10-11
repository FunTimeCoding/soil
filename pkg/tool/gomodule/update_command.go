package gomodule

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gomodule/constant"
	"github.com/spf13/cobra"
)

func updateCommand(t *terminal.Terminal) *cobra.Command {
	var continueOnError bool
	var exclusives []string
	var downgrades []string
	result := &cobra.Command{
		Use:   constant.UpdateUsage,
		Short: "Update direct dependencies, tidy, and sync the Go version into build files",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			update(t, exclusives, downgrades, continueOnError)
		},
	}
	f := result.Flags()
	f.BoolVar(
		&continueOnError,
		argumentConstant.Continue,
		false,
		"Continue on error",
	)
	f.StringSliceVar(
		&exclusives,
		argumentConstant.Exclusive,
		nil,
		"One or more matches to exclusively update, comma separated",
	)
	f.StringSliceVar(
		&downgrades,
		argumentConstant.Downgrade,
		nil,
		"One or more downgrades to apply after update, comma separated",
	)

	return result
}

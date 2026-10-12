package gosecret

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"github.com/spf13/cobra"
)

func manifestCommand(
	t *terminal.Terminal,
	use string,
	short string,
	mode constant.Mode,
) *cobra.Command {
	var directory string
	result := &cobra.Command{
		Use:   use,
		Short: short,
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			if e := Run(directory, mode); e != nil {
				t.Exitf("%v\n", e)
			}
		},
	}
	result.Flags().StringVar(
		&directory,
		argumentConstant.Directory,
		soil.CurrentDirectory,
		"Directory to scan for secret manifests",
	)

	return result
}

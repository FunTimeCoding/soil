package gobuild

import (
	"github.com/funtimecoding/soil/pkg/tool/gobuild/constant"
	"github.com/spf13/cobra"
)

func staleCommand() *cobra.Command {
	return &cobra.Command{
		Use:   constant.StaleCommand,
		Short: "Rebuild every stale binary on the PATH into the binary directory",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			rebuildStale()
		},
	}
}

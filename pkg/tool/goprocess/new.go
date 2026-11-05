package goprocess

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goprocess/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
	"github.com/spf13/cobra"
)

func New(
	open func(procfile string) face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{
		Use:           constant.Identity.Usage(),
		Short:         constant.Identity.Description(),
		SilenceErrors: true,
	}
	result.PersistentFlags().StringP(
		"file",
		"f",
		"Procfile",
		"path to Procfile",
	)
	result.AddCommand(checkCommand(t))
	result.AddCommand(runCommand(open, t))

	return result
}

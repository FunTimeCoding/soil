package gosalt

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosalt/constant"
	"github.com/spf13/cobra"
)

func New(t *terminal.Terminal) *cobra.Command {
	result := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	result.AddCommand(orphans(t))
	result.AddCommand(why(t))
	result.AddCommand(trace(t))

	return result
}

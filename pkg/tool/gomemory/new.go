package gomemory

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gomemory/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"github.com/spf13/cobra"
)

func New(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	x := &Context{Client: c, Terminal: t}
	result := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	result.AddCommand(profile(x))
	result.AddCommand(statistic(x))
	result.AddCommand(relations(x))

	return result
}

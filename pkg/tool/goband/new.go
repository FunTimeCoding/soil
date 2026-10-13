package goband

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goband/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/spf13/cobra"
)

func New(
	c *client.ClientWithResponses,
	t *terminal.Terminal,
) *cobra.Command {
	x := &Context{Client: c, Terminal: t}
	result := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	result.AddCommand(listInstances(x))
	result.AddCommand(status(x))
	result.AddCommand(start(x))
	result.AddCommand(shutdown(x))
	result.AddCommand(reboot(x))
	result.AddCommand(redirection(x))

	return result
}

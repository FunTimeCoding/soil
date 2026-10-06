package godirectory

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/spf13/cobra"
)

func group(
	c *client.ClientWithResponses,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{Use: "group", Short: "Manage directory groups"}
	result.AddCommand(groupList(c, t))
	result.AddCommand(groupCreate(c, t))
	result.AddCommand(groupDelete(c, t))
	result.AddCommand(groupAdd(c, t))
	result.AddCommand(groupRemove(c, t))

	return result
}

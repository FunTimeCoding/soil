package godirectory

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/spf13/cobra"
)

func user(
	c *client.ClientWithResponses,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{Use: "user", Short: "Manage directory users"}
	result.AddCommand(userList(c, t))
	result.AddCommand(userCreate(c, t))
	result.AddCommand(userDelete(c, t))
	result.AddCommand(userPassword(c, t))

	return result
}

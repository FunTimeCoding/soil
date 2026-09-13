package godirectory

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/spf13/cobra"
)

func user(c *client.ClientWithResponses) *cobra.Command {
	result := &cobra.Command{Use: "user", Short: "Manage directory users"}
	result.AddCommand(userList(c))
	result.AddCommand(userCreate(c))
	result.AddCommand(userDelete(c))
	result.AddCommand(userPassword(c))

	return result
}

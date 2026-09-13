package godirectory

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/spf13/cobra"
)

func group(c *client.ClientWithResponses) *cobra.Command {
	result := &cobra.Command{Use: "group", Short: "Manage directory groups"}
	result.AddCommand(groupList(c))
	result.AddCommand(groupCreate(c))
	result.AddCommand(groupDelete(c))
	result.AddCommand(groupAdd(c))
	result.AddCommand(groupRemove(c))

	return result
}

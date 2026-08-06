package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createPlatform(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "create-platform [name]",
		Short: "Create a platform (operating system designation)",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.CreatePlatform(arguments[0]))
		},
	}
}

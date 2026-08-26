package goalpine

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/client"
	"github.com/spf13/cobra"
)

func packages(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "packages [name]",
		Short: "List packages in the repository index",
		Args:  cobra.MaximumNArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			name := ""

			if len(arguments) > 0 {
				name = arguments[0]
			}

			fmt.Println(c.Packages(name))
		},
	}
}

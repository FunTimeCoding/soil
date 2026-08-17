package gopnsense

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	"github.com/spf13/cobra"
)

func interfaces(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "interfaces",
		Short: "List network interfaces",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.Interfaces())
		},
	}
}

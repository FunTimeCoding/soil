package gohabitica

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/client"
	"github.com/spf13/cobra"
)

func tags(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "tags",
		Short: "List tags",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.Tags())
		},
	}
}

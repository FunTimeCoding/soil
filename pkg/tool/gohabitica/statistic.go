package gohabitica

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/client"
	"github.com/spf13/cobra"
)

func statistic(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "statistic",
		Short: "Get user statistic",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.Statistic())
		},
	}
}

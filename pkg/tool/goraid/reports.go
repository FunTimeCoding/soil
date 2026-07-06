package goraid

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/raid"
	"github.com/spf13/cobra"
)

func reports(c *raid.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "reports",
		Short: "List generated reports",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.Reports())
		},
	}
}

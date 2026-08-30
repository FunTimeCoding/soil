package goraid

import (
	"github.com/funtimecoding/soil/pkg/console"
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
			console.Emit(c.Reports())
		},
	}
}

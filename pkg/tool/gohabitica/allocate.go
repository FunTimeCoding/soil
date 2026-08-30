package gohabitica

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/client"
	"github.com/spf13/cobra"
)

func allocate(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "allocate <stat>",
		Short: "Allocate a stat point: str, con, int, per",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.AllocateStat(arguments[0]))
		},
	}
}

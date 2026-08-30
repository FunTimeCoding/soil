package gohabitica

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/client"
	"github.com/spf13/cobra"
)

func equip(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "equip <key>",
		Short: "Equip a gear item by key",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.EquipGear(arguments[0]))
		},
	}
}

package gohabitica

import (
	"fmt"
	"github.com/spf13/cobra"
)

func equip(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "equip <key>",
		Short: "Equip a gear item by key",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(x.Client.EquipGear(arguments[0]))
		},
	}
}

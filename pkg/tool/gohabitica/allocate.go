package gohabitica

import (
	"fmt"
	"github.com/spf13/cobra"
)

func allocate(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "allocate <stat>",
		Short: "Allocate a stat point: str, con, int, per",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(x.Client.AllocateStat(arguments[0]))
		},
	}
}

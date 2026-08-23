package gohabitica

import (
	"fmt"
	"github.com/spf13/cobra"
)

func gear(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "gear",
		Short: "List owned gear",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(x.Client.Gear())
		},
	}
}

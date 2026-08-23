package gohabitica

import (
	"fmt"
	"github.com/spf13/cobra"
)

func statistic(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "statistic",
		Short: "Get user statistic",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(x.Client.Statistic())
		},
	}
}

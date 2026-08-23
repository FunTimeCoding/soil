package gohabitica

import (
	"fmt"
	"github.com/spf13/cobra"
)

func cron(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "cron",
		Short: "Check and run daily rollover if needed",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(x.Client.Cron())
		},
	}
}

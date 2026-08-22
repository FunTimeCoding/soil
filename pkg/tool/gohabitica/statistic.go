package gohabitica

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
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
			x.record(constant.GetStats)
		},
	}
}

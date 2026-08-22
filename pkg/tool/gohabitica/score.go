package gohabitica

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/spf13/cobra"
)

func score(x *Context) *cobra.Command {
	var direction string
	result := &cobra.Command{
		Use:   "score [identifier]",
		Short: "Score a task",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(x.Client.Score(arguments[0], direction))
			x.record(constant.ScoreTask)
		},
	}
	result.Flags().StringVar(
		&direction,
		"direction",
		"up",
		"Score direction: up or down",
	)

	return result
}

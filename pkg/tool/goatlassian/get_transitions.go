package goatlassian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func getTransitions(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "get-transitions [key]",
		Short: "List available transitions for a Jira issue",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.GetTransitions(arguments[0]))
		},
	}
}

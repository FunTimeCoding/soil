package goatlassian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func getIssue(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "get-issue [key]",
		Short: "Get a Jira issue by key",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.GetIssue(arguments[0]))
		},
	}
}

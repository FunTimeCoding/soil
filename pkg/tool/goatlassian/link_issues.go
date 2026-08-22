package goatlassian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func linkIssues(c *client.Client) *cobra.Command {
	var linkType string
	result := &cobra.Command{
		Use:   "link-issues [key] [target-key]",
		Short: "Link two Jira issues",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			c.LinkIssues(arguments[0], arguments[1], linkType)
			fmt.Printf("linked %s to %s\n", arguments[0], arguments[1])
		},
	}
	result.Flags().StringVar(
		&linkType,
		"type",
		"",
		"link type name (default Relates)",
	)

	return result
}

package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func getCreateMeta(c *client.Client) *cobra.Command {
	var project string
	var issueType string
	var expand string
	result := &cobra.Command{
		Use:   "get-create-meta",
		Short: "Get create metadata fields for a project and issue type",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.GetCreateMeta(project, issueType, expand))
		},
	}
	result.Flags().StringVar(&project, "project", "", "project key (required)")
	result.Flags().StringVar(
		&issueType,
		"type",
		"",
		"issue type name (required)",
	)
	result.Flags().StringVar(
		&expand,
		"expand",
		"",
		"comma-separated field names to expand allowed values",
	)
	errors.PanicOnError(result.MarkFlagRequired("project"))
	errors.PanicOnError(result.MarkFlagRequired("type"))

	return result
}

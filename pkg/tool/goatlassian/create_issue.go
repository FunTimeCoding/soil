package goatlassian

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func createIssue(c *client.Client) *cobra.Command {
	var project string
	var issueType string
	var description string
	var assignee string
	var labels string
	var fields string
	result := &cobra.Command{
		Use:   "create-issue [summary]",
		Short: "Create a Jira issue",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			var labelList []string

			if labels != "" {
				labelList = strings.Split(labels, ",")
			}

			var fieldMap map[string]any

			if fields != "" {
				if e := json.Unmarshal([]byte(fields), &fieldMap); e != nil {
					errors.Printf("fields must be a JSON object: %v\n", e)
					os.Exit(1)
				}
			}

			console.Emit(
				c.CreateIssue(
					project,
					issueType,
					arguments[0],
					description,
					assignee,
					labelList,
					fieldMap,
				),
			)
		},
	}
	result.Flags().StringVar(&project, "project", "", "project key (required)")
	result.Flags().StringVar(
		&issueType,
		"type",
		"",
		"issue type name (required)",
	)
	result.Flags().StringVar(&description, "description", "", "description")
	result.Flags().StringVar(&assignee, "assignee", "", "assignee name")
	result.Flags().StringVar(&labels, "labels", "", "comma-separated labels")
	result.Flags().StringVar(
		&fields,
		"fields",
		"",
		"additional fields as a JSON object",
	)
	errors.PanicOnError(result.MarkFlagRequired("project"))
	errors.PanicOnError(result.MarkFlagRequired("type"))

	return result
}

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

func updateIssue(c *client.Client) *cobra.Command {
	var summary string
	var description string
	var assignee string
	var reporter string
	var labels string
	var fields string
	var noDiff bool
	result := &cobra.Command{
		Use:   "update-issue [key]",
		Short: "Update fields on a Jira issue",
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
				c.UpdateIssue(
					arguments[0],
					summary,
					description,
					assignee,
					reporter,
					labelList,
					fieldMap,
					noDiff,
				),
			)
		},
	}
	result.Flags().StringVar(&summary, "summary", "", "new summary")
	result.Flags().StringVar(&description, "description", "", "new description")
	result.Flags().StringVar(&assignee, "assignee", "", "assignee name")
	result.Flags().StringVar(&reporter, "reporter", "", "reporter name")
	result.Flags().StringVar(&labels, "labels", "", "comma-separated labels")
	result.Flags().StringVar(
		&fields,
		"fields",
		"",
		"additional fields as a JSON object",
	)
	result.Flags().BoolVar(&noDiff, "no-diff", false, "skip the change diff")

	return result
}

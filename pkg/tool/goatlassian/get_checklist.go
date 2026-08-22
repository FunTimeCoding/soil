package goatlassian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func getChecklist(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "get-checklist [key]",
		Short: "Read an issue's checklist",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.GetChecklist(arguments[0]))
		},
	}
}

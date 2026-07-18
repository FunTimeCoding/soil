package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listVirtualJournalEntries(c *client.Client) *cobra.Command {
	var limit int32
	var offset int32
	result := &cobra.Command{
		Use:   "list-virtual-journal-entries [machine]",
		Short: "List journal entries for a virtual machine, newest first",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(
				c.ListVirtualJournalEntries(
					arguments[0],
					limit,
					offset,
				),
			)
		},
	}
	result.Flags().Int32Var(
		&limit,
		"limit",
		0,
		"maximum entries to return (optional)",
	)
	result.Flags().Int32Var(
		&offset,
		"offset",
		0,
		"entries to skip (optional)",
	)

	return result
}

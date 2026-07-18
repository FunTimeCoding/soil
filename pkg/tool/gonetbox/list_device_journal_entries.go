package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listDeviceJournalEntries(c *client.Client) *cobra.Command {
	var limit int32
	var offset int32
	result := &cobra.Command{
		Use:   "list-device-journal-entries [device]",
		Short: "List journal entries for a device, newest first",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(
				c.ListDeviceJournalEntries(
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

package goatlassian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func searchPages(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "search-pages [query]",
		Short: "Search Confluence pages using CQL or plain text",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.SearchPages(arguments[0]))
		},
	}
}

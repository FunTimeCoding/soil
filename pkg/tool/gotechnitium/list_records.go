package gotechnitium

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/technitium"
	"github.com/spf13/cobra"
)

func listRecords(c *technitium.Client) *cobra.Command {
	var all bool
	result := &cobra.Command{
		Use:   "list-records [domain]",
		Short: "List records for a zone or domain",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			for _, r := range c.MustRecords(arguments[0], all) {
				fmt.Printf(
					"%s %s %v\n",
					r.Name,
					r.Type,
					r.Payload,
				)
			}
		},
	}
	result.Flags().BoolVar(
		&all,
		"all",
		false,
		"list all records in the zone",
	)

	return result
}

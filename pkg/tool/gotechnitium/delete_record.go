package gotechnitium

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/technitium"
	"github.com/spf13/cobra"
)

func deleteRecord(c *technitium.Client) *cobra.Command {
	var recordType string
	var value string
	result := &cobra.Command{
		Use:   "delete-record [domain]",
		Short: "Delete a DNS record",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			c.MustDeleteRecord(arguments[0], recordType, value)
			fmt.Printf(
				"deleted: %s %s %s\n",
				arguments[0],
				recordType,
				value,
			)
		},
	}
	result.Flags().StringVar(
		&recordType,
		"type",
		"A",
		"record type (A, AAAA, CNAME, PTR, TXT)",
	)
	result.Flags().StringVar(
		&value,
		"value",
		"",
		"record value to delete",
	)
	errors.PanicOnError(result.MarkFlagRequired("value"))

	return result
}

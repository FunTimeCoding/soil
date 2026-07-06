package gotechnitium

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/technitium"
	"github.com/spf13/cobra"
)

func addRecord(c *technitium.Client) *cobra.Command {
	var recordType string
	var value string
	result := &cobra.Command{
		Use:   "add-record [domain]",
		Short: "Add a DNS record",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			r := c.MustAddRecord(arguments[0], recordType, value)
			fmt.Printf(
				"added: %s %s %v\n",
				r.Name,
				r.Type,
				r.Payload,
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
		"record value (IP address, domain name, or text)",
	)
	errors.PanicOnError(result.MarkFlagRequired("value"))

	return result
}

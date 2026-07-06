package gotechnitium

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/technitium"
	"github.com/spf13/cobra"
)

func createZone(c *technitium.Client) *cobra.Command {
	var zoneType string
	result := &cobra.Command{
		Use:   "create-zone [name]",
		Short: "Create a new authoritative zone",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			domain := c.MustCreateZone(arguments[0], zoneType)
			fmt.Printf("created zone: %s\n", domain)
		},
	}
	result.Flags().StringVar(
		&zoneType,
		"type",
		"Primary",
		"zone type (Primary, Secondary, Stub, Forwarder)",
	)

	return result
}

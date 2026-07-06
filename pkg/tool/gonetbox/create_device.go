package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createDevice(c *client.Client) *cobra.Command {
	var role string
	var deviceType string
	var site string
	var tenant string
	result := &cobra.Command{
		Use:   "create-device [name]",
		Short: "Create a NetBox device",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			var t *string

			if tenant != "" {
				t = &tenant
			}

			fmt.Println(
				c.CreateDevice(
					arguments[0],
					role,
					deviceType,
					site,
					t,
				),
			)
		},
	}
	result.Flags().StringVar(
		&role,
		"role",
		"",
		"device role name (required)",
	)
	result.Flags().StringVar(
		&deviceType,
		"type",
		"",
		"device type model (required)",
	)
	result.Flags().StringVar(&site, "site", "", "site name (required)")
	result.Flags().StringVar(&tenant, "tenant", "", "tenant name (optional)")
	errors.PanicOnError(result.MarkFlagRequired("role"))
	errors.PanicOnError(result.MarkFlagRequired("type"))
	errors.PanicOnError(result.MarkFlagRequired("site"))

	return result
}

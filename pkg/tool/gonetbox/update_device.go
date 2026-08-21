package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	generated "github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/spf13/cobra"
)

func updateDevice(c *client.Client) *cobra.Command {
	var location string
	var platform string
	var tenant string
	var serial string
	var description string
	result := &cobra.Command{
		Use:   "update-device [name]",
		Short: "Update device fields",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			body := generated.UpdateDeviceRequest{}

			if location != "" {
				body.Location = &location
			}

			if platform != "" {
				body.Platform = &platform
			}

			if tenant != "" {
				body.Tenant = &tenant
			}

			if serial != "" {
				body.Serial = &serial
			}

			if description != "" {
				body.Description = &description
			}

			fmt.Println(c.UpdateDevice(arguments[0], body))
		},
	}
	result.Flags().StringVar(
		&location,
		"location",
		"",
		"location name to assign (must already exist)",
	)
	result.Flags().StringVar(
		&platform,
		"platform",
		"",
		"platform name to assign (must already exist)",
	)
	result.Flags().StringVar(
		&tenant,
		"tenant",
		"",
		"tenant name to assign (must already exist)",
	)
	result.Flags().StringVar(&serial, "serial", "", "serial number")
	result.Flags().StringVar(
		&description,
		"description",
		"",
		"device description",
	)

	return result
}

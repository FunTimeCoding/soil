package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	generated "github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/spf13/cobra"
)

func updateVirtualMachine(c *client.Client) *cobra.Command {
	var platform string
	var tenant string
	var cores float32
	var memory int
	var status string
	result := &cobra.Command{
		Use:   "update-virtual-machine [name]",
		Short: "Update virtual machine fields",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			body := generated.UpdateVirtualMachineRequest{}

			if platform != "" {
				body.Platform = &platform
			}

			if tenant != "" {
				body.Tenant = &tenant
			}

			if cores > 0 {
				body.Cores = &cores
			}

			if memory > 0 {
				body.Memory = &memory
			}

			if status != "" {
				body.Status = &status
			}

			console.Emit(c.UpdateVirtualMachine(arguments[0], body))
		},
	}
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
	result.Flags().Float32Var(&cores, "cores", 0, "virtual CPU count")
	result.Flags().IntVar(&memory, "memory", 0, "memory in megabytes")
	result.Flags().StringVar(
		&status,
		"status",
		"",
		"status (active, offline, planned, staged, failed, decommissioning)",
	)

	return result
}

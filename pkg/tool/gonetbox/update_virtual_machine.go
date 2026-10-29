package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/spf13/cobra"
)

func updateVirtualMachine(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var platform string
	var role string
	var tenant string
	var cores float32
	var memory int
	var status string
	var customFields func() *map[string]any
	result := &cobra.Command{
		Use:   "update-virtual-machine [name]",
		Short: "Update virtual machine fields",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			body := client.UpdateVirtualMachineRequest{}

			if platform != "" {
				body.Platform = &platform
			}

			if role != "" {
				body.Role = &role
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

			body.CustomFields = customFields()
			t.EmitValue(
				c.UpdateVirtualMachine(context.Background(), arguments[0], body),
			)
		},
	}
	customFields = customFieldFlag(result)
	result.Flags().StringVar(
		&platform,
		"platform",
		"",
		"platform name to assign (must already exist)",
	)
	result.Flags().StringVar(
		&role,
		"role",
		"",
		"role name to assign (must already exist and allow virtual machines)",
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

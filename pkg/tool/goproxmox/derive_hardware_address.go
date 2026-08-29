package goproxmox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/spf13/cobra"
	"strconv"
)

func deriveHardwareAddress(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "derive-hardware-address [identifier]",
		Short: "Derive the MAC address a machine identifier maps to",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			identifier, e := strconv.Atoi(a[0])
			errors.PanicOnError(e)
			fmt.Println(c.Client().DeriveHardwareAddress(identifier))
		},
	}
}

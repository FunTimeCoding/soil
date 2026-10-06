package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
	"strconv"
)

func createVirtualDisk(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "create-virtual-disk [vm] [name] [size-megabytes]",
		Short: "Create a disk on a virtual machine",
		Args:  cobra.ExactArgs(3),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			size, e := strconv.Atoi(arguments[2])
			errors.PanicOnError(e)
			t.Emit(c.CreateVirtualDisk(arguments[0], arguments[1], int32(size)))
		},
	}
}

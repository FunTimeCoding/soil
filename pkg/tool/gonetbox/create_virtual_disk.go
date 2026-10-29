package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
	"strconv"
)

func createVirtualDisk(
	c face.Client,
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
			t.EmitValue(
				c.CreateVirtualDisk(
					context.Background(),
					arguments[0],
					arguments[1],
					int32(size),
				),
			)
		},
	}
}

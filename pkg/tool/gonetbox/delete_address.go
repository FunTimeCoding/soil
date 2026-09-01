package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
	"strconv"
)

func deleteAddress(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-address [identifier]",
		Short: "Delete an IP address by identifier",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])
			errors.PanicOnError(e)
			console.Emit(c.DeleteAddress(int32(identifier)))
			console.Line("address deleted")
		},
	}
}

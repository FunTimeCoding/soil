package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func addPowerCommands(
	t *terminal.Terminal,
	c face.Client,
	o *cobra.Command,
) {
	o.AddCommand(createPowerOutlet(c, t))
	o.AddCommand(createPowerPort(c, t))
	o.AddCommand(createPowerCable(c, t))
	o.AddCommand(listPowerOutlets(c, t))
}

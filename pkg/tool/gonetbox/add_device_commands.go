package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func addDeviceCommands(t *terminal.Terminal, c face.Client, o *cobra.Command) {
	o.AddCommand(listDevices(c, t))
	o.AddCommand(getDevice(c, t))
	o.AddCommand(getVirtualMachine(c, t))
	o.AddCommand(listInterfaces(c, t))
	o.AddCommand(listAddresses(c, t))
	o.AddCommand(listVirtualInterfaces(c, t))
	o.AddCommand(listVirtualAddresses(c, t))
	o.AddCommand(listVirtualDisks(c, t))
	o.AddCommand(createVirtualDisk(c, t))
	o.AddCommand(renameDevice(c, t))
	o.AddCommand(renameVirtualMachine(c, t))
	o.AddCommand(setDevicePrimaryAddress(c, t))
	o.AddCommand(setVirtualPrimaryAddress(c, t))
	o.AddCommand(updateDevice(c, t))
	o.AddCommand(updateVirtualMachine(c, t))
}

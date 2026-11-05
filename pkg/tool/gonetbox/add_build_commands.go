package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func addBuildCommands(t *terminal.Terminal, c face.Client, o *cobra.Command) {
	o.AddCommand(listWirelessNetworks(c, t))
	o.AddCommand(createWirelessNetwork(c, t))
	o.AddCommand(listSites(c, t))
	o.AddCommand(listTenants(c, t))
	o.AddCommand(listDeviceRoles(c, t))
	o.AddCommand(listManufacturers(c, t))
	o.AddCommand(listDeviceTypes(c, t))
	o.AddCommand(createSite(c, t))
	o.AddCommand(createTenant(c, t))
	o.AddCommand(createDeviceRole(c, t))
	o.AddCommand(createManufacturer(c, t))
	o.AddCommand(createDeviceType(c, t))
	o.AddCommand(createDevice(c, t))
	o.AddCommand(createInterface(c, t))
	o.AddCommand(setInterfacePhysicalAddress(c, t))
	o.AddCommand(setVirtualInterfacePhysicalAddress(c, t))
	o.AddCommand(createAddress(c, t))
	o.AddCommand(listClusterTypes(c, t))
	o.AddCommand(listClusters(c, t))
	o.AddCommand(listVirtualMachines(c, t))
	o.AddCommand(createClusterType(c, t))
	o.AddCommand(createCluster(c, t))
	o.AddCommand(createVirtualMachine(c, t))
	o.AddCommand(createVirtualInterface(c, t))
	o.AddCommand(createVirtualAddress(c, t))
}

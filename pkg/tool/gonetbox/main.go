package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/gonetbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	c := client.NewEnvironment()
	o := &cobra.Command{
		Use:     constant.Identity.Usage(),
		Short:   constant.Identity.Description(),
		Version: argument.CobraVersion(version, gitHash, buildDate),
	}
	o.AddCommand(listDevices(c))
	o.AddCommand(getDevice(c))
	o.AddCommand(listInterfaces(c))
	o.AddCommand(listAddresses(c))
	o.AddCommand(listSites(c))
	o.AddCommand(listTenants(c))
	o.AddCommand(listDeviceRoles(c))
	o.AddCommand(listManufacturers(c))
	o.AddCommand(listDeviceTypes(c))
	o.AddCommand(createSite(c))
	o.AddCommand(createTenant(c))
	o.AddCommand(createDeviceRole(c))
	o.AddCommand(createManufacturer(c))
	o.AddCommand(createDeviceType(c))
	o.AddCommand(createDevice(c))
	o.AddCommand(createInterface(c))
	o.AddCommand(createAddress(c))
	o.AddCommand(listClusterTypes(c))
	o.AddCommand(listClusters(c))
	o.AddCommand(listVirtualMachines(c))
	o.AddCommand(createClusterType(c))
	o.AddCommand(createCluster(c))
	o.AddCommand(createVirtualMachine(c))
	o.AddCommand(createVirtualInterface(c))
	o.AddCommand(createVirtualAddress(c))
	o.AddCommand(listTags(c))
	o.AddCommand(createTag(c))
	o.AddCommand(addDeviceTag(c))
	o.AddCommand(removeDeviceTag(c))
	o.AddCommand(addVirtualTag(c))
	o.AddCommand(removeVirtualTag(c))
	o.AddCommand(addDeviceJournalEntry(c))
	o.AddCommand(listDeviceJournalEntries(c))
	o.AddCommand(addVirtualJournalEntry(c))
	o.AddCommand(listVirtualJournalEntries(c))
	o.AddCommand(updateJournalEntry(c))
	o.AddCommand(deleteJournalEntry(c))
	o.AddCommand(setDeviceLabel(c))
	o.AddCommand(removeDeviceLabel(c))
	o.AddCommand(listDeviceLabels(c))
	o.AddCommand(setVirtualLabel(c))
	o.AddCommand(removeVirtualLabel(c))
	o.AddCommand(listVirtualLabels(c))
	o.AddCommand(listTunnelGroups(c))
	o.AddCommand(listTunnels(c))
	o.AddCommand(listTunnelTerminations(c))
	o.AddCommand(createTunnelGroup(c))
	o.AddCommand(createTunnel(c))
	o.AddCommand(createDeviceTunnelTermination(c))
	o.AddCommand(createVirtualTunnelTermination(c))
	o.AddCommand(listPrefixes(c))
	o.AddCommand(createPrefix(c))
	errors.PanicOnError(o.Execute())
}

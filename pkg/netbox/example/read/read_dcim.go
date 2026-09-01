package read

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox"
)

func readDCIM(
	n *netbox.Client,
	f *option.Format,
) {
	// Data Center Infrastructure Management
	for _, l := range n.MustLocations() {
		console.Format("Location: %s\n", l.Format(f))
	}

	for _, s := range n.MustSites() {
		console.Format("Site: %s\n", s.Format(f))
	}

	for _, r := range n.MustDeviceRoles() {
		console.Format("DeviceRole: %s\n", r.Format(f))
	}

	for _, p := range n.MustPlatforms() {
		console.Format("Platform: %s\n", p.Format(f))
	}

	for _, m := range n.MustManufacturers() {
		console.Format("Manufacturer: %s\n", m.Format(f))
	}

	for _, t := range n.MustDeviceTypes() {
		console.Format("DeviceType: %s\n", t.Format(f))
	}

	for _, d := range n.MustDevices() {
		console.Format("Device: %s\n", d.Format(f))

		for _, i := range n.MustDeviceInterfaces(d.Identifier) {
			console.Format("  Interface: %s\n", i.Format(f))
		}

		if false {
			// TODO: on load: panic: no value given for required property device
			//  Even if name is set
			//  But this worked yesterday..?
			for _, i := range n.MustDeviceModuleBays(d.Name) {
				console.Format("  ModuleBay: %s\n", i.Format(f))
			}
		}
	}

	for _, t := range n.MustRackTypes() {
		console.Format("RackType: %s\n", t.Format(f))
	}

	for _, r := range n.MustRackRoles() {
		console.Format("RackRole: %s\n", r.Format(f))
	}

	for _, r := range n.MustRacks() {
		console.Format("Rack: %s\n", r.Format(f))
	}

	for _, p := range n.MustCables() {
		console.Format("Cable: %s\n", p.Format(f))
	}

	for _, c := range n.MustVirtualChassis() {
		console.Format("VirtualChassis: %s\n", c.Format(f))
	}

	for _, t := range n.MustModuleTypes() {
		console.Format("ModuleType: %s\n", t.Format(f))
	}

	for _, m := range n.MustModules() {
		console.Format("Module: %s\n", m.Format(f))
	}

	for _, r := range n.MustInventoryItemRoles() {
		console.Format("InventoryItemRole: %s\n", r.Format(f))
	}

	for _, i := range n.MustInventoryItems() {
		console.Format("InventoryItem: %s\n", i.Format(f))
	}

	if false {
		// TODO: on load: panic: no value given for required property device
		//  This also worked - is something corrupt in the devices?
		for _, i := range n.MustModuleBays() {
			console.Format("ModuleBay: %s\n", i.Format(f))
		}
	}

	for _, e := range n.MustPowerFeeds() {
		console.Format("PowerFeed: %s\n", e.Format(f))
	}

	for _, p := range n.MustPowerPanels() {
		console.Format("PowerPanel: %s\n", p.Format(f))
	}

	for _, p := range n.MustConsoleServerPorts() {
		console.Format("ConsoleServerPort: %s\n", p.Format(f))
	}

	for _, p := range n.MustConsolePorts() {
		console.Format("ConsolePort: %s\n", p.Format(f))
	}

	for _, o := range n.MustPowerOutlets() {
		console.Format("PowerOutlet: %s\n", o.Format(f))
	}

	for _, p := range n.MustPowerPorts() {
		console.Format("PowerPort: %s\n", p.Format(f))
	}

	for _, p := range n.MustFrontPorts() {
		console.Format("FrontPort: %s\n", p.Format(f))
	}

	for _, p := range n.MustRearPorts() {
		console.Format("FrontPort: %s\n", p.Format(f))
	}

	for _, c := range n.MustVirtualDeviceContexts() {
		console.Format("VirtualDeviceContext: %s\n", c.Format(f))
	}

	// TODO: How to create DeviceBay? "device not compatible"
	//  And where to create DeviceBayTemplate? API only?
	for _, t := range n.MustDeviceBayTemplates() {
		console.Format("DeviceBayTemplate: %s\n", t.Format(f))
	}

	for _, p := range n.MustModuleTypeProfiles() {
		console.Format("ModuleTypeProfile: %s\n", p.Format(f))
	}
}

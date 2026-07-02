package read

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/netbox"
)

func readDCIM(
	n *netbox.Client,
	f *option.Format,
) {
	// Data Center Infrastructure Management
	for _, l := range n.MustLocations() {
		fmt.Printf("Location: %s\n", l.Format(f))
	}

	for _, s := range n.MustSites() {
		fmt.Printf("Site: %s\n", s.Format(f))
	}

	for _, r := range n.MustDeviceRoles() {
		fmt.Printf("DeviceRole: %s\n", r.Format(f))
	}

	for _, p := range n.MustPlatforms() {
		fmt.Printf("Platform: %s\n", p.Format(f))
	}

	for _, m := range n.MustManufacturers() {
		fmt.Printf("Manufacturer: %s\n", m.Format(f))
	}

	for _, t := range n.MustDeviceTypes() {
		fmt.Printf("DeviceType: %s\n", t.Format(f))
	}

	for _, d := range n.MustDevices() {
		fmt.Printf("Device: %s\n", d.Format(f))

		for _, i := range n.MustDeviceInterfaces(d.Identifier) {
			fmt.Printf("  Interface: %s\n", i.Format(f))
		}

		if false {
			// TODO: on load: panic: no value given for required property device
			//  Even if name is set
			//  But this worked yesterday..?
			for _, i := range n.MustDeviceModuleBays(d.Name) {
				fmt.Printf("  ModuleBay: %s\n", i.Format(f))
			}
		}
	}

	for _, t := range n.MustRackTypes() {
		fmt.Printf("RackType: %s\n", t.Format(f))
	}

	for _, r := range n.MustRackRoles() {
		fmt.Printf("RackRole: %s\n", r.Format(f))
	}

	for _, r := range n.MustRacks() {
		fmt.Printf("Rack: %s\n", r.Format(f))
	}

	for _, p := range n.MustCables() {
		fmt.Printf("Cable: %s\n", p.Format(f))
	}

	for _, c := range n.MustVirtualChassis() {
		fmt.Printf("VirtualChassis: %s\n", c.Format(f))
	}

	for _, t := range n.MustModuleTypes() {
		fmt.Printf("ModuleType: %s\n", t.Format(f))
	}

	for _, m := range n.MustModules() {
		fmt.Printf("Module: %s\n", m.Format(f))
	}

	for _, r := range n.MustInventoryItemRoles() {
		fmt.Printf("InventoryItemRole: %s\n", r.Format(f))
	}

	for _, i := range n.MustInventoryItems() {
		fmt.Printf("InventoryItem: %s\n", i.Format(f))
	}

	if false {
		// TODO: on load: panic: no value given for required property device
		//  This also worked - is something corrupt in the devices?
		for _, i := range n.MustModuleBays() {
			fmt.Printf("ModuleBay: %s\n", i.Format(f))
		}
	}

	for _, e := range n.MustPowerFeeds() {
		fmt.Printf("PowerFeed: %s\n", e.Format(f))
	}

	for _, p := range n.MustPowerPanels() {
		fmt.Printf("PowerPanel: %s\n", p.Format(f))
	}

	for _, p := range n.MustConsoleServerPorts() {
		fmt.Printf("ConsoleServerPort: %s\n", p.Format(f))
	}

	for _, p := range n.MustConsolePorts() {
		fmt.Printf("ConsolePort: %s\n", p.Format(f))
	}

	for _, o := range n.MustPowerOutlets() {
		fmt.Printf("PowerOutlet: %s\n", o.Format(f))
	}

	for _, p := range n.MustPowerPorts() {
		fmt.Printf("PowerPort: %s\n", p.Format(f))
	}

	for _, p := range n.MustFrontPorts() {
		fmt.Printf("FrontPort: %s\n", p.Format(f))
	}

	for _, p := range n.MustRearPorts() {
		fmt.Printf("FrontPort: %s\n", p.Format(f))
	}

	for _, c := range n.MustVirtualDeviceContexts() {
		fmt.Printf("VirtualDeviceContext: %s\n", c.Format(f))
	}

	// TODO: How to create DeviceBay? "device not compatible"
	//  And where to create DeviceBayTemplate? API only?
	for _, t := range n.MustDeviceBayTemplates() {
		fmt.Printf("DeviceBayTemplate: %s\n", t.Format(f))
	}

	for _, p := range n.MustModuleTypeProfiles() {
		fmt.Printf("ModuleTypeProfile: %s\n", p.Format(f))
	}
}

package read

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox"
)

func readDevices(n *netbox.Client, f *option.Format) {
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
			for _, i := range n.MustDeviceModuleBays(d.Name) {
				console.Format("  ModuleBay: %s\n", i.Format(f))
			}
		}
	}
}

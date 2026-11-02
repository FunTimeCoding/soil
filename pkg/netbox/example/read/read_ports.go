package read

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox"
)

func readPorts(n *netbox.Client, f *option.Format) {
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

	if false {
		for _, p := range n.MustFrontPorts() {
			console.Format("FrontPort: %s\n", p.Format(f))
		}
	}

	for _, p := range n.MustRearPorts() {
		console.Format("RearPort: %s\n", p.Format(f))
	}

	for _, c := range n.MustVirtualDeviceContexts() {
		console.Format("VirtualDeviceContext: %s\n", c.Format(f))
	}

	for _, t := range n.MustDeviceBayTemplates() {
		console.Format("DeviceBayTemplate: %s\n", t.Format(f))
	}

	for _, p := range n.MustModuleTypeProfiles() {
		console.Format("ModuleTypeProfile: %s\n", p.Format(f))
	}
}

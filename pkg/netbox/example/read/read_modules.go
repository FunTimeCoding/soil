package read

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox"
)

func readModules(n *netbox.Client, f *option.Format) {
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
		for _, i := range n.MustModuleBays() {
			console.Format("ModuleBay: %s\n", i.Format(f))
		}
	}
}

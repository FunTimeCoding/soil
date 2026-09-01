package read

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox"
)

func readTenant(
	n *netbox.Client,
	f *option.Format,
) {
	for _, g := range n.MustTenantGroups() {
		console.Format("TenantGroup: %s\n", g.Format(f))
	}

	for _, t := range n.MustTenants() {
		console.Format("Tenant: %s\n", t.Format(f))
	}

	for _, g := range n.MustContactGroups() {
		console.Format("ContactGroup: %s\n", g.Format(f))
	}

	for _, r := range n.MustContactRoles() {
		console.Format("ContactRole: %s\n", r.Format(f))
	}

	for _, c := range n.MustContacts() {
		console.Format("Contact: %s\n", c.Format(f))
	}
}

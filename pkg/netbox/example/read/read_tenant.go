package read

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox"
)

func readTenant(
	n *netbox.Client,
	f *option.Format,
) {
	for _, g := range n.MustTenantGroups() {
		fmt.Printf("TenantGroup: %s\n", g.Format(f))
	}

	for _, t := range n.MustTenants() {
		fmt.Printf("Tenant: %s\n", t.Format(f))
	}

	for _, g := range n.MustContactGroups() {
		fmt.Printf("ContactGroup: %s\n", g.Format(f))
	}

	for _, r := range n.MustContactRoles() {
		fmt.Printf("ContactRole: %s\n", r.Format(f))
	}

	for _, c := range n.MustContacts() {
		fmt.Printf("Contact: %s\n", c.Format(f))
	}
}

package read

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox"
)

func readUser(
	n *netbox.Client,
	f *option.Format,
) {
	for _, g := range n.MustUserGroups() {
		console.Format("UserGroup: %s\n", g.Format(f))
	}

	for _, g := range n.MustUsers() {
		console.Format("User: %s\n", g.Format(f))
	}
}

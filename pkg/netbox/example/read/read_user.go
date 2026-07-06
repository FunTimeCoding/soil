package read

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox"
)

func readUser(
	n *netbox.Client,
	f *option.Format,
) {
	for _, g := range n.MustUserGroups() {
		fmt.Printf("UserGroup: %s\n", g.Format(f))
	}

	for _, g := range n.MustUsers() {
		fmt.Printf("User: %s\n", g.Format(f))
	}
}

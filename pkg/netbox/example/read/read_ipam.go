package read

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/netbox"
)

func readIPAM(
	n *netbox.Client,
	f *option.Format,
) {
	// IP Address Management
	for _, t := range n.MustServiceTemplates() {
		fmt.Printf("ServiceTemplate: %s\n", t.Format(f))
	}

	for _, s := range n.MustServices() {
		fmt.Printf("Services: %s\n", s.Format(f))
	}

	if false {
		// TODO: on load: panic: json: cannot unmarshal number into Go struct field _PaginatedVLANGroupList.results.vid_ranges of type []int32
		for _, g := range n.MustVirtualNetworkGroups() {
			fmt.Printf("VirtualNetworkGroup: %s\n", g.Format(f))
		}
	}

	for _, e := range n.MustVirtualNetworks() {
		fmt.Printf("VirtualNetwork: %s\n", e.Format(f))
	}

	for _, u := range n.MustSystemNumbers() {
		fmt.Printf("SystemNumber: %s\n", u.Format(f))
	}
}

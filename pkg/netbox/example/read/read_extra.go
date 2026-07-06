package read

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox"
)

func readExtra(
	n *netbox.Client,
	f *option.Format,
) {
	for _, g := range n.MustNotificationGroups() {
		fmt.Printf("NotificationGroup: %s\n", g.Format(f))
	}

	for _, t := range n.MustTags() {
		fmt.Printf("Tag: %s\n", t.Format(f))
	}

	if false {
		// TODO: on load: panic: no value given for required property data_path
		for _, c := range n.MustConfigurationContexts() {
			fmt.Printf("ConfigContext: %s\n", c.Format(f))
		}
	}

	// TODO: Must specify either local content or a data file
	//  How, what is this for?
	for _, t := range n.MustConfigurationTemplates() {
		fmt.Printf("ConfigTemplate: %s\n", t.Format(f))
	}
}

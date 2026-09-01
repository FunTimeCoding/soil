package read

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox"
)

func readExtra(
	n *netbox.Client,
	f *option.Format,
) {
	for _, g := range n.MustNotificationGroups() {
		console.Format("NotificationGroup: %s\n", g.Format(f))
	}

	for _, t := range n.MustTags() {
		console.Format("Tag: %s\n", t.Format(f))
	}

	if false {
		// TODO: on load: panic: no value given for required property data_path
		for _, c := range n.MustConfigurationContexts() {
			console.Format("ConfigContext: %s\n", c.Format(f))
		}
	}

	// TODO: Must specify either local content or a data file
	//  How, what is this for?
	for _, t := range n.MustConfigurationTemplates() {
		console.Format("ConfigTemplate: %s\n", t.Format(f))
	}
}

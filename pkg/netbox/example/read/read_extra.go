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

	for _, b := range n.MustBookmarks() {
		console.Format("Bookmark: %s\n", b.Format(f))
	}

	if false {
		for _, c := range n.MustConfigurationContexts() {
			console.Format("ConfigContext: %s\n", c.Format(f))
		}
	}

	for _, t := range n.MustConfigurationTemplates() {
		console.Format("ConfigTemplate: %s\n", t.Format(f))
	}
}

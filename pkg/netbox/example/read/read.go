package read

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
)

func Read() {
	n := netbox.NewEnvironment()
	f := constant.Format
	readTenant(n, f)
	readDCIM(n, f)
	readIPAM(n, f)
	readVirtual(n, f)
	readUser(n, f)
	readExtra(n, f)
	readWireless(n, f)
	readTunnel(n, f)

	for _, s := range n.MustSources() {
		console.Format("DataSource: %s\n", s.Format(f))
	}

	for _, t := range n.MustExportTemplates() {
		console.Format("ExportTemplate: %s\n", t.Format(f))
	}

	for _, i := range n.MustCustomFields() {
		console.Format("CustomField: %s\n", i.Format(f))
	}

	for _, c := range n.MustCustomFieldChoices() {
		console.Format("CustomFieldChoice: %s\n", c.Format(f))
	}

	for _, l := range n.MustCustomLinks() {
		console.Format("CustomLink: %s\n", l.Format(f))
	}
}

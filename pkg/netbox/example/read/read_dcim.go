package read

import (
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox"
)

func readDCIM(
	n *netbox.Client,
	f *option.Format,
) {
	readDevices(n, f)
	readRacks(n, f)
	readModules(n, f)
	readPorts(n, f)
}

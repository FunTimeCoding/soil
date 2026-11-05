package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/opnsense"
)

func Read() {
	c := opnsense.NewEnvironment()
	defer c.Close()
	leases, e := c.Leases("")
	errors.PanicOnError(e)

	for _, l := range leases {
		console.Format(
			"%s %s %s reserved:%t\n",
			l.Address,
			l.Hostname,
			l.HardwareAddress,
			l.Reserved,
		)
	}

	rules, f := c.Rules("")
	errors.PanicOnError(f)

	for _, r := range rules {
		console.Format(
			"%s %s %s -> %s %s\n",
			r.Interface,
			r.Action,
			r.SourceNet,
			r.DestinationNet,
			r.Description,
		)
	}
}

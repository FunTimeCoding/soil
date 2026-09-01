package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/opnsense"
)

func Read() {
	c := opnsense.NewEnvironment()

	for _, l := range c.MustLeases("") {
		console.Format(
			"%s %s %s reserved:%t\n",
			l.Address,
			l.Hostname,
			l.HardwareAddress,
			l.Reserved,
		)
	}

	for _, r := range c.MustRules("") {
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

package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/opnsense"
)

func Read() {
	c := opnsense.NewEnvironment()

	for _, l := range c.MustLeases("") {
		fmt.Printf(
			"%s %s %s reserved:%t\n",
			l.Address,
			l.Hostname,
			l.HardwareAddress,
			l.Reserved,
		)
	}

	for _, r := range c.MustRules("") {
		fmt.Printf(
			"%s %s %s -> %s %s\n",
			r.Interface,
			r.Action,
			r.SourceNet,
			r.DestinationNet,
			r.Description,
		)
	}
}

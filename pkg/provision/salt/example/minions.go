package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/provision/salt"
)

func Minions() {
	c := salt.NewEnvironment()
	result, e := c.Minions()

	if e != nil {
		console.Format("error: %v\n", e)

		return
	}

	for _, m := range result {
		console.Format(
			"%s (%s %s)\n",
			m.Identifier,
			m.OperatingSystem,
			m.OperatingSystemRelease,
		)
	}
}

package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/fritz"
)

func Inventory() {
	c := fritz.NewEnvironment()
	device, e := c.Device()
	errors.PanicOnError(e)
	fmt.Printf(
		"%s %s, up %d seconds\n",
		device.Model,
		device.Software,
		device.UpTimeSeconds,
	)
	external, f := c.External()

	if f != nil {
		fmt.Printf("external IPv6: %v\n", f)
	} else {
		fmt.Printf(
			"external IPv6: %q /%d\n",
			external.Address,
			external.PrefixLength,
		)
	}

	prefix, g := c.Prefix()

	if g != nil {
		fmt.Printf("IPv6 prefix: %v\n", g)
	} else {
		fmt.Printf("IPv6 prefix: %q /%d\n", prefix.Prefix, prefix.PrefixLength)
	}
}

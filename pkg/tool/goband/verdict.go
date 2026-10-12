package goband

import "github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"

func verdict(
	v *client.Status,
	operation client.PowerRequestOperation,
	force bool,
) *client.Operation {
	for i := range v.Operations {
		o := &v.Operations[i]

		if string(o.Name) == string(operation) && o.Force == force {
			return o
		}
	}

	return nil
}

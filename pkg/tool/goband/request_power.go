package goband

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
)

func requestPower(
	x *Context,
	instance string,
	operation client.PowerRequestOperation,
	force bool,
	confirmed bool,
	description string,
) {
	v := readStatus(x, instance)

	if !confirmed {
		x.Terminal.EmitValue(
			describePower(v, operation, force, description),
			nil,
		)

		return
	}

	x.Terminal.EmitValue(
		x.Client.RequestPower(
			context.Background(),
			client.PowerRequest{
				Instance:  &v.Instance,
				Operation: operation,
				Force:     &force,
			},
		),
	)
}

package goband

import (
	"context"
	"fmt"
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
	x.Terminal.Line(fmt.Sprintf("Power state of %s: %s", v.Instance, v.Power))

	if !confirmed {
		if o := verdict(v, operation, force); o != nil && o.Refusal != nil {
			x.Terminal.Line(fmt.Sprintf("Would be refused: %s", *o.Refusal))

			return
		}

		x.Terminal.Line(description)

		return
	}

	r, e := x.Client.RequestPowerWithResponse(
		context.Background(),
		client.RequestPowerJSONRequestBody{
			Instance:  &v.Instance,
			Operation: operation,
			Force:     &force,
		},
	)

	if e != nil {
		x.Terminal.Exitf("error: %v\n", e)
	}

	if r.JSON200 == nil {
		x.Terminal.Reject(r.Status(), r.Body)

		return
	}

	x.Terminal.Line(
		fmt.Sprintf(
			"%s requested on %s",
			r.JSON200.Requested,
			r.JSON200.Instance,
		),
	)
}

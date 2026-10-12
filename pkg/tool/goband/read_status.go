package goband

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
)

func readStatus(
	x *Context,
	instance string,
) *client.Status {
	p := &client.GetStatusParams{}

	if instance != "" {
		p.Instance = &instance
	}

	r, e := x.Client.GetStatusWithResponse(context.Background(), p)

	if e != nil {
		x.Terminal.Exitf("error: %v\n", e)
	}

	if r.JSON200 == nil {
		x.Terminal.Reject(r.Status(), r.Body)
	}

	return r.JSON200
}

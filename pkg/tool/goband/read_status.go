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

	result, e := x.Client.GetStatus(context.Background(), p)

	if e != nil {
		x.Terminal.Exitln(e)
	}

	return result
}

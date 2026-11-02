package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/provision/types/change"
)

func guarded(
	c face.Client,
	g *peer.Guard,
) face.DownstreamFunction {
	return func(v *change.Change) error {
		return g.Attempt(
			context.Background(),
			func(context.Context) error {
				return c.PostTrigger(v)
			},
		)
	}
}

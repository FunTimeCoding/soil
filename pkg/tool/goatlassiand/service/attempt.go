package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
)

func attempt(
	g *peer.Guard,
	call func() error,
) error {
	return g.Attempt(
		context.Background(),
		func(context.Context) error {
			return call()
		},
	)
}

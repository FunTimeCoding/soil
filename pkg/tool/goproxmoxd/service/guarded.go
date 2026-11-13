package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/proxmox/face"
)

func guarded[T any](
	s *Service,
	instance string,
	f func(face.ProxmoxClient) (T, error),
) (T, error) {
	var result T
	p, e := s.pool.Client(instance)

	if e != nil {
		return result, e
	}

	return peer.Guarded(
		context.Background(),
		p.Guard(),
		func(context.Context) (T, error) { return f(p.Client()) },
	)
}

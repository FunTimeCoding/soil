package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chromium/face"
	"github.com/funtimecoding/soil/pkg/chromium/snapshot"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/view"
)

func (s *Service) observe(
	identifier string,
	p face.ProtocolClient,
) (*view.View, error) {
	nodes, e := peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*snapshot.Node, error) { return p.Snapshot() },
	)

	if e != nil {
		return nil, e
	}

	s.cacheSnapshot(identifier, nodes)

	return view.New(nodes, s.frames(identifier)), nil
}

package service

import (
	"github.com/funtimecoding/soil/pkg/chromium/face"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/view"
)

func (s *Service) observe(
	identifier string,
	p face.ProtocolClient,
) (*view.View, error) {
	nodes, e := guarded(s.guard, p.Snapshot)

	if e != nil {
		return nil, e
	}

	s.cacheSnapshot(identifier, nodes)

	return view.New(nodes, s.frames(identifier)), nil
}

package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gochromed/face"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/view"
)

func (s *Service) observe(
	identifier string,
	p face.Page,
) (*view.View, error) {
	nodes, e := p.Snapshot()

	if e != nil {
		return nil, e
	}

	s.cacheSnapshot(identifier, nodes)

	return view.New(nodes, s.frames(identifier)), nil
}

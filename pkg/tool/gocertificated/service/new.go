package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/publish"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
)

func New(
	s *store.Store,
	p *publish.Publisher,
) *Service {
	return &Service{store: s, publisher: p}
}

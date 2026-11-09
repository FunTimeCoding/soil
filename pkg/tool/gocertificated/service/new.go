package service

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/publish"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
	"time"
)

func New(
	s *store.Store,
	g *peer.Guard,
	p *publish.Publisher,
	clock func() time.Time,
) *Service {
	return &Service{store: s, guard: g, publisher: p, clock: clock}
}

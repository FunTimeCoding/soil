package system

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/debian"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store"
	"time"
)

func New(
	s *store.Store,
	d *debian.Service,
	clock func() time.Time,
	staleAfter time.Duration,
	trackerStaleAfter time.Duration,
) *Service {
	return &Service{
		store:             s,
		debian:            d,
		clock:             clock,
		staleAfter:        staleAfter,
		trackerStaleAfter: trackerStaleAfter,
	}
}

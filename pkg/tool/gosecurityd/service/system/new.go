package system

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/debian"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/score"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/system/matched"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store"
	"time"
)

func New(
	s *store.Store,
	d *debian.Service,
	c *score.Service,
	clock func() time.Time,
	staleAfter time.Duration,
	trackerStaleAfter time.Duration,
) *Service {
	return &Service{
		store:             s,
		debian:            d,
		score:             c,
		clock:             clock,
		staleAfter:        staleAfter,
		trackerStaleAfter: trackerStaleAfter,
		cache:             map[string]*matched.Matched{},
	}
}

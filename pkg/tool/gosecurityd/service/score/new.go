package score

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/mirror"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/score"
	"github.com/funtimecoding/soil/pkg/vulnerability/face"
	"time"
)

func New(
	s *store.Store,
	c face.NationalDatabase,
	g *peer.Guard,
	clock func() time.Time,
) *Service {
	return &Service{
		store:  s,
		client: c,
		guard:  g,
		clock:  clock,
		scores: map[string]*score.Score{},
		state:  mirror.Stub(),
	}
}

package service

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/provision/salt/face"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/runner"
)

func New(
	s *store.Store,
	r *runner.Runner,
	c face.SaltClient,
	g *peer.Guard,
) *Service {
	return &Service{store: s, runner: r, salt: c, guard: g}
}

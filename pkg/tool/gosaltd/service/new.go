package service

import (
	"github.com/funtimecoding/soil/pkg/provision/salt"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/runner"
)

func New(
	s *store.Store,
	r *runner.Runner,
	c *salt.Client,
) *Service {
	return &Service{store: s, runner: r, salt: c}
}

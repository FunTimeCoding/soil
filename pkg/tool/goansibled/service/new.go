package service

import (
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/runner"
)

func New(
	s *store.Store,
	r *runner.Runner,
) *Service {
	return &Service{store: s, runner: r}
}

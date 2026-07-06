package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
)

func New(
	s *store.Store,
	p *worker.Worker,
	r face.Reporter,
) *Server {
	return &Server{
		store:    s,
		worker:   p,
		reporter: r,
	}
}

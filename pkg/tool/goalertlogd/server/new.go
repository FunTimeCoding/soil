package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
)

func New(
	v *service.Service,
	p *worker.Worker,
	r face.Reporter,
) *Server {
	return &Server{service: v, worker: p, reporter: r}
}

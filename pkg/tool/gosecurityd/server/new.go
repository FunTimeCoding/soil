package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/debian"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/remediation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/score"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/vulnerability"
)

func New(
	v *vulnerability.Service,
	d *debian.Service,
	f *flag.Service,
	y *system.Service,
	c *score.Service,
	w *remediation.Service,
	r face.Reporter,
) *Server {
	return &Server{
		vulnerability: v,
		debian:        d,
		flag:          f,
		system:        y,
		score:         c,
		remediation:   w,
		reporter:      r,
	}
}

package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
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
	t face.Recorder,
) *Server {
	result := &Server{
		server:        server.New(constant.Identity, r).WithRecorder(t).Server(),
		vulnerability: v,
		debian:        d,
		flag:          f,
		system:        y,
		score:         c,
		remediation:   w,
		reporter:      r,
	}
	result.register()

	return result
}

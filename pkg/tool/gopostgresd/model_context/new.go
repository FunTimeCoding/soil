package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service"
)

func New(
	s *service.Service,
	r face.Reporter,
	t face.Recorder,
) *Server {
	result := &Server{
		server: server.New(constant.Identity).WithInstructions(
			instance_tool.Instructions(
				constant.Identity,
				len(s.Inventory().Instances),
			),
		).WithRecorder(t).WithSessionEnd(s.Selection().Clear).Server(),
		service:   s,
		instances: instance_tool.New(s.Inventory(), s.Selection()),
		reporter:  r,
	}
	result.register()

	return result
}

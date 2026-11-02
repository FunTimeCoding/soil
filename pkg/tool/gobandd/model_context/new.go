package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/service"
)

func New(
	s *service.Service,
	i *instance.Inventory,
	n *selection.Selection,
	r face.Reporter,
	t face.Recorder,
) *Server {
	result := &Server{
		server: server.New(constant.Identity).WithInstructions(
			instance_tool.Instructions(constant.Identity, len(i.Instances)),
		).WithRecorder(t).WithSessionEnd(n.Clear).Server(),
		service:   s,
		instances: instance_tool.New(i, n, r),
		reporter:  r,
	}
	result.register()

	return result
}

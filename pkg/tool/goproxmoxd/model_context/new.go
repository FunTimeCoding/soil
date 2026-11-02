package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

func New(
	v *service.Service,
	i *instance.Inventory,
	l *selection.Selection,
	r face.Reporter,
	t face.Recorder,
) *Server {
	result := &Server{
		server: server.New(constant.Identity).WithInstructions(
			instance_tool.Instructions(constant.Identity, len(i.Instances)),
		).WithRecorder(t).WithSessionEnd(l.Clear).Server(),
		service:   v,
		instances: instance_tool.New(i, l, r),
		reporter:  r,
	}
	result.register()

	return result
}

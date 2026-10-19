package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
)

func New(
	v *service.Service,
	r face.Reporter,
	t face.Recorder,
) *Server {
	result := &Server{
		server: server.New(constant.Identity).WithInstructions(
			instance_tool.Instructions(
				constant.Identity,
				len(v.Inventory().Instances),
			),
		).WithRecorder(t).WithSessionEnd(v.Selection().Clear).Server(),
		service:   v,
		instances: instance_tool.New(v.Inventory(), v.Selection(), r),
		reporter:  r,
	}
	result.register()

	return result
}

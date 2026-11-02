package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
)

func New(
	v *service.Service,
	i *instance.Inventory,
	n *selection.Selection,
	r face.Reporter,
	t face.Recorder,
) *Server {
	result := &Server{
		server: server.New(constant.Identity).WithInstructions(
			instance_tool.Instructions(constant.Identity, len(i.Instances)),
		).WithRecorder(t).WithSessionEnd(n.Clear).Server(),
		service:   v,
		instances: instance_tool.New(i, n, r),
		reporter:  r,
	}
	result.register()

	return result
}

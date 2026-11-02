package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/service"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"
)

func New(
	v *service.Service,
	i *instance.Inventory,
	n *selection.Selection,
	r face.Reporter,
	t face.Recorder,
) *Server {
	result := &Server{
		server: server.New(constant.Identity, r).WithInstructions(
			instance_tool.Instructions(constant.Identity, len(i.Instances)),
		).WithRecorder(t).WithSessionEnd(n.Clear).Server(),
		service:   v,
		instances: instance_tool.New(i, n, r),
		reporter:  r,
	}
	result.register()

	return result
}

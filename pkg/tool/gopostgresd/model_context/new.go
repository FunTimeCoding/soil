package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/backup"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/catalog"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/query"
)

func New(
	s *catalog.Service,
	q *query.Service,
	b *backup.Service,
	i *instance.Inventory,
	n *selection.Selection,
	r face.Reporter,
	t face.Recorder,
) *Server {
	result := &Server{
		server: server.New(constant.Identity, r).WithInstructions(
			instance_tool.Instructions(constant.Identity, len(i.Instances)),
		).WithRecorder(t).WithSessionEnd(n.Clear).Server(),
		catalog:   s,
		sql:       q,
		backup:    b,
		instances: instance_tool.New(i, n, r),
		reporter:  r,
	}
	result.register()

	return result
}

package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	proxFace "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
)

func New(
	v proxFace.Service,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithInstructions(
			constant.Identity.RenderInstructions(
				map[string]bool{
					constant.MultiInstance: len(v.Instances()) > 1,
				},
			),
		).WithRecorder(t).Server(),
		service:  v,
		reporter: r,
	}
	result.register()

	return result
}

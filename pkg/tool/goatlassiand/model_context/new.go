package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	atlassianFace "github.com/funtimecoding/soil/pkg/tool/goatlassiand/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/service"
)

func New(
	j atlassianFace.JiraSource,
	c atlassianFace.ConfluenceSource,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithRecorder(t).Server(),
		jira:       j,
		confluence: c,
		service:    service.New(c),
		reporter:   r,
	}
	result.register()

	return result
}

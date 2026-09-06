package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	atlassianFace "github.com/funtimecoding/soil/pkg/tool/goatlassiand/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/service"
)

func New(
	l atlassianFace.JiraSource,
	c atlassianFace.ConfluenceSource,
	r face.Reporter,
) *Server {
	return &Server{
		jira:       l,
		confluence: c,
		service:    service.New(l, c),
		reporter:   r,
	}
}

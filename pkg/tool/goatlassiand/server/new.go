package server

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/jira"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/service"
)

func New(
	l *jira.Client,
	c *confluence.Client,
	r face.Reporter,
) *Server {
	return &Server{
		jira:       l,
		confluence: c,
		service:    service.New(l, c),
		reporter:   r,
	}
}

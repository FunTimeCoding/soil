package server

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/jira"
	"github.com/funtimecoding/soil/pkg/face"
)

func New(
	l *jira.Client,
	c *confluence.Client,
	r face.Reporter,
) *Server {
	return &Server{
		jira:       l,
		confluence: c,
		reporter:   r,
	}
}

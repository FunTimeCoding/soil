package server

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/jira"
	"github.com/funtimecoding/soil/pkg/face"
)

type Server struct {
	jira       *jira.Client
	confluence *confluence.Client
	reporter   face.Reporter
}

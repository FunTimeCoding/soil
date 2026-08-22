package server

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/jira"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/service"
)

type Server struct {
	jira       *jira.Client
	confluence *confluence.Client
	service    *service.Service
	reporter   face.Reporter
}

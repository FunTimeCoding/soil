package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	atlassianFace "github.com/funtimecoding/soil/pkg/tool/goatlassiand/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/service"
)

type Server struct {
	jira       atlassianFace.JiraSource
	confluence atlassianFace.ConfluenceSource
	service    *service.Service
	reporter   face.Reporter
}

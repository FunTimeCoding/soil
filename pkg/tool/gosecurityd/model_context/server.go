package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/debian"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/remediation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/score"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/vulnerability"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server        *server.MCPServer
	vulnerability *vulnerability.Service
	debian        *debian.Service
	flag          *flag.Service
	system        *system.Service
	score         *score.Service
	remediation   *remediation.Service
	reporter      face.Reporter
}

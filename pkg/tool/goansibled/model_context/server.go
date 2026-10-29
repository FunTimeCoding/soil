package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	ansible "github.com/funtimecoding/soil/pkg/tool/goansibled/face"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	service  ansible.Ansible
	reporter face.Reporter
}

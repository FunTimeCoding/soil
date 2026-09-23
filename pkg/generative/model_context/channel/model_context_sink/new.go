package model_context_sink

import "github.com/mark3labs/mcp-go/server"

func New(s *server.MCPServer) *Sink {
	return &Sink{server: s}
}

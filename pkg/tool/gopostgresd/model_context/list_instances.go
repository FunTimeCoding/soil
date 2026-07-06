package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) listInstances(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListInstances,
) (*mcp.CallToolResult, error) {
	var active string

	if session := server.ClientSessionFromContext(x); session != nil {
		active, _ = s.service.ActiveInstance(session.SessionID())
	}

	type entry struct {
		Name     string `json:"name"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Database string `json:"database"`
		Active   bool   `json:"active"`
	}
	var result []entry

	for _, i := range s.service.Instances() {
		result = append(
			result,
			entry{
				Name:     i.Name,
				Host:     i.Host,
				Port:     i.Port,
				Database: i.Database,
				Active:   i.Name == active,
			},
		)
	}

	return response.SuccessAny(result)
}

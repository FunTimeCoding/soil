package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	proxResponse "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListStorages(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListStorages,
) (*mcp.CallToolResult, error) {
	if a.Node == "" {
		return response.Fail("node is required")
	}

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	storages, e := s.service.ListStorages(c, a.Node)

	if e != nil {
		return s.captureDetail(e)
	}

	rows := make([]proxResponse.Storage, len(storages))

	for i, v := range storages {
		rows[i] = proxResponse.Storage{
			Name:    v.Name,
			Type:    v.Type,
			Content: v.Content,
			Enabled: v.Enabled == 1,
			Shared:  v.Shared == 1,
			Active:  v.Active == 1,
			Avail:   v.Avail,
			Used:    v.Used,
			Total:   v.Total,
		}
	}

	return response.SuccessAny(rows)
}

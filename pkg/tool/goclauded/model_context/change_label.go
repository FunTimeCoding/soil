package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) changeLabel(
	sessionIdentifier string,
	name string,
	target string,
	key string,
	q mcp.CallToolRequest,
) (string, error) {
	if v := q.GetString(constant.LabelAdd, ""); v != "" {
		return s.label.AddLabel(sessionIdentifier, name, target, key, v)
	}

	if v := q.GetString(constant.LabelRemove, ""); v != "" {
		return s.label.RemoveLabel(sessionIdentifier, name, target, key, v)
	}

	return s.label.SetLabel(
		sessionIdentifier,
		name,
		target,
		key,
		q.GetString(constant.Value, ""),
	)
}

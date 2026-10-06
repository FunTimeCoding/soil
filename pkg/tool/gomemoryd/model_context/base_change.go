package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func baseChange(q mcp.CallToolRequest) *string {
	if q.GetBool(constant.ClearBase, false) {
		return new("")
	}

	if value := q.GetString(constant.Base, ""); value != "" {
		return &value
	}

	return nil
}

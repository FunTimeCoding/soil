package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func selectionOptions() []mcp.ToolOption {
	return []mcp.ToolOption{
		mcp.WithString(
			constant.ShowField,
			mcp.Description(
				"Coverage to narrow the systems to: supported, tracked or untracked; empty selects every coverage",
			),
		),
		mcp.WithObject(
			constant.LabelsArgument,
			mcp.Description(
				"Labels the systems must carry, one value per key, e.g. {\"exposure\": \"internet\", \"group\": \"red\"}",
			),
			mcp.AdditionalProperties(map[string]any{"type": "string"}),
		),
		mcp.WithString(
			constant.SearchField,
			mcp.Description("Part of the host name, ignoring case"),
		),
	}
}

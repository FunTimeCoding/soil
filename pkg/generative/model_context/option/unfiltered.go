package option

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func Unfiltered() mcp.ToolOption {
	return mcp.WithBoolean(
		constant.ParameterUnfiltered,
		mcp.Description(
			"Keep the fields normally stripped as noise, such as managedFields and last-applied-configuration (default false)",
		),
	)
}

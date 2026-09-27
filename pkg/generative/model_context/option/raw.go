package option

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func Raw() mcp.ToolOption {
	return mcp.WithBoolean(
		constant.ParameterRaw,
		mcp.Description(
			"Include the upstream object as received, which is large (default false)",
		),
	)
}

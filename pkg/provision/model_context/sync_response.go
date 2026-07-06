package model_context

import (
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/mark3labs/mcp-go/mcp"
)

func SyncResponse(
	result *runner.SyncResult,
	e error,
) (*mcp.CallToolResult, error) {
	if e != nil {
		return response.Fail(e.Error())
	}

	if result.Error != nil {
		return response.Fail(result.Error.Error())
	}

	if !result.Changed {
		return response.Success("already up to date")
	}

	if result.Diff == "" {
		return response.Success("synced, no relevant changes")
	}

	return response.SuccessAny(result)
}

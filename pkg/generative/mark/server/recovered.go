package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func recovered(r face.Reporter) server.ToolHandlerMiddleware {
	return func(next server.ToolHandlerFunc) server.ToolHandlerFunc {
		return func(
			x context.Context,
			q mcp.CallToolRequest,
		) (result *mcp.CallToolResult, e error) {
			defer func() {
				if v := recover(); v != nil {
					result, e = response.RecoverFail(r, v, q.Params.Name)
				}
			}()

			return next(x, q)
		}
	}
}

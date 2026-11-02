package server

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func recoveredResource(r face.Reporter) server.ResourceHandlerMiddleware {
	return func(next server.ResourceHandlerFunc) server.ResourceHandlerFunc {
		return func(
			x context.Context,
			q mcp.ReadResourceRequest,
		) (result []mcp.ResourceContents, e error) {
			defer func() {
				if v := recover(); v != nil {
					result = nil
					e = fmt.Errorf(
						"%s panicked, event %s",
						q.Params.URI,
						r.Recover(v),
					)
				}
			}()

			return next(x, q)
		}
	}
}

package unit

import (
	"bytes"
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/request"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"log/slog"
)

func hookLog(
	m mcp.MCPMethod,
	message any,
) string {
	var b bytes.Buffer
	h := &server.Hooks{}
	request.Hooks(h, slog.New(slog.NewTextHandler(&b, nil)), false)

	for _, f := range h.OnBeforeAny {
		f(context.Background(), 1, m, message)
	}

	return b.String()
}

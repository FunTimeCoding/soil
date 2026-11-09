package request

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"log/slog"
)

func Hooks(
	h *server.Hooks,
	l *slog.Logger,
	verbose bool,
) {
	h.AddBeforeAny(
		func(
			_ context.Context,
			_ any,
			m mcp.MCPMethod,
			a any,
		) {
			if q, okay := a.(*mcp.CallToolRequest); okay {
				l.Info(
					fmt.Sprintf("%s %s", q.Method, q.Params.Name),
					"arguments",
					q.Params.Arguments,
				)

				return
			}

			l.Info(string(m))
		},
	)

	if verbose {
		h.AddBeforeInitialize(
			func(
				_ context.Context,
				identifier any,
				q *mcp.InitializeRequest,
			) {
				l.Info(
					"Before initialize",
					constant.MarkIdentifierKey,
					identifier,
					"request",
					q,
				)
			},
		)
		h.AddOnRegisterSession(
			func(
				_ context.Context,
				s server.ClientSession,
			) {
				l.Info("Session registered", "session", s.SessionID())
			},
		)
	}
}

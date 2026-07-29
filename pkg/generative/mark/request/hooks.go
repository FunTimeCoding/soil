package request

import (
	"context"
	"errors"
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
			identifier any,
			m mcp.MCPMethod,
			a any,
		) {
			if q, okay := a.(*mcp.InitializeRequest); okay {
				l.Info(q.Method)

				return
			}

			if q, okay := a.(*mcp.ListToolsRequest); okay {
				l.Info(q.Method)

				return
			}

			if q, okay := a.(*mcp.CallToolRequest); okay {
				l.Info(
					fmt.Sprintf(
						"%s %s",
						q.Method,
						q.Params.Name,
					),
					"arguments",
					q.Params.Arguments,
				)

				return
			}

			if q, okay := a.(*mcp.PaginatedRequest); okay {
				l.Info(
					fmt.Sprintf("PaginatedRequest: %s", q.Method),
				)
				LogParameter(l, &q.Request)

				return
			}

			if q, okay := a.(*mcp.Request); okay {
				l.Info(fmt.Sprintf("Request: %s", q.Method))
				LogParameter(l, q)

				return
			}

			l.Error(
				fmt.Sprintf("Unexpected type: %T", a),
				constant.MarkIdentifierKey,
				identifier,
				"method",
				m,
				"value",
				a,
			)
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
				l.Info(
					"Session registered",
					"session",
					s.SessionID(),
				)
			},
		)
		h.AddOnRequestInitialization(
			func(
				_ context.Context,
				identifier any,
				a any,
			) error {
				if q, okay := a.(*mcp.ListToolsRequest); okay {
					if false {
						l.Info(
							"ListToolsRequest",
							constant.MarkIdentifierKey,
							identifier,
							"request",
							q,
						)

						return errors.New("unauthorized")
					}
				}

				return nil
			},
		)
	}
}

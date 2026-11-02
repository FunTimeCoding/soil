package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/turn"
	"github.com/mark3labs/mcp-go/server"
)

func turnContext(x context.Context) context.Context {
	if c := server.ClientSessionFromContext(x); c != nil {
		return turn.WithSession(x, c.SessionID())
	}

	return x
}

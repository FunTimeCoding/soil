package instance_tool

import (
	"context"
	"github.com/mark3labs/mcp-go/server"
)

func (t *Tool[T, P]) ActiveName(x context.Context) string {
	session := server.ClientSessionFromContext(x)

	if session == nil {
		return ""
	}

	result, _ := t.selection.Active(session.SessionID())

	return result
}

package selection

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/mark3labs/mcp-go/server"
)

func identifier(x context.Context) (string, error) {
	session := server.ClientSessionFromContext(x)

	if session == nil {
		return "", errors.New(constant.NoSession)
	}

	return session.SessionID(), nil
}

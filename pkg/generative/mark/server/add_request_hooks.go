package server

import (
	"github.com/funtimecoding/soil/pkg/generative/mark/request"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/mark3labs/mcp-go/server"
)

func addRequestHooks(
	h *server.Hooks,
	l *logger.Logger,
	verbose bool,
) {
	request.Hooks(h, l.Slog(), verbose)
}

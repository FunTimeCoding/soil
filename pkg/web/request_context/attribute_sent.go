package request_context

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"log/slog"
)

func (c *Context) AttributeSent(status int) []slog.Attr {
	result := c.Attribute()
	result = append(result, slog.Int(constant.TelemetryStatus, status))

	return result
}

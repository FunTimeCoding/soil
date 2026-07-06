package request_context

import (
	"github.com/funtimecoding/soil/pkg/web/telemetry/constant"
	"log/slog"
)

func (c *Context) AttributeSent(status int) []slog.Attr {
	result := c.Attribute()
	result = append(result, slog.Int(constant.Status, status))

	return result
}

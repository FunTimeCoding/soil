package request_context

import (
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"log/slog"
	"strings"
)

func (c *Context) WebhookAttribute() []slog.Attr {
	b := c.Body()
	result := append(
		c.Attribute(),
		slog.Int(constant.TelemetryBodySize, len(b)),
		slog.String(constant.TelemetryBody, b),
	)

	for k, v := range c.request.Header {
		result = append(
			result,
			slog.Any(
				key_value.Dot(
					constant.TelemetryHeaderPrefix,
					strings.ToLower(k),
				),
				v,
			),
		)
	}

	return result
}

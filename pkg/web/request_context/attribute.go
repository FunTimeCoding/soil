package request_context

import (
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"log/slog"
	"strings"
)

func (c *Context) Attribute() []slog.Attr {
	b := c.Body()
	result := []slog.Attr{
		slog.String(constant.TelemetryRequestMethod, c.request.Method),
		slog.String(constant.TelemetryPath, c.request.URL.Path),
		slog.String(constant.TelemetryScheme, c.Scheme()),
		slog.String(constant.TelemetryQuery, c.request.URL.RawQuery),
		slog.String(constant.TelemetryRoute, c.request.Pattern),
		slog.String(constant.TelemetryClient, c.ClientAddress()),
		slog.String(constant.TelemetryPeer, c.request.RemoteAddr),
		slog.String(constant.TelemetryProtocol, c.ProtocolVersion()),
		slog.String(constant.TelemetryServer, c.request.Host),
		slog.String(constant.TelemetryUserAgent, c.request.UserAgent()),
		slog.Int(constant.TelemetryBodySize, len(b)),
		slog.String(constant.TelemetryBody, b),
	}

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

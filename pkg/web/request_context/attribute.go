package request_context

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"log/slog"
)

func (c *Context) Attribute() []slog.Attr {
	return []slog.Attr{
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
	}
}

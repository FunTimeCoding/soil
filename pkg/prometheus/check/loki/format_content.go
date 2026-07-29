package loki

import (
	"github.com/funtimecoding/soil/pkg/console/status/option"
	prometheus "github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/message"
	"github.com/funtimecoding/soil/pkg/strings/join"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"strings"
)

func formatContent(
	v *message.Message,
	f *option.Format,
) string {
	route := v.Value(web.TelemetryRoute)
	body := strings.TrimSpace(v.Value(web.TelemetryBody))

	if route != "" && body != "" {
		return join.Empty(formatRoute(route, f), " ", body)
	}

	if route != "" {
		return formatRoute(route, f)
	}

	if body != "" {
		return body
	}

	if m := v.Value(prometheus.SlogMessage); m != "" {
		return m
	}

	if v.Text != "" {
		return strings.TrimSpace(v.Text)
	}

	return ""
}

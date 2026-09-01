package loki

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/message"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func printBody(messages []*message.Message) {
	for _, v := range messages {
		body := v.Value(constant.TelemetryBody)

		if body != "" {
			console.Line(body)
		}
	}
}

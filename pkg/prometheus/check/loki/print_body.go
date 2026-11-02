package loki

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/message"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
)

func printBody(messages []*message.Message) {
	for _, v := range messages {
		body := v.Value(webConstant.TelemetryBody)

		if body != "" {
			console.Line(body)
		}
	}
}

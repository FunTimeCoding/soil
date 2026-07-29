package loki

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/message"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func printBody(messages []*message.Message) {
	for _, v := range messages {
		body := v.Value(constant.TelemetryBody)

		if body != "" {
			fmt.Println(body)
		}
	}
}

package loki

import (
	"github.com/funtimecoding/soil/pkg/prometheus/loki/message"
	"time"
)

func latest(messages []*message.Message) time.Time {
	var result time.Time

	for _, v := range messages {
		if v.Time.After(result) {
			result = v.Time
		}
	}

	return result
}

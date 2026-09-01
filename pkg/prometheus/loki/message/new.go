package message

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/response"
	"github.com/funtimecoding/soil/pkg/time"
)

func New(
	e []string,
	r *response.Stream,
) *Message {
	if false {
		console.Format("Value: %s\n", r.Stream)
		console.Format("  Timestamp: %s\n", e[0])
		console.Format("  Line: %s\n", e[1])
	}

	var messageType string
	result := &Message{
		Time:      time.FromUnixNanoString(e[0]),
		Stream:    r.Stream,
		RawStream: r,
	}

	if v, okay := parseNotation(e[1]); okay {
		messageType = constant.LokiNotationType
		result.Values = v
	} else {
		messageType = constant.LokiTextType
		result.Text = e[1]
	}

	result.Type = messageType

	return result
}

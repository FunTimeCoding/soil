package message

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/response"
	"github.com/funtimecoding/soil/pkg/time"
)

func New(
	e []string,
	r *response.Stream,
) *Message {
	if false {
		fmt.Printf("Value: %s\n", r.Stream)
		fmt.Printf("  Timestamp: %s\n", e[0])
		fmt.Printf("  Line: %s\n", e[1])
	}

	var messageType string
	result := &Message{
		Time:      time.FromUnixNanoString(e[0]),
		Stream:    r.Stream,
		RawStream: r,
	}

	if v, okay := parseNotation(e[1]); okay {
		messageType = NotationType
		result.Values = v
	} else {
		messageType = TextType
		result.Text = e[1]
	}

	result.Type = messageType

	return result
}

package message

import (
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/response"
	"time"
)

type Message struct {
	Time      time.Time
	Stream    string
	Type      string
	Text      string
	Values    map[string]any
	RawStream *response.Stream
}

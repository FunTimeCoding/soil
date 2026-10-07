package stream_payload

import "github.com/funtimecoding/soil/pkg/prometheus/loki/basic/stream"

func New(streams ...*stream.Stream) *Payload {
	return &Payload{Streams: streams}
}

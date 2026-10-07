package stream_payload

import "github.com/funtimecoding/soil/pkg/prometheus/loki/basic/stream"

type Payload struct {
	Streams []*stream.Stream `json:"streams"`
}

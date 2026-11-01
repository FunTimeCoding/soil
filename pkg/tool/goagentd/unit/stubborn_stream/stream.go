package stubborn_stream

import "github.com/funtimecoding/soil/pkg/system/run/stream/mock_stream"

type Stream struct {
	*mock_stream.Stream
	refusal error
}

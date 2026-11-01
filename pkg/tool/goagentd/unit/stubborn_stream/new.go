package stubborn_stream

import "github.com/funtimecoding/soil/pkg/system/run/stream/mock_stream"

func New(refusal error) *Stream {
	return &Stream{Stream: mock_stream.New(nil, true, nil), refusal: refusal}
}

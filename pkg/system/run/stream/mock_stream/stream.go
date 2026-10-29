package mock_stream

import (
	"io"
	"sync"
)

type Stream struct {
	mutex  sync.Mutex
	reader *io.PipeReader
	writer *io.PipeWriter
	ending error
	killed int
}

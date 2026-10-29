package mock_stream

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Stream) Close() {
	errors.PanicOnError(s.reader.Close())
}

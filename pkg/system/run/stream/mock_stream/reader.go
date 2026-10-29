package mock_stream

import "io"

func (s *Stream) Reader() io.ReadCloser {
	return s.reader
}

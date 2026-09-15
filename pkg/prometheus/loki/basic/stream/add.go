package stream

import (
	"strconv"
	"time"
)

func (s *Stream) Add(
	t time.Time,
	line string,
) {
	s.Values = append(
		s.Values,
		[]string{strconv.FormatInt(t.UnixNano(), 10), line},
	)
}

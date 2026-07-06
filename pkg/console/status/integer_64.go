package status

import "github.com/funtimecoding/soil/pkg/integers64"

func (s *Status) Integer64(v ...int64) *Status {
	for _, e := range v {
		s.bubbles = append(s.bubbles, integers64.ToString(e))
	}

	return s
}

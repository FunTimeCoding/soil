package status

import "github.com/funtimecoding/soil/pkg/integers32"

func (s *Status) Integer32(v ...int32) *Status {
	for _, e := range v {
		s.bubbles = append(s.bubbles, integers32.ToString(e))
	}

	return s
}

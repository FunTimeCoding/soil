package status

import "github.com/funtimecoding/soil/pkg/integers"

func (s *Status) Integer(v ...int) *Status {
	for _, e := range v {
		s.bubbles = append(s.bubbles, integers.ToString(e))
	}

	return s
}

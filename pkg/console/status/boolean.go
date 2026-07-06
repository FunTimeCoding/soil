package status

import "github.com/funtimecoding/soil/pkg/booleans"

func (s *Status) Boolean(v ...bool) *Status {
	for _, e := range v {
		s.bubbles = append(s.bubbles, booleans.ToString(e))
	}

	return s
}

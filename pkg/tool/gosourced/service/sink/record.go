package sink

import "github.com/funtimecoding/soil/pkg/system"

func (s *Sink) record(o *operation) {
	if !system.InsideDirectory(s.root, o.path) {
		s.dropped = append(s.dropped, o.path)

		return
	}

	if o.target != "" && !system.InsideDirectory(s.root, o.target) {
		s.dropped = append(s.dropped, o.target)

		return
	}

	s.operations = append(s.operations, o)
}

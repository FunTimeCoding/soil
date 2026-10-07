package sink

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/sink_operation"
)

func (s *Sink) record(o *sink_operation.Operation) {
	if !system.InsideDirectory(s.root, o.Path) {
		s.dropped = append(s.dropped, o.Path)

		return
	}

	if o.Target != "" && !system.InsideDirectory(s.root, o.Target) {
		s.dropped = append(s.dropped, o.Target)

		return
	}

	s.operations = append(s.operations, o)
}

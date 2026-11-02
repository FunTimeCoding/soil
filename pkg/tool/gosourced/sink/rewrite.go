package sink

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/sink_operation"
)

func (s *Sink) Rewrite(
	path string,
	expected []byte,
	content []byte,
) {
	if !system.InsideDirectory(s.root, path) {
		s.dropped = append(s.dropped, path)

		return
	}

	o := sink_operation.New(constant.OperationWrite, path)
	o.Expected = expected
	o.Content = content
	s.rewrites = append(s.rewrites, o)
}

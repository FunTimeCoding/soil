package sink

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/sink_operation"
)

func (s *Sink) Rename(
	path string,
	target string,
) {
	o := sink_operation.New(constant.OperationRename, path)
	o.Target = target
	s.record(o)
}

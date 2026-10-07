package sink

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/sink_operation"
)

func (s *Sink) Write(
	path string,
	content []byte,
) {
	o := sink_operation.New(constant.OperationWrite, path)
	o.Content = content
	s.record(o)
}

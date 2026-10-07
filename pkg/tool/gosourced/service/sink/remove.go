package sink

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/sink_operation"
)

func (s *Sink) Remove(path string) {
	s.record(sink_operation.New(constant.OperationRemove, path))
}

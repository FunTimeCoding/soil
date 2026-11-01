package sink

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/sink_operation"
)

func (s *Sink) Renames() []*sink_operation.Operation {
	var result []*sink_operation.Operation

	for _, o := range s.operations {
		if o.Kind == constant.OperationRename {
			result = append(result, o)
		}
	}

	return result
}

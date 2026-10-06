package sink

import "github.com/funtimecoding/soil/pkg/tool/gosourced/constant"

func (s *Sink) Remove(path string) {
	s.record(&operation{kind: constant.OperationRemove, path: path})
}

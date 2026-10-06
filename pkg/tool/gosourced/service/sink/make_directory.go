package sink

import "github.com/funtimecoding/soil/pkg/tool/gosourced/constant"

func (s *Sink) MakeDirectory(path string) {
	s.record(&operation{kind: constant.OperationMakeDirectory, path: path})
}

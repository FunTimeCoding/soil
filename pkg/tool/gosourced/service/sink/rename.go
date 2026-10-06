package sink

import "github.com/funtimecoding/soil/pkg/tool/gosourced/constant"

func (s *Sink) Rename(
	path string,
	target string,
) {
	s.record(
		&operation{kind: constant.OperationRename, path: path, target: target},
	)
}

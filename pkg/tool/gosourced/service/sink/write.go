package sink

import "github.com/funtimecoding/soil/pkg/tool/gosourced/constant"

func (s *Sink) Write(
	path string,
	content []byte,
) {
	s.record(
		&operation{kind: constant.OperationWrite, path: path, content: content},
	)
}

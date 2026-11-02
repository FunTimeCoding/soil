package sink

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/sink_operation"
)

func (s *Sink) Relocate(
	path string,
	target string,
) {
	if !system.InsideDirectory(s.root, path) ||
		!system.InsideDirectory(s.root, target) {
		return
	}

	o := sink_operation.New(constant.OperationRename, path)
	o.Target = target
	s.relocations = append(s.relocations, o)
}

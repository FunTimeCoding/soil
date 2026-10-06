package sink

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"os"
)

func (s *Sink) Commit() error {
	for _, o := range s.operations {
		var e error

		switch o.kind {
		case constant.OperationWrite:
			e = os.WriteFile(o.path, o.content, 0644)
		case constant.OperationRemove:
			e = os.Remove(o.path)
		case constant.OperationMakeDirectory:
			e = os.MkdirAll(o.path, 0755)
		case constant.OperationRename:
			e = os.Rename(o.path, o.target)
		}

		if e != nil {
			return e
		}
	}

	return nil
}

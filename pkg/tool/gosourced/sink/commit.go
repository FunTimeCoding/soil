package sink

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"os"
	"slices"
)

func (s *Sink) Commit() error {
	for _, o := range slices.Concat(s.rewrites, s.operations) {
		var e error

		switch o.Kind {
		case constant.OperationWrite:
			e = os.WriteFile(o.Path, o.Content, 0644)
		case constant.OperationRemove:
			e = os.Remove(o.Path)
		case constant.OperationMakeDirectory:
			e = os.MkdirAll(o.Path, 0755)
		case constant.OperationRename:
			e = os.Rename(o.Path, o.Target)
		}

		if e != nil {
			return e
		}
	}

	return nil
}

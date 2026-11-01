package entry

import "github.com/funtimecoding/soil/pkg/web/constant"

func (e *Entry) carriesFile() bool {
	for _, f := range e.fields {
		if f.Kind() == constant.EntryFileKind {
			return true
		}
	}

	return false
}

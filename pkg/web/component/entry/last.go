package entry

import "github.com/funtimecoding/soil/pkg/web/component/entry/field"

func (e *Entry) last(apply func(*field.Field)) *Entry {
	if len(e.fields) > 0 {
		apply(e.fields[len(e.fields)-1])
	}

	return e
}

package entry

import "github.com/funtimecoding/soil/pkg/web/component/entry/field"

func (e *Entry) Focus() *Entry {
	return e.last(func(f *field.Field) { f.Focus() })
}

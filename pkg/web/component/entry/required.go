package entry

import "github.com/funtimecoding/soil/pkg/web/component/entry/field"

func (e *Entry) Required() *Entry {
	return e.last(func(f *field.Field) { f.Required() })
}

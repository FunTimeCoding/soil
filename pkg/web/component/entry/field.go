package entry

import "github.com/funtimecoding/soil/pkg/web/component/entry/field"

func (e *Entry) Field(f *field.Field) *Entry {
	return e.add(f)
}

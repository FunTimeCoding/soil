package entry

import "github.com/funtimecoding/soil/pkg/web/component/entry/field"

func (e *Entry) add(f *field.Field) *Entry {
	e.fields = append(e.fields, f)

	return e
}

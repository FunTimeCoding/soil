package entry

import "github.com/funtimecoding/soil/pkg/web/component/entry/field"

func (e *Entry) Max(value string) *Entry {
	return e.last(func(f *field.Field) { f.Max(value) })
}

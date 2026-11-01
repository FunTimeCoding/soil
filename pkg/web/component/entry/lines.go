package entry

import "github.com/funtimecoding/soil/pkg/web/component/entry/field"

func (e *Entry) Lines(count string) *Entry {
	return e.last(func(f *field.Field) { f.Lines(count) })
}

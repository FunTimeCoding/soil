package entry

import "github.com/funtimecoding/soil/pkg/web/component/entry/field"

func (e *Entry) Reveal(
	path string,
	target string,
) *Entry {
	return e.last(func(f *field.Field) { f.Reveal(path, target) })
}

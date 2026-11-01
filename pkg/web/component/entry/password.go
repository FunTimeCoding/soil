package entry

import (
	"github.com/funtimecoding/soil/pkg/web/component/entry/field"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func (e *Entry) Password(
	name string,
	label string,
) *Entry {
	return e.add(field.New(constant.EntryPasswordKind, name, label, ""))
}

package entry

import (
	"github.com/funtimecoding/soil/pkg/web/component/entry/field"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func (e *Entry) Mail(
	name string,
	label string,
	value string,
) *Entry {
	return e.add(field.New(constant.EntryMailKind, name, label, value))
}

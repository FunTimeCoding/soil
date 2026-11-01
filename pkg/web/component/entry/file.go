package entry

import (
	"github.com/funtimecoding/soil/pkg/web/component/entry/field"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func (e *Entry) File(
	name string,
	label string,
	accept string,
) *Entry {
	return e.add(
		field.New(constant.EntryFileKind, name, label, "").Accept(accept),
	)
}

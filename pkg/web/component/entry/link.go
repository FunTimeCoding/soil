package entry

import (
	"github.com/funtimecoding/soil/pkg/web/component/entry/field"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func (e *Entry) Link(
	name string,
	label string,
	value string,
) *Entry {
	return e.add(field.New(constant.EntryLinkKind, name, label, value))
}

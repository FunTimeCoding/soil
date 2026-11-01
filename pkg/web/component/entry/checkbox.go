package entry

import (
	"github.com/funtimecoding/soil/pkg/web/component/entry/field"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func (e *Entry) Checkbox(
	name string,
	label string,
	checked bool,
) *Entry {
	return e.add(
		field.New(
			constant.EntryCheckboxKind,
			name,
			label,
			constant.EntryCheckedValue,
		).Checked(checked),
	)
}

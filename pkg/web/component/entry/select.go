package entry

import (
	"github.com/funtimecoding/soil/pkg/web/component/choice"
	"github.com/funtimecoding/soil/pkg/web/component/entry/field"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func (e *Entry) Select(
	name string,
	label string,
	choices []*choice.Choice,
	selected string,
) *Entry {
	return e.add(
		field.New(constant.EntrySelectKind, name, label, selected).
			Choices(choices),
	)
}

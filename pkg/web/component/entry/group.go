package entry

import (
	"github.com/funtimecoding/soil/pkg/web/component/entry/field"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func (e *Entry) Group(legend string) *Entry {
	return e.add(field.New(constant.EntryGroupKind, "", legend, ""))
}

package entry

import (
	"github.com/funtimecoding/soil/pkg/web/component/entry/field"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func (e *Entry) Content(n ...gomponents.Node) *Entry {
	return e.add(field.New(constant.EntryContentKind, "", "", "").Content(n...))
}

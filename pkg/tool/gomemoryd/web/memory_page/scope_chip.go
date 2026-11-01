package memory_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/memory_list"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	"maragu.dev/gomponents"
)

func scopeChip(
	l *memory_list.List,
	label string,
	scope string,
	current bool,
) gomponents.Node {
	result := chip.New(label).Link(ListPath(l.Type, l.Tag, scope, 1))

	if current {
		result.Current()
	}

	return result.Node()
}

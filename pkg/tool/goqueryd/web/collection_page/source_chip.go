package collection_page

import (
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	"maragu.dev/gomponents"
)

func sourceChip(
	label string,
	name string,
	value string,
	current string,
) gomponents.Node {
	result := chip.New(label).Link(Path(name, value, 1))

	if value == current {
		result.Current()
	}

	return result.Node()
}

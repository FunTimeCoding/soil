package relation_page

import (
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	"maragu.dev/gomponents"
)

func filterChip(
	label string,
	value string,
	current string,
) gomponents.Node {
	result := chip.New(label).Link(Path(value))

	if value == current {
		result.Current()
	}

	return result.Node()
}

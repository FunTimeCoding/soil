package summary_page

import (
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	"maragu.dev/gomponents"
)

func choice(
	label string,
	path string,
	current bool,
) gomponents.Node {
	c := chip.New(label).Link(path)

	if current {
		c.Current()
	}

	return c.Node()
}

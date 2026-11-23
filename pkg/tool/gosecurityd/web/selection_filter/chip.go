package selection_filter

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/system_page"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	"maragu.dev/gomponents"
)

func (f *Filter) chip(
	label string,
	count int,
	s *system_selection.Selection,
	current bool,
) gomponents.Node {
	c := chip.New(fmt.Sprintf("%s %d", label, count)).
		Link(system_page.Query(f.path, f.query(s)))

	if current {
		c.Current()
	}

	return c.Node()
}

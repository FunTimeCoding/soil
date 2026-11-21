package system_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	"maragu.dev/gomponents"
)

func facetChip(
	label string,
	count int,
	s *system_selection.Selection,
	current bool,
) gomponents.Node {
	c := chip.New(fmt.Sprintf("%s %d", label, count)).
		Link(Query(constant.SystemsPath, s.Query()))

	if current {
		c.Current()
	}

	return c.Node()
}

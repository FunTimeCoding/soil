package selection_filter

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/facet_count"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"net/url"
)

func New(
	path string,
	s *system_selection.Selection,
	c *facet_count.Count,
) *Filter {
	return &Filter{
		path:      path,
		selection: s,
		count:     c,
		kept:      url.Values{},
		floors:    constant.SystemFloors,
	}
}

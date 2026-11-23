package selection_filter

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/facet_count"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"net/url"
)

type Filter struct {
	path      string
	selection *system_selection.Selection
	count     *facet_count.Count
	kept      url.Values
	fallback  float64
	floors    []float64
}

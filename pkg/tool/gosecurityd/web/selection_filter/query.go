package selection_filter

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"maps"
	"net/url"
)

func (f *Filter) query(s *system_selection.Selection) url.Values {
	result := s.Query(f.fallback)
	maps.Copy(result, f.kept)

	return result
}

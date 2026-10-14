package unclosed_resource

import "go/types"

func (s *Summaries) spentFor(f *types.Func) bool {
	if s == nil || f == nil {
		return false
	}

	return s.spent[functionKey(f)]
}

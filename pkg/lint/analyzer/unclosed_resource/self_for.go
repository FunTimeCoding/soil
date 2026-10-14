package unclosed_resource

import "go/types"

func (s *Summaries) selfFor(f *types.Func) bool {
	if s == nil || f == nil {
		return false
	}

	return s.self[functionKey(f)]
}

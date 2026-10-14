package unclosed_resource

import "go/types"

func (s *Summaries) lendsFor(f *types.Func) bool {
	if s == nil || f == nil {
		return false
	}

	return s.lends[functionKey(f)]
}

package unclosed_resource

import "go/types"

func (s *Summaries) arrangesFor(f *types.Func) bool {
	if s == nil || f == nil {
		return false
	}

	return s.arranges[f]
}

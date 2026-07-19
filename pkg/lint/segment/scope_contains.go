package segment

import "go/types"

func ScopeContains(
	s *types.Scope,
	name string,
) bool {
	return s.Lookup(name) != nil ||
		childScopeContains(s, name) ||
		ParentScopeContains(s, name)
}

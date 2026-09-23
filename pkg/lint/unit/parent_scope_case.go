package unit

import "go/types"

type parentScopeCase struct {
	name   string
	scope  *types.Scope
	target string
	want   bool
}

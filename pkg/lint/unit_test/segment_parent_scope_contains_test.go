package unit_test

import (
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"go/types"
	"testing"
)

func TestParentScopeContains(t *testing.T) {
	universe := types.Universe
	p := types.NewScope(universe, 0, 1, "package")
	function := types.NewScope(p, 0, 1, "function")
	block := types.NewScope(function, 0, 1, "block")
	inner := types.NewScope(block, 0, 1, "inner")
	p.Insert(types.NewVar(0, nil, "p", types.Typ[types.Int]))
	function.Insert(types.NewVar(0, nil, "f", types.Typ[types.Int]))
	block.Insert(types.NewVar(0, nil, "b", types.Typ[types.Int]))

	for _, c := range []struct {
		name   string
		scope  *types.Scope
		target string
		want   bool
	}{
		{"ParentHasName", block, "f", true},
		{"GrandparentHasName", inner, "f", true},
		{"PackageScopeHasName", inner, "p", true},
		{"SameScopeIgnored", function, "f", false},
		{"NoParentHasName", inner, "z", false},
		{"UniverseScopeIgnored", inner, "int", false},
	} {
		t.Run(
			c.name,
			func(t *testing.T) {
				got := segment.ParentScopeContains(c.scope, c.target)

				if got != c.want {
					t.Errorf(
						"parentScopeContains(%s, %q) = %v, want %v",
						c.scope.String(),
						c.target,
						got,
						c.want,
					)
				}
			},
		)
	}
}

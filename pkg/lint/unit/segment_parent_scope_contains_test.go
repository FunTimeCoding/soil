package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"testing"
)

func TestParentScopeContainsParent(t *testing.T) {
	assert.True(t, segment.ParentScopeContains(nestedScope().Parent(), "f"))
}

func TestParentScopeContainsGrandparent(t *testing.T) {
	assert.True(t, segment.ParentScopeContains(nestedScope(), "f"))
}

func TestParentScopeContainsPackageScope(t *testing.T) {
	assert.True(t, segment.ParentScopeContains(nestedScope(), "p"))
}

func TestParentScopeContainsSameScopeIgnored(t *testing.T) {
	assert.False(
		t,
		segment.ParentScopeContains(nestedScope().Parent().Parent(), "f"),
	)
}

func TestParentScopeContainsNoParent(t *testing.T) {
	assert.False(t, segment.ParentScopeContains(nestedScope(), "z"))
}

func TestParentScopeContainsUniverseIgnored(t *testing.T) {
	assert.False(t, segment.ParentScopeContains(nestedScope(), "int"))
}

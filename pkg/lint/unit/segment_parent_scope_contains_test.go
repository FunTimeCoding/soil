package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"github.com/funtimecoding/soil/pkg/lint/unit/lint_tester"
	"testing"
)

func TestParentScopeContainsParent(t *testing.T) {
	assert.True(
		t,
		segment.ParentScopeContains(lint_tester.NestedScope().Parent(), "f"),
	)
}

func TestParentScopeContainsGrandparent(t *testing.T) {
	assert.True(t, segment.ParentScopeContains(lint_tester.NestedScope(), "f"))
}

func TestParentScopeContainsPackageScope(t *testing.T) {
	assert.True(t, segment.ParentScopeContains(lint_tester.NestedScope(), "p"))
}

func TestParentScopeContainsSameScopeIgnored(t *testing.T) {
	assert.False(
		t,
		segment.ParentScopeContains(
			lint_tester.NestedScope().Parent().Parent(),
			"f",
		),
	)
}

func TestParentScopeContainsNoParent(t *testing.T) {
	assert.False(t, segment.ParentScopeContains(lint_tester.NestedScope(), "z"))
}

func TestParentScopeContainsUniverseIgnored(t *testing.T) {
	assert.False(
		t,
		segment.ParentScopeContains(lint_tester.NestedScope(), "int"),
	)
}

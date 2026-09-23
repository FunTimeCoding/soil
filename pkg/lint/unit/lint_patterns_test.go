package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/unit/lint_tester"
	"path/filepath"
	"testing"
)

func TestPatternsNone(t *testing.T) {
	root := lint_tester.ScopeTree(t)
	patterns, e := lint.Patterns(root, root, nil)
	assert.Nil(t, e)
	assert.Strings(t, nil, patterns)
}

func TestPatternsWholeFromRoot(t *testing.T) {
	root := lint_tester.ScopeTree(t)
	patterns, e := lint.Patterns(root, root, []string{"./..."})
	assert.Nil(t, e)
	assert.Strings(t, []string{"./..."}, patterns)
}

func TestPatternsWholeFromSubdirectory(t *testing.T) {
	root := lint_tester.ScopeTree(t)
	patterns, e := lint.Patterns(
		root,
		filepath.Join(root, constant.PackageDirectory, "lint"),
		[]string{"./...", ".", "../lint"},
	)
	assert.Nil(t, e)
	assert.Strings(
		t,
		[]string{"./pkg/lint/...", "./pkg/lint", "./pkg/lint"},
		patterns,
	)
}

func TestPatternsImportPath(t *testing.T) {
	root := lint_tester.ScopeTree(t)
	patterns, e := lint.Patterns(
		root,
		filepath.Join(root, constant.PackageDirectory),
		[]string{"github.com/example/module/...", "std"},
	)
	assert.Nil(t, e)
	assert.Strings(
		t,
		[]string{"github.com/example/module/...", "std"},
		patterns,
	)
}

func TestPatternsMissing(t *testing.T) {
	root := lint_tester.ScopeTree(t)
	_, e := lint.Patterns(root, root, []string{"./ghost/..."})
	assert.Error(t, e)
}

func TestPatternsRefusesFile(t *testing.T) {
	root := lint_tester.ScopeTree(t)
	_, e := lint.Patterns(root, root, []string{"pkg/lint/lint.go"})
	assert.Error(t, e)
	assert.StringContains(t, "package directory pattern", e.Error())
	_, e = lint.Patterns(root, root, []string{"./pkg/lint/lint.go"})
	assert.Error(t, e)
}

func TestPatternsOutside(t *testing.T) {
	root := lint_tester.ScopeTree(t)
	_, e := lint.Patterns(root, root, []string{"../..."})
	assert.Error(t, e)
}

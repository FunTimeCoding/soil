package lint_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
	"testing"
)

func StubbedSoil(t *testing.T) string {
	t.Helper()
	result := testutil.PrepareTestPackage(
		t,
		"testdata/src/example",
		"require github.com/funtimecoding/soil v1.0.0",
		"replace github.com/funtimecoding/soil => ./stub/soil",
	)
	stub := filepath.Join(result, "stub", "soil")
	testutil.CopyTree(t, "testdata/src/github.com/funtimecoding/soil", stub)
	system_tester.Write(
		t,
		stub,
		"go.mod",
		"module github.com/funtimecoding/soil\n\ngo 1.26.3\n",
	)

	return result
}

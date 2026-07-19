package forbidden_import

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/forbidden_import"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"os"
	"path/filepath"
	"testing"
)

func TestBlocked(t *testing.T) {
	temporary := testutil.PrepareTestPackage(
		t,
		"testdata/src/example",
		"require github.com/stretchr/testify v1.0.0",
		"replace github.com/stretchr/testify => ./stub/testify",
	)
	stubDirectory := filepath.Join(temporary, "stub", "testify")
	testutil.CopyTree(
		t,
		"testdata/src/github.com/stretchr/testify",
		stubDirectory,
	)
	e := os.WriteFile(
		filepath.Join(stubDirectory, "go.mod"),
		[]byte("module github.com/stretchr/testify\n\ngo 1.26.3\n"),
		0644,
	)

	if e != nil {
		t.Fatalf("write stub go.mod: %s", e)
	}

	p, results := testutil.LoadFromDirectory(t, temporary)
	forbidden_import.Check(p, results)
	testutil.AssertBlocked(t, results, 2)
	testutil.AssertBlockedContains(t, results, "pflag")
	testutil.AssertBlockedContains(t, results, "testify")
}

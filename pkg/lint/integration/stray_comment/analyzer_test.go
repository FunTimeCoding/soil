package stray_comment

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/stray_comment"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestStrayComment(t *testing.T) {
	directory := testutil.PrepareTestPackage(t, "testdata/src/example")
	configuration := &packages.Config{
		Mode: packages.LoadSyntax | packages.NeedModule,
		Fset: token.NewFileSet(),
		Dir:  directory,
	}
	loaded, e := packages.Load(configuration, "./...")

	if e != nil {
		t.Fatalf("load: %s", e)
	}

	results := output.NewResultsWithDirectory(fmt.Sprintf("%s/", directory))

	for _, p := range loaded {
		if len(p.Errors) > 0 {
			t.Fatalf("package errors: %v", p.Errors)
		}

		stray_comment.Check(p, results)
	}

	testutil.AssertBlocked(t, results, 4)
	testutil.AssertBlockedContains(t, results, "narrates the obvious")
	testutil.AssertBlockedContains(t, results, "pass the token along")
}

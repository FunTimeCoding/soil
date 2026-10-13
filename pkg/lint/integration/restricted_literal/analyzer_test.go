package restricted_literal

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/restricted_literal"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestCheckRules(t *testing.T) {
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

		restricted_literal.CheckRestrictions(
			p,
			results,
			constant.FixtureLiteralRestrictions,
		)
	}

	testutil.AssertBlocked(t, results, 1)
	testutil.AssertBlockedContains(
		t,
		results,
		"build clients in the blessed package",
	)
}

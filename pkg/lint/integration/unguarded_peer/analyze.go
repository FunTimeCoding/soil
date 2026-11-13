package unguarded_peer

import (
	"fmt"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unguarded_peer"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/token"
	"golang.org/x/tools/go/packages"
	"os"
	"path/filepath"
	"testing"
)

func analyze(t *testing.T) *output.Results {
	t.Helper()
	directory := testutil.PrepareTestPackage(
		t,
		"testdata/src/example",
		"require github.com/funtimecoding/soil v1.0.0",
		"replace github.com/funtimecoding/soil => ./stub/soil",
	)
	stub := filepath.Join(directory, "stub", "soil")
	testutil.CopyTree(t, "testdata/src/github.com/funtimecoding/soil", stub)
	e := os.WriteFile(
		filepath.Join(stub, "go.mod"),
		[]byte("module github.com/funtimecoding/soil\n\ngo 1.26.3\n"),
		0644,
	)

	if e != nil {
		t.Fatalf("write stub go.mod: %s", e)
	}

	configuration := &packages.Config{
		Mode:  packages.LoadSyntax | packages.NeedModule,
		Fset:  token.NewFileSet(),
		Dir:   directory,
		Tests: true,
	}
	loaded, f := packages.Load(configuration, goModConstant.AllPackages)

	if f != nil {
		t.Fatalf("load: %s", f)
	}

	result := output.NewResultsWithDirectory(fmt.Sprintf("%s/", directory))

	for _, p := range resolve.PreferTestVariants(loaded) {
		if len(p.Errors) > 0 {
			t.Fatalf("package errors: %v", p.Errors)
		}

		unguarded_peer.Check(p, result)
	}

	return result
}

package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/system/writer"
	"golang.org/x/tools/go/packages"
	"io"
)

func dropBroken(
	loaded []*packages.Package,
	notice io.Writer,
) []*packages.Package {
	broken := map[string]bool{}

	for _, p := range loaded {
		if len(p.Errors) > 0 && !broken[p.PkgPath] {
			broken[p.PkgPath] = true
			writer.Print(notice, "facts skipped: %s does not load\n", p.PkgPath)
		}
	}

	var result []*packages.Package

	for _, p := range loaded {
		if !broken[p.PkgPath] {
			result = append(result, p)
		}
	}

	return result
}

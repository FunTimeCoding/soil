package file_identity

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"path/filepath"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	for _, file := range p.Syntax {
		if ast.IsGenerated(file) {
			continue
		}

		name := filepath.Base(p.Fset.File(file.Pos()).Name())

		if skip(name) {
			continue
		}

		checkFile(p, results, file, name)
	}
}

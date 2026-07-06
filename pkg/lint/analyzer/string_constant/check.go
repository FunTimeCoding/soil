package string_constant

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
	if len(p.Syntax) == 0 {
		return
	}

	directory := filepath.Dir(p.Fset.File(p.Syntax[0].Pos()).Name())

	if filepath.Base(directory) == "constant" {
		return
	}

	constants := collectConstants(directory)

	if len(constants) == 0 {
		return
	}

	for _, file := range p.Syntax {
		if ast.IsGenerated(file) {
			continue
		}

		name := filepath.Base(p.Fset.File(file.Pos()).Name())

		if name == "constant.go" {
			continue
		}

		checkFile(p, results, file, constants)
	}
}

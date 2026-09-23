package file_identity

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"path/filepath"
	"strings"
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

		if strings.HasSuffix(name, constant.TestSuffix) {
			checkTestFile(p, results, file, name)

			continue
		}

		if skip(name) {
			continue
		}

		checkFile(p, results, file, name)
	}
}

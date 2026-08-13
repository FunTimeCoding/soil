package gofix

import (
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"os"
	"path/filepath"
)

func findFormatEdits(
	all []*packages.Package,
	r *output.Results,
	collapse bool,
) map[string]*dst.File {
	changed := make(map[string]*dst.File)
	sourceCache := make(map[string][]byte)

	for _, p := range all {
		generated := make(map[string]bool)

		for _, file := range p.Syntax {
			name := p.Fset.File(file.Pos()).Name()

			if filepath.Base(name) == constant.GeneratedFile {
				generated[name] = true
			}
		}

		dec := decorator.NewDecorator(p.Fset)

		for _, file := range p.Syntax {
			name := p.Fset.File(file.Pos()).Name()

			if generated[name] {
				continue
			}

			if ast.IsGenerated(file) {
				continue
			}

			if !filepath.IsAbs(name) || !fileExists(name) {
				continue
			}

			source, okay := sourceCache[name]

			if !okay {
				var e error
				source, e = os.ReadFile(name)

				if e != nil {
					continue
				}

				sourceCache[name] = source
			}

			destinationFile, e := dec.DecorateFile(file)

			if e != nil {
				continue
			}

			walkFormatEdits(
				destinationFile,
				dec,
				p.Fset,
				source,
				name,
				collapse,
				r,
				changed,
			)
		}
	}

	return changed
}

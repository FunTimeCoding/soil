package index

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/crap/complexity"
	"github.com/funtimecoding/soil/pkg/crap/function"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/ast"
	"strings"
)

func Load(
	root string,
	patterns ...string,
) *Index {
	loaded, set, e := resolve.LoadPackages(root, patterns...)
	errors.PanicOnError(e)
	result := New(root)
	seen := map[string]bool{}

	for _, p := range loaded {
		for _, file := range p.Syntax {
			path := set.File(file.Pos()).Name()

			if seen[path] ||
				strings.HasSuffix(path, constant.TestSuffix) ||
				ast.IsGenerated(file) {
				continue
			}

			seen[path] = true

			for _, d := range file.Decls {
				declaration, okay := d.(*ast.FuncDecl)

				if !okay || declaration.Body == nil {
					continue
				}

				f := function.New(
					p.PkgPath,
					path,
					set.Position(declaration.Pos()).Line,
					declaration.Name.Name,
					complexity.Complexity(declaration),
				)
				f.EndLine = set.Position(declaration.End()).Line
				f.Receiver = receiver(declaration)
				result.Add(f)
			}
		}
	}

	return result
}

package repository

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func declarationNames(directory string) map[string]bool {
	result := map[string]bool{}
	entries, e := os.ReadDir(directory)

	if e != nil {
		return result
	}

	set := token.NewFileSet()

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), constant.GoExtension) {
			continue
		}

		f, g := parser.ParseFile(
			set,
			filepath.Join(directory, entry.Name()),
			nil,
			parser.SkipObjectResolution,
		)

		if g != nil {
			continue
		}

		for _, d := range f.Decls {
			switch declaration := d.(type) {
			case *ast.FuncDecl:
				if declaration.Recv == nil {
					result[declaration.Name.Name] = true
				}
			case *ast.GenDecl:
				for _, s := range declaration.Specs {
					switch spec := s.(type) {
					case *ast.TypeSpec:
						result[spec.Name.Name] = true
					case *ast.ValueSpec:
						for _, n := range spec.Names {
							result[n.Name] = true
						}
					}
				}
			}
		}
	}

	return result
}

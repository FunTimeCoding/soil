package gofix

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func fixFileReferences(
	path string,
	renames []exportedRename,
	r *output.Results,
) {
	content, e := os.ReadFile(path)

	if e != nil {
		return
	}

	set := token.NewFileSet()
	file, e := parser.ParseFile(set, path, content, parser.ParseComments)

	if e != nil {
		return
	}

	renameMap := make(map[string]string)

	for _, rename := range renames {
		renameMap[rename.oldName] = rename.newName
	}

	selectorPositions := collectSelectorPositions(file)
	modified := make([]byte, len(content))
	copy(modified, content)
	offset := 0
	ast.Inspect(
		file,
		func(n ast.Node) bool {
			i, okay := n.(*ast.Ident)

			if !okay {
				return true
			}

			newName, exists := renameMap[i.Name]

			if !exists {
				return true
			}

			if selectorPositions[i.Pos()] {
				return true
			}

			start := set.Position(i.Pos()).Offset + offset
			end := start + len(i.Name)
			modified = append(
				modified[:start],
				append([]byte(newName), modified[end:]...)...,
			)
			offset += len(newName) - len(i.Name)
			r.AddConcern(
				concern.NewFile(
					"renamed",
					fmt.Sprintf(
						"renamed %s → %s (unloaded)",
						i.Name,
						newName,
					),
					path,
					true,
				),
			)

			return true
		},
	)

	if offset != 0 {
		errors.PanicOnError(os.WriteFile(path, modified, 0644))
	}
}

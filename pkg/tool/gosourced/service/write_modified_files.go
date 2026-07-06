package service

import (
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/ast"
	"go/token"
	"sort"
)

func writeModifiedFiles(
	fileSet *token.FileSet,
	references []resolve.Reference,
) ([]string, error) {
	fileMap := make(map[string]*ast.File)

	for _, r := range references {
		position := fileSet.Position(r.Ident.Pos())
		path := position.Filename

		if _, exists := fileMap[path]; exists {
			continue
		}

		for _, file := range r.Package.Syntax {
			filePosition := fileSet.Position(file.Pos())

			if filePosition.Filename == path {
				fileMap[path] = file

				break
			}
		}
	}

	var paths []string

	for path, file := range fileMap {
		e := writeFile(fileSet, file, path)

		if e != nil {
			return nil, e
		}

		paths = append(paths, path)
	}

	sort.Strings(paths)

	return paths, nil
}

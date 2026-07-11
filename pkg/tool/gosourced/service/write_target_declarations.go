package service

import (
	"github.com/funtimecoding/soil/pkg/source/imports"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
)

func writeTargetDeclarations(
	directory string,
	packageName string,
	fileName string,
	declarations []ast.Decl,
	carried []*ast.ImportSpec,
) (string, error) {
	targetPath := filepath.Join(directory, fileName)
	_, e := os.Stat(targetPath)

	if e != nil {
		return targetPath, writeExtractedFile(
			packageName,
			carried,
			declarations,
			targetPath,
		)
	}

	fileSet := token.NewFileSet()
	file, e := parser.ParseFile(
		fileSet,
		targetPath,
		nil,
		parser.ParseComments,
	)

	if e != nil {
		return targetPath, e
	}

	file.Decls = append(file.Decls, declarations...)
	names := importLocalNames(file)

	for _, spec := range carried {
		importPath, f := strconv.Unquote(spec.Path.Value)

		if f != nil {
			continue
		}

		if importPresent(names, importPath) {
			continue
		}

		alias := ""

		if spec.Name != nil {
			alias = spec.Name.Name
		}

		imports.Add(file, importPath, alias)
	}

	return targetPath, writeFile(fileSet, file, targetPath)
}

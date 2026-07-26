package service

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func writeTargetDeclarations(
	directory string,
	packageName string,
	fileName string,
	rendered string,
	carried []*ast.ImportSpec,
) (string, error) {
	targetPath := filepath.Join(directory, fileName)
	_, e := os.Stat(targetPath)

	if e != nil {
		return targetPath, writeCreatedFile(
			packageName,
			carried,
			rendered,
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

	names := importLocalNames(file)
	added := false

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
		added = true
	}

	var content []byte

	if added {
		var buffer bytes.Buffer

		if f := format.Node(&buffer, fileSet, file); f != nil {
			return targetPath, f
		}

		content = buffer.Bytes()
	} else {
		content, e = os.ReadFile(targetPath)

		if e != nil {
			return targetPath, e
		}
	}

	combined := join.Empty(
		strings.TrimRight(string(content), "\n"),
		"\n\n",
		rendered,
		"\n",
	)
	formatted, e := format.Source([]byte(combined))

	if e != nil {
		return targetPath, e
	}

	return targetPath, os.WriteFile(targetPath, formatted, 0644)
}

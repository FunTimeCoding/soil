package repository

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/repository/declaration_set"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func parseDeclarations(directory string) *declaration_set.Set {
	result := declaration_set.New()
	entries, e := os.ReadDir(directory)

	if e != nil {
		return result
	}

	set := token.NewFileSet()

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), soil.GoExtension) {
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
			result.Add(d)
		}
	}

	result.Promote()

	return result
}

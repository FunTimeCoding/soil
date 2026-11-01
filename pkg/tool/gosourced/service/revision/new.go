package revision

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	modConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"go/ast"
	"go/parser"
	"go/token"
	"maps"
	"slices"
	"strings"
)

func New(files map[string][]byte) *Revision {
	result := &Revision{Module: modulePath(files[modConstant.ModFile])}
	s := token.NewFileSet()
	var parsed []*ast.File

	for _, name := range slices.Sorted(maps.Keys(files)) {
		if !strings.HasSuffix(name, soil.GoExtension) {
			continue
		}

		f, e := parser.ParseFile(
			s,
			name,
			files[name],
			parser.SkipObjectResolution,
		)

		if e != nil {
			result.Unparsed = append(result.Unparsed, name)

			continue
		}

		parsed = append(parsed, f)
	}

	clauses := importClauses(s, parsed, result.Module)
	names := packageNames(s, parsed)

	for _, f := range parsed {
		result.Symbols = append(
			result.Symbols,
			fileSymbols(s, f, result.Module, clauses, names)...,
		)
	}

	distinguish(result.Symbols)
	attachText(result.Symbols, files)

	return result
}

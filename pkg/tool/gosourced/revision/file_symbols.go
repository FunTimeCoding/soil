package revision

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"
	"go/ast"
	"go/token"
	"strings"
)

func fileSymbols(
	s *token.FileSet,
	f *ast.File,
	module string,
	clauses map[string]string,
	names map[string]map[string]bool,
) []*symbol.Symbol {
	directory := packageDirectory(s, f)
	own := importPath(module, directory)

	if strings.HasSuffix(f.Name.Name, constant.TestClauseSuffix) {
		own = join.Empty(own, constant.TestClauseSuffix)
	}

	q := &qualifier{
		imports: importNames(f, clauses),
		own:     names[packageKey(directory, f.Name.Name)],
		path:    own,
	}
	var result []*symbol.Symbol

	for _, d := range f.Decls {
		switch d := d.(type) {
		case *ast.FuncDecl:
			result = append(result, functionSymbol(s, f, d, q))
		case *ast.GenDecl:
			result = append(result, declarationSymbols(s, f, d, q)...)
		}
	}

	for _, entry := range result {
		entry.Path = own
	}

	return result
}

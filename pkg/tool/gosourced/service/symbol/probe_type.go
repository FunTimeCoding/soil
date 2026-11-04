package symbol

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/package_importer"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"sort"
	"strings"
)

func probeType(
	source string,
	isType bool,
	qualifiers map[string]string,
	index map[string]*types.Package,
) (types.Type, error) {
	var lines []string

	for name, path := range qualifiers {
		lines = append(lines, fmt.Sprintf("\t%s %q", name, path))
	}

	sort.Strings(lines)
	declaration := fmt.Sprintf("var _ = %s", source)

	if isType {
		declaration = fmt.Sprintf("var _ %s", source)
	}

	content := fmt.Sprintf(
		"package probe\n\nimport (\n%s\n)\n\n%s\n",
		strings.Join(lines, "\n"),
		declaration,
	)
	set := token.NewFileSet()
	file, e := parser.ParseFile(set, "probe.go", content, 0)

	if e != nil {
		return nil, fmt.Errorf("%s does not parse: %w", source, e)
	}

	checked := &types.Info{Types: map[ast.Expr]types.TypeAndValue{}}
	c := types.Config{Importer: package_importer.New(index)}

	if _, f := c.Check("probe", set, []*ast.File{file}, checked); f != nil {
		return nil, fmt.Errorf("%s: %w", source, f)
	}

	spec := file.Decls[len(file.Decls)-1].(*ast.GenDecl).Specs[0].(*ast.ValueSpec)

	if isType {
		return checked.TypeOf(spec.Type), nil
	}

	return checked.TypeOf(spec.Values[0]), nil
}

package pattern

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/match"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func buildReplacement(
	replacement *match.Pattern,
	source string,
	querySymbol string,
	qualifiers map[string]string,
	set *decoration.Set,
	p *packages.Package,
	bindings map[string]ast.Expr,
	anchor ast.Node,
) (dst.Stmt, error) {
	file, e := decorator.Parse(join.Empty("package replacement\n\n", source))

	if e != nil {
		return nil, e
	}

	var statement dst.Stmt

	for _, d := range file.Decls {
		declaration, okay := d.(*dst.FuncDecl)

		if okay && len(declaration.Body.List) == 1 {
			statement = declaration.Body.List[0]
		}
	}

	if statement == nil {
		return nil, fmt.Errorf("replacement holds no statement")
	}

	return substitute(
		statement,
		replacement.Holes,
		querySymbol,
		qualifiers,
		set,
		p,
		bindings,
		anchor,
	).(dst.Stmt), nil
}

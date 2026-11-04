package symbol

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"sort"
)

func evaluate(
	set *token.FileSet,
	owner *packages.Package,
	position token.Pos,
	source string,
	isType bool,
	names map[string][]string,
	explicit map[string]string,
	index map[string]*types.Package,
) (types.Type, map[string]string, error) {
	expression, e := parser.ParseExpr(source)

	if e != nil {
		return nil, nil, fmt.Errorf("%s does not parse: %w", source, e)
	}

	scope := owner.Types.Scope().Innermost(position)

	if scope == nil {
		scope = owner.Types.Scope()
	}

	qualifiers := map[string]string{}
	missing := map[string]string{}
	local := false
	var failure error
	resolve := func(
		name string,
		qualifier bool,
	) {
		if name == "_" || failure != nil {
			return
		}

		_, o := scope.LookupParent(name, position)

		if p, okay := o.(*types.PkgName); okay {
			qualifiers[name] = p.Imported().Path()

			return
		}

		if o != nil {
			if o.Parent() != types.Universe {
				local = true
			}

			return
		}

		path, given := explicit[name]

		if !given && qualifier && len(names[name]) == 1 {
			path, given = names[name][0], true
		}

		if !qualifier || !given {
			failure = fmt.Errorf("%s is not in scope here", name)

			return
		}

		qualifiers[name] = path
		missing[name] = path
	}
	var visit func(n ast.Node) bool
	visit = func(n ast.Node) bool {
		switch v := n.(type) {
		case *ast.SelectorExpr:
			if x, plain := v.X.(*ast.Ident); plain {
				resolve(x.Name, true)
			} else {
				ast.Inspect(v.X, visit)
			}

			return false
		case *ast.KeyValueExpr:
			if _, field := v.Key.(*ast.Ident); field {
				ast.Inspect(v.Value, visit)

				return false
			}
		case *ast.Ident:
			resolve(v.Name, false)
		}

		return true
	}
	ast.Inspect(expression, visit)

	if failure != nil {
		return nil, nil, failure
	}

	if len(missing) == 0 {
		t, f := types.Eval(set, owner.Types, position, source)

		if f != nil {
			return nil, nil, fmt.Errorf("%s: %w", source, f)
		}

		if isType && !t.IsType() {
			return nil, nil, fmt.Errorf("%s is not a type", source)
		}

		if !isType && !t.IsValue() {
			return nil, nil, fmt.Errorf("%s is not a value", source)
		}

		return t.Type, qualifiers, nil
	}

	if local {
		var absent []string

		for name := range missing {
			absent = append(absent, name)
		}

		sort.Strings(absent)

		return nil, nil, fmt.Errorf(
			"%s mixes names from this scope with package %s, which is not imported here",
			source,
			absent[0],
		)
	}

	t, g := probeType(source, isType, qualifiers, index)

	if g != nil {
		return nil, nil, g
	}

	return t, qualifiers, nil
}

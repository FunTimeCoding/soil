package symbol

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/dstutil"
)

func decoratedExpression(
	source string,
	isType bool,
	qualifiers map[string]string,
) (dst.Expr, error) {
	declaration := fmt.Sprintf("var _ = %s", source)

	if isType {
		declaration = fmt.Sprintf("var _ %s", source)
	}

	file, e := decorator.Parse(fmt.Sprintf("package p\n\n%s\n", declaration))

	if e != nil {
		return nil, fmt.Errorf("%s does not parse: %w", source, e)
	}

	spec := file.Decls[0].(*dst.GenDecl).Specs[0].(*dst.ValueSpec)
	result := spec.Type

	if !isType {
		result = spec.Values[0]
	}

	qualified := dstutil.Apply(
		result,
		nil,
		func(c *dstutil.Cursor) bool {
			s, okay := c.Node().(*dst.SelectorExpr)

			if !okay {
				return true
			}

			x, plain := s.X.(*dst.Ident)

			if !plain {
				return true
			}

			path, qualifier := qualifiers[x.Name]

			if !qualifier {
				return true
			}

			c.Replace(&dst.Ident{Name: s.Sel.Name, Path: path})

			return true
		},
	)

	return qualified.(dst.Expr), nil
}

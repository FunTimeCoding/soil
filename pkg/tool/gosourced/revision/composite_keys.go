package revision

import "go/ast"

func compositeKeys(c *ast.CompositeLit) []*ast.Ident {
	var result []*ast.Ident

	for _, element := range c.Elts {
		pair, isPair := element.(*ast.KeyValueExpr)

		if !isPair {
			continue
		}

		if key, isIdent := pair.Key.(*ast.Ident); isIdent {
			result = append(result, key)
		}
	}

	return result
}

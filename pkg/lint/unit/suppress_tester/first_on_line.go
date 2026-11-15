package suppress_tester

import "go/ast"

func firstOnLine[T ast.Node](
	o *Tester,
	line int,
) T {
	o.t.Helper()
	var result T
	found := false
	ast.Inspect(
		o.file,
		func(n ast.Node) bool {
			if found {
				return false
			}

			if v, okay := n.(T); okay &&
				o.fileSet.Position(n.Pos()).Line == line {
				result = v
				found = true
			}

			return true
		},
	)

	if !found {
		o.t.Fatalf("no %T starting on line %d", result, line)
	}

	return result
}

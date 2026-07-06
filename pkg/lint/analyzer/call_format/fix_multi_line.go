package call_format

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/ast"
	"go/token"
)

func fixMultiLine(
	fileSet *token.FileSet,
	call *ast.CallExpr,
	argIndent string,
) []Fix {
	var result []Fix
	openLine := fileSet.Position(call.Lparen).Line
	firstArgLine := fileSet.Position(call.Args[0].Pos()).Line

	if openLine == firstArgLine {
		result = append(
			result,
			Fix{
				Position: call.Lparen + 1,
				End:      call.Args[0].Pos(),
				NewText:  join.Empty("\n", argIndent),
			},
		)
	}

	for i := 1; i < len(call.Args); i++ {
		previousEnd := fileSet.Position(call.Args[i-1].End()).Line
		currentStart := fileSet.Position(call.Args[i].Pos()).Line

		if previousEnd == currentStart {
			result = append(
				result,
				Fix{
					Position: call.Args[i-1].End(),
					End:      call.Args[i].Pos(),
					NewText:  join.Empty(",\n", argIndent),
				},
			)
		}
	}

	return result
}

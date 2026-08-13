package element_format

import (
	"github.com/dave/dst"
	"go/token"
)

func Apply(
	fileSet *token.FileSet,
	e *Elements,
	source []byte,
	items []dst.Expr,
	collapse bool,
) bool {
	openLine := fileSet.Position(e.Open).Line
	closeLine := fileSet.Position(e.Close).Line

	if openLine == closeLine {
		if !exceedsLineLength(fileSet, e, source) {
			return false
		}

		Split(items)

		return true
	}

	if collapse && !e.Ellipsis && !e.HasComments &&
		FitsSingleLine(fileSet, e, source) && !hasLineComments(items) {
		Collapse(items)

		return true
	}

	if IsMultiLineViolation(fileSet, e) {
		Split(items)

		return true
	}

	return false
}

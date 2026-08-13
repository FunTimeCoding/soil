package element_format

import "go/token"

func IsMultiLineViolation(
	fileSet *token.FileSet,
	e *Elements,
) bool {
	openLine := fileSet.Position(e.Open).Line
	firstItemLine := fileSet.Position(e.Items[0].Pos()).Line

	if openLine == firstItemLine {
		return true
	}

	for i := 1; i < len(e.Items); i++ {
		previousEnd := fileSet.Position(e.Items[i-1].End()).Line
		currentStart := fileSet.Position(e.Items[i].Pos()).Line

		if previousEnd == currentStart {
			return true
		}
	}

	return false
}

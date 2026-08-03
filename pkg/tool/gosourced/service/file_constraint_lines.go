package service

import (
	"go/ast"
	"go/build/constraint"
)

func fileConstraintLines(file *ast.File) []string {
	var result []string

	for _, group := range file.Comments {
		if group.Pos() >= file.Package {
			break
		}

		for _, c := range group.List {
			if constraint.IsGoBuild(c.Text) ||
				constraint.IsPlusBuild(c.Text) {
				result = append(result, c.Text)
			}
		}
	}

	return result
}

package service

import (
	"go/ast"
	"strings"
)

func buildConstraints(file *ast.File) []string {
	var result []string

	for _, group := range file.Comments {
		if group.Pos() >= file.Package {
			break
		}

		for _, c := range group.List {
			if strings.HasPrefix(c.Text, "//go:build ") {
				result = append(result, c.Text)
			}
		}
	}

	return result
}

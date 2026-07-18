package service

import (
	"go/ast"
	"go/token"
	"sort"
)

func writeModified(
	s *token.FileSet,
	modified map[string]*ast.File,
) error {
	var filenames []string

	for n := range modified {
		filenames = append(filenames, n)
	}

	sort.Strings(filenames)

	for _, n := range filenames {
		if e := writeFile(s, modified[n], n); e != nil {
			return e
		}
	}

	return nil
}

package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"path/filepath"
)

func renamePackageClauses(
	all []*packages.Package,
	set *token.FileSet,
	sourceDirectory string,
	oldName string,
	modified map[string]*ast.File,
) {
	testName := join.Empty(oldName, "_test")

	for _, loaded := range all {
		for _, file := range loaded.Syntax {
			filename := set.Position(file.Pos()).Filename

			if filepath.Dir(filename) != sourceDirectory {
				continue
			}

			if file.Name.Name == oldName || file.Name.Name == testName {
				modified[filename] = file
			}
		}
	}
}

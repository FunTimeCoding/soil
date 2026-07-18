package service

import (
	"fmt"
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
	newName string,
	modified map[string]*ast.File,
) {
	testName := fmt.Sprintf("%s_test", oldName)

	for _, loaded := range all {
		for _, file := range loaded.Syntax {
			filename := set.Position(file.Pos()).Filename

			if filepath.Dir(filename) != sourceDirectory {
				continue
			}

			if file.Name.Name == oldName {
				file.Name.Name = newName
				modified[filename] = file
			}

			if file.Name.Name == testName {
				file.Name.Name = fmt.Sprintf("%s_test", newName)
				modified[filename] = file
			}
		}
	}
}

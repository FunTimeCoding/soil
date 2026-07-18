package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/ast"
	"go/token"
	"path/filepath"
)

func (s *Service) RenamePackageClause(
	directory string,
	packagePath string,
	newName string,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)

	if !token.IsIdentifier(newName) {
		return failValidation(
			r,
			fmt.Sprintf("not a valid package name: %s", newName),
		)
	}

	all, set, e := resolve.LoadPackages(directory, "./...")

	if e != nil {
		return nil, e
	}

	p := findPackage(all, packagePath)

	if p == nil {
		return failValidation(
			r,
			fmt.Sprintf("package not found: %s", packagePath),
		)
	}

	if len(p.GoFiles) == 0 {
		return failValidation(
			r,
			fmt.Sprintf("package has no Go files: %s", packagePath),
		)
	}

	oldName := p.Types.Name()

	if oldName == newName {
		return failValidation(
			r,
			fmt.Sprintf("package is already named %s", newName),
		)
	}

	qualifiers, taken := collectPackageQualifiers(
		all,
		set,
		packagePath,
		oldName,
		newName,
	)

	if taken != "" {
		return failValidation(r, taken)
	}

	modified := make(map[string]*ast.File)
	renameQualifiers(r, set, qualifiers, oldName, newName, modified)
	renamePackageClauses(
		all,
		set,
		filepath.Dir(p.GoFiles[0]),
		oldName,
		newName,
		modified,
	)

	return r, writeModified(set, modified)
}

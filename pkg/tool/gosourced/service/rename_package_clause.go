package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"go/ast"
	"go/token"
	"path/filepath"
)

func (s *Service) RenamePackageClause(
	directory string,
	packagePath string,
	newName string,
	dryRun bool,
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
	renamePackageClauses(
		all,
		set,
		filepath.Dir(p.GoFiles[0]),
		oldName,
		modified,
	)
	decorations := decoration.NewSet()

	if e := decorateModified(decorations, set, all, modified); e != nil {
		return nil, e
	}

	renameDecoratedClauses(decorations, modified, oldName, newName)
	e = decorateQualifiers(r, decorations, set, qualifiers, oldName, newName)

	if e != nil {
		return nil, e
	}

	names := resolve.NewNames(all)
	names.Override(packagePath, newName)

	if e := restoreDecorations(decorations, names, nil, dryRun); e != nil {
		return nil, e
	}

	if dryRun {
		r.MarkPlanned()
	}

	return r, nil
}

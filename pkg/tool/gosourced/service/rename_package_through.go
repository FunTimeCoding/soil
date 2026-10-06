package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
	"go/ast"
	"go/token"
	"os"
	"path"
	"path/filepath"
)

func (s *Service) renamePackageThrough(
	directory string,
	packagePath string,
	newName string,
	out *sink.Sink,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)

	if !token.IsIdentifier(newName) {
		return failValidation(
			r,
			fmt.Sprintf("not a valid package name: %s", newName),
		)
	}

	all, set, e := s.subtreeLoad(directory, packagePath)

	if e != nil {
		return nil, e
	}

	p := findPackage(all, packagePath)

	if p == nil || len(p.GoFiles) == 0 {
		p = findTestPackage(all, packagePath)
	}

	if p == nil || len(p.GoFiles) == 0 {
		return failValidation(
			r,
			fmt.Sprintf("package not found: %s", packagePath),
		)
	}

	oldName := p.Types.Name()

	if oldName == newName {
		return failValidation(
			r,
			fmt.Sprintf("package is already named %s", newName),
		)
	}

	targetPackagePath := path.Join(path.Dir(packagePath), newName)

	if targetPackagePath == packagePath {
		return failValidation(
			r,
			fmt.Sprintf("import path would not change: %s", packagePath),
		)
	}

	moveDirectory, e := targetDirectory(p, nil, targetPackagePath)

	if e != nil {
		return failValidation(r, e.Error())
	}

	if _, f := os.Stat(moveDirectory); f == nil {
		return failValidation(
			r,
			fmt.Sprintf("target directory already exists: %s", moveDirectory),
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

	sourceDirectory := filepath.Dir(p.GoFiles[0])
	modified := make(map[string]*ast.File)
	renamePackageClauses(all, set, sourceDirectory, oldName, modified)
	decorations := decoration.NewSet()
	e = retargetImports(
		r,
		decorations,
		set,
		all,
		packagePath,
		targetPackagePath,
	)

	if e != nil {
		return nil, e
	}

	if e := decorateModified(decorations, set, all, modified); e != nil {
		return nil, e
	}

	renameDecoratedClauses(decorations, modified, oldName, newName)
	e = decorateQualifiers(r, decorations, set, qualifiers, oldName, newName)

	if e != nil {
		return nil, e
	}

	names := resolve.NewNames(all)
	names.Override(targetPackagePath, newName)

	if e := restoreDecorations(decorations, names, nil, out); e != nil {
		return nil, e
	}

	out.Rename(sourceDirectory, moveDirectory)
	r.AddConcern(
		concern.NewFile(
			"moved",
			fmt.Sprintf("%s → %s", packagePath, targetPackagePath),
			moveDirectory,
			true,
		),
	)

	return r, nil
}

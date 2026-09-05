package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func (s *Service) MovePackage(
	directory string,
	packagePath string,
	targetPackagePath string,
	dryRun bool,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)

	if packagePath == targetPackagePath {
		return failValidation(r, "source and target package are the same")
	}

	if path.Base(packagePath) != path.Base(targetPackagePath) {
		return failValidation(
			r,
			fmt.Sprintf(
				"package name would change from %s to %s - move keeps the name",
				path.Base(packagePath),
				path.Base(targetPackagePath),
			),
		)
	}

	prefix := fmt.Sprintf("%s/", packagePath)

	if strings.HasPrefix(targetPackagePath, prefix) {
		return failValidation(
			r,
			fmt.Sprintf("cannot move %s into itself", packagePath),
		)
	}

	all, set, e := resolve.LoadPackages(directory, "./...")

	if e != nil {
		return nil, e
	}

	p := findPackage(all, packagePath)

	if p == nil || len(p.GoFiles) == 0 {
		p = findTestPackage(all, packagePath)
	}

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

	sourceDirectory := filepath.Dir(p.GoFiles[0])
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

	names := resolve.NewNames(all)
	names.Override(targetPackagePath, p.Types.Name())
	e = restoreDecorations(decorations, names, nil, dryRun)

	if e != nil {
		return nil, e
	}

	if !dryRun {
		e = os.MkdirAll(filepath.Dir(moveDirectory), 0755)

		if e != nil {
			return nil, e
		}

		e = os.Rename(sourceDirectory, moveDirectory)

		if e != nil {
			return nil, e
		}
	}

	r.AddConcern(
		concern.NewFile(
			"moved",
			fmt.Sprintf("%s → %s", packagePath, targetPackagePath),
			moveDirectory,
			true,
		),
	)

	if dryRun {
		r.MarkPlanned()
	}

	return r, nil
}
